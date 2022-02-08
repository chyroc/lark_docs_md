# lark_docs_md

[![codecov](https://codecov.io/gh/chyroc/lark_docs_md/branch/master/graph/badge.svg?token=Z73T6YFF80)](https://codecov.io/gh/chyroc/lark_docs_md)
[![go report card](https://goreportcard.com/badge/github.com/chyroc/lark_docs_md "go report card")](https://goreportcard.com/report/github.com/chyroc/lark_docs_md)
[![test status](https://github.com/chyroc/lark_docs_md/actions/workflows/test.yml/badge.svg)](https://github.com/chyroc/lark_docs_md/actions)
[![Apache-2.0 license](https://img.shields.io/badge/License-Apache%202.0-brightgreen.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/chyroc/lark_docs_md)
[![Go project version](https://badge.fury.io/go/github.com%2Fchyroc%2Flark_docs_md.svg)](https://badge.fury.io/go/github.com%2Fchyroc%2Flark_docs_md)

![](./header.png)

## Install

```shell
go get github.com/chyroc/lark_docs_md
```

## Usage

```go
package main

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
	"github.com/chyroc/lark/larkext"
	"github.com/chyroc/lark_docs_md"
)

func main() {
	larkClient := lark.New(lark.WithAppCredential("app-id", "app-secret"))
	docToken := "doc-token"

	// 这一步是获取 doc 内容
	doc, err := larkext.NewDoc(larkClient, docToken).Content(context.Background())
	if err != nil {
		panic(err)
	}

	// 转化为 markdown
	result := lark_docs_md.DocMarkdown(context.Background(), doc, &lark_docs_md.FormatOpt{
		LarkClient: larkClient,

		// 如果需要下载图片等静态文件，请配置这两项
		// StaticDir:  "static",
		// FilePrefix: "static",

		// 如果不需要下载文件，而替换为 24 小时有效的链接，请配置这个项
		StaticAsURL: true,
	})

	// 输出
	fmt.Println(result)
}
```
