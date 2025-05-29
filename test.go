package main

import (
	"fmt"
	"log"

	"github.com/aviate-labs/agent-go"
	"github.com/aviate-labs/agent-go/principal"
)

type (
	Account struct {
		Account string `ic:"account"`
	}

	Balance struct {
		E8S uint64 `ic:"e8s"`
	}
)

func main() {
	var balance Balance
	a, _ := agent.New(agent.DefaultConfig)
	if err := a.Call(
		principal.MustDecode("xumdk-uaaaa-aaaam-aejwq-cai"), "getTotalVotes",
		[]any{},
		[]any{&balance},
	); err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)
}
