package main

import (
	// JSONを扱う
	"encoding/json"
	// テンプレート
	"html/template"
	// ログ出力
	"log"
	// サーバー標準パッケージ
	"net/http"
	// 時刻
	"time"
)

func main() {
	// Web用のハンドラーを登録
	// http.HandleFunc("/clock/", clockHandler)
	// API用のを登録する
	http.HandleFunc("/api/clock/", apiClockHandler)

	// 指定されたパス (/static/) に対してハンドラーを登録します
	http.Handle("/static/",
		// URLパスから接頭辞 (/static/) を取り除くミドルウェア的なハンドラーを生成します
		http.StripPrefix("/static/",
			// URLパスを元に静的ファイルを探して返すハンドラーを生成します
			http.FileServer(
				// 静的ファイルを配置したディレクトリ (/etc/gohttpserver/doc) を読み込みます
				// http.Dir("/etc/gohttpserver/doc"),
				http.Dir("./doc"),
			),
		),
	)
	// ListenAndServe は何かに失敗した時にエラーを返すので、それを記録します。
	// エラーが起きなければ ListenAndServe はブロッキングし続けるので log.Fatal は呼ばれません。
	// ここの log はデフォルトで標準出力に書き出されるので、内容はターミナルにそのまま表示されます
	log.Fatal(
		// 指定されたアドレス (:8080) でHTTPサーバーを起動します
		// 2つめの引数で nil を渡すことで、上の http.Handle で登録したハンドラーが使われるようになります
		http.ListenAndServe(":8080", nil),
	)
}

/// 現在時刻をJSONエンコードかけて返す
func apiClockHandler(w http.ResponseWriter, r *http.Request) {
	// JSONにする構造体
	type ResponseBody struct {
		Time time.Time `json:"current_time"`
	}
	rb := &ResponseBody{
		Time: time.Now(),
	}

	// ヘッダーをセット
	w.Header().Set("Content-type", "application/json")

	// JSONにエンコードしてレスポンスに書き込む
	if err := json.NewEncoder(w).Encode(rb); err != nil {
		log.Fatal(err)
	}
}

/// 現在時刻をHTML表示する
func clockHandler(w http.ResponseWriter, r *http.Request) {
	// テンプレートをパース
	// t := template.Must(template.ParseFiles("/etc/gohttpserver/templates/clock.html.tpl"))
	t := template.Must(template.ParseFiles("./templates/clock.html.tpl"))

	// テンプレートを描画
	if err := t.ExecuteTemplate(w, "clock.html.tpl", time.Now()); err != nil {
		log.Fatal(err)
	}
}

// func clockHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, `
//         <!DOCTYPE html>
//         <html>
//         <body>
//             It's %d o'clock now.
//         </body>
//         </html>
//     `, time.Now().Hour())
// }
