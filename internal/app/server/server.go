package server

import (
	"database/sql"
	"net/http"
	"os"
	"patreon/internal/app"
	"patreon/internal/app/handlers"
	"patreon/internal/app/sessions/repository"
	"patreon/internal/app/sessions/sessions_manager"
	_ "patreon/internal/app/store/sqlstore"

	redis "github.com/gomodule/redigo/redis"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	_ "patreon/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

type Server struct {
	config  *Config
	handler app.Handler
	logger  *log.Logger
}

func New(config *Config, handler app.Handler) *Server {
	return &Server{
		config:  config,
		logger:  log.New(),
		handler: handler,
	}
}

func newDB(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// @title Patreon
// @version 1.0
// @description Server for Patreon application.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @x-extension-openapi {"example": "value on a json format"}

func Start(config *Config) error {
	level, err := log.ParseLevel(config.LogLevel)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	logger := log.New()
	logger.SetLevel(level)

	handler := handlers.NewMainHandler()
	handler.SetLogger(logger)

	router := mux.NewRouter()
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	handler.SetRouter(router)

	//handler.RegisterHandlers()

	/*db, err := newDB(config.DataBaseUrl)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	store := sqlstore.New(db)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}*/

	registerHandler := handlers.NewRegisterHandler()
	//registerHandler.SetStore(store)

	loginHandler := handlers.NewLoginHandler()
	//loginHandler.SetStore(store)
	//joinedHandlers := []app.Joinable{
	//	handlers.NewRegisterHandler(),
	//}
	sessionLog := log.New()
	sessionLog.SetLevel(log.FatalLevel)
	redisConn := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	redisRepository := repository.NewRedisRepository(redisConn, sessionLog)
	sessionManager := sessions_manager.NewSessionManager(redisRepository)
	loginHandler.SetSessionManager(sessionManager)
	handler.JoinHandlers([]app.Joinable{registerHandler, loginHandler})

	s := New(config, handler)
	s.logger.Info("starting server")
	return http.ListenAndServe(config.BindAddr, s.handler)
}
