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
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
)

const (
	defaultRegion = "eu-central-1"
)

func getReservationCoverage(service string, start string, end string) (*float64, error) {
	sess, conf := getAwsSession()
	srv := costexplorer.New(sess, conf)

	rs, err := srv.GetReservationCoverage(&costexplorer.GetReservationCoverageInput{
		TimePeriod: &costexplorer.DateInterval{
			End:   aws.String(end),
			Start: aws.String(start),
		},
		Filter: &costexplorer.Expression{
			Dimensions: &costexplorer.DimensionValues{
				Key: aws.String("SERVICE"),
				Values: []*string{
					getServiceName(service),
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	rc, err := strconv.ParseFloat(*rs.Total.CoverageHours.CoverageHoursPercentage, 64)
	if err != nil {
		return nil, err
	}
	return &rc, nil
}

func getServiceName(key string) *string {
	if key == "ec2" {
		return aws.String("Amazon Elastic Compute Cloud - Compute")
	} else if key == "rds" {
		return aws.String("Amazon Relational Database Service")
	} else if key == "elasticache" {
		return aws.String("Amazon ElastiCache")
	} else if key == "redshift" {
		return aws.String("Amazon Redshift")
	} else if key == "elasticsearch" {
		return aws.String("Amazon Elasticsearch Service")
	}
	return nil
}

func getAwsSession() (*session.Session, *aws.Config) {
	conf := &aws.Config{Region: aws.String(defaultRegion)}
	if region := os.Getenv("REGION"); len(region) > 0 {
		log.Println("getAwsSession: set region to ", region)
		conf.Region = aws.String(region)
	}

	sess := session.Must(session.NewSession())
	if role := os.Getenv("ASSUME_ROLE"); len(role) > 0 {
		log.Println("getAwsSession: assume role ", role)
		creds := stscreds.NewCredentials(sess, role)
		conf.Credentials = creds
	}
	return sess, conf
}
