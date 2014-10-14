package main

import (
	"container/list"
	"encoding/json"
	log "github.com/cihub/seelog"
	"io/ioutil"
	"os"
	"time"
)

type Config struct {
	Address   string
	SiteId    string
	MongoHost string
}

func NewConfig() (config *Config, err error) {
	var file []byte
	file, err = ioutil.ReadFile("config.json")

	if err != nil {
		return nil, err
	}

	config = new(Config)
	if err = json.Unmarshal(file, config); err != nil {
		return nil, err
	}

	return config, nil
}

/*func alarmFetcher(chAlarmSent chan int) {

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

}*/

//fake alarm generator
//used for the system testing
func fakeAlarm(ListChain *list.List) {
	eventType := 1
	eventCode := "FAKEEVP"
	info := "Simple fake alarme"
	data := BuildAlarmPacket(SITE_ID, eventType, eventCode, info)

	for e := ListChain.Front(); e != nil; e = e.Next() {
		c := e.Value.(F1ComClient)
		c.Out <- data
	}
}

func schedule(what func(), delay time.Duration) chan bool {
	stop := make(chan bool)
	go func() {
		for {
			what()
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
		}
	}()

	return stop
}

func checkErr(err error) {
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
