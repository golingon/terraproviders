// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataConnectInstance(name string, args DataConnectInstanceArgs) *DataConnectInstance {
	return &DataConnectInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataConnectInstance)(nil)

type DataConnectInstance struct {
	Name string
	Args DataConnectInstanceArgs
}

func (ci *DataConnectInstance) DataSource() string {
	return "aws_connect_instance"
}

func (ci *DataConnectInstance) LocalName() string {
	return ci.Name
}

func (ci *DataConnectInstance) Configuration() interface{} {
	return ci.Args
}

func (ci *DataConnectInstance) Attributes() dataConnectInstanceAttributes {
	return dataConnectInstanceAttributes{ref: terra.ReferenceDataResource(ci)}
}

type DataConnectInstanceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceAlias: string, optional
	InstanceAlias terra.StringValue `hcl:"instance_alias,attr"`
	// InstanceId: string, optional
	InstanceId terra.StringValue `hcl:"instance_id,attr"`
}
type dataConnectInstanceAttributes struct {
	ref terra.Reference
}

func (ci dataConnectInstanceAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(ci.ref.Append("arn"))
}

func (ci dataConnectInstanceAttributes) AutoResolveBestVoicesEnabled() terra.BoolValue {
	return terra.ReferenceBool(ci.ref.Append("auto_resolve_best_voices_enabled"))
}

func (ci dataConnectInstanceAttributes) ContactFlowLogsEnabled() terra.BoolValue {
	return terra.ReferenceBool(ci.ref.Append("contact_flow_logs_enabled"))
}

func (ci dataConnectInstanceAttributes) ContactLensEnabled() terra.BoolValue {
	return terra.ReferenceBool(ci.ref.Append("contact_lens_enabled"))
}

func (ci dataConnectInstanceAttributes) CreatedTime() terra.StringValue {
	return terra.ReferenceString(ci.ref.Append("created_time"))
}

func (ci dataConnectInstanceAttributes) EarlyMediaEnabled() terra.BoolValue {
	return terra.ReferenceBool(ci.ref.Append("early_media_enabled"))
}

func (ci dataConnectInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ci.ref.Append("id"))
}

func (ci dataConnectInstanceAttributes) IdentityManagementType() terra.StringValue {
	return terra.ReferenceString(ci.ref.Append("identity_management_type"))
}

func (ci dataConnectInstanceAttributes) InboundCallsEnabled() terra.BoolValue {
	return terra.ReferenceBool(ci.ref.Append("inbound_calls_enabled"))
}

func (ci dataConnectInstanceAttributes) InstanceAlias() terra.StringValue {
	return terra.ReferenceString(ci.ref.Append("instance_alias"))
}

func (ci dataConnectInstanceAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceString(ci.ref.Append("instance_id"))
}

func (ci dataConnectInstanceAttributes) MultiPartyConferenceEnabled() terra.BoolValue {
	return terra.ReferenceBool(ci.ref.Append("multi_party_conference_enabled"))
}

func (ci dataConnectInstanceAttributes) OutboundCallsEnabled() terra.BoolValue {
	return terra.ReferenceBool(ci.ref.Append("outbound_calls_enabled"))
}

func (ci dataConnectInstanceAttributes) ServiceRole() terra.StringValue {
	return terra.ReferenceString(ci.ref.Append("service_role"))
}

func (ci dataConnectInstanceAttributes) Status() terra.StringValue {
	return terra.ReferenceString(ci.ref.Append("status"))
}
