package server

import (
	handler "golang-uber-fx/adapter/http"
	repository "golang-uber-fx/adapter/mysql/clienteRepository"
	service "golang-uber-fx/core/usecase"
	my "golang-uber-fx/adapter/mysql"
	

	"gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"

	log "golang-uber-fx/util/log"

	"go.uber.org/fx"
)

func Start() {

	fx.New(
		fx.Options(

			fx.Provide(
				repository.NewRepository,
				service.NewService,
				mux.NewRouter,
				handler.NewServer,
			),
			fx.Invoke(func(job handler.IClientServer) {
				job.RegisterRoutes()
			},
			),
		),
	).Run()

}

// outra forma de chamar o fx e a forma que vc cria uma anotação com o nome da função e informa o retorno dela dentro de um fx.As

var ModuleRepository = fx.Module("repository", fx.Provide(
	fx.Annotate(
		repository.NewRepository,
		fx.As(new(repository.Irepository)),
	),
),
)

var ModuleService = fx.Module("service", fx.Provide(
	fx.Annotate(
		service.NewService,
		fx.As(new(service.Iservice)),
	),
),
)

var ModuleHandler = fx.Module("handler", fx.Provide(
	fx.Annotate(
		handler.NewServer,
		fx.As(new(handler.IClientServer)),
	),

))
var ModuleLog = fx.Module("log", fx.Provide(
	fx.Annotate(
		log.NewLogLevel,
		fx.As(new(log.ILogLevel)),
	),
))

func Start2() {
	fx.New(
		fx.Provide(
			mux.NewRouter,
			my.NewConnectDB,
			
		),
		ModuleRepository,
		ModuleLog,
		ModuleService,
		ModuleHandler,

		fx.Invoke(func(job handler.IClientServer) error {

			job.RegisterRoutes()
			// OU ...
			// r := mux.NewRouter()
			// r.HandleFunc("/cliente", job.Save).Methods("POST")
			// r.HandleFunc("/cliente", job.Del).Methods("DELETE")
			// r.HandleFunc("/cliente/{id}", job.Find).Methods("GET")
			// http.ListenAndServe(":8080", r)
			// fmt.Println("on na porta 8080")
			return nil
		},
		),
	).Run()

}
