package main

type configuration struct {
	NodeDaemon     string
	IpfsConnection string
}

var Conf = configuration{}

func init() {
	Conf.NodeDaemon = "http://localhost:46657"
}
