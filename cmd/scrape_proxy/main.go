package main

import (
	"github.com/levelitta/scrape_proxy/pkg/api"
	"github.com/levelitta/scrape_proxy/pkg/api/http_client"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

var v *viper.Viper

func main() {
	InitConfig()

	addr := v.GetString("addr")

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Panic(err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(api.NewAuthInterceptor(v.GetString("auth_token")).Unary()),
	)
	api.RegisterScrapeProxyServer(s, api.NewImplementation(http_client.NewClient()))

	log.Println("Run app...")
	log.Printf("port %s/grpc", addr)

	if err := s.Serve(listener); err != nil {
		log.Panic(err)
	}
}

func InitConfig() {
	v = viper.New()
	v.SetEnvPrefix("scrape_proxy")

	input := []string{
		"addr",
		"auth_token",
	}

	for _, in := range input {
		err := v.BindEnv(in)
		if err != nil {
			log.Panicf("bindEnvVariables: %s", err)
		}
	}

}
