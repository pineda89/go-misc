package lib_gc_metrics_endpoint

import(
	"expvar"
	"net/http"
	"net"
)

var (
	Dummy dummy
)

type dummy struct{}

func init(){
	z:=expvar.NewString("METRICS")
	z.Set("OK")
}

func Start(_net,host,port string) error{
	sock, err := net.Listen(_net, host+":"+port)
	if err != nil {
		return err
	}
	go func() {

		http.Serve(sock, nil)
	}()
	return nil
}