package main

import (
	"github.com/exu/go-workshops/20-testing/todo-10-chromedriver/webdriver"

	"log"
	"os"
	"path/filepath"
)

const LOGIN_URL = "https://dev.ioki.com.pl/login"
const USERNAME = "knowak"
const PASSWORD = "ngml2"

func main() {
	cd := webdriver.NewChromeDriver()
	cd.Goto(LOGIN_URL)
	cd.Login(USERNAME, PASSWORD)

	for _, file := range files {
		cd.Goto(ADD_URL)
		cd.AddFile(file)
	}

	cd.Quit()
}
