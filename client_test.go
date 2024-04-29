package haproxy_test

import (
	"fmt"

	"github.com/ruansteve/go-haproxy"
)

func ExampleHAProxyClient_Stats() {
	client := &haproxy.HAProxyClient{
		Addr: "unix:///var/run/haproxy.sock",
	}
	stats, err := client.Stats()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, s := range stats {
		fmt.Printf("%s: %s\n", s.SvName, s.Status)
	}
}

func ExampleHAProxyClient_Info() {
	client := &haproxy.HAProxyClient{
		Addr: "unix:///var/run/haproxy.sock",
	}
	info, err := client.Info()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s version %s\n", info.Name, info.Version)
	// Output:
	//HAProxy version 2.8.9-1ppa1~jammy
}

func ExampleHAProxyClient_RunCommand() {
	client := &haproxy.HAProxyClient{
		Addr: "unix:///var/run/haproxy.sock",
	}
	result, err := client.RunCommand("show info")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result.String())
}
