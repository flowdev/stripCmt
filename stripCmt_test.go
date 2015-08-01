package main

import (
	"strings"
	"testing"
)

func TestAllReaderNoSpecial(t *testing.T) {
	lr := NewFormatter(NewStripCmtLineReader(NewSaneLineReader(strings.NewReader(fooN + bar))))
	expectLines(t, lr, foo, bar)
}
func TestAllReaderFormating(t *testing.T) {
	lr := NewFormatter(NewStripCmtLineReader(NewSaneLineReader(strings.NewReader("a  \t \n  \t \r\n \t\n \nb"))))
	expectLines(t, lr, "a", "", "b")
}
func TestAllReaderLineComment(t *testing.T) {
	lr := NewFormatter(NewStripCmtLineReader(NewSaneLineReader(strings.NewReader("a'// foo' \t \r // bar "))))
	expectLines(t, lr, "a'// foo'")
}
func TestAllReaderBlockComment(t *testing.T) {
	lr := NewFormatter(NewStripCmtLineReader(NewSaneLineReader(strings.NewReader("a'/* foo */' \t \r /* bar */ \t \r \n"))))
	expectLines(t, lr, "a'/* foo */'")
}
func TestAllReaderMixed(t *testing.T) {
	lr := NewFormatter(NewStripCmtLineReader(NewSaneLineReader(strings.NewReader(
		"a'/* // */' \t \r // bar */ \t \r \n" +
			"b\" // /* */\" /* bla\n" +
			" blabbi blaaaa!\r\n" +
			"*/c ' // boo' // bar\n" +
			"d ' \" ' "))))
	expectLines(t, lr, "a'/* // */'", "b\" // /* */\"", "c ' // boo'", "d ' \" '")
}

func TestSpecialReaderManager(t *testing.T) {
	// TODO: !!!
}
