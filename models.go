package main

import (
	"encoding/json"
	"labix.org/v2/mgo/bson"
)

type Event struct {
	Id         bson.ObjectId `json:"id"              bson:"_id"`
	AppId      string        `json:"appid"           bson:"appid"`
	AsteriskId string        `json:"asteriskid"	     bson:"asteriskid"`
	Code       string        `json:"code"			 bson:"code"`
	Data       string        `json:"data"			 bson:"data"`
	Type       int           `json:"type"			 bson:"type"`
}

func (event *Event) String() string {
	datas, _ := json.MarshalIndent(event, "", "")
	return string(datas)
}

func (event *Event) Json() []byte {
	datas, _ := json.MarshalIndent(event, "", "")
	return datas
}

type EventToSend struct {
	Id   bson.ObjectId
	Data []byte
}

func NewEventToSend(id bson.ObjectId, data []byte) *EventToSend {
	event := &EventToSend{
		Id:   id,
		Data: make([]byte, len(data)),
	}
	copy(event.Data, data[:])

	return event
}
