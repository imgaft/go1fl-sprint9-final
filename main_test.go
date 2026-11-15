package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-5, 0, 1, 10, 100} {
		data := generateRandomElements(size)
		wantLength := size
		if size < 0 {
			wantLength = 0
		}
		if len(data) != wantLength {
			t.Errorf("size=%d: got length %d, want %d", size, len(data), wantLength)
		}
		for _, value := range data {
			if value <= 0 || value > 1_000_000 {
				t.Errorf("generated value out of range: %d", value)
			}
		}
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"single", []int{7}, 7},
		{"increasing", []int{1, 2, 3}, 3},
		{"decreasing", []int{9, 4, 1}, 9},
		{"duplicates", []int{5, 5, 5}, 5},
		{"negative", []int{-8, -2, -5}, -2},
		{"mixed", []int{-5, 0, 4, 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"single", []int{9}, 9},
		{"short", []int{3, 8, 2}, 8},
		{"eight", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"remainder", []int{1, 2, 3, 4, 5, 6, 7, 8, 100, 9}, 100},
		{"negative", []int{-8, -3, -9}, -3},
		{"maximumFirst", []int{100, 1, 2, 3, 4, 5, 6, 7, 8}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaximumEqualsMaxChunks(t *testing.T) {
	sizes := []int{
		0,
		1,
		7,
		8,
		9,
		100,
		1000,
		10000,
	}

	for _, size := range sizes {
		data := generateRandomElements(size)

		expected := maximum(data)
		actual := maxChunks(data)

		if actual != expected {
			t.Errorf(
				"size=%d: maximum=%d, maxChunks=%d",
				size,
				expected,
				actual,
			)
		}
	}
}

// Benchmark последовательного поиска максимума.
func BenchmarkMaximum(b *testing.B) {
	data := generateRandomElements(1_000_000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		maximum(data)
	}
}

// Benchmark параллельного поиска максимума.
func BenchmarkMaxChunks(b *testing.B) {
	data := generateRandomElements(1_000_000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		maxChunks(data)
	}
}
