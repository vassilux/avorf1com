package main

import (
	//"encoding/hex"
	"bytes"
	"testing"
)

func Test_Config(t *testing.T) {

	config, err := NewConfig()
	if err != nil {
		t.Fatalf(" Failed : %s", err)
	}

	if len(config.Address) == 0 {
		t.Fatal(" Failed : Cannot find the address")
	}

	if len(config.SiteId) == 0 {
		t.Fatal(" Failed : Cannot find the siteid")
	}

}

/*
func Test_Mongo(t *testing.T) {

	mongoWorker := NewMongoWorker()

	err := mongoWorker.Open("127.0.0.1")

	if err != nil {
		t.Fatalf(" Failed : %s", err)
	}

	var events []Event

	events, err = mongoWorker.Fetch()
	if err != nil {
		t.Fatalf(" Failed : %s", err)
	}

	for _, event := range events {

		t.Logf("Event : %s", event.String())
		err = mongoWorker.Delete(event.Id)

	}

	defer mongoWorker.Close()
}
*/
func Test_F1ComPacket(t *testing.T) {
	//header := []byte{0x20, 0x20, 0x20, 0x20, 0x20, 0x55, 0x47, 0x49, 0x20, 0x12, 0x09}
	header := []byte{0x46, 0x31, 0x43, 0x4F, 0x4D, 0x54, 0x53, 0x54, 0x20, 0x12, 0x09}
	data := []byte{0x51, 0x51, 0x51, 0x48, 0x48, 0x4d, 0x4d, 0x53, 0x53}
	f1comPacket := NewF1ComPacket(header, data)
	if f1comPacket.Header.Type != 0x12 {
		t.Fatalf("Uncorrect type of header")
	}

	if f1comPacket.Header.Network != 0x20 {
		t.Fatal("Failed to check the network value.")
	}
}

func Test_BuildAlarmPacket(t *testing.T) {
	eventType := 1
	eventCode := "EVTAPP"
	info := "test, bla"
	data := BuildAlarmPacket(eventType, eventCode, info)

	if !bytes.Contains(data, []byte(eventCode)) {
		t.Fatal("Failed to find eventCode into the returned bytes.")
	}
}
