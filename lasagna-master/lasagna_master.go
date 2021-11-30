package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
    if timePerLayer==0 {
        timePerLayer = 2
    }
	return len(layers)*timePerLayer
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodle int, sauce float64) {
    for _,v := range layers {
        switch v {
            case "noodles":
        		noodle += 50
            case "sauce":
        		sauce +=0.2
        }
    }
	return
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) []string {
    return append(myList,  friendList[len(friendList)-1])
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
	scaled := make([]float64,0)
    for _,v := range amounts {
        scaled = append(scaled, v*float64(portions)/2)
    }
	return scaled
}