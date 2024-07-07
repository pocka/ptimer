// SPDX-FileCopyrightText: 2024 Shota FUJI <pockawoooh@gmail.com>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"tools/ptimer/create"
	"tools/ptimer/extract"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use: "ptimer",
	}

	cmd.AddCommand(create.Command())
	cmd.AddCommand(extract.Command())

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
