# Setting up a validator

## Installation Steps

> Prerequisite: go1.18+ required [ref](https://golang.org/doc/install)

```shell
sudo snap install --classic go
```

> Prerequisite: git [ref](https://github.com/git/git)

```shell
sudo apt install -y git gcc make
```

> Prerequisite: Set environment variables

```shell
sudo nano $HOME/.profile
# Add the following two lines at the end of the file
GOPATH=$HOME/go
PATH=$GOPATH/bin:$PATH
# Save the file and exit the editor
source $HOME/.profile
# Now you should be able to see your variables like this:
echo $GOPATH
/home/[your_username]/go
echo $PATH
/home/[your_username]/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin
```

> Recommended requirement: Increase 'number of open files' limit

```shell
sudo nano /etc/security/limits.conf
# Before the end of the file, add:
[your_username] soft nofile 4096
# Then reboot the instance for it to take effect and check with:
ulimit -Sn
```

> Optional requirement: GNU make [ref](https://www.gnu.org/software/make/manual/html_node/index.html)

- Clone git repository and Network

```shell
git clone https://github.com/FairBlock/fairyring
git clone https://github.com/FairBlock/networks
```

- Checkout release tag

```shell
cd fairyring
git fetch --tags
git checkout [tag_no]
```

- Install

```shell
cd fairyring
go mod tidy
make install
```

### Generate keys

`fairyringd keys add [key_name]`

or

`fairyringd keys add [key_name] --recover` to regenerate keys with your [BIP39](https://github.com/bitcoin/bips/tree/master/bip-0039) mnemonic

---

## Node start and validator setup

### Prior to genesis creation and network launch

- [Install](#installation-steps) FairyRing application
- Initialize node

```shell
fairyringd init {{NODE_NAME}} --chain-id {{CHAIN_ID}}
```

- Create a new key

```shell
fairyringd keys add <keyName>
```

- Add a genesis account with `100000000000stake tokens`

```shell
fairyringd add-genesis-account {{KEY_NAME}} 100000000000stake
```

- Make a genesis transaction to become a validator

```shell
fairyringd gentx \
  [account_key_name] \
  100000000000stake \
  --commission-max-change-rate 0.01 \
  --commission-max-rate 0.2 \
  --commission-rate 0.05 \
  --min-self-delegation 1 \
  --details [optional-details] \
  --identity [optional-identity] \
  --security-contact "[optional-security@example.com]" \
  --website [optional.web.page.com] \
  --moniker [node_moniker] \
  --chain-id fairytest-2
```

- Copy the contents of `${HOME}/.fairyring/config/gentx/gentx-XXXXXXXX.json`
- Clone the [fairyring repository](https://github.com/FairBlock/fairyring) and create a new branch
- Create a file `gentx-{{VALIDATOR_NAME}}.json` under the `networks/testnets/fairytest-2/gentxs` folder in the newly created branch, paste the copied text into the file (note: find reference file `gentx-examplexxxxxxxx.json` in the same folder)
- Run `fairyringd tendermint show-node-id` and copy your nodeID
- Run `ifconfig` or `curl ipinfo.io/ip` and copy your publicly reachable IP address
- Create a file `peers-{{VALIDATOR_NAME}}.json` under the `networks/testnets/fairytest-2/peers` folder in the new branch, paste the copied text from the last two steps into the file (note: find reference file `peers-examplexxxxxxxx.json` in the same folder)
- Create a Pull Request to the `main` branch of the [fairyring repository](https://github.com/FairBlock/fairyring)
  > **NOTE:** the Pull Request will be merged by the maintainers to confirm the inclusion of the validator at the genesis. The final genesis file will be published under the file `networks/testnets/fairytest-2/genesis.json`.
- Once the submission process has closed and the genesis file has been created, replace the contents of your `${HOME}/.fairyring/config/genesis.json` with that of `networks/testnets/fairytest-2/genesis.json`
- Add the required `persistent_peers` or `seeds` in `${HOME}/.fairyring/config/config.toml` from `networks/testnets/fairytest-2/peers-nodes.txt`
- Update the `~/.fairyring/config/config.toml` file and make sure that the peer-port (26656) is publicly reachable. If you are running sentry nodes, make sure that the peer-port of your sentry is publicly reachable
- Start node

```shell
fairyringd start
```

## Genesis Time

The genesis transactions should be sent before [date and time] and the same will be used to publish the `genesis.json` at [date and time]

<!-- > Submitting Gentx is now closed. Genesis has been published and block generation has started -->

---

## After genesis creation and network launch

### Step 1: Start a full node

- [Install](#installation-steps) Fairyring application
- Initialize node

```shell
fairyringd init {{NODE_NAME}}
```

- Replace the contents of your `${HOME}/.fairyring/config/genesis.json` with that of `fairytest-2/genesis.json` from the `master` branch of [network repository](https://github.com/FairBlock/networks)

```shell
curl https://github.com/FairBlock/blob/master/networks/fairyring-testnet-1/genesis.json > $HOME/.fairyring/config/genesis.json
```

- Add `persistent_peers` or `seeds` in `${HOME}/.fairyring/config/config.toml` from `fairytest-2/peers.txt` from the `master` branch of [network repository](https://github.com/Fairblock/fairyring/blob/main/networks/testnets/fairytest-2/peers-nodes.txt)
- Start node

```shell
fairyringd start
```

> Note: if you are only planning to run a full node, you can stop here

### Step 2: Create a validator

- Acquire stake tokens from the (coming soon)
- Wait for your full node to catch up to the latest block (compare to the (coming soon))
- Run `fairyringd tendermint show-validator` and copy your consensus public key
- Send a create-validator transaction

```shell
fairyringd tx staking create-validator \
  --amount 500000000stake \
  --commission-max-change-rate 0.01 \
  --commission-max-rate 0.2 \
  --commission-rate 0.1 \
  --from [account_key_name] \
  --fees 400000stake \
  --min-self-delegation 1 \
  --moniker [validator_moniker] \
  --pubkey $(fairyringd tendermint show-validator) \
  --chain-id fairytest-2 \
  -y
```

---

## Persistent Peers

The `persistent_peers` needs a comma-separated list of trusted peers on the network, you can acquire it from the [peers.txt](https://github.com/Fairblock/fairyring/blob/main/networks/testnets/fairytest-2/peers-nodes.txt) for example:

```shell
cafe8a3e08658fb0309c5ad7017297427ea19ebb@195.14.6.182:26656,525d605dc4cc1e92d6b1d9a6934b19066083e610@34.66.108.187:26656,cd1cbf64a3e85d511c2a40b9e3e7b2e9b40d5905@35.74.28.144:26656,24c65c37d4b4c4cd2688e28a8ea38c377dd0d7f6@65.21.163.231:26656,3cda3bebf7aaeeb0533734496158420dcd3da4ad@94.130.137.119:26666,f842253c4971e898247e054b5dd9c024503f593c@89.58.32.218:27675,6d394ad537476eea0b000e09a786a0400388fc2b@34.30.193.253:26656
```

## Version

This chain is currently running on fairyring [v0.2.1](https://github.com/FairBlock/fairyring/releases/tag/v0.2.1)
Commit Hash: [8e6f6deea6a04b260d190fcd5787bcc4ff85f149](https://github.com/FairBlock/fairyring/commit/8e6f6deea6a04b260d190fcd5787bcc4ff85f149)

## Binary

The binary can be downloaded from [here](https://github.com/FairBlock/fairyring/releases/tag/v0.2.1)

## Explorer

Coming soon!

## Faucet

Coming soon!

### Documentation

Coming soon!

### RPC & API

Coming soon!
