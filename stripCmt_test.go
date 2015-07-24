package main

import (
	"testing"
	"io"
	"strings"
//	"fmt"
)

const foo = "foo"
const fooN = "foo\n"
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
  lr := NewBufferedLineReader(strings.NewReader(strings.Join([]string{foo, bar}, "\n")))
  expectTwoLinesEof(t, lr, fooN, bar)
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
  eols := NewEolStripper(NewBufferedLineReader(strings.NewReader(strings.Join([]string{foo, "\r"}, ""))))
  expectEofLine(t, eols, 1, foo)
}
func TestEolStripperRN(t *testing.T) {
  eols := NewEolStripper(NewBufferedLineReader(strings.NewReader(strings.Join([]string{foo, "\r\n"}, ""))))
  expectOneLine(t, eols, foo)
}

func TestEofDelayerNoDelay(t *testing.T) {
  lcs := NewEofDelayer(NewBufferedLineReader(strings.NewReader(fooN)))
  expectOneLine(t, lcs, fooN)
}
func TestEofDelayerDelay(t *testing.T) {
  lcs := NewEofDelayer(NewBufferedLineReader(strings.NewReader(foo)))
  expectOneLine(t, lcs, foo)
}

func TestLineCommentStripperNoComment(t *testing.T) {
  lcs := NewLineCommentStripper(NewEofDelayer(NewEolStripper(NewBufferedLineReader(strings.NewReader(fooN)))))
  expectOneLine(t, lcs, foo)
}
func TestLineCommentStripperHalfComment(t *testing.T) {
  lcs := NewLineCommentStripper(NewEofDelayer(NewEolStripper(NewBufferedLineReader(strings.NewReader(strings.Join([]string{foo, "// blue\n"}, ""))))))
  expectOneLine(t, lcs, foo)
}
func TestLineCommentStripperFullComment(t *testing.T) {
  lcs := NewLineCommentStripper(NewEofDelayer(NewEolStripper(NewBufferedLineReader(strings.NewReader("// blue")))))
  expectOneLine(t, lcs, "")
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



func TestCellIsAlive(t *testing.T) {
	cell := NewCell()
	if !cell.isAlive() {
		t.Fatal("Test Cell is Alive: RED") 
	}
}

func TestCellIsDead(t *testing.T){
	cell := Cell{false}
	if cell.isAlive() {
		t.Fatal("Fatal: the cell is alive.")
	}
}

func TestCellDies(t *testing.T) {
	cell := Cell{true}
	cell.evolve()
	if cell.isAlive() {
		t.Fatal("Fatal: Cell is supposed to die")
	}
}

