package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, funcName func(sideLen float64) float64) float64 {
	if sideLen == 0 {
		return 0
	}

	return funcName(sideLen)
}

func SidesSquare(sideLen float64) float64 {
	return sideLen * sideLen
}

func SidesTriangle(sideLen float64) float64 {
	return sideLen * sideLen / 2
}

func SidesCircle(sideLen float64) float64 {
	return 2 * math.Pi * math.Pow(sideLen, 2)
}
