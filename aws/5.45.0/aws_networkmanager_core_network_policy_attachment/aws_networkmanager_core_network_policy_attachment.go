// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_networkmanager_core_network_policy_attachment

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

// Resource represents the Terraform resource aws_networkmanager_core_network_policy_attachment.
type Resource struct {
	Name      string
	Args      Args
	state     *awsNetworkmanagerCoreNetworkPolicyAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ancnpa *Resource) Type() string {
	return "aws_networkmanager_core_network_policy_attachment"
}

// LocalName returns the local name for [Resource].
func (ancnpa *Resource) LocalName() string {
	return ancnpa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ancnpa *Resource) Configuration() interface{} {
	return ancnpa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ancnpa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ancnpa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ancnpa *Resource) Dependencies() terra.Dependencies {
	return ancnpa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ancnpa *Resource) LifecycleManagement() *terra.Lifecycle {
	return ancnpa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ancnpa *Resource) Attributes() awsNetworkmanagerCoreNetworkPolicyAttachmentAttributes {
	return awsNetworkmanagerCoreNetworkPolicyAttachmentAttributes{ref: terra.ReferenceResource(ancnpa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ancnpa *Resource) ImportState(state io.Reader) error {
	ancnpa.state = &awsNetworkmanagerCoreNetworkPolicyAttachmentState{}
	if err := json.NewDecoder(state).Decode(ancnpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ancnpa.Type(), ancnpa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ancnpa *Resource) State() (*awsNetworkmanagerCoreNetworkPolicyAttachmentState, bool) {
	return ancnpa.state, ancnpa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ancnpa *Resource) StateMust() *awsNetworkmanagerCoreNetworkPolicyAttachmentState {
	if ancnpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ancnpa.Type(), ancnpa.LocalName()))
	}
	return ancnpa.state
}

// Args contains the configurations for aws_networkmanager_core_network_policy_attachment.
type Args struct {
	// CoreNetworkId: string, required
	CoreNetworkId terra.StringValue `hcl:"core_network_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyDocument: string, required
	PolicyDocument terra.StringValue `hcl:"policy_document,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsNetworkmanagerCoreNetworkPolicyAttachmentAttributes struct {
	ref terra.Reference
}

// CoreNetworkId returns a reference to field core_network_id of aws_networkmanager_core_network_policy_attachment.
func (ancnpa awsNetworkmanagerCoreNetworkPolicyAttachmentAttributes) CoreNetworkId() terra.StringValue {
	return terra.ReferenceAsString(ancnpa.ref.Append("core_network_id"))
}

// Id returns a reference to field id of aws_networkmanager_core_network_policy_attachment.
func (ancnpa awsNetworkmanagerCoreNetworkPolicyAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ancnpa.ref.Append("id"))
}

// PolicyDocument returns a reference to field policy_document of aws_networkmanager_core_network_policy_attachment.
func (ancnpa awsNetworkmanagerCoreNetworkPolicyAttachmentAttributes) PolicyDocument() terra.StringValue {
	return terra.ReferenceAsString(ancnpa.ref.Append("policy_document"))
}

// State returns a reference to field state of aws_networkmanager_core_network_policy_attachment.
func (ancnpa awsNetworkmanagerCoreNetworkPolicyAttachmentAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ancnpa.ref.Append("state"))
}

func (ancnpa awsNetworkmanagerCoreNetworkPolicyAttachmentAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ancnpa.ref.Append("timeouts"))
}

type awsNetworkmanagerCoreNetworkPolicyAttachmentState struct {
	CoreNetworkId  string         `json:"core_network_id"`
	Id             string         `json:"id"`
	PolicyDocument string         `json:"policy_document"`
	State          string         `json:"state"`
	Timeouts       *TimeoutsState `json:"timeouts"`
}
