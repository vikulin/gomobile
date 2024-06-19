package mobile

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
    }
    return fmt.Sprintf("Hello, %s!", s)
}