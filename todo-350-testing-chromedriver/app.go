package main

import (
	browser "testing-chromedriver/site"
)

const LOGIN_URL = "http://dev.ioki.com.pl/login"
const USERNAME = "knowak"
const PASSWORD = "ngml2"

func main() {
	nmel := browser.NewNmel()
	nmel.Goto(LOGIN_URL)
	nmel.Login(USERNAME, PASSWORD)
	nmel.Quit()
}
