package allergies

var list = [8]string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}

//Allergies returns allergies array from score
func Allergies(score uint) (allergies []string) {

	allergies = make([]string, 0)

	for i, v := range list {
		if score>>i%2 != 0 {
			allergies = append(allergies, v)
		}
	}

	return
}

//AllergicTo returns true if score contains candidate allergie
func AllergicTo(score uint, candidate string) bool {
	allergies := Allergies(score)
	for _, v := range allergies {
		if v == candidate {
			return true
		}
	}
	return false
}
