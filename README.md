# questions

## API

### POST /api/question

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
  "createdAt": "2023-03-08T16:57:58+09:00"
}
```

### GET /api/question

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
    "createdAt": "2023-03-08T16:57:58+09:00"
  }
]
```

### GET /api/question/:questionId

パラメーター

`questionId`: `bb76e15c-0335-485c-a264-4ccf1a1bbc15` とか

### POST /api/question/:questionId/answer

リクエスト

```json
{
  "answer": "かいとーう"
}
```

レスポンス

```json
{
  "id": "bb76e15c-0335-485c-a264-4ccf1a1bbc15",
  "question": "しつもーん",
  "answer": "かいとーう",
  "createdAt": "2023-03-08T16:57:58+09:00"
}
```

回答は上書きできる。
