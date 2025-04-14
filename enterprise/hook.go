package enterprise

import (
	"context"
	"fmt"

	"github.com/aquasecurity/trivy/pkg/flag"
	"github.com/aquasecurity/trivy/pkg/types"
)

type Hook struct{}

func (h Hook) Name() string {
	return "enterprise"
}

func (h Hook) PreRun(ctx context.Context, opts flag.Options) error {
	// Override options
	opts.Scanners = types.Scanners{types.MisconfigScanner}

	enterpriseOptions := ExtractOptions(opts)
	fmt.Println(enterpriseOptions.URL)
	fmt.Println(enterpriseOptions.AquaKey)
	fmt.Println(enterpriseOptions.AquaSecret)

	return nil
}

func (h Hook) PostRun(ctx context.Context, opts flag.Options) error {
	return nil
}

func (h Hook) PreReport(ctx context.Context, report *types.Report, opts flag.Options) error {
	return nil
}

func (h Hook) PostReport(ctx context.Context, report *types.Report, opts flag.Options) error {
	// Upload report to Aqua
	return nil
}
