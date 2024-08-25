package server

import (
	handler "golang-uber-fx/adapter/http"
	my "golang-uber-fx/adapter/mysql"
	repository "golang-uber-fx/adapter/mysql/repository"
	service "golang-uber-fx/core/usecase"
	"golang-uber-fx/routes"

	"gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"

	log "golang-uber-fx/util/log"

	"go.uber.org/fx"
)

func Start() {

	fx.New(
		fx.Options(

			fx.Provide(
				repository.NewClientRepository,
				service.NewService,
				mux.NewRouter,
				handler.NewServer,
			),
			fx.Invoke(func(job handler.IClientServer) {
				//job.RegisterRoutes()
			},
			),
		),
	).Run()

}

// outra forma de chamar o fx e a forma que vc cria uma anotação com o nome da função e informa o retorno dela dentro de um fx.As

var ModuleClientRepository = fx.Module("repository", fx.Provide(
	fx.Annotate(
		repository.NewClientRepository,
		fx.As(new(repository.IClientRepository)),
	),
),
)
var ModuleUserRepository = fx.Module("UserRepository", fx.Provide(
	fx.Annotate(
		repository.NewUserRepository,
		fx.As(new(repository.IUserRepository)),
	),
),
)

var ModuleClientService = fx.Module("service", fx.Provide(
	fx.Annotate(
		service.NewService,
		fx.As(new(service.IClientService)),
	),
),
)
var ModuleUserService = fx.Module("Userservice", fx.Provide(
	fx.Annotate(
		service.NewUserService,
		fx.As(new(service.IUserService)),
	),
),
)

var ModuleClientHandler = fx.Module("Clienthandler", fx.Provide(
	fx.Annotate(
		handler.NewServer,
		fx.As(new(handler.IClientServer)),
	),
))

var ModuleUserHandler = fx.Module("Userhandler", fx.Provide(
	fx.Annotate(
		handler.NewUserServer,
		fx.As(new(handler.IUserServer)),
	),
))
var ModuleRoutes = fx.Module("routes", fx.Provide(
	fx.Annotate(
		routes.NewRoutes,
		fx.As(new(routes.IRoutes)),
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
		ModuleClientRepository,
		ModuleUserRepository,
		ModuleLog,
		ModuleClientService,
		ModuleUserService,
		ModuleClientHandler,
		ModuleRoutes,
		ModuleUserHandler,

		fx.Invoke(func(job routes.IRoutes) error {

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
