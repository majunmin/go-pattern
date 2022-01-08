// Created By: junmin.ma
// Description: <description>
// Date: 2022-01-08 11:03
package adapter

// Target 目标接口  provider an  interface with which system should work.
type Target interface {
	Request() string
}

// Adaptee 一般是是三方系统提供的工具,需要被适配
type Adaptee struct {
}

func (a *Adaptee) SpecialRequest() string {
	return "Specual request."
}

// Adapter 适配器类
type Adapter struct {
	Adaptee
}

func NewAdapter(a Adaptee) *Adapter {
	return &Adapter{a}
}

// Request is an Adapter Method
func (a *Adapter) Request() string {
	// do Before
	res := a.SpecialRequest()
	// do After
	return res
}
