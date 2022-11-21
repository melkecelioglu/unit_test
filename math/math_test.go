package math

import "testing"

type AddData struct {
	x, y   int
	result int
}

func TestAdd(t *testing.T) {
	/*
		result := Add(1, 3)
		if result != 4 {
			t.Errorf("Add(1,3) FAILED. Expected %d, got %d\n", 4, result)
		} else {
			t.Logf("Add(1,3) PASSED. Expected %d, got %d\n", 4, result)
		}*/
	testData := []AddData{
		{1, 2, 3},
		{3, 5, 8},
		{7, -4, 3},
	}
	for _, datum := range testData {
		result := Add(datum.x, datum.y)
		if result != datum.result {
			t.Errorf("Add(%d, %d) FAILED. Expected %d got %d\n",
				datum.x, datum.y, datum.result, result)
		}
	}

}
func TestSubs(t *testing.T) {
	result := Subs(5, 4)
	if result != 1 {
		t.Errorf("Substract(5,4) FAILED. Expected %d, got %d\n", 1, result)
	} else {
		t.Logf("Substract(5,4) PASSED. Expected %d, got %d\n", 1, result)
	}
}

func TestDivide(t *testing.T) {
	result := Divide(6, 0)
	if result != 0 {
		t.Errorf("Divide(6,3) FAILED. Expected %d, got %f\n", 0, result)
	} else {
		t.Logf("Substract(6,3) PASSED. Expected %d, got %f\n", 0, result)
	}
}

func TestMultiply(t *testing.T) {
	testData := []AddData{
		{3, 1, 3},
		{4, 2, 8},
		{7, 0, 0},
	}
	for _, datum := range testData {
		result := Multiply(datum.x, datum.y)
		if result != datum.result {
			t.Errorf("Multiply(%d, %d) FAILED. Expected %d got %d\n",
				datum.x, datum.y, datum.result, result)
		}
	}
}
