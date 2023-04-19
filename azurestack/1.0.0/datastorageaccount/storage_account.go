// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datastorageaccount

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CustomDomain struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type CustomDomainAttributes struct {
	ref terra.Reference
}

func (cd CustomDomainAttributes) InternalRef() (terra.Reference, error) {
	return cd.ref, nil
}

func (cd CustomDomainAttributes) InternalWithRef(ref terra.Reference) CustomDomainAttributes {
	return CustomDomainAttributes{ref: ref}
}

func (cd CustomDomainAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cd.ref.InternalTokens()
}

func (cd CustomDomainAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cd.ref.Append("name"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type CustomDomainState struct {
	Name string `json:"name"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
