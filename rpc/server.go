package rpc

import (
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"shendu.com/log"
	// _ "shendu.com/metrics/influxdb"
)

type Server struct {
	*grpc.Server
}

func (s *Server) Close() error {
	s.Server.Stop()
	return nil
}

func NewServer() *Server {
	srv := &Server{Server: grpc.NewServer()}
	return srv
}

func (s *Server) Serve(port int) {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Error(err)
	}
	fmt.Println("listen on port ===>:", port)
	s.Server.Serve(listen)
}

func NewClient(addrs ...string) *grpc.ClientConn {
	dialOpts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBackoffMaxDelay(time.Second),
		grpc.WithTimeout(time.Second),
		grpc.WithBalancer(grpc.RoundRobin(NewResolver(addrs...))),
	}
	conn, err := grpc.Dial("", dialOpts...)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
