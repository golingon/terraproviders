// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	flowlog "github.com/golingon/terraproviders/aws/4.60.0/flowlog"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewFlowLog(name string, args FlowLogArgs) *FlowLog {
	return &FlowLog{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FlowLog)(nil)

type FlowLog struct {
	Name  string
	Args  FlowLogArgs
	state *flowLogState
}

func (fl *FlowLog) Type() string {
	return "aws_flow_log"
}

func (fl *FlowLog) LocalName() string {
	return fl.Name
}

func (fl *FlowLog) Configuration() interface{} {
	return fl.Args
}

func (fl *FlowLog) Attributes() flowLogAttributes {
	return flowLogAttributes{ref: terra.ReferenceResource(fl)}
}

func (fl *FlowLog) ImportState(av io.Reader) error {
	fl.state = &flowLogState{}
	if err := json.NewDecoder(av).Decode(fl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fl.Type(), fl.LocalName(), err)
	}
	return nil
}

func (fl *FlowLog) State() (*flowLogState, bool) {
	return fl.state, fl.state != nil
}

func (fl *FlowLog) StateMust() *flowLogState {
	if fl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fl.Type(), fl.LocalName()))
	}
	return fl.state
}

func (fl *FlowLog) DependOn() terra.Reference {
	return terra.ReferenceResource(fl)
}

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
	// DependsOn contains resources that FlowLog depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type flowLogAttributes struct {
	ref terra.Reference
}

func (fl flowLogAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(fl.ref.Append("arn"))
}

func (fl flowLogAttributes) DeliverCrossAccountRole() terra.StringValue {
	return terra.ReferenceString(fl.ref.Append("deliver_cross_account_role"))
}

func (fl flowLogAttributes) EniId() terra.StringValue {
	return terra.ReferenceString(fl.ref.Append("eni_id"))
}

func (fl flowLogAttributes) IamRoleArn() terra.StringValue {
	return terra.ReferenceString(fl.ref.Append("iam_role_arn"))
}

func (fl flowLogAttributes) Id() terra.StringValue {
	return terra.ReferenceString(fl.ref.Append("id"))
}

func (fl flowLogAttributes) LogDestination() terra.StringValue {
	return terra.ReferenceString(fl.ref.Append("log_destination"))
}

func (fl flowLogAttributes) LogDestinationType() terra.StringValue {
	return terra.ReferenceString(fl.ref.Append("log_destination_type"))
}

func (fl flowLogAttributes) LogFormat() terra.StringValue {
	return terra.ReferenceString(fl.ref.Append("log_format"))
}

func (fl flowLogAttributes) LogGroupName() terra.StringValue {
	return terra.ReferenceString(fl.ref.Append("log_group_name"))
}

func (fl flowLogAttributes) MaxAggregationInterval() terra.NumberValue {
	return terra.ReferenceNumber(fl.ref.Append("max_aggregation_interval"))
}

func (fl flowLogAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceString(fl.ref.Append("subnet_id"))
}

func (fl flowLogAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](fl.ref.Append("tags"))
}

func (fl flowLogAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](fl.ref.Append("tags_all"))
}

func (fl flowLogAttributes) TrafficType() terra.StringValue {
	return terra.ReferenceString(fl.ref.Append("traffic_type"))
}

func (fl flowLogAttributes) TransitGatewayAttachmentId() terra.StringValue {
	return terra.ReferenceString(fl.ref.Append("transit_gateway_attachment_id"))
}

func (fl flowLogAttributes) TransitGatewayId() terra.StringValue {
	return terra.ReferenceString(fl.ref.Append("transit_gateway_id"))
}

func (fl flowLogAttributes) VpcId() terra.StringValue {
	return terra.ReferenceString(fl.ref.Append("vpc_id"))
}

func (fl flowLogAttributes) DestinationOptions() terra.ListValue[flowlog.DestinationOptionsAttributes] {
	return terra.ReferenceList[flowlog.DestinationOptionsAttributes](fl.ref.Append("destination_options"))
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
