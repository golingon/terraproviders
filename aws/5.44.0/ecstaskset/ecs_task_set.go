// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package ecstaskset

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type CapacityProviderStrategy struct {
	// Base: number, optional
	Base terra.NumberValue `hcl:"base,attr"`
	// CapacityProvider: string, required
	CapacityProvider terra.StringValue `hcl:"capacity_provider,attr" validate:"required"`
	// Weight: number, required
	Weight terra.NumberValue `hcl:"weight,attr" validate:"required"`
}

type LoadBalancer struct {
	// ContainerName: string, required
	ContainerName terra.StringValue `hcl:"container_name,attr" validate:"required"`
	// ContainerPort: number, optional
	ContainerPort terra.NumberValue `hcl:"container_port,attr"`
	// LoadBalancerName: string, optional
	LoadBalancerName terra.StringValue `hcl:"load_balancer_name,attr"`
	// TargetGroupArn: string, optional
	TargetGroupArn terra.StringValue `hcl:"target_group_arn,attr"`
}

type NetworkConfiguration struct {
	// AssignPublicIp: bool, optional
	AssignPublicIp terra.BoolValue `hcl:"assign_public_ip,attr"`
	// SecurityGroups: set of string, optional
	SecurityGroups terra.SetValue[terra.StringValue] `hcl:"security_groups,attr"`
	// Subnets: set of string, required
	Subnets terra.SetValue[terra.StringValue] `hcl:"subnets,attr" validate:"required"`
}

type Scale struct {
	// Unit: string, optional
	Unit terra.StringValue `hcl:"unit,attr"`
	// Value: number, optional
	Value terra.NumberValue `hcl:"value,attr"`
}

type ServiceRegistries struct {
	// ContainerName: string, optional
	ContainerName terra.StringValue `hcl:"container_name,attr"`
	// ContainerPort: number, optional
	ContainerPort terra.NumberValue `hcl:"container_port,attr"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// RegistryArn: string, required
	RegistryArn terra.StringValue `hcl:"registry_arn,attr" validate:"required"`
}

type CapacityProviderStrategyAttributes struct {
	ref terra.Reference
}

func (cps CapacityProviderStrategyAttributes) InternalRef() (terra.Reference, error) {
	return cps.ref, nil
}

func (cps CapacityProviderStrategyAttributes) InternalWithRef(ref terra.Reference) CapacityProviderStrategyAttributes {
	return CapacityProviderStrategyAttributes{ref: ref}
}

func (cps CapacityProviderStrategyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cps.ref.InternalTokens()
}

func (cps CapacityProviderStrategyAttributes) Base() terra.NumberValue {
	return terra.ReferenceAsNumber(cps.ref.Append("base"))
}

func (cps CapacityProviderStrategyAttributes) CapacityProvider() terra.StringValue {
	return terra.ReferenceAsString(cps.ref.Append("capacity_provider"))
}

func (cps CapacityProviderStrategyAttributes) Weight() terra.NumberValue {
	return terra.ReferenceAsNumber(cps.ref.Append("weight"))
}

type LoadBalancerAttributes struct {
	ref terra.Reference
}

func (lb LoadBalancerAttributes) InternalRef() (terra.Reference, error) {
	return lb.ref, nil
}

func (lb LoadBalancerAttributes) InternalWithRef(ref terra.Reference) LoadBalancerAttributes {
	return LoadBalancerAttributes{ref: ref}
}

func (lb LoadBalancerAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lb.ref.InternalTokens()
}

func (lb LoadBalancerAttributes) ContainerName() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("container_name"))
}

func (lb LoadBalancerAttributes) ContainerPort() terra.NumberValue {
	return terra.ReferenceAsNumber(lb.ref.Append("container_port"))
}

func (lb LoadBalancerAttributes) LoadBalancerName() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("load_balancer_name"))
}

func (lb LoadBalancerAttributes) TargetGroupArn() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("target_group_arn"))
}

type NetworkConfigurationAttributes struct {
	ref terra.Reference
}

func (nc NetworkConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return nc.ref, nil
}

func (nc NetworkConfigurationAttributes) InternalWithRef(ref terra.Reference) NetworkConfigurationAttributes {
	return NetworkConfigurationAttributes{ref: ref}
}

func (nc NetworkConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nc.ref.InternalTokens()
}

func (nc NetworkConfigurationAttributes) AssignPublicIp() terra.BoolValue {
	return terra.ReferenceAsBool(nc.ref.Append("assign_public_ip"))
}

func (nc NetworkConfigurationAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](nc.ref.Append("security_groups"))
}

func (nc NetworkConfigurationAttributes) Subnets() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](nc.ref.Append("subnets"))
}

type ScaleAttributes struct {
	ref terra.Reference
}

func (s ScaleAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s ScaleAttributes) InternalWithRef(ref terra.Reference) ScaleAttributes {
	return ScaleAttributes{ref: ref}
}

func (s ScaleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s ScaleAttributes) Unit() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("unit"))
}

func (s ScaleAttributes) Value() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("value"))
}

type ServiceRegistriesAttributes struct {
	ref terra.Reference
}

func (sr ServiceRegistriesAttributes) InternalRef() (terra.Reference, error) {
	return sr.ref, nil
}

func (sr ServiceRegistriesAttributes) InternalWithRef(ref terra.Reference) ServiceRegistriesAttributes {
	return ServiceRegistriesAttributes{ref: ref}
}

func (sr ServiceRegistriesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sr.ref.InternalTokens()
}

func (sr ServiceRegistriesAttributes) ContainerName() terra.StringValue {
	return terra.ReferenceAsString(sr.ref.Append("container_name"))
}

func (sr ServiceRegistriesAttributes) ContainerPort() terra.NumberValue {
	return terra.ReferenceAsNumber(sr.ref.Append("container_port"))
}

func (sr ServiceRegistriesAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(sr.ref.Append("port"))
}

func (sr ServiceRegistriesAttributes) RegistryArn() terra.StringValue {
	return terra.ReferenceAsString(sr.ref.Append("registry_arn"))
}

type CapacityProviderStrategyState struct {
	Base             float64 `json:"base"`
	CapacityProvider string  `json:"capacity_provider"`
	Weight           float64 `json:"weight"`
}

type LoadBalancerState struct {
	ContainerName    string  `json:"container_name"`
	ContainerPort    float64 `json:"container_port"`
	LoadBalancerName string  `json:"load_balancer_name"`
	TargetGroupArn   string  `json:"target_group_arn"`
}

type NetworkConfigurationState struct {
	AssignPublicIp bool     `json:"assign_public_ip"`
	SecurityGroups []string `json:"security_groups"`
	Subnets        []string `json:"subnets"`
}

type ScaleState struct {
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

type ServiceRegistriesState struct {
	ContainerName string  `json:"container_name"`
	ContainerPort float64 `json:"container_port"`
	Port          float64 `json:"port"`
	RegistryArn   string  `json:"registry_arn"`
}
