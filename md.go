package lark_docs_md

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/chyroc/lark"
)

func Unmarshal(content string) (*Docs, error) {
	res := new(Docs)
	err := json.Unmarshal([]byte(content), res)
	return res, err
}

func (r *Docs) Markdown(cli *lark.Lark, dir, staticPrefix string) string {
	opt := &formatOpt{cli: cli, dir: dir, staticPrefix: staticPrefix}
	parchOpt(r, opt)

	buf := strings.Builder{}
	buf.WriteString(r.Title.Markdown())
	buf.WriteString("\n")
	buf.WriteString(r.Body.Markdown(false))
	return buf.String()
}

func (r *Style) Markdown() string {
	if r == nil {
		return ""
	}
	s := ""
	if r.Quote {
		s = ">" + s
	}
	for i := r.HeadingLevel; i >= 1; i-- {
		s = "#" + s
	}
	// r.List.Type

	if r.List != nil && r.List.Type != "" {
		switch r.List.Type {
		case "checkBox":
			s = "- [ ]" + s
		case "checkedBox":
			s = "- [x]" + s
		case "number":
			s = r.List.ListTag() + s
		case "bullet":
			s = r.List.ListTag() + s
		default:
			panic(fmt.Sprintf("style list: %s", r.List.Type))
		}
	}
	if r.List != nil {
		for i := r.List.IndentLevel; i >= 2; i-- {
			s = "   " + s
		}
	}
	if s == "" {
		return s
	}
	return s + ""
}

func (r *TextRun) Markdown() string {
	if r.Style.Link.URL != "" {
		x, _ := url.QueryUnescape(r.Style.Link.URL)
		return fmt.Sprintf("[%s](%s)", r.Text, x)
	}
	s := r.Style.Markdown()
	if s == "" {
		return r.Text
	}
	return s + " " + r.Text
}

func (r *Element) Markdown() string {
	switch r.Type {
	case "textRun":
		return r.TextRun.Markdown()
	case "docsLink":
		return r.DocsLink.Markdown()
	default:
		panic(fmt.Sprintf("Element type %s", r.Type))
	}
}

func (r *Elements) Markdown() string {
	buf := strings.Builder{}

	for idx, v := range *r {
		if idx > 0 {
			buf.WriteString("\n")
		}
		buf.WriteString(v.Markdown())
	}
	return buf.String()
}

func (r *Title) Markdown() string {
	return r.Elements.Markdown()
}

func (r *Paragraph) Markdown(block bool) string {
	p := r.Elements.Markdown()
	s := r.Style.Markdown()
	if s == "" {
		if !block {
			return p + "\n"
		}
		return p
	}
	return s + " " + p
}

func (r *Callout) Markdown() string {
	return "```\n" + r.Body.Markdown(true) + "\n```"
}

func (r *Code) Markdown() string {
	buf := strings.Builder{}
	buf.WriteString("```")
	buf.WriteString(r.Language)
	buf.WriteString("\n")
	buf.WriteString(r.Body.Markdown(true))
	buf.WriteString("\n```")
	return buf.String()
}

func (r *Image) Markdown() string {
	path := saveMedia(r.opt, r.FileToken, ".png")
	return fmt.Sprintf("<img src=\"%s\" width=%d height=%d>", path, r.Width, r.Height)
}

func (r *Images) Markdown() string {
	buf := strings.Builder{}

	for idx, v := range *r {
		if idx > 0 {
			buf.WriteString(" ")
		}
		buf.WriteString(v.Markdown())
	}
	return buf.String()
}

func (r *Gallery) Markdown() string {
	return r.ImageList.Markdown()
}

func (r *Table) Markdown() string {
	return "table"
}

func (r *Block) Markdown(block bool) string {
	switch r.Type {
	case "callout":
		return r.Callout.Markdown()
	case "horizontalLine":
		return "\n---"
	case "code":
		return r.Code.Markdown()
	case "gallery":
		return r.Gallery.Markdown()
	case "table":
		return r.Table.Markdown()
	case "paragraph":
		return r.Paragraph.Markdown(block)
	case "sheet":
		return r.Sheet.Markdown()
	case "undefinedBlock":
		return fmt.Sprintf("不支持的飞书文档组件: 未知组件")
	case "chatGroup":
		return r.ChatGroup.Markdown()
	case "file":
		return r.File.Markdown()
	default:
		panic(fmt.Sprintf("block: %s", r.Type))
	}
}

func (r *Sheet) Markdown() string {
	return "sheet"
}

func (r *DocsLink) Markdown() string {
	return r.URL
}

func (r *ChatGroup) Markdown() string {
	return fmt.Sprintf("不支持的飞书文档组件: 群名片")
}

func (r *File) Markdown() string {
	path := saveMedia(r.opt, r.FileToken, filepath.Ext(r.FileName))
	return fmt.Sprintf("[%s](%s)", r.FileName, path)
}

func (r *Blocks) Markdown(block bool) string {
	buf := strings.Builder{}
	for idx, v := range *r {
		if idx > 0 {
			buf.WriteString("\n")
			if v.Paragraph != nil && v.Paragraph.Style != nil {
				if v.Paragraph.Style.Quote && (*r)[idx-1].Paragraph.Style.Quote {
					buf.WriteString(">\n")
				}
			}

		}
		if v.Paragraph != nil && v.Paragraph.Style != nil && v.Paragraph.Style.HeadingLevel >= 1 && !strings.HasSuffix(buf.String(), "\n\n") {
			buf.WriteString("\n")
		}
		buf.WriteString(v.Markdown(block))
	}
	return buf.String()
}

func (r *Body) Markdown(block bool) string {
	return r.Blocks.Markdown(block)
}

func saveMedia(opt *formatOpt, fileToken string, ext string) string {
	res, _, err := opt.cli.Drive.DownloadDriveMedia(context.Background(), &lark.DownloadDriveMediaReq{
		FileToken: fileToken,
	})
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(opt.dir, 0o777)
	if err != nil {
		panic(err)
	}

	storePath := strings.TrimRight(opt.dir, "/") + "/" + fileToken + ext
	mdPath := opt.staticPrefix + fileToken + ext
	f, err := os.OpenFile(storePath, os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.File)

	return mdPath
}
