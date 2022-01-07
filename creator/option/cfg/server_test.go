// Created By: junmin.ma
// Description: <description>
// Date: 2021-11-29 11:45
package cfg

import (
	"testing"
	"time"
)

func Test_NewServer(t *testing.T) {
	srv := New(Config{
		host:    "localhost",
		port:    2080,
		timeout: time.Minute,
		maxConn: 10,
	})

	srv.Start()
}
