// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	networkmanagervpcattachment "github.com/golingon/terraproviders/aws/4.63.0/networkmanagervpcattachment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkmanagerVpcAttachment creates a new instance of [NetworkmanagerVpcAttachment].
func NewNetworkmanagerVpcAttachment(name string, args NetworkmanagerVpcAttachmentArgs) *NetworkmanagerVpcAttachment {
	return &NetworkmanagerVpcAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkmanagerVpcAttachment)(nil)

// NetworkmanagerVpcAttachment represents the Terraform resource aws_networkmanager_vpc_attachment.
type NetworkmanagerVpcAttachment struct {
	Name      string
	Args      NetworkmanagerVpcAttachmentArgs
	state     *networkmanagerVpcAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkmanagerVpcAttachment].
func (nva *NetworkmanagerVpcAttachment) Type() string {
	return "aws_networkmanager_vpc_attachment"
}

// LocalName returns the local name for [NetworkmanagerVpcAttachment].
func (nva *NetworkmanagerVpcAttachment) LocalName() string {
	return nva.Name
}

// Configuration returns the configuration (args) for [NetworkmanagerVpcAttachment].
func (nva *NetworkmanagerVpcAttachment) Configuration() interface{} {
	return nva.Args
}

// DependOn is used for other resources to depend on [NetworkmanagerVpcAttachment].
func (nva *NetworkmanagerVpcAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(nva)
}

// Dependencies returns the list of resources [NetworkmanagerVpcAttachment] depends_on.
func (nva *NetworkmanagerVpcAttachment) Dependencies() terra.Dependencies {
	return nva.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkmanagerVpcAttachment].
func (nva *NetworkmanagerVpcAttachment) LifecycleManagement() *terra.Lifecycle {
	return nva.Lifecycle
}

// Attributes returns the attributes for [NetworkmanagerVpcAttachment].
func (nva *NetworkmanagerVpcAttachment) Attributes() networkmanagerVpcAttachmentAttributes {
	return networkmanagerVpcAttachmentAttributes{ref: terra.ReferenceResource(nva)}
}

// ImportState imports the given attribute values into [NetworkmanagerVpcAttachment]'s state.
func (nva *NetworkmanagerVpcAttachment) ImportState(av io.Reader) error {
	nva.state = &networkmanagerVpcAttachmentState{}
	if err := json.NewDecoder(av).Decode(nva.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nva.Type(), nva.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkmanagerVpcAttachment] has state.
func (nva *NetworkmanagerVpcAttachment) State() (*networkmanagerVpcAttachmentState, bool) {
	return nva.state, nva.state != nil
}

// StateMust returns the state for [NetworkmanagerVpcAttachment]. Panics if the state is nil.
func (nva *NetworkmanagerVpcAttachment) StateMust() *networkmanagerVpcAttachmentState {
	if nva.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nva.Type(), nva.LocalName()))
	}
	return nva.state
}

// NetworkmanagerVpcAttachmentArgs contains the configurations for aws_networkmanager_vpc_attachment.
type NetworkmanagerVpcAttachmentArgs struct {
	// CoreNetworkId: string, required
	CoreNetworkId terra.StringValue `hcl:"core_network_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SubnetArns: set of string, required
	SubnetArns terra.SetValue[terra.StringValue] `hcl:"subnet_arns,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VpcArn: string, required
	VpcArn terra.StringValue `hcl:"vpc_arn,attr" validate:"required"`
	// Options: optional
	Options *networkmanagervpcattachment.Options `hcl:"options,block"`
	// Timeouts: optional
	Timeouts *networkmanagervpcattachment.Timeouts `hcl:"timeouts,block"`
}
type networkmanagerVpcAttachmentAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_networkmanager_vpc_attachment.
func (nva networkmanagerVpcAttachmentAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(nva.ref.Append("arn"))
}

// AttachmentPolicyRuleNumber returns a reference to field attachment_policy_rule_number of aws_networkmanager_vpc_attachment.
func (nva networkmanagerVpcAttachmentAttributes) AttachmentPolicyRuleNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(nva.ref.Append("attachment_policy_rule_number"))
}

// AttachmentType returns a reference to field attachment_type of aws_networkmanager_vpc_attachment.
func (nva networkmanagerVpcAttachmentAttributes) AttachmentType() terra.StringValue {
	return terra.ReferenceAsString(nva.ref.Append("attachment_type"))
}

// CoreNetworkArn returns a reference to field core_network_arn of aws_networkmanager_vpc_attachment.
func (nva networkmanagerVpcAttachmentAttributes) CoreNetworkArn() terra.StringValue {
	return terra.ReferenceAsString(nva.ref.Append("core_network_arn"))
}

// CoreNetworkId returns a reference to field core_network_id of aws_networkmanager_vpc_attachment.
func (nva networkmanagerVpcAttachmentAttributes) CoreNetworkId() terra.StringValue {
	return terra.ReferenceAsString(nva.ref.Append("core_network_id"))
}

// EdgeLocation returns a reference to field edge_location of aws_networkmanager_vpc_attachment.
func (nva networkmanagerVpcAttachmentAttributes) EdgeLocation() terra.StringValue {
	return terra.ReferenceAsString(nva.ref.Append("edge_location"))
}

// Id returns a reference to field id of aws_networkmanager_vpc_attachment.
func (nva networkmanagerVpcAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nva.ref.Append("id"))
}

// OwnerAccountId returns a reference to field owner_account_id of aws_networkmanager_vpc_attachment.
func (nva networkmanagerVpcAttachmentAttributes) OwnerAccountId() terra.StringValue {
	return terra.ReferenceAsString(nva.ref.Append("owner_account_id"))
}

// ResourceArn returns a reference to field resource_arn of aws_networkmanager_vpc_attachment.
func (nva networkmanagerVpcAttachmentAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(nva.ref.Append("resource_arn"))
}

// SegmentName returns a reference to field segment_name of aws_networkmanager_vpc_attachment.
func (nva networkmanagerVpcAttachmentAttributes) SegmentName() terra.StringValue {
	return terra.ReferenceAsString(nva.ref.Append("segment_name"))
}

// State returns a reference to field state of aws_networkmanager_vpc_attachment.
func (nva networkmanagerVpcAttachmentAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(nva.ref.Append("state"))
}

// SubnetArns returns a reference to field subnet_arns of aws_networkmanager_vpc_attachment.
func (nva networkmanagerVpcAttachmentAttributes) SubnetArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](nva.ref.Append("subnet_arns"))
}

// Tags returns a reference to field tags of aws_networkmanager_vpc_attachment.
func (nva networkmanagerVpcAttachmentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nva.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_networkmanager_vpc_attachment.
func (nva networkmanagerVpcAttachmentAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nva.ref.Append("tags_all"))
}

// VpcArn returns a reference to field vpc_arn of aws_networkmanager_vpc_attachment.
func (nva networkmanagerVpcAttachmentAttributes) VpcArn() terra.StringValue {
	return terra.ReferenceAsString(nva.ref.Append("vpc_arn"))
}

func (nva networkmanagerVpcAttachmentAttributes) Options() terra.ListValue[networkmanagervpcattachment.OptionsAttributes] {
	return terra.ReferenceAsList[networkmanagervpcattachment.OptionsAttributes](nva.ref.Append("options"))
}

func (nva networkmanagerVpcAttachmentAttributes) Timeouts() networkmanagervpcattachment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanagervpcattachment.TimeoutsAttributes](nva.ref.Append("timeouts"))
}

type networkmanagerVpcAttachmentState struct {
	Arn                        string                                     `json:"arn"`
	AttachmentPolicyRuleNumber float64                                    `json:"attachment_policy_rule_number"`
	AttachmentType             string                                     `json:"attachment_type"`
	CoreNetworkArn             string                                     `json:"core_network_arn"`
	CoreNetworkId              string                                     `json:"core_network_id"`
	EdgeLocation               string                                     `json:"edge_location"`
	Id                         string                                     `json:"id"`
	OwnerAccountId             string                                     `json:"owner_account_id"`
	ResourceArn                string                                     `json:"resource_arn"`
	SegmentName                string                                     `json:"segment_name"`
	State                      string                                     `json:"state"`
	SubnetArns                 []string                                   `json:"subnet_arns"`
	Tags                       map[string]string                          `json:"tags"`
	TagsAll                    map[string]string                          `json:"tags_all"`
	VpcArn                     string                                     `json:"vpc_arn"`
	Options                    []networkmanagervpcattachment.OptionsState `json:"options"`
	Timeouts                   *networkmanagervpcattachment.TimeoutsState `json:"timeouts"`
}
