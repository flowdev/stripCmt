package main

import (
	"testing"
	"strings"
)

func TestSpaceTrimmerNoSpace(t *testing.T) {
	st := NewSpaceTrimmer(NewSaneLineReader(strings.NewReader(foo)))
	expectOneLine(t, st, foo)
}
func TestSpaceTrimmerAllSpace(t *testing.T) {
	st := NewSpaceTrimmer(NewSaneLineReader(strings.NewReader(" \t  \r \t")))
	expectOneLine(t, st, "")
}

func TestEmptyLineStripperNoEmptyLines(t *testing.T) {
	els := NewEmptyLineStripper(NewSaneLineReader(strings.NewReader("\n" + bar)))
	expectTwoLines(t, els, "", bar)
}
func TestEmptyLineStripperAllEmptyLines(t *testing.T) {
	els := NewEmptyLineStripper(NewSaneLineReader(strings.NewReader("\n\r\n\n\r\n\n")))
	expectOneLine(t, els, "")
}

