package lib

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
)

// output:
// Get Conn: golang.org:443
// DNS Info: {Host:golang.org}
// DNS Info: {Addrs:[{IP:172.217.160.81 Zone:} {IP:2404:6800:4008:803::2011 Zone:}] Err:<nil> Coalesced:false}
// ConnectStart Info: tcp, 172.217.160.81:443
// ConnectDone Info: tcp, 172.217.160.81:443, <nil>
// Got Conn: {Conn:0xc00006e000 Reused:false WasIdle:false IdleTime:0s}
// GotFirstResponseByte:
func HttpTrace() {
	req, _ := http.NewRequest("GET", "https://golang.org/", nil)
	trace := &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			fmt.Printf("Get Conn: %+v\n", hostPort)
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},
		PutIdleConn: func(err error) {
			fmt.Printf("PutIdleConn: %+v\n", err)
		},
		GotFirstResponseByte: func() {
			fmt.Printf("GotFirstResponseByte: \n")
		},
		Got100Continue: func() {
			fmt.Printf("Got100Continue: \n")
		},
		DNSStart: func(dnsInfo httptrace.DNSStartInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
		ConnectStart: func(network, addr string) {
			fmt.Printf("ConnectStart Info: %+v, %+v\n", network, addr)
		},
		ConnectDone: func(network, addr string, err error) {
			fmt.Printf("ConnectDone Info: %+v, %+v, %+v\n", network, addr, err)
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	_, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}
}
