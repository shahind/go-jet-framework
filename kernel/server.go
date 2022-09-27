package kernel

import (
	"github.com/shahind/go-jet-framework/register"
	"log"
	"net"
	"net/http"
	"strconv"
)

type ServerConf struct {
	Name    string
	Port    int
	SSL     bool
	SSLCert string
	SSLKey  string
	Key     string
}

var config *ServerConf

// RunServer this command
func RunServer(conf ServerConf, r []register.HTTPRouter) {
	config = &conf
	router := WebRouter(r)
	server := GetHttpServer(router, conf)

	if err := startServer(server, conf); err != nil {
		log.Fatal(err)
	}
}

// startServer will run the Go HTTP web server
func startServer(srv *http.Server, conf ServerConf) error {
	webListener, _ := net.Listen("tcp4", ":"+strconv.Itoa(conf.Port))

	if conf.SSL {
		if err := srv.ServeTLS(webListener, conf.SSLCert, conf.SSLKey); err != nil {
			return err
		}
	} else {
		if err := srv.Serve(webListener); err != nil {
			return err
		}
	}

	return nil
}
