package main

import (
	"testing"
)

func TestHelper(t *testing.T) {
	t.Helper()
	t.Log("test finished")
}

func TestSum(t *testing.T) {
	t.Cleanup(func() {
		TestHelper(t)
	})

	a, b := 11, 5
	expected := 16
	actual := Sum(a, b)
	if actual != expected {
		t.Errorf("Sum(%d, %d) expected %d, but got %d\n", a, b, expected, actual)
	}
}

func TestEqual(t *testing.T) {
	t.Cleanup(func() {
		TestHelper(t)
	})

	a, b := 10, 10
	expected := true
	actual := a == b
	if actual != expected {
		t.Errorf("Equal(%d, %d) expected %v, but got %v", a, b, expected, actual)
	}
}

func TestTableStyle(t *testing.T) {
	testData := []struct{
		a, b int
		exp  int
	}{
		{10, 20, 30},
		{5, 5, 10},
		{2, 3, 5},
		{4, 4, 8},
	}

	for _, data := range testData {
		if actual := Sum(data.a, data.b); actual != data.exp {
			t.Errorf("Sum(%d, %d) expected %d, but got %d", data.a, data.b, data.exp, actual)
			t.FailNow()
		} 
	}
}
