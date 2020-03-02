package apiserver

import (
	"github.com/fasthttp/router"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"http-rest-api/internal/libs/calculator"
)

var (
	requestChan chan *fasthttp.RequestCtx
	donechan    chan string
)

type APIServer struct {
	serverconfig *Config
	logger       *logrus.Logger
	router       *router.Router
}

func New(config *Config) *APIServer {
	/*	for i := 0; i < 16; i++ { // threads
		go newWorker()
	}*/
	go newWorker()
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
	requestChan = make(chan *fasthttp.RequestCtx, 1000000)
	s.router.GET("/:id", saveRequest)
}

func saveRequest(ctx *fasthttp.RequestCtx) {
	go func() { requestChan <- ctx }()
}

func handleRequest(ctx *fasthttp.RequestCtx) {
	input := ctx.UserValue("id").(string)
	output, _, err := calculator.GetAmicableNumberv2(input)
	if err != nil {
		output = "Error in handleRequest function"
	}
	ctx.WriteString(output)
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
