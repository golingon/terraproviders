// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datalbtargetgroup "github.com/golingon/terraproviders/aws/4.60.0/datalbtargetgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataLbTargetGroup(name string, args DataLbTargetGroupArgs) *DataLbTargetGroup {
	return &DataLbTargetGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLbTargetGroup)(nil)

type DataLbTargetGroup struct {
	Name string
	Args DataLbTargetGroupArgs
}

func (ltg *DataLbTargetGroup) DataSource() string {
	return "aws_lb_target_group"
}

func (ltg *DataLbTargetGroup) LocalName() string {
	return ltg.Name
}

func (ltg *DataLbTargetGroup) Configuration() interface{} {
	return ltg.Args
}

func (ltg *DataLbTargetGroup) Attributes() dataLbTargetGroupAttributes {
	return dataLbTargetGroupAttributes{ref: terra.ReferenceDataResource(ltg)}
}

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

func (ltg dataLbTargetGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(ltg.ref.Append("arn"))
}

func (ltg dataLbTargetGroupAttributes) ArnSuffix() terra.StringValue {
	return terra.ReferenceString(ltg.ref.Append("arn_suffix"))
}

func (ltg dataLbTargetGroupAttributes) ConnectionTermination() terra.BoolValue {
	return terra.ReferenceBool(ltg.ref.Append("connection_termination"))
}

func (ltg dataLbTargetGroupAttributes) DeregistrationDelay() terra.NumberValue {
	return terra.ReferenceNumber(ltg.ref.Append("deregistration_delay"))
}

func (ltg dataLbTargetGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ltg.ref.Append("id"))
}

func (ltg dataLbTargetGroupAttributes) LambdaMultiValueHeadersEnabled() terra.BoolValue {
	return terra.ReferenceBool(ltg.ref.Append("lambda_multi_value_headers_enabled"))
}

func (ltg dataLbTargetGroupAttributes) LoadBalancingAlgorithmType() terra.StringValue {
	return terra.ReferenceString(ltg.ref.Append("load_balancing_algorithm_type"))
}

func (ltg dataLbTargetGroupAttributes) LoadBalancingCrossZoneEnabled() terra.StringValue {
	return terra.ReferenceString(ltg.ref.Append("load_balancing_cross_zone_enabled"))
}

func (ltg dataLbTargetGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ltg.ref.Append("name"))
}

func (ltg dataLbTargetGroupAttributes) Port() terra.NumberValue {
	return terra.ReferenceNumber(ltg.ref.Append("port"))
}

func (ltg dataLbTargetGroupAttributes) PreserveClientIp() terra.StringValue {
	return terra.ReferenceString(ltg.ref.Append("preserve_client_ip"))
}

func (ltg dataLbTargetGroupAttributes) Protocol() terra.StringValue {
	return terra.ReferenceString(ltg.ref.Append("protocol"))
}

func (ltg dataLbTargetGroupAttributes) ProtocolVersion() terra.StringValue {
	return terra.ReferenceString(ltg.ref.Append("protocol_version"))
}

func (ltg dataLbTargetGroupAttributes) ProxyProtocolV2() terra.BoolValue {
	return terra.ReferenceBool(ltg.ref.Append("proxy_protocol_v2"))
}

func (ltg dataLbTargetGroupAttributes) SlowStart() terra.NumberValue {
	return terra.ReferenceNumber(ltg.ref.Append("slow_start"))
}

func (ltg dataLbTargetGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ltg.ref.Append("tags"))
}

func (ltg dataLbTargetGroupAttributes) TargetType() terra.StringValue {
	return terra.ReferenceString(ltg.ref.Append("target_type"))
}

func (ltg dataLbTargetGroupAttributes) VpcId() terra.StringValue {
	return terra.ReferenceString(ltg.ref.Append("vpc_id"))
}

func (ltg dataLbTargetGroupAttributes) HealthCheck() terra.ListValue[datalbtargetgroup.HealthCheckAttributes] {
	return terra.ReferenceList[datalbtargetgroup.HealthCheckAttributes](ltg.ref.Append("health_check"))
}

func (ltg dataLbTargetGroupAttributes) Stickiness() terra.ListValue[datalbtargetgroup.StickinessAttributes] {
	return terra.ReferenceList[datalbtargetgroup.StickinessAttributes](ltg.ref.Append("stickiness"))
}

func (ltg dataLbTargetGroupAttributes) Timeouts() datalbtargetgroup.TimeoutsAttributes {
	return terra.ReferenceSingle[datalbtargetgroup.TimeoutsAttributes](ltg.ref.Append("timeouts"))
}
