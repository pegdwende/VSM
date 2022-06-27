package database

import "fmt"

type Dns struct {
	username string
	password string
	protocol string
	host     string
	port     string
	database string
}

func (dns Dns) getDnsString() string {

	dnsString := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=true", dns.username, dns.password, dns.protocol, dns.host, dns.port, dns.database)

	return dnsString
}
