// Created By: junmin.ma
// Description: <description>
// Date: 2021-11-29 11:53
package option

import (
	"testing"
	"time"
)

func TestOption(t *testing.T) {
	srv := New(WithHost("localhsot"), WithPort(8081), WithTimeout(time.Minute), WithMaxConn(20))
	_ = srv.Start()
}
