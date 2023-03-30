// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package lambdafunctioneventinvokeconfig

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DestinationConfig struct {
	// OnFailure: optional
	OnFailure *OnFailure `hcl:"on_failure,block"`
	// OnSuccess: optional
	OnSuccess *OnSuccess `hcl:"on_success,block"`
}

type OnFailure struct {
	// Destination: string, required
	Destination terra.StringValue `hcl:"destination,attr" validate:"required"`
}

type OnSuccess struct {
	// Destination: string, required
	Destination terra.StringValue `hcl:"destination,attr" validate:"required"`
}

type DestinationConfigAttributes struct {
	ref terra.Reference
}

func (dc DestinationConfigAttributes) InternalRef() terra.Reference {
	return dc.ref
}

func (dc DestinationConfigAttributes) InternalWithRef(ref terra.Reference) DestinationConfigAttributes {
	return DestinationConfigAttributes{ref: ref}
}

func (dc DestinationConfigAttributes) InternalTokens() hclwrite.Tokens {
	return dc.ref.InternalTokens()
}

func (dc DestinationConfigAttributes) OnFailure() terra.ListValue[OnFailureAttributes] {
	return terra.ReferenceList[OnFailureAttributes](dc.ref.Append("on_failure"))
}

func (dc DestinationConfigAttributes) OnSuccess() terra.ListValue[OnSuccessAttributes] {
	return terra.ReferenceList[OnSuccessAttributes](dc.ref.Append("on_success"))
}

type OnFailureAttributes struct {
	ref terra.Reference
}

func (of OnFailureAttributes) InternalRef() terra.Reference {
	return of.ref
}

func (of OnFailureAttributes) InternalWithRef(ref terra.Reference) OnFailureAttributes {
	return OnFailureAttributes{ref: ref}
}

func (of OnFailureAttributes) InternalTokens() hclwrite.Tokens {
	return of.ref.InternalTokens()
}

func (of OnFailureAttributes) Destination() terra.StringValue {
	return terra.ReferenceString(of.ref.Append("destination"))
}

type OnSuccessAttributes struct {
	ref terra.Reference
}

func (os OnSuccessAttributes) InternalRef() terra.Reference {
	return os.ref
}

func (os OnSuccessAttributes) InternalWithRef(ref terra.Reference) OnSuccessAttributes {
	return OnSuccessAttributes{ref: ref}
}

func (os OnSuccessAttributes) InternalTokens() hclwrite.Tokens {
	return os.ref.InternalTokens()
}

func (os OnSuccessAttributes) Destination() terra.StringValue {
	return terra.ReferenceString(os.ref.Append("destination"))
}

type DestinationConfigState struct {
	OnFailure []OnFailureState `json:"on_failure"`
	OnSuccess []OnSuccessState `json:"on_success"`
}

type OnFailureState struct {
	Destination string `json:"destination"`
}

type OnSuccessState struct {
	Destination string `json:"destination"`
}
