# Company X Golang Style Guide

# Introduction
このスタイルガイドは、Company X で開発される Golang コードのコーディング規約を示します。
Go の標準的な規約とベストプラクティスに基づいており、組織内の特定のニーズや好みに対応するために一部修正されています。

# Key Principles
* **読みやすさ:** コードは全てのチームメンバーにとって理解しやすいものであるべきです。
* **保守性:** コードは簡単に修正・拡張できる必要があります。
* **一貫性:** すべてのプロジェクトで一貫したスタイルに従うことで、コラボレーションを改善しエラーを減らします。
* **パフォーマンス:** Go はパフォーマンスとシンプルさを重視しています。コードは効率的であると同時に、不要な複雑さを避けるべきです。
* **Go 言語の流儀 (Idiomatic Go):** Go の規約に従い、過度な設計や標準から逸脱したパターンを避けましょう。

# コードフォーマット (Code Formatting)
* **`gofmt` の使用:** 全てのコードは `gofmt` を使用してフォーマットされる必要があります。これにより一貫性と可読性が確保されます。
* **行長:** 明確な制限はありませんが、可読性を保つために適切な長さに保ちましょう。
* **インデント:** インデントにはタブを使用し、必要に応じてスペースで整列します (Go の標準)。

# インポート (Imports)
* **インポートのグループ化:**
  * 標準ライブラリのインポート
  * サードパーティパッケージのインポート
  * 内部パッケージのインポート
* **インポートの順序:** グループごとに空行で分け、各グループ内でアルファベット順に並べます。

# 命名規則 (Naming Conventions)
* **変数:** 小文字のキャメルケースを使用 (例: `userName`, `totalCount`)。
* **定数:** エクスポートする定数には PascalCase を、エクスポートしないものには camelCase を使用 (例: `MaxValue`, `databaseName`)。
* **関数:** エクスポートする関数には PascalCase を、エクスポートしないものには camelCase を使用 (例: `CalculateTotal()`, `processData()`)。
* **パッケージ:** 短く小文字のみで表記 (例: `userutil`, `paymentgateway`)。アンダースコアや大文字の使用は避ける。
* **型と構造体:** エクスポートする型は PascalCase で表記 (例: `UserManager`, `PaymentProcessor`)。

# コメント (Comments)
* **完全な文を使用:** コメントは文法的に正しい文章で記述し、ピリオドで終わるようにします。
* **ドキュメントコメント:** エクスポートする要素には宣言の直前にコメントを記載 (例: `// CalculateTotal は全ての値を合計して結果を返します。`)。
* **インラインコメント:** 必要な場合に限り、複雑なロジックを説明するために使用します。

# ログ出力 (Logging)
* **標準的なログフレームワークの使用:** `log` または構造化ログを扱う `zap` を推奨します。
* **適切なログレベル:** INFO、DEBUG、WARNING、ERROR、FATAL を使い分けましょう。
* **コンテキストの追加:** デバッグを助けるために適切な情報を含めるようにしましょう。

# エラーハンドリング (Error Handling)
* **`error` の使用:** 失敗する可能性のある関数は `error` 型を最後の戻り値として返すべきです。
* **`panic` の使用を避ける:** `panic` は致命的なエラーやプログラムエラー時のみ使用してください。
* **`fmt.Errorf()` や `errors.New()` の使用:** 有益なエラーメッセージを提供するようにしましょう。

# ツール (Tooling)
* **コードフォーマッタ:** `gofmt` - 一貫したフォーマットを自動的に適用します。
* **リンター:** `golint`, `staticcheck`, `errcheck` - 潜在的な問題やスタイル違反を検出します。
* **テスト:** `go test` を使用してユニットテストを記述および実行します。

# サンプルコード (Example)
```go
// Package userauth はユーザー認証に関するユーティリティを提供します。

package userauth

import (
    "crypto/sha256"
    "encoding/hex"
    "errors"
    "log"
)

// HashPassword はパスワードを SHA-256 でハッシュ化します。
func HashPassword(password string) (string, error) {
    if password == "" {
        return "", errors.New("パスワードは空ではいけません")
    }

    hash := sha256.New()
    hash.Write([]byte(password))
    return hex.EncodeToString(hash.Sum(nil)), nil
}

// AuthenticateUser は提供されたパスワードが保存されたハッシュと一致するかを確認します。
func AuthenticateUser(storedHash, password string) bool {
    if storedHash == "" || password == "" {
        log.Println("認証に無効な入力が含まれています")
        return false
    }

    hash, err := HashPassword(password)
    if err != nil {
        log.Println("ハッシュ生成中のエラー:", err)
        return false
    }

    return storedHash == hash
}
```

