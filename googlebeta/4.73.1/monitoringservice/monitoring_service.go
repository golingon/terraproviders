// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package monitoringservice

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Telemetry struct{}

type BasicService struct {
	// ServiceLabels: map of string, optional
	ServiceLabels terra.MapValue[terra.StringValue] `hcl:"service_labels,attr"`
	// ServiceType: string, optional
	ServiceType terra.StringValue `hcl:"service_type,attr"`
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

func (t TelemetryAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TelemetryAttributes) InternalWithRef(ref terra.Reference) TelemetryAttributes {
	return TelemetryAttributes{ref: ref}
}

func (t TelemetryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TelemetryAttributes) ResourceName() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("resource_name"))
}

type BasicServiceAttributes struct {
	ref terra.Reference
}

func (bs BasicServiceAttributes) InternalRef() (terra.Reference, error) {
	return bs.ref, nil
}

func (bs BasicServiceAttributes) InternalWithRef(ref terra.Reference) BasicServiceAttributes {
	return BasicServiceAttributes{ref: ref}
}

func (bs BasicServiceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bs.ref.InternalTokens()
}

func (bs BasicServiceAttributes) ServiceLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bs.ref.Append("service_labels"))
}

func (bs BasicServiceAttributes) ServiceType() terra.StringValue {
	return terra.ReferenceAsString(bs.ref.Append("service_type"))
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

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type TelemetryState struct {
	ResourceName string `json:"resource_name"`
}

type BasicServiceState struct {
	ServiceLabels map[string]string `json:"service_labels"`
	ServiceType   string            `json:"service_type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
