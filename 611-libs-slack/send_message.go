package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/user"

	"flag"
	"log"

	"github.com/nlopes/slack"
)

type Config struct {
	Slack struct {
		Token string `json:"pgpt"`
	} `json:"slack"`
}

var channel = flag.String("channel", "G6DRKQ9DF", "Slack Channel")
var token = flag.String("token", "", `Slack API Token (if not set it'll try to read from ~/.auth.json file {"slack":{"pgpt":"XXX"}}`)
var buffer = flag.Int("buffer", 10*1024, `max message size`)
var messagePrefix = flag.String("prefix", "", `message prefix`)

func getTokenFromAuthFile() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	configFile := fmt.Sprintf("%s/.auth.json", usr.HomeDir)
	log.Println("Loading config from file:", configFile)

	file, err := os.Open(configFile)
	if err != nil {
		log.Println("Can't load config data from:", configFile)
	}
	decoder := json.NewDecoder(file)
	config := &Config{}

	err = decoder.Decode(&config)
	if err != nil {
		log.Println("Decode ./config.json error:", err)
	}

	return config.Slack.Token
}

func sendMessage(api *slack.Client, message string) error {
	params := slack.PostMessageParameters{}
	channelID, timestamp, err := api.PostMessage("G6DRKQ9DF", message, params)
	if err != nil {
		return err
	}
	fmt.Printf("%s\nSuccessfully sent to channel %s at %s\n", message, channelID, timestamp)

	return nil
}

func init() {
	flag.Parse()
}

func main() {
	if len(*token) == 0 {
		*token = getTokenFromAuthFile()
	}
	api := slack.New(*token)

	scanner := bufio.NewScanner(os.Stdin)
	messageText := ""

	if len(*messagePrefix) > 0 {
		*messagePrefix = "`" + *messagePrefix + "` "
	}

	for scanner.Scan() {

		messageText += *messagePrefix + scanner.Text() + "\n"

		if len(messageText) >= *buffer {
			err := sendMessage(api, messageText)
			if err != nil {
				panic(err)
			}

			messageText = ""
		}
	}

	if len(messageText) >= 0 {
		err := sendMessage(api, messageText)
		if err != nil {
			panic(err)
		}
	}

}
