Simple Request API

Goで実装したシンプルな REST API です。

本リポジトリは DB実装に依存しない設計 を重視し、
永続化層を interface として切り出すことで、
実装詳細からビジネスロジックを分離した構成を採用しています。

---

概要
- Go（標準ライブラリのみ）で構成
- net/http によるシンプルな REST API
- 永続化層を interface として定義
- in-memory 実装により DBなしで動作可能
- レイヤードアーキテクチャを意識した設計

---

ディレクトリ構成

```bash
├── cmd/
│   └── api/
│       └── main.go        # エントリーポイント
├── internal/
│   ├── handler/           # HTTPリクエストの受付
│   ├── service/           # ビジネスロジック
│   ├── repository/        # 永続化層（interface + 実装）
│   └── model/             # ドメインモデル
└── README.md
```

---

設計方針

レイヤード構成
```bash
handler → service → repository(interface)
```

各レイヤーの責務

handler
- HTTPリクエスト / レスポンスの制御
- URL・メソッド・パラメータの解釈
- Service の呼び出し

service
- ビジネスロジックの責務
- Repository の interface にのみ依存
- 永続化の詳細を知らない

repository
- データ取得・保存の抽象化
- interface として定義
- 将来的な DB 実装差し替えを想定

---

Repository の設計

永続化層は interface として定義しています。
```go
type RequestRepository interface {
	FindAll() ([]model.Request, error)
	FindByID(id int64) (*model.Request, error)
	Create(req *model.Request) error
}
```

この設計により、
- DB実装が未完成でも API を動かせる
- Service 層のコードを変更せずに DB を差し替えられる

というメリットがあります。

現在は in-memory 実装 を使用しています。

---

API エンドポイント

GET /requests

リクエスト一覧を取得します。

```bash
curl http://localhost:8080/requests
```

---
GET /requests/{id}

指定した ID のリクエストを取得します。

```bash
curl http://localhost:8080/requests/1
```

- 存在しない ID の場合は 404 Not Found を返します

---

POST /requests

新しいリクエストを作成します。

```bash
curl -X POST http://localhost:8080/requests \
  -H "Content-Type: application/json" \
  -d '{
    "title": "sample title",
    "description": "sample description"
  }'
```

---

起動方法

通常起動
```bash
go run ./cmd/api
```

開発用（ホットリロード）

Goは標準でホットリロードを行わないため、
開発時は air を利用しています。

```bash
air
```


---
今後の拡張予定
- PUT /requests/{id} の追加
- DELETE /requests/{id} の追加
- SQLite / RDB 実装の追加
- テストコードの実装

---

学習目的
- Go による Web API の基本構造理解
- interface を用いた依存関係逆転（DIP）の体験
- レイヤードアーキテクチャの実践
- 実務を意識したディレクトリ設計の習得

---

補足

本リポジトリは「まず動かす → 設計を改善する」ことを目的とした学習用プロジェクトです。

小さく作りながら、設計上の判断や気づきを言語化することを重視しています。