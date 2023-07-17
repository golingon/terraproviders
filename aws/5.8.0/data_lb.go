// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datalb "github.com/golingon/terraproviders/aws/5.8.0/datalb"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataLb creates a new instance of [DataLb].
func NewDataLb(name string, args DataLbArgs) *DataLb {
	return &DataLb{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLb)(nil)

// DataLb represents the Terraform data resource aws_lb.
type DataLb struct {
	Name string
	Args DataLbArgs
}

// DataSource returns the Terraform object type for [DataLb].
func (l *DataLb) DataSource() string {
	return "aws_lb"
}

// LocalName returns the local name for [DataLb].
func (l *DataLb) LocalName() string {
	return l.Name
}

// Configuration returns the configuration (args) for [DataLb].
func (l *DataLb) Configuration() interface{} {
	return l.Args
}

// Attributes returns the attributes for [DataLb].
func (l *DataLb) Attributes() dataLbAttributes {
	return dataLbAttributes{ref: terra.ReferenceDataResource(l)}
}

// DataLbArgs contains the configurations for aws_lb.
type DataLbArgs struct {
	// Arn: string, optional
	Arn terra.StringValue `hcl:"arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// AccessLogs: min=0
	AccessLogs []datalb.AccessLogs `hcl:"access_logs,block" validate:"min=0"`
	// SubnetMapping: min=0
	SubnetMapping []datalb.SubnetMapping `hcl:"subnet_mapping,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datalb.Timeouts `hcl:"timeouts,block"`
}
type dataLbAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_lb.
func (l dataLbAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("arn"))
}

// ArnSuffix returns a reference to field arn_suffix of aws_lb.
func (l dataLbAttributes) ArnSuffix() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("arn_suffix"))
}

// CustomerOwnedIpv4Pool returns a reference to field customer_owned_ipv4_pool of aws_lb.
func (l dataLbAttributes) CustomerOwnedIpv4Pool() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("customer_owned_ipv4_pool"))
}

// DesyncMitigationMode returns a reference to field desync_mitigation_mode of aws_lb.
func (l dataLbAttributes) DesyncMitigationMode() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("desync_mitigation_mode"))
}

// DnsName returns a reference to field dns_name of aws_lb.
func (l dataLbAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("dns_name"))
}

// DropInvalidHeaderFields returns a reference to field drop_invalid_header_fields of aws_lb.
func (l dataLbAttributes) DropInvalidHeaderFields() terra.BoolValue {
	return terra.ReferenceAsBool(l.ref.Append("drop_invalid_header_fields"))
}

// EnableCrossZoneLoadBalancing returns a reference to field enable_cross_zone_load_balancing of aws_lb.
func (l dataLbAttributes) EnableCrossZoneLoadBalancing() terra.BoolValue {
	return terra.ReferenceAsBool(l.ref.Append("enable_cross_zone_load_balancing"))
}

// EnableDeletionProtection returns a reference to field enable_deletion_protection of aws_lb.
func (l dataLbAttributes) EnableDeletionProtection() terra.BoolValue {
	return terra.ReferenceAsBool(l.ref.Append("enable_deletion_protection"))
}

// EnableHttp2 returns a reference to field enable_http2 of aws_lb.
func (l dataLbAttributes) EnableHttp2() terra.BoolValue {
	return terra.ReferenceAsBool(l.ref.Append("enable_http2"))
}

// EnableTlsVersionAndCipherSuiteHeaders returns a reference to field enable_tls_version_and_cipher_suite_headers of aws_lb.
func (l dataLbAttributes) EnableTlsVersionAndCipherSuiteHeaders() terra.BoolValue {
	return terra.ReferenceAsBool(l.ref.Append("enable_tls_version_and_cipher_suite_headers"))
}

// EnableWafFailOpen returns a reference to field enable_waf_fail_open of aws_lb.
func (l dataLbAttributes) EnableWafFailOpen() terra.BoolValue {
	return terra.ReferenceAsBool(l.ref.Append("enable_waf_fail_open"))
}

// EnableXffClientPort returns a reference to field enable_xff_client_port of aws_lb.
func (l dataLbAttributes) EnableXffClientPort() terra.BoolValue {
	return terra.ReferenceAsBool(l.ref.Append("enable_xff_client_port"))
}

// Id returns a reference to field id of aws_lb.
func (l dataLbAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("id"))
}

// IdleTimeout returns a reference to field idle_timeout of aws_lb.
func (l dataLbAttributes) IdleTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(l.ref.Append("idle_timeout"))
}

// Internal returns a reference to field internal of aws_lb.
func (l dataLbAttributes) Internal() terra.BoolValue {
	return terra.ReferenceAsBool(l.ref.Append("internal"))
}

// IpAddressType returns a reference to field ip_address_type of aws_lb.
func (l dataLbAttributes) IpAddressType() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("ip_address_type"))
}

// LoadBalancerType returns a reference to field load_balancer_type of aws_lb.
func (l dataLbAttributes) LoadBalancerType() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("load_balancer_type"))
}

// Name returns a reference to field name of aws_lb.
func (l dataLbAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("name"))
}

// PreserveHostHeader returns a reference to field preserve_host_header of aws_lb.
func (l dataLbAttributes) PreserveHostHeader() terra.BoolValue {
	return terra.ReferenceAsBool(l.ref.Append("preserve_host_header"))
}

// SecurityGroups returns a reference to field security_groups of aws_lb.
func (l dataLbAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](l.ref.Append("security_groups"))
}

// Subnets returns a reference to field subnets of aws_lb.
func (l dataLbAttributes) Subnets() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](l.ref.Append("subnets"))
}

// Tags returns a reference to field tags of aws_lb.
func (l dataLbAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](l.ref.Append("tags"))
}

// VpcId returns a reference to field vpc_id of aws_lb.
func (l dataLbAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("vpc_id"))
}

// XffHeaderProcessingMode returns a reference to field xff_header_processing_mode of aws_lb.
func (l dataLbAttributes) XffHeaderProcessingMode() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("xff_header_processing_mode"))
}

// ZoneId returns a reference to field zone_id of aws_lb.
func (l dataLbAttributes) ZoneId() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("zone_id"))
}

func (l dataLbAttributes) AccessLogs() terra.ListValue[datalb.AccessLogsAttributes] {
	return terra.ReferenceAsList[datalb.AccessLogsAttributes](l.ref.Append("access_logs"))
}

func (l dataLbAttributes) SubnetMapping() terra.SetValue[datalb.SubnetMappingAttributes] {
	return terra.ReferenceAsSet[datalb.SubnetMappingAttributes](l.ref.Append("subnet_mapping"))
}

func (l dataLbAttributes) Timeouts() datalb.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datalb.TimeoutsAttributes](l.ref.Append("timeouts"))
}
