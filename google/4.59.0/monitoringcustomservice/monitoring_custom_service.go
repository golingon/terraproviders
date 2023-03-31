// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package monitoringcustomservice

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Telemetry struct {
	// ResourceName: string, optional
	ResourceName terra.StringValue `hcl:"resource_name,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type TelemetryAttributes struct {
	ref terra.Reference
}

func (t TelemetryAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TelemetryAttributes) InternalWithRef(ref terra.Reference) TelemetryAttributes {
	return TelemetryAttributes{ref: ref}
}

func (t TelemetryAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TelemetryAttributes) ResourceName() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("resource_name"))
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
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type TelemetryState struct {
	ResourceName string `json:"resource_name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
