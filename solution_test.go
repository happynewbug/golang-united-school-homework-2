package square

import (
	"math"
	"testing"
)

func TestSidesSquare(t *testing.T) {
	var expects float64 = 100
	actual := SidesSquare(10.0)
	if expects != actual {
		t.Errorf("actual %f expect %f", actual, expects)
	}

	t.Log("OK!")
}

func TestSidesTriangle(t *testing.T) {
	var expects float64 = 50
	actual := SidesTriangle(10.0)
	if expects != actual {
		t.Errorf("actual %f expect %f", actual, expects)
	}

	t.Log("OK!")
}

func TestSidesCircle(t *testing.T) {
	var expects = math.Pi * 100
	actual := SidesCircle(10.0)
	if expects != actual {
		t.Errorf("actual %f expect %f", actual, expects)
	}

	t.Log("OK!")
}
