# Finatextインターン サーバーエンジニア 選考課題

## 実装内容

1. ローカル環境で `GET http://localhost:8080` でアクセスできるAPIサーバを⽴ち上げる．
1. 以下の仕様を満たすAPIをローカル環境で実装する．

### リクエスト

- メソッド： `GET`
- URL： `/address`
- パラメータ： `postal_code`
  - 説明：郵便番号7桁（ハイフン無し）
  - 例：1050011
- 住所の取得には外部API（[http://zip.cgis.biz/xml/zip.php?zn=1020073](http://zip.cgis.biz/xml/zip.php?zn=1020073)）を用いる．

### レスポンス

- フィールド

|  要素  |  型  |  説明  |
| ---- | ---- | ---- |
|  postal_code  |  string  |  リクエストパラメータで与えた郵便番号7桁（ハイフン無し）  |
|  address  |  string  |  外部APIから取得した住所  |
|  address_kana  |  string  |  外部APIから取得した住所の読み仮名  |

### サンプルレスポンス

```
{
    "postal_code": "1050011",
    "address": "東京都港区芝公園",
    "address_kana": "トウキョウトミナトクシバコウエン"
}
```

### エラー

|  コード  |  説明  |
| ---- | ---- |
| 400 | 無効なパラメータや無効な郵便番号が渡された場合に返される． |

### エラーレスポンス

```
400 Bad Request
```

## 工夫した点

- 命名規則
  - Effective Goの[Names](https://go.dev/doc/effective_go#names)を参考にした．
- 変数や構造体，関数の可視性
  - 不用意な参照を避けるため，大文字から始まる名前を最低限にした．
- package構成
  - プロジェクトの開発が続き，コードが肥大化してしまうことを想定して，packageを切り分けた．
  - 具体的には，構造体の定義を `types` というpackage内に記述した．
- 関数化
  - 適切な粒度の関数になるように処理をまとめた．
  - 具体的には，各種 `Handler` や `GET` メソッドを関数に切り出した．
- エラーハンドリング
  - 無効なパラメータや無効な郵便番号が渡された場合， `400 Bad Request` を返すようにした．
  - その他は[Error handling and Go](https://go.dev/blog/error-handling-and-go)を参考にした．

## 感想
- 意外にも手軽にローカルサーバを立てられることに感動しました．
- Go初心者のため，基本文法を学ぶのに[A Tour of Go](https://go-tour-jp.appspot.com/)を利用しました．新しい言語に触れるのはやはり楽しかったです．
- Goの基本文法を学ぶ中で，ポインタの理解を深められたことがためになりました．
- APIサーバの開発はやはり楽しいです．
  - 特に，データを適切な形式に整形してそれを返すという流れが好きです．
  - インターンではより親切なエラーハンドリングの書き方も学びたいです．

**引き続きよろしくお願いいたします🙇‍♂️**
