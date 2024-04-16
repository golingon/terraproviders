// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_ecs_service

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
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
	// ServiceConnectConfigurationLogConfiguration: optional
	LogConfiguration *ServiceConnectConfigurationLogConfiguration `hcl:"log_configuration,block"`
	// ServiceConnectConfigurationService: min=0
	Service []ServiceConnectConfigurationService `hcl:"service,block" validate:"min=0"`
}

type ServiceConnectConfigurationLogConfiguration struct {
	// LogDriver: string, required
	LogDriver terra.StringValue `hcl:"log_driver,attr" validate:"required"`
	// Options: map of string, optional
	Options terra.MapValue[terra.StringValue] `hcl:"options,attr"`
	// ServiceConnectConfigurationLogConfigurationSecretOption: min=0
	SecretOption []ServiceConnectConfigurationLogConfigurationSecretOption `hcl:"secret_option,block" validate:"min=0"`
}

type ServiceConnectConfigurationLogConfigurationSecretOption struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ValueFrom: string, required
	ValueFrom terra.StringValue `hcl:"value_from,attr" validate:"required"`
}

type ServiceConnectConfigurationService struct {
	// DiscoveryName: string, optional
	DiscoveryName terra.StringValue `hcl:"discovery_name,attr"`
	// IngressPortOverride: number, optional
	IngressPortOverride terra.NumberValue `hcl:"ingress_port_override,attr"`
	// PortName: string, required
	PortName terra.StringValue `hcl:"port_name,attr" validate:"required"`
	// ServiceConnectConfigurationServiceClientAlias: optional
	ClientAlias *ServiceConnectConfigurationServiceClientAlias `hcl:"client_alias,block"`
	// ServiceConnectConfigurationServiceTimeout: optional
	Timeout *ServiceConnectConfigurationServiceTimeout `hcl:"timeout,block"`
	// ServiceConnectConfigurationServiceTls: optional
	Tls *ServiceConnectConfigurationServiceTls `hcl:"tls,block"`
}

type ServiceConnectConfigurationServiceClientAlias struct {
	// DnsName: string, optional
	DnsName terra.StringValue `hcl:"dns_name,attr"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
}

type ServiceConnectConfigurationServiceTimeout struct {
	// IdleTimeoutSeconds: number, optional
	IdleTimeoutSeconds terra.NumberValue `hcl:"idle_timeout_seconds,attr"`
	// PerRequestTimeoutSeconds: number, optional
	PerRequestTimeoutSeconds terra.NumberValue `hcl:"per_request_timeout_seconds,attr"`
}

type ServiceConnectConfigurationServiceTls struct {
	// KmsKey: string, optional
	KmsKey terra.StringValue `hcl:"kms_key,attr"`
	// RoleArn: string, optional
	RoleArn terra.StringValue `hcl:"role_arn,attr"`
	// ServiceConnectConfigurationServiceTlsIssuerCertAuthority: required
	IssuerCertAuthority *ServiceConnectConfigurationServiceTlsIssuerCertAuthority `hcl:"issuer_cert_authority,block" validate:"required"`
}

type ServiceConnectConfigurationServiceTlsIssuerCertAuthority struct {
	// AwsPcaAuthorityArn: string, required
	AwsPcaAuthorityArn terra.StringValue `hcl:"aws_pca_authority_arn,attr" validate:"required"`
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

func (a AlarmsAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AlarmsAttributes) InternalWithRef(ref terra.Reference) AlarmsAttributes {
	return AlarmsAttributes{ref: ref}
}

func (a AlarmsAttributes) InternalTokens() (hclwrite.Tokens, error) {
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

type DeploymentCircuitBreakerAttributes struct {
	ref terra.Reference
}

func (dcb DeploymentCircuitBreakerAttributes) InternalRef() (terra.Reference, error) {
	return dcb.ref, nil
}

func (dcb DeploymentCircuitBreakerAttributes) InternalWithRef(ref terra.Reference) DeploymentCircuitBreakerAttributes {
	return DeploymentCircuitBreakerAttributes{ref: ref}
}

func (dcb DeploymentCircuitBreakerAttributes) InternalTokens() (hclwrite.Tokens, error) {
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

func (dc DeploymentControllerAttributes) InternalRef() (terra.Reference, error) {
	return dc.ref, nil
}

func (dc DeploymentControllerAttributes) InternalWithRef(ref terra.Reference) DeploymentControllerAttributes {
	return DeploymentControllerAttributes{ref: ref}
}

func (dc DeploymentControllerAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dc.ref.InternalTokens()
}

func (dc DeploymentControllerAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("type"))
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

func (lb LoadBalancerAttributes) ElbName() terra.StringValue {
	return terra.ReferenceAsString(lb.ref.Append("elb_name"))
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

type OrderedPlacementStrategyAttributes struct {
	ref terra.Reference
}

func (ops OrderedPlacementStrategyAttributes) InternalRef() (terra.Reference, error) {
	return ops.ref, nil
}

func (ops OrderedPlacementStrategyAttributes) InternalWithRef(ref terra.Reference) OrderedPlacementStrategyAttributes {
	return OrderedPlacementStrategyAttributes{ref: ref}
}

func (ops OrderedPlacementStrategyAttributes) InternalTokens() (hclwrite.Tokens, error) {
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

func (pc PlacementConstraintsAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc PlacementConstraintsAttributes) InternalWithRef(ref terra.Reference) PlacementConstraintsAttributes {
	return PlacementConstraintsAttributes{ref: ref}
}

func (pc PlacementConstraintsAttributes) InternalTokens() (hclwrite.Tokens, error) {
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

func (scc ServiceConnectConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return scc.ref, nil
}

func (scc ServiceConnectConfigurationAttributes) InternalWithRef(ref terra.Reference) ServiceConnectConfigurationAttributes {
	return ServiceConnectConfigurationAttributes{ref: ref}
}

func (scc ServiceConnectConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return scc.ref.InternalTokens()
}

func (scc ServiceConnectConfigurationAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(scc.ref.Append("enabled"))
}

func (scc ServiceConnectConfigurationAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("namespace"))
}

func (scc ServiceConnectConfigurationAttributes) LogConfiguration() terra.ListValue[ServiceConnectConfigurationLogConfigurationAttributes] {
	return terra.ReferenceAsList[ServiceConnectConfigurationLogConfigurationAttributes](scc.ref.Append("log_configuration"))
}

func (scc ServiceConnectConfigurationAttributes) Service() terra.ListValue[ServiceConnectConfigurationServiceAttributes] {
	return terra.ReferenceAsList[ServiceConnectConfigurationServiceAttributes](scc.ref.Append("service"))
}

type ServiceConnectConfigurationLogConfigurationAttributes struct {
	ref terra.Reference
}

func (lc ServiceConnectConfigurationLogConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return lc.ref, nil
}

func (lc ServiceConnectConfigurationLogConfigurationAttributes) InternalWithRef(ref terra.Reference) ServiceConnectConfigurationLogConfigurationAttributes {
	return ServiceConnectConfigurationLogConfigurationAttributes{ref: ref}
}

func (lc ServiceConnectConfigurationLogConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lc.ref.InternalTokens()
}

func (lc ServiceConnectConfigurationLogConfigurationAttributes) LogDriver() terra.StringValue {
	return terra.ReferenceAsString(lc.ref.Append("log_driver"))
}

func (lc ServiceConnectConfigurationLogConfigurationAttributes) Options() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lc.ref.Append("options"))
}

func (lc ServiceConnectConfigurationLogConfigurationAttributes) SecretOption() terra.ListValue[ServiceConnectConfigurationLogConfigurationSecretOptionAttributes] {
	return terra.ReferenceAsList[ServiceConnectConfigurationLogConfigurationSecretOptionAttributes](lc.ref.Append("secret_option"))
}

type ServiceConnectConfigurationLogConfigurationSecretOptionAttributes struct {
	ref terra.Reference
}

func (so ServiceConnectConfigurationLogConfigurationSecretOptionAttributes) InternalRef() (terra.Reference, error) {
	return so.ref, nil
}

func (so ServiceConnectConfigurationLogConfigurationSecretOptionAttributes) InternalWithRef(ref terra.Reference) ServiceConnectConfigurationLogConfigurationSecretOptionAttributes {
	return ServiceConnectConfigurationLogConfigurationSecretOptionAttributes{ref: ref}
}

func (so ServiceConnectConfigurationLogConfigurationSecretOptionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return so.ref.InternalTokens()
}

func (so ServiceConnectConfigurationLogConfigurationSecretOptionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("name"))
}

func (so ServiceConnectConfigurationLogConfigurationSecretOptionAttributes) ValueFrom() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("value_from"))
}

type ServiceConnectConfigurationServiceAttributes struct {
	ref terra.Reference
}

func (s ServiceConnectConfigurationServiceAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s ServiceConnectConfigurationServiceAttributes) InternalWithRef(ref terra.Reference) ServiceConnectConfigurationServiceAttributes {
	return ServiceConnectConfigurationServiceAttributes{ref: ref}
}

func (s ServiceConnectConfigurationServiceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s ServiceConnectConfigurationServiceAttributes) DiscoveryName() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("discovery_name"))
}

func (s ServiceConnectConfigurationServiceAttributes) IngressPortOverride() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("ingress_port_override"))
}

func (s ServiceConnectConfigurationServiceAttributes) PortName() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("port_name"))
}

func (s ServiceConnectConfigurationServiceAttributes) ClientAlias() terra.ListValue[ServiceConnectConfigurationServiceClientAliasAttributes] {
	return terra.ReferenceAsList[ServiceConnectConfigurationServiceClientAliasAttributes](s.ref.Append("client_alias"))
}

func (s ServiceConnectConfigurationServiceAttributes) Timeout() terra.ListValue[ServiceConnectConfigurationServiceTimeoutAttributes] {
	return terra.ReferenceAsList[ServiceConnectConfigurationServiceTimeoutAttributes](s.ref.Append("timeout"))
}

func (s ServiceConnectConfigurationServiceAttributes) Tls() terra.ListValue[ServiceConnectConfigurationServiceTlsAttributes] {
	return terra.ReferenceAsList[ServiceConnectConfigurationServiceTlsAttributes](s.ref.Append("tls"))
}

type ServiceConnectConfigurationServiceClientAliasAttributes struct {
	ref terra.Reference
}

func (ca ServiceConnectConfigurationServiceClientAliasAttributes) InternalRef() (terra.Reference, error) {
	return ca.ref, nil
}

func (ca ServiceConnectConfigurationServiceClientAliasAttributes) InternalWithRef(ref terra.Reference) ServiceConnectConfigurationServiceClientAliasAttributes {
	return ServiceConnectConfigurationServiceClientAliasAttributes{ref: ref}
}

func (ca ServiceConnectConfigurationServiceClientAliasAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ca.ref.InternalTokens()
}

func (ca ServiceConnectConfigurationServiceClientAliasAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("dns_name"))
}

func (ca ServiceConnectConfigurationServiceClientAliasAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(ca.ref.Append("port"))
}

type ServiceConnectConfigurationServiceTimeoutAttributes struct {
	ref terra.Reference
}

func (t ServiceConnectConfigurationServiceTimeoutAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t ServiceConnectConfigurationServiceTimeoutAttributes) InternalWithRef(ref terra.Reference) ServiceConnectConfigurationServiceTimeoutAttributes {
	return ServiceConnectConfigurationServiceTimeoutAttributes{ref: ref}
}

func (t ServiceConnectConfigurationServiceTimeoutAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t ServiceConnectConfigurationServiceTimeoutAttributes) IdleTimeoutSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(t.ref.Append("idle_timeout_seconds"))
}

func (t ServiceConnectConfigurationServiceTimeoutAttributes) PerRequestTimeoutSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(t.ref.Append("per_request_timeout_seconds"))
}

type ServiceConnectConfigurationServiceTlsAttributes struct {
	ref terra.Reference
}

func (t ServiceConnectConfigurationServiceTlsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t ServiceConnectConfigurationServiceTlsAttributes) InternalWithRef(ref terra.Reference) ServiceConnectConfigurationServiceTlsAttributes {
	return ServiceConnectConfigurationServiceTlsAttributes{ref: ref}
}

func (t ServiceConnectConfigurationServiceTlsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t ServiceConnectConfigurationServiceTlsAttributes) KmsKey() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("kms_key"))
}

func (t ServiceConnectConfigurationServiceTlsAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("role_arn"))
}

func (t ServiceConnectConfigurationServiceTlsAttributes) IssuerCertAuthority() terra.ListValue[ServiceConnectConfigurationServiceTlsIssuerCertAuthorityAttributes] {
	return terra.ReferenceAsList[ServiceConnectConfigurationServiceTlsIssuerCertAuthorityAttributes](t.ref.Append("issuer_cert_authority"))
}

type ServiceConnectConfigurationServiceTlsIssuerCertAuthorityAttributes struct {
	ref terra.Reference
}

func (ica ServiceConnectConfigurationServiceTlsIssuerCertAuthorityAttributes) InternalRef() (terra.Reference, error) {
	return ica.ref, nil
}

func (ica ServiceConnectConfigurationServiceTlsIssuerCertAuthorityAttributes) InternalWithRef(ref terra.Reference) ServiceConnectConfigurationServiceTlsIssuerCertAuthorityAttributes {
	return ServiceConnectConfigurationServiceTlsIssuerCertAuthorityAttributes{ref: ref}
}

func (ica ServiceConnectConfigurationServiceTlsIssuerCertAuthorityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ica.ref.InternalTokens()
}

func (ica ServiceConnectConfigurationServiceTlsIssuerCertAuthorityAttributes) AwsPcaAuthorityArn() terra.StringValue {
	return terra.ReferenceAsString(ica.ref.Append("aws_pca_authority_arn"))
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
	Enabled          bool                                               `json:"enabled"`
	Namespace        string                                             `json:"namespace"`
	LogConfiguration []ServiceConnectConfigurationLogConfigurationState `json:"log_configuration"`
	Service          []ServiceConnectConfigurationServiceState          `json:"service"`
}

type ServiceConnectConfigurationLogConfigurationState struct {
	LogDriver    string                                                         `json:"log_driver"`
	Options      map[string]string                                              `json:"options"`
	SecretOption []ServiceConnectConfigurationLogConfigurationSecretOptionState `json:"secret_option"`
}

type ServiceConnectConfigurationLogConfigurationSecretOptionState struct {
	Name      string `json:"name"`
	ValueFrom string `json:"value_from"`
}

type ServiceConnectConfigurationServiceState struct {
	DiscoveryName       string                                               `json:"discovery_name"`
	IngressPortOverride float64                                              `json:"ingress_port_override"`
	PortName            string                                               `json:"port_name"`
	ClientAlias         []ServiceConnectConfigurationServiceClientAliasState `json:"client_alias"`
	Timeout             []ServiceConnectConfigurationServiceTimeoutState     `json:"timeout"`
	Tls                 []ServiceConnectConfigurationServiceTlsState         `json:"tls"`
}

type ServiceConnectConfigurationServiceClientAliasState struct {
	DnsName string  `json:"dns_name"`
	Port    float64 `json:"port"`
}

type ServiceConnectConfigurationServiceTimeoutState struct {
	IdleTimeoutSeconds       float64 `json:"idle_timeout_seconds"`
	PerRequestTimeoutSeconds float64 `json:"per_request_timeout_seconds"`
}

type ServiceConnectConfigurationServiceTlsState struct {
	KmsKey              string                                                          `json:"kms_key"`
	RoleArn             string                                                          `json:"role_arn"`
	IssuerCertAuthority []ServiceConnectConfigurationServiceTlsIssuerCertAuthorityState `json:"issuer_cert_authority"`
}

type ServiceConnectConfigurationServiceTlsIssuerCertAuthorityState struct {
	AwsPcaAuthorityArn string `json:"aws_pca_authority_arn"`
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
