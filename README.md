# DemoGoHttpServer

This is 学習用ソースコード with Golang


## 実行方法

```sh
$ go run main.go
```

## レスポンス

Webブラウザから`http://localhost:8080` にアクセスすれば確認できる。

### JSONパース例

`http://localhost:8080/api/clock`にリクエストすると現在時刻をJSON形式で返す。

```JSON
{
	"current_time": "2019-05-26T14:29:48.904973+09:00"
}
```
