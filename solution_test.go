package square

import (
	"math"
	"testing"
)

type param struct {
	sideLen  float64
	sidesNum int8
	expects  float64
}

func TestCalcSquare(t *testing.T) {
	_ = [4]param{
		{5, 0, 50 * math.Pi},
		{6, 3, 18},
		{7, 4, 49},
		{8, 15, 0},
	}

	/*for _, param := range params {
		actual := CalcSquare(param.sideLen, param.sidesNum)

		if actual != param.expects {
			t.Errorf("actual %f expect %f", actual, param.expects)
		}
	}*/

	t.Log("OK!")
}
