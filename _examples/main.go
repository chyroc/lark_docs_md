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
