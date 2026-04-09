# toy-turn

[pion/turn](https://github.com/pion/turn) v5 を使ったシンプルな TURN サーバー。

## Requirements

- [mise](https://mise.jdx.dev/)

## Setup

```bash
mise install
cp .env.example .env
# .env を編集して PUBLIC_IP, TURN_USER, TURN_PASSWORD を設定
```

## Usage

```bash
mise exec -- go run .
```

## Environment Variables

| 変数 | 必須 | デフォルト | 説明 |
|---|---|---|---|
| `PUBLIC_IP` | yes | - | サーバーの公開 IP アドレス |
| `TURN_USER` | yes | - | TURN 認証ユーザー名 |
| `TURN_PASSWORD` | yes | - | TURN 認証パスワード |
| `REALM` | no | `toy-turn` | TURN レルム |
| `PORT` | no | `3478` | UDP リスンポート |

## License

[MIT](LICENSE)
