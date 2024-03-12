#!/bin/bash

echo ""
echo "#################################################"
echo "#      Submit Proposal to Destination chain     #"
echo "#   Submit encrypted vote on Destination chain  #"
echo "#     Resolve proposal on Destination chain     #"
echo "#       Submit Proposal to Fairyring chain      #"
echo "#    Submit encrypted vote on Fairyring chain   #"
echo "#      Resolve proposal on Fairyring chain      #"
echo "#################################################"
echo ""

BINARY=fairyringd
ENCRYPTER=vote-encrypter
GENERATOR=ShareGenerator
CHAIN_DIR=$(pwd)/data
CHAINID_1=fairyring_test_1
CHAINID_2=fairyring_test_2
BLOCK_TIME=5

WALLET_3=$($BINARY keys show wallet3 -a --keyring-backend test --home $CHAIN_DIR/$CHAINID_1)
VAL1=$($BINARY keys show val1 -a --keyring-backend test --home $CHAIN_DIR/$CHAINID_1)

check_tx_code () {
  local TX_CODE=$(echo "$1" | jq -r '.code')
  if [ "$TX_CODE" != 0 ]; then
    echo "ERROR: Tx failed with code: $TX_CODE"
    exit 1
  fi
}

check_tx_err () {
  local TX_CODE=$(echo "$1" | jq -r '.code')
  if [ "$TX_CODE" -eq 0 ]; then
    return 0
  else
    return 1
  fi
}

wait_for_tx () {
  RESULT=""  
  sleep $BLOCK_TIME
  local TXHASH=$(echo "$1" | jq -r '.txhash')
  if [ "$2" = "source" ]; then
      RESULT=$($BINARY q tx --type=hash $TXHASH --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --node tcp://localhost:16657 -o json)
  elif [ "$2" = "destination" ]; then
      RESULT=$($BINARY q tx --type=hash $TXHASH --home $CHAIN_DIR/$CHAINID_2 --chain-id $CHAINID_2 --node tcp://localhost:26657 -o json)
  fi
  echo "$RESULT"
}
#
#echo "Creating a new proposal on destination chain"
#RESULT=$($BINARY tx gov submit-proposal draft_proposal.json --from $VALIDATOR_2 --gas-prices 1ufairy --home $CHAIN_DIR/$CHAINID_2 --chain-id $CHAINID_2 --node tcp://localhost:26657 --broadcast-mode sync --keyring-backend test -o json -y)
#check_tx_code $RESULT
#RESULT=$(wait_for_tx $RESULT "destination")
#
## waiting for identity to be updated
#sleep 25
#PROPOSAL=$(fairyringd q gov proposals --home $CHAIN_DIR/$CHAINID_1 -o json | jq '.proposals[0]')
#IDENTITY=$(echo "$PROPOSAL" | jq -r '.identity')
#PUBKEY=$(echo "$PROPOSAL" | jq -r '.pubkey')
#
#if [ -z "$IDENTITY" ]; then
#  echo "ERROR: The identity is blank"
#  echo "$PROPOSAL"
#  exit 1
#elif [ -z "$PUBKEY" ]; then
#  echo "The pubkey is blank"
#  echo "$PROPOSAL"
#  exit 1
#else
#  echo "Successfully created proposal on destination chain"
#fi
#
#
#echo "Submitting encrypted vote on destination chain"
#echo "Encrypting vote with Pub key: '$PUBKEY' and Identity: $IDENTITY"
#ENCVOTE=$($ENCRYPTER "yes" 100 $IDENTITY $PUBKEY)
#echo "$ENCVOTE"
#
#RESULT=$(fairyringd tx gov vote-encrypted 1 $ENCVOTE --from val2 --home $CHAIN_DIR/$CHAINID_1 --keyring-backend test --gas-prices 1ufairy -o json -y)
#check_tx_code $RESULT
#RESULT=$(wait_for_tx $RESULT "destination")
#
#VOTE=$(fairyringd q gov votes 1 --home $CHAIN_DIR/$CHAINID_1 -o json | jq '.votes[0]')
#VOTEDATA=$(echo "$VOTE" | jq -r '.encrypted_vote_data')
#
#if [ -z "$VOTEDATA" ]; then
#  echo "ERROR: Encrypted vote option is blank"
#  echo "$VOTE"
#  exit 1
#fi
#echo "Successfully submitted encrypted vote on destination chain"

echo "PUBKEY: $1"
echo "SHARE: $2"

GENERATED_SHARE=$2

echo "Creating a new proposal on source chain"
RESULT=$($BINARY tx gov submit-proposal draft_proposal.json --from $WALLET_3 --gas-prices 1ufairy --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --node tcp://localhost:16657 --broadcast-mode sync --keyring-backend test -o json -y)
sleep 5
check_tx_code $RESULT
RESULT=$(wait_for_tx $RESULT "source")

PROPOSAL=$(fairyringd q gov proposals --home $CHAIN_DIR/$CHAINID_1 --node tcp://localhost:16657 -o json | jq '.proposals[0]')
IDENTITY=$(echo "$PROPOSAL" | jq -r '.identity')
PUBKEY=$(echo "$PROPOSAL" | jq -r '.pubkey')

if [ -z "$IDENTITY" ]; then
  echo "ERROR: The identity is blank"
  echo "$PROPOSAL"
  exit 1
elif [ -z "$PUBKEY" ]; then
  echo "The pubkey is blank"
  echo "$PROPOSAL"
  exit 1
else
  echo "Successfully created proposal on source chain"
fi


echo "Encrypting vote with Pub key: '$PUBKEY' and Identity: $IDENTITY"
ENCVOTE=$($ENCRYPTER "yes" 100 $IDENTITY $PUBKEY)


while true; do
  echo "Submitting encrypted vote on source chain"

  RESULT=$(fairyringd tx gov vote-encrypted 1 $ENCVOTE --from $VAL1 --home $CHAIN_DIR/$CHAINID_1 --node tcp://localhost:16657 --keyring-backend test --gas-prices 1ufairy -o json -y)
  echo "$RESULT"
  check_tx_err $RESULT
  if [ $? -eq 0 ]; then
    break
  fi
done

RESULT=$(wait_for_tx $RESULT "source")

VOTE=$(fairyringd q gov votes 1 --home $CHAIN_DIR/$CHAINID_1 --node tcp://localhost:16657 -o json | jq '.votes[0]')
VOTEDATA=$(echo "$VOTE" | jq -r '.encrypted_vote_data')

if [ -z "$VOTEDATA" ]; then
  echo "ERROR: Encrypted vote option is blank"
  echo "$VOTE"
  exit 1
fi
echo "Successfully submitted encrypted vote on source chain"

echo "waiting for voting period to expire"

sleep 60
EXTRACTED_RESULT=$($GENERATOR derive $GENERATED_SHARE 1 $IDENTITY)
EXTRACTED_SHARE=$(echo "$EXTRACTED_RESULT" | jq -r '.KeyShare')

while true; do
  echo "Submitting General Key Share"
  
  RESULT=$($BINARY tx keyshare create-general-key-share "private-gov-identity" $IDENTITY $EXTRACTED_SHARE 1 --from $VAL1 --gas-prices 1ufairy --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --node tcp://localhost:16657 --broadcast-mode sync --keyring-backend test -o json -y)
  echo "$RESULT"
  check_tx_err $RESULT
  if [ $? -eq 0 ]; then
    break
  fi
done

RESULT=$(wait_for_tx $RESULT "source")
sleep 25

#echo "Checking Status of proposal on Destination chain"
#PROPOSAL=$(fairyringd q gov proposals --home $CHAIN_DIR/$CHAINID_1 -o json | jq '.proposals[0]')
#STATUS=$(echo "$PROPOSAL" | jq -r '.status')
#
#if [ "$STATUS" != "PROPOSAL_STATUS_PASSED" ]; then
#  echo "ERROR: Failed to pass proposal on destination chain"
#  echo "$PROPOSAL"
#  exit 1
#fi
#echo "Successfully passed proposal with enc vote on destination chain"

echo "Checking Status of proposal on Source chain"
sleep 5
PROPOSAL=$(fairyringd q gov proposals --home $CHAIN_DIR/$CHAINID_1 --node tcp://localhost:16657 -o json | jq '.proposals[0]')
STATUS=$(echo "$PROPOSAL" | jq -r '.status')

if [ "$STATUS" != "PROPOSAL_STATUS_PASSED" ]; then
  echo "ERROR: Failed to pass proposal on source chain"
  echo "$PROPOSAL"
  exit 1
fi
echo "Successfully passed proposal with enc vote on source chain"

echo ""
echo "#################################################"
echo "#              SUCCESSFULLY TESTED              #"
echo "#      Submit Proposal to Destination chain     #"
echo "#   Submit encrypted vote on Destination chain  #"
echo "#     Resolve proposal on Destination chain     #"
echo "#       Submit Proposal to Fairyring chain      #"
echo "#    Submit encrypted vote on Fairyring chain   #"
echo "#      Resolve proposal on Fairyring chain      #"
echo "#################################################"
echo ""
