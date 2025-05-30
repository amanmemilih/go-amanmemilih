package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/zinct/amanmemilih/config"
	"github.com/zinct/amanmemilih/internal/domain/entities"
	"github.com/zinct/amanmemilih/internal/domain/interfaces"
	"github.com/zinct/amanmemilih/internal/domain/usecases"
	"github.com/zinct/amanmemilih/internal/infrastructure/ipfs"
	"github.com/zinct/amanmemilih/pkg/logger"
)

type DocumentUsecase struct {
	client interfaces.BlockchainClient
	cfg    *config.Config
	log    *logger.Logger
	ipfs   ipfs.IPFS
}

func NewDocumentUsecase(client interfaces.BlockchainClient, cfg *config.Config, log *logger.Logger, ipfs ipfs.IPFS) usecases.DocumentUsecase {
	return &DocumentUsecase{client: client, cfg: cfg, log: log, ipfs: ipfs}
}

func (u *DocumentUsecase) FindAll(ctx context.Context, userId int) ([]interfaces.CheckDocumentResponse, error) {
	documents, err := u.client.CheckDocument(ctx, uint32(userId))
	if err != nil {
		return nil, err
	}

	return documents, nil
}

func (u *DocumentUsecase) Find(ctx context.Context, id int, electionType string) (*interfaces.PresidentialDocumentDetailResponse, error) {
	documents, err := u.client.GetDetailDocument(ctx, uint32(id), electionType)
	if err != nil {
		return nil, err
	}

	return documents, nil
}

func (u *DocumentUsecase) Verify(ctx context.Context, id int, electionType string) error {
	status, err := u.client.VerifyDocument(ctx, uint32(id), electionType)
	if err != nil {
		return err
	}

	if !status {
		return errors.Wrap(errors.New("document not verified"), err.Error())
	}

	return nil
}

func (u *DocumentUsecase) Create(ctx context.Context, userId int, electionType string, votes []entities.DocumentVote, documents []string, documentNames []string) error {
	// Upload documents to IPFS with election type as group
	ipfsHashes, err := u.ipfs.UploadMultiple(ctx, documents, "document-c1")
	if err != nil {
		return err
	}

	// Send data to ICP blockchain
	if electionType == "presidential" {
		// Convert votes to PresidentialVoteParams
		presidentialVotes := make([]interfaces.PresidentialVoteParams, len(votes))
		for i, vote := range votes {
			presidentialVotes[i] = interfaces.PresidentialVoteParams{
				CandidateId: uint32(vote.CandidateId),
				TotalVotes:  uint32(vote.TotalVote),
			}
		}

		// Create blockchain params
		params := interfaces.CreatePresidentialDocumentParams{
			CreatedAt:     time.Now().Format(time.RFC3339),
			DocumentC1Url: ipfsHashes,
			UserId:        uint32(userId),
			Vote:          presidentialVotes,
		}

		// Send to blockchain
		if err := u.client.CreatePresidentialDocument(ctx, params); err != nil {
			u.log.Error("Failed to create presidential document in blockchain", map[string]interface{}{
				"error":  err,
				"params": params,
			})
			return fmt.Errorf("failed to create presidential document in blockchain: %w", err)
		}
	} else {
		// Create blockchain params for non-presidential documents
		params := interfaces.CreateDocumentParams{
			CreatedAt:     time.Now().Format(time.RFC3339),
			DocumentC1Url: ipfsHashes,
			ElectionType:  electionType,
			UserID:        uint32(userId),
		}

		// Send to blockchain
		if err := u.client.CreateDocument(ctx, params); err != nil {
			u.log.Error("Failed to create document in blockchain", map[string]interface{}{
				"error":  err,
				"params": params,
			})
			return fmt.Errorf("failed to create document in blockchain: %w", err)
		}
	}

	// Log the IPFS hashes
	u.log.Info("Documents uploaded to IPFS", map[string]interface{}{
		"election_type": electionType,
		"ipfs_hashes":   ipfsHashes,
		"votes":         votes,
		"file_names":    documentNames,
	})

	return nil
}

func (u *DocumentUsecase) Summary(ctx context.Context) ([]interfaces.VotePercentage, error) {
	summary, err := u.client.GetTotalVotes(ctx)
	if err != nil {
		return nil, err
	}

	return summary, nil
}

func (u *DocumentUsecase) Dashboard(ctx context.Context, userId int) (*interfaces.DashboardResponse, error) {
	dashboard, err := u.client.GetDashboard(ctx, uint32(userId))
	if err != nil {
		return nil, err
	}

	return dashboard, nil
}
