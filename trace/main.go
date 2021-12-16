package main

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"log"
	"net/http"
)

func main() {
	tracer := opentracing.GlobalTracer()
	http.HandleFunc("/publish", func(w http.ResponseWriter, r *http.Request) {
		// Extract the context from the headers
		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		serverSpan := tracer.StartSpan("server", ext.RPCServerOption(spanCtx))
		defer serverSpan.Finish()
	})

	log.Fatal(http.ListenAndServe(":8082", nil))
}
