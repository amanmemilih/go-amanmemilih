package icp

import (
	"context"

	"github.com/aviate-labs/agent-go"
	"github.com/aviate-labs/agent-go/principal"
	"github.com/pkg/errors"
	"github.com/zinct/amanmemilih/internal/domain/interfaces"
)

// Types according to amanmemilih.did
type CheckDocumentResponse struct {
	ElectionType string `ic:"electionType"`
	ID           uint32 `ic:"id"`
	Name         string `ic:"name"`
	Status       uint8  `ic:"status"`
}

type VotePercentage struct {
	Name           string `ic:"name"`
	No             uint8  `ic:"no"`
	Image          string `ic:"image"`
	ID             uint32 `ic:"id"`
	VotePercentage string `ic:"vote_percentage"`
}

type DashboardResponse struct {
	NotUploaded uint8 `ic:"notUploaded"`
	Uploaded    uint8 `ic:"uploaded"`
	Verified    uint8 `ic:"verified"`
}

type Result struct {
	Err *string                             `ic:"Err,variant"`
	Ok  *PresidentialDocumentDetailResponse `ic:"Ok,variant"`
}

type ResultText struct {
	Err *string `ic:"Err,variant"`
	Ok  *string `ic:"Ok,variant"`
}

type PresidentialVoteParams struct {
	CandidateID uint32 `ic:"candidateId"`
	TotalVotes  uint32 `ic:"totalVotes"`
}

type PresidentialDocumentDetailVote struct {
	CandidateName string `ic:"candidateName"`
	CandidateNo   uint8  `ic:"candidateNo"`
	TotalVotes    uint32 `ic:"totalVotes"`
}

type PresidentialDocumentDetailResponse struct {
	CreatedAt    string                           `ic:"createdAt"`
	DocumentC1   []string                         `ic:"documentC1"`
	ElectionDate string                           `ic:"electionDate"`
	Status       uint8                            `ic:"status"`
	Votes        []PresidentialDocumentDetailVote `ic:"votes"`
}

type PresidentialDocument struct {
	CreatedAt  string   `ic:"createdAt"`
	DocumentC1 []string `ic:"documentC1"`
	ID         uint32   `ic:"id"`
	Status     uint8    `ic:"status"`
	UserID     uint32   `ic:"userId"`
}

type CreatePresidentialDocumentParams struct {
	CreatedAt     string                   `ic:"createdAt"`
	DocumentC1Url []string                 `ic:"documentC1Url"`
	UserID        uint32                   `ic:"userId"`
	Vote          []PresidentialVoteParams `ic:"vote"`
}

type CreateDocumentParams struct {
	CreatedAt     string   `ic:"createdAt"`
	DocumentC1Url []string `ic:"documentC1Url"`
	ElectionType  string   `ic:"electionType"`
	UserID        uint32   `ic:"userId"`
}

type GetUserResponse struct {
	ID     uint32 `ic:"id"`
	UserID uint32 `ic:"userId"`
}

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
		canisterId: principal.MustDecode("mj2qz-yiaaa-aaaad-aanoq-cai"),
	}, nil
}

func (c *Client) CheckDocument(ctx context.Context, userId uint32) ([]interfaces.CheckDocumentResponse, error) {
	var output []CheckDocumentResponse
	if err := c.agent.Call(
		c.canisterId, "checkDocument",
		[]any{userId},
		[]any{&output},
	); err != nil {
		return nil, errors.Wrap(err, "Client.CheckDocument")
	}

	result := make([]interfaces.CheckDocumentResponse, len(output))
	for i, doc := range output {
		result[i] = interfaces.CheckDocumentResponse{
			ElectionType: doc.ElectionType,
			ID:           doc.ID,
			Name:         doc.Name,
			Status:       doc.Status,
		}
	}
	return result, nil
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

func (c *Client) CreateDocument(ctx context.Context, params interfaces.CreateDocumentParams) error {
	icpParams := CreateDocumentParams{
		CreatedAt:     params.CreatedAt,
		DocumentC1Url: params.DocumentC1Url,
		ElectionType:  params.ElectionType,
		UserID:        params.UserID,
	}

	var result ResultText
	if err := c.agent.Call(
		c.canisterId, "createDocument",
		[]any{icpParams},
		[]any{&result},
	); err != nil {
		return errors.Wrap(err, "Client.CreateDocument")
	}
	if result.Err != nil {
		return errors.New(*result.Err)
	}
	return nil
}

func (c *Client) CreatePresidentialDocument(ctx context.Context, params interfaces.CreatePresidentialDocumentParams) error {
	icpParams := CreatePresidentialDocumentParams{
		CreatedAt:     params.CreatedAt,
		DocumentC1Url: params.DocumentC1Url,
		UserID:        params.UserId,
		Vote:          make([]PresidentialVoteParams, len(params.Vote)),
	}

	for i, v := range params.Vote {
		icpParams.Vote[i] = PresidentialVoteParams{
			CandidateID: v.CandidateId,
			TotalVotes:  v.TotalVotes,
		}
	}

	var result ResultText
	if err := c.agent.Call(
		c.canisterId, "createPresidentialDocument",
		[]any{icpParams},
		[]any{&result},
	); err != nil {
		return errors.Wrap(err, "Client.CreatePresidentialDocument")
	}
	if result.Err != nil {
		return errors.New(*result.Err)
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

func (c *Client) GetDetailDocument(ctx context.Context, documentId uint32, electionType string) (*interfaces.PresidentialDocumentDetailResponse, error) {
	var result Result
	if err := c.agent.Call(
		c.canisterId, "getDetailDocument",
		[]any{documentId, electionType},
		[]any{&result},
	); err != nil {
		return nil, errors.Wrap(err, "Client.GetDetailPresidentialDocument")
	}
	if result.Err != nil {
		return nil, errors.New(*result.Err)
	}

	if result.Ok == nil {
		return nil, nil
	}

	response := &interfaces.PresidentialDocumentDetailResponse{
		CreatedAt:    result.Ok.CreatedAt,
		DocumentC1:   result.Ok.DocumentC1,
		ElectionType: electionType,
		ElectionDate: result.Ok.ElectionDate,
		Status:       result.Ok.Status,
		Votes:        make([]interfaces.PresidentialDocumentDetailVote, len(result.Ok.Votes)),
	}

	for i, vote := range result.Ok.Votes {
		response.Votes[i] = interfaces.PresidentialDocumentDetailVote{
			CandidateName: vote.CandidateName,
			CandidateNo:   vote.CandidateNo,
			TotalVotes:    vote.TotalVotes,
		}
	}

	return response, nil
}

func (c *Client) GetPresidentialDocuments(ctx context.Context) ([]interfaces.PresidentialDocument, error) {
	var output []PresidentialDocument
	if err := c.agent.Call(
		c.canisterId, "getPresidentialDocuments",
		[]any{},
		[]any{&output},
	); err != nil {
		return nil, errors.Wrap(err, "Client.GetPresidentialDocuments")
	}

	result := make([]interfaces.PresidentialDocument, len(output))
	for i, doc := range output {
		result[i] = interfaces.PresidentialDocument{
			CreatedAt:  doc.CreatedAt,
			DocumentC1: doc.DocumentC1,
			ID:         doc.ID,
			Status:     doc.Status,
			UserID:     doc.UserID,
		}
	}
	return result, nil
}

func (c *Client) GetTotalVotes(ctx context.Context) ([]interfaces.VotePercentage, error) {
	var output []VotePercentage
	if err := c.agent.Call(
		c.canisterId, "getTotalVotes",
		[]any{},
		[]any{&output},
	); err != nil {
		return nil, errors.Wrap(err, "Client.GetTotalVotes")
	}

	result := make([]interfaces.VotePercentage, len(output))
	for i, vote := range output {
		result[i] = interfaces.VotePercentage{
			Name:           vote.Name,
			No:             vote.No,
			Image:          vote.Image,
			ID:             vote.ID,
			VotePercentage: vote.VotePercentage,
		}
	}
	return result, nil
}

func (c *Client) GetDashboard(ctx context.Context, userId uint32) (*interfaces.DashboardResponse, error) {
	var output DashboardResponse
	if err := c.agent.Call(
		c.canisterId, "getDashboard",
		[]any{userId},
		[]any{&output},
	); err != nil {
		return nil, errors.Wrap(err, "Client.GetDashboard")
	}

	return &interfaces.DashboardResponse{
		NotUploaded: output.NotUploaded,
		Uploaded:    output.Uploaded,
		Verified:    output.Verified,
	}, nil
}

func (c *Client) GetDocumentUser(ctx context.Context, electionType string) ([]interfaces.GetUserResponse, error) {
	var output []GetUserResponse
	if err := c.agent.Call(
		c.canisterId, "getDocumentUser",
		[]any{electionType},
		[]any{&output},
	); err != nil {
		return nil, errors.Wrap(err, "Client.GetDocumentUser")
	}

	result := make([]interfaces.GetUserResponse, len(output))
	for i, doc := range output {
		result[i] = interfaces.GetUserResponse{
			ID:     doc.ID,
			UserID: doc.UserID,
		}
	}
	return result, nil
}

func (c *Client) VerifyDocument(ctx context.Context, documentId uint32, electionType string) (bool, error) {
	var output bool
	if err := c.agent.Call(
		c.canisterId, "verifyDocument",
		[]any{documentId, electionType},
		[]any{&output},
	); err != nil {
		return false, errors.Wrap(err, "Client.VerifyDocument")
	}
	return output, nil
}
