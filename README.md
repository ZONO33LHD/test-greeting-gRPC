# gRPC サーバーとクライアントの起動手順

| ステップ | 説明 | コマンド |
|----------|------|----------|
| 1        | `grpc-server` ディレクトリに移動 | `cd backend/grpc-server` |
| 2        | gRPC サーバーを起動 | `go run main.go` |
| 3        | 新しいターミナルを開き、`grpc-client` ディレクトリに移動 | `cd backend/grpc-client` |
| 4        | gRPC クライアントを起動 | `go run main.go` |
| 5        | (オプション) クライアントに名前を指定してリクエストを送信 | `go run main.go YourName` |
