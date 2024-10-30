package main

import (
	"github.com/go-chi/chi"
	"github.com/rakyll/statik/fs"
	"github.com/rs/zerolog/log"
	claylog "github.com/utrack/clay/v2/log"
	"github.com/utrack/clay/v2/transport/middlewares/mwgrpc"
	"github.com/utrack/clay/v2/transport/server"
	"net/http"

	// files of Swagger UI
	_ "github.com/utrack/clay/doc/example/static/statik"

	"proto-example/internal/app/api/example"
	"proto-example/internal/app/service"
)

func main() {
	//Swagger UI
	staticFS, err := fs.New()
	if err != nil {
		log.Fatal().Err(err).Msgf("%v", err)
	}
	hmux := chi.NewRouter()
	hmux.Mount("/", http.FileServer(staticFS))

	svc := service.NewService()

	impl := example.NewExampleAPI(svc)

	srv := server.NewServer(
		85,
		server.WithHTTPPort(80),
		server.WithHTTPMux(hmux),

		server.WithGRPCUnaryMiddlewares(mwgrpc.UnaryPanicHandler(claylog.Default)),
		server.WithGRPCStreamMiddlewares(mwgrpc.StreamPanicHandler(claylog.Default)),
	)

	if err := srv.Run(impl); err != nil {
		log.Fatal().Err(err).Msgf("%v", err)
	}
}
