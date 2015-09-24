package main

import "fmt"

func main() {
	userFavouriteDrinks := map[string][]string{
		"jacek": {"śliwowica", "browar"},
		"kazik": {"metanol", "denaturat", "woda z jeziora"},
		"john":  {"whiskey"},
	}

	// krótka składnia w wartościach weszła od go 1.4? jakoś tak
	// @todo sprawdź to :)
	// w starszych wersjach trzeba by []string{"śliwowica", "browar"}

	for user, drinks := range userFavouriteDrinks {
		fmt.Println("Ulubione picie użytkownika ", user)
		for _, drink := range drinks {
			fmt.Println("-", drink)
		}
	}
}
