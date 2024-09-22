package main

import (
	"flag"
	"life-telemetry/lib/utils"
	"life-telemetry/proto/services/eczema/v1/eczemav1connect"
	"life-telemetry/services/eczema"
	eczemadb "life-telemetry/services/eczema/db"
	"net/http"

	"connectrpc.com/connect"
)

func main() {
	verbose := flag.Bool("v", false, "enable verbose logging")
	flag.Parse()

	utils.InitTintedSlog(*verbose)

	mux := http.NewServeMux()

	db, err := utils.OpenSqlite(eczemadb.Schema, "eczema.db")
	if err != nil {
		utils.Fatal(err)
	}
	mux.Handle(eczemav1connect.NewEczemaServiceHandler(
		eczema.NewService(db),
		connect.WithInterceptors(utils.SlogInterceptor()),
	))

	utils.ListenGRPC(mux, 8000)
}
