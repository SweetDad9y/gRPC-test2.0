package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	hello "grpc-test2.0/grpc"
	"grpc-test2.0/protogen"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных", err)
	}

	db.AutoMigrate(&hello.Table{})

	server := grpc.NewServer()
	serve := &hello.GRPCServer{DB: db}
	protogen.RegisterSayHelloServer(server, serve)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Serve(l); err != nil {
		log.Fatal(err)
	}
}
