package icp

import (
	"context"

	"github.com/aviate-labs/agent-go"
	"github.com/aviate-labs/agent-go/principal"
	"github.com/pkg/errors"
	"github.com/zinct/amanmemilih/internal/domain/interfaces"
)

type Client struct {
	agent      *agent.Agent
	canisterId principal.Principal
}

func NewClient() (interfaces.BlockchainClient, error) {
	a, err := agent.New(agent.DefaultConfig)
	if err != nil {
		return nil, errors.Wrap(err, "Client.New")
	}

	return &Client{
		agent:      a,
		canisterId: principal.MustDecode("xumdk-uaaaa-aaaam-aejwq-cai"),
	}, nil
}

func (c *Client) CheckIfUserHasVoted(ctx context.Context, userId uint32) (bool, error) {
	var output bool
	if err := c.agent.Call(
		c.canisterId, "checkIfUserHasVoted",
		[]any{userId},
		[]any{&output},
	); err != nil {
		return false, errors.Wrap(err, "Client.CheckIfUserHasVoted")
	}
	return output, nil
}

func (c *Client) CreatePresidentialDocument(ctx context.Context, params interfaces.CreatePresidentialDocumentParams) error {
	if err := c.agent.Call(
		c.canisterId, "createPresidentialDocument",
		[]any{params},
		[]any{},
	); err != nil {
		return errors.Wrap(err, "Client.CreatePresidentialDocument")
	}
	return nil
}

func (c *Client) DeleteAllData(ctx context.Context) error {
	if err := c.agent.Call(
		c.canisterId, "deleteAllData",
		[]any{},
		[]any{},
	); err != nil {
		return errors.Wrap(err, "Client.DeleteAllData")
	}
	return nil
}

func (c *Client) GetPresidentialDocuments(ctx context.Context) ([]interface{}, error) {
	var output []interface{}
	if err := c.agent.Call(
		c.canisterId, "getPresidentialDocuments",
		[]any{},
		[]any{&output},
	); err != nil {
		return nil, errors.Wrap(err, "Client.GetPresidentialDocuments")
	}
	return output, nil
}

func (c *Client) GetPresidentialVotes(ctx context.Context) ([]interface{}, error) {
	var output []interface{}
	if err := c.agent.Call(
		c.canisterId, "getPresidentialVotes",
		[]any{},
		[]any{&output},
	); err != nil {
		return nil, errors.Wrap(err, "Client.GetPresidentialVotes")
	}
	return output, nil
}

func (c *Client) GetTotalVotes(ctx context.Context) ([]interface{}, error) {
	var output []interface{}
	if err := c.agent.Call(
		c.canisterId, "getTotalVotes",
		[]any{},
		[]any{&output},
	); err != nil {
		return nil, errors.Wrap(err, "Client.GetTotalVotes")
	}
	return output, nil
}
