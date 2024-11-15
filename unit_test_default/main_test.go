package main

import "testing"

type TestCase struct { // membuat struct untuk menyimpan data testing
	a        int
	b        int
	expected int
}

func TestSumTwoNumber(t *testing.T) {
	var listTest = []TestCase{ // menyiapkan semua data test case
		{a: 1, b: 2, expected: 3},
		{a: -2, b: 4, expected: 2},
		{a: 2, b: -7, expected: -5},
		{a: -4, b: -5, expected: -9},
	}

	for _, test := range listTest { // melakukan perulangan untuk menjalankan semua test case
		result := SumTwoNumber(test.a, test.b)
		if result != test.expected {
			t.Errorf("Expected %d, but got %d", test.expected, result)
		}
	}
}
