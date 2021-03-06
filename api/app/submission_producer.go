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

package app

import (
	"github.com/infomark-org/infomark/configuration"
	"github.com/infomark-org/infomark/service"
)

// Producer is interface to pipe the workload over AMPQ to the backend workers
type Producer interface {
	Publish(body []byte) error
}

// DefaultSubmissionProducer is the producer which broadcasts all submissions
// amongst the workers.
var DefaultSubmissionProducer Producer

// VoidProducer acts like a real producer, but will not trigger any background worker
// if you do not need these or within tests
type VoidProducer struct{}

// Publish of VoidProducer does nothing on purpose (used in unit tests).
func (t *VoidProducer) Publish(body []byte) error { return nil }

func InitSubmissionProducer() {
	var err error

	cfg := service.NewConfig(&configuration.Configuration.Server.Services.RabbitMQ)

	if configuration.Configuration.Server.DistributeJobs {
		DefaultSubmissionProducer, err = service.NewProducer(cfg)
		if err != nil {
			panic(err)
		}
	} else {
		DefaultSubmissionProducer = &VoidProducer{}

	}

}
