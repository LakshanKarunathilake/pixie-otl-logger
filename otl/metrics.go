package metrics

import (
	v1 "awesomeProject/proto/otlp/collector/metrics/v1"
	"encoding/json"
	"fmt"
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) Export(ctx context.Context, request *v1.ExportMetricsServiceRequest) (*v1.ExportMetricsServiceResponse, error) {
	b, err := json.MarshalIndent(request.ResourceMetrics, " ", "\t")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("printing request")
	fmt.Println(string(b))
	resp := v1.ExportMetricsServiceResponse{}
	return &resp, nil
}
