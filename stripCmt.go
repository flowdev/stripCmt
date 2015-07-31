package main

import (
	"io"
	"strings"
)

type SpecialReaderManager struct {
	lr    LineReader
	srs   []SpecialReader
	cursr SpecialReader
}

func NewSpecialReaderManager(lr LineReader, srs ...SpecialReader) *SpecialReaderManager {
	return &SpecialReaderManager{lr, srs, nil}
}
func NewStripCmtLineReader(lr LineReader) *SpecialReaderManager {
	return &SpecialReaderManager{lr, []SpecialReader{NewSingleQuoteSpecialReader(), NewDoubleQuoteSpecialReader(), NewLineCommentSpecialReader(), NewBlockCommentSpecialReader()}, nil}
}
func (srm *SpecialReaderManager) ReadLine() (line string, err error) {
	line, err = srm.lr.ReadLine()
	if err == nil || (err == io.EOF && len(line) > 0) {
		line = srm.readSpecial(line)
	}
	return line, err
}
func (srm *SpecialReaderManager) readSpecial(line string) string {
	result := ""
	rest := line
	firstLine := false

	for len(rest) > 0 {
		if srm.cursr == nil {
			firstLine = true
			firstsr, pos := firstSpecialReader(rest, srm.srs)
			result += rest[0:pos]
			rest = rest[pos:]
			srm.cursr = firstsr
		} else {
			substr, restPos, done := srm.cursr.ReadSpecial(rest, firstLine)
			if restPos > 0 {
				result += substr[0:restPos]
				rest = substr[restPos:]
			} else {
				rest = substr
			}
			if done {
				srm.cursr = nil
			}
		}
	}
	return result
}

func firstSpecialReader(line string, srs []SpecialReader) (firstsr SpecialReader, pos int) {
	min := len(line)
	for i := 0; i < len(srs) && min > 0; i++ {
		pos = strings.Index(line, srs[i].ConstStart())
		if pos >= 0 && pos < min {
			firstsr = srs[i]
			min = pos
		}
	}
	return firstsr, min
}
