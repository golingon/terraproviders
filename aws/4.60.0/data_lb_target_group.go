// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datalbtargetgroup "github.com/golingon/terraproviders/aws/4.60.0/datalbtargetgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataLbTargetGroup creates a new instance of [DataLbTargetGroup].
func NewDataLbTargetGroup(name string, args DataLbTargetGroupArgs) *DataLbTargetGroup {
	return &DataLbTargetGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLbTargetGroup)(nil)

// DataLbTargetGroup represents the Terraform data resource aws_lb_target_group.
type DataLbTargetGroup struct {
	Name string
	Args DataLbTargetGroupArgs
}

// DataSource returns the Terraform object type for [DataLbTargetGroup].
func (ltg *DataLbTargetGroup) DataSource() string {
	return "aws_lb_target_group"
}

// LocalName returns the local name for [DataLbTargetGroup].
func (ltg *DataLbTargetGroup) LocalName() string {
	return ltg.Name
}

// Configuration returns the configuration (args) for [DataLbTargetGroup].
func (ltg *DataLbTargetGroup) Configuration() interface{} {
	return ltg.Args
}

// Attributes returns the attributes for [DataLbTargetGroup].
func (ltg *DataLbTargetGroup) Attributes() dataLbTargetGroupAttributes {
	return dataLbTargetGroupAttributes{ref: terra.ReferenceDataResource(ltg)}
}

// DataLbTargetGroupArgs contains the configurations for aws_lb_target_group.
type DataLbTargetGroupArgs struct {
	// Arn: string, optional
	Arn terra.StringValue `hcl:"arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// HealthCheck: min=0
	HealthCheck []datalbtargetgroup.HealthCheck `hcl:"health_check,block" validate:"min=0"`
	// Stickiness: min=0
	Stickiness []datalbtargetgroup.Stickiness `hcl:"stickiness,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datalbtargetgroup.Timeouts `hcl:"timeouts,block"`
}
type dataLbTargetGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_lb_target_group.
func (ltg dataLbTargetGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("arn"))
}

// ArnSuffix returns a reference to field arn_suffix of aws_lb_target_group.
func (ltg dataLbTargetGroupAttributes) ArnSuffix() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("arn_suffix"))
}

// ConnectionTermination returns a reference to field connection_termination of aws_lb_target_group.
func (ltg dataLbTargetGroupAttributes) ConnectionTermination() terra.BoolValue {
	return terra.ReferenceAsBool(ltg.ref.Append("connection_termination"))
}

// DeregistrationDelay returns a reference to field deregistration_delay of aws_lb_target_group.
func (ltg dataLbTargetGroupAttributes) DeregistrationDelay() terra.NumberValue {
	return terra.ReferenceAsNumber(ltg.ref.Append("deregistration_delay"))
}

// Id returns a reference to field id of aws_lb_target_group.
func (ltg dataLbTargetGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("id"))
}

// LambdaMultiValueHeadersEnabled returns a reference to field lambda_multi_value_headers_enabled of aws_lb_target_group.
func (ltg dataLbTargetGroupAttributes) LambdaMultiValueHeadersEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ltg.ref.Append("lambda_multi_value_headers_enabled"))
}

// LoadBalancingAlgorithmType returns a reference to field load_balancing_algorithm_type of aws_lb_target_group.
func (ltg dataLbTargetGroupAttributes) LoadBalancingAlgorithmType() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("load_balancing_algorithm_type"))
}

// LoadBalancingCrossZoneEnabled returns a reference to field load_balancing_cross_zone_enabled of aws_lb_target_group.
func (ltg dataLbTargetGroupAttributes) LoadBalancingCrossZoneEnabled() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("load_balancing_cross_zone_enabled"))
}

// Name returns a reference to field name of aws_lb_target_group.
func (ltg dataLbTargetGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("name"))
}

// Port returns a reference to field port of aws_lb_target_group.
func (ltg dataLbTargetGroupAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(ltg.ref.Append("port"))
}

// PreserveClientIp returns a reference to field preserve_client_ip of aws_lb_target_group.
func (ltg dataLbTargetGroupAttributes) PreserveClientIp() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("preserve_client_ip"))
}

// Protocol returns a reference to field protocol of aws_lb_target_group.
func (ltg dataLbTargetGroupAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("protocol"))
}

// ProtocolVersion returns a reference to field protocol_version of aws_lb_target_group.
func (ltg dataLbTargetGroupAttributes) ProtocolVersion() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("protocol_version"))
}

// ProxyProtocolV2 returns a reference to field proxy_protocol_v2 of aws_lb_target_group.
func (ltg dataLbTargetGroupAttributes) ProxyProtocolV2() terra.BoolValue {
	return terra.ReferenceAsBool(ltg.ref.Append("proxy_protocol_v2"))
}

// SlowStart returns a reference to field slow_start of aws_lb_target_group.
func (ltg dataLbTargetGroupAttributes) SlowStart() terra.NumberValue {
	return terra.ReferenceAsNumber(ltg.ref.Append("slow_start"))
}

// Tags returns a reference to field tags of aws_lb_target_group.
func (ltg dataLbTargetGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ltg.ref.Append("tags"))
}

// TargetType returns a reference to field target_type of aws_lb_target_group.
func (ltg dataLbTargetGroupAttributes) TargetType() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("target_type"))
}

// VpcId returns a reference to field vpc_id of aws_lb_target_group.
func (ltg dataLbTargetGroupAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("vpc_id"))
}

func (ltg dataLbTargetGroupAttributes) HealthCheck() terra.ListValue[datalbtargetgroup.HealthCheckAttributes] {
	return terra.ReferenceAsList[datalbtargetgroup.HealthCheckAttributes](ltg.ref.Append("health_check"))
}

func (ltg dataLbTargetGroupAttributes) Stickiness() terra.ListValue[datalbtargetgroup.StickinessAttributes] {
	return terra.ReferenceAsList[datalbtargetgroup.StickinessAttributes](ltg.ref.Append("stickiness"))
}

func (ltg dataLbTargetGroupAttributes) Timeouts() datalbtargetgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datalbtargetgroup.TimeoutsAttributes](ltg.ref.Append("timeouts"))
}