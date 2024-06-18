package mobile

import "fmt"
import "github.com/wlynxg/anet"

func Greetings(name string) string {
    allifaces, err := anet.Interfaces()
    if err != nil {
        return fmt.Sprintf("Error: %s!", err)
    }
    s := ""
    for _, iface := range allifaces {
        s += iface.Name
        addrs, err := iface.Addrs()
        if err != nil {
            s += err + "\n"
        }
    }
	return fmt.Sprintf("Hello, %s!", s)
}