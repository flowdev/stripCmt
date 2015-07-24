package main

import (
  "io"
  "bufio"
  "strings"
)

type LineReader interface {
  ReadLine() (line string, err error)
}

type BufferedLineReader struct {
  br *bufio.Reader
}
func NewBufferedLineReader(r io.Reader) LineReader {
  return BufferedLineReader{bufio.NewReader(r)}
}
func (lr BufferedLineReader) ReadLine() (line string, err error) {
  return lr.br.ReadString('\n')
}

type EolStripper struct {
  lr LineReader
}
func NewEolStripper(lr LineReader) EolStripper {
  return EolStripper{lr}
}
func (eols EolStripper) ReadLine() (line string, err error) {
  line, err = eols.lr.ReadLine()
  if err == nil || err == io.EOF {
    line = strings.TrimRight(line, "\r\n")
  }
  return line, err
}

type EofDelayer struct {
  lr LineReader
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

type LineCommentStripper struct {
  lr LineReader
}
func NewLineCommentStripper(lr LineReader) LineCommentStripper {
  return LineCommentStripper{lr}
}
func (lcs LineCommentStripper) ReadLine() (line string, err error) {
  line, err = lcs.lr.ReadLine()
  if err == nil || err == io.EOF {
    idx := strings.Index(line, "//")
    if idx >= 0 {
      line = line[0:idx]
    }
  }
  return line, err
}


/*
type BlockCommentStripper struct {
  lr LineReader
  inComment bool
}
func NewBlockCommentStripper(lr LineReader) CommentStripper {
  return BlockCommentStripper{lr, false}
}

type EndOfLineTrimmer struct {
  cs CommentStripper
}

type EmptyLineStripper struct {
  eolt EndOfLineTrimmer
  emptyLineReturned bool
}
*/

type Cell struct {
	state bool
}

func (c *Cell) isAlive() bool {
	return c.state 
}

func NewCell() *Cell {
	return &Cell{true}
}

func (c *Cell) evolve() {
	c.state = false
}

