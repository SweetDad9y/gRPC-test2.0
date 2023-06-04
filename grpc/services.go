package hello

import (
	"context"
	"log"
	"strings"

	"gorm.io/gorm"
	"grpc-test2.0/protogen"
)

type Table struct {
	gorm.Model
	Str1 string
	Str2 string
	Str3 string
}

type GRPCServer struct {
	protogen.UnimplementedSayHelloServer
	DB *gorm.DB
}

func (s *GRPCServer) Say(ctx context.Context, req *protogen.SayRequest) (*protogen.SayResponse, error) {
	str := req.GetRequest()
	strs := strings.Split(str, " ")

	newRecord := &Table{Str1: strs[0], Str2: strs[1], Str3: strs[2]}

	if err := s.DB.Create(newRecord).Error; err != nil {
		log.Printf("Не удалось записать данные в базу данных: %v", err)
		return nil, err
	}

	log.Printf("Данные записаны: %s", str)

	return &protogen.SayResponse{Response: str}, nil
}
