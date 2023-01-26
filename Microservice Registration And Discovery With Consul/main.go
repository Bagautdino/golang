package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
	_ "github.com/hashicorp/consul/api/watch"
	"log"
	"net"
	"time"
)

const (
	ttl     = time.Second * 8
	checkID = "check_health"
)

type Service struct {
	consulClient *api.Client
}

func NewService() *Service {
	client, err := api.NewClient(&api.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return &Service{
		consulClient: client,
	}
}

func (s *Service) Start() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	s.registerService()
	go s.updateHealthCheck()
	s.acceptLoop(ln)

}

func (s *Service) acceptLoop(ln net.Listener) {
	for {
		_, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
	}
}
func (s *Service) updateHealthCheck() {
	ticker := time.NewTicker(time.Second * 5)
	for {
		err := s.consulClient.Agent().UpdateTTL(checkID, "online", api.HealthPassing)
		if err != nil {
			log.Fatal(err)
		}
		<-ticker.C
	}
}

func (s *Service) registerService() {
	check := &api.AgentServiceCheck{
		DeregisterCriticalServiceAfter: ttl.String(),
		TLSSkipVerify:                  true,
		TTL:                            ttl.String(),
		CheckID:                        "checkalive",
	}
	register := &api.AgentServiceRegistration{
		ID:      "login_service",
		Name:    "mycluster",
		Tags:    []string{"login"},
		Address: "127.0.0.1",
		Port:    8080,
		Check:   check,
	}

	query := map[string]any{
		"type":      "service",
		"service":   "mycluster",
		"passingly": true,
	}
	plan, err := watch.Parse(query)
	if err != nil {
		log.Fatal(err)
	}

	plan.HybridHandler = func(index watch.BlockingParamVal, result any) {
		switch msg := result.(type) {
		case []*api.ServiceEntry:
			for _, service := range msg {
				fmt.Println("new member joined", service.Service.Address)
			}
		}
		fmt.Println("update cluster", result)
	}
	go func() {
		if err := plan.RunWithConfig("", &api.Config{}); err != nil {
			log.Fatal(err)
		}
	}()

	if err := s.consulClient.Agent().ServiceRegister(register); err != nil {
		log.Fatal(err)
	}
}

func main() {
	s := NewService()
	s.Start()
}
