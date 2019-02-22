package no_math

import "testing"

var neededPaperTests = []struct {
	input          string
	expectedOutput int
}{
	{"2x3x4", 58},
	{"1x1x10", 43},
}

func TestNeededPaper(t *testing.T) {
	for _, tc := range neededPaperTests {
		t.Run(tc.input, func(t *testing.T) {
			output := NeededPaper(tc.input)
			if output != tc.expectedOutput {
				t.Errorf("Wrong needed paper: %v, expected: %v\n", output, tc.expectedOutput)
			}
		})
	}
}

var neededRibbonTests = []struct {
	input          string
	expectedOutput int
}{
	{"2x3x4", 34},
	{"1x1x10", 14},
}

func TestNeededRibbon(t *testing.T) {
	for _, tc := range neededRibbonTests {
		t.Run(tc.input, func(t *testing.T) {
			output := NeededRibbon(tc.input)
			if output != tc.expectedOutput {
				t.Errorf("Wrong needed ribbon: %v, expected: %v\n", output, tc.expectedOutput)
			}
		})
	}
}

var neededResourceTests = []struct {
	input          string
	resource       string
	expectedOutput int
}{
	{"input_paper", "paper", 1606483},
	{"input_ribbon", "ribbon", 3842356},
}

func TestNeededResourceFromFile(t *testing.T) {
	for _, tc := range neededResourceTests {
		t.Run(tc.input, func(t *testing.T) {
			output := NeededResourceFromFile(tc.input, tc.resource)
			if output != tc.expectedOutput {
				t.Errorf(
					"Wrong needed %v: %v, from file: %v, expected: %v\n",
					tc.resource,
					output,
					tc.input,
					tc.expectedOutput,
				)
			}
		})
	}
}
