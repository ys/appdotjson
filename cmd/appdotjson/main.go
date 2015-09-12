package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/ys/appdotjson/appjson"
)

func main() {
	app := cli.NewApp()
	app.Name = "appdotjson"
	app.Usage = "Play with app.json"
	app.Authors = []cli.Author{
		{
			Name:  "Tools team",
			Email: "tools@heroku.com",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "env",
			Aliases: []string{"e"},
			Usage:   "Create .env file from app.json",
			Action: func(c *cli.Context) {
				path := c.Args().Get(0)
				if path == "" {
					path = "app.json"
				}
				appjsonToEnv(path)
			},
		},
	}
	app.Run(os.Args)
}

func appjsonToEnv(appFile string) {
	appJson, err := appjson.FromFile(appFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Creating .env file ...")
	envContent := ""
	for key, value := range appJson.Env {
		if value.Value != "" {
			fmt.Printf("%v (default: %v):\n", key, value.Value)
		} else {
			fmt.Printf("%v:\n", key)
		}
		var input string
		fmt.Scanln(&input)
		if input == "" {
			input = value.Value
		}
		envContent += fmt.Sprintf("%v=%v\n", key, input)
	}
	fmt.Println(envContent)
}
