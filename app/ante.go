package app

import (
	"github.com/Fairblock/fairyring/blockbuster"
	pepante "github.com/Fairblock/fairyring/x/pep/ante"
	pepkeeper "github.com/Fairblock/fairyring/x/pep/keeper"

	storetypes "cosmossdk.io/store/types"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
)

type FairyringHandlerOptions struct {
	BaseOptions       ante.HandlerOptions
	wasmConfig        wasmtypes.WasmConfig
	txCounterStoreKey storetypes.StoreKey
	Mempool           blockbuster.Mempool
	KeyShareLane      pepante.KeyShareLane
	TxDecoder         sdk.TxDecoder
	TxEncoder         sdk.TxEncoder
	PepKeeper         pepkeeper.Keeper
}

// NewFairyringAnteHandler wraps all of the default Cosmos SDK AnteDecorators with the Fairyring AnteHandler.
func NewFairyringAnteHandler(options FairyringHandlerOptions) sdk.AnteHandler {
	if options.BaseOptions.AccountKeeper == nil {
		panic("account keeper is required for ante builder")
	}

	if options.BaseOptions.BankKeeper == nil {
		panic("bank keeper is required for ante builder")
	}

	if options.BaseOptions.SignModeHandler == nil {
		panic("sign mode handler is required for ante builder")
	}

	if options.BaseOptions.FeegrantKeeper == nil {
		panic("fee grant keeper is required for ante builder")
	}

	anteDecorators := []sdk.AnteDecorator{
		ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		wasmkeeper.NewLimitSimulationGasDecorator(options.wasmConfig.SimulationGasLimit),
		wasmkeeper.NewCountTXDecorator(options.txCounterStoreKey),
		ante.NewExtensionOptionsDecorator(options.BaseOptions.ExtensionOptionChecker),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(options.BaseOptions.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.BaseOptions.AccountKeeper),
		ante.NewDeductFeeDecorator(
			options.BaseOptions.AccountKeeper,
			options.BaseOptions.BankKeeper,
			options.BaseOptions.FeegrantKeeper,
			options.BaseOptions.TxFeeChecker,
		),
		ante.NewSetPubKeyDecorator(options.BaseOptions.AccountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(options.BaseOptions.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.BaseOptions.AccountKeeper, options.BaseOptions.SigGasConsumer),
		ante.NewSigVerificationDecorator(options.BaseOptions.AccountKeeper, options.BaseOptions.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.BaseOptions.AccountKeeper),
		pepante.NewPepDecorator(options.PepKeeper, options.TxEncoder, options.KeyShareLane, options.Mempool),
	}

	return sdk.ChainAnteDecorators(anteDecorators...)
}
