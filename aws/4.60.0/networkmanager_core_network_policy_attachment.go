// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	networkmanagercorenetworkpolicyattachment "github.com/golingon/terraproviders/aws/4.60.0/networkmanagercorenetworkpolicyattachment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkmanagerCoreNetworkPolicyAttachment creates a new instance of [NetworkmanagerCoreNetworkPolicyAttachment].
func NewNetworkmanagerCoreNetworkPolicyAttachment(name string, args NetworkmanagerCoreNetworkPolicyAttachmentArgs) *NetworkmanagerCoreNetworkPolicyAttachment {
	return &NetworkmanagerCoreNetworkPolicyAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkmanagerCoreNetworkPolicyAttachment)(nil)

// NetworkmanagerCoreNetworkPolicyAttachment represents the Terraform resource aws_networkmanager_core_network_policy_attachment.
type NetworkmanagerCoreNetworkPolicyAttachment struct {
	Name      string
	Args      NetworkmanagerCoreNetworkPolicyAttachmentArgs
	state     *networkmanagerCoreNetworkPolicyAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkmanagerCoreNetworkPolicyAttachment].
func (ncnpa *NetworkmanagerCoreNetworkPolicyAttachment) Type() string {
	return "aws_networkmanager_core_network_policy_attachment"
}

// LocalName returns the local name for [NetworkmanagerCoreNetworkPolicyAttachment].
func (ncnpa *NetworkmanagerCoreNetworkPolicyAttachment) LocalName() string {
	return ncnpa.Name
}

// Configuration returns the configuration (args) for [NetworkmanagerCoreNetworkPolicyAttachment].
func (ncnpa *NetworkmanagerCoreNetworkPolicyAttachment) Configuration() interface{} {
	return ncnpa.Args
}

// DependOn is used for other resources to depend on [NetworkmanagerCoreNetworkPolicyAttachment].
func (ncnpa *NetworkmanagerCoreNetworkPolicyAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(ncnpa)
}

// Dependencies returns the list of resources [NetworkmanagerCoreNetworkPolicyAttachment] depends_on.
func (ncnpa *NetworkmanagerCoreNetworkPolicyAttachment) Dependencies() terra.Dependencies {
	return ncnpa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkmanagerCoreNetworkPolicyAttachment].
func (ncnpa *NetworkmanagerCoreNetworkPolicyAttachment) LifecycleManagement() *terra.Lifecycle {
	return ncnpa.Lifecycle
}

// Attributes returns the attributes for [NetworkmanagerCoreNetworkPolicyAttachment].
func (ncnpa *NetworkmanagerCoreNetworkPolicyAttachment) Attributes() networkmanagerCoreNetworkPolicyAttachmentAttributes {
	return networkmanagerCoreNetworkPolicyAttachmentAttributes{ref: terra.ReferenceResource(ncnpa)}
}

// ImportState imports the given attribute values into [NetworkmanagerCoreNetworkPolicyAttachment]'s state.
func (ncnpa *NetworkmanagerCoreNetworkPolicyAttachment) ImportState(av io.Reader) error {
	ncnpa.state = &networkmanagerCoreNetworkPolicyAttachmentState{}
	if err := json.NewDecoder(av).Decode(ncnpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ncnpa.Type(), ncnpa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkmanagerCoreNetworkPolicyAttachment] has state.
func (ncnpa *NetworkmanagerCoreNetworkPolicyAttachment) State() (*networkmanagerCoreNetworkPolicyAttachmentState, bool) {
	return ncnpa.state, ncnpa.state != nil
}

// StateMust returns the state for [NetworkmanagerCoreNetworkPolicyAttachment]. Panics if the state is nil.
func (ncnpa *NetworkmanagerCoreNetworkPolicyAttachment) StateMust() *networkmanagerCoreNetworkPolicyAttachmentState {
	if ncnpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ncnpa.Type(), ncnpa.LocalName()))
	}
	return ncnpa.state
}

// NetworkmanagerCoreNetworkPolicyAttachmentArgs contains the configurations for aws_networkmanager_core_network_policy_attachment.
type NetworkmanagerCoreNetworkPolicyAttachmentArgs struct {
	// CoreNetworkId: string, required
	CoreNetworkId terra.StringValue `hcl:"core_network_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyDocument: string, required
	PolicyDocument terra.StringValue `hcl:"policy_document,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *networkmanagercorenetworkpolicyattachment.Timeouts `hcl:"timeouts,block"`
}
type networkmanagerCoreNetworkPolicyAttachmentAttributes struct {
	ref terra.Reference
}

// CoreNetworkId returns a reference to field core_network_id of aws_networkmanager_core_network_policy_attachment.
func (ncnpa networkmanagerCoreNetworkPolicyAttachmentAttributes) CoreNetworkId() terra.StringValue {
	return terra.ReferenceAsString(ncnpa.ref.Append("core_network_id"))
}

// Id returns a reference to field id of aws_networkmanager_core_network_policy_attachment.
func (ncnpa networkmanagerCoreNetworkPolicyAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ncnpa.ref.Append("id"))
}

// PolicyDocument returns a reference to field policy_document of aws_networkmanager_core_network_policy_attachment.
func (ncnpa networkmanagerCoreNetworkPolicyAttachmentAttributes) PolicyDocument() terra.StringValue {
	return terra.ReferenceAsString(ncnpa.ref.Append("policy_document"))
}

// State returns a reference to field state of aws_networkmanager_core_network_policy_attachment.
func (ncnpa networkmanagerCoreNetworkPolicyAttachmentAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ncnpa.ref.Append("state"))
}

func (ncnpa networkmanagerCoreNetworkPolicyAttachmentAttributes) Timeouts() networkmanagercorenetworkpolicyattachment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanagercorenetworkpolicyattachment.TimeoutsAttributes](ncnpa.ref.Append("timeouts"))
}

type networkmanagerCoreNetworkPolicyAttachmentState struct {
	CoreNetworkId  string                                                   `json:"core_network_id"`
	Id             string                                                   `json:"id"`
	PolicyDocument string                                                   `json:"policy_document"`
	State          string                                                   `json:"state"`
	Timeouts       *networkmanagercorenetworkpolicyattachment.TimeoutsState `json:"timeouts"`
}
