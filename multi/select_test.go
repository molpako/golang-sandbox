package multi

import (
	"testing"
)

func TestSuccessive(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Successive(); got != tt.want {
				t.Errorf("Successive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcurrency(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Concurrency(); got != tt.want {
				t.Errorf("Cunccurency() = %v, want %v", got, tt.want)
			}
		})
	}
}
