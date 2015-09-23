package browser

import (
	"log"
	"time"

	wd "github.com/fedesog/webdriver"
)

type Nmel struct {
	Driver  wd.ChromeDriver
	Session wd.Session
}

func NewNmel() *Nmel {
	chromeDriver := wd.NewChromeDriver("chromedriver")
	err := chromeDriver.Start()
	if err != nil {
		log.Println(err)
	}

	desired := wd.Capabilities{"Platform": "Linux"}
	required := wd.Capabilities{}
	session, err := chromeDriver.NewSession(desired, required)

	if err != nil {
		log.Println(err)
	}

	return &Nmel{Driver: *chromeDriver, Session: *session}
}

func (this *Nmel) Login(user string, password string) (bool, error) {
	element, err := this.Session.FindElement(wd.CSS_Selector, "#username")
	if err != nil {
		return false, err
	}
	element.SendKeys(user)

	element, err = this.Session.FindElement(wd.CSS_Selector, "#password")
	if err != nil {
		return false, err
	}
	element.SendKeys(password)

	log.Printf("Login'in as %s:%s\n", user, password)

	element, err = this.Session.FindElement(wd.CSS_Selector, "#login_button")
	if err != nil {
		return false, err
	}
	element.Click()

	return true, nil
}

func (this *Nmel) Goto(url string) (bool, error) {
	log.Printf("Going to: %s\n", url)
	this.Session.Url(url)

	return true, nil
}

func (this *Nmel) Quit() {
	log.Println("Byebye!")
	time.Sleep(1 * time.Second)
	this.Session.Delete()
	this.Driver.Stop()
	log.Println("Gracefully shuted down!")
}
