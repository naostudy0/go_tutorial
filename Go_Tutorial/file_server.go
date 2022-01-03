package main

import (
    "log"
    "net/http"
    "text/template"
)

// テストページの表示
func testHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("public/test.html")
    if err != nil {
        panic(err.Error())
    }
    if err := t.Execute(w, nil); err != nil {
        panic(err.Error())
    }
}

func main() {
    // 公開ディレクトリ
    fs := http.FileServer(http.Dir("public"))
    // URIとの対応
    http.Handle("/", fs)
    http.HandleFunc("/test", testHandler)
    // サーバー起動（致命的なエラーが発生した場合はターミナルへ出力）
    log.Fatal(http.ListenAndServe(":8888", nil))
}
