package webdriver

import (
	wd "github.com/fedesog/webdriver"
	"log"
	"time"
)

func NewChromeSession(chromeDriver *wd.ChromeDriver) *wd.Session {
	desired := wd.Capabilities{"Platform": "Linux"}
	required := wd.Capabilities{}
	session, err := chromeDriver.NewSession(desired, required)

	if err != nil {
		log.Println(err)
	}

	return session
}

func NewChromeDriver() *ChromeDriver {
	chromeDriver := wd.NewChromeDriver("chromedriver")
	err := chromeDriver.Start()
	if err != nil {
		log.Println(err)
	}

	session := NewChromeSession(chromeDriver)

	return &ChromeDriver{Session: *session, ChromeDriver: *chromeDriver}
}

type ChromeDriver struct {
	Session      wd.Session
	ChromeDriver wd.ChromeDriver
}

func (this *ChromeDriver) Login(user string, password string) (bool, error) {
	element, _ := this.Session.FindElement(wd.CSS_Selector, "#identity")
	element.SendKeys(user)

	element, _ = this.Session.FindElement(wd.CSS_Selector, "#credential")
	element.SendKeys(password)

	log.Printf("Login'in as %s:%s\n", user, password)

	element, _ = this.Session.FindElement(wd.CSS_Selector, "#sign-in-submit")
	element.Click()

	return true, nil
}

func (this *ChromeDriver) Goto(url string) (bool, error) {
	log.Printf("Going to: %s\n", url)
	this.Session.Url(url)

	return true, nil
}

func (this *ChromeDriver) Quit() {
	log.Println("Byebye!")
	time.Sleep(1 * time.Second)
	this.Session.Delete()
	this.ChromeDriver.Stop()
}
