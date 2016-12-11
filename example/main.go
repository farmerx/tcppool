package main

import (
	"fmt"
	"net"
	"time"

	"github.com/bibinbin/tcppool"
)

func main() {

	//factory 创建连接的方法
	factory := func() (interface{}, error) { return net.Dial("tcp", "127.0.0.1:80") }

	//close 关闭链接的方法
	close := func(v interface{}) error { return v.(net.Conn).Close() }

	//创建一个连接池： 初始化5，最大链接30
	poolConfig := tcppool.InitOptions{
		InitialCap: 5,
		MaxCap:     30,
		Factory:    factory,
		Close:      close,
		//链接最大空闲时间，超过该时间的链接 将会关闭，可避免空闲时链接EOF，自动失效的问题
		IdleTimeout: 15 * time.Second,
	}
	p, err := tcppool.NewChannelPool(poolConfig)
	if err != nil {
		fmt.Println("err=", err)
	}

	//从连接池中取得一个链接
	v, err := p.Get()

	//do something
	//conn=v.(net.Conn)

	//将链接放回连接池中
	p.Put(v)

	//释放连接池中的所有链接
	//p.Release()

	//查看当前链接中的数量
	current := p.Len()
	fmt.Println("len=", current)
}
