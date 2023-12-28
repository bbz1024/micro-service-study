package discovery

import (
	"context"
	"fmt"
)

type Service struct {
	Name string
	Host string
	Port int

	Weight int
}

func (s Service) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

type Discovery interface {
	// Register register service
	Register(ctx context.Context, service Service) error
	// Deregister deregister service
	Deregister(ctx context.Context, name string) error
	// GetService get service
	GetService(ctx context.Context, name string) (string, error)
}
