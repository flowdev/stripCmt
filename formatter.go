package main

import (
	"io"
	"strings"
)

type SpaceTrimmer struct {
	lr LineReader
}

func NewSpaceTrimmer(lr LineReader) *SpaceTrimmer {
	return &SpaceTrimmer{lr}
}
func (st *SpaceTrimmer) ReadLine() (line string, err error) {
	line, err = st.lr.ReadLine()
	if err == nil || err == io.EOF {
		line = strings.TrimRight(line, "\r\n\t ")
	}
	return line, err
}

type EmptyLineStripper struct {
	lr            LineReader
	lastLineEmpty bool
}

func NewEmptyLineStripper(lr LineReader) *EmptyLineStripper {
	return &EmptyLineStripper{lr, false}
}
func (els *EmptyLineStripper) ReadLine() (line string, err error) {
	line, err = els.lr.ReadLine()
	for els.lastLineEmpty {
		if err == nil && len(line) <= 0 {
			line, err = els.lr.ReadLine()
		} else {
			els.lastLineEmpty = false
		}
	}
	if len(line) <= 0 {
		els.lastLineEmpty = true
	}
	return line, err
}

func NewFormatter(lr LineReader) LineReader {
	return NewEmptyLineStripper(NewSpaceTrimmer(lr))
}
