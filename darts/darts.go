package darts

//Score returns the dart score from an x/y position
func Score(x, y float64) int {

	switch {
	case x*x+y*y <= 1:
		return 10
	case x*x+y*y <= 25:
		return 5
	case x*x+y*y <= 100:
		return 1
	default:
		return 0
	}
}
