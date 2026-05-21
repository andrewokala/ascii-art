package main

import (
	"ascii-art/engine"
	"testing"
)

// ---------------------------------------------------------------------------
// Tests for Split
// ---------------------------------------------------------------------------

func TestSplit_EmptyString(t *testing.T) {
	result := engine.SplitInput("")
	if len(result) != 1 || result[0] != "" {
		t.Errorf("SplitInput(\"\") = %v; want [\"\"]", result)
	}
}

func TestSplit_NoNewline(t *testing.T) {
	result := engine.SplitInput("Hello")
	if len(result) != 1 || result[0] != "Hello" {
		t.Errorf("SplitInput(\"Hello\") = %v; want [\"Hello\"]", result)
	}
}

func TestSplit_SingleNewline(t *testing.T) {
	result := engine.SplitInput(`\n`)
	if len(result) != 2 || result[0] != "" || result[1] != "" {
		t.Errorf(`SplitInput("\\n") = %v; want [""]`, result)
	}
}

func TestSplit_TrailingNewline(t *testing.T) {
	result := engine.SplitInput(`Hello\n`)
	if len(result) != 2 || result[0] != "Hello" || result[1] != "" {
		t.Errorf(`SplitInput("Hello\\n") = %v; want ["Hello", ""]`, result)
	}
}

func TestSplit_DoubleNewline(t *testing.T) {
	result := engine.SplitInput(`Hello\n\nThere`)
	if len(result) != 3 || result[0] != "Hello" || result[1] != "" || result[2] != "There" {
		t.Errorf(`SplitInput("Hello\\n\\nThere") = %v; want ["Hello", "", "There"]`, result)
	}
}

func TestSplit_MultipleSegments(t *testing.T) {
	result := engine.SplitInput(`A\nB\nC`)
	expected := []string{"A", "B", "C"}
	if len(result) != len(expected) {
		t.Fatalf("Split length = %d; want %d", len(result), len(expected))
	}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Split[%d] = %q; want %q", i, result[i], v)
		}
	}
}

// ---------------------------------------------------------------------------
// Tests for Validator
// ---------------------------------------------------------------------------

func TestValidator_TwoArgs_Standard(t *testing.T) {
	input, banner, err := engine.Validator([]string{"prog", "Hello"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if input != "Hello" {
		t.Errorf("input = %q; want \"Hello\"", input)
	}
	if banner != "standard" {
		t.Errorf("banner = %q; want \"standard\"", banner)
	}
}

func TestValidator_ThreeArgs_Shadow(t *testing.T) {
	input, banner, err := engine.Validator([]string{"prog", "Hi", "shadow"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if input != "Hi" {
		t.Errorf("input = %q; want \"Hi\"", input)
	}
	if banner != "shadow" {
		t.Errorf("banner = %q; want \"shadow\"", banner)
	}
}

func TestValidator_ThreeArgs_Thinkertoy(t *testing.T) {
	_, banner, err := engine.Validator([]string{"prog", "Hi", "thinkertoy"})
	if err != nil || banner != "thinkertoy" {
		t.Errorf("expected thinkertoy, got %q, err: %v", banner, err)
	}
}

func TestValidator_TooFewArgs(t *testing.T) {
	_, _, err := engine.Validator([]string{"prog"})
	if err == nil {
		t.Error("expected error for too few args, got nil")
	}
}

func TestValidator_TooManyArgs(t *testing.T) {
	_, _, err := engine.Validator([]string{"prog", "a", "standard", "extra"})
	if err == nil {
		t.Error("expected error for too many args, got nil")
	}
}

func TestValidator_InvalidBanner(t *testing.T) {
	_, _, err := engine.Validator([]string{"prog", "Hello", "matrix"})
	if err == nil {
		t.Error("expected error for invalid banner name, got nil")
	}
}

func TestValidator_InvalidChar(t *testing.T) {
	_, _, err := engine.Validator([]string{"prog", "Hello\tWorld"})
	if err == nil {
		t.Error("expected error for tab character, got nil")
	}
}

func TestValidator_ValidNewlineEscape(t *testing.T) {
	_, _, err := engine.Validator([]string{"prog", `Hello\nWorld`})
	if err != nil {
		t.Errorf("unexpected error for valid \\n escape: %v", err)
	}
}

func TestValidator_EmptyInput(t *testing.T) {
	input, banner, err := engine.Validator([]string{"prog", ""})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if input != "" || banner != "standard" {
		t.Errorf("got input=%q banner=%q", input, banner)
	}
}

// ---------------------------------------------------------------------------
// Tests for LoadBanner
// ---------------------------------------------------------------------------

func TestLoadBanner_Standard(t *testing.T) {
	bmap, err := engine.LoadBanner("standard")
	if err != nil {
		t.Fatalf("LoadBanner(standard) error: %v", err)
	}
	if len(bmap) != 95 {
		t.Errorf("len(bannerMap) = %d; want 95", len(bmap))
	}
	if _, ok := bmap[' ']; !ok {
		t.Error("space character missing from bannerMap")
	}
	if _, ok := bmap['A']; !ok {
		t.Error("'A' character missing from bannerMap")
	}
}

func TestLoadBanner_Shadow(t *testing.T) {
	_, err := engine.LoadBanner("shadow")
	if err != nil {
		t.Fatalf("LoadBanner(shadow) error: %v", err)
	}
}

func TestLoadBanner_Thinkertoy(t *testing.T) {
	_, err := engine.LoadBanner("thinkertoy")
	if err != nil {
		t.Fatalf("LoadBanner(thinkertoy) error: %v", err)
	}
}

func TestLoadBanner_NonExistent(t *testing.T) {
	_, err := engine.LoadBanner("doesnotexist")
	if err == nil {
		t.Error("expected error for non-existent banner, got nil")
	}
}

// ---------------------------------------------------------------------------
// Tests for GenerateArt
// ---------------------------------------------------------------------------

func TestGenerateArt_EmptyLine(t *testing.T) {
	bmap, _ := engine.LoadBanner("standard")
	art := engine.GenerateArt([]string{""}, bmap)

	if art == nil {
		return
	}
	for _, row := range art {
		if row != "" {
			t.Errorf("expected only empty rows, got %v", art)
			return
		}
	}
}

func TestGenerateArt_NonEmpty(t *testing.T) {
	bmap, _ := engine.LoadBanner("standard")
	art := engine.GenerateArt([]string{"A"}, bmap)
	if len(art) != 8 {
		t.Fatalf("GenerateArt(\"A\") returned %d rows; want 8", len(art))
	}
}

func TestGenerateArt_Space(t *testing.T) {
	bmap, _ := engine.LoadBanner("standard")
	art := engine.GenerateArt([]string{" "}, bmap)
	if len(art) != 8 {
		t.Fatalf("GenerateArt(\" \") returned %d rows; want 8", len(art))
	}
}

// ---------------------------------------------------------------------------
// Integration
// ---------------------------------------------------------------------------

func TestRender_EmptyInput(t *testing.T) {
	bmap, _ := engine.LoadBanner("standard")
	lines := engine.SplitInput("")
	art := engine.GenerateArt(lines, bmap)
	engine.Render(art)
}

func TestRender_OnlyNewline(t *testing.T) {
	bmap, _ := engine.LoadBanner("standard")
	lines := engine.SplitInput(`\n`)
	art := engine.GenerateArt(lines, bmap)
	engine.Render(art)
}

func TestRender_HelloWorld(t *testing.T) {
	bmap, _ := engine.LoadBanner("standard")
	lines := engine.SplitInput("Hello")
	art := engine.GenerateArt(lines, bmap)
	engine.Render(art)
}
