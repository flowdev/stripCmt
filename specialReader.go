package main

import (
	"strings"
)

type SpecialReader interface {
	SpecialStart(line string, start int) int
	ReadSpecial(line string, start int, firstLine bool) (newLine string, restPos int, done bool)
}

type LineCommentSpecialReader struct {
}

func NewLineCommentSpecialReader() *LineCommentSpecialReader {
	return &LineCommentSpecialReader{}
}
func (lcsr *LineCommentSpecialReader) SpecialStart(line string, start int) int {
	return index(line, start, "//")
}
func (lcsr *LineCommentSpecialReader) ReadSpecial(line string, start int, firstLine bool) (newLine string, restPos int, done bool) {
	newLine = strings.Trim(line[0:start], "\r\n\t ")
	if len(newLine) <= 0 {
		return "", -1, true
	} else {
		return line[0:start], start, true
	}
}

type BlockCommentSpecialReader struct {
}

func NewBlockCommentSpecialReader() *BlockCommentSpecialReader {
	return &BlockCommentSpecialReader{}
}
func (bcsr *BlockCommentSpecialReader) SpecialStart(line string, start int) int {
	return index(line, start, "/*")
}
func (bcsr *BlockCommentSpecialReader) ReadSpecial(line string, start int, firstLine bool) (newLine string, restPos int, done bool) {
	pos := strings.Index(line[start:], "*/")
	if pos >= 0 {
		newLine = line[0:start] + line[start+pos+2:]
	} else if firstLine {
		newLine = line[0:start]
	} else {
		newLine = ""
	}
	tmpLine := strings.Trim(newLine, "\r\n\t ")
	if len(tmpLine) <= 0 {
		return "", -1, (pos>=0)
	} else {
		return newLine, start, (pos>=0)
	}
}

type SingleQuoteSpecialReader struct {
}

func NewSingleQuoteSpecialReader() *SingleQuoteSpecialReader {
	return &SingleQuoteSpecialReader{}
}
func (sqsr *SingleQuoteSpecialReader) SpecialStart(line string, start int) int {
	return index(line, start, "'")
}
func (sqsr *SingleQuoteSpecialReader) ReadSpecial(line string, start int, firstLine bool) (newLine string, restPos int, done bool) {
	return readSpecialQuote(line, start, firstLine, "'")
}

type DoubleQuoteSpecialReader struct {
}

func NewDoubleQuoteSpecialReader() *DoubleQuoteSpecialReader {
	return &DoubleQuoteSpecialReader{}
}
func (dqsr *DoubleQuoteSpecialReader) SpecialStart(line string, start int) int {
	return index(line, start, "\"")
}
func (dqsr *DoubleQuoteSpecialReader) ReadSpecial(line string, start int, firstLine bool) (newLine string, restPos int, done bool) {
	return readSpecialQuote(line, start, firstLine, "\"")
}

func readSpecialQuote(line string, start int, firstLine bool, quote string) (newLine string, restPos int, done bool) {
	oldPos := start
	if firstLine {
		oldPos += len(quote)
	}
	substr := line[oldPos:]
	pos := strings.IndexAny(substr, quote+"\\")
	for pos >= 0 && substr[pos] == '\\' {
		oldPos += pos + 2
		if len(substr) > pos+2 {
			substr = substr[pos+2:]
		} else {
			substr = ""
		}
		pos = strings.IndexAny(substr, quote+"\\")
	}
	if pos >= 0 {
		return line, oldPos + pos + 1, true
	} else {
		return line, len(line), false
	}
}

func index(line string, start int, substr string) int {
	pos := strings.Index(line[start:], substr)
	if pos >= 0 {
		return start + pos
	} else {
		return len(line)
	}
}

