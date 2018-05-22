package httpd

import (
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/Symantec/Dominator/lib/log"
	"github.com/Symantec/cloud-gate/broker"
	"github.com/Symantec/cloud-gate/broker/appconfiguration"
	"github.com/Symantec/cloud-gate/broker/configuration"
)

type HtmlWriter interface {
	WriteHtml(writer io.Writer)
}

type Server struct {
	brokers     map[string]broker.Broker
	appConfig   *appconfiguration.AppConfiguration
	config      *configuration.Configuration
	htmlWriters []HtmlWriter
	logger      log.DebugLogger
}

func StartServer(appConfig *appconfiguration.AppConfiguration, brokers map[string]broker.Broker,
	logger log.DebugLogger) (*Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", appConfig.Base.ListenPort))
	if err != nil {
		return nil, err
	}
	server := &Server{
		brokers:   brokers,
		logger:    logger,
		appConfig: appConfig,
	}
	http.HandleFunc("/", server.rootHandler)
	http.HandleFunc("/status", server.statusHandler)
	go http.Serve(listener, nil)
	return server, nil
}

func (s *Server) AddHtmlWriter(htmlWriter HtmlWriter) {
	s.htmlWriters = append(s.htmlWriters, htmlWriter)
}

func (s *Server) UpdateConfiguration(
	config *configuration.Configuration) error {
	s.config = config
	return nil
}
