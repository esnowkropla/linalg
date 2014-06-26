package linalg

import (
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	eye := Ident(3)
	eye2 := eye.Copy()
	if !eye.Eq(eye2) {
		t.Error("Copy not equal to original")
	}
}

func ExampleIdent() {
	eye := Ident(3)
	fmt.Println(eye)
	// Output:
	// 1 0j, 0 0j, 0 0j
	// 0 0j, 1 0j, 0 0j
	// 0 0j, 0 0j, 1 0j
}

func TestMul(t *testing.T) {
	eye := Ident(3)
	test := Init(3, 3, []complex128{1, 4, 7, 2, 5, 8, 3, 6, 9})
	out := Zero(3, 3)
	Mul(test, eye, out)
	if !test.Eq(out) {
		t.Error("Multiplied", test, eye, "expected", test, "got", out)
	}
}

func TestAdd(t *testing.T) {
	first := Init(3, 3, []complex128{2, 0, 0, 0, 2, 0, 0, 0, 2})
	second := Ident(3)
	final := Init(3, 3, []complex128{3, 0, 0, 0, 3, 0, 0, 0, 3})
	out := Zero(3, 3)
	Add(first, second, out)
	if !out.Eq(final) {
		t.Error("Added", first, second, "expected", final, "got", out)
	}
}

func TestScale(t *testing.T) {
	eye := Ident(3)
	test := Init(3, 3, []complex128{2, 0, 0, 0, 2, 0, 0, 0, 2})
	eye.Scale(2)
	if !eye.Eq(test) {
		t.Error("Multiplied", Ident(3), "by 2, expected", test, "got", eye)
	}
}
