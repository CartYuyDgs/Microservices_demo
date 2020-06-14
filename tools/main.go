package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {

	var opt Option

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "f",
			Usage:       "idl filename",
			Value:       "./test.proto",
			Destination: &opt.Proto3Filename,
		},

		cli.StringFlag{
			Name:        "o",
			Value:       "./output",
			Usage:       "output dir",
			Destination: &opt.Output,
		},

		cli.BoolFlag{
			Name:        "c",
			Usage:       "generate grpc client code",
			Destination: &opt.GenClientCode,
		},

		cli.BoolFlag{
			Name:        "s",
			Usage:       "generate grpc server code",
			Destination: &opt.GenServerCode,
		},
	}

	app.Action = func(c *cli.Context) error {
		name := "Nefertiti"
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}
		fmt.Println("cmd : ", name)
		err := generatorMgr.Run(&opt)
		if err != nil {
			fmt.Println("code generator failed: ", err)
			return err
		}
		fmt.Println("code generator success!")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
