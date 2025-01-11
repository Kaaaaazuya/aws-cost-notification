package cost

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
)

type explorerClient interface {
	GetCostAndUsage(ctx context.Context, params *costexplorer.GetCostAndUsageInput, optFns ...func(*costexplorer.Options)) (*costexplorer.GetCostAndUsageOutput, error)
}

// NewExplorerClient returns a new cost explorer client
func NewExplorerClient(region string) (*costexplorer.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		return nil, err
	}

	client := costexplorer.NewFromConfig(cfg)
	return client, nil
}

// GetCostAndUsage returns the cost for the current month
func GetCostAndUsage(ctx context.Context, client explorerClient, start *string, end *string) (*costexplorer.GetCostAndUsageOutput, error) {
	return client.GetCostAndUsage(ctx, createGetCostAndUsageInput(start, end))
}

// createGetCostAndUsageInput returns the input for the cost explorer
func createGetCostAndUsageInput(start *string, end *string) *costexplorer.GetCostAndUsageInput {
	return &costexplorer.GetCostAndUsageInput{
		Granularity: "MONTHLY",
		Metrics: []string{
			"UnblendedCost",
		},
		TimePeriod: &types.DateInterval{
			Start: start,
			End:   end,
		},
	}
}
