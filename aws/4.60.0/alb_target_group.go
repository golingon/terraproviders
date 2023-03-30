// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	albtargetgroup "github.com/golingon/terraproviders/aws/4.60.0/albtargetgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewAlbTargetGroup(name string, args AlbTargetGroupArgs) *AlbTargetGroup {
	return &AlbTargetGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AlbTargetGroup)(nil)

type AlbTargetGroup struct {
	Name  string
	Args  AlbTargetGroupArgs
	state *albTargetGroupState
}

func (atg *AlbTargetGroup) Type() string {
	return "aws_alb_target_group"
}

func (atg *AlbTargetGroup) LocalName() string {
	return atg.Name
}

func (atg *AlbTargetGroup) Configuration() interface{} {
	return atg.Args
}

func (atg *AlbTargetGroup) Attributes() albTargetGroupAttributes {
	return albTargetGroupAttributes{ref: terra.ReferenceResource(atg)}
}

func (atg *AlbTargetGroup) ImportState(av io.Reader) error {
	atg.state = &albTargetGroupState{}
	if err := json.NewDecoder(av).Decode(atg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", atg.Type(), atg.LocalName(), err)
	}
	return nil
}

func (atg *AlbTargetGroup) State() (*albTargetGroupState, bool) {
	return atg.state, atg.state != nil
}

func (atg *AlbTargetGroup) StateMust() *albTargetGroupState {
	if atg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", atg.Type(), atg.LocalName()))
	}
	return atg.state
}

func (atg *AlbTargetGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(atg)
}

type AlbTargetGroupArgs struct {
	// ConnectionTermination: bool, optional
	ConnectionTermination terra.BoolValue `hcl:"connection_termination,attr"`
	// DeregistrationDelay: string, optional
	DeregistrationDelay terra.StringValue `hcl:"deregistration_delay,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpAddressType: string, optional
	IpAddressType terra.StringValue `hcl:"ip_address_type,attr"`
	// LambdaMultiValueHeadersEnabled: bool, optional
	LambdaMultiValueHeadersEnabled terra.BoolValue `hcl:"lambda_multi_value_headers_enabled,attr"`
	// LoadBalancingAlgorithmType: string, optional
	LoadBalancingAlgorithmType terra.StringValue `hcl:"load_balancing_algorithm_type,attr"`
	// LoadBalancingCrossZoneEnabled: string, optional
	LoadBalancingCrossZoneEnabled terra.StringValue `hcl:"load_balancing_cross_zone_enabled,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// Port: number, optional
	Port terra.NumberValue `hcl:"port,attr"`
	// PreserveClientIp: string, optional
	PreserveClientIp terra.StringValue `hcl:"preserve_client_ip,attr"`
	// Protocol: string, optional
	Protocol terra.StringValue `hcl:"protocol,attr"`
	// ProtocolVersion: string, optional
	ProtocolVersion terra.StringValue `hcl:"protocol_version,attr"`
	// ProxyProtocolV2: bool, optional
	ProxyProtocolV2 terra.BoolValue `hcl:"proxy_protocol_v2,attr"`
	// SlowStart: number, optional
	SlowStart terra.NumberValue `hcl:"slow_start,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TargetType: string, optional
	TargetType terra.StringValue `hcl:"target_type,attr"`
	// VpcId: string, optional
	VpcId terra.StringValue `hcl:"vpc_id,attr"`
	// HealthCheck: optional
	HealthCheck *albtargetgroup.HealthCheck `hcl:"health_check,block"`
	// Stickiness: optional
	Stickiness *albtargetgroup.Stickiness `hcl:"stickiness,block"`
	// TargetFailover: min=0
	TargetFailover []albtargetgroup.TargetFailover `hcl:"target_failover,block" validate:"min=0"`
	// DependsOn contains resources that AlbTargetGroup depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type albTargetGroupAttributes struct {
	ref terra.Reference
}

func (atg albTargetGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(atg.ref.Append("arn"))
}

func (atg albTargetGroupAttributes) ArnSuffix() terra.StringValue {
	return terra.ReferenceString(atg.ref.Append("arn_suffix"))
}

func (atg albTargetGroupAttributes) ConnectionTermination() terra.BoolValue {
	return terra.ReferenceBool(atg.ref.Append("connection_termination"))
}

func (atg albTargetGroupAttributes) DeregistrationDelay() terra.StringValue {
	return terra.ReferenceString(atg.ref.Append("deregistration_delay"))
}

func (atg albTargetGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceString(atg.ref.Append("id"))
}

func (atg albTargetGroupAttributes) IpAddressType() terra.StringValue {
	return terra.ReferenceString(atg.ref.Append("ip_address_type"))
}

func (atg albTargetGroupAttributes) LambdaMultiValueHeadersEnabled() terra.BoolValue {
	return terra.ReferenceBool(atg.ref.Append("lambda_multi_value_headers_enabled"))
}

func (atg albTargetGroupAttributes) LoadBalancingAlgorithmType() terra.StringValue {
	return terra.ReferenceString(atg.ref.Append("load_balancing_algorithm_type"))
}

func (atg albTargetGroupAttributes) LoadBalancingCrossZoneEnabled() terra.StringValue {
	return terra.ReferenceString(atg.ref.Append("load_balancing_cross_zone_enabled"))
}

func (atg albTargetGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceString(atg.ref.Append("name"))
}

func (atg albTargetGroupAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceString(atg.ref.Append("name_prefix"))
}

func (atg albTargetGroupAttributes) Port() terra.NumberValue {
	return terra.ReferenceNumber(atg.ref.Append("port"))
}

func (atg albTargetGroupAttributes) PreserveClientIp() terra.StringValue {
	return terra.ReferenceString(atg.ref.Append("preserve_client_ip"))
}

func (atg albTargetGroupAttributes) Protocol() terra.StringValue {
	return terra.ReferenceString(atg.ref.Append("protocol"))
}

func (atg albTargetGroupAttributes) ProtocolVersion() terra.StringValue {
	return terra.ReferenceString(atg.ref.Append("protocol_version"))
}

func (atg albTargetGroupAttributes) ProxyProtocolV2() terra.BoolValue {
	return terra.ReferenceBool(atg.ref.Append("proxy_protocol_v2"))
}

func (atg albTargetGroupAttributes) SlowStart() terra.NumberValue {
	return terra.ReferenceNumber(atg.ref.Append("slow_start"))
}

func (atg albTargetGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](atg.ref.Append("tags"))
}

func (atg albTargetGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](atg.ref.Append("tags_all"))
}

func (atg albTargetGroupAttributes) TargetType() terra.StringValue {
	return terra.ReferenceString(atg.ref.Append("target_type"))
}

func (atg albTargetGroupAttributes) VpcId() terra.StringValue {
	return terra.ReferenceString(atg.ref.Append("vpc_id"))
}

func (atg albTargetGroupAttributes) HealthCheck() terra.ListValue[albtargetgroup.HealthCheckAttributes] {
	return terra.ReferenceList[albtargetgroup.HealthCheckAttributes](atg.ref.Append("health_check"))
}

func (atg albTargetGroupAttributes) Stickiness() terra.ListValue[albtargetgroup.StickinessAttributes] {
	return terra.ReferenceList[albtargetgroup.StickinessAttributes](atg.ref.Append("stickiness"))
}

func (atg albTargetGroupAttributes) TargetFailover() terra.ListValue[albtargetgroup.TargetFailoverAttributes] {
	return terra.ReferenceList[albtargetgroup.TargetFailoverAttributes](atg.ref.Append("target_failover"))
}

type albTargetGroupState struct {
	Arn                            string                               `json:"arn"`
	ArnSuffix                      string                               `json:"arn_suffix"`
	ConnectionTermination          bool                                 `json:"connection_termination"`
	DeregistrationDelay            string                               `json:"deregistration_delay"`
	Id                             string                               `json:"id"`
	IpAddressType                  string                               `json:"ip_address_type"`
	LambdaMultiValueHeadersEnabled bool                                 `json:"lambda_multi_value_headers_enabled"`
	LoadBalancingAlgorithmType     string                               `json:"load_balancing_algorithm_type"`
	LoadBalancingCrossZoneEnabled  string                               `json:"load_balancing_cross_zone_enabled"`
	Name                           string                               `json:"name"`
	NamePrefix                     string                               `json:"name_prefix"`
	Port                           float64                              `json:"port"`
	PreserveClientIp               string                               `json:"preserve_client_ip"`
	Protocol                       string                               `json:"protocol"`
	ProtocolVersion                string                               `json:"protocol_version"`
	ProxyProtocolV2                bool                                 `json:"proxy_protocol_v2"`
	SlowStart                      float64                              `json:"slow_start"`
	Tags                           map[string]string                    `json:"tags"`
	TagsAll                        map[string]string                    `json:"tags_all"`
	TargetType                     string                               `json:"target_type"`
	VpcId                          string                               `json:"vpc_id"`
	HealthCheck                    []albtargetgroup.HealthCheckState    `json:"health_check"`
	Stickiness                     []albtargetgroup.StickinessState     `json:"stickiness"`
	TargetFailover                 []albtargetgroup.TargetFailoverState `json:"target_failover"`
}
