# SEIYU 5% OFF Web API

西友(SEIYU) - 5% OFF 開催日の一覧を JSON 形式で取得できる Web API です。

開催日のデータは、[西友 - 5%OFF開催日カレンダー | SEIYU](https://www.seiyu.co.jp/service/5off/) から取得しています。


## ドキュメント

### 月ごとの 5% OFF 開催日

`https://seiyu5off.mahoroi.com/api/v1/{year}/{month}.json`

#### URL の例

`GET` : `https://seiyu5off.mahoroi.com/api/v1/2019/03.json`

#### レスポンスの例

```json
{
    "updated": "2019-02-18 14:51:32",
    "date": [
        "2019-03-02",
        "2019-03-16"
    ]
}
```
