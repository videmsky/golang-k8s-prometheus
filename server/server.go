package server

import (
	"io"
	"net/http"
	"time"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// HTTPServer creates a http server and can be reached through the porvided port
type HTTPServer struct {
	port string
}

// NewHTTPServer initializes variables
func NewHTTPServer(port string) *HTTPServer {
	return &HTTPServer{port}
}

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "kubi_assignment_processed_ops_total",
		Help: "The total number of processed events",
	})
)

// Open creates the http server
func (s HTTPServer) Open() error {
	http.HandleFunc("/", home)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(s.port, nil)
	
	return nil
}

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Kubiya!!!")
	recordMetrics()
}

