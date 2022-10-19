package complex

import "math"

// Define the Number type here.
type Number struct {
	a, b float64
}

func (n Number) Real() float64 {
	return n.a
}

func (n Number) Imaginary() float64 {
	return n.b
}

func (n1 Number) Add(n2 Number) Number {
	return Number{n1.a + n2.a, n1.b + n2.b}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{n1.a - n2.a, n1.b - n2.b}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{(n1.a*n2.a - n1.b*n2.b), (n1.b*n2.a + n1.a*n2.b)}
}

func (n Number) Times(factor float64) Number {
	return Number{n.a * factor, n.b * factor}
}

func (n1 Number) Divide(n2 Number) Number {
	return Number{(n1.a*n2.a + n1.b*n2.b) / (math.Pow(n2.a, 2) + math.Pow(n2.b, 2)), (n1.b*n2.a - n1.a*n2.b) / (math.Pow(n2.a, 2) + math.Pow(n2.b, 2))}
}

func (n Number) Conjugate() Number {
	return Number{n.a, -n.b}
}

func (n Number) Abs() float64 {
	return math.Sqrt(math.Pow(n.a, 2) + math.Pow(n.b, 2))
}

func (n Number) Exp() Number {
	return Number{math.Exp(n.a) * math.Cos(n.b), math.Exp(n.a) * math.Sin(n.b)}
}
