package main

import (
	"context"
	"log"
	"net"
	"net/http"
	monolith "tutours/soa/gateway/proto/monolith"
	ms_encounters "tutours/soa/gateway/proto/ms-encounters"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	monolith.MonolithServer
}

type EncountersServer struct {
	ms_encounters.EncountersServer
}

func enableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		h.ServeHTTP(w, r)
	})
}

func main() {
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
	monolith.RegisterMonolithServer(s, &server{})
	ms_encounters.RegisterEncountersServer(s, &EncountersServer{})
	log.Println("Serving gRPC on 0.0.0.0:9999")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	conn, err := grpc.NewClient(
		"0.0.0.0:5172",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	encounter_conn, err := grpc.NewClient(
		"0.0.0.0:8000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	client := monolith.NewMonolithClient(conn)
	encounterClient := ms_encounters.NewEncountersClient(encounter_conn)
	err = monolith.RegisterMonolithHandlerClient(context.Background(), gwmux, client)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	err = ms_encounters.RegisterEncountersHandlerClient(context.Background(), gwmux, encounterClient)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}
	gwServer.Handler = enableCors(gwmux)

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
