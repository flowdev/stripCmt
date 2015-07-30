package main

import (
  // "strings"
)

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

// Renate (Owe): 04635/1386
