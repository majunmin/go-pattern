// Created By: junmin.ma
// Description: <description>
// Date: 2021-11-29 11:38
package factory

import "testing"

func TestNew(t *testing.T) {
	server := New("localhsot", 8081)
	_ = server.Start()
}
