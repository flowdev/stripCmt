package main

import (
	"testing"
)

func TestSpecialReaderManager(t *testing.T) {
  // TODO: !!!
}

func TestLineCommentSpecialReaderStart(t *testing.T) {
  lcsr := NewLineCommentSpecialReader()
  start := lcsr.ConstStart()
  if start != "//" {
    t.Error("ERROR: Unexpected start: '", start, "' (expected is ' //  ').")
  }
}
func TestLineCommentSpecialReaderComment(t *testing.T) {
  lcsr := NewLineCommentSpecialReader()
  substr, restPos, done := lcsr.ReadSpecial("// bla comment\n", true)
  if substr != "" {
    t.Error("ERROR: Unexpected substr: '", substr, "' (expected is '  ').")
  }
  if restPos != 0 {
    t.Error("ERROR: Unexpected restPos:", restPos, "(expected is 0).")
  }
  if !done {
    t.Error("ERROR: Unexpected done:", done, "(expected is: true).")
  }
}

func TestBlockCommentSpecialReaderStart(t *testing.T) {
  bcsr := NewBlockCommentSpecialReader()
  start := bcsr.ConstStart()
  if start != "/*" {
    t.Error("ERROR: Unexpected start: '", start, "' (expected is ' /*  ').")
  }
}
func TestBlockCommentSpecialReaderCommentDone(t *testing.T) {
  bcsr := NewBlockCommentSpecialReader()
  substr, restPos, done := bcsr.ReadSpecial("/* bla comment */ blue\n", true)
  if substr != " blue\n" {
    t.Error("ERROR: Unexpected substr: '", substr, "' (expected is '  blue\\n ').")
  }
  if restPos != 0 {
    t.Error("ERROR: Unexpected restPos:", restPos, "(expected is 0).")
  }
  if !done {
    t.Error("ERROR: Unexpected done:", done, "(expected is: true).")
  }
}
func TestBlockCommentSpecialReaderCommentNotDone(t *testing.T) {
  bcsr := NewBlockCommentSpecialReader()
  substr, restPos, done := bcsr.ReadSpecial("/* bla comment \n", true)
  if substr != "" {
    t.Error("ERROR: Unexpected substr: '", substr, "' (expected is '  blue\\n ').")
  }
  if restPos != 0 {
    t.Error("ERROR: Unexpected restPos:", restPos, "(expected is 0).")
  }
  if done {
    t.Error("ERROR: Unexpected done:", done, "(expected is: false).")
  }
}
func TestBlockCommentSpecialReaderCommentContinued(t *testing.T) {
  bcsr := NewBlockCommentSpecialReader()
  substr, restPos, done := bcsr.ReadSpecial("*/ bla blue \n", false)
  if substr != " bla blue \n" {
    t.Error("ERROR: Unexpected substr: '", substr, "' (expected is '  bla blue \\n ').")
  }
  if restPos != 0 {
    t.Error("ERROR: Unexpected restPos:", restPos, "(expected is 0).")
  }
  if !done {
    t.Error("ERROR: Unexpected done:", done, "(expected is: true).")
  }
}

func TestSingleQuoteSpecialReaderStart(t *testing.T) {
  sqsr := NewSingleQuoteSpecialReader()
  start := sqsr.ConstStart()
  if start != "'" {
    t.Error("ERROR: Unexpected start: '", start, "' (expected is ' ' ').")
  }
}
func TestSingleQuoteSpecialReaderDone(t *testing.T) {
  sqsr := NewSingleQuoteSpecialReader()
  line := "'foo\\r \\n \\\\ \\t \\'' bar"
  substr, restPos, done := sqsr.ReadSpecial(line, true)
  if substr != line {
    t.Error("ERROR: Unexpected substr: '", substr, "' (expected is '", line, "').")
  }
  if restPos != 19 {
    t.Error("ERROR: Unexpected restPos:", restPos, "(expected is 19).")
  }
  if !done {
    t.Error("ERROR: Unexpected done:", done, "(expected is: true).")
  }
}
func TestSingleQuoteSpecialReaderNotDone(t *testing.T) {
  sqsr := NewSingleQuoteSpecialReader()
  line := "'foo\\r \\n \\\\ \\t \\\" \\'"
  substr, restPos, done := sqsr.ReadSpecial(line, true)
  if substr != line {
    t.Error("ERROR: Unexpected substr: '", substr, "' (expected is '", line, "').")
  }
  if restPos != len(line) {
    t.Error("ERROR: Unexpected restPos:", restPos, "(expected is", len(line), ").")
  }
  if done {
    t.Error("ERROR: Unexpected done:", done, "(expected is: false).")
  }
}
func TestSingleQuoteSpecialReaderContinued(t *testing.T) {
  sqsr := NewSingleQuoteSpecialReader()
  line := "'bar"
  substr, restPos, done := sqsr.ReadSpecial(line, false)
  if substr != line {
    t.Error("ERROR: Unexpected substr: '", substr, "' (expected is '", line, "').")
  }
  if restPos != 1 {
    t.Error("ERROR: Unexpected restPos:", restPos, "(expected is 1).")
  }
  if !done {
    t.Error("ERROR: Unexpected done:", done, "(expected is: true).")
  }
}

func TestDoubleQuoteSpecialReaderStart(t *testing.T) {
  sqsr := NewDoubleQuoteSpecialReader()
  start := sqsr.ConstStart()
  if start != "\"" {
    t.Error("ERROR: Unexpected start: '", start, "' (expected is ' \" ').")
  }
}
func TestDoubleQuoteSpecialReaderDone(t *testing.T) {
  sqsr := NewDoubleQuoteSpecialReader()
  line := "\"foo\\r \\n \\\\ \\t \\\"\" bar"
  substr, restPos, done := sqsr.ReadSpecial(line, true)
  if substr != line {
    t.Error("ERROR: Unexpected substr: '", substr, "' (expected is '", line, "').")
  }
  if restPos != 19 {
    t.Error("ERROR: Unexpected restPos:", restPos, "(expected is 19).")
  }
  if !done {
    t.Error("ERROR: Unexpected done:", done, "(expected is: true).")
  }
}
func TestDoubleQuoteSpecialReaderNotDone(t *testing.T) {
  sqsr := NewDoubleQuoteSpecialReader()
  line := "\"foo\\r \\n \\\\ \\t \\\" \\'"
  substr, restPos, done := sqsr.ReadSpecial(line, true)
  if substr != line {
    t.Error("ERROR: Unexpected substr: '", substr, "' (expected is '", line, "').")
  }
  if restPos != len(line) {
    t.Error("ERROR: Unexpected restPos:", restPos, "(expected is", len(line), ").")
  }
  if done {
    t.Error("ERROR: Unexpected done:", done, "(expected is: false).")
  }
}
func TestDoubleQuoteSpecialReaderContinued(t *testing.T) {
  sqsr := NewDoubleQuoteSpecialReader()
  line := "\"bar"
  substr, restPos, done := sqsr.ReadSpecial(line, false)
  if substr != line {
    t.Error("ERROR: Unexpected substr: '", substr, "' (expected is '", line, "').")
  }
  if restPos != 1 {
    t.Error("ERROR: Unexpected restPos:", restPos, "(expected is 1).")
  }
  if !done {
    t.Error("ERROR: Unexpected done:", done, "(expected is: true).")
  }
}

