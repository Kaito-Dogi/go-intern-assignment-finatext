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

## 工夫した点

- 命名規則
  - Effective Goの[Names](https://go.dev/doc/effective_go#names)を参考にした．
- 可視性
  - 不用意な参照を避けるため，大文字から始まる名前は最低限にした．
- package構成
  - 構造体
