// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	lbtargetgroup "github.com/golingon/terraproviders/aws/5.6.2/lbtargetgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLbTargetGroup creates a new instance of [LbTargetGroup].
func NewLbTargetGroup(name string, args LbTargetGroupArgs) *LbTargetGroup {
	return &LbTargetGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LbTargetGroup)(nil)

// LbTargetGroup represents the Terraform resource aws_lb_target_group.
type LbTargetGroup struct {
	Name      string
	Args      LbTargetGroupArgs
	state     *lbTargetGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LbTargetGroup].
func (ltg *LbTargetGroup) Type() string {
	return "aws_lb_target_group"
}

// LocalName returns the local name for [LbTargetGroup].
func (ltg *LbTargetGroup) LocalName() string {
	return ltg.Name
}

// Configuration returns the configuration (args) for [LbTargetGroup].
func (ltg *LbTargetGroup) Configuration() interface{} {
	return ltg.Args
}

// DependOn is used for other resources to depend on [LbTargetGroup].
func (ltg *LbTargetGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(ltg)
}

// Dependencies returns the list of resources [LbTargetGroup] depends_on.
func (ltg *LbTargetGroup) Dependencies() terra.Dependencies {
	return ltg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LbTargetGroup].
func (ltg *LbTargetGroup) LifecycleManagement() *terra.Lifecycle {
	return ltg.Lifecycle
}

// Attributes returns the attributes for [LbTargetGroup].
func (ltg *LbTargetGroup) Attributes() lbTargetGroupAttributes {
	return lbTargetGroupAttributes{ref: terra.ReferenceResource(ltg)}
}

// ImportState imports the given attribute values into [LbTargetGroup]'s state.
func (ltg *LbTargetGroup) ImportState(av io.Reader) error {
	ltg.state = &lbTargetGroupState{}
	if err := json.NewDecoder(av).Decode(ltg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ltg.Type(), ltg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LbTargetGroup] has state.
func (ltg *LbTargetGroup) State() (*lbTargetGroupState, bool) {
	return ltg.state, ltg.state != nil
}

// StateMust returns the state for [LbTargetGroup]. Panics if the state is nil.
func (ltg *LbTargetGroup) StateMust() *lbTargetGroupState {
	if ltg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ltg.Type(), ltg.LocalName()))
	}
	return ltg.state
}

// LbTargetGroupArgs contains the configurations for aws_lb_target_group.
type LbTargetGroupArgs struct {
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
	HealthCheck *lbtargetgroup.HealthCheck `hcl:"health_check,block"`
	// Stickiness: optional
	Stickiness *lbtargetgroup.Stickiness `hcl:"stickiness,block"`
	// TargetFailover: min=0
	TargetFailover []lbtargetgroup.TargetFailover `hcl:"target_failover,block" validate:"min=0"`
}
type lbTargetGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_lb_target_group.
func (ltg lbTargetGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("arn"))
}

// ArnSuffix returns a reference to field arn_suffix of aws_lb_target_group.
func (ltg lbTargetGroupAttributes) ArnSuffix() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("arn_suffix"))
}

// ConnectionTermination returns a reference to field connection_termination of aws_lb_target_group.
func (ltg lbTargetGroupAttributes) ConnectionTermination() terra.BoolValue {
	return terra.ReferenceAsBool(ltg.ref.Append("connection_termination"))
}

// DeregistrationDelay returns a reference to field deregistration_delay of aws_lb_target_group.
func (ltg lbTargetGroupAttributes) DeregistrationDelay() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("deregistration_delay"))
}

// Id returns a reference to field id of aws_lb_target_group.
func (ltg lbTargetGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("id"))
}

// IpAddressType returns a reference to field ip_address_type of aws_lb_target_group.
func (ltg lbTargetGroupAttributes) IpAddressType() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("ip_address_type"))
}

// LambdaMultiValueHeadersEnabled returns a reference to field lambda_multi_value_headers_enabled of aws_lb_target_group.
func (ltg lbTargetGroupAttributes) LambdaMultiValueHeadersEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ltg.ref.Append("lambda_multi_value_headers_enabled"))
}

// LoadBalancingAlgorithmType returns a reference to field load_balancing_algorithm_type of aws_lb_target_group.
func (ltg lbTargetGroupAttributes) LoadBalancingAlgorithmType() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("load_balancing_algorithm_type"))
}

// LoadBalancingCrossZoneEnabled returns a reference to field load_balancing_cross_zone_enabled of aws_lb_target_group.
func (ltg lbTargetGroupAttributes) LoadBalancingCrossZoneEnabled() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("load_balancing_cross_zone_enabled"))
}

// Name returns a reference to field name of aws_lb_target_group.
func (ltg lbTargetGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_lb_target_group.
func (ltg lbTargetGroupAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("name_prefix"))
}

// Port returns a reference to field port of aws_lb_target_group.
func (ltg lbTargetGroupAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(ltg.ref.Append("port"))
}

// PreserveClientIp returns a reference to field preserve_client_ip of aws_lb_target_group.
func (ltg lbTargetGroupAttributes) PreserveClientIp() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("preserve_client_ip"))
}

// Protocol returns a reference to field protocol of aws_lb_target_group.
func (ltg lbTargetGroupAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("protocol"))
}

// ProtocolVersion returns a reference to field protocol_version of aws_lb_target_group.
func (ltg lbTargetGroupAttributes) ProtocolVersion() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("protocol_version"))
}

// ProxyProtocolV2 returns a reference to field proxy_protocol_v2 of aws_lb_target_group.
func (ltg lbTargetGroupAttributes) ProxyProtocolV2() terra.BoolValue {
	return terra.ReferenceAsBool(ltg.ref.Append("proxy_protocol_v2"))
}

// SlowStart returns a reference to field slow_start of aws_lb_target_group.
func (ltg lbTargetGroupAttributes) SlowStart() terra.NumberValue {
	return terra.ReferenceAsNumber(ltg.ref.Append("slow_start"))
}

// Tags returns a reference to field tags of aws_lb_target_group.
func (ltg lbTargetGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ltg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_lb_target_group.
func (ltg lbTargetGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ltg.ref.Append("tags_all"))
}

// TargetType returns a reference to field target_type of aws_lb_target_group.
func (ltg lbTargetGroupAttributes) TargetType() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("target_type"))
}

// VpcId returns a reference to field vpc_id of aws_lb_target_group.
func (ltg lbTargetGroupAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(ltg.ref.Append("vpc_id"))
}

func (ltg lbTargetGroupAttributes) HealthCheck() terra.ListValue[lbtargetgroup.HealthCheckAttributes] {
	return terra.ReferenceAsList[lbtargetgroup.HealthCheckAttributes](ltg.ref.Append("health_check"))
}

func (ltg lbTargetGroupAttributes) Stickiness() terra.ListValue[lbtargetgroup.StickinessAttributes] {
	return terra.ReferenceAsList[lbtargetgroup.StickinessAttributes](ltg.ref.Append("stickiness"))
}

func (ltg lbTargetGroupAttributes) TargetFailover() terra.ListValue[lbtargetgroup.TargetFailoverAttributes] {
	return terra.ReferenceAsList[lbtargetgroup.TargetFailoverAttributes](ltg.ref.Append("target_failover"))
}

type lbTargetGroupState struct {
	Arn                            string                              `json:"arn"`
	ArnSuffix                      string                              `json:"arn_suffix"`
	ConnectionTermination          bool                                `json:"connection_termination"`
	DeregistrationDelay            string                              `json:"deregistration_delay"`
	Id                             string                              `json:"id"`
	IpAddressType                  string                              `json:"ip_address_type"`
	LambdaMultiValueHeadersEnabled bool                                `json:"lambda_multi_value_headers_enabled"`
	LoadBalancingAlgorithmType     string                              `json:"load_balancing_algorithm_type"`
	LoadBalancingCrossZoneEnabled  string                              `json:"load_balancing_cross_zone_enabled"`
	Name                           string                              `json:"name"`
	NamePrefix                     string                              `json:"name_prefix"`
	Port                           float64                             `json:"port"`
	PreserveClientIp               string                              `json:"preserve_client_ip"`
	Protocol                       string                              `json:"protocol"`
	ProtocolVersion                string                              `json:"protocol_version"`
	ProxyProtocolV2                bool                                `json:"proxy_protocol_v2"`
	SlowStart                      float64                             `json:"slow_start"`
	Tags                           map[string]string                   `json:"tags"`
	TagsAll                        map[string]string                   `json:"tags_all"`
	TargetType                     string                              `json:"target_type"`
	VpcId                          string                              `json:"vpc_id"`
	HealthCheck                    []lbtargetgroup.HealthCheckState    `json:"health_check"`
	Stickiness                     []lbtargetgroup.StickinessState     `json:"stickiness"`
	TargetFailover                 []lbtargetgroup.TargetFailoverState `json:"target_failover"`
}