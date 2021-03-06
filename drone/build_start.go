package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/drone/drone/model"
)

var buildStartCmd = cli.Command{
	Name:  "start",
	Usage: "start a build",
	Action: func(c *cli.Context) {
		if err := buildStart(c); err != nil {
			log.Fatalln(err)
		}
	},
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "fork",
			Usage: "fork the build",
		},
	},
}

func buildStart(c *cli.Context) (err error) {
	repo := c.Args().First()
	owner, name, err := parseRepo(repo)
	if err != nil {
		return err
	}
	number, err := strconv.Atoi(c.Args().Get(1))
	if err != nil {
		return err
	}

	client, err := newClient(c)
	if err != nil {
		return err
	}

	var build *model.Build
	if c.Bool("fork") {
		build, err = client.BuildFork(owner, name, number)
	} else {
		build, err = client.BuildStart(owner, name, number)
	}
	if err != nil {
		return err
	}

	fmt.Printf("Starting build %s/%s#%d\n", owner, name, build.Number)
	return nil
}
