package mobile

import "fmt"
import "net"

func Greetings(name string) string {
    allifaces, err := net.Interfaces()
    if err != nil {
        return fmt.Sprintf("Error: %s!", err)
    }
    s := ""
    for _, iface := range allifaces {
        s += iface.Name+"\n"
    }
	return fmt.Sprintf("Hello, %s!", s)
}