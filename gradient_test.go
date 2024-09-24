package IceGradient

import (
	"strings"
	"testing"
)

// Copilot自动生成的测试用例

// TestBlendColors tests the BlendColors function
func TestBlendColors(t *testing.T) {
	colors := []interface{}{"#FF0000", "#00FF00", "#0000FF"}
	expected := []int{85, 85, 85}
	result, err := BlendColors(colors...)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("expected %d, got %d", expected[i], v)
		}
	}
}

// TestGradientTextWithTemplate tests the GradientTextWithTemplate function
func TestGradientTextWithTemplate(t *testing.T) {
	text := "Hello, World!"
	templateName := "Rainbow"
	result := GradientTextWithTemplate(text, templateName)
	if !strings.Contains(result, "\033[38;2;") {
		t.Errorf("expected ANSI color codes in result, got %s", result)
	}
}

// TestGradientBackgroundTextWithTemplate tests the GradientBackgroundTextWithTemplate function
func TestGradientBackgroundTextWithTemplate(t *testing.T) {
	text := "Hello, World!"
	templateName := "Rainbow"
	result := GradientBackgroundTextWithTemplate(text, templateName)
	if !strings.Contains(result, "\033[48;2;") {
		t.Errorf("expected ANSI background color codes in result, got %s", result)
	}
}

// BenchmarkBlendColors benchmarks the BlendColors function
func BenchmarkBlendColors(b *testing.B) {
	colors := []interface{}{"#FF0000", "#00FF00", "#0000FF"}
	for i := 0; i < b.N; i++ {
		_, _ = BlendColors(colors...)
	}
}

// BenchmarkGradientTextWithTemplate benchmarks the GradientTextWithTemplate function
func BenchmarkGradientTextWithTemplate(b *testing.B) {
	text := "Hello, World!"
	templateName := "Rainbow"
	for i := 0; i < b.N; i++ {
		_ = GradientTextWithTemplate(text, templateName)
	}
}

// BenchmarkGradientBackgroundTextWithTemplate benchmarks the GradientBackgroundTextWithTemplate function
func BenchmarkGradientBackgroundTextWithTemplate(b *testing.B) {
	text := "Hello, World!"
	templateName := "Rainbow"
	for i := 0; i < b.N; i++ {
		_ = GradientBackgroundTextWithTemplate(text, templateName)
	}
}
