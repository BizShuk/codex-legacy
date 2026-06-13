package fundation

import "testing"

type StringTest struct {
	name string
	s    string
}

func GenerateTestCase_String() []StringTest {
	return []StringTest{
		{name: "Simple token line", s: "a\nb\nc\nd\n\n\n\n"},
		{name: "Simple token line", s: "a b\nccc\nd eeeee f\ng\n\n\n\n"},
	}
}

func TestStringSplit_1(t *testing.T) {
	tests := GenerateTestCase_String()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StringSplit_1(tt.s)
		})
	}
}

func TestStringSplit_2(t *testing.T) {
	tests := GenerateTestCase_String()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StringSplit_2(tt.s)
		})
	}
}

func TestStringSplit_3(t *testing.T) {
	tests := GenerateTestCase_String()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StringSplit_3(tt.s)
		})
	}
}
