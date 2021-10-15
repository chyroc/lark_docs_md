package lark_docs_md

import (
	"fmt"
	"strconv"
)

type Docs struct {
	Title Title `json:"title"`
	Body  Body  `json:"body"`
}

type Style struct {
	Quote        bool      `json:"quote"`
	HeadingLevel int       `json:"headingLevel"`
	List         StyleList `json:"list"`
	Link         struct {
		URL string `json:"url"`
	}
}

type StyleList struct {
	Type        string `json:"type"` // number bullet checkBox checkedBox
	IndentLevel int    `json:"indentLevel"`
	Number      int    `json:"number"`
}

func (r StyleList) ListTag() string {
	switch r.Type {
	case "number":
		return strconv.FormatInt(int64(r.Number), 10) + "."
	case "bullet":
		return "-"
	default:
		panic(fmt.Sprintf("style list"))
	}
}

type Location struct {
	ZoneID     string `json:"zoneId"`
	StartIndex int    `json:"startIndex"`
	EndIndex   int    `json:"endIndex"`
}

type TextRun struct {
	Text     string   `json:"text"`
	Style    Style    `json:"style"`
	Location Location `json:"location"`
}

type Element struct {
	Type     string   `json:"type"` // textRun
	TextRun  TextRun  `json:"textRun"`
	DocsLink DocsLink `json:"docsLink"`
}

type Elements []Element

type Title struct {
	Elements Elements `json:"elements"`
	Location Location `json:"location"`
	LineID   string   `json:"lineId"`
}

type Paragraph struct {
	Elements Elements `json:"elements"`
	Style    Style    `json:"style"`
	Location Location `json:"location"`
	LineID   string   `json:"lineId"`
}

type Color struct {
	Red   int `json:"red"`
	Green int `json:"green"`
	Blue  int `json:"blue"`
	Alpha int `json:"alpha"`
}

type Callout struct {
	Location               Location `json:"location"`
	CalloutEmojiID         string   `json:"calloutEmojiId"`
	CalloutBackgroundColor Color    `json:"calloutBackgroundColor"`
	CalloutBorderColor     Color    `json:"calloutBorderColor"`
	ZoneID                 string   `json:"zoneId"`
	Body                   Body     `json:"body"`
}

type HorizontalLine struct {
	Location Location `json:"location"`
}

type Code struct {
	Language string   `json:"language"`
	Location Location `json:"location"`
	ZoneID   string   `json:"zoneId"`
	Body     Body     `json:"body"`
}

type GalleryStyle struct{}

type Image struct {
	FileToken string `json:"fileToken"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
}

type Images []Image

type Gallery struct {
	GalleryStyle GalleryStyle `json:"galleryStyle"`
	ImageList    Images       `json:"imageList"`
	Location     Location     `json:"location"`
}

type TableCell struct {
	ZoneID      string `json:"zoneId"`
	ColumnIndex int    `json:"columnIndex"`
	Body        Body   `json:"body"`
}

type TableRow struct {
	RowIndex   int         `json:"rowIndex"`
	TableCells []TableCell `json:"tableCells"`
}

type TableColumnProperty struct {
	Width int `json:"width"`
}

type TableStyle struct {
	TableColumnProperties []TableColumnProperty `json:"tableColumnProperties"`
}

type Table struct {
	TableID     string        `json:"tableId"`
	RowSize     int           `json:"rowSize"`
	ColumnSize  int           `json:"columnSize"`
	TableRows   []TableRow    `json:"tableRows"`
	TableStyle  TableStyle    `json:"tableStyle"`
	MergedCells []interface{} `json:"mergedCells"`
	Location    Location      `json:"location"`
}

type Block struct {
	Type           string         `json:"type"`
	Callout        Callout        `json:"callout,omitempty"`
	HorizontalLine HorizontalLine `json:"horizontalLine,omitempty"`
	Code           Code           `json:"code,omitempty"`
	Gallery        Gallery        `json:"gallery,omitempty"`
	Table          Table          `json:"table,omitempty"`
	Paragraph      Paragraph      `json:"paragraph,omitempty"`
	Sheet          Sheet          `json:"sheet,omitempty"`
	ChatGroup      ChatGroup      `json:"chatGroup"`
	File           File           `json:"file"`
}

type Sheet struct {
	Token    string   `json:"token"`
	Location Location `json:"location"`
}

type DocsLink struct {
	URL      string   `json:"url"`
	Location Location `json:"location"`
}

type ChatGroup struct {
	OpenChatID string   `json:"openChatId"`
	Location   Location `json:"location"`
}

type File struct {
	FileToken string   `json:"fileToken"`
	ViewType  string   `json:"viewType"`
	FileName  string   `json:"fileName"`
	Location  Location `json:"location"`
}

type Blocks []Block

type Body struct {
	Blocks Blocks `json:"blocks"`
}
