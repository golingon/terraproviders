// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package external

import "github.com/golingon/lingon/pkg/terra"

var _ terra.Provider = (*Provider)(nil)

// Provider contains the configurations for provider.
type Provider struct{}

// LocalName returns the provider local name for [Provider].
func (p *Provider) LocalName() string {
	return "external"
}

// Source returns the provider source for [Provider].
func (p *Provider) Source() string {
	return "hashicorp/external"
}

// Version returns the provider version for [Provider].
func (p *Provider) Version() string {
	return "2.3.3"
}

// Configuration returns the provider configuration for [Provider].
func (p *Provider) Configuration() interface{} {
	return p
}
