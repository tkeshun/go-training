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

Goでは変数名の先頭が小文字だとプライベートな値、大文字だとパブリックな値となる。
プライベートの値はパッケージ外からアクセスできなくなる。


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

### 構造体と型

### 型

[言語仕様](https://go.dev/ref/spec#Types)では以下のように書かれている。
> A type determines a set of values together with operations and methods specific to those values. A type may be denoted by a type name, if it has one, which must be followed by type arguments if the type is generic. A type may also be specified using a type literal, which composes a type from existing types.
> (翻訳)　型は、特定の値の集合と、それらの値に特有の操作やメソッドを決定します。型には名前が付いている場合があり（型名）、その場合、型がジェネリックであれば型引数を伴う必要があります。また、型リテラルを使用して、既存の型を組み合わせて新しい型を指定することもできます。 

型に名前をつけたり、複数の型を組み合わせて新しい型を作ることが可能

型に名前をつける例を次に示す。
```
type MyInt int
```

このコードはMyIntという新しい型を定義している。この型はintをベースにメソッドを拡張することができる。
このとき定義した型はint型とは別物扱いになる。
以下のコードをPlaygraundに貼り付けて試してほしい。

```
package main

import "fmt"

// 型定義
type MyInt int

func main() {
    var a int = 10
    var b MyInt = 10

    // 型を直接比較（エラーになります）
    // fmt.Println(a == b) // コンパイルエラー: mismatched types int and MyInt

    // 明示的に型変換すれば比較可能
    fmt.Println(a == int(b)) // true

    // 型の判定
    fmt.Printf("Type of a: %T\n", a) // int
    fmt.Printf("Type of b: %T\n", b) // main.MyInt
}
```

int と MyInt は異なる型として扱われ、直接比較するとコンパイルエラーになる。
明示的に型変換することで比較が可能になる。
次節で複数の型をまとめた型について説明する。

### 構造体

構造体について、[言語仕様](https://go.dev/ref/spec#Types)では以下のように書かれている。

> A struct is a sequence of named elements, called fields, each of which has a name and a type. Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField). Within a struct, non-blank field names must be unique.
> (翻訳)　構造体（struct）は、フィールドと呼ばれる名前付き要素の列で構成され、それぞれに名前と型があります。フィールド名は明示的に指定することも（IdentifierList）、暗黙的に指定することも（EmbeddedField）できます。同じ構造体内では、空白でないフィールド名は一意でなければなりません。

#### 基本

前節で触れた、複数の型をまとめた型が構造体である。
Javaでいうクラスみたいなもの。
以下のようにかく。

```
type StructName Struct {
    FieldName1 int
    Fieldname2 string
}
```

実際には以下のように使う。

```
package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	gradYear int
}

func main() {
	var p = Person{Name: "Alice", Age: 30, gradYear: 2023}
	fmt.Println(p.Name)     // Alice
	fmt.Println(p.Age)      // 30
	fmt.Println(p.gradYear) // 2023
}

```

宣言するときは、フィールドを指定して値を入れる。
フィールドの先頭が小文字だとプライベートフィールドになる。
プライベートフィールドの場合は、パッケージ外からはアクセスできない。

※　割と実践的な例（実用Go言語より）

Statusコードによって出力する内容を変えられる。

```
package main

import "fmt"

type HTTPStatus int

const (
	StatusOK              HTTPStatus = 200
	StatusUnauthorized    HTTPStatus = 401
	StatusPaymentRequired HTTPStatus = 402
	StatusForbidden       HTTPStatus = 403
)

// String()メソッドが実装されてると、fmt.Print系で呼び出される
func (s HTTPStatus) String() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusUnauthorized:
		return "Unauthorized"
	case StatusPaymentRequired:
		return "Payment Required"
	case StatusForbidden:
		return "Forbidden"
	default:
		return fmt.Sprintf("HTTPStatus(%d)", s)
	}
}

func main() {
	fmt.Println(StatusOK) // 正しい
	printHTTPStatus(200)  // int型の拡張なので通る
	// printHTTPStatus("200") // stringだとコンパイルエラーになる
}

func printHTTPStatus(s HTTPStatus) {
	fmt.Println(s)
}
```

#### メソッドを生やす

構造体には関数を紐付けることでメソッドの定義ができる。
具体的には、レシーバー(pythonでいうself, javaでいうthis)を指定して関数を定義する。

基本形
```
func (レシーバ名 レシーバ型) メソッド名(引数) 戻り値型 {
    // メソッドの処理
}
```

- レシーバ名: メソッドが属する構造体のインスタンスを表す変数名（慣例として1文字が多い）
- レシーバ型: メソッドが関連付けられる構造体の型
- 引数と戻り値: 通常の関数と同じように定義可能

具体例を以下に示す。

```
package main

import "fmt"

// 構造体の定義
type Person struct {
    Name string
    Age  int
}

// 値レシーバのメソッド
func (p Person) Greet() string {
    return "Hello, my name is " + p.Name
}

func main() {
    p := Person{Name: "Alice", Age: 30}
    fmt.Println(p.Greet()) // Hello, my name is Alice
}
```

上記のうち、レシーバーは２つに分類できる。
値レシーバとポインタレシーバだ。

- 値レシーバ
    構造体のコピーが渡される。
    フィールドの変更が元の構造体に伝播しない。
    ```
    // 値レシーバのメソッド
    func (p Person) Greet() string {
        return "Hello, my name is " + p.Name
    }
    ```

- ポインタレシーバ
    構造体のポインタが渡される。
    フィールドの変更が元の構造体に伝播する。
    レシーバにアスタリスク(*)を前置するとポインタレシーバになる。
    ```
    // 値レシーバのメソッド
    func (p *Person) Greet() string {
        return "Hello, my name is " + p.Name
    }
    ```

値レシーバとポインタレシーバの比較の例を下に示す。
```
package main

import "fmt"

// 構造体の定義
type Person struct {
    Name string
    Age  int
}

// 値レシーバのメソッド（コピーが渡される）
func (p Person) CompareAddressValue(other Person) bool {
    return &p == &other
}

// ポインタレシーバのメソッド（元のアドレスが渡される）
func (p *Person) CompareAddressPointer(other *Person) bool {
    return p == other
}

func main() {
    // 2つの構造体インスタンスを作成
    p1 := Person{Name: "Alice", Age: 30}
    p2 := Person{Name: "Alice", Age: 30}

    // 値レシーバを使用
    fmt.Println("Value Receiver Address Comparison:")
    fmt.Printf("p1 == p2: %v\n", p1.CompareAddressValue(p2)) // false（異なるコピーのアドレス）

    // ポインタレシーバを使用
    fmt.Println("Pointer Receiver Address Comparison:")
    fmt.Printf("p1 == &p2: %v\n", (&p1).CompareAddressPointer(&p2)) // false（異なるインスタンス）

    // 同じアドレスを比較
    fmt.Printf("p1 == &p1: %v\n", (&p1).CompareAddressPointer(&p1)) // true（同じインスタンス）
}
```

Playgroundで実行してみてほしい。
出力は以下のようになるはず。(手元では検証済み)

```
Value Receiver Address Comparison:
p1 == p2: false
Pointer Receiver Address Comparison:
p1 == &p2: false
p1 == &p1: true
```

値レシーバでは、アドレスが異なっているのが確認できる。

- 値レシーバとポインタレシーバーのまとめ

|項目|	値レシーバ	|ポインタレシーバ|
|----|---|---|
|**渡されるもの**|構造体のコピー|構造体のポインタ|
|**元の値の変更**|できない|できる|
|**用途**|値を変更しない操作に適する|値を変更する操作に適する

※ポインタとは？（参考：実用Go言語P404ページ）
変数の値を格納するメモリ上の場所（アドレス）を指す。
アドレスを扱うのがポインタ。
- 変数のポインタ型にはアスタリスク（*）を前置する。
- 既存の変数のポインタを取り出すには＆を利用する
- ポインタから参照先の値を取り出す（dereference：デリファレンス. リファレンスをたどって指しているデータにアクセスすること[参考](https://gihyo.jp/dev/serial/01/perl-hackers-hub/002001#:~:text=%E3%83%AA%E3%83%95%E3%82%A1%E3%83%AC%E3%83%B3%E3%82%B9%E3%82%92%E3%81%9F%E3%81%A9%E3%81%A3%E3%81%A6%E6%8C%87%E3%81%97,%EF%BC%88dereference%EF%BC%89%E3%81%A8%E8%A8%80%E3%81%84%E3%81%BE%E3%81%99%E3%80%82)）には*を利用する


サンプルコード（実用Go言語記載コードをベースに拡張した）
```
package main

import "fmt"

func main() {
	var i int = 10
	var p *int
	p = &i
	fmt.Println(&i) // コピー元
	fmt.Println(p)  // アドレス
	fmt.Println(*p) // 格納された値
}
```

## if文

言語仕様は[ここ](https://go.dev/ref/spec#If_statements).
条件判定に使われる。  
書き方は以下の通り。  

- if

    ```
    if 条件 {
        // 条件がtrueの場合の処理
    }
    ```

    ```
    package main

    import "fmt"

    func main() {
        x := 10
        if x > 5 {
            fmt.Println("xは5より大きい")
        }
    }
    ```

- if else

    ```
    if 条件 {
        // 条件がtrueの場合の処理
    } else {
        // 条件がfalseの場合の処理
    }
    ```

    ```
    x := 3
    if x > 5 {
        fmt.Println("xは5より大きい")
    } else {
        fmt.Println("xは5以下")
    }
    ```

- if else if

    ```
    if 条件1 {
    // 条件1がtrueの場合の処理
    } else if 条件2 {
        // 条件2がtrueの場合の処理
    } else {
        // どちらの条件もfalseの場合の処理
    }
    ```

    ```
    x := 7
    if x > 10 {
        fmt.Println("xは10より大きい")
    } else if x > 5 {
        fmt.Println("xは5より大きいが、10以下")
    } else {
        fmt.Println("xは5以下")
    }
    ```

if文の中で変数宣言もできる。

```
if x := 5; x > 3 {
    fmt.Println("xは3より大きい")
} else {
    fmt.Println("xは3以下")
}
// ここでxを参照しようとするとエラーになる
```

エラーハンドリングなんかでよく使う。  
コードを簡潔にかけるメリット、err文を間違って上書きするミス、エラーハンドリングを忘れるミスを減らせる。  
```
if err := doFunc(); err != nil {
    // エラー時の処理　
}
```

## Switch

言語仕様は[ここ](https://go.dev/ref/spec#Switch_statements)
switch文は上から順に評価される

- 基本的な構文

```
switch 式 {
case 値1:
    // 値1に一致する場合の処理
case 値2:
    // 値2に一致する場合の処理
default:
    // どのケースにも一致しない場合の処理
}
```

-  条件判定付きのSwitch文

```
x := 7
switch {
case x > 10:
    fmt.Println("xは10より大きい")
case x > 5:
    fmt.Println("xは5より大きいが、10以下")
default:
    fmt.Println("xは5以下")
}

```

- 複数の値を使う

```
day := "土曜日"
switch day {
case "土曜日", "日曜日":
    fmt.Println("週末です")
default:
    fmt.Println("平日です")
}
```

- fallthrough

switchは一致したケースを実行したら終了するが、`fallthrough`を指定すると、マッチした以降のケースを実行する。
```
switch x := 2; x {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
    fallthrough
case 3:
    fmt.Println("Three")
}

```
## breakとcontinue

- break
  ループを終了させる。
  ```
  package main
    import "fmt"

    func main() {
        for i := 1; i <= 10; i++ {
            if i == 5 {
                fmt.Println("5でループ終了")
                break
            }
            fmt.Println(i)
        }
    }
  ```

- continue
  現在のループを中断して、次のループに行く
  ```
  package main

    import "fmt"

    func main() {
        for i := 1; i <= 10; i++ {
            if i%2 == 0 { // 偶数の場合はスキップ
                continue
            }
            fmt.Println(i) // 奇数のみ出力
        }
    }
  ```


## for文

言語仕様は[ここ](https://go.dev/ref/spec#For_statements).
Goの繰り返し構文はforのみ。

for文の条件判定は３つのパターンがある。
１つ目は条件判定がfalseになるまで、ループする構文。

```
for a < b {
  print(a)
}
```

無限ループもできる

```
a =: 0
for {
    a++
    if(a > 10) {
        break
    }
}
```
２つ目は値の増減と評価を含むfor文
iのスコープはfor文内だけ、ループを抜けたら無効になる
```
for i := 0; i < 10; i++ {
    print(i)
}
```

次のようにfor文の外で宣言して、for文のスコープで新たに宣言すると別物扱い。
```
package main

func main() {
	i := 999999999
	for i := 0; i < 10; i++ {
		print(i)
	}
	print("\n")
	print(i)
}

// 出力例
// 0123456789
// 999999999
```

ただし、スコープ内で宣言していないと、ループ外の変数が参照される。
```
package main

func main() {
	i := 999999999
	for j := 0; j < 10; j++ {
		print(i)
	}
	print("\n")
	print(i)
}
```
この仕様でバグを作り込むことがあるので、ループ外とループ内で宣言する変数名は変えたほうがいい。  
多重ループ内のエラーハンドリングが特にミスりやすい（実際にやったことある）  

３つ目の書き方はrangeを使う方法。  
Pythonっぽい感じでかける。  
以下サンプル。

```
package main

import "fmt"

func main() {
    nums := []int{1, 2, 3, 4, 5}
    for index, value := range nums {
        fmt.Printf("index: %d, value: %d\n", index, value)
    }
}
```


使える条件は[言語仕様](https://go.dev/ref/spec#For_range)にかいてある。  
以下の通り。  

- 配列、スライス、文字列、マップ

```
// 配列
arr := [3]int{1, 2, 3}
for i, v := range arr {
    fmt.Printf("Index: %d, Value: %d\n", i, v)
}

```

```
// スライス
slice := []string{"Go", "Python", "Rust"}
for i, v := range slice {
    fmt.Printf("Index: %d, Value: %s\n", i, v)
}

```

```
// 文字列
str := "Gopherkunn"
for i, v := range str {
    fmt.Printf("Index: %d, Rune: %c\n", i, v)
}
```

```
// map
m := map[string]int{"Alice": 25, "Bob": 30}
for k, v := range m {
    fmt.Printf("Key: %s, Value: %d\n", k, v)
}
```

- チャネルで受信した値

```
ch := make(chan int, 3)
ch <- 10
ch <- 20
ch <- 30
close(ch)

for v := range ch {
    fmt.Printf("Received: %d\n", v)
}

```

- ゼロから上限値までの整数値(Go 1.22以降)

```
for i := range 5 { // 0から4まで
    fmt.Printf("Value: %d\n", i)
}
```

- イテレータ関数のyield関数に渡された値(Go1.23以降)

ぶっちゃけ業務で使ったことない。
理解するまでクソむずいので一旦放置でもOK

```
package main

import (
	"fmt"
	"golang.org/x/exp/iter"
)

func main() {
	// yield を使ってカスタムイテレータを作成
	customIterator := iter.New(func(yield func(int)) {
		for i := 0; i < 10; i += 2 { // 0から10未満、ステップ2
			yield(i)
		}
	})

	// イテレータを反復処理
	for v := range customIterator {
		fmt.Printf("Yielded: %d\n", v)
	}
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

#### 補足：init関数
## net/http

## 並行処理

### gorutine


### channel

### select

評価順はランダム

## ジェネリクス

## イテレータ


## テスト

### Unitテスト
### 並行テスト
### runnを使ったテスト

## おまけ

### go embed

### Opentelemetry+clickhouse+Grafana

#### Opentlemetry

#### clickhouse

#### Grafana

## 参考文献
  - https://go.dev/doc/tutorial/getting-started
  - 実用Go言語
  - https://koko206.hatenablog.com/entry/2024/01/06/055112
