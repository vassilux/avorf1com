package main

import (
	"flag"
	"fmt"
	log "github.com/cihub/seelog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	VERSION = "1.0.1"
)

var (
	f1comSrv *F1ComServer
	config   *Config
	fake     = flag.Bool("fake", false, "generate fake alarm") // flag to generate fake alarm
	version  = flag.Bool("version", false, "show version")
)

func loadLogger() {
	logger, err := log.LoggerFromConfigAsFile("logger.xml")

	if err != nil {
		log.Error("Can not load the logger configuration file, Please check if the file logger.xml exists on current directory", err)
		os.Exit(1)
	} else {
		log.ReplaceLogger(logger)
		logger.Flush()
	}

}

func init() {
	//
	loadLogger()
}

func main() {
	//
	flag.Parse()
	//
	if *version {
		fmt.Printf("Version : %s\n", VERSION)
		fmt.Println("Get fun! Live well !")
		return
	}

	var err error
	config, err = NewConfig()
	checkErr(err)
	if err != nil {
		log.Errorf(" %s", err)
		return
	}

	if len(config.SiteId) != 8 {
		log.Errorf("Please check the configuration file, siteId parameter must be equal to 8 char.", err)
		return
	}
	addr := config.Address
	out := make(chan []byte)
	in := make(chan *F1ComPacket)

	laddr, err := net.ResolveTCPAddr("tcp", addr)
	if nil != err {
		log.Error(err)
		os.Exit(1)
	}
	listener, err := net.ListenTCP("tcp", laddr)
	if nil != err {
		log.Error(err)
		os.Exit(1)
	}
	log.Infof("Listening on : %s", listener.Addr().String())

	f1comSrv = NewF1ComServer(in, out, config)
	//
	if *fake {
		log.Info("Fake alarm activate.")
		duration := time.Duration(10) * time.Second
		ticker := time.NewTicker(duration)
		go func() {
			for {
				select {
				case <-ticker.C:
					fakeAlarm(f1comSrv.Clientlist)
				}
			}
		}()
	}

	//Run server
	f1comSrv.Run(listener)
	// Handle SIGINT and SIGTERM.
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Info(<-ch)

	// Stop the service gracefully.
	f1comSrv.Stop()
}
