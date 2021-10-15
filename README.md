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
	"fmt"

	"github.com/chyroc/lark_docs_md"
)

func main() {
	res, err := lark_docs_md.Unmarshal(content)
	if err != nil {
		panic(err)
    }
	
	fmt.Println(res.Markdown())
}
```
