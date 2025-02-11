package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type custom int

const SidesTriangle custom = 3
const SidesSquare custom = 4
const SidesCircle custom = 0

func CalcSquare(sideLen float64, sidesNum custom) float64 {
	switch sidesNum {
	case SidesTriangle:
		return math.Pow(sideLen, 2) * math.Sqrt(3) / 4
	case SidesSquare:
		return sideLen * sideLen
	case SidesCircle:
		return math.Pi * (sideLen * sideLen)
	}

	return 0
}
