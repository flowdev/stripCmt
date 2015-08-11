package main

import (
	"io"
	"os"
	"fmt"
)

type SpecialReaderManager struct {
	lr    LineReader
	srs   []SpecialReader
	cursr SpecialReader
}

func main() {
	lr := NewFormatter(NewStripCmtLineReader(NewSaneLineReader(os.Stdin)))
	line, err := lr.ReadLine()
	for ; err == nil; line, err = lr.ReadLine() {
		fmt.Println(line)
	}
	if err != io.EOF {
		fmt.Fprintln(os.Stderr, "ERROR:", err.Error())
		os.Exit(1)
	}
}

func NewSpecialReaderManager(lr LineReader, srs ...SpecialReader) *SpecialReaderManager {
	return &SpecialReaderManager{lr, srs, nil}
}
func NewStripCmtLineReader(lr LineReader) *SpecialReaderManager {
	return &SpecialReaderManager{lr, []SpecialReader{NewSingleQuoteSpecialReader(), NewDoubleQuoteSpecialReader(), NewLineCommentSpecialReader(), NewBlockCommentSpecialReader()}, nil}
}
func (srm *SpecialReaderManager) ReadLine() (line string, err error) {
	for deleted := true; err == nil && deleted; {
		line, err = srm.lr.ReadLine()
		if err == nil || (err == io.EOF && len(line) > 0) {
			line, deleted = srm.readSpecial(line)
		} else {
			deleted = false
		}
	}
	return line, err
}
func (srm *SpecialReaderManager) readSpecial(line string) (newLine string, deleted bool) {
	restPos := 0
	firstLine := false
	done := false

	for restPos < len(line) {
		if srm.cursr == nil {
			firstLine = true
			srm.cursr, restPos = firstSpecialReader(line, restPos, srm.srs)
		} else {
			line, restPos, done = srm.cursr.ReadSpecial(line, restPos, firstLine)
			if restPos < 0 {
				deleted = true
				restPos = len(line)
			}
			if done {
				srm.cursr = nil
			}
		}
	}
	return line, deleted
}

func firstSpecialReader(line string, start int, srs []SpecialReader) (firstsr SpecialReader, pos int) {
	min := len(line)
	for i := 0; i < len(srs) && min > start; i++ {
		pos = srs[i].SpecialStart(line, start)
		if pos < min {
			firstsr = srs[i]
			min = pos
		}
	}
	return firstsr, min
}
