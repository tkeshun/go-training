# TODOWebアプリケーションを作る



## 目標：テストが通るWebアプリケーションを作る

## 制限 
- go workを使うこと
- slogを使ってロガーを作成すること
    - json形式で構造化ログを出力すること
    - error発生時はErrorログ出力する
    - infoやwarnも使ってみて

- pkg/loggerは別パッケージにすること
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

### 事前準備

1. 各ディレクトリでgo.modを作る  
    `go mod init todo`、`go mod init pkg`

2. go.workを作る    
    `go work init todo pkg`


### 1. pkg/logger.goでロガーを作る。

#### 実装編
package pkgでパッケージをきる。
設定用の構造体を作る。
Loggerを実装。


#### Test編

testの雛形を作れるツールがあるので利用する。  
まだパッケージをインストールしていない場合は、`go install github.com/cweill/gotests/...@latest`でパッケージをインストールする(もしくはgo getする)。　　
`gotests -w -all logger.go`テストコードを生成する(tools.goの場合：`go run github.com/cweill/gotests/gotests -w -all logger.go`)。  
[Testify](https://github.com/stretchr/testify)を使ってUnitテストを書いていく。  
`go get github.com/stretchr/testify`でインストール。  

### 2. todoアプリのディレクトリ構造を決めて実装

今回、api,usecase,domain,infraに分けた。

