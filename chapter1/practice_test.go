package chapter1_test

import (
	"reflect"
	"testing"

	"github.com/nasjp/algorithms-and-data-structures/chapter1"
)

func TestPractice1_1(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want bool
	}{
		{"Case1", 1, false},
		{"Case2", 2, true},
		{"Case3", 3, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := chapter1.Practice1_1(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestPractice1_2(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want bool
	}{
		{"Case1", 4, false},
		{"Case2", 5, true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := chapter1.Practice1_2(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestPractice1_3(t *testing.T) {
	tests := []struct {
		name string
		arg  []int
		want int
	}{
		{"Case1", []int{1, 2, 3, 4, 5}, 5},
		{"Case2", []int{5, 3, 1, 2, 4}, 5},
		{"Case3", []int{-10, 2, 99999999, -1111111, 4}, 99999999},
		{"Case4", []int{}, 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := chapter1.Practice1_3(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestPractice1_4Mean(t *testing.T) {
	tests := []struct {
		name string
		arg  []int
		want float64
	}{
		{"Case1", []int{1, 2, 3, 4, 5}, 3},
		{"Case2", []int{5, 3, 1, 2, 4}, 3},
		{"Case3", []int{1, 1, 1, 1, 5}, 1.8},
		{"Case4", []int{}, 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := chapter1.Practice1_4Mean(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestPractice1_4Variance(t *testing.T) {
	tests := []struct {
		name string
		arg  []int
		want float64
	}{
		{"Case1", []int{1, 2, 3, 4, 5}, 2},
		{"Case2", []int{5, 3, 1, 2, 4}, 2},
		{"Case3", []int{-3, -2, -1, 0, 1}, 2},
		{"Case4", []int{}, 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := chapter1.Practice1_4Variance(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestPractice1_5(t *testing.T) {
	type arg struct {
		a *chapter1.Point
		b *chapter1.Point
	}

	tests := []struct {
		name string
		arg  arg
		want float64
	}{
		{"Case1", arg{a: &chapter1.Point{X: 0, Y: 0}, b: &chapter1.Point{X: 0, Y: 1}}, 1},
		{"Case2", arg{a: &chapter1.Point{X: 0, Y: 0}, b: &chapter1.Point{X: 1, Y: 1}}, 1.4142135623730951},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := chapter1.Practice1_5(tt.arg.a, tt.arg.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestPractice1_6(t *testing.T) {
	type arg struct {
		a *chapter1.Point
		b *chapter1.Point
		c *chapter1.Point
	}

	tests := []struct {
		name string
		arg  arg
		want float64
	}{
		{"Case1", arg{a: &chapter1.Point{X: 0, Y: 0}, b: &chapter1.Point{X: 2, Y: 0}, c: &chapter1.Point{X: 0, Y: 2}}, 2},
		{"Case2", arg{a: &chapter1.Point{X: -1, Y: -1}, b: &chapter1.Point{X: 1, Y: -1}, c: &chapter1.Point{X: 1, Y: 1}}, 2},
		{"Case3", arg{a: &chapter1.Point{X: 0, Y: 0}, b: &chapter1.Point{X: 0, Y: 0}, c: &chapter1.Point{X: 0, Y: 0}}, 0},
		{"Case4", arg{a: &chapter1.Point{X: -1, Y: -1}, b: &chapter1.Point{X: 0, Y: 0}, c: &chapter1.Point{X: 1, Y: 1}}, 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := chapter1.Practice1_6(tt.arg.a, tt.arg.b, tt.arg.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestPractice1_7(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want bool
	}{
		{"Case1", 2, false},
		{"Case2", 4, true},
		{"Case3", 49, true},
		{"Case4", 169, true},
		{"Case5", 0, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := chapter1.Practice1_7(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestPractice1_8(t *testing.T) {
	type arg struct {
		a int
		b int
	}

	tests := []struct {
		name string
		arg  arg
		want int
	}{
		{"Case1", arg{a: 2, b: 3}, 6},
		{"Case2", arg{a: 5, b: 12}, 60},
		{"Case3", arg{a: 12, b: 5}, 60},
		{"Case4", arg{a: 11, b: 34}, 374},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := chapter1.Practice1_8(tt.arg.a, tt.arg.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestPractice1_10(t *testing.T) {
	type arg struct {
		x int
		y int
	}

	type want struct {
		crane  int
		turtle int
	}

	tests := []struct {
		name string
		arg  arg
		want want
	}{
		{"Case1", arg{x: 20, y: 6}, want{crane: 2, turtle: 4}},
		{"Case2", arg{x: 200, y: 80}, want{crane: 60, turtle: 20}},
		{"Case3", arg{x: 1, y: 1}, want{crane: 0, turtle: 0}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotCrane, gotTurtle := chapter1.Practice1_10(tt.arg.x, tt.arg.y)
			if !reflect.DeepEqual(gotTurtle, tt.want.turtle) {
				t.Errorf("gotTurtle: %v, wantTurtle: %v", gotTurtle, tt.want.turtle)
			}

			if !reflect.DeepEqual(gotCrane, tt.want.crane) {
				t.Errorf("gotCrane: %v, wantCrane: %v", gotCrane, tt.want.crane)
			}
		})
	}
}

func TestPractice1_13(t *testing.T) {
	tests := []struct {
		name string
		arg  []int
		want int
	}{
		{"Case1", []int{1, 2, 3, 4, 5}, 5},
		{"Case2", []int{5, 3, 1, 2, 4}, 5},
		{"Case3", []int{-10, 2, 99999999, -1111111, 4}, 99999999},
		{"Case4", []int{}, 0},
		{"Case5", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := chapter1.Practice1_13(tt.arg...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
