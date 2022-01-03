package main

import (
    "net/http"
)

func main() {
    // 公開ディレクトリ
    fs := http.FileServer(http.Dir("public"))
    // URIとの対応
    http.Handle("/", fs)
    // サーバー起動
    http.ListenAndServe(":8888", nil)
}
