package redisx

import "time"

type Config = ClusterConf

type ClusterConf []NodeConf

// NodeConf is a single redis node configuration
type NodeConf struct {
	Host         string
	Password     string
	DB           int
	MinIdleConn  int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolSize     int
	PoolTimeout  time.Duration
	// tracing switch
	EnableTrace bool
}
