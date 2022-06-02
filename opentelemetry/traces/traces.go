package traces

import (
	v1trace "awesomeProject/proto/otlp/collector/trace/v1"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (s Server) Export(ctx context.Context, request *v1trace.ExportTraceServiceRequest) (*v1trace.ExportTraceServiceResponse, error) {
	b, err := json.MarshalIndent(request.ResourceSpans, " ", "\t")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("printing request")
	fmt.Println(string(b))
	resp := v1trace.ExportTraceServiceResponse{}
	return &resp, nil
}
