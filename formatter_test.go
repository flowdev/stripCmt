package main

import (
	"strings"
	"testing"
)

func TestSpaceTrimmerNoSpace(t *testing.T) {
	st := NewSpaceTrimmer(NewSaneLineReader(strings.NewReader(foo)))
	expectLines(t, st, foo)
}
func TestSpaceTrimmerAllSpace(t *testing.T) {
	st := NewSpaceTrimmer(NewSaneLineReader(strings.NewReader(" \t  \r \t")))
	expectLines(t, st, "")
}

func TestEmptyLineStripperNoEmptyLines(t *testing.T) {
	els := NewEmptyLineStripper(NewSaneLineReader(strings.NewReader("\n" + bar)))
	expectLines(t, els, "", bar)
}
func TestEmptyLineStripperAllEmptyLines(t *testing.T) {
	els := NewEmptyLineStripper(NewSaneLineReader(strings.NewReader("\n\r\n\n\r\n\n")))
	expectLines(t, els, "")
}
