package str

import "testing"

func TestPrettyCmplx(t *testing.T) {
	tests := []struct {
		name       string
		inputValue complex128
		want       string
	}{
		// Test cases.
		{"real", 1, "1"},
		{"imaginary", 2i, "2i"},
		{"complex", complex(3, 4), "(3+4i)"},
		{"complex zero", complex(0, 0), "0"},
		{"complex real", complex(5, 0), "5"},
		{"complex imaginary", complex(0, 6), "6i"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrettyCmplx(tt.inputValue); got != tt.want {
				t.Errorf("PrettyCmplx() = %v, want %v", got, tt.want)
			}
		})
	}
}
