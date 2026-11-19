package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"zero", 0},
		{"one", 1},
		{"ten", 10},
		{"thousand", 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)

			if len(got) != tt.size {
				t.Fatalf("generateRandomElements(%d): длина = %d, ожидалось %d",
					tt.size, len(got), tt.size)
			}

			for i, value := range got {
				if value <= 0 {
					t.Fatalf("generateRandomElements(%d): элемент %d = %d, ожидалось положительное",
						tt.size, i, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"empty", nil, 0},
		{"one element", []int{42}, 42},
		{"positive", []int{3, 1, 7, 2}, 7},
		{"negative", []int{-5, -2, -9}, -2},
		{"all equal", []int{5, 5, 5}, 5},
		{"zero", []int{0}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.input)
			if got != tt.want {
				t.Fatalf("maximum(%v) = %d, ожидалось %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"empty", nil, 0},
		{"one element", []int{42}, 42},
		{"less than chunks", []int{4, 1, 9, 2}, 9},
		{"exactly chunks", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"more than chunks", []int{10, 2, 30, 4, 50, 6, 70, 8, 90, 1}, 90},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.input)
			if got != tt.want {
				t.Fatalf("maxChunks(%v) = %d, ожидалось %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	data := generateRandomElements(10_000)

	got := maxChunks(data)
	want := maximum(data)

	if got != want {
		t.Fatalf("maxChunks() = %d, maximum() = %d", got, want)
	}
}