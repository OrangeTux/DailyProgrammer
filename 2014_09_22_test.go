package main

import (
	"testing"
)

func Test_parse_equation_ax_b(t *testing.T) {
	a, b := parse_equation("2x+3")
	if a != 2 && b != 3 {
		t.Error("parse_eqation(2x+3) does not return 2, 3.")
	}
}

func Test_find_intersection(t *testing.T) {
	x, y := find_intersection("2x+2", "5x-4")
	if x != 2 || y != 6 {
		t.Error("find_intersection(2x+2 5x-4) does not return 2, 6.")
	}
}

func Test_find_x_intersect(t *testing.T) {
	if find_x_intersect(2, 5, 2, -4) != 2 {
		t.Error("find_y_intersection(2, 3, 3) does not return 2.")
	}
}

func Test_find__y_intersect(t *testing.T) {
	if find_y_intersect(2, 3, 3) != 9 {
		t.Error("find_y_intersection(2, 3, 3) does not return 9.")
	}
}
