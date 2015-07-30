package main

import (
	"bufio"
	"io"
	"strings"
)

type LineReader interface {
	ReadLine() (line string, err error)
}

type BufferedLineReader struct {
	br *bufio.Reader
}

func NewBufferedLineReader(r io.Reader) *BufferedLineReader {
	return &BufferedLineReader{bufio.NewReader(r)}
}
func (lr *BufferedLineReader) ReadLine() (line string, err error) {
	return lr.br.ReadString('\n')
}

type EofDelayer struct {
	lr    LineReader
	isEof bool
}

func NewEofDelayer(lr LineReader) *EofDelayer {
	return &EofDelayer{lr, false}
}
func (eofd *EofDelayer) ReadLine() (line string, err error) {
	if eofd.isEof {
		return "", io.EOF
	} else {
		line, err = eofd.lr.ReadLine()
		if err == io.EOF && len(line) > 0 {
			eofd.isEof = true
			return line, nil
		}
		return line, err
	}
}

type EolStripper struct {
	lr LineReader
}

func NewEolStripper(lr LineReader) *EolStripper {
	return &EolStripper{lr}
}
func (eols *EolStripper) ReadLine() (line string, err error) {
	line, err = eols.lr.ReadLine()
	if err == nil || err == io.EOF {
		line = strings.TrimRight(line, "\r\n")
	}
	return line, err
}

func NewSaneLineReader(r io.Reader) LineReader {
	return NewEolStripper(NewEofDelayer(NewBufferedLineReader(r)))
}
