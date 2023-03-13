# Development

## 環境変数

`.env`を作って環境変数を指定できます。

例

```sh
TRAQ_CLIENT_ID=AS9vdsra85jioFai8HENFEJ3nrfvR43bdfgo
TRAQ_REDIRECT_URL=http://localhost:3000/api/oauth2/callback 
WEBHOOK_ID=967eb02b-23b0-ma94-b933-e976benbos98
WEBHOOK_SECRET=traQ-A
CLIENT_URL=http://localhost:5173
```

- `TRAQ_CLIENT_ID`
  - traQのクライアントID
  - [BOT Console](https://bot-console.trap.jp/)からクライアントを作って取得できる
  - 質問に回答するときの認証に必要

- `TRAQ_REDIRECT_URL`
  - traQのOAuthのリダイレクト先
  - `/api/oauth2/callback`にすると、traQの認証画面に飛び、ブラウザで使えるようになる。Postmanなどで使いたいときは違う適当なURLを入れるとよい

- `WEBHOOK_ID`
  - traQのWebhookのID。無い場合はWebhookを飛ばさない

- `WEBHOOK_SECRET`
  - Webhookを作ったときにSecretを設定したなら必要

- `CLIENT_URL`
  - クライアント側のURL。CORSの解決に必要になることがある。何も設定しなかった場合は`http://localhost:5173`になる

## Docker Compose

Docker Composeで動かせます。Airによるホットリロードが効いてます。

```sh
docker compose up --build
```

## MySQL

初期データは`mysql/init/0_init.sql`で定義しています。