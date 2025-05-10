package model

type Stats struct {
	stats *[]Stat
}

type Stat struct {
	Clouvider   ClouviderStats
	Smsactivate SmsActivateStats
	Asocks      AsocksStats
	Dataimpulse DataImpulseStats
	Iphoster    IpHosterStats
	Aeza        AezaStats
}

type ClouviderStats struct {
	Balance      string
	ServersCount int8
	Servers      interface{}
	Error        string
}
type SmsActivateStats struct {
	Balance string
	Error   string
}
type IpHosterStats struct {
	Balance string
	Error   string
}
type DataImpulseStats struct {
	TrafficLeft string
	Error       string
}
type AsocksStats struct {
	Balance string
	Error   string
}

type AezaStats struct {
	Balance string
	Error   string
}
