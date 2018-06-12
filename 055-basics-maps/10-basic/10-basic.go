package main

import "fmt"

func main() {

	userFavouriteDrinks := map[string][]string{}

	userFavouriteDrinks["jacek"] = []string{"śliwowica", "beer"}
	userFavouriteDrinks["kazik"] = []string{"metanol", "etanol", "water from the lake"}
	userFavouriteDrinks["john"] = []string{"whiskey"}

	for user, drinks := range userFavouriteDrinks {
		fmt.Println("User", user, "favourite drinks:")
		for _, drink := range drinks {
			fmt.Println("-", drink)
		}
	}
}
