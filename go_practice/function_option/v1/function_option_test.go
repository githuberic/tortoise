package v1

import (
	"crypto/tls"
	"fmt"
	"testing"
	"time"
)

/**
增加一个配置对象
*/
type Config struct {
	Protocol string
	Timeout  time.Duration
	Maxconns int
	TLS      *tls.Config
}
type Server struct {
	Addr string
	Port int
	Conf *Config
}

func NewServer(addr string, port int, conf *Config) (*Server, error) {
	return &Server{Addr: addr, Port: port, Conf: conf}, nil
}

func TestVerify(t *testing.T) {
	//Using the default configuratrion
	srv1, _ := NewServer("localhost", 9000, nil)
	fmt.Print(srv1)

	conf := Config{Protocol: "tcp", Timeout: 60 * time.Duration, 300, nil}
	srv2, _ := NewServer("locahost", 9000, &conf)
	fmt.Print(srv2)
}
