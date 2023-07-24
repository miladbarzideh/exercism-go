package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTime int) int {
	if prepTime == 0 {
		prepTime = 2
	}
	return len(layers) * prepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64){
	var noodle int
	var sauce float64
	for _,item := range layers {
		if item == "noodles" {
			noodle += 50
		} else if item == "sauce" {
			sauce += 0.2
		}
	}
	return noodle, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
	lastElement := friendList[len(friendList) - 1]
	myList[len(myList) - 1] = lastElement
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(slice []float64, amount int) []float64 {
	result := []float64{}
	for _, item := range slice {
		value := item * float64(amount) / 2
		result = append(result, value)
	}
	return result
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
