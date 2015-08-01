package main

import (
	"io"
	"strings"
	"testing"
)

const foo = "foo"
const fooN = foo + "\n"
const bar = "bar"

func TestLineReader0(t *testing.T) {
	lr := NewBufferedLineReader(strings.NewReader(""))
	expectEof(t, lr)
}
func TestLineReader1(t *testing.T) {
	lr := NewBufferedLineReader(strings.NewReader(fooN))
	expectLines(t, lr, fooN)
}
func TestLineReader2(t *testing.T) {
	lr := NewBufferedLineReader(strings.NewReader(fooN + bar))
	expectLinesEof(t, lr, fooN, bar)
}

func TestEofDelayerNoDelay(t *testing.T) {
	lcs := NewEofDelayer(NewBufferedLineReader(strings.NewReader(fooN)))
	expectLines(t, lcs, fooN)
}
func TestEofDelayerDelay(t *testing.T) {
	lcs := NewEofDelayer(NewBufferedLineReader(strings.NewReader(foo)))
	expectLines(t, lcs, foo)
}

func TestEolStripper0(t *testing.T) {
	eols := NewEolStripper(NewBufferedLineReader(strings.NewReader(foo)))
	expectLinesEof(t, eols, foo)
}
func TestEolStripperN(t *testing.T) {
	eols := NewEolStripper(NewBufferedLineReader(strings.NewReader(fooN)))
	expectLines(t, eols, foo)
}
func TestEolStripperR(t *testing.T) {
	eols := NewEolStripper(NewBufferedLineReader(strings.NewReader(foo + "\r")))
	expectLinesEof(t, eols, foo)
}
func TestEolStripperRN(t *testing.T) {
	eols := NewEolStripper(NewBufferedLineReader(strings.NewReader(foo + "\r\n")))
	expectLines(t, eols, foo)
}

func expectLines(t *testing.T, lr LineReader, expected ...string) {
	for i := 0; i < len(expected); i++ {
		expectLine(t, lr, i+1, expected[i])
	}
	expectEof(t, lr)
}
func expectLinesEof(t *testing.T, lr LineReader, expected ...string) {
	len := len(expected)
	for i := 0; i < len-1; i++ {
		expectLine(t, lr, i+1, expected[i])
	}
	expectEofLine(t, lr, len, expected[len-1])
}
func expectLine(t *testing.T, lr LineReader, lineNum int, expected string) {
	line, err := lr.ReadLine()
	if err != nil {
		t.Fatal("ERROR: Unexpected error for line (", lineNum, "): ", err.Error())
	}
	if line != expected {
		t.Error("ERROR: Unexpected line (", lineNum, "): '", line, "' (expected is '", expected, "').")
	}
}
func expectEofLine(t *testing.T, lr LineReader, lineNum int, expected string) {
	line, err := lr.ReadLine()
	if err != io.EOF {
		t.Error("ERROR: Expected EOF but got: ", err)
	}
	if line != expected {
		t.Error("ERROR: Unexpected line (", lineNum, "): '", line, "' (expected is '", expected, "').")
	}
}
func expectEof(t *testing.T, lr LineReader) {
	_, err := lr.ReadLine()
	if err != io.EOF {
		t.Fatal("ERROR: Expected EOF but got: ", err)
	}
}
