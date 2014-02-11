/*
 * Copyright (C) 2014 Cloudius Systems, Ltd.
 *
 * This work is open source software, licensed under the terms of the
 * BSD license as described in the LICENSE file in the top-level directory.
 */

package main

import "github.com/cloudius-systems/capstan"
import "github.com/cloudius-systems/capstan/qemu"
import "github.com/codegangsta/cli"
import "os"

var (
	VERSION string
)

func main() {
	repo := capstan.NewRepo()
	app := cli.NewApp()
	app.Name = "capstan"
	app.Version = VERSION
	app.Usage = "pack, ship, and run applications in light-weight VMs"
	app.Commands = []cli.Command{
		{
			Name:  "push",
			Usage: "push an image to a repository",
			Action: func(c *cli.Context) {
				if len(c.Args()) != 2 {
					println("usage: capstan push [image-name] [image-file]")
					return
				}
				err := repo.PushImage(c.Args()[0], c.Args()[1])
				if err != nil {
					println(err.Error())
				}
			},
		},
		{
			Name:  "pull",
			Usage: "pull an image to the repository",
			Action: func(c *cli.Context) {
				err := repo.PullImage(c.Args().First())
				if err != nil {
					println(err.Error())
				}
			},
		},
		{
			Name:  "rmi",
			Usage: "delete an image from an repository",
			Action: func(c *cli.Context) {
				repo.RemoveImage(c.Args().First())
			},
		},
		{
			Name:  "run",
			Usage: "launch a VM",
			Action: func(c *cli.Context) {
				cmd := qemu.LaunchVM(repo, c.Args().First())
				cmd.Wait()
			},
		},
		{
			Name:  "build",
			Usage: "build an image",
			Action: func(c *cli.Context) {
				qemu.BuildImage(repo, c.Args().First())
			},
		},
		{
			Name:  "images",
			Usage: "list images",
			Action: func(c *cli.Context) {
				repo.ListImages()
			},
		},
	}
	app.Run(os.Args)
}