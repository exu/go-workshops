package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

const IntroDirectory = "000-intro"
const GithubRoot = "https://github.com/exu/go-workshops/tree/master"

func getTitleFromDir(input string) (title string) {
	title = strings.Replace(input[4:], "-", " ", -1)
	title = strings.Title(title)
	return
}

func getLink(dir string) string {
	title := getTitleFromDir(dir)
	return fmt.Sprintf("([code for: %s](%s/%s))", title, GithubRoot, dir)
}

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err.Error())
	}

	content := "# Golang Training Resources\n\n"

	r, _ := regexp.Compile("^[0-9]{3}-[a-z0-9-]+")

	contentBytes, err := ioutil.ReadFile(IntroDirectory + "/README.md")
	content += string(contentBytes)
	if err != nil {
		panic(err.Error())
	}

	for _, file := range files {
		if file.IsDir() && r.MatchString(file.Name()) && file.Name() != IntroDirectory {
			dir := file.Name()
			readmeBytes, err := ioutil.ReadFile(dir + "/README.md")
			title := getTitleFromDir(dir)
			link := getLink(dir)
			if err != nil {
				content += fmt.Sprintf("\n\n## %s\n\n", title)
				content += fmt.Sprintf("\n\nThere is no README.md for %s use %s link to follow code examples\n\n", title, link)
			} else {
				readme := string(readmeBytes)
				content += fmt.Sprintf("\n\n%s\n\n\n%s", readme, getLink(dir))
			}
		}
	}

	ioutil.WriteFile("README.md", []byte(content), 0600)
}
