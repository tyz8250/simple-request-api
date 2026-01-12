#「申請・受付を想定したシンプルな業務API」

なぜこのテーマがベストか

① 自治体・業務改善と直結
•申請
•受付
•状態管理

これは
•自治体
•社内システム
•DX案件

で一生使われる構造。

② フロント不要（重要）
•API中心
•curl / Postmanで完結
•UIに時間を吸われない

③ Goの基礎が全部入る
•HTTP
•JSON
•struct
•DB
•エラーハンドリング

👉 「Goで何ができるか」を一通り示せる。

④ 説明しやすい
•業務フローが想像しやすい
•READMEが書きやすい
•面接で話しやすい

⸻

成果物の具体像（最小構成）

📦 システム名（仮）

Simple Request API

🧱 機能（最初はこれだけ）

エンティティ：Request（申請）

フィールド	説明
id	申請ID
title	件名
description	内容
status	pending / approved / rejected
created_at	作成日時

API一覧（最小）
•POST /requests
→ 申請を作成
•GET /requests
→ 一覧取得
•GET /requests/{id}
→ 詳細取得

※ 更新・削除は最初は作らない

⸻

🧠 設計で考えるポイント（評価される）
•なぜ status は string か？
•struct と DB の責務分離
•handler / service / repository の分離

技術スタック（決め打ち）
•Go
•net/http
•encoding/json
•SQLite（最初は）
•sqlx or database/sql

README に必ず書く項目
1.何を作ったか
2.想定業務
3.API構成
4.設計で意識したこと
5.詰まった点と学び

👉 コードよりREADMEが主役

# Simple Request API

申請・受付業務を想定したシンプルなAPI。
GoでのHTTP処理と業務設計の練習を目的とする。

cmd/api/main.go

役割
•アプリの入口
•サーバ起動
•依存関係の組み立て

書かないこと
•ビジネスロジック
•SQL
•JSON変換

internal/model
役割
•業務データの定義
•Request構造体

type Request struct {
    ID          int64
    Title       string
    Description string
    Status      string
    CreatedAt   time.Time
}

internal/repository

役割
•DBとのやりとり
•SQLを書く場所
func (r *RequestRepository) FindAll() ([]model.Request, error)

internal/service

役割
•業務ルール
•状態の初期値など
func (s *RequestService) Create(title, description string) (*model.Request, error)

internal/handler

役割
•HTTPリクエストを受ける
•JSON変換
•ステータスコード
func (h *RequestHandler) Create(w http.ResponseWriter, r *http.Request)