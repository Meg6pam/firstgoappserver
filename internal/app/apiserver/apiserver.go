package apiserver

import (
	"github.com/fasthttp/router"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"http-rest-api/internal/libs/calculator"
	"strings"
)

var (
	requestChan chan fasthttp.RequestCtx
	donechan    chan string
)

type APIServer struct {
	serverconfig *Config
	logger       *logrus.Logger
	router       *router.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		serverconfig: config,
		logger:       logrus.New(),
		router:       router.New(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureLogger()
	s.configureRouter()
	for i := 0; i < 16; i++ { // threads
		go newWorker()
	}
	s.logger.Info("starting API server...")

	return fasthttp.ListenAndServe(s.serverconfig.BinAddr, s.router.Handler)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.serverconfig.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	requestChan = make(chan fasthttp.RequestCtx, 1000000)
	s.router.GET("/:id", saveRequest)
}

func saveRequest(ctx *fasthttp.RequestCtx) {
	go func() { requestChan <- *ctx }()
}

func handleRequest(ctx fasthttp.RequestCtx) {
	logrus.Info("entering in handleRequest")
	var input string
	if ctx.UserValue("id") != nil {
		input = ctx.UserValue("id").(string)
	} else {
		input = strings.TrimLeft("/", string(ctx.Request.URI().PathOriginal()))
	}
	in := strings.TrimLeft("/", string(ctx.Request.Header.RequestURI()))
	logrus.Info("in + " + in)
	inp := strings.TrimLeft("/", string(ctx.Request.URI().Path()))
	logrus.Info("inp " + inp)
	//fmt.Fprintf(ctx, "Hello, %s!\n", ctx.UserValue("id"))
	logrus.Info("got input" + input)
	output, _, err := calculator.GetAmicableNumberv2(input)
	logrus.Info("got output")
	if err != nil {
		output = "Error in handleRequest function"
	}
	//ctx.WriteString(output)
	ctx.Response.SetBodyString(output)
	logrus.Info("Sent response" + output)
	//	s.logger.Info("Got number " + input["id"] + ". Sent response " + output)
}

func newWorker() {
	for true {
		select {
		case ctx := <-requestChan:
			{
				handleRequest(ctx)
			}
		}
	}
}
