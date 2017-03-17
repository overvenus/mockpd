package main

import (
	"flag"

	"strings"

	"github.com/overvenus/mockpd/cases"
	"github.com/overvenus/mockpd/server"
)

var eps = flag.String("endpoints", "http://127.0.0.1:42379,http://127.0.0.1:52379", "mock PD server endpoints")

func main() {
	es := strings.Split(*eps, ",")
	lcc := cases.NewLeaderChange(es)
	server.Serve(es, lcc)
	select {}
}
