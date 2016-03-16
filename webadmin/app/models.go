package app

import "github.com/zubairhamed/betwixt"

type StatsModel struct {
	ClientsCount int
	MemUsage     string
	Requests     int
	Errors       int
}

type ClientModel struct {
	Endpoint         string
	RegistrationID   string
	RegistrationDate string
	LastUpdate       string
	Objects          map[string]ObjectModel
}

type ResourceModel struct {

}

type ContentValueModel struct {
	Id    uint16
	Value interface{}
}

type ExecuteResponseModel struct {
	Content []*ContentValueModel
}

type ObjectModel struct {
	Instances  []int
	Definition betwixt.ObjectDefinition
}
