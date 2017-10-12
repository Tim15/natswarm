package main

import (
  "github.com/docker/swarmkit/agent"
  "fmt"
  "os"
)

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
