// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	networkmanagersitetositevpnattachment "github.com/golingon/terraproviders/aws/5.0.1/networkmanagersitetositevpnattachment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkmanagerSiteToSiteVpnAttachment creates a new instance of [NetworkmanagerSiteToSiteVpnAttachment].
func NewNetworkmanagerSiteToSiteVpnAttachment(name string, args NetworkmanagerSiteToSiteVpnAttachmentArgs) *NetworkmanagerSiteToSiteVpnAttachment {
	return &NetworkmanagerSiteToSiteVpnAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkmanagerSiteToSiteVpnAttachment)(nil)

// NetworkmanagerSiteToSiteVpnAttachment represents the Terraform resource aws_networkmanager_site_to_site_vpn_attachment.
type NetworkmanagerSiteToSiteVpnAttachment struct {
	Name      string
	Args      NetworkmanagerSiteToSiteVpnAttachmentArgs
	state     *networkmanagerSiteToSiteVpnAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkmanagerSiteToSiteVpnAttachment].
func (nstsva *NetworkmanagerSiteToSiteVpnAttachment) Type() string {
	return "aws_networkmanager_site_to_site_vpn_attachment"
}

// LocalName returns the local name for [NetworkmanagerSiteToSiteVpnAttachment].
func (nstsva *NetworkmanagerSiteToSiteVpnAttachment) LocalName() string {
	return nstsva.Name
}

// Configuration returns the configuration (args) for [NetworkmanagerSiteToSiteVpnAttachment].
func (nstsva *NetworkmanagerSiteToSiteVpnAttachment) Configuration() interface{} {
	return nstsva.Args
}

// DependOn is used for other resources to depend on [NetworkmanagerSiteToSiteVpnAttachment].
func (nstsva *NetworkmanagerSiteToSiteVpnAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(nstsva)
}

// Dependencies returns the list of resources [NetworkmanagerSiteToSiteVpnAttachment] depends_on.
func (nstsva *NetworkmanagerSiteToSiteVpnAttachment) Dependencies() terra.Dependencies {
	return nstsva.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkmanagerSiteToSiteVpnAttachment].
func (nstsva *NetworkmanagerSiteToSiteVpnAttachment) LifecycleManagement() *terra.Lifecycle {
	return nstsva.Lifecycle
}

// Attributes returns the attributes for [NetworkmanagerSiteToSiteVpnAttachment].
func (nstsva *NetworkmanagerSiteToSiteVpnAttachment) Attributes() networkmanagerSiteToSiteVpnAttachmentAttributes {
	return networkmanagerSiteToSiteVpnAttachmentAttributes{ref: terra.ReferenceResource(nstsva)}
}

// ImportState imports the given attribute values into [NetworkmanagerSiteToSiteVpnAttachment]'s state.
func (nstsva *NetworkmanagerSiteToSiteVpnAttachment) ImportState(av io.Reader) error {
	nstsva.state = &networkmanagerSiteToSiteVpnAttachmentState{}
	if err := json.NewDecoder(av).Decode(nstsva.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nstsva.Type(), nstsva.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkmanagerSiteToSiteVpnAttachment] has state.
func (nstsva *NetworkmanagerSiteToSiteVpnAttachment) State() (*networkmanagerSiteToSiteVpnAttachmentState, bool) {
	return nstsva.state, nstsva.state != nil
}

// StateMust returns the state for [NetworkmanagerSiteToSiteVpnAttachment]. Panics if the state is nil.
func (nstsva *NetworkmanagerSiteToSiteVpnAttachment) StateMust() *networkmanagerSiteToSiteVpnAttachmentState {
	if nstsva.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nstsva.Type(), nstsva.LocalName()))
	}
	return nstsva.state
}

// NetworkmanagerSiteToSiteVpnAttachmentArgs contains the configurations for aws_networkmanager_site_to_site_vpn_attachment.
type NetworkmanagerSiteToSiteVpnAttachmentArgs struct {
	// CoreNetworkId: string, required
	CoreNetworkId terra.StringValue `hcl:"core_network_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VpnConnectionArn: string, required
	VpnConnectionArn terra.StringValue `hcl:"vpn_connection_arn,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *networkmanagersitetositevpnattachment.Timeouts `hcl:"timeouts,block"`
}
type networkmanagerSiteToSiteVpnAttachmentAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_networkmanager_site_to_site_vpn_attachment.
func (nstsva networkmanagerSiteToSiteVpnAttachmentAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(nstsva.ref.Append("arn"))
}

// AttachmentPolicyRuleNumber returns a reference to field attachment_policy_rule_number of aws_networkmanager_site_to_site_vpn_attachment.
func (nstsva networkmanagerSiteToSiteVpnAttachmentAttributes) AttachmentPolicyRuleNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(nstsva.ref.Append("attachment_policy_rule_number"))
}

// AttachmentType returns a reference to field attachment_type of aws_networkmanager_site_to_site_vpn_attachment.
func (nstsva networkmanagerSiteToSiteVpnAttachmentAttributes) AttachmentType() terra.StringValue {
	return terra.ReferenceAsString(nstsva.ref.Append("attachment_type"))
}

// CoreNetworkArn returns a reference to field core_network_arn of aws_networkmanager_site_to_site_vpn_attachment.
func (nstsva networkmanagerSiteToSiteVpnAttachmentAttributes) CoreNetworkArn() terra.StringValue {
	return terra.ReferenceAsString(nstsva.ref.Append("core_network_arn"))
}

// CoreNetworkId returns a reference to field core_network_id of aws_networkmanager_site_to_site_vpn_attachment.
func (nstsva networkmanagerSiteToSiteVpnAttachmentAttributes) CoreNetworkId() terra.StringValue {
	return terra.ReferenceAsString(nstsva.ref.Append("core_network_id"))
}

// EdgeLocation returns a reference to field edge_location of aws_networkmanager_site_to_site_vpn_attachment.
func (nstsva networkmanagerSiteToSiteVpnAttachmentAttributes) EdgeLocation() terra.StringValue {
	return terra.ReferenceAsString(nstsva.ref.Append("edge_location"))
}

// Id returns a reference to field id of aws_networkmanager_site_to_site_vpn_attachment.
func (nstsva networkmanagerSiteToSiteVpnAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nstsva.ref.Append("id"))
}

// OwnerAccountId returns a reference to field owner_account_id of aws_networkmanager_site_to_site_vpn_attachment.
func (nstsva networkmanagerSiteToSiteVpnAttachmentAttributes) OwnerAccountId() terra.StringValue {
	return terra.ReferenceAsString(nstsva.ref.Append("owner_account_id"))
}

// ResourceArn returns a reference to field resource_arn of aws_networkmanager_site_to_site_vpn_attachment.
func (nstsva networkmanagerSiteToSiteVpnAttachmentAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(nstsva.ref.Append("resource_arn"))
}

// SegmentName returns a reference to field segment_name of aws_networkmanager_site_to_site_vpn_attachment.
func (nstsva networkmanagerSiteToSiteVpnAttachmentAttributes) SegmentName() terra.StringValue {
	return terra.ReferenceAsString(nstsva.ref.Append("segment_name"))
}

// State returns a reference to field state of aws_networkmanager_site_to_site_vpn_attachment.
func (nstsva networkmanagerSiteToSiteVpnAttachmentAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(nstsva.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_networkmanager_site_to_site_vpn_attachment.
func (nstsva networkmanagerSiteToSiteVpnAttachmentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nstsva.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_networkmanager_site_to_site_vpn_attachment.
func (nstsva networkmanagerSiteToSiteVpnAttachmentAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nstsva.ref.Append("tags_all"))
}

// VpnConnectionArn returns a reference to field vpn_connection_arn of aws_networkmanager_site_to_site_vpn_attachment.
func (nstsva networkmanagerSiteToSiteVpnAttachmentAttributes) VpnConnectionArn() terra.StringValue {
	return terra.ReferenceAsString(nstsva.ref.Append("vpn_connection_arn"))
}

func (nstsva networkmanagerSiteToSiteVpnAttachmentAttributes) Timeouts() networkmanagersitetositevpnattachment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanagersitetositevpnattachment.TimeoutsAttributes](nstsva.ref.Append("timeouts"))
}

type networkmanagerSiteToSiteVpnAttachmentState struct {
	Arn                        string                                               `json:"arn"`
	AttachmentPolicyRuleNumber float64                                              `json:"attachment_policy_rule_number"`
	AttachmentType             string                                               `json:"attachment_type"`
	CoreNetworkArn             string                                               `json:"core_network_arn"`
	CoreNetworkId              string                                               `json:"core_network_id"`
	EdgeLocation               string                                               `json:"edge_location"`
	Id                         string                                               `json:"id"`
	OwnerAccountId             string                                               `json:"owner_account_id"`
	ResourceArn                string                                               `json:"resource_arn"`
	SegmentName                string                                               `json:"segment_name"`
	State                      string                                               `json:"state"`
	Tags                       map[string]string                                    `json:"tags"`
	TagsAll                    map[string]string                                    `json:"tags_all"`
	VpnConnectionArn           string                                               `json:"vpn_connection_arn"`
	Timeouts                   *networkmanagersitetositevpnattachment.TimeoutsState `json:"timeouts"`
}
