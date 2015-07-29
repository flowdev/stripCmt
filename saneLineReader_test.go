package main

import (
	"testing"
	"io"
	"strings"
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
  expectOneLine(t, lr, fooN)
}
func TestLineReader2(t *testing.T) {
  lr := NewBufferedLineReader(strings.NewReader(fooN + bar))
  expectTwoLinesEof(t, lr, fooN, bar)
}

func TestEofDelayerNoDelay(t *testing.T) {
  lcs := NewEofDelayer(NewBufferedLineReader(strings.NewReader(fooN)))
  expectOneLine(t, lcs, fooN)
}
func TestEofDelayerDelay(t *testing.T) {
  lcs := NewEofDelayer(NewBufferedLineReader(strings.NewReader(foo)))
  expectOneLine(t, lcs, foo)
}

func TestEolStripper0(t *testing.T) {
  eols := NewEolStripper(NewBufferedLineReader(strings.NewReader(foo)))
  expectEofLine(t, eols, 1, foo)
}
func TestEolStripperN(t *testing.T) {
  eols := NewEolStripper(NewBufferedLineReader(strings.NewReader(fooN)))
  expectOneLine(t, eols, foo)
}
func TestEolStripperR(t *testing.T) {
  eols := NewEolStripper(NewBufferedLineReader(strings.NewReader(foo + "\r")))
  expectEofLine(t, eols, 1, foo)
}
func TestEolStripperRN(t *testing.T) {
  eols := NewEolStripper(NewBufferedLineReader(strings.NewReader(foo + "\r\n")))
  expectOneLine(t, eols, foo)
}


func expectOneLine(t *testing.T, lr LineReader, expected string) {
  expectLine(t, lr, 1, expected)
  expectEof(t, lr)
}
func expectTwoLines(t *testing.T, lr LineReader, expected1 string, expected2 string) {
  expectLine(t, lr, 1, expected1)
  expectLine(t, lr, 2, expected2)
  expectEof(t, lr)
}
func expectTwoLinesEof(t *testing.T, lr LineReader, expected1 string, expected2 string) {
  expectLine(t, lr, 1, expected1)
  expectEofLine(t, lr, 2, expected2)
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

