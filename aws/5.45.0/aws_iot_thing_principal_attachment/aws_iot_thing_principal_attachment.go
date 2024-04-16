// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_iot_thing_principal_attachment

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

// Resource represents the Terraform resource aws_iot_thing_principal_attachment.
type Resource struct {
	Name      string
	Args      Args
	state     *awsIotThingPrincipalAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aitpa *Resource) Type() string {
	return "aws_iot_thing_principal_attachment"
}

// LocalName returns the local name for [Resource].
func (aitpa *Resource) LocalName() string {
	return aitpa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aitpa *Resource) Configuration() interface{} {
	return aitpa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aitpa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aitpa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aitpa *Resource) Dependencies() terra.Dependencies {
	return aitpa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aitpa *Resource) LifecycleManagement() *terra.Lifecycle {
	return aitpa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aitpa *Resource) Attributes() awsIotThingPrincipalAttachmentAttributes {
	return awsIotThingPrincipalAttachmentAttributes{ref: terra.ReferenceResource(aitpa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aitpa *Resource) ImportState(state io.Reader) error {
	aitpa.state = &awsIotThingPrincipalAttachmentState{}
	if err := json.NewDecoder(state).Decode(aitpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aitpa.Type(), aitpa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aitpa *Resource) State() (*awsIotThingPrincipalAttachmentState, bool) {
	return aitpa.state, aitpa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aitpa *Resource) StateMust() *awsIotThingPrincipalAttachmentState {
	if aitpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aitpa.Type(), aitpa.LocalName()))
	}
	return aitpa.state
}

// Args contains the configurations for aws_iot_thing_principal_attachment.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Principal: string, required
	Principal terra.StringValue `hcl:"principal,attr" validate:"required"`
	// Thing: string, required
	Thing terra.StringValue `hcl:"thing,attr" validate:"required"`
}

type awsIotThingPrincipalAttachmentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_iot_thing_principal_attachment.
func (aitpa awsIotThingPrincipalAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aitpa.ref.Append("id"))
}

// Principal returns a reference to field principal of aws_iot_thing_principal_attachment.
func (aitpa awsIotThingPrincipalAttachmentAttributes) Principal() terra.StringValue {
	return terra.ReferenceAsString(aitpa.ref.Append("principal"))
}

// Thing returns a reference to field thing of aws_iot_thing_principal_attachment.
func (aitpa awsIotThingPrincipalAttachmentAttributes) Thing() terra.StringValue {
	return terra.ReferenceAsString(aitpa.ref.Append("thing"))
}

type awsIotThingPrincipalAttachmentState struct {
	Id        string `json:"id"`
	Principal string `json:"principal"`
	Thing     string `json:"thing"`
}
