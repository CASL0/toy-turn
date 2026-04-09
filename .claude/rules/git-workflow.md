---
description: Git操作（コミット、ブランチ作成、PR作成）を行う際に適用する
paths:
  - "**/*"
---

# Git Workflow

## Branching Strategy: GitHub Flow

- `main` ブランチは常にデプロイ可能な状態を保つ
- 新しい作業はすべて `main` から feature ブランチを切って行う
- ブランチ名は `feat/`, `fix/`, `docs/`, `refactor/`, `chore/` などのプレフィックスを付ける
- 作業が完了したら Pull Request を作成し、レビュー後に `main` へマージする

## Commit Message: Conventional Commits

コミットメッセージは [Conventional Commits](https://www.conventionalcommits.org/) に従う。

```
<type>(<scope>): <description>

[optional body]

[optional footer(s)]
```

### Type

- `feat`: 新機能
- `fix`: バグ修正
- `docs`: ドキュメントのみの変更
- `style`: コードの意味に影響しない変更（空白、フォーマット等）
- `refactor`: バグ修正でも機能追加でもないコード変更
- `test`: テストの追加・修正
- `chore`: ビルドプロセスや補助ツールの変更
- `ci`: CI 設定の変更

### Rules

- description は英語で簡潔に書く
- body が必要な場合は空行を挟んで詳細を記述する
- 破壊的変更がある場合は `BREAKING CHANGE:` フッターまたは `!` を type の後に付ける
