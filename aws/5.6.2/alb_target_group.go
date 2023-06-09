// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	albtargetgroup "github.com/golingon/terraproviders/aws/5.6.2/albtargetgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAlbTargetGroup creates a new instance of [AlbTargetGroup].
func NewAlbTargetGroup(name string, args AlbTargetGroupArgs) *AlbTargetGroup {
	return &AlbTargetGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AlbTargetGroup)(nil)

// AlbTargetGroup represents the Terraform resource aws_alb_target_group.
type AlbTargetGroup struct {
	Name      string
	Args      AlbTargetGroupArgs
	state     *albTargetGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AlbTargetGroup].
func (atg *AlbTargetGroup) Type() string {
	return "aws_alb_target_group"
}

// LocalName returns the local name for [AlbTargetGroup].
func (atg *AlbTargetGroup) LocalName() string {
	return atg.Name
}

// Configuration returns the configuration (args) for [AlbTargetGroup].
func (atg *AlbTargetGroup) Configuration() interface{} {
	return atg.Args
}

// DependOn is used for other resources to depend on [AlbTargetGroup].
func (atg *AlbTargetGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(atg)
}

// Dependencies returns the list of resources [AlbTargetGroup] depends_on.
func (atg *AlbTargetGroup) Dependencies() terra.Dependencies {
	return atg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AlbTargetGroup].
func (atg *AlbTargetGroup) LifecycleManagement() *terra.Lifecycle {
	return atg.Lifecycle
}

// Attributes returns the attributes for [AlbTargetGroup].
func (atg *AlbTargetGroup) Attributes() albTargetGroupAttributes {
	return albTargetGroupAttributes{ref: terra.ReferenceResource(atg)}
}

// ImportState imports the given attribute values into [AlbTargetGroup]'s state.
func (atg *AlbTargetGroup) ImportState(av io.Reader) error {
	atg.state = &albTargetGroupState{}
	if err := json.NewDecoder(av).Decode(atg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", atg.Type(), atg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AlbTargetGroup] has state.
func (atg *AlbTargetGroup) State() (*albTargetGroupState, bool) {
	return atg.state, atg.state != nil
}

// StateMust returns the state for [AlbTargetGroup]. Panics if the state is nil.
func (atg *AlbTargetGroup) StateMust() *albTargetGroupState {
	if atg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", atg.Type(), atg.LocalName()))
	}
	return atg.state
}

// AlbTargetGroupArgs contains the configurations for aws_alb_target_group.
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
}
type albTargetGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_alb_target_group.
func (atg albTargetGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(atg.ref.Append("arn"))
}

// ArnSuffix returns a reference to field arn_suffix of aws_alb_target_group.
func (atg albTargetGroupAttributes) ArnSuffix() terra.StringValue {
	return terra.ReferenceAsString(atg.ref.Append("arn_suffix"))
}

// ConnectionTermination returns a reference to field connection_termination of aws_alb_target_group.
func (atg albTargetGroupAttributes) ConnectionTermination() terra.BoolValue {
	return terra.ReferenceAsBool(atg.ref.Append("connection_termination"))
}

// DeregistrationDelay returns a reference to field deregistration_delay of aws_alb_target_group.
func (atg albTargetGroupAttributes) DeregistrationDelay() terra.StringValue {
	return terra.ReferenceAsString(atg.ref.Append("deregistration_delay"))
}

// Id returns a reference to field id of aws_alb_target_group.
func (atg albTargetGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(atg.ref.Append("id"))
}

// IpAddressType returns a reference to field ip_address_type of aws_alb_target_group.
func (atg albTargetGroupAttributes) IpAddressType() terra.StringValue {
	return terra.ReferenceAsString(atg.ref.Append("ip_address_type"))
}

// LambdaMultiValueHeadersEnabled returns a reference to field lambda_multi_value_headers_enabled of aws_alb_target_group.
func (atg albTargetGroupAttributes) LambdaMultiValueHeadersEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(atg.ref.Append("lambda_multi_value_headers_enabled"))
}

// LoadBalancingAlgorithmType returns a reference to field load_balancing_algorithm_type of aws_alb_target_group.
func (atg albTargetGroupAttributes) LoadBalancingAlgorithmType() terra.StringValue {
	return terra.ReferenceAsString(atg.ref.Append("load_balancing_algorithm_type"))
}

// LoadBalancingCrossZoneEnabled returns a reference to field load_balancing_cross_zone_enabled of aws_alb_target_group.
func (atg albTargetGroupAttributes) LoadBalancingCrossZoneEnabled() terra.StringValue {
	return terra.ReferenceAsString(atg.ref.Append("load_balancing_cross_zone_enabled"))
}

// Name returns a reference to field name of aws_alb_target_group.
func (atg albTargetGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(atg.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_alb_target_group.
func (atg albTargetGroupAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(atg.ref.Append("name_prefix"))
}

// Port returns a reference to field port of aws_alb_target_group.
func (atg albTargetGroupAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(atg.ref.Append("port"))
}

// PreserveClientIp returns a reference to field preserve_client_ip of aws_alb_target_group.
func (atg albTargetGroupAttributes) PreserveClientIp() terra.StringValue {
	return terra.ReferenceAsString(atg.ref.Append("preserve_client_ip"))
}

// Protocol returns a reference to field protocol of aws_alb_target_group.
func (atg albTargetGroupAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(atg.ref.Append("protocol"))
}

// ProtocolVersion returns a reference to field protocol_version of aws_alb_target_group.
func (atg albTargetGroupAttributes) ProtocolVersion() terra.StringValue {
	return terra.ReferenceAsString(atg.ref.Append("protocol_version"))
}

// ProxyProtocolV2 returns a reference to field proxy_protocol_v2 of aws_alb_target_group.
func (atg albTargetGroupAttributes) ProxyProtocolV2() terra.BoolValue {
	return terra.ReferenceAsBool(atg.ref.Append("proxy_protocol_v2"))
}

// SlowStart returns a reference to field slow_start of aws_alb_target_group.
func (atg albTargetGroupAttributes) SlowStart() terra.NumberValue {
	return terra.ReferenceAsNumber(atg.ref.Append("slow_start"))
}

// Tags returns a reference to field tags of aws_alb_target_group.
func (atg albTargetGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](atg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_alb_target_group.
func (atg albTargetGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](atg.ref.Append("tags_all"))
}

// TargetType returns a reference to field target_type of aws_alb_target_group.
func (atg albTargetGroupAttributes) TargetType() terra.StringValue {
	return terra.ReferenceAsString(atg.ref.Append("target_type"))
}

// VpcId returns a reference to field vpc_id of aws_alb_target_group.
func (atg albTargetGroupAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(atg.ref.Append("vpc_id"))
}

func (atg albTargetGroupAttributes) HealthCheck() terra.ListValue[albtargetgroup.HealthCheckAttributes] {
	return terra.ReferenceAsList[albtargetgroup.HealthCheckAttributes](atg.ref.Append("health_check"))
}

func (atg albTargetGroupAttributes) Stickiness() terra.ListValue[albtargetgroup.StickinessAttributes] {
	return terra.ReferenceAsList[albtargetgroup.StickinessAttributes](atg.ref.Append("stickiness"))
}

func (atg albTargetGroupAttributes) TargetFailover() terra.ListValue[albtargetgroup.TargetFailoverAttributes] {
	return terra.ReferenceAsList[albtargetgroup.TargetFailoverAttributes](atg.ref.Append("target_failover"))
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
