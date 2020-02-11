# SEIYU 5% OFF Web API & Google カレンダー

西友(SEIYU) - 5% OFF 開催日の一覧を JSON 形式で取得できる Web API と、開催日が登録されている Google カレンダーです。

開催日のデータは、[西友 - 5%OFF開催日カレンダー | SEIYU](https://www.seiyu.co.jp/service/5off/) から取得しています。


## Web API ドキュメント

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

## Google カレンダー

5% OFF 開催日が登録されている Google カレンダーです。

カレンダー ID: `o7vmmt5nqggrbavlleq5vklnbk@group.calendar.google.com`

公開 URL: [https://calendar.google.com/calendar/embed?src=o7vmmt5nqggrbavlleq5vklnbk%40group.calendar.google.com](https://calendar.google.com/calendar/embed?src=o7vmmt5nqggrbavlleq5vklnbk%40group.calendar.google.com)

### カレンダー追加方法

1. [Google カレンダー](https://calendar.google.com/calendar/)にアクセス
2. 左下メニュー「他のカレンダー」->「カレンダーに登録」
3. フォームに上記のカレンダー ID を入力
