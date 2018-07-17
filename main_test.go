package main

import (
	"regexp"
	"strconv"
	"testing"
)

func Test_shapes(t *testing.T) {
	type expShape struct {
		rand  int
		shape string
	}
	tests := []expShape{
		{0, "chevrons"},
		{1, "concentric-circles"},
		{2, "diamonds"},
		{3, "hexagons"},
		{4, "mosaic-squares"},
		{5, "nested-squares"},
		{6, "octagons"},
		{7, "overlapping-circles"},
		{8, "overlapping-rings"},
		{9, "plaid"},
		{10, "plus-signs"},
		{11, "sine-waves"},
		{12, "squares"},
		{13, "tessellation"},
		{14, "triangles"},
		{15, "xes"},
	}

	for _, tt := range tests {
		got := shapes(tt.rand)
		if got != tt.shape {
			t.Errorf("shape() = %s, want %s", got, tt.shape)
		}
	}
}
func Test_phraseGenerator(t *testing.T) {
	randoms := []int{5, 10, 15, 30}
	for _, rand := range randoms {
		got := phraseGenerator(rand)
		random := strconv.Itoa(rand)
		reg := regexp.MustCompile("(A-Z0-9a-z!?$&){" + random + "}")
		if reg.MatchString(got) {
			t.Errorf("phrase() = %s does not match regexp %s", got, reg)
		}
	}
}

func Test_hexGenerator(t *testing.T) {
	//hexCodes := []string{"#ff5733", "#76eec6", "#8a2be2", "#ff7f24", "#b8860b"}
	for i := 0; i < 10; i++ {
		got := hexGenerator()
		reg := regexp.MustCompile("#(0-9a-e){6}")
		if reg.MatchString(got) {
			t.Errorf("hexCode() = %s does not match regexp %s", got, reg)
		}
	}
}
