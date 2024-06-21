package mobile

import (
	"fmt"
	"log"
	"net"
	"net/netip"

	"github.com/wlynxg/anet"
)

func Greetings(name string) string {
	anet.SetAndroidVersion(14)
	allifaces, err := anet.Interfaces()
	if err != nil {
		return fmt.Sprintf("Error: %s!", err)
	}
	s := ""

	for _, iface := range allifaces {
		s += iface.Name
		addrs, err := anet.InterfaceAddrsByInterface(&iface)
		if err != nil {
			s += fmt.Sprintf("Error: %s", err) + "\n"
		} else {
			s += fmt.Sprintf("%s \n", addrs)
		}

		for _, addr := range addrs {
			parseAddr, err := netip.ParsePrefix(addr.String())
			if err != nil {
				continue
			}

			if parseAddr.Addr().Is4() || parseAddr.Addr().IsMulticast() {
				continue
			}

			addr := fmt.Sprintf("[%s%%%s]:51234", parseAddr.Addr().String(), iface.Name)

			l, err := net.Listen("tcp", addr)
			if err != nil {
				log.Printf("listen %s error: %s", addr, err)
				continue
			}
			defer l.Close()

			go func(l net.Listener) {
				for {
					conn, err := l.Accept()
					if err != nil {
						log.Printf("accept error: %s", err)
						continue
					}
					// Handle connection in a new goroutine.
					go handleConnection(conn)
				}
			}(l)

			dialer := &net.Dialer{}
			addrIP, _, _ := net.ParseCIDR(parseAddr.Addr().String())
			dialer.LocalAddr = &net.TCPAddr{
				IP:   addrIP,
				Port: 0,
				Zone: iface.Name,
			}

			_, err = dialer.Dial("tcp", addr)
			if err != nil {
				log.Printf("dial %s error: %s", dialer, err)
			}
		}
	}
	return fmt.Sprintf("Hello, %s!", s)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// Implement your connection handling logic here
	log.Printf("Handling connection from %s", conn.RemoteAddr().String())
}
