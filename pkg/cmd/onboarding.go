// Temporary manual command for testing the onboarding endpoint.
// This will be replaced by Stainless-generated code once the OpenAPI spec is synced.

package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/MercuryTechnologies/mercury-cli/internal/auth"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var onboardingSubmit = cli.Command{
	Name:    "submit",
	Usage:   "Submit onboarding data to pre-fill a Mercury application",
	Suggest: true,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "file",
			Aliases:  []string{"f"},
			Usage:    "Path to JSON file with onboarding data (or use stdin)",
			Required: false,
		},
	},
	Action:          handleOnboardingSubmit,
	HideHelpCommand: true,
}

func handleOnboardingSubmit(ctx context.Context, cmd *cli.Command) error {
	// Read input from --file or stdin
	var input []byte
	var err error

	if filePath := cmd.String("file"); filePath != "" {
		input, err = os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("reading file: %w", err)
		}
	} else if isInputPiped() {
		input, err = io.ReadAll(os.Stdin)
		if err != nil {
			return fmt.Errorf("reading stdin: %w", err)
		}
	} else {
		return fmt.Errorf("provide onboarding data via --file or stdin")
	}

	// Validate it's valid JSON
	if !json.Valid(input) {
		return fmt.Errorf("input is not valid JSON")
	}

	// Resolve base URL
	baseURL := "https://api.mercury.com"
	if override := cmd.Root().String("base-url"); override != "" {
		baseURL = override
	}

	url := baseURL + "/api/v1/submit-onboarding-data"

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(input))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Auth is optional for onboarding — include if available but don't require it
	apiKey := cmd.Root().String("api-key")
	if apiKey == "" {
		environment := cmd.Root().String("environment")
		if environment == "" {
			environment = "production"
		}
		token, _ := auth.GetToken(environment)
		apiKey = token
	}
	if apiKey != "" {
		req.Header.Set("Api-Secret-Key", apiKey)
	}

	if cmd.Root().Bool("debug") {
		fmt.Fprintf(os.Stderr, "POST %s\n", url)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error (HTTP %d): %s", resp.StatusCode, string(body))
	}

	obj := gjson.ParseBytes(body)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "onboarding submit",
		Transform:      transform,
	})
}
