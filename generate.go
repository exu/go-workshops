package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

const INTRO_DIR = "000-intro"

func getLink(dir string) string {
	title := strings.Replace(dir[4:], "-", " ", -1)
	title = strings.ToTitle(title)
	return fmt.Sprintf("[%s CODE](%s)", title, dir)
}

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err.Error())
	}

	content := "# Goland Training\n\n"

	r, _ := regexp.Compile("^[0-9]{3}-[a-z0-9-]+")

	contentBytes, err := ioutil.ReadFile(INTRO_DIR + "/README.md")
	content += string(contentBytes)
	if err != nil {
		panic(err.Error())
	}

	for _, file := range files {
		if file.IsDir() && r.MatchString(file.Name()) && file.Name() != INTRO_DIR {
			dir := file.Name()
			readmeBytes, err := ioutil.ReadFile(dir + "/README.md")
			if err != nil {
				content += fmt.Sprintf("\n\n# %s\n\n", getLink(dir))
			} else {
				readme := string(readmeBytes)
				readme = strings.Replace(readme, "\n", " "+getLink(dir)+"\n", 1)
				content += readme
			}
		}
	}

	ioutil.WriteFile("README.md", []byte(content), 0600)
}
