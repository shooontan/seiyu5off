# SEIYU 5% OFF Web API

西友(SEIYU) - 5% OFF 開催日の一覧を JSON 形式で取得できる Web API です。

開催日のデータは、[西友 - 5%OFF開催日カレンダー | SEIYU](https://www.seiyu.co.jp/service/5off/) から取得しています。


## ドキュメント

### 年ごとの 5% OFF 開催日

`https://seiyu5off.mahoroi.com/api/v1/:year.json`

#### URL の例

`GET` : `https://seiyu5off.mahoroi.com/api/v1/2019.json`

#### レスポンスの例

```json
{
    "updated": "2019-05-19 07:29:24",
    "date": [
        "2019-01-05",
        "2019-01-19",
        "2019-02-02",
        "2019-02-16",
        "2019-03-02",
        "2019-03-16",
        "2019-04-06",
        "2019-04-20",
        "2019-05-04",
        "2019-05-11",
        "2019-05-18",
        "2019-06-01",
        "2019-06-15",
        "2019-06-22",
        "2019-06-29",
        "2019-07-06",
        "2019-07-13",
        "2019-07-20"
    ]
}
```

### 月ごとの 5% OFF 開催日

`https://seiyu5off.mahoroi.com/api/v1/:year/:month.json`

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
