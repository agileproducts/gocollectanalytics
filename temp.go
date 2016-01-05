package gocollectanalytics

import (
	"encoding/json"
	"log"
)

type Kollector struct {
	store Daytastore
}

type Daytastore interface {
	LogIt(interface{}) error
}

type Bogstore struct{}

func NewKollector(ds Daytastore) *Kollector {
	return &Kollector{
		store: ds,
	}
}

func (k *Kollector) record(datatype interface{}) string {
	err := k.store.LogIt(datatype)
	if err != nil {
		return "boo!"
	}
	return "ok"
}

func NewBogstore() *Bogstore {
	b := Bogstore{}
	return &b
}

func (b *Bogstore) LogIt(datatype interface{}) error {
	json, err := json.Marshal(datatype)
	if err != nil {
		return err
	}
	log.Printf("Logging: %+s", json)
	return nil
}
