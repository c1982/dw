package whois

import (
	"io/ioutil"
	"net"
	"regexp"
	"time"
)

const defaultWhoisPort = "43"
const defaultTimeOut = time.Second * 8

func whoisConn(nic, domain string) (data string, err error) {

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(nic, defaultWhoisPort), defaultTimeOut)

	if err != nil {
		return data, err
	}

	defer conn.Close()

	conn.Write([]byte(domain + "\r\n"))

	bf, err := ioutil.ReadAll(conn)

	if err != nil {
		return data, err
	}

	data = string(bf[:])

	return
}

func parseNameServers(data string) (ns []string) {

	//regex pattern copy from https://github.com/undiabler/golang-whois/blob/master/extra.go#L37
	rx := regexp.MustCompile(`(?i)Name Server:\s+(.*?)(\s|$)`)

	list := rx.FindAllStringSubmatch(data, -1)

	for i := 0; i < len(list); i++ {
		if len(list[i]) >= 1 {
			ns = append(ns, list[i][1])
		}
	}

	return
}

func parserRegistatWhoisServer(data string) (server string) {
	rx := regexp.MustCompile(`(?i)Registrar WHOIS Server:\s+(.*?)(\s|$)`)

	list := rx.FindAllStringSubmatch(data, 1)

	if len(list) == 0 {
		return
	}

	if len(list[0]) == 0 {
		return
	}

	server = list[0][1]

	return
}

//FIXME: in error for two parts domain extension.
func parseDomainNameExtension(data string) string {

	rx := regexp.MustCompile(`(\.[^.]+)$`)
	list := rx.FindAllStringSubmatch(data, -1)

	if len(list) == 0 {
		return ""
	}

	if len(list[0]) == 0 {
		return ""
	}

	return list[0][1]
}
