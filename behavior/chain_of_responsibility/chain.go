// Created By: junmin.ma
// Description: <description>
// Date: 2022-01-08 20:47
package chain_of_responsibility

import "strings"

// Context  携带消息的 上下文信息
type Context struct {
	param string
}

// Handler provider a handle interface
type Handler interface {
	// handle msg: 上下文信息
	handle(msg *Context) string
}

type ConcreteHandlerA struct {
	next Handler
}

func (c *ConcreteHandlerA) handle(ctx *Context) string {
	if strings.Contains(ctx.param, "3") {
		// ConcreteHandlerA doSomething
		return "c"
	}
	if c.next != nil {
		return c.next.handle(ctx)
	}

	return ""
}

type ConcreteHandlerB struct {
	next Handler
}

func (c *ConcreteHandlerB) handle(ctx *Context) string {
	if strings.Contains(ctx.param, "2") {
		// ConcreteHandlerB doSomething
		return "b"
	}
	if c.next != nil {
		return c.next.handle(ctx)
	}

	return ""
}

type ConcreteHandlerC struct {
	next Handler
}

func (c *ConcreteHandlerC) handle(ctx *Context) string {
	if strings.Contains(ctx.param, "3") {
		// ConcreteHandlerC doSomething
		return "c"
	}
	if c.next != nil {
		return c.next.handle(ctx)
	}

	return ""
}
