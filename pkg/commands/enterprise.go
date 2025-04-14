// While it should add a build tag to avoid importing enterprise package when not needed,
// it is not added now as this is a quick PoC.
// //go:build enterprise
package commands

import (
	"os"

	enterprise "github.com/aquasecurity/trivy-enterprise"
	"github.com/aquasecurity/trivy/pkg/extension"
)

func init() {
	extension.RegisterFlagExtension(enterprise.FlagExtension{})

	// Options are not parsed at this point, so we need to get them manually.
	aquaKey := os.Getenv("TRIVY_AQUA_KEY")
	aquaSecret := os.Getenv("TRIVY_AQUA_SECRET")

	if aquaKey != "" && aquaSecret != "" {
		extension.RegisterHook(enterprise.Hook{})
	}
}
