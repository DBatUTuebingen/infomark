// InfoMark - a platform for managing courses with
//            distributing exercise sheets and testing exercise submissions
// Copyright (C) 2019 ComputerGraphics Tuebingen
//               2020-present InfoMark.org
// Authors: Patrick Wieschollek
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"log"

	"github.com/infomark-org/infomark/api"
	"github.com/infomark-org/infomark/configuration"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start the infomark RESTful JSON backend server with configured api",
	Long:  `Starts a http server and serves the configured api`,
	Run: func(cmd *cobra.Command, args []string) {

		configuration.MustFindAndReadConfiguration()

		server, err := api.NewServer(&configuration.Configuration.Server)
		if err != nil {
			log.Fatal(err)
		}
		server.Start()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
