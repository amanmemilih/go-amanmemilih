package controllers

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zinct/amanmemilih/config"
	"github.com/zinct/amanmemilih/internal/domain/entities"
	"github.com/zinct/amanmemilih/internal/domain/usecases"
	apperr "github.com/zinct/amanmemilih/internal/errors"
	"github.com/zinct/amanmemilih/internal/interface/delivery/http/v1/response"
	"github.com/zinct/amanmemilih/pkg/logger"
)

type Vote struct {
	CandidateID int `form:"candidat_id"`
	TotalVotes  int `form:"total_votes"`
}

type DocumentController struct {
	usecase usecases.DocumentUsecase
	config  *config.Config
	logger  *logger.Logger
}

func NewDocumentController(usecase usecases.DocumentUsecase, config *config.Config, logger *logger.Logger) *DocumentController {
	return &DocumentController{
		usecase: usecase,
		config:  config,
		logger:  logger,
	}
}

func (c *DocumentController) FindAll(ctx *gin.Context) {
	panic("not implemented")
}

func (c *DocumentController) Find(ctx *gin.Context) {
	panic("not implemented")
}

func (c *DocumentController) Verify(ctx *gin.Context) {
	panic("not implemented")
}

func (c *DocumentController) Create(ctx *gin.Context) {
	// Parse multipart form with 32MB max memory
	if err := ctx.Request.ParseMultipartForm(32 << 20); err != nil {
		ctx.JSON(400, gin.H{"error": "Failed to parse form data"})
		return
	}

	// Debug: Print all form values
	fmt.Println("All form values:")
	for key, values := range ctx.Request.PostForm {
		fmt.Printf("Key: %s, Values: %v\n", key, values)
	}

	// Get election type
	electionType := ctx.PostForm("election_type")

	// Validate election type
	if electionType == "" {
		apperr := apperr.NewValidationError("invalid request format", map[string][]string{
			"election_type": {"The election type field is required"},
		})
		response.JSONError(ctx, c.config, c.logger, apperr)
		return
	}

	// Validate election type value
	validElectionTypes := map[string]bool{
		"presidential": true,
		"legislative":  true,
		"regional":     true,
	}
	if !validElectionTypes[electionType] {
		apperr := apperr.NewValidationError("invalid request format", map[string][]string{
			"election_type": {"The election type must be one of: presidential, legislative, regional"},
		})
		response.JSONError(ctx, c.config, c.logger, apperr)
		return
	}

	// Handle file uploads
	form, err := ctx.MultipartForm()
	if err != nil {
		apperr := apperr.NewValidationError("invalid request format", map[string][]string{
			"documents": {"Failed to parse multipart form"},
		})
		response.JSONError(ctx, c.config, c.logger, apperr)
		return
	}

	// Get uploaded files
	files := form.File["documents[]"]
	if len(files) == 0 {
		apperr := apperr.NewValidationError("invalid request format", map[string][]string{
			"documents": {"The documents field is required"},
		})
		response.JSONError(ctx, c.config, c.logger, apperr)
		return
	}

	// Process uploaded files
	var documentContents []string
	var documentNames []string
	for _, file := range files {
		// Open the file
		src, err := file.Open()
		if err != nil {
			apperr := apperr.NewValidationError("invalid request format", map[string][]string{
				"documents": {fmt.Sprintf("Failed to open file %s", file.Filename)},
			})
			response.JSONError(ctx, c.config, c.logger, apperr)
			return
		}
		defer src.Close()

		// Read file content
		buf := new(bytes.Buffer)
		if _, err := io.Copy(buf, src); err != nil {
			apperr := apperr.NewValidationError("invalid request format", map[string][]string{
				"documents": {fmt.Sprintf("Failed to read file %s", file.Filename)},
			})
			response.JSONError(ctx, c.config, c.logger, apperr)
			return
		}

		// Store file content and name
		documentContents = append(documentContents, buf.String())
		documentNames = append(documentNames, file.Filename)
	}

	// Create a map to store votes
	votes := make(map[int]Vote)
	validationErrors := make(map[string][]string)

	// Process vote data from PostForm
	for key, values := range ctx.Request.PostForm {
		if len(values) == 0 {
			continue
		}

		// Check if the key matches the vote pattern
		if strings.HasPrefix(key, "vote[") && strings.Contains(key, "][") {
			// Extract index and field name
			parts := strings.Split(strings.Trim(key, "vote[]"), "][")
			if len(parts) != 2 {
				continue
			}

			index, err := strconv.Atoi(parts[0])
			if err != nil {
				continue
			}

			field := parts[1]
			value, err := strconv.Atoi(values[0])
			if err != nil {
				validationErrors[fmt.Sprintf("vote[%d][%s]", index, field)] = append(
					validationErrors[fmt.Sprintf("vote[%d][%s]", index, field)],
					"must be a number",
				)
				continue
			}

			// Create or update vote entry
			vote, exists := votes[index]
			if !exists {
				vote = Vote{}
			}

			switch field {
			case "candidat_id":
				vote.CandidateID = value
			case "total_votes":
				vote.TotalVotes = value
			}

			votes[index] = vote
		}
	}

	// Validate if vote array exists
	if len(votes) == 0 {
		apperr := apperr.NewValidationError("invalid request format", map[string][]string{
			"vote": {"The vote field is required"},
		})
		response.JSONError(ctx, c.config, c.logger, apperr)
		return
	}

	// Validate required fields for each vote
	for index, vote := range votes {
		if vote.CandidateID == 0 {
			validationErrors[fmt.Sprintf("vote[%d][candidat_id]", index)] = append(
				validationErrors[fmt.Sprintf("vote[%d][candidat_id]", index)],
				"is required",
			)
		}
		if vote.TotalVotes == 0 {
			validationErrors[fmt.Sprintf("vote[%d][total_votes]", index)] = append(
				validationErrors[fmt.Sprintf("vote[%d][total_votes]", index)],
				"is required",
			)
		}
	}

	// If there are validation errors, return them
	if len(validationErrors) > 0 {
		apperr := apperr.NewValidationError("invalid request format", validationErrors)
		response.JSONError(ctx, c.config, c.logger, apperr)
		return
	}

	// Convert map to slice
	var voteSlice []Vote
	for i := 0; i < len(votes); i++ {
		if vote, exists := votes[i]; exists {
			voteSlice = append(voteSlice, vote)
		}
	}

	// Map Vote slice to DocumentVote slice
	documentVotes := make([]entities.DocumentVote, len(voteSlice))
	for i, vote := range voteSlice {
		documentVotes[i] = entities.DocumentVote{
			CandidateId: vote.CandidateID,
			TotalVote:   vote.TotalVotes,
		}
	}

	userId := ctx.GetInt("user_id")

	// Process the data with usecase
	err = c.usecase.Create(ctx, userId, electionType, documentVotes, documentContents, documentNames)
	if err != nil {
		response.JSONError(ctx, c.config, c.logger, err)
		return
	}

	response.JSONSuccess(ctx, "Document created successfully", nil)
}
