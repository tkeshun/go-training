# 課題：並行アクセス処理を作る

## 目標：サンプルを改良して、アクセス時間を縮める

初期状態：1分程度、アクセスするエンドポイントは20個。
レスポンス内容を番号順にソートして出力する。
テストコードを実行して速度計測する。
access-appに初期状態のコードがあるので改変してください。

### 出力例

```
$ go run main.go
Response from /api/0: {"message":"Endpoint 0 OK","delayTime":3}

Response from /api/1: {"message":"Endpoint 1 OK","delayTime":4}

Response from /api/2: {"message":"Endpoint 2 OK","delayTime":2}

Response from /api/3: {"message":"Endpoint 3 OK","delayTime":5}

Response from /api/4: {"message":"Endpoint 4 OK","delayTime":3}

Response from /api/5: {"message":"Endpoint 5 OK","delayTime":1}

Response from /api/6: {"message":"Endpoint 6 OK","delayTime":6}

Response from /api/7: {"message":"Endpoint 7 OK","delayTime":3}

Response from /api/8: {"message":"Endpoint 8 OK","delayTime":2}

Response from /api/9: {"message":"Endpoint 9 OK","delayTime":4}

Response from /api/10: {"message":"Endpoint 10 OK","delayTime":3}

Response from /api/11: {"message":"Endpoint 11 OK","delayTime":1}

Response from /api/12: {"message":"Endpoint 12 OK","delayTime":2}

Response from /api/13: {"message":"Endpoint 13 OK","delayTime":5}

Response from /api/14: {"message":"Endpoint 14 OK","delayTime":3}

Response from /api/15: {"message":"Endpoint 15 OK","delayTime":2}

Response from /api/16: {"message":"Endpoint 16 OK","delayTime":3}

Response from /api/17: {"message":"Endpoint 17 OK","delayTime":4}

Response from /api/18: {"message":"Endpoint 18 OK","delayTime":2}

Response from /api/19: {"message":"Endpoint 19 OK","delayTime":2}

Total elapsed time: 6.001512784s
```

## 制限
serverは改変禁止！
抜け道みたいなのは禁止！
実際にはアクセスしないとか