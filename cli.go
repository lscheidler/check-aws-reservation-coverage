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
	"flag"
	"fmt"
	"os"
)

const (
	criticalUsage           = "set critical threshold"
	debugUsage              = "set debug mode"
	debugDefaultVal         = false
	warningUsage            = "set warning threshold"
	ec2Usage                = "check ec2"
	ec2DefaultVal           = true
	rdsUsage                = "check rds"
	rdsDefaultVal           = false
	elasticacheUsage        = "check elasticache"
	elasticacheDefaultVal   = false
	redshiftUsage           = "check redshift"
	redshiftDefaultVal      = false
	elasticsearchUsage      = "check elasticsearch"
	elasticsearchDefaultVal = false
)

func parseArguments(c *check) {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s (%s):\n", os.Args[0], version)
		flag.PrintDefaults()

		examples := `
    `
		fmt.Fprintf(flag.CommandLine.Output(), "%s\n", examples)
	}

	flag.Float64Var(&c.Critical, "critical", 25.0, criticalUsage)
	flag.Float64Var(&c.Critical, "c", 25.0, criticalUsage)
	flag.BoolVar(&c.Debug, "d", debugDefaultVal, debugUsage)
	flag.Float64Var(&c.Warning, "warning", 50.0, warningUsage)
	flag.Float64Var(&c.Warning, "w", 50.0, warningUsage)
	flag.BoolVar(&c.ec2, "ec2", ec2DefaultVal, ec2Usage)
	flag.BoolVar(&c.rds, "rds", rdsDefaultVal, rdsUsage)
	flag.BoolVar(&c.elasticache, "elasticache", elasticacheDefaultVal, elasticacheUsage)
	flag.BoolVar(&c.redshift, "redshift", redshiftDefaultVal, redshiftUsage)
	flag.BoolVar(&c.elasticsearch, "elasticsearch", elasticsearchDefaultVal, elasticsearchUsage)

	flag.Parse()
}
