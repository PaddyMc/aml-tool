### aml-tool

This tool is a POC for the aml tracking for the cosmos cash protocol

[comsmos-cash protocol](https://github.com/allinbits/cosmos-cash)

### How to run

Clone the [comsmos-cash protocol](https://github.com/allinbits/cosmos-cash)

In the root directory of the `cosmos-cash` protocol run:

```sh
make start-dev
make seed
```

This will set up some users and an issuer to track aml

Then in the `aml-tool` repo directory root run:

```sh
go run main.go db.go
```

Return to the `cosmos-cash` directory root and run:

```sh
./scripts/seeds/05_money_mule.sh
```

The output in the `aml-tool` should look as follows

```sh
sender: did:cash:cosmos1x77pefahg0ts8vcd45u4cug769mrqcwd0z7ewr is sus, transaction total is more than AML profile suggests
sender: did:cash:cosmos1putuftmjg0vpyh6l7gd48r6pzca3wpayzl575v is sus, number of txs is more than AML profile suggests
sender: did:cash:cosmos1putuftmjg0vpyh6l7gd48r6pzca3wpayzl575v is sus, number of txs is more than AML profile suggests
sender: did:cash:cosmos1putuftmjg0vpyh6l7gd48r6pzca3wpayzl575v is sus, number of txs is more than AML profile suggests
sender: did:cash:cosmos1putuftmjg0vpyh6l7gd48r6pzca3wpayzl575v is sus, number of txs is more than AML profile suggests
sender: did:cash:cosmos1putuftmjg0vpyh6l7gd48r6pzca3wpayzl575v is sus, number of txs is more than AML profile suggests
recipient: did:cash:cosmos1qcdx9qxz780uytqn6y0fzk0aett0a8wys00kmn is sus, number of txs is more than AML profile suggests
recipient: did:cash:cosmos1qcdx9qxz780uytqn6y0fzk0aett0a8wys00kmn is sus, number of txs is more than AML profile suggests
```

### How to investigate addresses

Query the did document in cosmos-cash

```sh
cosmos-cashd query identifier identifier did:cash:cosmos1qcdx9qxz780uytqn6y0fzk0aett0a8wys00kmn --output json | jq
```

Query the verifiable credential in cosmos cash

```sh
cosmos-cashd query verifiablecredentialservice verifiable-credential origin-cred-2 --output json | jq
```
