// Copyright (c) 2015-2024 MinIO, Inc.
//
// This file is part of MinIO Object Storage stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"os"

	"github.com/minio/cli"
)

// AppName is the name of the MinIO server application.
const AppName = "minio"

// AppVersion is set at build time via ldflags.
var AppVersion = "DEVELOPMENT.GOGET"

// Main is the entry point for the minio server.
func Main(args []string) {
	// Set the minio app info.
	app := registerApp(AppName, appCmds)
	app.Before = func(ctx *cli.Context) error {
		return nil
	}

	// Run the app - exit with code 1 on any error.
	if err := app.Run(args); err != nil {
		os.Exit(1) //nolint:gocritic
	}
}

// appCmds contains the list of commands supported by the minio server.
var appCmds = []cli.Command{
	serverCmd,
	gatewayCmd,
}

// registerApp creates and configures the CLI application.
func registerApp(name string, appCmds []cli.Command) *cli.App {
	app := cli.NewApp()
	app.Name = name
	app.Author = "MinIO, Inc."
	app.Version = AppVersion
	app.Usage = "High Performance Object Storage"
	app.Description = `MinIO is a High Performance Object Storage released under GNU AGPLv3.
  It is API compatible with Amazon S3 cloud storage service. Use MinIO to build
  high performance infrastructure for machine learning, analytics and application
  data workloads.`
	app.Commands = appCmds
	app.CustomAppHelpTemplate = minioHelpTemplate
	// Hide the version flag from the help output to reduce noise.
	app.HideVersion = true
	// Show help if no subcommand is provided, rather than silently doing nothing.
	app.Action = cli.ShowAppHelp
	// Enable bash/zsh completion support.
	app.EnableBashCompletion = true
	return app
}

// minioHelpTemplate is the custom help template for the minio CLI.
// NOTE(personal): Added EXAMPLES section to make it easier to remember
// common invocations without digging through the docs every time.
var minioHelpTemplate = `NAME:
  {{.Name}} - {{.Usage}}

DESCRIPTION:
  {{.Description}}

USAGE:
  {{.Name}} - {{.Usage}}

COMMANDS:
  {{range .Commands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
  {{end}}
EXAMPLES:
  Start a standalone server:
    $ minio server /data

  Start a server with a custom address:
    $ minio server --address :9090 /data

  Start a distributed setup with 4 nodes (erasure coding):
    $ minio server http://node{1...4}/data

  Enable TLS with custom certs directory:
    $ minio server --certs-dir /etc/minio/certs /data

VERSION:
  {{.Version}}

{{"COPYRIGHT:"}}
  Copyright (c) 2015-2024 MinIO, Inc.
`
