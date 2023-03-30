// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewLightsailLbCertificateAttachment(name string, args LightsailLbCertificateAttachmentArgs) *LightsailLbCertificateAttachment {
	return &LightsailLbCertificateAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LightsailLbCertificateAttachment)(nil)

type LightsailLbCertificateAttachment struct {
	Name  string
	Args  LightsailLbCertificateAttachmentArgs
	state *lightsailLbCertificateAttachmentState
}

func (llca *LightsailLbCertificateAttachment) Type() string {
	return "aws_lightsail_lb_certificate_attachment"
}

func (llca *LightsailLbCertificateAttachment) LocalName() string {
	return llca.Name
}

func (llca *LightsailLbCertificateAttachment) Configuration() interface{} {
	return llca.Args
}

func (llca *LightsailLbCertificateAttachment) Attributes() lightsailLbCertificateAttachmentAttributes {
	return lightsailLbCertificateAttachmentAttributes{ref: terra.ReferenceResource(llca)}
}

func (llca *LightsailLbCertificateAttachment) ImportState(av io.Reader) error {
	llca.state = &lightsailLbCertificateAttachmentState{}
	if err := json.NewDecoder(av).Decode(llca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", llca.Type(), llca.LocalName(), err)
	}
	return nil
}

func (llca *LightsailLbCertificateAttachment) State() (*lightsailLbCertificateAttachmentState, bool) {
	return llca.state, llca.state != nil
}

func (llca *LightsailLbCertificateAttachment) StateMust() *lightsailLbCertificateAttachmentState {
	if llca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", llca.Type(), llca.LocalName()))
	}
	return llca.state
}

func (llca *LightsailLbCertificateAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(llca)
}

type LightsailLbCertificateAttachmentArgs struct {
	// CertificateName: string, required
	CertificateName terra.StringValue `hcl:"certificate_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LbName: string, required
	LbName terra.StringValue `hcl:"lb_name,attr" validate:"required"`
	// DependsOn contains resources that LightsailLbCertificateAttachment depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type lightsailLbCertificateAttachmentAttributes struct {
	ref terra.Reference
}

func (llca lightsailLbCertificateAttachmentAttributes) CertificateName() terra.StringValue {
	return terra.ReferenceString(llca.ref.Append("certificate_name"))
}

func (llca lightsailLbCertificateAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceString(llca.ref.Append("id"))
}

func (llca lightsailLbCertificateAttachmentAttributes) LbName() terra.StringValue {
	return terra.ReferenceString(llca.ref.Append("lb_name"))
}

type lightsailLbCertificateAttachmentState struct {
	CertificateName string `json:"certificate_name"`
	Id              string `json:"id"`
	LbName          string `json:"lb_name"`
}
