package main

import (
  "github.com/docker/swarmkit/agent"
  "github.com/docker/swarmkit/api"
  "fmt"
  "os"
  "context"
)

// LocalExecutor implements agent/exec#Executor
type LocalExecutor struct

func (*LocalExecutor) Describe(ctx context.Context) (*api.NodeDescription, error) {
  return api.NodeDescription{}
}
// Configure uses the node object state to propagate node
// state to the underlying executor.
Configure(ctx context.Context, node *api.Node) error

// Controller provides a controller for the given task.
Controller(t *api.Task) (Controller, error)

// SetNetworkBootstrapKeys passes the symmetric keys from the
// manager to the executor.
SetNetworkBootstrapKeys([]*api.EncryptionKey) error

func main() {
  fmt.Println("Starting natswarm")
  fmt.Println("Bootstrapping local agent")
  hostname := os.Hostname()
  cfg := agent.Config{
    Hostname: hostname,
  }
  a, err := agent.Agent.New(cfg)
  if err != nil {
    fmt.Println(err)
  }
}
