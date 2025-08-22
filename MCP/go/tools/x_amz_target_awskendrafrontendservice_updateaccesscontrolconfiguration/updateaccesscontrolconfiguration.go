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

func UpdateaccesscontrolconfigurationHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.UpdateAccessControlConfigurationRequest
		
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
		url := fmt.Sprintf("%s/#X-Amz-Target=AWSKendraFrontendService.UpdateAccessControlConfiguration", cfg.BaseURL)
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
		var result models.UpdateAccessControlConfigurationResponse
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

func CreateUpdateaccesscontrolconfigurationTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_#X-Amz-Target=AWSKendraFrontendService_UpdateAccessControlConfiguration",
		mcp.WithDescription("<p>Updates an access control configuration for your documents in an index. This includes user and group access information for your documents. This is useful for user context filtering, where search results are filtered based on the user or their group access to documents.</p> <p>You can update an access control configuration you created without indexing all of your documents again. For example, your index contains top-secret company documents that only certain employees or users should access. You created an 'allow' access control configuration for one user who recently joined the 'top-secret' team, switching from a team with 'deny' access to top-secret documents. However, the user suddenly returns to their previous team and should no longer have access to top secret documents. You can update the access control configuration to re-configure access control for your documents as circumstances change.</p> <p>You call the <a href="https://docs.aws.amazon.com/kendra/latest/dg/API_BatchPutDocument.html">BatchPutDocument</a> API to apply the updated access control configuration, with the <code>AccessControlConfigurationId</code> included in the <a href="https://docs.aws.amazon.com/kendra/latest/dg/API_Document.html">Document</a> object. If you use an S3 bucket as a data source, you synchronize your data source to apply the <code>AccessControlConfigurationId</code> in the <code>.metadata.json</code> file. Amazon Kendra currently only supports access control configuration for S3 data sources and documents indexed using the <code>BatchPutDocument</code> API.</p>"),
		mcp.WithString("X-Amz-Target", mcp.Required(), mcp.Description("")),
		mcp.WithString("AccessControlList", mcp.Description("")),
		mcp.WithString("Description", mcp.Description("")),
		mcp.WithString("HierarchicalAccessControlList", mcp.Description("")),
		mcp.WithString("Id", mcp.Required(), mcp.Description("")),
		mcp.WithString("IndexId", mcp.Required(), mcp.Description("")),
		mcp.WithString("Name", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdateaccesscontrolconfigurationHandler(cfg),
	}
}
