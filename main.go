package main

import (
	"github.com/nlopes/slack"
)

func main() {
	// API Clientを作成する
	api := slack.New("Dummy")

	// WebSocketでSlack RTM APIに接続する
	rtm := api.NewRTM()
	// goroutineで並列化する
	go rtm.ManageConnection()

	// イベントを取得する
	for msg := range rtm.IncomingEvents {
		// 型swtichで型を比較する
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			// MessageEventだったら、「Hello」と応答する
			rtm.SendMessage(rtm.NewOutgoingMessage("は？", ev.Channel))
		}
	}
}
