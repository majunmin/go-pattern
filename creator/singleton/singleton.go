// Created By: junmin.ma
// Description: <description>
// Date: 2021-11-24 16:14
package singleton

import "sync"

// 首字母小写 避免外部初始化
type singleton struct {
}

var once sync.Once
var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		once.Do(func() {
			instance = new(singleton)
		})
	}
	return instance
}
