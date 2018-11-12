package main

import (
	"log"

	"github.com/nlopes/slack"
)

func main() {
	// API Client を作成する
	api := slack.New("xoxb-ぎゃおーーん")

	// WebSocket で Slack RTM API に接続する
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			// メッセージイベントの場合は、入力メッセージを Log 出力
			log.Printf("Message: %v\n", ev.Text)
			rtm.SendMessage(rtm.NewOutgoingMessage("は？", ev.Channel))
		}
	}
}
