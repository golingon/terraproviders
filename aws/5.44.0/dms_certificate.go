// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewDmsCertificate creates a new instance of [DmsCertificate].
func NewDmsCertificate(name string, args DmsCertificateArgs) *DmsCertificate {
	return &DmsCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DmsCertificate)(nil)

// DmsCertificate represents the Terraform resource aws_dms_certificate.
type DmsCertificate struct {
	Name      string
	Args      DmsCertificateArgs
	state     *dmsCertificateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DmsCertificate].
func (dc *DmsCertificate) Type() string {
	return "aws_dms_certificate"
}

// LocalName returns the local name for [DmsCertificate].
func (dc *DmsCertificate) LocalName() string {
	return dc.Name
}

// Configuration returns the configuration (args) for [DmsCertificate].
func (dc *DmsCertificate) Configuration() interface{} {
	return dc.Args
}

// DependOn is used for other resources to depend on [DmsCertificate].
func (dc *DmsCertificate) DependOn() terra.Reference {
	return terra.ReferenceResource(dc)
}

// Dependencies returns the list of resources [DmsCertificate] depends_on.
func (dc *DmsCertificate) Dependencies() terra.Dependencies {
	return dc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DmsCertificate].
func (dc *DmsCertificate) LifecycleManagement() *terra.Lifecycle {
	return dc.Lifecycle
}

// Attributes returns the attributes for [DmsCertificate].
func (dc *DmsCertificate) Attributes() dmsCertificateAttributes {
	return dmsCertificateAttributes{ref: terra.ReferenceResource(dc)}
}

// ImportState imports the given attribute values into [DmsCertificate]'s state.
func (dc *DmsCertificate) ImportState(av io.Reader) error {
	dc.state = &dmsCertificateState{}
	if err := json.NewDecoder(av).Decode(dc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dc.Type(), dc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DmsCertificate] has state.
func (dc *DmsCertificate) State() (*dmsCertificateState, bool) {
	return dc.state, dc.state != nil
}

// StateMust returns the state for [DmsCertificate]. Panics if the state is nil.
func (dc *DmsCertificate) StateMust() *dmsCertificateState {
	if dc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dc.Type(), dc.LocalName()))
	}
	return dc.state
}

// DmsCertificateArgs contains the configurations for aws_dms_certificate.
type DmsCertificateArgs struct {
	// CertificateId: string, required
	CertificateId terra.StringValue `hcl:"certificate_id,attr" validate:"required"`
	// CertificatePem: string, optional
	CertificatePem terra.StringValue `hcl:"certificate_pem,attr"`
	// CertificateWallet: string, optional
	CertificateWallet terra.StringValue `hcl:"certificate_wallet,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type dmsCertificateAttributes struct {
	ref terra.Reference
}

// CertificateArn returns a reference to field certificate_arn of aws_dms_certificate.
func (dc dmsCertificateAttributes) CertificateArn() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("certificate_arn"))
}

// CertificateId returns a reference to field certificate_id of aws_dms_certificate.
func (dc dmsCertificateAttributes) CertificateId() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("certificate_id"))
}

// CertificatePem returns a reference to field certificate_pem of aws_dms_certificate.
func (dc dmsCertificateAttributes) CertificatePem() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("certificate_pem"))
}

// CertificateWallet returns a reference to field certificate_wallet of aws_dms_certificate.
func (dc dmsCertificateAttributes) CertificateWallet() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("certificate_wallet"))
}

// Id returns a reference to field id of aws_dms_certificate.
func (dc dmsCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dc.ref.Append("id"))
}

// Tags returns a reference to field tags of aws_dms_certificate.
func (dc dmsCertificateAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_dms_certificate.
func (dc dmsCertificateAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dc.ref.Append("tags_all"))
}

type dmsCertificateState struct {
	CertificateArn    string            `json:"certificate_arn"`
	CertificateId     string            `json:"certificate_id"`
	CertificatePem    string            `json:"certificate_pem"`
	CertificateWallet string            `json:"certificate_wallet"`
	Id                string            `json:"id"`
	Tags              map[string]string `json:"tags"`
	TagsAll           map[string]string `json:"tags_all"`
}
