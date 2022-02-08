package lark_docs_md_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/chyroc/lark"
	"github.com/chyroc/lark/larkext"
	"github.com/chyroc/lark_docs_md"
	"github.com/stretchr/testify/assert"
)

func TestDocMarkdown(t *testing.T) {
	as := assert.New(t)
	larkClient := lark.New(
		lark.WithAppCredential(os.Getenv("LARK_CHYROC_HEYMAN_APP_ID"), os.Getenv("LARK_CHYROC_HEYMAN_APP_SECRET")),
		lark.WithTimeout(time.Minute),
	)

	tests := []struct {
		name string
	}{
		{"1"},
		{"2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt := &lark_docs_md.FormatOpt{
				LarkClient: larkClient,
				StaticDir:  fmt.Sprintf("testdata/%s/static", tt.name),
				FilePrefix: "static",
			}

			doc := loadDocContent(t, tt.name)
			md := loadMarkdown(t, tt.name)
			result := lark_docs_md.DocMarkdown(context.Background(), doc, opt)
			as.Equal(md, result, tt.name)
		})
	}
}

func Test_FromDocToken(t *testing.T) {
	t.Skip()

	as := assert.New(t)

	// 	https://rs6qnacjws.feishu.cn/docs/doccng7C51ULsOrHtbmWBa3a4Hf
	larkClient := lark.New(
		lark.WithAppCredential(os.Getenv("LARK_CHYROC_HEYMAN_APP_ID"), os.Getenv("LARK_CHYROC_HEYMAN_APP_SECRET")),
	)
	docToken := "doccng7C51ULsOrHtbmWBa3a4Hf"

	doc, err := larkext.NewDoc(larkClient, docToken).Content(context.Background())
	as.Nil(err)

	result := lark_docs_md.DocMarkdown(context.Background(), doc, &lark_docs_md.FormatOpt{
		LarkClient: larkClient,
		// StaticDir:  "/tmp/tmp-static",
		// FilePrefix: "static",
		StaticAsURL: true,
	})
	fmt.Println(result)
}

func loadDocContent(t *testing.T, testCase string) *lark.DocContent {
	as := assert.New(t)

	bs, err := ioutil.ReadFile(fmt.Sprintf("testdata/%s/data.json", testCase))
	as.Nil(err)

	doc := &lark.DocContent{}
	as.Nil(json.Unmarshal(bs, doc))

	return doc
}

func loadMarkdown(t *testing.T, testCase string) string {
	as := assert.New(t)

	bs, err := ioutil.ReadFile(fmt.Sprintf("testdata/%s/data.md", testCase))
	as.Nil(err)

	return string(bs)
}
