// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datalboutboundrule

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type FrontendIpConfiguration struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type FrontendIpConfigurationAttributes struct {
	ref terra.Reference
}

func (fic FrontendIpConfigurationAttributes) InternalRef() terra.Reference {
	return fic.ref
}

func (fic FrontendIpConfigurationAttributes) InternalWithRef(ref terra.Reference) FrontendIpConfigurationAttributes {
	return FrontendIpConfigurationAttributes{ref: ref}
}

func (fic FrontendIpConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return fic.ref.InternalTokens()
}

func (fic FrontendIpConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceString(fic.ref.Append("id"))
}

func (fic FrontendIpConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceString(fic.ref.Append("name"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

type FrontendIpConfigurationState struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
