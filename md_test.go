package lark_docs_md_test

import (
	"io/ioutil"
	"testing"

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
	ioutil.WriteFile("testdata/1.md", []byte(res.Markdown()), 0o666)
	// fmt.Println(res.Markdown())
	as.Equal(res.Markdown(), string(bs2))
}
