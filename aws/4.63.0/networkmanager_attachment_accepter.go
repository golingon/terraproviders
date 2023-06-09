// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	networkmanagerattachmentaccepter "github.com/golingon/terraproviders/aws/4.63.0/networkmanagerattachmentaccepter"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkmanagerAttachmentAccepter creates a new instance of [NetworkmanagerAttachmentAccepter].
func NewNetworkmanagerAttachmentAccepter(name string, args NetworkmanagerAttachmentAccepterArgs) *NetworkmanagerAttachmentAccepter {
	return &NetworkmanagerAttachmentAccepter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkmanagerAttachmentAccepter)(nil)

// NetworkmanagerAttachmentAccepter represents the Terraform resource aws_networkmanager_attachment_accepter.
type NetworkmanagerAttachmentAccepter struct {
	Name      string
	Args      NetworkmanagerAttachmentAccepterArgs
	state     *networkmanagerAttachmentAccepterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkmanagerAttachmentAccepter].
func (naa *NetworkmanagerAttachmentAccepter) Type() string {
	return "aws_networkmanager_attachment_accepter"
}

// LocalName returns the local name for [NetworkmanagerAttachmentAccepter].
func (naa *NetworkmanagerAttachmentAccepter) LocalName() string {
	return naa.Name
}

// Configuration returns the configuration (args) for [NetworkmanagerAttachmentAccepter].
func (naa *NetworkmanagerAttachmentAccepter) Configuration() interface{} {
	return naa.Args
}

// DependOn is used for other resources to depend on [NetworkmanagerAttachmentAccepter].
func (naa *NetworkmanagerAttachmentAccepter) DependOn() terra.Reference {
	return terra.ReferenceResource(naa)
}

// Dependencies returns the list of resources [NetworkmanagerAttachmentAccepter] depends_on.
func (naa *NetworkmanagerAttachmentAccepter) Dependencies() terra.Dependencies {
	return naa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkmanagerAttachmentAccepter].
func (naa *NetworkmanagerAttachmentAccepter) LifecycleManagement() *terra.Lifecycle {
	return naa.Lifecycle
}

// Attributes returns the attributes for [NetworkmanagerAttachmentAccepter].
func (naa *NetworkmanagerAttachmentAccepter) Attributes() networkmanagerAttachmentAccepterAttributes {
	return networkmanagerAttachmentAccepterAttributes{ref: terra.ReferenceResource(naa)}
}

// ImportState imports the given attribute values into [NetworkmanagerAttachmentAccepter]'s state.
func (naa *NetworkmanagerAttachmentAccepter) ImportState(av io.Reader) error {
	naa.state = &networkmanagerAttachmentAccepterState{}
	if err := json.NewDecoder(av).Decode(naa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", naa.Type(), naa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkmanagerAttachmentAccepter] has state.
func (naa *NetworkmanagerAttachmentAccepter) State() (*networkmanagerAttachmentAccepterState, bool) {
	return naa.state, naa.state != nil
}

// StateMust returns the state for [NetworkmanagerAttachmentAccepter]. Panics if the state is nil.
func (naa *NetworkmanagerAttachmentAccepter) StateMust() *networkmanagerAttachmentAccepterState {
	if naa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", naa.Type(), naa.LocalName()))
	}
	return naa.state
}

// NetworkmanagerAttachmentAccepterArgs contains the configurations for aws_networkmanager_attachment_accepter.
type NetworkmanagerAttachmentAccepterArgs struct {
	// AttachmentId: string, required
	AttachmentId terra.StringValue `hcl:"attachment_id,attr" validate:"required"`
	// AttachmentType: string, required
	AttachmentType terra.StringValue `hcl:"attachment_type,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *networkmanagerattachmentaccepter.Timeouts `hcl:"timeouts,block"`
}
type networkmanagerAttachmentAccepterAttributes struct {
	ref terra.Reference
}

// AttachmentId returns a reference to field attachment_id of aws_networkmanager_attachment_accepter.
func (naa networkmanagerAttachmentAccepterAttributes) AttachmentId() terra.StringValue {
	return terra.ReferenceAsString(naa.ref.Append("attachment_id"))
}

// AttachmentPolicyRuleNumber returns a reference to field attachment_policy_rule_number of aws_networkmanager_attachment_accepter.
func (naa networkmanagerAttachmentAccepterAttributes) AttachmentPolicyRuleNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(naa.ref.Append("attachment_policy_rule_number"))
}

// AttachmentType returns a reference to field attachment_type of aws_networkmanager_attachment_accepter.
func (naa networkmanagerAttachmentAccepterAttributes) AttachmentType() terra.StringValue {
	return terra.ReferenceAsString(naa.ref.Append("attachment_type"))
}

// CoreNetworkArn returns a reference to field core_network_arn of aws_networkmanager_attachment_accepter.
func (naa networkmanagerAttachmentAccepterAttributes) CoreNetworkArn() terra.StringValue {
	return terra.ReferenceAsString(naa.ref.Append("core_network_arn"))
}

// CoreNetworkId returns a reference to field core_network_id of aws_networkmanager_attachment_accepter.
func (naa networkmanagerAttachmentAccepterAttributes) CoreNetworkId() terra.StringValue {
	return terra.ReferenceAsString(naa.ref.Append("core_network_id"))
}

// EdgeLocation returns a reference to field edge_location of aws_networkmanager_attachment_accepter.
func (naa networkmanagerAttachmentAccepterAttributes) EdgeLocation() terra.StringValue {
	return terra.ReferenceAsString(naa.ref.Append("edge_location"))
}

// Id returns a reference to field id of aws_networkmanager_attachment_accepter.
func (naa networkmanagerAttachmentAccepterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(naa.ref.Append("id"))
}

// OwnerAccountId returns a reference to field owner_account_id of aws_networkmanager_attachment_accepter.
func (naa networkmanagerAttachmentAccepterAttributes) OwnerAccountId() terra.StringValue {
	return terra.ReferenceAsString(naa.ref.Append("owner_account_id"))
}

// ResourceArn returns a reference to field resource_arn of aws_networkmanager_attachment_accepter.
func (naa networkmanagerAttachmentAccepterAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(naa.ref.Append("resource_arn"))
}

// SegmentName returns a reference to field segment_name of aws_networkmanager_attachment_accepter.
func (naa networkmanagerAttachmentAccepterAttributes) SegmentName() terra.StringValue {
	return terra.ReferenceAsString(naa.ref.Append("segment_name"))
}

// State returns a reference to field state of aws_networkmanager_attachment_accepter.
func (naa networkmanagerAttachmentAccepterAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(naa.ref.Append("state"))
}

func (naa networkmanagerAttachmentAccepterAttributes) Timeouts() networkmanagerattachmentaccepter.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanagerattachmentaccepter.TimeoutsAttributes](naa.ref.Append("timeouts"))
}

type networkmanagerAttachmentAccepterState struct {
	AttachmentId               string                                          `json:"attachment_id"`
	AttachmentPolicyRuleNumber float64                                         `json:"attachment_policy_rule_number"`
	AttachmentType             string                                          `json:"attachment_type"`
	CoreNetworkArn             string                                          `json:"core_network_arn"`
	CoreNetworkId              string                                          `json:"core_network_id"`
	EdgeLocation               string                                          `json:"edge_location"`
	Id                         string                                          `json:"id"`
	OwnerAccountId             string                                          `json:"owner_account_id"`
	ResourceArn                string                                          `json:"resource_arn"`
	SegmentName                string                                          `json:"segment_name"`
	State                      string                                          `json:"state"`
	Timeouts                   *networkmanagerattachmentaccepter.TimeoutsState `json:"timeouts"`
}
