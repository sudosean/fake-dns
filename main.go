package main
import (
	"fmt"

	"github.com/miekg/dns"
)
// type Question struct {
//     Name   string `dns:"cdomain-name"` // "cdomain-name" specifies encoding (and may be compressed)
//     Qtype  uint16
//     Qclass uint16
// }

func main()  {
	fmt.Println("DNS Query and Response from localhost to  z.tiwaz.net")
	for index := 0; index < 100; index++ {
		Query()
	}
}

func Query()  {
	m1 := new(dns.Msg)
	m1.Id = dns.Id()
	m1.RecursionDesired = true
	m1.Question = make([]dns.Question, 1)
	m1.Question[0] = dns.Question{"z.tiwaz.", dns.TypeA, dns.ClassINET}	

	c := new(dns.Client)
	in, rtt, err := c.Exchange(m1, "127.0.0.1:53")

	if err != nil {
		panic(err)
	}
	fmt.Println("----- IN ----- ",in)
	fmt.Println("----- RTT ----- ",rtt)

}