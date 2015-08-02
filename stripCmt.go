package main

import (
	"io"
	"strings"
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
func (srm *SpecialReaderManager) readSpecial(line string) (substr string, deleted bool) {
	substr = ""
	rest := line
	firstLine := false
	i := 0

	for len(rest) > 0 {
		if srm.cursr == nil {
			firstLine = true
			firstsr, pos := firstSpecialReader(rest, srm.srs)
			substr += rest[0:pos]
			rest = rest[pos:]
			srm.cursr = firstsr
		} else {
			subspec, restPos, done := srm.cursr.ReadSpecial(rest, firstLine)
			switch {
			case restPos > 0:
				substr += subspec[0:restPos]
				rest = subspec[restPos:]
			case restPos == 0:
				rest = subspec
			default:
				rest = ""
				if len(substr) <= 0 {
					deleted = true
				}
			}
			if done {
				srm.cursr = nil
			}
		}
		i++
	}
	return substr, deleted
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
