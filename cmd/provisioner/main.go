package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

const (
	flagLVName    = "lvname"
	flagLVSize    = "lvsize"
	flagVGName    = "vgname"
	flagPVS       = "pvs"
	flagDirectory = "directory"
)

func cmdNotFound(c *cli.Context, command string) {
	panic(fmt.Errorf("Unrecognized command: %s", command))
}

func onUsageError(c *cli.Context, err error, isSubcommand bool) error {
	panic(fmt.Errorf("Usage error, please check your command"))
}

func main() {
	p := cli.NewApp()
	p.Usage = "LVM Provisioner Pod"
	p.Commands = []cli.Command{
		createLVCmd(),
		deleteLVCmd(),
	}
	p.CommandNotFound = cmdNotFound
	p.OnUsageError = onUsageError

	if err := p.Run(os.Args); err != nil {
		log.Fatalf("Critical error: %v", err)
	}
}