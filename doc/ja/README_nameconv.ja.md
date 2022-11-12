[![UnitTest](https://github.com/nao1215/nameconv/actions/workflows/unit_test.yml/badge.svg)](https://github.com/nao1215/nameconv/actions/workflows/unit_test.yml)
[![reviewdog](https://github.com/nao1215/nameconv/actions/workflows/review_dog.yml/badge.svg)](https://github.com/nao1215/nameconv/actions/workflows/review_dog.yml)
[![codecov](https://codecov.io/gh/nao1215/nameconv/branch/main/graph/badge.svg?token=DO641K2SOW)](https://codecov.io/gh/nao1215/nameconv)
[![Go Reference](https://pkg.go.dev/badge/github.com/nao1215/nameconv.svg)](https://pkg.go.dev/github.com/nao1215/nameconv)
[![Go Report Card](https://goreportcard.com/badge/github.com/nao1215/nameconv)](https://goreportcard.com/report/github.com/nao1215/nameconv)  

# nameconv
**nameconv**は、文字列の命名規則を変換するライブラリです。
    
## コンバート関数
下表に示す関数は、文字列を指定の命名規則に変換します。
| 関数|説明|
|:--|:--|
|ToSnakeCase()| 文字列をスネークケース（例：snake_case）に変換|
|ToChainCase()| 文字列をチェインケース（例：chain-case）に変換|
|ToKebabCase()| チェインケースと同じ|
|ToCamelCase()| 文字列をキャメルケース（例：camelCase）に変換|
|ToPascalCase()| 文字列をパスカルケース（例：PascalCase）に変換|
|ToFlatCase()| 文字列をフラットケース（例：flatcase）に変換|
|ToUpperCase()| 文字列をアッパーケース（例：UPPER_CASE）に変換|  
    
## 命名規則チェック
下表に示す関数は、文字列が指定の命名規則かどうかを返します。期待どおりであればtrueを返し、それ以外の場合はfalseを返します。
| 関数|説明|
|:--|:--|
|IsSnakeCase()| スネークケースかどうかをチェック|
|IsChainCase()| チェインケースかどうかをチェック|
|IsKebabCase()| IsChainCase()と同じ|
|IsCamelCase()| キャメルケースかどうかをチェック |
|IsPascalCase()| パスカルケースかどうかをチェック |
|IsFlatCase()| フラットケースかどうかをチェック|
|IsUpperCase()| アッパーケースかどうかをチェック |  

IsCamelCase(), IsPascalCase(), IsFlatCase()は、最初の文字が数値の場合、必ずfalseを返します。数値の場合は、大文字／小文字の区別が出来ないからです。

# 実装例
```
package main

import (
	"fmt"

	"github.com/nao1215/nameconv"
)

func main() {
	fmt.Println(nameconv.ToSnakeCase("dave mustaine"))            // dave_mustaine
	fmt.Println(nameconv.ToChainCase("TifaLockhart"))             // tifa-lockhart
	fmt.Println(nameconv.ToKebabCase("The_Shawshank_Redemption")) // the-shawshank-redemption
	fmt.Println(nameconv.ToCamelCase("master-of_puppets"))        // masterOfPuppets
	fmt.Println(nameconv.ToPascalCase("Miracles_outOf-nowhere"))  // MiraclesOutOfNowhere
	fmt.Println(nameconv.ToFlatCase("SolBadguy"))                 // Sol badguy
	fmt.Println(nameconv.ToUpperCase("heaven and hell"))          // HEAVEN_AND_HELL

	fmt.Println(nameconv.IsSnakeCase("snake_case"))  // true
	fmt.Println(nameconv.IsChainCase("chain-case"))  // true
	fmt.Println(nameconv.IsKebabCase("kebab case"))  // false
	fmt.Println(nameconv.IsCamelCase("CamelCase"))   // false
	fmt.Println(nameconv.IsPascalCase("pascalCase")) // false
	fmt.Println(nameconv.IsFlatCase("flat-case"))    // false
	fmt.Println(nameconv.IsUpperCase("0_UPPER"))     // true
}
```
  
# 連絡先
開発者に対して「バグ報告」や「機能の追加要望」がある場合は、コメントをください。その際、以下の連絡先を使用してください。

[GitHub Issue](https://github.com/nao1215/gup/nameconv)

# ライセンス
nameconvライブラリは、混合ライセンスです。

- MIT License（casee*.go および camelcase*.go）
- Apache License Version 2.0（上記のコード以外全て）

MITライセンス部分の作者は、[pinzolo氏](https://github.com/pinzolo)および[Fatih Arslan氏](https://github.com/fatih)です。当該コードには、MIT License条文およびCopyrightを明記しています。
