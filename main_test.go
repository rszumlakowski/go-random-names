package main

import (
	"testing"
)

func TestPickWordEmpty(t *testing.T) {
	word := pickWord("")

	if word != "" {
		t.Error("pickWord returned non-empty string.")
	}
}

func TestPickWordAllWhitespace(t *testing.T) {
	word := pickWord("  \n\n   ")

	if word != "" {
		t.Error("pickWord returned non-empty string.")
	}
}

func TestPickWordOne(t *testing.T) {
	word := pickWord("one")

	if word != "one" {
		t.Errorf("pickWord returned '%s'. Expected 'one'.", word)
	}
}

func TestPickWordTwoOrThree(t *testing.T) {
	word := pickWord("two three")

	if word != "two" && word != "three" {
		t.Errorf("pickWord returned '%s'. Expected 'two' or 'three'.", word)
	}
}

func TestComponents(t *testing.T) {
	components := generateComponents("blue", "sky")

	if components.adjective != "blue" {
		t.Errorf("generateComponents returned adjective '%s'. Expected 'blue'.", components.adjective)
	}

	if components.noun != "sky" {
		t.Errorf("generateComponents returned noun '%s'. Expected 'sky'.", components.adjective)
	}
}

func TestFormatComponents(t *testing.T) {
	c := Components{"green", "frog", 65535}
	result := c.format()

	if result != "green-frog-0000ffff" {
		t.Errorf("Components.format returned '%s'. Expected 'green-frog-0000ffff'.", result)
	}
}

func TestGenerate(t *testing.T) {
	name := generate()

	if name == "" {
		t.Error("Generate returned empty string.")
	}
}
