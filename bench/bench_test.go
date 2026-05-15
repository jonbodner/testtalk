package bench

import (
	"fmt"
	"testing"
)

func TestFileLen(t *testing.T) {
	result, err := FileLen("testdata/data.txt", 1)
	if err != nil {
		t.Fatal(err)
	}
	if result != 65204 {
		t.Error("Expected 65204, got", result)
	}
}

func BenchmarkFileLen1(b *testing.B) {
	for b.Loop() {
		_, err := FileLen("testdata/data.txt", 1)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFileLen(b *testing.B) {
	for _, v := range []int{1, 10, 100, 1000, 10000, 100000} {
		b.Run(fmt.Sprintf("FileLen-%d", v), func(b *testing.B) {
			for b.Loop() {
				_, err := FileLen("testdata/data.txt", v)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
