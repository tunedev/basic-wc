package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	exp := 4
	res := count(b, false, false)
	throwErrIfNotEqual(t, exp, res)
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("line 1 word1 word2\nline2 word1 word2\nline3 word1 word2")

	exp := 3
	res := count(b, true, false)
	throwErrIfNotEqual(t, exp, res)
}

func TestCountBytes(t *testing.T) {
	
	b := bytes.NewBufferString("so many things are happening in this life")
	exp := b.Len()

	res := count(b, false, true)
	throwErrIfNotEqual(t, exp, res)
}

func throwErrIfNotEqual(t *testing.T, exp, actual interface{}) {
	if actual != exp {
		t.Errorf("Expected %d, got %d intead.\n", exp, actual)
	}
}