// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	apigeeendpointattachment "github.com/golingon/terraproviders/googlebeta/4.64.0/apigeeendpointattachment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApigeeEndpointAttachment creates a new instance of [ApigeeEndpointAttachment].
func NewApigeeEndpointAttachment(name string, args ApigeeEndpointAttachmentArgs) *ApigeeEndpointAttachment {
	return &ApigeeEndpointAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApigeeEndpointAttachment)(nil)

// ApigeeEndpointAttachment represents the Terraform resource google_apigee_endpoint_attachment.
type ApigeeEndpointAttachment struct {
	Name      string
	Args      ApigeeEndpointAttachmentArgs
	state     *apigeeEndpointAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApigeeEndpointAttachment].
func (aea *ApigeeEndpointAttachment) Type() string {
	return "google_apigee_endpoint_attachment"
}

// LocalName returns the local name for [ApigeeEndpointAttachment].
func (aea *ApigeeEndpointAttachment) LocalName() string {
	return aea.Name
}

// Configuration returns the configuration (args) for [ApigeeEndpointAttachment].
func (aea *ApigeeEndpointAttachment) Configuration() interface{} {
	return aea.Args
}

// DependOn is used for other resources to depend on [ApigeeEndpointAttachment].
func (aea *ApigeeEndpointAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(aea)
}

// Dependencies returns the list of resources [ApigeeEndpointAttachment] depends_on.
func (aea *ApigeeEndpointAttachment) Dependencies() terra.Dependencies {
	return aea.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApigeeEndpointAttachment].
func (aea *ApigeeEndpointAttachment) LifecycleManagement() *terra.Lifecycle {
	return aea.Lifecycle
}

// Attributes returns the attributes for [ApigeeEndpointAttachment].
func (aea *ApigeeEndpointAttachment) Attributes() apigeeEndpointAttachmentAttributes {
	return apigeeEndpointAttachmentAttributes{ref: terra.ReferenceResource(aea)}
}

// ImportState imports the given attribute values into [ApigeeEndpointAttachment]'s state.
func (aea *ApigeeEndpointAttachment) ImportState(av io.Reader) error {
	aea.state = &apigeeEndpointAttachmentState{}
	if err := json.NewDecoder(av).Decode(aea.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aea.Type(), aea.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApigeeEndpointAttachment] has state.
func (aea *ApigeeEndpointAttachment) State() (*apigeeEndpointAttachmentState, bool) {
	return aea.state, aea.state != nil
}

// StateMust returns the state for [ApigeeEndpointAttachment]. Panics if the state is nil.
func (aea *ApigeeEndpointAttachment) StateMust() *apigeeEndpointAttachmentState {
	if aea.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aea.Type(), aea.LocalName()))
	}
	return aea.state
}

// ApigeeEndpointAttachmentArgs contains the configurations for google_apigee_endpoint_attachment.
type ApigeeEndpointAttachmentArgs struct {
	// EndpointAttachmentId: string, required
	EndpointAttachmentId terra.StringValue `hcl:"endpoint_attachment_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// OrgId: string, required
	OrgId terra.StringValue `hcl:"org_id,attr" validate:"required"`
	// ServiceAttachment: string, required
	ServiceAttachment terra.StringValue `hcl:"service_attachment,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *apigeeendpointattachment.Timeouts `hcl:"timeouts,block"`
}
type apigeeEndpointAttachmentAttributes struct {
	ref terra.Reference
}

// ConnectionState returns a reference to field connection_state of google_apigee_endpoint_attachment.
func (aea apigeeEndpointAttachmentAttributes) ConnectionState() terra.StringValue {
	return terra.ReferenceAsString(aea.ref.Append("connection_state"))
}

// EndpointAttachmentId returns a reference to field endpoint_attachment_id of google_apigee_endpoint_attachment.
func (aea apigeeEndpointAttachmentAttributes) EndpointAttachmentId() terra.StringValue {
	return terra.ReferenceAsString(aea.ref.Append("endpoint_attachment_id"))
}

// Host returns a reference to field host of google_apigee_endpoint_attachment.
func (aea apigeeEndpointAttachmentAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(aea.ref.Append("host"))
}

// Id returns a reference to field id of google_apigee_endpoint_attachment.
func (aea apigeeEndpointAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aea.ref.Append("id"))
}

// Location returns a reference to field location of google_apigee_endpoint_attachment.
func (aea apigeeEndpointAttachmentAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(aea.ref.Append("location"))
}

// Name returns a reference to field name of google_apigee_endpoint_attachment.
func (aea apigeeEndpointAttachmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aea.ref.Append("name"))
}

// OrgId returns a reference to field org_id of google_apigee_endpoint_attachment.
func (aea apigeeEndpointAttachmentAttributes) OrgId() terra.StringValue {
	return terra.ReferenceAsString(aea.ref.Append("org_id"))
}

// ServiceAttachment returns a reference to field service_attachment of google_apigee_endpoint_attachment.
func (aea apigeeEndpointAttachmentAttributes) ServiceAttachment() terra.StringValue {
	return terra.ReferenceAsString(aea.ref.Append("service_attachment"))
}

func (aea apigeeEndpointAttachmentAttributes) Timeouts() apigeeendpointattachment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apigeeendpointattachment.TimeoutsAttributes](aea.ref.Append("timeouts"))
}

type apigeeEndpointAttachmentState struct {
	ConnectionState      string                                  `json:"connection_state"`
	EndpointAttachmentId string                                  `json:"endpoint_attachment_id"`
	Host                 string                                  `json:"host"`
	Id                   string                                  `json:"id"`
	Location             string                                  `json:"location"`
	Name                 string                                  `json:"name"`
	OrgId                string                                  `json:"org_id"`
	ServiceAttachment    string                                  `json:"service_attachment"`
	Timeouts             *apigeeendpointattachment.TimeoutsState `json:"timeouts"`
}
