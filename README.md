# Simple Request API

Goで実装したシンプルなREST APIです。

本リポジトリは、DB実装に依存しない設計を意識し、
永続化層を interface として切り出した構成を採用しています。

⸻

概要
•Go（標準ライブラリのみ）で構成
•net/http によるシンプルな API
•永続化層を interface として定義
•in-memory 実装により DB なしで動作可能

⸻

ディレクトリ構成
``` bash
.
├── cmd/
│   └── api/
│       └── main.go        // エントリーポイント
├── internal/
│   ├── handler/           // HTTPリクエストの受付
│   ├── service/           // ビジネスロジック
│   ├── repository/        // 永続化層（interface + 実装）
│   └── model/             // ドメインモデル
└── README.md
```

⸻

設計方針

レイヤード構成
``` bash
handler → service → repository(interface)
```

•handler
    •HTTPリクエスト / レスポンスの責務
    •URL パラメータの解釈
•service
    •業務ロジックの責務
    •永続化の詳細は知らない
•repository
    •データ取得・保存の抽象化
    •interface として定義

⸻

Repository の設計

永続化層は interface として定義しています。

``` go
type RequestRepository interface {
	FindAll() ([]model.Request, error)
	FindByID(id int64) (*model.Request, error)
}
```

これにより、DB 実装が未完成な段階でも
アプリケーション全体を動かすことができます。

現在は in-memory 実装 を使用しています。

⸻

API エンドポイント

GET /requests

リクエスト一覧を取得します。

``` bash
curl http://localhost:8080/requests
```

⸻

GET /requests/{id}

指定した ID のリクエストを取得します。

``` bash
curl http://localhost:8080/requests/1
```

存在しない ID の場合は 404 を返します。

⸻

起動方法
``` bash
go run ./cmd/api
```

⸻

今後の拡張予定
•	POST /requests の追加
•	SQLite / RDB 実装の追加
•	テストコードの実装

⸻

学習目的
•	Go による Web API の基本構造理解
•	interface を用いた依存関係逆転の体験
•	実務を意識したディレクトリ設計

⸻

