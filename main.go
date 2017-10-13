package main

import (
	"context"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/docker/swarmkit/agent"
	"github.com/docker/swarmkit/agent/exec"
	"github.com/docker/swarmkit/api"
	"google.golang.org/grpc/credentials"
	"os"
)

// LocalExecutor implements agent/exec#Executor
type LocalExecutor struct{}

func (*LocalExecutor) Describe(ctx context.Context) (*api.NodeDescription, error) {
	return &api.NodeDescription{}, nil
}

// Configure uses the node object state to propagate node
// state to the underlying executor.
func (*LocalExecutor) Configure(ctx context.Context, node *api.Node) error {
	return nil
}

// Controller provides a controller for the given task.
func (*LocalExecutor) Controller(t *api.Task) (exec.Controller, error) {
	return nil, nil
}

// SetNetworkBootstrapKeys passes the symmetric keys from the
// manager to the executor.
func (*LocalExecutor) SetNetworkBootstrapKeys([]*api.EncryptionKey) error {
	return nil
}

func main() {
	fmt.Println("Starting natswarm")
	fmt.Println("Bootstrapping local agent")
	hostname, _ := os.Hostname()
	b := &x509.Certificate{}

	cert, err := x509.CreateCertificate(rand.Reader, nil, nil, nil, nil)
	if err != nil {
		fmt.Println(err)
	}
	tcert, err := tls.X509KeyPair(cert, keyPEMBlock)
	if err != nil {
		fmt.Println(err)
	}
	cfg := agent.Config{
		Hostname:    hostname,
		Credentials: credentials.NewServerTLSFromCert(&tcert),
	}
	a, err := agent.New(&cfg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}
