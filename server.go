package main

import (
	"awesomeProject/opentelemetry/traces"
	metrics "awesomeProject/otl"
	v1 "awesomeProject/proto/otlp/collector/metrics/v1"
	v1trace "awesomeProject/proto/otlp/collector/trace/v1"
	"fmt"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
	"log"
	"net"
)

func main() {

	fmt.Println("Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := metrics.Server{}
	s2 := traces.Server{}

	grpcServer := grpc.NewServer()

	v1.RegisterMetricsServiceServer(grpcServer, &s)
	v1trace.RegisterTraceServiceServer(grpcServer, &s2)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
