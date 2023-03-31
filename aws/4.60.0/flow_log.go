// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	flowlog "github.com/golingon/terraproviders/aws/4.60.0/flowlog"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFlowLog creates a new instance of [FlowLog].
func NewFlowLog(name string, args FlowLogArgs) *FlowLog {
	return &FlowLog{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FlowLog)(nil)

// FlowLog represents the Terraform resource aws_flow_log.
type FlowLog struct {
	Name      string
	Args      FlowLogArgs
	state     *flowLogState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FlowLog].
func (fl *FlowLog) Type() string {
	return "aws_flow_log"
}

// LocalName returns the local name for [FlowLog].
func (fl *FlowLog) LocalName() string {
	return fl.Name
}

// Configuration returns the configuration (args) for [FlowLog].
func (fl *FlowLog) Configuration() interface{} {
	return fl.Args
}

// DependOn is used for other resources to depend on [FlowLog].
func (fl *FlowLog) DependOn() terra.Reference {
	return terra.ReferenceResource(fl)
}

// Dependencies returns the list of resources [FlowLog] depends_on.
func (fl *FlowLog) Dependencies() terra.Dependencies {
	return fl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FlowLog].
func (fl *FlowLog) LifecycleManagement() *terra.Lifecycle {
	return fl.Lifecycle
}

// Attributes returns the attributes for [FlowLog].
func (fl *FlowLog) Attributes() flowLogAttributes {
	return flowLogAttributes{ref: terra.ReferenceResource(fl)}
}

// ImportState imports the given attribute values into [FlowLog]'s state.
func (fl *FlowLog) ImportState(av io.Reader) error {
	fl.state = &flowLogState{}
	if err := json.NewDecoder(av).Decode(fl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fl.Type(), fl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FlowLog] has state.
func (fl *FlowLog) State() (*flowLogState, bool) {
	return fl.state, fl.state != nil
}

// StateMust returns the state for [FlowLog]. Panics if the state is nil.
func (fl *FlowLog) StateMust() *flowLogState {
	if fl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fl.Type(), fl.LocalName()))
	}
	return fl.state
}

// FlowLogArgs contains the configurations for aws_flow_log.
type FlowLogArgs struct {
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
	DestinationOptions *flowlog.DestinationOptions `hcl:"destination_options,block"`
}
type flowLogAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_flow_log.
func (fl flowLogAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(fl.ref.Append("arn"))
}

// DeliverCrossAccountRole returns a reference to field deliver_cross_account_role of aws_flow_log.
func (fl flowLogAttributes) DeliverCrossAccountRole() terra.StringValue {
	return terra.ReferenceAsString(fl.ref.Append("deliver_cross_account_role"))
}

// EniId returns a reference to field eni_id of aws_flow_log.
func (fl flowLogAttributes) EniId() terra.StringValue {
	return terra.ReferenceAsString(fl.ref.Append("eni_id"))
}

// IamRoleArn returns a reference to field iam_role_arn of aws_flow_log.
func (fl flowLogAttributes) IamRoleArn() terra.StringValue {
	return terra.ReferenceAsString(fl.ref.Append("iam_role_arn"))
}

// Id returns a reference to field id of aws_flow_log.
func (fl flowLogAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fl.ref.Append("id"))
}

// LogDestination returns a reference to field log_destination of aws_flow_log.
func (fl flowLogAttributes) LogDestination() terra.StringValue {
	return terra.ReferenceAsString(fl.ref.Append("log_destination"))
}

// LogDestinationType returns a reference to field log_destination_type of aws_flow_log.
func (fl flowLogAttributes) LogDestinationType() terra.StringValue {
	return terra.ReferenceAsString(fl.ref.Append("log_destination_type"))
}

// LogFormat returns a reference to field log_format of aws_flow_log.
func (fl flowLogAttributes) LogFormat() terra.StringValue {
	return terra.ReferenceAsString(fl.ref.Append("log_format"))
}

// LogGroupName returns a reference to field log_group_name of aws_flow_log.
func (fl flowLogAttributes) LogGroupName() terra.StringValue {
	return terra.ReferenceAsString(fl.ref.Append("log_group_name"))
}

// MaxAggregationInterval returns a reference to field max_aggregation_interval of aws_flow_log.
func (fl flowLogAttributes) MaxAggregationInterval() terra.NumberValue {
	return terra.ReferenceAsNumber(fl.ref.Append("max_aggregation_interval"))
}

// SubnetId returns a reference to field subnet_id of aws_flow_log.
func (fl flowLogAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(fl.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of aws_flow_log.
func (fl flowLogAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fl.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_flow_log.
func (fl flowLogAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fl.ref.Append("tags_all"))
}

// TrafficType returns a reference to field traffic_type of aws_flow_log.
func (fl flowLogAttributes) TrafficType() terra.StringValue {
	return terra.ReferenceAsString(fl.ref.Append("traffic_type"))
}

// TransitGatewayAttachmentId returns a reference to field transit_gateway_attachment_id of aws_flow_log.
func (fl flowLogAttributes) TransitGatewayAttachmentId() terra.StringValue {
	return terra.ReferenceAsString(fl.ref.Append("transit_gateway_attachment_id"))
}

// TransitGatewayId returns a reference to field transit_gateway_id of aws_flow_log.
func (fl flowLogAttributes) TransitGatewayId() terra.StringValue {
	return terra.ReferenceAsString(fl.ref.Append("transit_gateway_id"))
}

// VpcId returns a reference to field vpc_id of aws_flow_log.
func (fl flowLogAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(fl.ref.Append("vpc_id"))
}

func (fl flowLogAttributes) DestinationOptions() terra.ListValue[flowlog.DestinationOptionsAttributes] {
	return terra.ReferenceAsList[flowlog.DestinationOptionsAttributes](fl.ref.Append("destination_options"))
}

type flowLogState struct {
	Arn                        string                            `json:"arn"`
	DeliverCrossAccountRole    string                            `json:"deliver_cross_account_role"`
	EniId                      string                            `json:"eni_id"`
	IamRoleArn                 string                            `json:"iam_role_arn"`
	Id                         string                            `json:"id"`
	LogDestination             string                            `json:"log_destination"`
	LogDestinationType         string                            `json:"log_destination_type"`
	LogFormat                  string                            `json:"log_format"`
	LogGroupName               string                            `json:"log_group_name"`
	MaxAggregationInterval     float64                           `json:"max_aggregation_interval"`
	SubnetId                   string                            `json:"subnet_id"`
	Tags                       map[string]string                 `json:"tags"`
	TagsAll                    map[string]string                 `json:"tags_all"`
	TrafficType                string                            `json:"traffic_type"`
	TransitGatewayAttachmentId string                            `json:"transit_gateway_attachment_id"`
	TransitGatewayId           string                            `json:"transit_gateway_id"`
	VpcId                      string                            `json:"vpc_id"`
	DestinationOptions         []flowlog.DestinationOptionsState `json:"destination_options"`
}
