package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTimePerLayer int)int{
    if(avgPrepTimePerLayer == 0){
        return len(layers) * 2
    }
    return len(layers) * avgPrepTimePerLayer
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
    noodles = 0
    sauce = 0.0
    for _,val := range layers {
        if val == "noodles" {
            noodles += 50
        }else if val == "sauce" {
            sauce += 0.2
        }
    }
    return
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string){
     myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	factor := float64(portions) / 2.0   

	scaled := make([]float64, len(quantities))

	for i, q := range quantities {
		scaled[i] = q * factor
	}

	return scaled
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
