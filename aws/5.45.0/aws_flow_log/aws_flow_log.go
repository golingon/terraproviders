// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_flow_log

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_flow_log.
type Resource struct {
	Name      string
	Args      Args
	state     *awsFlowLogState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (afl *Resource) Type() string {
	return "aws_flow_log"
}

// LocalName returns the local name for [Resource].
func (afl *Resource) LocalName() string {
	return afl.Name
}

// Configuration returns the configuration (args) for [Resource].
func (afl *Resource) Configuration() interface{} {
	return afl.Args
}

// DependOn is used for other resources to depend on [Resource].
func (afl *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(afl)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (afl *Resource) Dependencies() terra.Dependencies {
	return afl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (afl *Resource) LifecycleManagement() *terra.Lifecycle {
	return afl.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (afl *Resource) Attributes() awsFlowLogAttributes {
	return awsFlowLogAttributes{ref: terra.ReferenceResource(afl)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (afl *Resource) ImportState(state io.Reader) error {
	afl.state = &awsFlowLogState{}
	if err := json.NewDecoder(state).Decode(afl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", afl.Type(), afl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (afl *Resource) State() (*awsFlowLogState, bool) {
	return afl.state, afl.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (afl *Resource) StateMust() *awsFlowLogState {
	if afl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", afl.Type(), afl.LocalName()))
	}
	return afl.state
}

// Args contains the configurations for aws_flow_log.
type Args struct {
	// DeliverCrossAccountRole: string, optional
	DeliverCrossAccountRole terra.StringValue `hcl:"deliver_cross_account_role,attr"`
	// EniId: string, optional
	EniId terra.StringValue `hcl:"eni_id,attr"`
	// IamRoleArn: string, optional
	IamRoleArn terra.StringValue `hcl:"iam_role_arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogDestination: string, optional
	LogDestination terra.StringValue `hcl:"log_destination,attr"`
	// LogDestinationType: string, optional
	LogDestinationType terra.StringValue `hcl:"log_destination_type,attr"`
	// LogFormat: string, optional
	LogFormat terra.StringValue `hcl:"log_format,attr"`
	// LogGroupName: string, optional
	LogGroupName terra.StringValue `hcl:"log_group_name,attr"`
	// MaxAggregationInterval: number, optional
	MaxAggregationInterval terra.NumberValue `hcl:"max_aggregation_interval,attr"`
	// SubnetId: string, optional
	SubnetId terra.StringValue `hcl:"subnet_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TrafficType: string, optional
	TrafficType terra.StringValue `hcl:"traffic_type,attr"`
	// TransitGatewayAttachmentId: string, optional
	TransitGatewayAttachmentId terra.StringValue `hcl:"transit_gateway_attachment_id,attr"`
	// TransitGatewayId: string, optional
	TransitGatewayId terra.StringValue `hcl:"transit_gateway_id,attr"`
	// VpcId: string, optional
	VpcId terra.StringValue `hcl:"vpc_id,attr"`
	// DestinationOptions: optional
	DestinationOptions *DestinationOptions `hcl:"destination_options,block"`
}

type awsFlowLogAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_flow_log.
func (afl awsFlowLogAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(afl.ref.Append("arn"))
}

// DeliverCrossAccountRole returns a reference to field deliver_cross_account_role of aws_flow_log.
func (afl awsFlowLogAttributes) DeliverCrossAccountRole() terra.StringValue {
	return terra.ReferenceAsString(afl.ref.Append("deliver_cross_account_role"))
}

// EniId returns a reference to field eni_id of aws_flow_log.
func (afl awsFlowLogAttributes) EniId() terra.StringValue {
	return terra.ReferenceAsString(afl.ref.Append("eni_id"))
}

// IamRoleArn returns a reference to field iam_role_arn of aws_flow_log.
func (afl awsFlowLogAttributes) IamRoleArn() terra.StringValue {
	return terra.ReferenceAsString(afl.ref.Append("iam_role_arn"))
}

// Id returns a reference to field id of aws_flow_log.
func (afl awsFlowLogAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(afl.ref.Append("id"))
}

// LogDestination returns a reference to field log_destination of aws_flow_log.
func (afl awsFlowLogAttributes) LogDestination() terra.StringValue {
	return terra.ReferenceAsString(afl.ref.Append("log_destination"))
}

// LogDestinationType returns a reference to field log_destination_type of aws_flow_log.
func (afl awsFlowLogAttributes) LogDestinationType() terra.StringValue {
	return terra.ReferenceAsString(afl.ref.Append("log_destination_type"))
}

// LogFormat returns a reference to field log_format of aws_flow_log.
func (afl awsFlowLogAttributes) LogFormat() terra.StringValue {
	return terra.ReferenceAsString(afl.ref.Append("log_format"))
}

// LogGroupName returns a reference to field log_group_name of aws_flow_log.
func (afl awsFlowLogAttributes) LogGroupName() terra.StringValue {
	return terra.ReferenceAsString(afl.ref.Append("log_group_name"))
}

// MaxAggregationInterval returns a reference to field max_aggregation_interval of aws_flow_log.
func (afl awsFlowLogAttributes) MaxAggregationInterval() terra.NumberValue {
	return terra.ReferenceAsNumber(afl.ref.Append("max_aggregation_interval"))
}

// SubnetId returns a reference to field subnet_id of aws_flow_log.
func (afl awsFlowLogAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(afl.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of aws_flow_log.
func (afl awsFlowLogAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](afl.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_flow_log.
func (afl awsFlowLogAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](afl.ref.Append("tags_all"))
}

// TrafficType returns a reference to field traffic_type of aws_flow_log.
func (afl awsFlowLogAttributes) TrafficType() terra.StringValue {
	return terra.ReferenceAsString(afl.ref.Append("traffic_type"))
}

// TransitGatewayAttachmentId returns a reference to field transit_gateway_attachment_id of aws_flow_log.
func (afl awsFlowLogAttributes) TransitGatewayAttachmentId() terra.StringValue {
	return terra.ReferenceAsString(afl.ref.Append("transit_gateway_attachment_id"))
}

// TransitGatewayId returns a reference to field transit_gateway_id of aws_flow_log.
func (afl awsFlowLogAttributes) TransitGatewayId() terra.StringValue {
	return terra.ReferenceAsString(afl.ref.Append("transit_gateway_id"))
}

// VpcId returns a reference to field vpc_id of aws_flow_log.
func (afl awsFlowLogAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(afl.ref.Append("vpc_id"))
}

func (afl awsFlowLogAttributes) DestinationOptions() terra.ListValue[DestinationOptionsAttributes] {
	return terra.ReferenceAsList[DestinationOptionsAttributes](afl.ref.Append("destination_options"))
}

type awsFlowLogState struct {
	Arn                        string                    `json:"arn"`
	DeliverCrossAccountRole    string                    `json:"deliver_cross_account_role"`
	EniId                      string                    `json:"eni_id"`
	IamRoleArn                 string                    `json:"iam_role_arn"`
	Id                         string                    `json:"id"`
	LogDestination             string                    `json:"log_destination"`
	LogDestinationType         string                    `json:"log_destination_type"`
	LogFormat                  string                    `json:"log_format"`
	LogGroupName               string                    `json:"log_group_name"`
	MaxAggregationInterval     float64                   `json:"max_aggregation_interval"`
	SubnetId                   string                    `json:"subnet_id"`
	Tags                       map[string]string         `json:"tags"`
	TagsAll                    map[string]string         `json:"tags_all"`
	TrafficType                string                    `json:"traffic_type"`
	TransitGatewayAttachmentId string                    `json:"transit_gateway_attachment_id"`
	TransitGatewayId           string                    `json:"transit_gateway_id"`
	VpcId                      string                    `json:"vpc_id"`
	DestinationOptions         []DestinationOptionsState `json:"destination_options"`
}
