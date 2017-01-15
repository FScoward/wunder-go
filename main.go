package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "wunder-go"

	app.Commands = []cli.Command{
		{
			Name:    "authentication",
			Aliases: []string{"a"},
			Usage:   "authentication",
			Action: func() {
				//return nil
			},
		},
		{
			Name:    "new task",
			Aliases: []string{"n"},
			Usage:   "create new task",
			Action: func(c *cli.Context) error {
				fmt.Println("added task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "show task",
			Aliases: []string{"t"},
			Usage:   "show task",
			Action: func(c *cli.Context) error {
				fmt.Println("show task list...")
				var tasks []Task
				var values = url.Values{}
				values.Add("list_id", "276635689")
				execute("https://a.wunderlist.com/api/v1/tasks", values, &tasks)
				fmt.Println(tasks)
				return nil
			},
		},
		{
			Name:    "show list",
			Aliases: []string{"l"},
			Usage:   "show list",
			Action: func(c *cli.Context) error {
				var list []List
				execute("https://a.wunderlist.com/api/v1/lists", nil, &list)
				fmt.Println(list)
				fmt.Println("show list: ", c.Args().First())
				return nil
			},
		},
	}

	app.Run(os.Args)
}

func execute(url string, values url.Values, typedef interface{}) {
	client := new(http.Client)
	get(client, url, values, &typedef)
}
