package enterprise

import (
	"github.com/aquasecurity/trivy/pkg/flag"
)

type enterpriseOptionKey struct{}

var (
	url = flag.Flag[string]{
		Name:       "enterprise-url",
		ConfigName: "enterprise-url",
		Usage:      "enterprise URL",
		Default:    "https://dummy.aquasec.com",
	}
	sastConfig = flag.Flag[string]{
		Name:       "sast-config",
		ConfigName: "sast-config",
		Usage:      "SAST config",
		Default:    "sast-config.yaml",
	}
	aquaKey = flag.Flag[string]{
		Name:       "aqua-key",
		ConfigName: "aqua-key",
		Usage:      "Aqua key",
		Default:    "",
	}
	aquaSecret = flag.Flag[string]{
		Name:       "aqua-secret",
		ConfigName: "aqua-secret",
		Usage:      "Aqua secret",
		Default:    "",
	}
)

type FlagGroup struct {
	URL        *flag.Flag[string]
	SASTConfig *flag.Flag[string]
	AquaKey    *flag.Flag[string]
	AquaSecret *flag.Flag[string]
}

type Options struct {
	URL        string
	SASTConfig string
	AquaKey    string
	AquaSecret string
}

func (e FlagGroup) Name() string {
	return "Enterprise"
}

func (e FlagGroup) Flags() []flag.Flagger {
	return []flag.Flagger{
		e.URL,
		e.SASTConfig,
		e.AquaKey,
		e.AquaSecret,
	}
}

// ToOptions sets the enterprise options to the options
func (e FlagGroup) ToOptions(opts *flag.Options) error {
	if opts.CustomOptions == nil {
		opts.CustomOptions = make(map[any]any)
	}
	opts.CustomOptions[enterpriseOptionKey{}] = Options{
		URL:        e.URL.Value(),
		SASTConfig: e.SASTConfig.Value(),
		AquaKey:    e.AquaKey.Value(),
		AquaSecret: e.AquaSecret.Value(),
	}
	return nil
}

type FlagExtension struct{}

func (e FlagExtension) Name() string {
	return "enterprise"
}

func (e FlagExtension) CustomFlagGroup(command string) flag.FlagGroup {
	return &FlagGroup{
		URL:        url.Clone(),
		SASTConfig: sastConfig.Clone(),
		AquaKey:    aquaKey.Clone(),
		AquaSecret: aquaSecret.Clone(),
	}
}

func ExtractOptions(opts flag.Options) Options {
	if opts.CustomOptions == nil {
		return Options{}
	}
	return opts.CustomOptions[enterpriseOptionKey{}].(Options)
}
