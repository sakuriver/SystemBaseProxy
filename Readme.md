# 本リポジトリの立ち位置

クライアント　→　プロキシサーバー（本リポジトリ　ー継承してアプリ専用のURLを作る）　→　フロント専用ロジック（BFF）　⇄　バックエンド共通機能（統合システムの領域）　⇄　闇DB系

## 実装内容

基本的なAPI提供
　httpリクエスト向け（REST一式）
      webviewer向け（HTMLビューワー）
      アプリ用JSon関連
　    バイナリ機械用通信（ProtocolBUffer、GRPC、RabbitMQ)
　　　　他、各種通信機構が増えたら随時追加


## api一覧

### サーバーの起動コマンド

### json用のget api

url
curl localhost:8080/masterdata/version
レスポンス内容
{"id":1,"title":"0.01","created_date":"2023-05-21T12:30:56.190983+09:00","updated_date":"2023-05-21T12:30:56.190983+09:00"}

### json用のpost api
url 
curl -X post localhost:8080/playdata/upload

レスポンス内容

{"result":true}

                      
