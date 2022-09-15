package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	Maxconns int
	TLS      *tls.Config
}

type Option func(*Server)

func main() {
	s1, _ := NerServer("localhost", 1024)
	fmt.Printf("S1: %+v \n", s1)

	s2, _ := NerServer("localhost", 2048, Protocol("udp"))
	fmt.Printf("S2: %+v \n", s2)

	s3, _ := NerServer("localhost", 3000, Timeout(300*time.Second), MaxConns(1000))
	fmt.Printf("S3: %+v \n", s3)
}

func Protocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}

func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func MaxConns(maxconns int) Option {
	return func(s *Server) {
		s.Maxconns = maxconns
	}
}

func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

func NerServer(addr string, port int, options ...func(*Server)) (*Server, error) {
	server := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		Maxconns: 1000,
		TLS:      nil,
	}

	for _, option := range options {
		option(&server)
	}

	return &server, nil
}
