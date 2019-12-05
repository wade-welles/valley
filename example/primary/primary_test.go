package primary_test

import (
	"testing"

	"github.com/seeruk/valley/example/primary"
	"github.com/seeruk/valley/valley"
)

func BenchmarkExample_ValidateHappy(b *testing.B) {
	var example primary.Example
	var violations []valley.ConstraintViolation

	example.Text = "Hello"
	example.Texts = []string{"Hello", "World!"}
	example.TextMap = map[string]string{"hello": "world"}
	example.Int = 999
	example.Ints = []int{1}
	example.Nested = primary.NestedExample{Text: "Hello, World!"}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		violations = example.Validate()
	}

	if len(violations) > 0 {
		b.Error("expected no violations")
	}
}

func BenchmarkExample_ValidateUnhappy(b *testing.B) {
	var example primary.Example
	var violations []valley.ConstraintViolation

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		violations = example.Validate()
	}

	if len(violations) == 0 {
		b.Error("expected no violations")
	}
}
