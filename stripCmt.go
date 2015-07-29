package main

import (
  "strings"
)

type SpecialReader interface {
  ConstStart() string
  ReadSpecial(line string, firstLine bool) (substring string, restPos int, done bool)
}
type SpecialReaderManager struct {
  lr LineReader
  srs []SpecialReader
  currentSpecialReader SpecialReader
}
func NewSpecialReaderManager(lr LineReader, srs ...SpecialReader) *SpecialReaderManager {
  return &SpecialReaderManager{lr, srs, nil}
}
func (srm *SpecialReaderManager) ReadLine() (line string, err error) {
  line, err = srm.lr.ReadLine()
  return line, err
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
  } else {
    return "", 0, false
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
  orgLine := line
  orgPos := 0
  substr := line
  if firstLine {
    substr = line[1:]
    orgPos = 1
  }
  pos := strings.IndexAny(substr, "'\\")
  for pos >= 0 && substr[pos] == '\\' {
    orgPos += pos+2
    substr = substr[pos+2:]
    pos = strings.IndexAny(substr, "'\\")
  }
  if pos >= 0 {
    return orgLine, orgPos+pos+1, true
  } else {
    return orgLine, len(orgLine), false
  }
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
  orgLine := line
  orgPos := 0
  substr := line
  if firstLine {
    substr = line[1:]
    orgPos = 1
  }
  pos := strings.IndexAny(substr, "\"\\")
  for pos >= 0 && substr[pos] == '\\' {
    orgPos += pos+2
    substr = substr[pos+2:]
    pos = strings.IndexAny(substr, "\"\\")
  }
  if pos >= 0 {
    return orgLine, orgPos+pos+1, true
  } else {
    return orgLine, len(orgLine), false
  }
}

