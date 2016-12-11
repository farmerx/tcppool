package tcppool

import (
	"net"
	"reflect"
	"testing"
	"time"
)

// 初始化参数
var options InitOptions

// factory 创建连接的方法
var factory = func() (interface{}, error) { return net.Dial("tcp", "127.0.0.1:80") }

// close 关闭链接的方法
var close = func(v interface{}) error { return v.(net.Conn).Close() }

func init() {
	options = InitOptions{
		InitialCap:  5,
		MaxCap:      30,
		Factory:     factory,
		Close:       close,
		IdleTimeout: 15 * time.Second, // 链接最大空闲时间，超过该时间的链接 将会关闭，可避免空闲时链接EOF，自动失效的问题
	}
}

func TestNewChannelPool(t *testing.T) {

	type args struct {
		options InitOptions
	}
	tests := []struct {
		name    string
		args    args
		want    TCPPool
		wantErr bool
	}{
		// TODO: Add test cases.
		{"tcppool", args{options}, func() TCPPool { pool, _ := NewChannelPool(options); return pool }(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewChannelPool(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewChannelPool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, got) {
				t.Errorf("NewChannelPool() = %v, want %v", got, tt.want)
			}
		})
	}
}
