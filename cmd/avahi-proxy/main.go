package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/godbus/dbus/v5"
	"github.com/holoplot/go-avahi"
	"github.com/miekg/dns"
)

func handleDnsRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Compress = false

	switch r.Opcode {
	case dns.OpcodeQuery:
		parseQuery(m)
	}

	w.WriteMsg(m)
}

func mDNSResolver(fqdn string, rrtype int32) avahi.HostName {
	// split the host.domain.local and get the host string only
        spltstr := strings.Split(fqdn, ".")
	host := spltstr[0]
	fqdn = host + ".local"

        conn, err := dbus.SystemBus()
        if err != nil {
                log.Fatalf("Cannot get system bus: %v", err)
        }

        server, err := avahi.ServerNew(conn)
        if err != nil {
                log.Fatalf("Avahi new failed: %v", err)
        }

        hn, err := server.ResolveHostName(avahi.InterfaceUnspec, rrtype, fqdn, rrtype, 0)
        if err != nil {
                log.Println("ResolveHostName() failed: %v", err)
        }

	return hn
}

func parseQuery(m *dns.Msg) {
	for _, q := range m.Question {
		switch q.Qtype {
		case dns.TypeA:
			result := mDNSResolver(q.Name, 0) // 0 integer is ProtoInet (see go-avahi source types.go) 
			ip := result.Address
			if ip != "" {
				rr, err := dns.NewRR(fmt.Sprintf("%s A %s", q.Name, ip))
				if err == nil {
					m.Answer = append(m.Answer, rr)
				}
			}
		case dns.TypeAAAA:
                        result := mDNSResolver(q.Name, 1) // 1 integer is ProtoInet6 (see go-avahi source types.go) 
                        ip := result.Address
                        if ip != "" {
                                rr, err := dns.NewRR(fmt.Sprintf("%s AAAA %s", q.Name, ip))
                                if err == nil {
                                        m.Answer = append(m.Answer, rr)
                                }
                        }
		}
	}
}

func main() {
	// attach request handler func
	dns.HandleFunc("home.lan.", handleDnsRequest)

	// start server
	port := 5354
	server := &dns.Server{Addr: ":" + strconv.Itoa(port), Net: "udp"}
	log.Printf("Starting at %d\n", port)
	err := server.ListenAndServe()
	defer server.Shutdown()
	if err != nil {
		log.Fatalf("Failed to start server: %s\n ", err.Error())
	}

}
