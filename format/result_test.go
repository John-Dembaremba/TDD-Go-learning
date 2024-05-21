package format

import (
	"fmt"
	"strings"
	"testing"
)

func TestResult(t *testing.T) {
	// Arrange
	result := 5.0
	expression := "2+3"

	// Act
	wrappedResult := Result(expression, result)

	if !strings.ContainsAny(wrappedResult, expression) {
		t.Errorf("error does not contain: got %q, want %q", wrappedResult, expression)
	}

	if !strings.Contains(wrappedResult, fmt.Sprint(result)) {
		t.Errorf("error does not contain: got %s, want %s", wrappedResult, fmt.Sprint(result))
	}

}
