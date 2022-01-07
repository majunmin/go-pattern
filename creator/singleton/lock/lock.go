// Created By: junmin.ma
// Description: <description>
// Date: 2021-11-24 16:18
package lock

import (
	"sync"
	"sync/atomic"
)

type singleton struct {
}

var instance *singleton
var mux sync.Mutex

func GetInstance() *singleton {
	if instance == nil { // not perfect,  since it's not fully atomic
		mux.Lock()
		defer mux.Unlock()
		if instance == nil {
			instance = &singleton{} // 这一步非原子操作
		}
	}

	return instance
}

var initialize uint32

// 使用  原子变量
func GetInstance2() *singleton {
	if atomic.LoadUint32(&initialize) == 1 {
		return instance
	}
	mux.Lock()
	defer mux.Unlock()
	if initialize == 0 {
		instance = &singleton{}
		atomic.StoreUint32(&initialize, 1)
	}
	return instance
}
