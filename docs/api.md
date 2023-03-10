# API

## GET /api/oauth2/authorize

レスポンス
303
環境変数の`TRAQ_REDIRECT_URL`にリダイレクト

## GET /api/oauth2/callback

レスポンス
200
成功

401
失敗

404
codeがない

## POST /api/question

リクエスト

```json
{
  "question": "しつもーん"
}
```

レスポンス
200

```json
{
  "id": "bb76e15c-0335-485c-a264-4ccf1a1bbc15",
  "question": "しつもーん",
  "answer": "",
  "answerer": "",
  "createdAt": "2023-03-08T16:57:58+09:00",
  "updatedAt": "2023-03-08T16:57:58+09:00"
}
```

## GET /api/question

パラメーター

- `limit` デフォルト 10
- `offset` デフォルト 0

レスポンス
200

```json
[
  {
    "id": "bb76e15c-0335-485c-a264-4ccf1a1bbc15",
    "question": "しつもーん",
    "answer": "かいとうー",
    "answerer": "7265b13d-9e06-42f6-98e3-41ea742f8fb2", //traQのユーザーuuid
    "createdAt": "2023-03-08T16:57:58+09:00",
    "updatedAt": "2023-03-09T08:57:58+09:00"
  }
]
```

## GET /api/question/:questionId

パラメーター

`questionId`: `bb76e15c-0335-485c-a264-4ccf1a1bbc15` とか

## POST /api/question/:questionId/answer

traQでのログインが必要

リクエスト

```json
{
  "answer": "かいとーう"
}
```

レスポンス

200

```json
{
  "id": "bb76e15c-0335-485c-a264-4ccf1a1bbc15",
  "question": "しつもーん",
  "answer": "かいとーう",
  "answerer": "7265b13d-9e06-42f6-98e3-41ea742f8fb2", //traQのユーザーuuid
  "createdAt": "2023-03-08T16:57:58+09:00",
  "updatedAt": "2023-03-09T08:00:58+09:00"
}
```

401
traQでログインしていない

回答は上書きできる。

全体として、時間のフォーマットがずれそう
