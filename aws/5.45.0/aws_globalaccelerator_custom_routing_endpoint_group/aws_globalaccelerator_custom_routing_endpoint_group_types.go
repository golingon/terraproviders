// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_globalaccelerator_custom_routing_endpoint_group

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DestinationConfiguration struct {
	// FromPort: number, required
	FromPort terra.NumberValue `hcl:"from_port,attr" validate:"required"`
	// Protocols: set of string, required
	Protocols terra.SetValue[terra.StringValue] `hcl:"protocols,attr" validate:"required"`
	// ToPort: number, required
	ToPort terra.NumberValue `hcl:"to_port,attr" validate:"required"`
}

type EndpointConfiguration struct {
	// EndpointId: string, optional
	EndpointId terra.StringValue `hcl:"endpoint_id,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type DestinationConfigurationAttributes struct {
	ref terra.Reference
}

func (dc DestinationConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return dc.ref, nil
}

func (dc DestinationConfigurationAttributes) InternalWithRef(ref terra.Reference) DestinationConfigurationAttributes {
	return DestinationConfigurationAttributes{ref: ref}
}

func (dc DestinationConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dc.ref.InternalTokens()
}

func (dc DestinationConfigurationAttributes) FromPort() terra.NumberValue {
	return terra.ReferenceAsNumber(dc.ref.Append("from_port"))
}

func (dc DestinationConfigurationAttributes) Protocols() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dc.ref.Append("protocols"))
}

func (dc DestinationConfigurationAttributes) ToPort() terra.NumberValue {
	return terra.ReferenceAsNumber(dc.ref.Append("to_port"))
}

type EndpointConfigurationAttributes struct {
	ref terra.Reference
}

func (ec EndpointConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ec.ref, nil
}

func (ec EndpointConfigurationAttributes) InternalWithRef(ref terra.Reference) EndpointConfigurationAttributes {
	return EndpointConfigurationAttributes{ref: ref}
}

func (ec EndpointConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ec.ref.InternalTokens()
}

func (ec EndpointConfigurationAttributes) EndpointId() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("endpoint_id"))
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

type DestinationConfigurationState struct {
	FromPort  float64  `json:"from_port"`
	Protocols []string `json:"protocols"`
	ToPort    float64  `json:"to_port"`
}

type EndpointConfigurationState struct {
	EndpointId string `json:"endpoint_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}
