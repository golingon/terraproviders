// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewIotThingPrincipalAttachment creates a new instance of [IotThingPrincipalAttachment].
func NewIotThingPrincipalAttachment(name string, args IotThingPrincipalAttachmentArgs) *IotThingPrincipalAttachment {
	return &IotThingPrincipalAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IotThingPrincipalAttachment)(nil)

// IotThingPrincipalAttachment represents the Terraform resource aws_iot_thing_principal_attachment.
type IotThingPrincipalAttachment struct {
	Name      string
	Args      IotThingPrincipalAttachmentArgs
	state     *iotThingPrincipalAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IotThingPrincipalAttachment].
func (itpa *IotThingPrincipalAttachment) Type() string {
	return "aws_iot_thing_principal_attachment"
}

// LocalName returns the local name for [IotThingPrincipalAttachment].
func (itpa *IotThingPrincipalAttachment) LocalName() string {
	return itpa.Name
}

// Configuration returns the configuration (args) for [IotThingPrincipalAttachment].
func (itpa *IotThingPrincipalAttachment) Configuration() interface{} {
	return itpa.Args
}

// DependOn is used for other resources to depend on [IotThingPrincipalAttachment].
func (itpa *IotThingPrincipalAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(itpa)
}

// Dependencies returns the list of resources [IotThingPrincipalAttachment] depends_on.
func (itpa *IotThingPrincipalAttachment) Dependencies() terra.Dependencies {
	return itpa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IotThingPrincipalAttachment].
func (itpa *IotThingPrincipalAttachment) LifecycleManagement() *terra.Lifecycle {
	return itpa.Lifecycle
}

// Attributes returns the attributes for [IotThingPrincipalAttachment].
func (itpa *IotThingPrincipalAttachment) Attributes() iotThingPrincipalAttachmentAttributes {
	return iotThingPrincipalAttachmentAttributes{ref: terra.ReferenceResource(itpa)}
}

// ImportState imports the given attribute values into [IotThingPrincipalAttachment]'s state.
func (itpa *IotThingPrincipalAttachment) ImportState(av io.Reader) error {
	itpa.state = &iotThingPrincipalAttachmentState{}
	if err := json.NewDecoder(av).Decode(itpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", itpa.Type(), itpa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IotThingPrincipalAttachment] has state.
func (itpa *IotThingPrincipalAttachment) State() (*iotThingPrincipalAttachmentState, bool) {
	return itpa.state, itpa.state != nil
}

// StateMust returns the state for [IotThingPrincipalAttachment]. Panics if the state is nil.
func (itpa *IotThingPrincipalAttachment) StateMust() *iotThingPrincipalAttachmentState {
	if itpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", itpa.Type(), itpa.LocalName()))
	}
	return itpa.state
}

// IotThingPrincipalAttachmentArgs contains the configurations for aws_iot_thing_principal_attachment.
type IotThingPrincipalAttachmentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Principal: string, required
	Principal terra.StringValue `hcl:"principal,attr" validate:"required"`
	// Thing: string, required
	Thing terra.StringValue `hcl:"thing,attr" validate:"required"`
}
type iotThingPrincipalAttachmentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_iot_thing_principal_attachment.
func (itpa iotThingPrincipalAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(itpa.ref.Append("id"))
}

// Principal returns a reference to field principal of aws_iot_thing_principal_attachment.
func (itpa iotThingPrincipalAttachmentAttributes) Principal() terra.StringValue {
	return terra.ReferenceAsString(itpa.ref.Append("principal"))
}

// Thing returns a reference to field thing of aws_iot_thing_principal_attachment.
func (itpa iotThingPrincipalAttachmentAttributes) Thing() terra.StringValue {
	return terra.ReferenceAsString(itpa.ref.Append("thing"))
}

type iotThingPrincipalAttachmentState struct {
	Id        string `json:"id"`
	Principal string `json:"principal"`
	Thing     string `json:"thing"`
}
