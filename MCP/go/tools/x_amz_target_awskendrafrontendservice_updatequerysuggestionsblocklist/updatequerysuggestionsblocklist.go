package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/awskendrafrontendservice/mcp-server/config"
	"github.com/awskendrafrontendservice/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func UpdatequerysuggestionsblocklistHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.UpdateQuerySuggestionsBlockListRequest
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/#X-Amz-Target=AWSKendraFrontendService.UpdateQuerySuggestionsBlockList", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["X-Amz-Target"]; ok {
			req.Header.Set("X-Amz-Target", fmt.Sprintf("%v", val))
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateUpdatequerysuggestionsblocklistTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_#X-Amz-Target=AWSKendraFrontendService_UpdateQuerySuggestionsBlockList",
		mcp.WithDescription("<p>Updates a block list used for query suggestions for an index.</p> <p>Updates to a block list might not take effect right away. Amazon Kendra needs to refresh the entire suggestions list to apply any updates to the block list. Other changes not related to the block list apply immediately.</p> <p>If a block list is updating, then you need to wait for the first update to finish before submitting another update.</p> <p>Amazon Kendra supports partial updates, so you only need to provide the fields you want to update.</p> <p> <code>UpdateQuerySuggestionsBlockList</code> is currently not supported in the Amazon Web Services GovCloud (US-West) region.</p>"),
		mcp.WithString("X-Amz-Target", mcp.Required(), mcp.Description("")),
		mcp.WithString("Id", mcp.Required(), mcp.Description("")),
		mcp.WithString("IndexId", mcp.Required(), mcp.Description("")),
		mcp.WithString("Name", mcp.Description("")),
		mcp.WithString("RoleArn", mcp.Description("")),
		mcp.WithString("SourceS3Path", mcp.Description("")),
		mcp.WithString("Description", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdatequerysuggestionsblocklistHandler(cfg),
	}
}
