// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	apigeeinstanceattachment "github.com/golingon/terraproviders/google/4.75.1/apigeeinstanceattachment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApigeeInstanceAttachment creates a new instance of [ApigeeInstanceAttachment].
func NewApigeeInstanceAttachment(name string, args ApigeeInstanceAttachmentArgs) *ApigeeInstanceAttachment {
	return &ApigeeInstanceAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApigeeInstanceAttachment)(nil)

// ApigeeInstanceAttachment represents the Terraform resource google_apigee_instance_attachment.
type ApigeeInstanceAttachment struct {
	Name      string
	Args      ApigeeInstanceAttachmentArgs
	state     *apigeeInstanceAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApigeeInstanceAttachment].
func (aia *ApigeeInstanceAttachment) Type() string {
	return "google_apigee_instance_attachment"
}

// LocalName returns the local name for [ApigeeInstanceAttachment].
func (aia *ApigeeInstanceAttachment) LocalName() string {
	return aia.Name
}

// Configuration returns the configuration (args) for [ApigeeInstanceAttachment].
func (aia *ApigeeInstanceAttachment) Configuration() interface{} {
	return aia.Args
}

// DependOn is used for other resources to depend on [ApigeeInstanceAttachment].
func (aia *ApigeeInstanceAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(aia)
}

// Dependencies returns the list of resources [ApigeeInstanceAttachment] depends_on.
func (aia *ApigeeInstanceAttachment) Dependencies() terra.Dependencies {
	return aia.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApigeeInstanceAttachment].
func (aia *ApigeeInstanceAttachment) LifecycleManagement() *terra.Lifecycle {
	return aia.Lifecycle
}

// Attributes returns the attributes for [ApigeeInstanceAttachment].
func (aia *ApigeeInstanceAttachment) Attributes() apigeeInstanceAttachmentAttributes {
	return apigeeInstanceAttachmentAttributes{ref: terra.ReferenceResource(aia)}
}

// ImportState imports the given attribute values into [ApigeeInstanceAttachment]'s state.
func (aia *ApigeeInstanceAttachment) ImportState(av io.Reader) error {
	aia.state = &apigeeInstanceAttachmentState{}
	if err := json.NewDecoder(av).Decode(aia.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aia.Type(), aia.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApigeeInstanceAttachment] has state.
func (aia *ApigeeInstanceAttachment) State() (*apigeeInstanceAttachmentState, bool) {
	return aia.state, aia.state != nil
}

// StateMust returns the state for [ApigeeInstanceAttachment]. Panics if the state is nil.
func (aia *ApigeeInstanceAttachment) StateMust() *apigeeInstanceAttachmentState {
	if aia.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aia.Type(), aia.LocalName()))
	}
	return aia.state
}

// ApigeeInstanceAttachmentArgs contains the configurations for google_apigee_instance_attachment.
type ApigeeInstanceAttachmentArgs struct {
	// Environment: string, required
	Environment terra.StringValue `hcl:"environment,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *apigeeinstanceattachment.Timeouts `hcl:"timeouts,block"`
}
type apigeeInstanceAttachmentAttributes struct {
	ref terra.Reference
}

// Environment returns a reference to field environment of google_apigee_instance_attachment.
func (aia apigeeInstanceAttachmentAttributes) Environment() terra.StringValue {
	return terra.ReferenceAsString(aia.ref.Append("environment"))
}

// Id returns a reference to field id of google_apigee_instance_attachment.
func (aia apigeeInstanceAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aia.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of google_apigee_instance_attachment.
func (aia apigeeInstanceAttachmentAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(aia.ref.Append("instance_id"))
}

// Name returns a reference to field name of google_apigee_instance_attachment.
func (aia apigeeInstanceAttachmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aia.ref.Append("name"))
}

func (aia apigeeInstanceAttachmentAttributes) Timeouts() apigeeinstanceattachment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apigeeinstanceattachment.TimeoutsAttributes](aia.ref.Append("timeouts"))
}

type apigeeInstanceAttachmentState struct {
	Environment string                                  `json:"environment"`
	Id          string                                  `json:"id"`
	InstanceId  string                                  `json:"instance_id"`
	Name        string                                  `json:"name"`
	Timeouts    *apigeeinstanceattachment.TimeoutsState `json:"timeouts"`
}
