package args

import (
	"testing"
)

func TestNextArg(t *testing.T) {
	input := []string{"FILE_PATH", "minigit", "add", "HelloWorld.MD"}
	tests := []struct {
		expectedLiteral string
	}{
		{"minigit"},
		{"add"},
		{"HelloWorld.MD"},
	}

	args := New(input)

	for i, tt := range tests {
		arg, err := args.Next()
		if err != nil {
			t.Fatalf("tests[%d] - error recieved. err=%q",
				i, err.Error())
		}
		if arg != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, arg)
		}
	}
}

func TestEmpty(t *testing.T) {
	input := []string{}

	args := New(input)

	_, err := args.Next()
	if err.Error() != "err No next argument"{
		t.Fatalf("Expected 'No next argument' err. got: err=%q", err.Error())
	}
}

func TestOneArg(t *testing.T) {
	input := []string{"FILE_PATH"}

	args := New(input)

	_, err := args.Next()
	if err.Error() != "err No next argument"{
		t.Fatalf("Expected 'No next argument' err. got: err=%q", err.Error())
	}
}
