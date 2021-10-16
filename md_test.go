package lark_docs_md_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/chyroc/lark"
	"github.com/chyroc/lark_docs_md"
	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	as := assert.New(t)

	bs, err := ioutil.ReadFile("testdata/1.json")
	as.Nil(err)
	bs2, err := ioutil.ReadFile("testdata/1.md")
	_ = bs2
	as.Nil(err)
	res, err := lark_docs_md.Unmarshal(string(bs))
	as.Nil(err)
	cli := lark.New(
		lark.WithAppCredential(os.Getenv("LARK_CHYROC_HEYMAN_APP_ID"), os.Getenv("LARK_CHYROC_HEYMAN_APP_SECRET")),
	)
	data := res.Markdown(cli, "testdata/1-static/", "./1-static/")
	ioutil.WriteFile("testdata/1.md", []byte(data), 0o666)
	// fmt.Println(res.Markdown())
	as.Equal(data, string(bs2))
}
