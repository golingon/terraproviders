// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataalb "github.com/golingon/terraproviders/aws/5.8.0/dataalb"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAlb creates a new instance of [DataAlb].
func NewDataAlb(name string, args DataAlbArgs) *DataAlb {
	return &DataAlb{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAlb)(nil)

// DataAlb represents the Terraform data resource aws_alb.
type DataAlb struct {
	Name string
	Args DataAlbArgs
}

// DataSource returns the Terraform object type for [DataAlb].
func (a *DataAlb) DataSource() string {
	return "aws_alb"
}

// LocalName returns the local name for [DataAlb].
func (a *DataAlb) LocalName() string {
	return a.Name
}

// Configuration returns the configuration (args) for [DataAlb].
func (a *DataAlb) Configuration() interface{} {
	return a.Args
}

// Attributes returns the attributes for [DataAlb].
func (a *DataAlb) Attributes() dataAlbAttributes {
	return dataAlbAttributes{ref: terra.ReferenceDataResource(a)}
}

// DataAlbArgs contains the configurations for aws_alb.
type DataAlbArgs struct {
	// Arn: string, optional
	Arn terra.StringValue `hcl:"arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// AccessLogs: min=0
	AccessLogs []dataalb.AccessLogs `hcl:"access_logs,block" validate:"min=0"`
	// SubnetMapping: min=0
	SubnetMapping []dataalb.SubnetMapping `hcl:"subnet_mapping,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataalb.Timeouts `hcl:"timeouts,block"`
}
type dataAlbAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_alb.
func (a dataAlbAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("arn"))
}

// ArnSuffix returns a reference to field arn_suffix of aws_alb.
func (a dataAlbAttributes) ArnSuffix() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("arn_suffix"))
}

// CustomerOwnedIpv4Pool returns a reference to field customer_owned_ipv4_pool of aws_alb.
func (a dataAlbAttributes) CustomerOwnedIpv4Pool() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("customer_owned_ipv4_pool"))
}

// DesyncMitigationMode returns a reference to field desync_mitigation_mode of aws_alb.
func (a dataAlbAttributes) DesyncMitigationMode() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("desync_mitigation_mode"))
}

// DnsName returns a reference to field dns_name of aws_alb.
func (a dataAlbAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("dns_name"))
}

// DropInvalidHeaderFields returns a reference to field drop_invalid_header_fields of aws_alb.
func (a dataAlbAttributes) DropInvalidHeaderFields() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("drop_invalid_header_fields"))
}

// EnableCrossZoneLoadBalancing returns a reference to field enable_cross_zone_load_balancing of aws_alb.
func (a dataAlbAttributes) EnableCrossZoneLoadBalancing() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("enable_cross_zone_load_balancing"))
}

// EnableDeletionProtection returns a reference to field enable_deletion_protection of aws_alb.
func (a dataAlbAttributes) EnableDeletionProtection() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("enable_deletion_protection"))
}

// EnableHttp2 returns a reference to field enable_http2 of aws_alb.
func (a dataAlbAttributes) EnableHttp2() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("enable_http2"))
}

// EnableTlsVersionAndCipherSuiteHeaders returns a reference to field enable_tls_version_and_cipher_suite_headers of aws_alb.
func (a dataAlbAttributes) EnableTlsVersionAndCipherSuiteHeaders() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("enable_tls_version_and_cipher_suite_headers"))
}

// EnableWafFailOpen returns a reference to field enable_waf_fail_open of aws_alb.
func (a dataAlbAttributes) EnableWafFailOpen() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("enable_waf_fail_open"))
}

// EnableXffClientPort returns a reference to field enable_xff_client_port of aws_alb.
func (a dataAlbAttributes) EnableXffClientPort() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("enable_xff_client_port"))
}

// Id returns a reference to field id of aws_alb.
func (a dataAlbAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("id"))
}

// IdleTimeout returns a reference to field idle_timeout of aws_alb.
func (a dataAlbAttributes) IdleTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(a.ref.Append("idle_timeout"))
}

// Internal returns a reference to field internal of aws_alb.
func (a dataAlbAttributes) Internal() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("internal"))
}

// IpAddressType returns a reference to field ip_address_type of aws_alb.
func (a dataAlbAttributes) IpAddressType() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("ip_address_type"))
}

// LoadBalancerType returns a reference to field load_balancer_type of aws_alb.
func (a dataAlbAttributes) LoadBalancerType() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("load_balancer_type"))
}

// Name returns a reference to field name of aws_alb.
func (a dataAlbAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("name"))
}

// PreserveHostHeader returns a reference to field preserve_host_header of aws_alb.
func (a dataAlbAttributes) PreserveHostHeader() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("preserve_host_header"))
}

// SecurityGroups returns a reference to field security_groups of aws_alb.
func (a dataAlbAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](a.ref.Append("security_groups"))
}

// Subnets returns a reference to field subnets of aws_alb.
func (a dataAlbAttributes) Subnets() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](a.ref.Append("subnets"))
}

// Tags returns a reference to field tags of aws_alb.
func (a dataAlbAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](a.ref.Append("tags"))
}

// VpcId returns a reference to field vpc_id of aws_alb.
func (a dataAlbAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("vpc_id"))
}

// XffHeaderProcessingMode returns a reference to field xff_header_processing_mode of aws_alb.
func (a dataAlbAttributes) XffHeaderProcessingMode() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("xff_header_processing_mode"))
}

// ZoneId returns a reference to field zone_id of aws_alb.
func (a dataAlbAttributes) ZoneId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("zone_id"))
}

func (a dataAlbAttributes) AccessLogs() terra.ListValue[dataalb.AccessLogsAttributes] {
	return terra.ReferenceAsList[dataalb.AccessLogsAttributes](a.ref.Append("access_logs"))
}

func (a dataAlbAttributes) SubnetMapping() terra.SetValue[dataalb.SubnetMappingAttributes] {
	return terra.ReferenceAsSet[dataalb.SubnetMappingAttributes](a.ref.Append("subnet_mapping"))
}

func (a dataAlbAttributes) Timeouts() dataalb.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataalb.TimeoutsAttributes](a.ref.Append("timeouts"))
}