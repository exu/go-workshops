package main

import (
	"github.com/exu/go-workshops/20-testing/todo-10-chromedriver/webdriver"
)

const LOGIN_URL = "http://dev.ioki.com.pl/login"
const USERNAME = "knowak"
const PASSWORD = "ngml2"

func main() {
	nmel := webdriver.NewNmel()
	nmel.Goto(LOGIN_URL)
	nmel.Login(USERNAME, PASSWORD)
	nmel.Quit()
}
