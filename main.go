package main

// APC AP7900 Controller
//   Network Management Card AOS v3.5.8
//   Rack PDU APP v3.5.7

import (
	"flag"
	"fmt"
	expect "github.com/google/goexpect"
	"github.com/google/goterm/term"
	"github.com/ziutek/telnet"
	"log"
	"os"
	"time"
)

var timeout time.Duration

func main() {
	var (
		fc  = flag.String("c", "ap7900.yaml", "config file")
		fo  = flag.Int("o", -1, "OutLet")
		fs  = flag.Bool("s", false, "State")
		fl  = flag.Bool("l", false, "Load")
		fa1 = flag.Bool("on", false, "Immediate On")
		fa2 = flag.Bool("off", false, "Immediate Off")
		fa3 = flag.Bool("reboot", false, "Immediate Reboot")
		fa4 = flag.Bool("don", false, "Delayed On")
		fa5 = flag.Bool("doff", false, "Delayed Off")
		fa6 = flag.Bool("dreboot", false, "Delayed Reboot")
	)
	flag.Parse()

	err := LoadConfigs(*fc)
	if err != nil {
		log.Fatal(err)
	}
	timeout = time.Duration(Conf().Timeout) * time.Second

	exp, _, err := telnetSpawn(Conf().Address, timeout, expect.Verbose(Conf().Debug))
	if err != nil {
		fmt.Println(term.Redf("telnetSpawn(%q,%v) failed: %v", Conf().Address, timeout, err))
		os.Exit(1)
	}

	if *fl {
		fmt.Println(loadStatus(exp))
		return
	}

	if (*fo < 1) || (*fo > 8) {
		fmt.Println(term.Redf("-o (OutLet) must be 1 ~ 8"))
		flag.PrintDefaults()
		os.Exit(1)
	}

	switch {
	case *fs:
		fmt.Println(outputStatus(exp, *fo))
	case *fa1:
		outputManagement(exp, *fo, 1)
	case *fa2:
		outputManagement(exp, *fo, 2)
	case *fa3:
		outputManagement(exp, *fo, 3)
	case *fa4:
		outputManagement(exp, *fo, 4)
	case *fa5:
		outputManagement(exp, *fo, 5)
	case *fa6:
		outputManagement(exp, *fo, 6)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func telnetSpawn(addr string, timeout time.Duration, opts ...expect.Option) (expect.Expecter, <-chan error, error) {
	conn, err := telnet.Dial("tcp", addr)
	if err != nil {
		return nil, nil, err
	}

	resCh := make(chan error)

	return expect.SpawnGeneric(&expect.GenOptions{
		In:  conn,
		Out: conn,
		Wait: func() error {
			return <-resCh
		},
		Close: func() error {
			close(resCh)
			return conn.Close()
		},
		Check: func() bool { return true },
	}, timeout, opts...)
}
