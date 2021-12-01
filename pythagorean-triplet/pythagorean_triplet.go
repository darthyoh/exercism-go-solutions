package pythagorean

import "math"

//Triplet type
type Triplet [3]int

func (t Triplet) isSum(p int) bool {
	return t[0]+t[1]+t[2] == p
}

//Range function get all triplets in a range
func Range(min, max int) []Triplet {

	triplets := make([]Triplet, 0)

	for i := min; i < max-1; i++ {
		for j := i + 1; j < max; j++ {
			for k := j + 1; k <= max; k++ {
				if math.Pow(float64(i), 2)+math.Pow(float64(j), 2) == math.Pow(float64(k), 2) {
					triplets = append(triplets, [3]int{i, j, k})
				}
			}
		}
	}
	return triplets
}

//Sum function get all triplets that satisfy a+b+c=perimeter
func Sum(p int) []Triplet {

	triplets := make([]Triplet, 0)

	for _, t := range Range(0, p) {
		if t.isSum(p) {
			triplets = append(triplets, t)
		}
	}
	return triplets
}
