// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package ecsservice

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Alarms struct {
	// AlarmNames: set of string, required
	AlarmNames terra.SetValue[terra.StringValue] `hcl:"alarm_names,attr" validate:"required"`
	// Enable: bool, required
	Enable terra.BoolValue `hcl:"enable,attr" validate:"required"`
	// Rollback: bool, required
	Rollback terra.BoolValue `hcl:"rollback,attr" validate:"required"`
}

type CapacityProviderStrategy struct {
	// Base: number, optional
	Base terra.NumberValue `hcl:"base,attr"`
	// CapacityProvider: string, required
	CapacityProvider terra.StringValue `hcl:"capacity_provider,attr" validate:"required"`
	// Weight: number, optional
	Weight terra.NumberValue `hcl:"weight,attr"`
}

type DeploymentCircuitBreaker struct {
	// Enable: bool, required
	Enable terra.BoolValue `hcl:"enable,attr" validate:"required"`
	// Rollback: bool, required
	Rollback terra.BoolValue `hcl:"rollback,attr" validate:"required"`
}

type DeploymentController struct {
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
}

type LoadBalancer struct {
	// ContainerName: string, required
	ContainerName terra.StringValue `hcl:"container_name,attr" validate:"required"`
	// ContainerPort: number, required
	ContainerPort terra.NumberValue `hcl:"container_port,attr" validate:"required"`
	// ElbName: string, optional
	ElbName terra.StringValue `hcl:"elb_name,attr"`
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

type OrderedPlacementStrategy struct {
	// Field: string, optional
	Field terra.StringValue `hcl:"field,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type PlacementConstraints struct {
	// Expression: string, optional
	Expression terra.StringValue `hcl:"expression,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type ServiceConnectConfiguration struct {
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
	// LogConfiguration: optional
	LogConfiguration *LogConfiguration `hcl:"log_configuration,block"`
	// Service: min=0
	Service []Service `hcl:"service,block" validate:"min=0"`
}

type LogConfiguration struct {
	// LogDriver: string, required
	LogDriver terra.StringValue `hcl:"log_driver,attr" validate:"required"`
	// Options: map of string, optional
	Options terra.MapValue[terra.StringValue] `hcl:"options,attr"`
	// SecretOption: min=0
	SecretOption []SecretOption `hcl:"secret_option,block" validate:"min=0"`
}

type SecretOption struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ValueFrom: string, required
	ValueFrom terra.StringValue `hcl:"value_from,attr" validate:"required"`
}

type Service struct {
	// DiscoveryName: string, optional
	DiscoveryName terra.StringValue `hcl:"discovery_name,attr"`
	// IngressPortOverride: number, optional
	IngressPortOverride terra.NumberValue `hcl:"ingress_port_override,attr"`
	// PortName: string, required
	PortName terra.StringValue `hcl:"port_name,attr" validate:"required"`
	// ClientAlias: optional
	ClientAlias *ClientAlias `hcl:"client_alias,block"`
}

type ClientAlias struct {
	// DnsName: string, optional
	DnsName terra.StringValue `hcl:"dns_name,attr"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
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

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AlarmsAttributes struct {
	ref terra.Reference
}

func (a AlarmsAttributes) InternalRef() terra.Reference {
	return a.ref
}

func (a AlarmsAttributes) InternalWithRef(ref terra.Reference) AlarmsAttributes {
	return AlarmsAttributes{ref: ref}
}

func (a AlarmsAttributes) InternalTokens() hclwrite.Tokens {
	return a.ref.InternalTokens()
}

func (a AlarmsAttributes) AlarmNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](a.ref.Append("alarm_names"))
}

func (a AlarmsAttributes) Enable() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("enable"))
}

func (a AlarmsAttributes) Rollback() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("rollback"))
}

type CapacityProviderStrategyAttributes struct {
	ref terra.Reference
}

func (cps CapacityProviderStrategyAttributes) InternalRef() terra.Reference {
	return cps.ref
}

func (cps CapacityProviderStrategyAttributes) InternalWithRef(ref terra.Reference) CapacityProviderStrategyAttributes {
	return CapacityProviderStrategyAttributes{ref: ref}
}

func (cps CapacityProviderStrategyAttributes) InternalTokens() hclwrite.Tokens {
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

type DeploymentCircuitBreakerAttributes struct {
	ref terra.Reference
}

func (dcb DeploymentCircuitBreakerAttributes) InternalRef() terra.Reference {
	return dcb.ref
}

func (dcb DeploymentCircuitBreakerAttributes) InternalWithRef(ref terra.Reference) DeploymentCircuitBreakerAttributes {
	return DeploymentCircuitBreakerAttributes{ref: ref}
}

func (dcb DeploymentCircuitBreakerAttributes) InternalTokens() hclwrite.Tokens {
	return dcb.ref.InternalTokens()
}

func (dcb DeploymentCircuitBreakerAttributes) Enable() terra.BoolValue {
	return terra.ReferenceAsBool(dcb.ref.Append("enable"))
}

func (dcb DeploymentCircuitBreakerAttributes) Rollback() terra.BoolValue {
	return terra.ReferenceAsBool(dcb.ref.Append("rollback"))
}

type DeploymentControllerAttributes struct {
	ref terra.Reference
}

func (dc DeploymentControllerAttributes) InternalRef() terra.Reference {
	return dc.ref
}

func (dc DeploymentControllerAttributes) InternalWithRef(ref terra.Reference) DeploymentControllerAttributes {
	return DeploymentControllerAttributes{ref: ref}
}

func (dc DeploymentControllerAttributes) InternalTokens() hclwrite.Tokens {
	return dc.ref.InternalTokens()
}

func (dc DeploymentControllerAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("type"))
}

type LoadBalancerAttributes struct {
	ref terra.Reference
}

func (lb LoadBalancerAttributes) InternalRef() terra.Reference {
	return lb.ref
}

func (lb LoadBalancerAttributes) InternalWithRef(ref terra.Reference) LoadBalancerAttributes {
	return LoadBalancerAttributes{ref: ref}
}

func (lb LoadBalancerAttributes) InternalTokens() hclwrite.Tokens {
	return lb.ref.InternalTokens()
}

func (lb LoadBalancerAttributes) ContainerName() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("container_name"))
}

func (lb LoadBalancerAttributes) ContainerPort() terra.NumberValue {
	return terra.ReferenceAsNumber(lb.ref.Append("container_port"))
}

func (lb LoadBalancerAttributes) ElbName() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("elb_name"))
}

func (lb LoadBalancerAttributes) TargetGroupArn() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("target_group_arn"))
}

type NetworkConfigurationAttributes struct {
	ref terra.Reference
}

func (nc NetworkConfigurationAttributes) InternalRef() terra.Reference {
	return nc.ref
}

func (nc NetworkConfigurationAttributes) InternalWithRef(ref terra.Reference) NetworkConfigurationAttributes {
	return NetworkConfigurationAttributes{ref: ref}
}

func (nc NetworkConfigurationAttributes) InternalTokens() hclwrite.Tokens {
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

type OrderedPlacementStrategyAttributes struct {
	ref terra.Reference
}

func (ops OrderedPlacementStrategyAttributes) InternalRef() terra.Reference {
	return ops.ref
}

func (ops OrderedPlacementStrategyAttributes) InternalWithRef(ref terra.Reference) OrderedPlacementStrategyAttributes {
	return OrderedPlacementStrategyAttributes{ref: ref}
}

func (ops OrderedPlacementStrategyAttributes) InternalTokens() hclwrite.Tokens {
	return ops.ref.InternalTokens()
}

func (ops OrderedPlacementStrategyAttributes) Field() terra.StringValue {
	return terra.ReferenceAsString(ops.ref.Append("field"))
}

func (ops OrderedPlacementStrategyAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ops.ref.Append("type"))
}

type PlacementConstraintsAttributes struct {
	ref terra.Reference
}

func (pc PlacementConstraintsAttributes) InternalRef() terra.Reference {
	return pc.ref
}

func (pc PlacementConstraintsAttributes) InternalWithRef(ref terra.Reference) PlacementConstraintsAttributes {
	return PlacementConstraintsAttributes{ref: ref}
}

func (pc PlacementConstraintsAttributes) InternalTokens() hclwrite.Tokens {
	return pc.ref.InternalTokens()
}

func (pc PlacementConstraintsAttributes) Expression() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("expression"))
}

func (pc PlacementConstraintsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("type"))
}

type ServiceConnectConfigurationAttributes struct {
	ref terra.Reference
}

func (scc ServiceConnectConfigurationAttributes) InternalRef() terra.Reference {
	return scc.ref
}

func (scc ServiceConnectConfigurationAttributes) InternalWithRef(ref terra.Reference) ServiceConnectConfigurationAttributes {
	return ServiceConnectConfigurationAttributes{ref: ref}
}

func (scc ServiceConnectConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return scc.ref.InternalTokens()
}

func (scc ServiceConnectConfigurationAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(scc.ref.Append("enabled"))
}

func (scc ServiceConnectConfigurationAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("namespace"))
}

func (scc ServiceConnectConfigurationAttributes) LogConfiguration() terra.ListValue[LogConfigurationAttributes] {
	return terra.ReferenceAsList[LogConfigurationAttributes](scc.ref.Append("log_configuration"))
}

func (scc ServiceConnectConfigurationAttributes) Service() terra.ListValue[ServiceAttributes] {
	return terra.ReferenceAsList[ServiceAttributes](scc.ref.Append("service"))
}

type LogConfigurationAttributes struct {
	ref terra.Reference
}

func (lc LogConfigurationAttributes) InternalRef() terra.Reference {
	return lc.ref
}

func (lc LogConfigurationAttributes) InternalWithRef(ref terra.Reference) LogConfigurationAttributes {
	return LogConfigurationAttributes{ref: ref}
}

func (lc LogConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return lc.ref.InternalTokens()
}

func (lc LogConfigurationAttributes) LogDriver() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("log_driver"))
}

func (lc LogConfigurationAttributes) Options() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lc.ref.Append("options"))
}

func (lc LogConfigurationAttributes) SecretOption() terra.ListValue[SecretOptionAttributes] {
	return terra.ReferenceAsList[SecretOptionAttributes](lc.ref.Append("secret_option"))
}

type SecretOptionAttributes struct {
	ref terra.Reference
}

func (so SecretOptionAttributes) InternalRef() terra.Reference {
	return so.ref
}

func (so SecretOptionAttributes) InternalWithRef(ref terra.Reference) SecretOptionAttributes {
	return SecretOptionAttributes{ref: ref}
}

func (so SecretOptionAttributes) InternalTokens() hclwrite.Tokens {
	return so.ref.InternalTokens()
}

func (so SecretOptionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("name"))
}

func (so SecretOptionAttributes) ValueFrom() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("value_from"))
}

type ServiceAttributes struct {
	ref terra.Reference
}

func (s ServiceAttributes) InternalRef() terra.Reference {
	return s.ref
}

func (s ServiceAttributes) InternalWithRef(ref terra.Reference) ServiceAttributes {
	return ServiceAttributes{ref: ref}
}

func (s ServiceAttributes) InternalTokens() hclwrite.Tokens {
	return s.ref.InternalTokens()
}

func (s ServiceAttributes) DiscoveryName() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("discovery_name"))
}

func (s ServiceAttributes) IngressPortOverride() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("ingress_port_override"))
}

func (s ServiceAttributes) PortName() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("port_name"))
}

func (s ServiceAttributes) ClientAlias() terra.ListValue[ClientAliasAttributes] {
	return terra.ReferenceAsList[ClientAliasAttributes](s.ref.Append("client_alias"))
}

type ClientAliasAttributes struct {
	ref terra.Reference
}

func (ca ClientAliasAttributes) InternalRef() terra.Reference {
	return ca.ref
}

func (ca ClientAliasAttributes) InternalWithRef(ref terra.Reference) ClientAliasAttributes {
	return ClientAliasAttributes{ref: ref}
}

func (ca ClientAliasAttributes) InternalTokens() hclwrite.Tokens {
	return ca.ref.InternalTokens()
}

func (ca ClientAliasAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("dns_name"))
}

func (ca ClientAliasAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(ca.ref.Append("port"))
}

type ServiceRegistriesAttributes struct {
	ref terra.Reference
}

func (sr ServiceRegistriesAttributes) InternalRef() terra.Reference {
	return sr.ref
}

func (sr ServiceRegistriesAttributes) InternalWithRef(ref terra.Reference) ServiceRegistriesAttributes {
	return ServiceRegistriesAttributes{ref: ref}
}

func (sr ServiceRegistriesAttributes) InternalTokens() hclwrite.Tokens {
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type AlarmsState struct {
	AlarmNames []string `json:"alarm_names"`
	Enable     bool     `json:"enable"`
	Rollback   bool     `json:"rollback"`
}

type CapacityProviderStrategyState struct {
	Base             float64 `json:"base"`
	CapacityProvider string  `json:"capacity_provider"`
	Weight           float64 `json:"weight"`
}

type DeploymentCircuitBreakerState struct {
	Enable   bool `json:"enable"`
	Rollback bool `json:"rollback"`
}

type DeploymentControllerState struct {
	Type string `json:"type"`
}

type LoadBalancerState struct {
	ContainerName  string  `json:"container_name"`
	ContainerPort  float64 `json:"container_port"`
	ElbName        string  `json:"elb_name"`
	TargetGroupArn string  `json:"target_group_arn"`
}

type NetworkConfigurationState struct {
	AssignPublicIp bool     `json:"assign_public_ip"`
	SecurityGroups []string `json:"security_groups"`
	Subnets        []string `json:"subnets"`
}

type OrderedPlacementStrategyState struct {
	Field string `json:"field"`
	Type  string `json:"type"`
}

type PlacementConstraintsState struct {
	Expression string `json:"expression"`
	Type       string `json:"type"`
}

type ServiceConnectConfigurationState struct {
	Enabled          bool                    `json:"enabled"`
	Namespace        string                  `json:"namespace"`
	LogConfiguration []LogConfigurationState `json:"log_configuration"`
	Service          []ServiceState          `json:"service"`
}

type LogConfigurationState struct {
	LogDriver    string              `json:"log_driver"`
	Options      map[string]string   `json:"options"`
	SecretOption []SecretOptionState `json:"secret_option"`
}

type SecretOptionState struct {
	Name      string `json:"name"`
	ValueFrom string `json:"value_from"`
}

type ServiceState struct {
	DiscoveryName       string             `json:"discovery_name"`
	IngressPortOverride float64            `json:"ingress_port_override"`
	PortName            string             `json:"port_name"`
	ClientAlias         []ClientAliasState `json:"client_alias"`
}

type ClientAliasState struct {
	DnsName string  `json:"dns_name"`
	Port    float64 `json:"port"`
}

type ServiceRegistriesState struct {
	ContainerName string  `json:"container_name"`
	ContainerPort float64 `json:"container_port"`
	Port          float64 `json:"port"`
	RegistryArn   string  `json:"registry_arn"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}