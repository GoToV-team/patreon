package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"os"
	server "patreon/internal/app/server"
	"patreon/internal/app/server/attachable_handler"
	"patreon/internal/app/store"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to config file")
}
func main() {
	flag.Parse()

	config := server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	level, err := log.ParseLevel(config.LogLevel)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	st := store.New(config.Store)
	err = st.Open()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	logger := log.New()
	logger.SetLevel(level)

	router := mux.NewRouter()
	handler := server.NewMainHandler(router, []attachable_handler.IAttachable{server.NewRegisterHandler(*st,
		[]attachable_handler.IAttachable{})})
	handler.SetLogger(logger)

	handler.Attach()

	s := server.New(config, handler)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
