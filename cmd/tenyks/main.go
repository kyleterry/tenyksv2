package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"sync"

	"github.com/kyleterry/tenyks/pkg/adapter"
	"github.com/kyleterry/tenyks/pkg/adapter/irc"
	"github.com/kyleterry/tenyks/pkg/config"
	"github.com/kyleterry/tenyks/pkg/logger"
	"github.com/kyleterry/tenyks/pkg/service"
)

func main() {
	configPath := flag.String("config", "", "path to a configuration file")

	flag.Parse()

	cfg, err := config.NewConfigFromFile(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	loggerConf := logger.StandardLoggerConfig{
		Debug:         cfg.Logging.Debug,
		ShowTimestamp: true,
	}
	standardLogger := logger.NewStandardLogger("tenyks", loggerConf)

	adapterRegistry := adapter.NewRegistry()

	for _, sc := range cfg.Servers {
		at, ok := adapter.AdapterTypeMapping[sc.Kind]
		if !ok {
			log.Fatalf("no such adapter %s", sc.Kind)
		}

		switch at {
		case adapter.AdapterTypeIRC:
			ircConfig := sc.Config.(config.IRCServerConfig)

			c, err := irc.New(irc.Config{
				Name:     ircConfig.Name,
				Server:   ircConfig.ServerAddr,
				UseTLS:   ircConfig.UseTLS,
				Password: ircConfig.Password,
				User:     ircConfig.User,
				RealName: ircConfig.RealName,
				Nicks:    ircConfig.Nicks,
				Channels: ircConfig.Channels,
				Commands: ircConfig.Commands,
				Logger:   standardLogger,
			})

			if err != nil {
				log.Fatal(err)
			}

			if err := c.Dial(context.Background()); err != nil {
				log.Fatal(err)
			}

			adapterRegistry.RegisterAdapter(conn)
		}
	}

	ws := service.NewWebsocketServer(adapterRegistry)
	ws.RegisterHandler()

	http.ListenAndServe("0.0.0.0:9999", ws)

	wg := sync.WaitGroup{}

	wg.Add(1)

	wg.Wait()
}
