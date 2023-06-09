// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLightsailLbCertificateAttachment creates a new instance of [LightsailLbCertificateAttachment].
func NewLightsailLbCertificateAttachment(name string, args LightsailLbCertificateAttachmentArgs) *LightsailLbCertificateAttachment {
	return &LightsailLbCertificateAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LightsailLbCertificateAttachment)(nil)

// LightsailLbCertificateAttachment represents the Terraform resource aws_lightsail_lb_certificate_attachment.
type LightsailLbCertificateAttachment struct {
	Name      string
	Args      LightsailLbCertificateAttachmentArgs
	state     *lightsailLbCertificateAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LightsailLbCertificateAttachment].
func (llca *LightsailLbCertificateAttachment) Type() string {
	return "aws_lightsail_lb_certificate_attachment"
}

// LocalName returns the local name for [LightsailLbCertificateAttachment].
func (llca *LightsailLbCertificateAttachment) LocalName() string {
	return llca.Name
}

// Configuration returns the configuration (args) for [LightsailLbCertificateAttachment].
func (llca *LightsailLbCertificateAttachment) Configuration() interface{} {
	return llca.Args
}

// DependOn is used for other resources to depend on [LightsailLbCertificateAttachment].
func (llca *LightsailLbCertificateAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(llca)
}

// Dependencies returns the list of resources [LightsailLbCertificateAttachment] depends_on.
func (llca *LightsailLbCertificateAttachment) Dependencies() terra.Dependencies {
	return llca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LightsailLbCertificateAttachment].
func (llca *LightsailLbCertificateAttachment) LifecycleManagement() *terra.Lifecycle {
	return llca.Lifecycle
}

// Attributes returns the attributes for [LightsailLbCertificateAttachment].
func (llca *LightsailLbCertificateAttachment) Attributes() lightsailLbCertificateAttachmentAttributes {
	return lightsailLbCertificateAttachmentAttributes{ref: terra.ReferenceResource(llca)}
}

// ImportState imports the given attribute values into [LightsailLbCertificateAttachment]'s state.
func (llca *LightsailLbCertificateAttachment) ImportState(av io.Reader) error {
	llca.state = &lightsailLbCertificateAttachmentState{}
	if err := json.NewDecoder(av).Decode(llca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", llca.Type(), llca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LightsailLbCertificateAttachment] has state.
func (llca *LightsailLbCertificateAttachment) State() (*lightsailLbCertificateAttachmentState, bool) {
	return llca.state, llca.state != nil
}

// StateMust returns the state for [LightsailLbCertificateAttachment]. Panics if the state is nil.
func (llca *LightsailLbCertificateAttachment) StateMust() *lightsailLbCertificateAttachmentState {
	if llca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", llca.Type(), llca.LocalName()))
	}
	return llca.state
}

// LightsailLbCertificateAttachmentArgs contains the configurations for aws_lightsail_lb_certificate_attachment.
type LightsailLbCertificateAttachmentArgs struct {
	// CertificateName: string, required
	CertificateName terra.StringValue `hcl:"certificate_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LbName: string, required
	LbName terra.StringValue `hcl:"lb_name,attr" validate:"required"`
}
type lightsailLbCertificateAttachmentAttributes struct {
	ref terra.Reference
}

// CertificateName returns a reference to field certificate_name of aws_lightsail_lb_certificate_attachment.
func (llca lightsailLbCertificateAttachmentAttributes) CertificateName() terra.StringValue {
	return terra.ReferenceAsString(llca.ref.Append("certificate_name"))
}

// Id returns a reference to field id of aws_lightsail_lb_certificate_attachment.
func (llca lightsailLbCertificateAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(llca.ref.Append("id"))
}

// LbName returns a reference to field lb_name of aws_lightsail_lb_certificate_attachment.
func (llca lightsailLbCertificateAttachmentAttributes) LbName() terra.StringValue {
	return terra.ReferenceAsString(llca.ref.Append("lb_name"))
}

type lightsailLbCertificateAttachmentState struct {
	CertificateName string `json:"certificate_name"`
	Id              string `json:"id"`
	LbName          string `json:"lb_name"`
}
