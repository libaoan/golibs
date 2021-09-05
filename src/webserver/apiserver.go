package webserver

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
)

var (
	port      = flag.Uint("port", 8080, "The port to listen on.  Default 8080.")
	address   = flag.String("address", "127.0.0.1", "The address on the local server to listen to. Default 127.0.0.1")
	certFile  = flag.String("cert_file", "ssl/cert.pem", "Cert File, default ssl/cert.pem")
	keyFile  = flag.String("key_file", "ssl/key.pem", "Cert File, default ssl/key.pem")
)

func RunServer(){
	log.Println("Start ApiServer ...")
	flag.Parse()

	err := http.ListenAndServeTLS(
		fmt.Sprintf("%s:%d", *address, *port),
		*certFile,
		*keyFile,
		NewApiServer())
	if err != nil {
		log.Fatal(err)
	}
}


// ApiServer object.
// 'storage' cache the server data.
type ApiServer struct {
	storage map[string]bool
	lock sync.Mutex
}

func NewApiServer() *ApiServer {
	return &ApiServer{
		storage: make(map[string]bool, 0),
	}
}

// HTTP Handler interface
func (server *ApiServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s", req.Method, req.RequestURI)
	urls, err := url.ParseRequestURI(req.RequestURI)
	if err != nil {
		server.error(err, w)
		return
	}

	if urls.Path != "/strings" {
		server.notFound(req, w)
		return
	}else {
		server.handleREST(urls, req, w)
	}
}

func (server *ApiServer) handleREST(url *url.URL, req *http.Request, w http.ResponseWriter) {
	switch req.Method {
	case "POST", "PUT":
		body, err := server.readBody(req)
		if err != nil {
			server.error(err, w)
			return
		}
		data := []string{}
		err = json.Unmarshal([]byte(body), &data)
		if err != nil {
			server.error(err, w)
			return
		}
		res := server.sevice(data)
		server.write(200, res, w)
		return
	default:
		server.notPermit(req, w)
	}
}

func (server *ApiServer) readBody(req *http.Request) (string, error) {
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	return string(body), err
}

func (server *ApiServer) write(statusCode int, object interface{}, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	output, err := json.Marshal(object)
	if err != nil {
		server.error(err, w)
		return
	}
	_, err = w.Write(output)
	if err != nil {
		log.Println(err)
	}
}

func (server *ApiServer) error(err error, w http.ResponseWriter) {
	w.WriteHeader(500)
	_, err = fmt.Fprintf(w, "Internal Error: %#v", err)
	if err != nil {
		log.Println(err)
	}
}

func (server *ApiServer) notFound(req *http.Request, w http.ResponseWriter) {
	w.WriteHeader(404)
	_, err := fmt.Fprintf(w, "Not Found: %#v", req)
	if err != nil {
		log.Println(err)
	}
}

func (server *ApiServer) notPermit(req *http.Request, w http.ResponseWriter) {
	w.WriteHeader(405)
	_, err := fmt.Fprintf(w, "Not Permit: %#v", req)
	if err != nil {
		log.Println(err)
	}
}

func (server *ApiServer) sevice(data []string) []bool {
	res := []bool{}
	server.lock.Lock()
	defer server.lock.Unlock()
	for _, item := range data {
		if _, ok := server.storage[item]; ok {
			res = append(res, true)
		}else{
			res = append(res, false)
			server.storage[item] = true
		}
	}
	return res
}