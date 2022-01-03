package main

import (
    "log"
    "net/http"
)

func main() {
    // 公開ディレクトリ
    fs := http.FileServer(http.Dir("public"))
    // URIとの対応
    http.Handle("/", fs)
    // サーバー起動（致命的なエラーが発生した場合はターミナルへ出力）
    log.Fatal(http.ListenAndServe(":8888", nil))
}
