package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
import (
	"math"
)

const (
	SidesCircle   int = 0
	SidesTriangle int = 3
	SidesSquare   int = 4
)

func CalcSquare(sideLen float64, sidesNum int) float64 {

	if sidesNum == SidesCircle {
		return math.Pi * math.Pow(sideLen, 2)
	} else if sidesNum == SidesSquare {
		return math.Pow(sideLen, 2)
	} else if sidesNum == SidesTriangle {
		return math.Sqrt(3) / 4 * math.Pow(sideLen, 2)
	}
	return 0
}
