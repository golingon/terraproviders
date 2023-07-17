// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	networkmanagerconnectattachment "github.com/golingon/terraproviders/aws/5.8.0/networkmanagerconnectattachment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkmanagerConnectAttachment creates a new instance of [NetworkmanagerConnectAttachment].
func NewNetworkmanagerConnectAttachment(name string, args NetworkmanagerConnectAttachmentArgs) *NetworkmanagerConnectAttachment {
	return &NetworkmanagerConnectAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkmanagerConnectAttachment)(nil)

// NetworkmanagerConnectAttachment represents the Terraform resource aws_networkmanager_connect_attachment.
type NetworkmanagerConnectAttachment struct {
	Name      string
	Args      NetworkmanagerConnectAttachmentArgs
	state     *networkmanagerConnectAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkmanagerConnectAttachment].
func (nca *NetworkmanagerConnectAttachment) Type() string {
	return "aws_networkmanager_connect_attachment"
}

// LocalName returns the local name for [NetworkmanagerConnectAttachment].
func (nca *NetworkmanagerConnectAttachment) LocalName() string {
	return nca.Name
}

// Configuration returns the configuration (args) for [NetworkmanagerConnectAttachment].
func (nca *NetworkmanagerConnectAttachment) Configuration() interface{} {
	return nca.Args
}

// DependOn is used for other resources to depend on [NetworkmanagerConnectAttachment].
func (nca *NetworkmanagerConnectAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(nca)
}

// Dependencies returns the list of resources [NetworkmanagerConnectAttachment] depends_on.
func (nca *NetworkmanagerConnectAttachment) Dependencies() terra.Dependencies {
	return nca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkmanagerConnectAttachment].
func (nca *NetworkmanagerConnectAttachment) LifecycleManagement() *terra.Lifecycle {
	return nca.Lifecycle
}

// Attributes returns the attributes for [NetworkmanagerConnectAttachment].
func (nca *NetworkmanagerConnectAttachment) Attributes() networkmanagerConnectAttachmentAttributes {
	return networkmanagerConnectAttachmentAttributes{ref: terra.ReferenceResource(nca)}
}

// ImportState imports the given attribute values into [NetworkmanagerConnectAttachment]'s state.
func (nca *NetworkmanagerConnectAttachment) ImportState(av io.Reader) error {
	nca.state = &networkmanagerConnectAttachmentState{}
	if err := json.NewDecoder(av).Decode(nca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nca.Type(), nca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkmanagerConnectAttachment] has state.
func (nca *NetworkmanagerConnectAttachment) State() (*networkmanagerConnectAttachmentState, bool) {
	return nca.state, nca.state != nil
}

// StateMust returns the state for [NetworkmanagerConnectAttachment]. Panics if the state is nil.
func (nca *NetworkmanagerConnectAttachment) StateMust() *networkmanagerConnectAttachmentState {
	if nca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nca.Type(), nca.LocalName()))
	}
	return nca.state
}

// NetworkmanagerConnectAttachmentArgs contains the configurations for aws_networkmanager_connect_attachment.
type NetworkmanagerConnectAttachmentArgs struct {
	// CoreNetworkId: string, required
	CoreNetworkId terra.StringValue `hcl:"core_network_id,attr" validate:"required"`
	// EdgeLocation: string, required
	EdgeLocation terra.StringValue `hcl:"edge_location,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TransportAttachmentId: string, required
	TransportAttachmentId terra.StringValue `hcl:"transport_attachment_id,attr" validate:"required"`
	// Options: required
	Options *networkmanagerconnectattachment.Options `hcl:"options,block" validate:"required"`
	// Timeouts: optional
	Timeouts *networkmanagerconnectattachment.Timeouts `hcl:"timeouts,block"`
}
type networkmanagerConnectAttachmentAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_networkmanager_connect_attachment.
func (nca networkmanagerConnectAttachmentAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(nca.ref.Append("arn"))
}

// AttachmentId returns a reference to field attachment_id of aws_networkmanager_connect_attachment.
func (nca networkmanagerConnectAttachmentAttributes) AttachmentId() terra.StringValue {
	return terra.ReferenceAsString(nca.ref.Append("attachment_id"))
}

// AttachmentPolicyRuleNumber returns a reference to field attachment_policy_rule_number of aws_networkmanager_connect_attachment.
func (nca networkmanagerConnectAttachmentAttributes) AttachmentPolicyRuleNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(nca.ref.Append("attachment_policy_rule_number"))
}

// AttachmentType returns a reference to field attachment_type of aws_networkmanager_connect_attachment.
func (nca networkmanagerConnectAttachmentAttributes) AttachmentType() terra.StringValue {
	return terra.ReferenceAsString(nca.ref.Append("attachment_type"))
}

// CoreNetworkArn returns a reference to field core_network_arn of aws_networkmanager_connect_attachment.
func (nca networkmanagerConnectAttachmentAttributes) CoreNetworkArn() terra.StringValue {
	return terra.ReferenceAsString(nca.ref.Append("core_network_arn"))
}

// CoreNetworkId returns a reference to field core_network_id of aws_networkmanager_connect_attachment.
func (nca networkmanagerConnectAttachmentAttributes) CoreNetworkId() terra.StringValue {
	return terra.ReferenceAsString(nca.ref.Append("core_network_id"))
}

// EdgeLocation returns a reference to field edge_location of aws_networkmanager_connect_attachment.
func (nca networkmanagerConnectAttachmentAttributes) EdgeLocation() terra.StringValue {
	return terra.ReferenceAsString(nca.ref.Append("edge_location"))
}

// Id returns a reference to field id of aws_networkmanager_connect_attachment.
func (nca networkmanagerConnectAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nca.ref.Append("id"))
}

// OwnerAccountId returns a reference to field owner_account_id of aws_networkmanager_connect_attachment.
func (nca networkmanagerConnectAttachmentAttributes) OwnerAccountId() terra.StringValue {
	return terra.ReferenceAsString(nca.ref.Append("owner_account_id"))
}

// ResourceArn returns a reference to field resource_arn of aws_networkmanager_connect_attachment.
func (nca networkmanagerConnectAttachmentAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(nca.ref.Append("resource_arn"))
}

// SegmentName returns a reference to field segment_name of aws_networkmanager_connect_attachment.
func (nca networkmanagerConnectAttachmentAttributes) SegmentName() terra.StringValue {
	return terra.ReferenceAsString(nca.ref.Append("segment_name"))
}

// State returns a reference to field state of aws_networkmanager_connect_attachment.
func (nca networkmanagerConnectAttachmentAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(nca.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_networkmanager_connect_attachment.
func (nca networkmanagerConnectAttachmentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nca.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_networkmanager_connect_attachment.
func (nca networkmanagerConnectAttachmentAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nca.ref.Append("tags_all"))
}

// TransportAttachmentId returns a reference to field transport_attachment_id of aws_networkmanager_connect_attachment.
func (nca networkmanagerConnectAttachmentAttributes) TransportAttachmentId() terra.StringValue {
	return terra.ReferenceAsString(nca.ref.Append("transport_attachment_id"))
}

func (nca networkmanagerConnectAttachmentAttributes) Options() terra.ListValue[networkmanagerconnectattachment.OptionsAttributes] {
	return terra.ReferenceAsList[networkmanagerconnectattachment.OptionsAttributes](nca.ref.Append("options"))
}

func (nca networkmanagerConnectAttachmentAttributes) Timeouts() networkmanagerconnectattachment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanagerconnectattachment.TimeoutsAttributes](nca.ref.Append("timeouts"))
}

type networkmanagerConnectAttachmentState struct {
	Arn                        string                                         `json:"arn"`
	AttachmentId               string                                         `json:"attachment_id"`
	AttachmentPolicyRuleNumber float64                                        `json:"attachment_policy_rule_number"`
	AttachmentType             string                                         `json:"attachment_type"`
	CoreNetworkArn             string                                         `json:"core_network_arn"`
	CoreNetworkId              string                                         `json:"core_network_id"`
	EdgeLocation               string                                         `json:"edge_location"`
	Id                         string                                         `json:"id"`
	OwnerAccountId             string                                         `json:"owner_account_id"`
	ResourceArn                string                                         `json:"resource_arn"`
	SegmentName                string                                         `json:"segment_name"`
	State                      string                                         `json:"state"`
	Tags                       map[string]string                              `json:"tags"`
	TagsAll                    map[string]string                              `json:"tags_all"`
	TransportAttachmentId      string                                         `json:"transport_attachment_id"`
	Options                    []networkmanagerconnectattachment.OptionsState `json:"options"`
	Timeouts                   *networkmanagerconnectattachment.TimeoutsState `json:"timeouts"`
}
