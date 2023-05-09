// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	apigeeenvgroupattachment "github.com/golingon/terraproviders/google/4.64.0/apigeeenvgroupattachment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApigeeEnvgroupAttachment creates a new instance of [ApigeeEnvgroupAttachment].
func NewApigeeEnvgroupAttachment(name string, args ApigeeEnvgroupAttachmentArgs) *ApigeeEnvgroupAttachment {
	return &ApigeeEnvgroupAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApigeeEnvgroupAttachment)(nil)

// ApigeeEnvgroupAttachment represents the Terraform resource google_apigee_envgroup_attachment.
type ApigeeEnvgroupAttachment struct {
	Name      string
	Args      ApigeeEnvgroupAttachmentArgs
	state     *apigeeEnvgroupAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApigeeEnvgroupAttachment].
func (aea *ApigeeEnvgroupAttachment) Type() string {
	return "google_apigee_envgroup_attachment"
}

// LocalName returns the local name for [ApigeeEnvgroupAttachment].
func (aea *ApigeeEnvgroupAttachment) LocalName() string {
	return aea.Name
}

// Configuration returns the configuration (args) for [ApigeeEnvgroupAttachment].
func (aea *ApigeeEnvgroupAttachment) Configuration() interface{} {
	return aea.Args
}

// DependOn is used for other resources to depend on [ApigeeEnvgroupAttachment].
func (aea *ApigeeEnvgroupAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(aea)
}

// Dependencies returns the list of resources [ApigeeEnvgroupAttachment] depends_on.
func (aea *ApigeeEnvgroupAttachment) Dependencies() terra.Dependencies {
	return aea.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApigeeEnvgroupAttachment].
func (aea *ApigeeEnvgroupAttachment) LifecycleManagement() *terra.Lifecycle {
	return aea.Lifecycle
}

// Attributes returns the attributes for [ApigeeEnvgroupAttachment].
func (aea *ApigeeEnvgroupAttachment) Attributes() apigeeEnvgroupAttachmentAttributes {
	return apigeeEnvgroupAttachmentAttributes{ref: terra.ReferenceResource(aea)}
}

// ImportState imports the given attribute values into [ApigeeEnvgroupAttachment]'s state.
func (aea *ApigeeEnvgroupAttachment) ImportState(av io.Reader) error {
	aea.state = &apigeeEnvgroupAttachmentState{}
	if err := json.NewDecoder(av).Decode(aea.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aea.Type(), aea.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApigeeEnvgroupAttachment] has state.
func (aea *ApigeeEnvgroupAttachment) State() (*apigeeEnvgroupAttachmentState, bool) {
	return aea.state, aea.state != nil
}

// StateMust returns the state for [ApigeeEnvgroupAttachment]. Panics if the state is nil.
func (aea *ApigeeEnvgroupAttachment) StateMust() *apigeeEnvgroupAttachmentState {
	if aea.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aea.Type(), aea.LocalName()))
	}
	return aea.state
}

// ApigeeEnvgroupAttachmentArgs contains the configurations for google_apigee_envgroup_attachment.
type ApigeeEnvgroupAttachmentArgs struct {
	// EnvgroupId: string, required
	EnvgroupId terra.StringValue `hcl:"envgroup_id,attr" validate:"required"`
	// Environment: string, required
	Environment terra.StringValue `hcl:"environment,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *apigeeenvgroupattachment.Timeouts `hcl:"timeouts,block"`
}
type apigeeEnvgroupAttachmentAttributes struct {
	ref terra.Reference
}

// EnvgroupId returns a reference to field envgroup_id of google_apigee_envgroup_attachment.
func (aea apigeeEnvgroupAttachmentAttributes) EnvgroupId() terra.StringValue {
	return terra.ReferenceAsString(aea.ref.Append("envgroup_id"))
}

// Environment returns a reference to field environment of google_apigee_envgroup_attachment.
func (aea apigeeEnvgroupAttachmentAttributes) Environment() terra.StringValue {
	return terra.ReferenceAsString(aea.ref.Append("environment"))
}

// Id returns a reference to field id of google_apigee_envgroup_attachment.
func (aea apigeeEnvgroupAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aea.ref.Append("id"))
}

// Name returns a reference to field name of google_apigee_envgroup_attachment.
func (aea apigeeEnvgroupAttachmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aea.ref.Append("name"))
}

func (aea apigeeEnvgroupAttachmentAttributes) Timeouts() apigeeenvgroupattachment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apigeeenvgroupattachment.TimeoutsAttributes](aea.ref.Append("timeouts"))
}

type apigeeEnvgroupAttachmentState struct {
	EnvgroupId  string                                  `json:"envgroup_id"`
	Environment string                                  `json:"environment"`
	Id          string                                  `json:"id"`
	Name        string                                  `json:"name"`
	Timeouts    *apigeeenvgroupattachment.TimeoutsState `json:"timeouts"`
}
