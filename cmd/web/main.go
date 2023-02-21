package main

import (
	"log"
	"net/http"
)

type application struct{}

func main() {
	// アプリの設定を行う
	app := application{}

	// アプリケーションルートを取得
	mux := app.routes()

	// メッセージの出力
	log.Println("Starting server on port 8080...")

	// サーバーの起動
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
