package square

import (
	"math"
	"testing"
)

func TestSidesSquare(t *testing.T) {
	var expects float64 = 100
	actual := CalcSquare(10.0, SidesSquare)

	diff(t, expects, actual)
}

func TestSidesTriangle(t *testing.T) {
	expects := 43.3
	actual := CalcSquare(10.0, SidesTriangle)

	diff(t, expects, actual)
}

func TestSidesCircle(t *testing.T) {
	var expects = math.Pi * 100
	actual := CalcSquare(10.0, SidesCircle)

	diff(t, expects, actual)
}

func diff(t *testing.T, expects, actual float64) {
	if expects-actual > 0.01 {
		t.Errorf("actual %f expect %f", actual, expects)
	}

	t.Log("OK!")
}
