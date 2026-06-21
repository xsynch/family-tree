package main

import (
	"net/http"
	"time"
)

func NewServer(address string, handler http.Handler) (*http.Server) {
	return &http.Server{
		Addr: address ,
		Handler: handler ,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	} 
}