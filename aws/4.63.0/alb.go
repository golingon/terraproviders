// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	alb "github.com/golingon/terraproviders/aws/4.63.0/alb"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAlb creates a new instance of [Alb].
func NewAlb(name string, args AlbArgs) *Alb {
	return &Alb{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Alb)(nil)

// Alb represents the Terraform resource aws_alb.
type Alb struct {
	Name      string
	Args      AlbArgs
	state     *albState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Alb].
func (a *Alb) Type() string {
	return "aws_alb"
}

// LocalName returns the local name for [Alb].
func (a *Alb) LocalName() string {
	return a.Name
}

// Configuration returns the configuration (args) for [Alb].
func (a *Alb) Configuration() interface{} {
	return a.Args
}

// DependOn is used for other resources to depend on [Alb].
func (a *Alb) DependOn() terra.Reference {
	return terra.ReferenceResource(a)
}

// Dependencies returns the list of resources [Alb] depends_on.
func (a *Alb) Dependencies() terra.Dependencies {
	return a.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Alb].
func (a *Alb) LifecycleManagement() *terra.Lifecycle {
	return a.Lifecycle
}

// Attributes returns the attributes for [Alb].
func (a *Alb) Attributes() albAttributes {
	return albAttributes{ref: terra.ReferenceResource(a)}
}

// ImportState imports the given attribute values into [Alb]'s state.
func (a *Alb) ImportState(av io.Reader) error {
	a.state = &albState{}
	if err := json.NewDecoder(av).Decode(a.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", a.Type(), a.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Alb] has state.
func (a *Alb) State() (*albState, bool) {
	return a.state, a.state != nil
}

// StateMust returns the state for [Alb]. Panics if the state is nil.
func (a *Alb) StateMust() *albState {
	if a.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", a.Type(), a.LocalName()))
	}
	return a.state
}

// AlbArgs contains the configurations for aws_alb.
type AlbArgs struct {
	// CustomerOwnedIpv4Pool: string, optional
	CustomerOwnedIpv4Pool terra.StringValue `hcl:"customer_owned_ipv4_pool,attr"`
	// DesyncMitigationMode: string, optional
	DesyncMitigationMode terra.StringValue `hcl:"desync_mitigation_mode,attr"`
	// DropInvalidHeaderFields: bool, optional
	DropInvalidHeaderFields terra.BoolValue `hcl:"drop_invalid_header_fields,attr"`
	// EnableCrossZoneLoadBalancing: bool, optional
	EnableCrossZoneLoadBalancing terra.BoolValue `hcl:"enable_cross_zone_load_balancing,attr"`
	// EnableDeletionProtection: bool, optional
	EnableDeletionProtection terra.BoolValue `hcl:"enable_deletion_protection,attr"`
	// EnableHttp2: bool, optional
	EnableHttp2 terra.BoolValue `hcl:"enable_http2,attr"`
	// EnableTlsVersionAndCipherSuiteHeaders: bool, optional
	EnableTlsVersionAndCipherSuiteHeaders terra.BoolValue `hcl:"enable_tls_version_and_cipher_suite_headers,attr"`
	// EnableWafFailOpen: bool, optional
	EnableWafFailOpen terra.BoolValue `hcl:"enable_waf_fail_open,attr"`
	// EnableXffClientPort: bool, optional
	EnableXffClientPort terra.BoolValue `hcl:"enable_xff_client_port,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdleTimeout: number, optional
	IdleTimeout terra.NumberValue `hcl:"idle_timeout,attr"`
	// Internal: bool, optional
	Internal terra.BoolValue `hcl:"internal,attr"`
	// IpAddressType: string, optional
	IpAddressType terra.StringValue `hcl:"ip_address_type,attr"`
	// LoadBalancerType: string, optional
	LoadBalancerType terra.StringValue `hcl:"load_balancer_type,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// PreserveHostHeader: bool, optional
	PreserveHostHeader terra.BoolValue `hcl:"preserve_host_header,attr"`
	// SecurityGroups: set of string, optional
	SecurityGroups terra.SetValue[terra.StringValue] `hcl:"security_groups,attr"`
	// Subnets: set of string, optional
	Subnets terra.SetValue[terra.StringValue] `hcl:"subnets,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// XffHeaderProcessingMode: string, optional
	XffHeaderProcessingMode terra.StringValue `hcl:"xff_header_processing_mode,attr"`
	// AccessLogs: optional
	AccessLogs *alb.AccessLogs `hcl:"access_logs,block"`
	// SubnetMapping: min=0
	SubnetMapping []alb.SubnetMapping `hcl:"subnet_mapping,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *alb.Timeouts `hcl:"timeouts,block"`
}
type albAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_alb.
func (a albAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("arn"))
}

// ArnSuffix returns a reference to field arn_suffix of aws_alb.
func (a albAttributes) ArnSuffix() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("arn_suffix"))
}

// CustomerOwnedIpv4Pool returns a reference to field customer_owned_ipv4_pool of aws_alb.
func (a albAttributes) CustomerOwnedIpv4Pool() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("customer_owned_ipv4_pool"))
}

// DesyncMitigationMode returns a reference to field desync_mitigation_mode of aws_alb.
func (a albAttributes) DesyncMitigationMode() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("desync_mitigation_mode"))
}

// DnsName returns a reference to field dns_name of aws_alb.
func (a albAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("dns_name"))
}

// DropInvalidHeaderFields returns a reference to field drop_invalid_header_fields of aws_alb.
func (a albAttributes) DropInvalidHeaderFields() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("drop_invalid_header_fields"))
}

// EnableCrossZoneLoadBalancing returns a reference to field enable_cross_zone_load_balancing of aws_alb.
func (a albAttributes) EnableCrossZoneLoadBalancing() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("enable_cross_zone_load_balancing"))
}

// EnableDeletionProtection returns a reference to field enable_deletion_protection of aws_alb.
func (a albAttributes) EnableDeletionProtection() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("enable_deletion_protection"))
}

// EnableHttp2 returns a reference to field enable_http2 of aws_alb.
func (a albAttributes) EnableHttp2() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("enable_http2"))
}

// EnableTlsVersionAndCipherSuiteHeaders returns a reference to field enable_tls_version_and_cipher_suite_headers of aws_alb.
func (a albAttributes) EnableTlsVersionAndCipherSuiteHeaders() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("enable_tls_version_and_cipher_suite_headers"))
}

// EnableWafFailOpen returns a reference to field enable_waf_fail_open of aws_alb.
func (a albAttributes) EnableWafFailOpen() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("enable_waf_fail_open"))
}

// EnableXffClientPort returns a reference to field enable_xff_client_port of aws_alb.
func (a albAttributes) EnableXffClientPort() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("enable_xff_client_port"))
}

// Id returns a reference to field id of aws_alb.
func (a albAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("id"))
}

// IdleTimeout returns a reference to field idle_timeout of aws_alb.
func (a albAttributes) IdleTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(a.ref.Append("idle_timeout"))
}

// Internal returns a reference to field internal of aws_alb.
func (a albAttributes) Internal() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("internal"))
}

// IpAddressType returns a reference to field ip_address_type of aws_alb.
func (a albAttributes) IpAddressType() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("ip_address_type"))
}

// LoadBalancerType returns a reference to field load_balancer_type of aws_alb.
func (a albAttributes) LoadBalancerType() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("load_balancer_type"))
}

// Name returns a reference to field name of aws_alb.
func (a albAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_alb.
func (a albAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("name_prefix"))
}

// PreserveHostHeader returns a reference to field preserve_host_header of aws_alb.
func (a albAttributes) PreserveHostHeader() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("preserve_host_header"))
}

// SecurityGroups returns a reference to field security_groups of aws_alb.
func (a albAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](a.ref.Append("security_groups"))
}

// Subnets returns a reference to field subnets of aws_alb.
func (a albAttributes) Subnets() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](a.ref.Append("subnets"))
}

// Tags returns a reference to field tags of aws_alb.
func (a albAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](a.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_alb.
func (a albAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](a.ref.Append("tags_all"))
}

// VpcId returns a reference to field vpc_id of aws_alb.
func (a albAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("vpc_id"))
}

// XffHeaderProcessingMode returns a reference to field xff_header_processing_mode of aws_alb.
func (a albAttributes) XffHeaderProcessingMode() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("xff_header_processing_mode"))
}

// ZoneId returns a reference to field zone_id of aws_alb.
func (a albAttributes) ZoneId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("zone_id"))
}

func (a albAttributes) AccessLogs() terra.ListValue[alb.AccessLogsAttributes] {
	return terra.ReferenceAsList[alb.AccessLogsAttributes](a.ref.Append("access_logs"))
}

func (a albAttributes) SubnetMapping() terra.SetValue[alb.SubnetMappingAttributes] {
	return terra.ReferenceAsSet[alb.SubnetMappingAttributes](a.ref.Append("subnet_mapping"))
}

func (a albAttributes) Timeouts() alb.TimeoutsAttributes {
	return terra.ReferenceAsSingle[alb.TimeoutsAttributes](a.ref.Append("timeouts"))
}

type albState struct {
	Arn                                   string                   `json:"arn"`
	ArnSuffix                             string                   `json:"arn_suffix"`
	CustomerOwnedIpv4Pool                 string                   `json:"customer_owned_ipv4_pool"`
	DesyncMitigationMode                  string                   `json:"desync_mitigation_mode"`
	DnsName                               string                   `json:"dns_name"`
	DropInvalidHeaderFields               bool                     `json:"drop_invalid_header_fields"`
	EnableCrossZoneLoadBalancing          bool                     `json:"enable_cross_zone_load_balancing"`
	EnableDeletionProtection              bool                     `json:"enable_deletion_protection"`
	EnableHttp2                           bool                     `json:"enable_http2"`
	EnableTlsVersionAndCipherSuiteHeaders bool                     `json:"enable_tls_version_and_cipher_suite_headers"`
	EnableWafFailOpen                     bool                     `json:"enable_waf_fail_open"`
	EnableXffClientPort                   bool                     `json:"enable_xff_client_port"`
	Id                                    string                   `json:"id"`
	IdleTimeout                           float64                  `json:"idle_timeout"`
	Internal                              bool                     `json:"internal"`
	IpAddressType                         string                   `json:"ip_address_type"`
	LoadBalancerType                      string                   `json:"load_balancer_type"`
	Name                                  string                   `json:"name"`
	NamePrefix                            string                   `json:"name_prefix"`
	PreserveHostHeader                    bool                     `json:"preserve_host_header"`
	SecurityGroups                        []string                 `json:"security_groups"`
	Subnets                               []string                 `json:"subnets"`
	Tags                                  map[string]string        `json:"tags"`
	TagsAll                               map[string]string        `json:"tags_all"`
	VpcId                                 string                   `json:"vpc_id"`
	XffHeaderProcessingMode               string                   `json:"xff_header_processing_mode"`
	ZoneId                                string                   `json:"zone_id"`
	AccessLogs                            []alb.AccessLogsState    `json:"access_logs"`
	SubnetMapping                         []alb.SubnetMappingState `json:"subnet_mapping"`
	Timeouts                              *alb.TimeoutsState       `json:"timeouts"`
}
