// Simple package for space age
package space

type Planet string

// Given an age in seconds and a planet, returns the Earth age in years
func Age(seconds float64, planet Planet) float64 {
	coef := 1.0
	switch planet {
	case "Mercury":
		coef = 0.2408467
	case "Venus":
		coef = 0.61519726
	case "Mars":
		coef = 1.8808158
	case "Jupiter":
		coef = 11.862615
	case "Saturn":
		coef = 29.447498
	case "Uranus":
		coef = 84.016846
	case "Neptune":
		coef = 164.79132
	}

	return seconds / (coef * 31557600)
}

/*
//Here's another solution using a Map
//This solution shows less performances in benchmark

var planetMap = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.000000,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	 return seconds / (31557600 * planetMap[planet])
}
*/
