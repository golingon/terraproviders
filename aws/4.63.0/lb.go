// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	lb "github.com/golingon/terraproviders/aws/4.63.0/lb"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLb creates a new instance of [Lb].
func NewLb(name string, args LbArgs) *Lb {
	return &Lb{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Lb)(nil)

// Lb represents the Terraform resource aws_lb.
type Lb struct {
	Name      string
	Args      LbArgs
	state     *lbState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Lb].
func (l *Lb) Type() string {
	return "aws_lb"
}

// LocalName returns the local name for [Lb].
func (l *Lb) LocalName() string {
	return l.Name
}

// Configuration returns the configuration (args) for [Lb].
func (l *Lb) Configuration() interface{} {
	return l.Args
}

// DependOn is used for other resources to depend on [Lb].
func (l *Lb) DependOn() terra.Reference {
	return terra.ReferenceResource(l)
}

// Dependencies returns the list of resources [Lb] depends_on.
func (l *Lb) Dependencies() terra.Dependencies {
	return l.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Lb].
func (l *Lb) LifecycleManagement() *terra.Lifecycle {
	return l.Lifecycle
}

// Attributes returns the attributes for [Lb].
func (l *Lb) Attributes() lbAttributes {
	return lbAttributes{ref: terra.ReferenceResource(l)}
}

// ImportState imports the given attribute values into [Lb]'s state.
func (l *Lb) ImportState(av io.Reader) error {
	l.state = &lbState{}
	if err := json.NewDecoder(av).Decode(l.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", l.Type(), l.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Lb] has state.
func (l *Lb) State() (*lbState, bool) {
	return l.state, l.state != nil
}

// StateMust returns the state for [Lb]. Panics if the state is nil.
func (l *Lb) StateMust() *lbState {
	if l.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", l.Type(), l.LocalName()))
	}
	return l.state
}

// LbArgs contains the configurations for aws_lb.
type LbArgs struct {
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
	AccessLogs *lb.AccessLogs `hcl:"access_logs,block"`
	// SubnetMapping: min=0
	SubnetMapping []lb.SubnetMapping `hcl:"subnet_mapping,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *lb.Timeouts `hcl:"timeouts,block"`
}
type lbAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_lb.
func (l lbAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("arn"))
}

// ArnSuffix returns a reference to field arn_suffix of aws_lb.
func (l lbAttributes) ArnSuffix() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("arn_suffix"))
}

// CustomerOwnedIpv4Pool returns a reference to field customer_owned_ipv4_pool of aws_lb.
func (l lbAttributes) CustomerOwnedIpv4Pool() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("customer_owned_ipv4_pool"))
}

// DesyncMitigationMode returns a reference to field desync_mitigation_mode of aws_lb.
func (l lbAttributes) DesyncMitigationMode() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("desync_mitigation_mode"))
}

// DnsName returns a reference to field dns_name of aws_lb.
func (l lbAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("dns_name"))
}

// DropInvalidHeaderFields returns a reference to field drop_invalid_header_fields of aws_lb.
func (l lbAttributes) DropInvalidHeaderFields() terra.BoolValue {
	return terra.ReferenceAsBool(l.ref.Append("drop_invalid_header_fields"))
}

// EnableCrossZoneLoadBalancing returns a reference to field enable_cross_zone_load_balancing of aws_lb.
func (l lbAttributes) EnableCrossZoneLoadBalancing() terra.BoolValue {
	return terra.ReferenceAsBool(l.ref.Append("enable_cross_zone_load_balancing"))
}

// EnableDeletionProtection returns a reference to field enable_deletion_protection of aws_lb.
func (l lbAttributes) EnableDeletionProtection() terra.BoolValue {
	return terra.ReferenceAsBool(l.ref.Append("enable_deletion_protection"))
}

// EnableHttp2 returns a reference to field enable_http2 of aws_lb.
func (l lbAttributes) EnableHttp2() terra.BoolValue {
	return terra.ReferenceAsBool(l.ref.Append("enable_http2"))
}

// EnableTlsVersionAndCipherSuiteHeaders returns a reference to field enable_tls_version_and_cipher_suite_headers of aws_lb.
func (l lbAttributes) EnableTlsVersionAndCipherSuiteHeaders() terra.BoolValue {
	return terra.ReferenceAsBool(l.ref.Append("enable_tls_version_and_cipher_suite_headers"))
}

// EnableWafFailOpen returns a reference to field enable_waf_fail_open of aws_lb.
func (l lbAttributes) EnableWafFailOpen() terra.BoolValue {
	return terra.ReferenceAsBool(l.ref.Append("enable_waf_fail_open"))
}

// EnableXffClientPort returns a reference to field enable_xff_client_port of aws_lb.
func (l lbAttributes) EnableXffClientPort() terra.BoolValue {
	return terra.ReferenceAsBool(l.ref.Append("enable_xff_client_port"))
}

// Id returns a reference to field id of aws_lb.
func (l lbAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("id"))
}

// IdleTimeout returns a reference to field idle_timeout of aws_lb.
func (l lbAttributes) IdleTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(l.ref.Append("idle_timeout"))
}

// Internal returns a reference to field internal of aws_lb.
func (l lbAttributes) Internal() terra.BoolValue {
	return terra.ReferenceAsBool(l.ref.Append("internal"))
}

// IpAddressType returns a reference to field ip_address_type of aws_lb.
func (l lbAttributes) IpAddressType() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("ip_address_type"))
}

// LoadBalancerType returns a reference to field load_balancer_type of aws_lb.
func (l lbAttributes) LoadBalancerType() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("load_balancer_type"))
}

// Name returns a reference to field name of aws_lb.
func (l lbAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_lb.
func (l lbAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("name_prefix"))
}

// PreserveHostHeader returns a reference to field preserve_host_header of aws_lb.
func (l lbAttributes) PreserveHostHeader() terra.BoolValue {
	return terra.ReferenceAsBool(l.ref.Append("preserve_host_header"))
}

// SecurityGroups returns a reference to field security_groups of aws_lb.
func (l lbAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](l.ref.Append("security_groups"))
}

// Subnets returns a reference to field subnets of aws_lb.
func (l lbAttributes) Subnets() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](l.ref.Append("subnets"))
}

// Tags returns a reference to field tags of aws_lb.
func (l lbAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](l.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_lb.
func (l lbAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](l.ref.Append("tags_all"))
}

// VpcId returns a reference to field vpc_id of aws_lb.
func (l lbAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("vpc_id"))
}

// XffHeaderProcessingMode returns a reference to field xff_header_processing_mode of aws_lb.
func (l lbAttributes) XffHeaderProcessingMode() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("xff_header_processing_mode"))
}

// ZoneId returns a reference to field zone_id of aws_lb.
func (l lbAttributes) ZoneId() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("zone_id"))
}

func (l lbAttributes) AccessLogs() terra.ListValue[lb.AccessLogsAttributes] {
	return terra.ReferenceAsList[lb.AccessLogsAttributes](l.ref.Append("access_logs"))
}

func (l lbAttributes) SubnetMapping() terra.SetValue[lb.SubnetMappingAttributes] {
	return terra.ReferenceAsSet[lb.SubnetMappingAttributes](l.ref.Append("subnet_mapping"))
}

func (l lbAttributes) Timeouts() lb.TimeoutsAttributes {
	return terra.ReferenceAsSingle[lb.TimeoutsAttributes](l.ref.Append("timeouts"))
}

type lbState struct {
	Arn                                   string                  `json:"arn"`
	ArnSuffix                             string                  `json:"arn_suffix"`
	CustomerOwnedIpv4Pool                 string                  `json:"customer_owned_ipv4_pool"`
	DesyncMitigationMode                  string                  `json:"desync_mitigation_mode"`
	DnsName                               string                  `json:"dns_name"`
	DropInvalidHeaderFields               bool                    `json:"drop_invalid_header_fields"`
	EnableCrossZoneLoadBalancing          bool                    `json:"enable_cross_zone_load_balancing"`
	EnableDeletionProtection              bool                    `json:"enable_deletion_protection"`
	EnableHttp2                           bool                    `json:"enable_http2"`
	EnableTlsVersionAndCipherSuiteHeaders bool                    `json:"enable_tls_version_and_cipher_suite_headers"`
	EnableWafFailOpen                     bool                    `json:"enable_waf_fail_open"`
	EnableXffClientPort                   bool                    `json:"enable_xff_client_port"`
	Id                                    string                  `json:"id"`
	IdleTimeout                           float64                 `json:"idle_timeout"`
	Internal                              bool                    `json:"internal"`
	IpAddressType                         string                  `json:"ip_address_type"`
	LoadBalancerType                      string                  `json:"load_balancer_type"`
	Name                                  string                  `json:"name"`
	NamePrefix                            string                  `json:"name_prefix"`
	PreserveHostHeader                    bool                    `json:"preserve_host_header"`
	SecurityGroups                        []string                `json:"security_groups"`
	Subnets                               []string                `json:"subnets"`
	Tags                                  map[string]string       `json:"tags"`
	TagsAll                               map[string]string       `json:"tags_all"`
	VpcId                                 string                  `json:"vpc_id"`
	XffHeaderProcessingMode               string                  `json:"xff_header_processing_mode"`
	ZoneId                                string                  `json:"zone_id"`
	AccessLogs                            []lb.AccessLogsState    `json:"access_logs"`
	SubnetMapping                         []lb.SubnetMappingState `json:"subnet_mapping"`
	Timeouts                              *lb.TimeoutsState       `json:"timeouts"`
}
