// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package apprunnerobservabilityconfiguration

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type TraceConfiguration struct {
	// Vendor: string, optional
	Vendor terra.StringValue `hcl:"vendor,attr"`
}

type TraceConfigurationAttributes struct {
	ref terra.Reference
}

func (tc TraceConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return tc.ref, nil
}

func (tc TraceConfigurationAttributes) InternalWithRef(ref terra.Reference) TraceConfigurationAttributes {
	return TraceConfigurationAttributes{ref: ref}
}

func (tc TraceConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return tc.ref.InternalTokens()
}

func (tc TraceConfigurationAttributes) Vendor() terra.StringValue {
	return terra.ReferenceAsString(tc.ref.Append("vendor"))
}

type TraceConfigurationState struct {
	Vendor string `json:"vendor"`
}
