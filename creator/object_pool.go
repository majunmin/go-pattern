// Created By: junmin.ma
// Description: <description>
// Date: 2022-02-08 20:12
package creator

import "sync"

type Pool struct {
	sync.Mutex
	inuse     []interface{}
	available []interface{}
	new       func() interface{}
}

// NewPool Constructor of Pool
func NewPool(new func() interface{}) *Pool {
	return &Pool{new: new}
}

// Acquire acquires a PoolObject from objectPool
func (p *Pool) Acquire() interface{} {
	p.Lock()
	defer p.Unlock()

	var obj interface{}
	if len(p.available) == 0 {
		obj = p.new()
	} else {
		obj = p.available[0]
		p.available = append(p.available[:0], p.available[1:]) // 防止 地址 变
	}

	p.inuse = append(p.inuse, obj)
	return obj
}

// Release releases a PooledObject to objectPool
func (p *Pool) Release(obj interface{}) {
	p.Lock()
	defer p.Unlock()

	for i := range p.inuse {
		if p.inuse[i] == obj {
			p.inuse = append(p.inuse[:i], p.inuse[i+1:])
		}
	}
	p.available = append(p.available, obj)
}
