/*
Copyright 2021 Lars Eric Scheidler

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"time"

	"github.com/lscheidler/go-nagios"
)

// Check data structure
type check struct {
	Critical      float64
	Debug         bool
	nagios        *nagios.Nagios
	Warning       float64
	ec2           bool
	rds           bool
	elasticache   bool
	redshift      bool
	elasticsearch bool
}

func new() *check {
	c := check{
		Critical:      25.0,
		Debug:         true,
		nagios:        nagios.New(),
		Warning:       50.0,
		ec2:           true,
		rds:           false,
		elasticache:   false,
		redshift:      false,
		elasticsearch: false,
	}
	parseArguments(&c)
	return &c
}

func (c *check) Run() {
	t := time.Now()
	thisMonthFirst := t.AddDate(0, 0, -1*(t.Day()-1)).Format("2006-01-02")
	lastMonthFirst := t.AddDate(0, -1, -1*(t.Day()-1)).Format("2006-01-02")

	if c.ec2 {
		c.checkService("ec2", "EC2NotCovered", lastMonthFirst, thisMonthFirst)
	}
	if c.rds {
		c.checkService("rds", "RDSNotCovered", lastMonthFirst, thisMonthFirst)
	}
	if c.elasticache {
		c.checkService("elasticache", "ElasticacheNotCovered", lastMonthFirst, thisMonthFirst)
	}
	if c.redshift {
		c.checkService("redshift", "RedshiftNotCovered", lastMonthFirst, thisMonthFirst)
	}
	if c.elasticsearch {
		c.checkService("elasticsearch", "ElasticsearchNotCovered", lastMonthFirst, thisMonthFirst)
	}
}

func (c *check) checkService(key string, name string, start string, stop string) {
	if rc, err := getReservationCoverage(key, start, stop); err != nil {
		c.nagios.Unknown(err.Error())
	} else {
		c.nagios.CheckThreshold(name, 100-*rc, 100-c.Warning, 100-c.Critical)
	}
}

// Exit defer function
func (c *check) Exit() {
	c.nagios.Exit()
}
