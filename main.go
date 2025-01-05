package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/etum-dev/paramarama/utils"
	"github.com/urfave/cli/v3"
)

func init() {
	cli.VersionFlag = &cli.BoolFlag{Name: "print-version", Aliases: []string{"V"}}

	cli.VersionPrinter = func(cmd *cli.Command) {
		fmt.Fprintf(cmd.Root().Writer, "version=%s\n", cmd.Root().Version)
	}
}

func main() {
	cmd := &cli.Command{
		Name:      "",
		Version:   "v0.0.0",
		Usage:     "cli",
		ArgsUsage: "[args and such]",
		Commands:  []*cli.Command{},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "path",
				Usage: "Path to JS file",
			},
			&cli.StringFlag{
				Name:  "url",
				Usage: "Path to JS url",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if cmd.NArg() <= 0 {
				log.Println("no arg")
			}
			if cmd.String("path") == "qwe" {
				filename := cmd.Args().Get(1)
				log.Println("reading file:", filename)
			}
			if cmd.String("url") == "" {
				utils.MakeWebRequest("https://example.com")
			} else {
				log.Println("???")
			}
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
