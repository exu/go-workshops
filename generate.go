package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

const IntroDirectory = "000-intro"
const GithubRoot = "https://github.com/exu/go-workshops"

func getLink(dir string) string {
	title := strings.Replace(dir[4:], "-", " ", -1)
	title = strings.Title(title)
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
			if err != nil {
				content += fmt.Sprintf("\n\n# %s\n\n", getLink(dir))
			} else {
				readme := string(readmeBytes)
				readme = strings.Replace(readme, "\n", " "+getLink(dir)+"\n", 1)
				content += "\n\n" + readme
			}
		}
	}

	ioutil.WriteFile("README.md", []byte(content), 0600)
}
