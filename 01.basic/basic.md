# 基礎

## Goの特徴

公式サイトより引用

> An open-source programming language supported by Google
> Easy to learn and great for teams
> Built-in concurrency and a robust standard library
> Large ecosystem of partners, communities, and tools

Googleがサポートするオープンソースのプログラミング言語
簡単に習得でき、チームに最適
組み込みの並行処理と堅牢な標準ライブラリ
パートナー、コミュニティ、ツールの大規模なエコシステム

簡単に試したい場合は、[Go Playground](https://go.dev/play/)がある

## 基本構文

- 参考
  - https://go.dev/doc/tutorial/getting-started
  - 実用Go言語
  - https://koko206.hatenablog.com/entry/2024/01/06/055112

### Hello World

[Go Playground](https://go.dev/play/)を開き、右上のRunを押してみる。

以下プログラムが実行されるはずである。
```
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
```

ここでは簡単に解説する。
`//`はコメントアウト。　　
`package`はGoでいうプログラムをくくる単位。`go run`、`go build`時にエントリーポイントになる。　　
`import fmt`はfmt packageをプログラム内で使用可能にしている。パッケージ外のプログラムを使用する際はimportして利用する。　　
Goは`func`で関数宣言する。
`fmt`は[標準パッケージ](https://pkg.go.dev/std)であり、出力関連のライブラリである。　　

> Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
訳：fmtパッケージは、C言語のprintfやscanfに類似した関数を用いてフォーマットされた入出力を実現する機能を提供する。
ちなみにGoでは変数宣言しておいて使わない変数があるとエラーになる。

### main関数とエントリーポイント

Goでは**mainパッケージ**の**main関数**がエントリーポイントになる。
main関数は必ずmain package内で定義されている必要がある。
main関数は以下の特徴がある。

1. パラメータを受け取らない
2. 戻り値を返しません
3. プログラムの状態初期化やプログラムロジックの実行を担当する


パッケージ名がmainでない例をもとに観察する。
下記サンプルを用意する。`package main`ではなく`package example1`となっている。
```
package example1

import "fmt"

func main() {
	fmt.Println("Hello Go")
}
```
実行(go run)するとコンパイルエラーがでる。
mainパッケージでないと下記エラーになる。
```
go run example.go 
package command-line-arguments is not a main package
```

ビルドするとエントリーポイントではないのでなにも作られない。
```
go build example.go 
```
名前付きビルドするとファイルはできるが実行できない。
```
go build -o example1 example.go 
shun@shun-ThinkPad-P14s-Gen-4:~/workspace/go-training/basic/example1$ ls -al
合計 32
drwxrwxr-x 2 shun shun  4096 12月 18 23:04 .
drwxrwxr-x 3 shun shun  4096 12月 18 22:49 ..
-rw-rw-r-- 1 shun shun    73 12月 18 22:50 example.go
-rw-rw-r-- 1 shun shun 20140 12月 18 23:04 example1
```

実行権限をつけて無理やり実行してもエラーになる。

```
shun@shun-ThinkPad-P14s-Gen-4:~/workspace/go-training/basic/example1$ chmod +x example1 
shun@shun-ThinkPad-P14s-Gen-4:~/workspace/go-training/basic/example1$ ls -al example1
-rwxrwxr-x 1 shun shun 20140 12月 18 23:04 example1
shun@shun-ThinkPad-P14s-Gen-4:~/workspace/go-training/basic/example1$ ./example1 
./example1: 行 1: 予期しないトークン `newline' 周辺に構文エラーがあります
./example1: 行 1: `!<arch>'
```

このように`package main`がないと実行可能なプログラムが生成されない。

次に`main関数`がない場合を観察する。
以下のコードを用意する。
`func main`ではなく、`func example2`となってる
```
package main

import "fmt"

func example2() {
	fmt.Println("Hello Go")
}
```

実行すると以下のエラーがでる。
```
shun@shun-ThinkPad-P14s-Gen-4:~/workspace/go-training/basic/example2$ go run example2.go 
# command-line-arguments
runtime.main_main·f: function main is undeclared in the main package
```

main package内に`main` 関数がないと言ってる。
この状態でビルドすると、そもそもコンパイルエラーとなる。
```
go build example2.go 
# command-line-arguments
runtime.main_main·f: function main is undeclared in the main package
```

今回の場合は、名前付きビルドでも、コンパイルエラーでビルドが止まり、そもそもビルド成果物を出力できない。

```
 go build -o example2 example2.go 
# command-line-arguments
runtime.main_main·f: function main is undeclared in the main package
```

では最後に正しいプログラムを書いてビルド・実行してみる。

```
package main

import "fmt"

func main() {
	fmt.Println("Hello Go")
}
```
下記コマンドで実行する。
実行が確認できた。
```
shun@shun-ThinkPad-P14s-Gen-4:~/workspace/go-training/basic/example3$ go run example3.go 
Hello Go
```

#### クイズ

- Q1. パッケージを使えるようにするにはなにが必要？


- Q2: mainパッケージの必要性について
    次のうち、Goプログラムを正しくビルドおよび実行するために正しいものをすべて選択してください。

    1. エントリーポイントのパッケージ名はmainである必要がある。
    2. main関数はどのパッケージでも定義できる。
    3. main関数は必ずmainパッケージ内に定義する必要がある。
    4. パッケージ名がmainでなくても実行可能なプログラムを生成できる。

解答１，３

- Q3. 以下のコードをコンパイルしようとすると何が起こりますか？

    ```
    // main.go
    package example

    import "fmt"

    func main() {
        fmt.Println("Hello Go")
    }
    ```

    ```
    go run main.go
    ```

### 変数宣言と基本のデータ型

変数宣言の仕方のバリエーションと基本的なデータ型について教える。

#### 変数宣言

変数宣言は下記の方法でできる。

- `var 変数名 型 = 値`
    varに続いて変数名・型・値を指定する。
- `var 変数名 型`
    値を指定しない場合型のゼロ値がはいる。
- `var 変数名 = 値`
    値からの型推論で変数の型を決める
- `変数名 := 値`
    varを省略できる。値からの型推論で変数の型を決める。
    一番使われる（主観）
    グローバル変数では使用できない。

    エラーになる例：https://go.dev/play/p/uEX9T2lWktc

    ```
    package main

    import "fmt"

    year := 2024
    　
    func main() {
       fmt.Pirnln(year) 
    }
    ./prog.go:5:5: syntax error: non-declaration statement outside function body
    ./prog.go:6:5: invalid character U+3000 in identifier
    ```

- `変数名 := make(型, 長さ, 容量)`
    特定の組み込み型（スライス、マップ、チャネル）を初期化するために使用される関数
    ```
    s := make([]int, 5, 10)
    ```

#### 基本的なデータ型

Webプログラミングで使いそうなものに絞って教える。
他にも興味がある人は[言語仕様](https://go.dev/ref/spec#Types)をみるといい。

- string
    文字列型。`+`で連結できる。連結時はコピーを作成して渡す。ループで回しまくるとパフォーマンスが落ちるので`strings`パッケージの`strings.Builder`を使うべき。
    ```
    package main

    import (
        "fmt"
        "strings"
    )

    func main() {
        // strings.Builderを初期化
        var builder strings.Builder

        // 複数の文字列を追加
        builder.WriteString("Hello")
        builder.WriteString(", ")
        builder.WriteString("Go")
        builder.WriteString(" World!")

        // 連結された文字列を取得して表示
        fmt.Println(builder.String()) // Hello, Go World!
    }
    ```

- int
    `int8`、`int16`、`int32`、`int64`のバリエーションがある。bit数を指定しない場合、環境依存になる。

- float64
    浮動小数点。32bitもあるけど基本使わない。

- bool
    真偽値。    
    trueとfalseがある

- slice
    配列みたいなもん。格納したデータにインデックスでアクセスできる。0オリジン。
    構文は`[]型`
    - 要素への参照先: 実データを参照するポインタ。
    - 長さ (len): スライスに含まれる要素数。
    - 容量 (cap): スライスの内部配列が保持できる最大要素数。
    
    型と長さと容量が定義されてる。
    容量を超えると元の容量の2倍くらい確保する
    1024を超えると1.25倍になるらしい。
    容量確保時は新たなスライスにコピーするので遅い。
    容量がわかってる場合や予測できる場合、予め確保しておくとよい。
    
    ```
    s := make([]int, 0, 2) // 長さ0、容量2のスライス
    fmt.Println(cap(s))    // 出力: 2
    s = append(s, 1, 2, 3) // 容量を超える3つの要素を追加
    fmt.Println(cap(s))    // 出力: 4（容量が倍増）
    ```
- map
    キーとペアを管理するデータ構造
    構文は`map[キー型]値型`。
- []byte
    gRPCやhttpレスポンスの書き込み、メールなどで地味に使う。

- time.Time
    プリミティブな型じゃないけどよく使う。


上記の型を使ったサンプルを以下に示す。

```
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	// string型: APIレスポンスやユーザー入力データを扱う
	name := "John Doe"

	// int型: ユーザーIDやページネーション用の値を扱う
	id := 123

	// float64型: 金額や割合の計算に使用
	price := 99.99

	// bool型: 認証状態やフラグの管理に使用
	isLoggedIn := true

	// slice型: 複数のデータを動的に扱う（JSON配列やクエリパラメータの処理に便利）
	users := []string{"Alice", "Bob", "Charlie"}

	// map型: キーと値のペアを扱う（HTTPヘッダーやクエリパラメータの管理）
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer token",
	}

	// []byte型: JSONのエンコード・デコードやバイナリデータの処理に使用
	data := []byte(`{"message":"Hello, Go!"}`)

	// time.Time型: APIのタイムスタンプや期限管理に使用
	now := time.Now()


	// 各値を出力
	fmt.Println("Name:", name)
	fmt.Println("ID:", id)
	fmt.Println("Price:", price)
	fmt.Println("Is Logged In:", isLoggedIn)
	fmt.Println("Users:", users)
	fmt.Println("Headers:", headers)
	fmt.Println("Raw JSON:", string(data))
	fmt.Println("Current Time:", now)
}
```


#### ちょっと特殊な型(anyとstruct{})

- any
    何入れてもOKな型。
    anyはjsonのパースなどに使われる。
    以下に例を示す。

    ```
    var value any = "Hello" // 型を限定せず任意の値を入れることが可能
    str, ok := value.(string) // 型アサーションで取り出す。値と取り出せたかのbool値が入る。
    if ok {
        fmt.Println("String value:", str)
    }
    ```
- struct{}
    structは型が不要なときに使われる。`struct{}`は全部の型の中で一番データ容量が軽いはず。
    chatgpt調べ
    ```
    struct{}は空の構造体。
    サイズが0バイトの型として扱われ、メモリ効率が非常に高い。
    用途:   
    マーカーやシグナルとして使用。
    データが不要な場合のメモリ効率化。
    ```    

    以下に例を示す。
    メソッドだけ実装したいときに使う（主観）
    ```
    package main

    import "net/http"

    type CustomHandler struct{} // anyだとエラーになる

    // http.Handlerのインターフェースを実装する
    func (h *CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Custom handler\n"))
    }

    func main() {
        handler := &CustomHandler{}
        // DefaultServeMuxにカスタムハンドラーを登録する
        http.Handle("/", handler)
        http.ListenAndServe(":8080", nil)
    }
    ```

### 構造体と型

### 関数宣言

- `func`で関数宣言できる

    ```
    func 関数名(引数) 型 {
        処理
        return 値
    }
    ```
- 返り値なしのパターン

    ```
    func 関数名(引数) 型 {
        処理
        return 値
    }
    ```

- 名前付き返り値

    関数のローカル内で返り値に定義した変数が使える。
    関数内で一番長いライフタイムを持つ
    ```
    func 関数名(引数) (変数名 型) {
        処理
        return 値
    }
    ```

- 複数の返り値を受け取るパターン
    制限なく複数の返り値を指定できるが2~3個くらいしか見たことない。
    いっぱい返すなら構造体定義したほうがいい。
    ```
    func 関数名(引数) (型, 型) {
        処理
        return 値,値
    }
    ```

- 可変長引数を受け取るパターン
    正直あんまり使わない。ライブラリ作る人とかは使うかも

    ```
    func 関数名(引数 ...型名) {
        処理
        return 
    }
    ```

## エラーハンドリング

## インターフェース

## モジュールとパッケージ


## 外部ライブラリの利用

### pgx(DBドライバー)

### database/sql

DBアクセスを抽象化する汎用的なインターフェースを提供するライブラリ。
実際にDBとアクセスを行う実装として、DBドライバーが必要。

## net/http

## テスト


## 並行処理

