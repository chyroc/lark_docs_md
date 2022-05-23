package lark_docs_md

import (
	"context"
	"fmt"
	"strings"

	"github.com/chyroc/lark"
)

type FormatOpt struct {
	ctx         context.Context
	LarkClient  *lark.Lark // lark 客户端
	StaticDir   string     // 如果需要下载静态文件，那么需要指定静态文件的目录
	FilePrefix  string     // 针对静态文件，需要指定文件在 Markdown 中的前缀
	StaticAsURL bool       // 不下载静态文件，直接把静态文件的 URL 插入到 Markdown 中
}

func DocMarkdown(ctx context.Context, doc *lark.DocContent, opt *FormatOpt) string {
	if opt == nil {
		opt = new(FormatOpt)
	}
	opt.ctx = ctx
	if opt.StaticDir == "" {
		opt.StaticDir = "static"
	}
	if opt.FilePrefix == "" {
		opt.FilePrefix = "static"
	}

	buf := new(strings.Builder)

	buf.WriteString("# ")
	buf.WriteString(DocParagraphMarkdown(doc.Title, opt))
	buf.WriteString("\n\n")

	buf.WriteString(DocBodyMarkdown(doc.Body, opt))

	return buf.String()
}

func DocBodyMarkdown(r *lark.DocBody, opt *FormatOpt) string {
	buf := new(strings.Builder)

	for idx, v := range r.Blocks {
		buf.WriteString(DocBlockMarkdown(v, opt))
		buf.WriteString("\n")

		// 如果 v 是 list，并且是最后一个，那么需要换行
		if idx == len(r.Blocks)-1 && v.Type == lark.DocBlockTypeParagraph && v.Paragraph.Style != nil && v.Paragraph.Style.List != nil && v.Paragraph.Style.List.IndentLevel > 0 {
			buf.WriteString("\n")
		}
	}

	return buf.String()
}

func DocBlockMarkdown(r *lark.DocBlock, opt *FormatOpt) string {
	switch r.Type {
	case "paragraph":
		return DocParagraphMarkdown(r.Paragraph, opt)
	case "gallery":
		return DocGalleryMarkdown(r.Gallery, opt)
	case "file":
		return DocFileMarkdown(r.File, opt)
	case "chatGroup":
		return DocChatGroupMarkdown(r.ChatGroup, opt)
	case "table":
		return DocTableMarkdown(r.Table, opt)
	case "horizontalLine":
		return DocHorizontalLineMarkdown(r.HorizontalLine, opt)
	case "embeddedPage":
		return DocEmbeddedPageMarkdown(r.EmbeddedPage, opt)
	case "sheet":
		return DocSheetMarkdown(r.Sheet, opt)
	case "bitable":
		return DocBitableMarkdown(r.Bitable, opt)
	case "diagram":
		return DocDiagramMarkdown(r.Diagram, opt)
	case "jira":
		return DocJiraMarkdown(r.Jira, opt)
	case "poll":
		return DocPollMarkdown(r.Poll, opt)
	case "code":
		return DocCodeMarkdown(r.Code, opt)
	case "docsApp":
		return DocDocsAppMarkdown(r.DocsApp, opt)
	case "callout":
		return DocCalloutMarkdown(r.Callout, opt)
	case "undefinedBlock":
		return DocUndefinedBlockMarkdown(r.UndefinedBlock, opt)
	default:
		return fmt.Sprintf("<!-- unknown block type %s -->", r.Type)
	}
}

func DocParagraphMarkdown(r *lark.DocParagraph, opt *FormatOpt) string {
	buf := new(strings.Builder)

	if r.Style != nil {
		if r.Style.HeadingLevel > 0 {
			buf.WriteString(strings.Repeat("#", int(r.Style.HeadingLevel)))
			buf.WriteString(" ")
		}
		if r.Style.List != nil {
			buf.WriteString(r.Style.List.ListTag())
			buf.WriteString(" ")
		}
		switch {
		case r.Style.Quote:
			buf.WriteString("> ")
		}
	}

	for _, v := range r.Elements {
		buf.WriteString(DocParagraphElementMarkdown(v, opt))
	}

	if r.Style != nil {
		if r.Style.List != nil && r.Style.List.IndentLevel > 0 {
			//
		} else {
			buf.WriteString("\n")
		}
	}

	return buf.String()
}

func DocGalleryMarkdown(r *lark.DocGallery, opt *FormatOpt) string {
	buf := new(strings.Builder)

	for _, v := range r.ImageList {
		buf.WriteString(DocImageItemMarkdown(v, opt))
	}

	return buf.String()
}

func DocImageItemMarkdown(r *lark.DocImageItem, opt *FormatOpt) string {
	if opt.LarkClient == nil {
		return fmt.Sprintf("`[image: %s]`", r.FileToken)
	}

	target := downloadFile(opt.ctx, r.FileToken, r.FileToken+".jpg", opt)
	return fmt.Sprintf("<img src=%q width=\"%d\" height=\"%d\"/>", target, r.Width, r.Height)
}

func DocChatGroupMarkdown(r *lark.DocChatGroup, opt *FormatOpt) string {
	return fmt.Sprintf("`[chat: %s]`", r.OpenChatID)
}

func DocTableMarkdown(r *lark.DocTable, opt *FormatOpt) string {
	// todo
	buf := new(strings.Builder)

	return buf.String()
}

func DocHorizontalLineMarkdown(r *lark.DocHorizontalLine, opt *FormatOpt) string {
	return "\n---\n"
}

func DocEmbeddedPageMarkdown(r *lark.DocEmbeddedPage, opt *FormatOpt) string {
	return fmt.Sprintf("[embedded-page: %s](%s)", r.Type, r.Url)
}

func DocSheetMarkdown(r *lark.DocSheet, opt *FormatOpt) string {
	return fmt.Sprintf("`[sheet: %s]`", r.Token)
}

func DocBitableMarkdown(r *lark.DocBitable, opt *FormatOpt) string {
	return fmt.Sprintf("`[bitable: %s / %s]`", r.ViewType, r.Token)
}

func DocDiagramMarkdown(r *lark.DocDiagram, opt *FormatOpt) string {
	return fmt.Sprintf("`[diagram: %s / %s]`", r.DiagramType, r.Token)
}

func DocPollMarkdown(r *lark.DocPoll, opt *FormatOpt) string {
	return fmt.Sprintf("`[poll: %s]`", r.Token)
}

func DocCodeMarkdown(r *lark.DocCode, opt *FormatOpt) string {
	buf := new(strings.Builder)

	buf.WriteString("```")
	if r.Language != "" {
		buf.WriteString(r.Language)
	}
	buf.WriteString("\n")

	buf.WriteString(DocBodyMarkdown(r.Body, opt))

	buf.WriteString("\n")
	buf.WriteString("```")

	return buf.String()
}

func DocDocsAppMarkdown(r *lark.DocDocsApp, opt *FormatOpt) string {
	return fmt.Sprintf("`[docs-app: %s]`", r.TypeID)
}

func DocCalloutMarkdown(r *lark.DocCallout, opt *FormatOpt) string {
	buf := new(strings.Builder)

	buf.WriteString("```")
	buf.WriteString("\n")

	buf.WriteString(DocBodyMarkdown(r.Body, opt))

	buf.WriteString("\n")
	buf.WriteString("```")

	return buf.String()
}

func DocUndefinedBlockMarkdown(r *lark.DocUndefinedBlock, opt *FormatOpt) string {
	return "`[undefined-block]`"
}

func DocParagraphElementMarkdown(r *lark.DocParagraphElement, opt *FormatOpt) string {
	switch r.Type {
	case "textRun":
		return DocTextRunMarkdown(r.TextRun, opt)
	case "docsLink":
		return DocDocsLinkMarkdown(r.DocsLink, opt)
	case "person":
		return DocPersonMarkdown(r.Person, opt)
	case "equation":
		return DocEquationMarkdown(r.Equation, opt)
	case "reminder":
		return DocReminderMarkdown(r.Reminder, opt)
	case "file":
		return DocFileMarkdown(r.File, opt)
	case "jira":
		return DocJiraMarkdown(r.Jira, opt)
	case "undefinedElement":
		return DocUndefinedElementMarkdown(r.UndefinedElement, opt)
	default:
		return fmt.Sprintf("<!-- unknown doc paragrapg element type %s -->", r.Type)
	}
}

func DocTextRunMarkdown(r *lark.DocTextRun, opt *FormatOpt) string {
	// **加粗**
	// _斜体_
	// ~~删除线~~
	// <u>下划线</u>
	// `行内代码`
	// [title](url)
	buf := new(strings.Builder)
	if r.Style != nil {
		switch {
		case r.Style.Bold:
			buf.WriteString("**")
		case r.Style.Italic:
			buf.WriteString("_")
		case r.Style.StrikeThrough:
			buf.WriteString("~~")
		case r.Style.Underline:
			buf.WriteString("<u>")
		case r.Style.CodeInline:
			buf.WriteString("`")
		case r.Style.Link != nil && r.Style.Link.URL != "":
			buf.WriteString("[")
		}
	}

	buf.WriteString(r.Text)

	if r.Style != nil {
		switch {
		case r.Style.Bold:
			buf.WriteString("**")
		case r.Style.Italic:
			buf.WriteString("_")
		case r.Style.StrikeThrough:
			buf.WriteString("~~")
		case r.Style.Underline:
			buf.WriteString("</u>")
		case r.Style.CodeInline:
			buf.WriteString("`")
		case r.Style.Link != nil && r.Style.Link.URL != "":
			buf.WriteString("](")
			buf.WriteString(r.Style.Link.URL)
			buf.WriteString(")")
		}
	}

	return buf.String()
}

func DocDocsLinkMarkdown(r *lark.DocDocsLink, opt *FormatOpt) string {
	return fmt.Sprintf("[%s](%s)", "云文档", r.URL)
}

func DocPersonMarkdown(r *lark.DocPerson, opt *FormatOpt) string {
	return "@" + r.OpenID
}

func DocEquationMarkdown(r *lark.DocEquation, opt *FormatOpt) string {
	return fmt.Sprintf("$$%s$$", r.Equation)
}

func DocReminderMarkdown(r *lark.DocReminder, opt *FormatOpt) string {
	return "`[doc-reminder]`"
}

func DocFileMarkdown(r *lark.DocFile, opt *FormatOpt) string {
	if opt.LarkClient == nil {
		return fmt.Sprintf("`[file: %s / %s]`", r.FileName, r.FileToken)
	}

	target := downloadFile(opt.ctx, r.FileToken, r.FileName, opt)
	return fmt.Sprintf("[file: %s](%s)", r.FileName, target)
}

func DocJiraMarkdown(r *lark.DocJira, opt *FormatOpt) string {
	return fmt.Sprintf("`[jira: %s / %s]`", r.JiraType, r.Token)
}

func DocUndefinedElementMarkdown(r *lark.DocUndefinedElement, opt *FormatOpt) string {
	return "`[undefined-element]`"
}
