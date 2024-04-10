// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	connectinstance "github.com/golingon/terraproviders/aws/5.44.0/connectinstance"
	"io"
)

// NewConnectInstance creates a new instance of [ConnectInstance].
func NewConnectInstance(name string, args ConnectInstanceArgs) *ConnectInstance {
	return &ConnectInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConnectInstance)(nil)

// ConnectInstance represents the Terraform resource aws_connect_instance.
type ConnectInstance struct {
	Name      string
	Args      ConnectInstanceArgs
	state     *connectInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConnectInstance].
func (ci *ConnectInstance) Type() string {
	return "aws_connect_instance"
}

// LocalName returns the local name for [ConnectInstance].
func (ci *ConnectInstance) LocalName() string {
	return ci.Name
}

// Configuration returns the configuration (args) for [ConnectInstance].
func (ci *ConnectInstance) Configuration() interface{} {
	return ci.Args
}

// DependOn is used for other resources to depend on [ConnectInstance].
func (ci *ConnectInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(ci)
}

// Dependencies returns the list of resources [ConnectInstance] depends_on.
func (ci *ConnectInstance) Dependencies() terra.Dependencies {
	return ci.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConnectInstance].
func (ci *ConnectInstance) LifecycleManagement() *terra.Lifecycle {
	return ci.Lifecycle
}

// Attributes returns the attributes for [ConnectInstance].
func (ci *ConnectInstance) Attributes() connectInstanceAttributes {
	return connectInstanceAttributes{ref: terra.ReferenceResource(ci)}
}

// ImportState imports the given attribute values into [ConnectInstance]'s state.
func (ci *ConnectInstance) ImportState(av io.Reader) error {
	ci.state = &connectInstanceState{}
	if err := json.NewDecoder(av).Decode(ci.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ci.Type(), ci.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConnectInstance] has state.
func (ci *ConnectInstance) State() (*connectInstanceState, bool) {
	return ci.state, ci.state != nil
}

// StateMust returns the state for [ConnectInstance]. Panics if the state is nil.
func (ci *ConnectInstance) StateMust() *connectInstanceState {
	if ci.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ci.Type(), ci.LocalName()))
	}
	return ci.state
}

// ConnectInstanceArgs contains the configurations for aws_connect_instance.
type ConnectInstanceArgs struct {
	// AutoResolveBestVoicesEnabled: bool, optional
	AutoResolveBestVoicesEnabled terra.BoolValue `hcl:"auto_resolve_best_voices_enabled,attr"`
	// ContactFlowLogsEnabled: bool, optional
	ContactFlowLogsEnabled terra.BoolValue `hcl:"contact_flow_logs_enabled,attr"`
	// ContactLensEnabled: bool, optional
	ContactLensEnabled terra.BoolValue `hcl:"contact_lens_enabled,attr"`
	// DirectoryId: string, optional
	DirectoryId terra.StringValue `hcl:"directory_id,attr"`
	// EarlyMediaEnabled: bool, optional
	EarlyMediaEnabled terra.BoolValue `hcl:"early_media_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdentityManagementType: string, required
	IdentityManagementType terra.StringValue `hcl:"identity_management_type,attr" validate:"required"`
	// InboundCallsEnabled: bool, required
	InboundCallsEnabled terra.BoolValue `hcl:"inbound_calls_enabled,attr" validate:"required"`
	// InstanceAlias: string, optional
	InstanceAlias terra.StringValue `hcl:"instance_alias,attr"`
	// MultiPartyConferenceEnabled: bool, optional
	MultiPartyConferenceEnabled terra.BoolValue `hcl:"multi_party_conference_enabled,attr"`
	// OutboundCallsEnabled: bool, required
	OutboundCallsEnabled terra.BoolValue `hcl:"outbound_calls_enabled,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *connectinstance.Timeouts `hcl:"timeouts,block"`
}
type connectInstanceAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_connect_instance.
func (ci connectInstanceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("arn"))
}

// AutoResolveBestVoicesEnabled returns a reference to field auto_resolve_best_voices_enabled of aws_connect_instance.
func (ci connectInstanceAttributes) AutoResolveBestVoicesEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ci.ref.Append("auto_resolve_best_voices_enabled"))
}

// ContactFlowLogsEnabled returns a reference to field contact_flow_logs_enabled of aws_connect_instance.
func (ci connectInstanceAttributes) ContactFlowLogsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ci.ref.Append("contact_flow_logs_enabled"))
}

// ContactLensEnabled returns a reference to field contact_lens_enabled of aws_connect_instance.
func (ci connectInstanceAttributes) ContactLensEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ci.ref.Append("contact_lens_enabled"))
}

// CreatedTime returns a reference to field created_time of aws_connect_instance.
func (ci connectInstanceAttributes) CreatedTime() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("created_time"))
}

// DirectoryId returns a reference to field directory_id of aws_connect_instance.
func (ci connectInstanceAttributes) DirectoryId() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("directory_id"))
}

// EarlyMediaEnabled returns a reference to field early_media_enabled of aws_connect_instance.
func (ci connectInstanceAttributes) EarlyMediaEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ci.ref.Append("early_media_enabled"))
}

// Id returns a reference to field id of aws_connect_instance.
func (ci connectInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("id"))
}

// IdentityManagementType returns a reference to field identity_management_type of aws_connect_instance.
func (ci connectInstanceAttributes) IdentityManagementType() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("identity_management_type"))
}

// InboundCallsEnabled returns a reference to field inbound_calls_enabled of aws_connect_instance.
func (ci connectInstanceAttributes) InboundCallsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ci.ref.Append("inbound_calls_enabled"))
}

// InstanceAlias returns a reference to field instance_alias of aws_connect_instance.
func (ci connectInstanceAttributes) InstanceAlias() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("instance_alias"))
}

// MultiPartyConferenceEnabled returns a reference to field multi_party_conference_enabled of aws_connect_instance.
func (ci connectInstanceAttributes) MultiPartyConferenceEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ci.ref.Append("multi_party_conference_enabled"))
}

// OutboundCallsEnabled returns a reference to field outbound_calls_enabled of aws_connect_instance.
func (ci connectInstanceAttributes) OutboundCallsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ci.ref.Append("outbound_calls_enabled"))
}

// ServiceRole returns a reference to field service_role of aws_connect_instance.
func (ci connectInstanceAttributes) ServiceRole() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("service_role"))
}

// Status returns a reference to field status of aws_connect_instance.
func (ci connectInstanceAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ci.ref.Append("status"))
}

func (ci connectInstanceAttributes) Timeouts() connectinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[connectinstance.TimeoutsAttributes](ci.ref.Append("timeouts"))
}

type connectInstanceState struct {
	Arn                          string                         `json:"arn"`
	AutoResolveBestVoicesEnabled bool                           `json:"auto_resolve_best_voices_enabled"`
	ContactFlowLogsEnabled       bool                           `json:"contact_flow_logs_enabled"`
	ContactLensEnabled           bool                           `json:"contact_lens_enabled"`
	CreatedTime                  string                         `json:"created_time"`
	DirectoryId                  string                         `json:"directory_id"`
	EarlyMediaEnabled            bool                           `json:"early_media_enabled"`
	Id                           string                         `json:"id"`
	IdentityManagementType       string                         `json:"identity_management_type"`
	InboundCallsEnabled          bool                           `json:"inbound_calls_enabled"`
	InstanceAlias                string                         `json:"instance_alias"`
	MultiPartyConferenceEnabled  bool                           `json:"multi_party_conference_enabled"`
	OutboundCallsEnabled         bool                           `json:"outbound_calls_enabled"`
	ServiceRole                  string                         `json:"service_role"`
	Status                       string                         `json:"status"`
	Timeouts                     *connectinstance.TimeoutsState `json:"timeouts"`
}
