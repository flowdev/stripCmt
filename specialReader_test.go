package main

import (
  "testing"
)

func TestLineCommentSpecialReaderStart(t *testing.T) {
  expectStart(t, NewLineCommentSpecialReader(), "//")
}
func TestLineCommentSpecialReaderComment(t *testing.T) {
  expectReadSpecial(t, NewLineCommentSpecialReader(), "// bla comment\n", true, "", 0, true)
}

func TestBlockCommentSpecialReaderStart(t *testing.T) {
  expectStart(t, NewBlockCommentSpecialReader(), "/*")
}
func TestBlockCommentSpecialReaderCommentDone(t *testing.T) {
  expectReadSpecial(t, NewBlockCommentSpecialReader(), "/* bla comment */ blue\n", true, " blue\n", 0, true)
}
func TestBlockCommentSpecialReaderCommentNotDone(t *testing.T) {
  expectReadSpecial(t, NewBlockCommentSpecialReader(), "/* bla comment \n", true, "", 0, false)
}
func TestBlockCommentSpecialReaderCommentContinued(t *testing.T) {
  expectReadSpecial(t, NewBlockCommentSpecialReader(), "*/ bla blue \n", false, " bla blue \n", 0, true)
}

func TestSingleQuoteSpecialReaderStart(t *testing.T) {
  expectStart(t, NewSingleQuoteSpecialReader(), "'")
}
func TestSingleQuoteSpecialReaderDone(t *testing.T) {
  line := "'foo\\r \\n \\\\ \\t \\'' bar"
  expectReadSpecial(t, NewSingleQuoteSpecialReader(), line, true, line, 19, true)
}
func TestSingleQuoteSpecialReaderNotDone(t *testing.T) {
  line := "'foo\\r \\n \\\\ \\t \\\" \\'"
  expectReadSpecial(t, NewSingleQuoteSpecialReader(), line, true, line, len(line), false)
}
func TestSingleQuoteSpecialReaderContinued(t *testing.T) {
  line := "'bar"
  expectReadSpecial(t, NewSingleQuoteSpecialReader(), line, false, line, 1, true)
}

func TestDoubleQuoteSpecialReaderStart(t *testing.T) {
  expectStart(t, NewDoubleQuoteSpecialReader(), "\"")
}
func TestDoubleQuoteSpecialReaderDone(t *testing.T) {
  line := "\"foo\\r \\n \\\\ \\t \\\"\" bar"
  expectReadSpecial(t, NewDoubleQuoteSpecialReader(), line, true, line, 19, true)
}
func TestDoubleQuoteSpecialReaderNotDone(t *testing.T) {
  line := "\"foo\\r \\n \\\\ \\t \\\" \\'"
  expectReadSpecial(t, NewDoubleQuoteSpecialReader(), line, true, line, len(line), false)
}
func TestDoubleQuoteSpecialReaderContinued(t *testing.T) {
  line := "\"bar"
  expectReadSpecial(t, NewDoubleQuoteSpecialReader(), line, false, line, 1, true)
}

func expectStart(t *testing.T, sr SpecialReader, expected string) {
  start := sr.ConstStart()
  if start != expected {
    t.Error("ERROR: Unexpected start: '", start, "' (expected is '", expected, "').")
  }
}
func expectReadSpecial(t *testing.T,
                       sr SpecialReader, line string, firstLine bool,
                       expectedSubstring string, expectedRestPos int, expectedDone bool) {
  substr, restPos, done := sr.ReadSpecial(line, firstLine)
  if substr != expectedSubstring {
    t.Error("ERROR: Unexpected substr: '", substr, "' (expected is '", expectedSubstring, "').")
  }
  if restPos != expectedRestPos {
    t.Error("ERROR: Unexpected restPos:", restPos, "(expected is", expectedRestPos, ").")
  }
  if done != expectedDone {
    t.Error("ERROR: Unexpected done:", done, "(expected is:", expectedDone, ").")
  }
}

