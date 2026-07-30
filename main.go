/*
Copyright © 2021 Daniel Bos <corani@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"log"
	"os"

	"github.com/corani/adr/cmd"
	"github.com/corani/adr/internal/config"
	"github.com/spf13/cobra"
)

func newRootCommand() *cobra.Command {
	conf, err := config.ReadConfig()
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("couldn't read config: %v", err)
	}

	//nolint:exhaustruct
	root := &cobra.Command{
		Use:   os.Args[0],
		Short: "A command line tool to maintain Architecture Decision Records",
		PersistentPreRunE: func(c *cobra.Command, _ []string) error {
			switch c.Name() {
			case "init", "version", "help", "completion":
				return nil
			}

			if conf == nil {
				log.Fatal("no ADR configuration found, run 'adr init' first")
			}

			return nil
		},
	}

	root.AddCommand(cmd.AdrCommands(conf)...)

	return root
}

func main() {
	if err := newRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
