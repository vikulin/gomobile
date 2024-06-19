package mobile

import "net"
import "fmt"
import "github.com/wlynxg/anet"

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
            addrIP, _, _ := net.ParseCIDR(addr.String())
            // Ignore IPv4 addresses
            if addrIP.To4() != nil {
                continue
            }
            // Ignore non-link-local addresses
            if !addrIP.IsLinkLocalUnicast() {
                continue
            }
            _, err = net.Listen("tcp", fmt.Sprintf("[%s]:0", addrIP))
            if err != nil {
                s += fmt.Sprintf("Listen error: %s", err) + "\n"
            }
        }
    }
    return fmt.Sprintf("Hello, %s!", s)
}