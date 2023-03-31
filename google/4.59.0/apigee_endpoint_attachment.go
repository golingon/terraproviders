// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	apigeeendpointattachment "github.com/golingon/terraproviders/google/4.59.0/apigeeendpointattachment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewApigeeEndpointAttachment(name string, args ApigeeEndpointAttachmentArgs) *ApigeeEndpointAttachment {
	return &ApigeeEndpointAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApigeeEndpointAttachment)(nil)

type ApigeeEndpointAttachment struct {
	Name  string
	Args  ApigeeEndpointAttachmentArgs
	state *apigeeEndpointAttachmentState
}

func (aea *ApigeeEndpointAttachment) Type() string {
	return "google_apigee_endpoint_attachment"
}

func (aea *ApigeeEndpointAttachment) LocalName() string {
	return aea.Name
}

func (aea *ApigeeEndpointAttachment) Configuration() interface{} {
	return aea.Args
}

func (aea *ApigeeEndpointAttachment) Attributes() apigeeEndpointAttachmentAttributes {
	return apigeeEndpointAttachmentAttributes{ref: terra.ReferenceResource(aea)}
}

func (aea *ApigeeEndpointAttachment) ImportState(av io.Reader) error {
	aea.state = &apigeeEndpointAttachmentState{}
	if err := json.NewDecoder(av).Decode(aea.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aea.Type(), aea.LocalName(), err)
	}
	return nil
}

func (aea *ApigeeEndpointAttachment) State() (*apigeeEndpointAttachmentState, bool) {
	return aea.state, aea.state != nil
}

func (aea *ApigeeEndpointAttachment) StateMust() *apigeeEndpointAttachmentState {
	if aea.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aea.Type(), aea.LocalName()))
	}
	return aea.state
}

func (aea *ApigeeEndpointAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(aea)
}

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
	// DependsOn contains resources that ApigeeEndpointAttachment depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type apigeeEndpointAttachmentAttributes struct {
	ref terra.Reference
}

func (aea apigeeEndpointAttachmentAttributes) ConnectionState() terra.StringValue {
	return terra.ReferenceString(aea.ref.Append("connection_state"))
}

func (aea apigeeEndpointAttachmentAttributes) EndpointAttachmentId() terra.StringValue {
	return terra.ReferenceString(aea.ref.Append("endpoint_attachment_id"))
}

func (aea apigeeEndpointAttachmentAttributes) Host() terra.StringValue {
	return terra.ReferenceString(aea.ref.Append("host"))
}

func (aea apigeeEndpointAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceString(aea.ref.Append("id"))
}

func (aea apigeeEndpointAttachmentAttributes) Location() terra.StringValue {
	return terra.ReferenceString(aea.ref.Append("location"))
}

func (aea apigeeEndpointAttachmentAttributes) Name() terra.StringValue {
	return terra.ReferenceString(aea.ref.Append("name"))
}

func (aea apigeeEndpointAttachmentAttributes) OrgId() terra.StringValue {
	return terra.ReferenceString(aea.ref.Append("org_id"))
}

func (aea apigeeEndpointAttachmentAttributes) ServiceAttachment() terra.StringValue {
	return terra.ReferenceString(aea.ref.Append("service_attachment"))
}

func (aea apigeeEndpointAttachmentAttributes) Timeouts() apigeeendpointattachment.TimeoutsAttributes {
	return terra.ReferenceSingle[apigeeendpointattachment.TimeoutsAttributes](aea.ref.Append("timeouts"))
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
