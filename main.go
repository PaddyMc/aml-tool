package main

import (
	"encoding/base64"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sacOO7/gowebsocket"
	"github.com/tidwall/gjson"
	"time"
)

var (
	SenderKey         = []byte{0x1}
	SenderTotalKey    = []byte{0x2}
	RecipientKey      = []byte{0x3}
	RecipientTotalKey = []byte{0x4}
	SusKey            = []byte{0x5}
)

type AmlProfile struct {
	token  string
	amount int
	txs    int
}

func main() {
	socket := gowebsocket.New("ws://localhost:26657/websocket")
	db := NewDB("./aml")
	defer db.leveldb.Close()

	amlMoneyMule := AmlProfile{
		"seuro",
		5000,
		5,
	}

	socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		result := gjson.Get(message, `result.data.value.TxResult.result.events.#(type=="transfer").attributes`)
		var sender, recip, denom string
		var amount sdk.Int
		for _, json := range result.Array() {
			key, _ := base64.StdEncoding.DecodeString(json.Get("key").String())
			value, _ := base64.StdEncoding.DecodeString(json.Get("value").String())
			switch string(key) {
			case "recipient":
				recip = string(value)
			case "sender":
				sender = string(value)
			case "amount":
				coins, _ := sdk.ParseCoinsNormalized(string(value))
				amount = coins[0].Amount
				denom = coins[0].GetDenom()
			}
		}

		if sender != "" {
			if amlMoneyMule.token == denom {
				//println(sender, recip, denom, amount.String())
				senderTxCount, _ := db.GetData(append(SenderKey, sender...))
				if senderTxCount == "" {
					db.PutData(append(SenderKey, sender...), []byte("1"))
				} else {
					count, _ := strconv.Atoi(senderTxCount)
					count += 1
					if count > amlMoneyMule.txs {
						println("sender: " + "did:cash:" + sender + " is sus, number of txs is more than AML profile suggests")

					}

					db.PutData(append(SenderKey, sender...), []byte(strconv.Itoa(count)))
				}

				senderTotal, _ := db.GetData(append(SenderTotalKey, sender...))
				if senderTotal == "" {
					db.PutData(append(SenderTotalKey, sender...), []byte(amount.String()))
				} else {
					total, _ := strconv.Atoi(senderTotal)
					newTransfer, _ := strconv.Atoi(amount.String())
					total += newTransfer
					if total > amlMoneyMule.amount {
						println("sender: " + "did:cash:" + sender + " is sus, transaction total is more than AML profile suggests")

					}
					db.PutData(append(SenderTotalKey, sender...), []byte(strconv.Itoa(total)))
				}

				recipientTxCount, _ := db.GetData(append(RecipientKey, recip...))
				if recipientTxCount == "" {
					db.PutData(append(RecipientKey, recip...), []byte("1"))
				} else {
					count, _ := strconv.Atoi(recipientTxCount)
					count += 1
					if count > amlMoneyMule.txs {
						println("recipient: " + "did:cash:" + recip + " is sus, number of txs is more than AML profile suggests")

					}
					db.PutData(append(RecipientKey, recip...), []byte(strconv.Itoa(count)))
				}

				recipientTotal, _ := db.GetData(append(RecipientTotalKey, recip...))
				if recipientTotal == "" {
					db.PutData(append(RecipientTotalKey, recip...), []byte(amount.String()))
				} else {
					total, _ := strconv.Atoi(recipientTotal)
					newTransfer, _ := strconv.Atoi(amount.String())
					total += newTransfer
					if total > amlMoneyMule.amount {
						println("recipient: " + "did:cash:" + recip + " is sus, transaction total is more than AML profile suggests")

					}
					db.PutData(append(RecipientKey, recip...), []byte(strconv.Itoa(total)))
				}
			}

		}
	}

	socket.Connect()

	socket.SendText(`{ "jsonrpc": "2.0", "method": "subscribe", "params": ["tm.event='Tx'"], "id": 1 }`)

	for {
		time.Sleep(time.Second)
	}
}
