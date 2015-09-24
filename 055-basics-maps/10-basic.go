package main

import "fmt"

func main() {
	userFavouriteDrinks := map[string][]string{}

	userFavouriteDrinks["jacek"] = []string{"śliwowica", "browar"}
	userFavouriteDrinks["kazik"] = []string{"metanol", "denaturat", "woda z jeziora"}
	userFavouriteDrinks["john"] = []string{"whiskey"}

	for user, drinks := range userFavouriteDrinks {
		fmt.Println("Ulubione picie użytkownika ", user)
		for _, drink := range drinks {
			fmt.Println("-", drink)
		}
	}
}
