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

## GET /api/me

レスポンス
200
成功

```json
{
  "id":"bb76e15c-0335-485c-a264-4ccf1a1bbc15",
  "name": "ikura-hamu",
  "displayName":"いくら・はむ"
}
```

401
ログインしていない

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

回答済み・未回答含めた全質問

パラメーター

- `limit` デフォルト 10
- `offset` デフォルト 0

レスポンス
200

```json
{
  "count":1, //総件数
  "questions":[
    {
      "id": "bb76e15c-0335-485c-a264-4ccf1a1bbc15",
      "question": "しつもーん",
      "answer": "かいとうー",
      "answerer": "ikura-hamu", //traQのユーザーid
      "createdAt": "2023-03-08T16:57:58+09:00",
      "updatedAt": "2023-03-09T08:57:58+09:00"
    }
  ]
}

```

## GET /api/question/answered

回答済みの質問一覧

パラメーター

- `limit` デフォルト 10
- `offset` デフォルト 0

レスポンス
200

```json
{
  "count":1, //総件数
  "questions":[
    {
      "id": "bb76e15c-0335-485c-a264-4ccf1a1bbc15",
      "question": "しつもーん",
      "answer": "かいとうー",
      "answerer": "ikura-hamu", //traQのユーザーid
      "createdAt": "2023-03-08T16:57:58+09:00",
      "updatedAt": "2023-03-09T08:57:58+09:00"
    }
  ]
}

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
