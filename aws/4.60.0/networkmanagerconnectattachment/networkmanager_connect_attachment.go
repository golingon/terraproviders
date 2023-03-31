// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package networkmanagerconnectattachment

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Options struct {
	// Protocol: string, optional
	Protocol terra.StringValue `hcl:"protocol,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type OptionsAttributes struct {
	ref terra.Reference
}

func (o OptionsAttributes) InternalRef() terra.Reference {
	return o.ref
}

func (o OptionsAttributes) InternalWithRef(ref terra.Reference) OptionsAttributes {
	return OptionsAttributes{ref: ref}
}

func (o OptionsAttributes) InternalTokens() hclwrite.Tokens {
	return o.ref.InternalTokens()
}

func (o OptionsAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("protocol"))
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

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

type OptionsState struct {
	Protocol string `json:"protocol"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}
