package discovery

import (
	"context"
	"fmt"
	capi "github.com/hashicorp/consul/api"
	"log"
)

/**
 * @Author: bbz
 * @Email: 2632141215@qq.com
 * @File: consul
 * @Date:
 * @Desc: ...
 *
 */

type ConsulDiscovery struct {
	prefix  string
	Address string
	client  *capi.Client
}

func NewConsulDiscovery(address string, prefix string) *ConsulDiscovery {
	cfg := capi.DefaultConfig()
	cfg.Address = address
	client, err := capi.NewClient(cfg)
	if err != nil {
		log.Printf("Connect Consul happens error: %v", err)
	}
	return &ConsulDiscovery{
		prefix:  prefix,
		Address: address,
		client:  client,
	}

}
func (c *ConsulDiscovery) Register(ctx context.Context, service Service) error {

	reg := &capi.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", c.prefix, service.Name, service.Port),
		Name:    service.Name,
		Address: service.Host,
		Port:    service.Port,
		//Http检查
		//Check: &capi.AgentServiceCheck{
		//	Interval:                       "5s",
		//	Timeout:                        "5s",
		//	DeregisterCriticalServiceAfter: "10s",
		//	HTTP:                           fmt.Sprintf("http://%s:%d/health", service.Address, port),
		//},
	}
	if err := c.client.Agent().ServiceRegister(reg); err != nil {
		return err
	}
	log.Printf("register service %s success", service.Name)
	return nil
}

func (c *ConsulDiscovery) Deregister(ctx context.Context, name string) error {
	panic("implement me")
}

func (c *ConsulDiscovery) GetService(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("consul://%s/%s?wait=15s", c.Address, name), nil
}

var Consul = NewConsulDiscovery("192.168.40.129:8500", "test")
