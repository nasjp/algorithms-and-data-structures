package chapter1_test

import (
	"reflect"
	"testing"

	"github.com/nasjp/algorithms-and-data-structures/chapter1"
)

func TestSieveOfEratosthenes(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want []int
	}{
		{"Case1", 5, []int{2, 3, 5}},
		{"Case2", 30, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := chapter1.SieveOfEratosthenes(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
