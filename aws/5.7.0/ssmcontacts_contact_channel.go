// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ssmcontactscontactchannel "github.com/golingon/terraproviders/aws/5.7.0/ssmcontactscontactchannel"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSsmcontactsContactChannel creates a new instance of [SsmcontactsContactChannel].
func NewSsmcontactsContactChannel(name string, args SsmcontactsContactChannelArgs) *SsmcontactsContactChannel {
	return &SsmcontactsContactChannel{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SsmcontactsContactChannel)(nil)

// SsmcontactsContactChannel represents the Terraform resource aws_ssmcontacts_contact_channel.
type SsmcontactsContactChannel struct {
	Name      string
	Args      SsmcontactsContactChannelArgs
	state     *ssmcontactsContactChannelState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SsmcontactsContactChannel].
func (scc *SsmcontactsContactChannel) Type() string {
	return "aws_ssmcontacts_contact_channel"
}

// LocalName returns the local name for [SsmcontactsContactChannel].
func (scc *SsmcontactsContactChannel) LocalName() string {
	return scc.Name
}

// Configuration returns the configuration (args) for [SsmcontactsContactChannel].
func (scc *SsmcontactsContactChannel) Configuration() interface{} {
	return scc.Args
}

// DependOn is used for other resources to depend on [SsmcontactsContactChannel].
func (scc *SsmcontactsContactChannel) DependOn() terra.Reference {
	return terra.ReferenceResource(scc)
}

// Dependencies returns the list of resources [SsmcontactsContactChannel] depends_on.
func (scc *SsmcontactsContactChannel) Dependencies() terra.Dependencies {
	return scc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SsmcontactsContactChannel].
func (scc *SsmcontactsContactChannel) LifecycleManagement() *terra.Lifecycle {
	return scc.Lifecycle
}

// Attributes returns the attributes for [SsmcontactsContactChannel].
func (scc *SsmcontactsContactChannel) Attributes() ssmcontactsContactChannelAttributes {
	return ssmcontactsContactChannelAttributes{ref: terra.ReferenceResource(scc)}
}

// ImportState imports the given attribute values into [SsmcontactsContactChannel]'s state.
func (scc *SsmcontactsContactChannel) ImportState(av io.Reader) error {
	scc.state = &ssmcontactsContactChannelState{}
	if err := json.NewDecoder(av).Decode(scc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scc.Type(), scc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SsmcontactsContactChannel] has state.
func (scc *SsmcontactsContactChannel) State() (*ssmcontactsContactChannelState, bool) {
	return scc.state, scc.state != nil
}

// StateMust returns the state for [SsmcontactsContactChannel]. Panics if the state is nil.
func (scc *SsmcontactsContactChannel) StateMust() *ssmcontactsContactChannelState {
	if scc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scc.Type(), scc.LocalName()))
	}
	return scc.state
}

// SsmcontactsContactChannelArgs contains the configurations for aws_ssmcontacts_contact_channel.
type SsmcontactsContactChannelArgs struct {
	// ContactId: string, required
	ContactId terra.StringValue `hcl:"contact_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// DeliveryAddress: required
	DeliveryAddress *ssmcontactscontactchannel.DeliveryAddress `hcl:"delivery_address,block" validate:"required"`
}
type ssmcontactsContactChannelAttributes struct {
	ref terra.Reference
}

// ActivationStatus returns a reference to field activation_status of aws_ssmcontacts_contact_channel.
func (scc ssmcontactsContactChannelAttributes) ActivationStatus() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("activation_status"))
}

// Arn returns a reference to field arn of aws_ssmcontacts_contact_channel.
func (scc ssmcontactsContactChannelAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("arn"))
}

// ContactId returns a reference to field contact_id of aws_ssmcontacts_contact_channel.
func (scc ssmcontactsContactChannelAttributes) ContactId() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("contact_id"))
}

// Id returns a reference to field id of aws_ssmcontacts_contact_channel.
func (scc ssmcontactsContactChannelAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("id"))
}

// Name returns a reference to field name of aws_ssmcontacts_contact_channel.
func (scc ssmcontactsContactChannelAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("name"))
}

// Type returns a reference to field type of aws_ssmcontacts_contact_channel.
func (scc ssmcontactsContactChannelAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(scc.ref.Append("type"))
}

func (scc ssmcontactsContactChannelAttributes) DeliveryAddress() terra.ListValue[ssmcontactscontactchannel.DeliveryAddressAttributes] {
	return terra.ReferenceAsList[ssmcontactscontactchannel.DeliveryAddressAttributes](scc.ref.Append("delivery_address"))
}

type ssmcontactsContactChannelState struct {
	ActivationStatus string                                           `json:"activation_status"`
	Arn              string                                           `json:"arn"`
	ContactId        string                                           `json:"contact_id"`
	Id               string                                           `json:"id"`
	Name             string                                           `json:"name"`
	Type             string                                           `json:"type"`
	DeliveryAddress  []ssmcontactscontactchannel.DeliveryAddressState `json:"delivery_address"`
}
