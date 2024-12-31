# TODOWebアプリケーションを作る



## 目標：テストが通るWebアプリケーションを作る

## 制限 
- go workを使うこと
- slogを使ってロガーを作成すること
    - json形式で構造化ログを出力すること
    - error発生時はErrorログ出力する
    - infoやwarnも使ってみて

- pkg/loggerは別パッケージにすること
- apiにベアラートークンで認証をかけること
- DBアクセスをすること


## 参考にするといい資料

- slog

    https://qiita.com/Imamotty/items/3fbe8ce6da4f1a653fae
    https://pkg.go.dev/log/slog



- go work

    https://zenn.dev/kimuson13/articles/go-workspace-mode-impressions
    https://qiita.com/Rqixy/items/6bdead71dc02eb233376
    https://future-architect.github.io/articles/20220216a/

## 解答例

### 1. pkg/loggerパッケージを作る。

package pkgでパッケージをきる。


### 2. go workにtodoモジュールとpkgモジュールを追加する

### 3. ディレクトリ構造を決定する

今回、ui,usecase,domain,infraに分けた。