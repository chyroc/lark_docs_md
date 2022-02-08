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
