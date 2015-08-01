package main

import (
	"strings"
)

type SpecialReader interface {
	ConstStart() string
	ReadSpecial(line string, firstLine bool) (substring string, restPos int, done bool)
}

type LineCommentSpecialReader struct {
}

func NewLineCommentSpecialReader() *LineCommentSpecialReader {
	return &LineCommentSpecialReader{}
}
func (lcsr *LineCommentSpecialReader) ConstStart() string {
	return "//"
}
func (lcsr *LineCommentSpecialReader) ReadSpecial(line string, firstLine bool) (substring string, restPos int, done bool) {
	return "", 0, true
}

type BlockCommentSpecialReader struct {
}

func NewBlockCommentSpecialReader() *BlockCommentSpecialReader {
	return &BlockCommentSpecialReader{}
}
func (bcsr *BlockCommentSpecialReader) ConstStart() string {
	return "/*"
}
func (bcsr *BlockCommentSpecialReader) ReadSpecial(line string, firstLine bool) (substring string, restPos int, done bool) {
	pos := strings.Index(line, "*/")
	if pos >= 0 {
		return line[pos+2:], 0, true
	} else if firstLine {
		return "", 0, false
	} else {
		return "", -1, false
	}
}

type SingleQuoteSpecialReader struct {
}

func NewSingleQuoteSpecialReader() *SingleQuoteSpecialReader {
	return &SingleQuoteSpecialReader{}
}
func (sqsr *SingleQuoteSpecialReader) ConstStart() string {
	return "'"
}
func (sqsr *SingleQuoteSpecialReader) ReadSpecial(line string, firstLine bool) (substring string, restPos int, done bool) {
	return ReadSpecialQuote(line, firstLine, sqsr.ConstStart())
}

type DoubleQuoteSpecialReader struct {
}

func NewDoubleQuoteSpecialReader() *DoubleQuoteSpecialReader {
	return &DoubleQuoteSpecialReader{}
}
func (dqsr *DoubleQuoteSpecialReader) ConstStart() string {
	return "\""
}
func (dqsr *DoubleQuoteSpecialReader) ReadSpecial(line string, firstLine bool) (substring string, restPos int, done bool) {
	return ReadSpecialQuote(line, firstLine, dqsr.ConstStart())
}

func ReadSpecialQuote(line string, firstLine bool, quote string) (substring string, restPos int, done bool) {
	orgLine := line
	orgPos := 0
	substr := line
	if firstLine {
		substr = line[len(quote):]
		orgPos = len(quote)
	}
	pos := strings.IndexAny(substr, quote+"\\")
	for pos >= 0 && substr[pos] == '\\' {
		orgPos += pos + 2
		substr = substr[pos+2:]
		pos = strings.IndexAny(substr, quote+"\\")
	}
	if pos >= 0 {
		return orgLine, orgPos + pos + 1, true
	} else {
		return orgLine, len(orgLine), false
	}
}
