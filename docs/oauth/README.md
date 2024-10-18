# OAuth概要

## 認証の流れ
### 1.SignUPでユーザーがGoogleの認証ボタンを押下
フロント側で実装されたGoogleの認証ボタン

### 2.フロントエンドでGoogleに対してリクエスト(Googleの認証エンドポイント)
フロントからGoogleに対してのGETリクエストの部分

URL: `https://accounts.google.com/o/oauth2/auth?client_id=YOUR_CLIENT_ID&redirect_uri=https://your-backend.com/oauth/callback&response_type=code&scope=email%20profile`

クエリパラメータ中身
```
client_id: Google Cloudで取得したOAuthクライアントID
redirect_uri: Google認証が成功した後に、ユーザーをリダイレクトするバックエンドのエンドポイント（例: https://your-backend.com/oauth/callback）
response_type: code（認可コードを取得するため）
scope: email, profile（取得するGoogleのユーザーデータ）
```

### 3.ユーザーがGoogleに対して認証をする。
ユーザーがemailとパスワードで認証

### 4.認証が完了するとGoogleからリダイレクトで認可コードが返ってくる。
認証が成功すると、Googleはフロントエンドにリダイレクトし、URLのクエリパラメータとして認可コード（code）を付与。
このリダイレクト先はフロントエンドが事前に指定したURL（通常バックエンドのエンドポイント）。

URL: `https://your-backend.com/oauth/callback?code=AUTHORIZATION_CODE`

### 5.その認可コードを使ってバックエンドからGoogleに対してリクエスト

URL: `https://oauth2.googleapis.com/token`

**POSTリクエストの中身**
```
code: Googleから送られてきた認可コード
client_id: Google Cloudで発行されたクライアントID
client_secret: Google Cloudで発行されたクライアントシークレット
redirect_uri: 認可コードを受け取ったバックエンドのエンドポイントと同じURI
grant_type: authorization_code
```

レスポンスの例
```
{
  "access_token": "ya29.A0AR...",
  "expires_in": 3599,
  "scope": "email profile",
  "token_type": "Bearer",
  "id_token": "eyJhbGciOiJSUzI1NiIsImtpZCI6...",
  "refresh_token": "1//0g..."
}
```

### 6.レスポンスとして、Googleからアクセストークンが返ってくる。

### 7.アクセストークンを使ってGoogle APIからユーザー情報を取得し
リクエスト
```
GET https://www.googleapis.com/oauth2/v2/userinfo
Authorization: Bearer YOUR_ACCESS_TOKEN
```
レスポンス例
```
{
  "id": "1234567890",
  "email": "user@example.com",
  "verified_email": true,
  "name": "User Name",
  "given_name": "User",
  "family_name": "Name",
  "picture": "https://lh3.googleusercontent.com/a-/AOh14G...",
  "locale": "en"
}
```

### 8.データベースにユーザーを登録する（または既存ユーザーならログイン処理を行う）。
サインアップ: まだユーザーが存在しない場合、取得した情報をもとに新規ユーザーをデータベースに登録します。
ログイン: 既にユーザーが登録済みであれば、ログイン処理を行います。

### 9.最終的にJWTを生成し、フロントエンドに返す。
