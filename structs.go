package gogo

import "sync"

// Juju defines the cluster name, which bundle to use, and the manifest for credentials and cloud
type Juju struct {
	Kind   string // should be "maas" or "aws" - will be used to figure out which creds and cloud to set
	Name   string
	Bundle string // ex "cs:bundle/canonical-kubernetes-193"
	p      Parallel
	MaasCl MaasCloud
	MaasCr MaasCredentials
	AwsCl  AWSCloud
	AwsCr  AWSCredentials
}

// MaasCloud information
type MaasCloud struct {
	Type     string
	Endpoint string
}

// MaasCredentials to be used with MaasCloud
type MaasCredentials struct {
	CloudName string
	Username  string
	MaasOauth string
}

// AWSCredentials information
type AWSCredentials struct {
	Username  string
	AccessKey string
	SecretKey string
}

// AWSCloud information to be used with AWS Creds
type AWSCloud struct {
	Region string
}

// Parallel sets the waitgroup if user wishes to bring up several clusters at once
type Parallel struct {
	wg sync.WaitGroup
}

// the following structs are for json parsing used with GetJujuStatus()

type jujuStatus struct {
	ApplicationResults map[string]applications `json:"applications"`
	Machines           map[string]machines     `json:"machines"`
}

type machines struct {
	MachStatus map[string]string `json:"juju-status"`
}

type applications struct {
	AppStatus map[string]string `json:"application-status"`
}