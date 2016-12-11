/*
	tcppool interface是对这个连接池的一个描述。也暴露了对外的三个接口。
	一个优秀的连接池要能实现对池子的大小控制，线程取用安全，简单等。
*/

package tcppool

import "errors"

var (
	//ErrClosed 连接池已经关闭Error
	ErrClosed = errors.New("pool is closed")
)

//Pool 基本方法
type TcpPool interface {
	Get() (interface{}, error)

	Put(interface{}) error

	Close(interface{}) error

	Release()

	Len() int
}
