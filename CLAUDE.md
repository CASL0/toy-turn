# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

toy-turn は pion/turn v5 を使ったシンプルな TURN サーバー。単一バイナリで UDP TURN サーバーを起動する。

## Development Environment

mise でツールチェインを管理。環境変数は `.env` ファイルから mise が自動で読み込む。

```bash
mise install          # Go 1.26.2 + golangci-lint 2.11.4 をインストール
cp .env.example .env  # 環境変数を設定
```

## Build / Run / Test / Lint

```bash
mise exec -- go build ./...              # ビルド
mise exec -- go run .                    # サーバー起動（.env の環境変数が必要）
mise exec -- go test -v ./...            # テスト実行
mise exec -- go test -v -run TestXxx .   # 単一テスト実行
mise exec -- golangci-lint run ./...     # lint + format チェック
```

## Environment Variables

| 変数 | 必須 | デフォルト | 説明 |
|---|---|---|---|
| `PUBLIC_IP` | yes | - | サーバーの公開 IP アドレス |
| `TURN_USER` | yes | - | TURN 認証ユーザー名 |
| `TURN_PASSWORD` | yes | - | TURN 認証パスワード |
| `REALM` | no | `toy-turn` | TURN レルム |
| `PORT` | no | `3478` | UDP リスンポート |

## Architecture

単一の `main.go` で構成。pion/turn v5 の `turn.NewServer` に `ServerConfig` を渡してサーバーを起動する。認証は長期クレデンシャル方式（`turn.GenerateAuthKey`）。シグナル（SIGINT/SIGTERM）でグレースフルシャットダウン。
