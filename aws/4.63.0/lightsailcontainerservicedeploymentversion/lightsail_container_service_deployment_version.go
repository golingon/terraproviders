// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package lightsailcontainerservicedeploymentversion

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Container struct {
	// Command: list of string, optional
	Command terra.ListValue[terra.StringValue] `hcl:"command,attr"`
	// ContainerName: string, required
	ContainerName terra.StringValue `hcl:"container_name,attr" validate:"required"`
	// Environment: map of string, optional
	Environment terra.MapValue[terra.StringValue] `hcl:"environment,attr"`
	// Image: string, required
	Image terra.StringValue `hcl:"image,attr" validate:"required"`
	// Ports: map of string, optional
	Ports terra.MapValue[terra.StringValue] `hcl:"ports,attr"`
}

type PublicEndpoint struct {
	// ContainerName: string, required
	ContainerName terra.StringValue `hcl:"container_name,attr" validate:"required"`
	// ContainerPort: number, required
	ContainerPort terra.NumberValue `hcl:"container_port,attr" validate:"required"`
	// HealthCheck: required
	HealthCheck *HealthCheck `hcl:"health_check,block" validate:"required"`
}

type HealthCheck struct {
	// HealthyThreshold: number, optional
	HealthyThreshold terra.NumberValue `hcl:"healthy_threshold,attr"`
	// IntervalSeconds: number, optional
	IntervalSeconds terra.NumberValue `hcl:"interval_seconds,attr"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// SuccessCodes: string, optional
	SuccessCodes terra.StringValue `hcl:"success_codes,attr"`
	// TimeoutSeconds: number, optional
	TimeoutSeconds terra.NumberValue `hcl:"timeout_seconds,attr"`
	// UnhealthyThreshold: number, optional
	UnhealthyThreshold terra.NumberValue `hcl:"unhealthy_threshold,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
}

type ContainerAttributes struct {
	ref terra.Reference
}

func (c ContainerAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ContainerAttributes) InternalWithRef(ref terra.Reference) ContainerAttributes {
	return ContainerAttributes{ref: ref}
}

func (c ContainerAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ContainerAttributes) Command() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("command"))
}

func (c ContainerAttributes) ContainerName() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("container_name"))
}

func (c ContainerAttributes) Environment() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](c.ref.Append("environment"))
}

func (c ContainerAttributes) Image() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("image"))
}

func (c ContainerAttributes) Ports() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](c.ref.Append("ports"))
}

type PublicEndpointAttributes struct {
	ref terra.Reference
}

func (pe PublicEndpointAttributes) InternalRef() (terra.Reference, error) {
	return pe.ref, nil
}

func (pe PublicEndpointAttributes) InternalWithRef(ref terra.Reference) PublicEndpointAttributes {
	return PublicEndpointAttributes{ref: ref}
}

func (pe PublicEndpointAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pe.ref.InternalTokens()
}

func (pe PublicEndpointAttributes) ContainerName() terra.StringValue {
	return terra.ReferenceAsString(pe.ref.Append("container_name"))
}

func (pe PublicEndpointAttributes) ContainerPort() terra.NumberValue {
	return terra.ReferenceAsNumber(pe.ref.Append("container_port"))
}

func (pe PublicEndpointAttributes) HealthCheck() terra.ListValue[HealthCheckAttributes] {
	return terra.ReferenceAsList[HealthCheckAttributes](pe.ref.Append("health_check"))
}

type HealthCheckAttributes struct {
	ref terra.Reference
}

func (hc HealthCheckAttributes) InternalRef() (terra.Reference, error) {
	return hc.ref, nil
}

func (hc HealthCheckAttributes) InternalWithRef(ref terra.Reference) HealthCheckAttributes {
	return HealthCheckAttributes{ref: ref}
}

func (hc HealthCheckAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hc.ref.InternalTokens()
}

func (hc HealthCheckAttributes) HealthyThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(hc.ref.Append("healthy_threshold"))
}

func (hc HealthCheckAttributes) IntervalSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(hc.ref.Append("interval_seconds"))
}

func (hc HealthCheckAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("path"))
}

func (hc HealthCheckAttributes) SuccessCodes() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("success_codes"))
}

func (hc HealthCheckAttributes) TimeoutSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(hc.ref.Append("timeout_seconds"))
}

func (hc HealthCheckAttributes) UnhealthyThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(hc.ref.Append("unhealthy_threshold"))
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

type ContainerState struct {
	Command       []string          `json:"command"`
	ContainerName string            `json:"container_name"`
	Environment   map[string]string `json:"environment"`
	Image         string            `json:"image"`
	Ports         map[string]string `json:"ports"`
}

type PublicEndpointState struct {
	ContainerName string             `json:"container_name"`
	ContainerPort float64            `json:"container_port"`
	HealthCheck   []HealthCheckState `json:"health_check"`
}

type HealthCheckState struct {
	HealthyThreshold   float64 `json:"healthy_threshold"`
	IntervalSeconds    float64 `json:"interval_seconds"`
	Path               string  `json:"path"`
	SuccessCodes       string  `json:"success_codes"`
	TimeoutSeconds     float64 `json:"timeout_seconds"`
	UnhealthyThreshold float64 `json:"unhealthy_threshold"`
}

type TimeoutsState struct {
	Create string `json:"create"`
}
