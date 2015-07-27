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
func NewBufferedLineReader(r io.Reader) *BufferedLineReader {
  return &BufferedLineReader{bufio.NewReader(r)}
}
func (lr *BufferedLineReader) ReadLine() (line string, err error) {
  return lr.br.ReadString('\n')
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
  lr LineReader
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

type LineCommentStripper struct {
  lr LineReader
}
func NewLineCommentStripper(lr LineReader) *LineCommentStripper {
  return &LineCommentStripper{lr}
}
func (lcs *LineCommentStripper) ReadLine() (line string, err error) {
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
func NewBlockCommentStripper(lr *LineReader) *BlockCommentStripper {
  return &BlockCommentStripper{lr, false}
}
*/

