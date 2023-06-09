// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTransferCertificate creates a new instance of [TransferCertificate].
func NewTransferCertificate(name string, args TransferCertificateArgs) *TransferCertificate {
	return &TransferCertificate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TransferCertificate)(nil)

// TransferCertificate represents the Terraform resource aws_transfer_certificate.
type TransferCertificate struct {
	Name      string
	Args      TransferCertificateArgs
	state     *transferCertificateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TransferCertificate].
func (tc *TransferCertificate) Type() string {
	return "aws_transfer_certificate"
}

// LocalName returns the local name for [TransferCertificate].
func (tc *TransferCertificate) LocalName() string {
	return tc.Name
}

// Configuration returns the configuration (args) for [TransferCertificate].
func (tc *TransferCertificate) Configuration() interface{} {
	return tc.Args
}

// DependOn is used for other resources to depend on [TransferCertificate].
func (tc *TransferCertificate) DependOn() terra.Reference {
	return terra.ReferenceResource(tc)
}

// Dependencies returns the list of resources [TransferCertificate] depends_on.
func (tc *TransferCertificate) Dependencies() terra.Dependencies {
	return tc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TransferCertificate].
func (tc *TransferCertificate) LifecycleManagement() *terra.Lifecycle {
	return tc.Lifecycle
}

// Attributes returns the attributes for [TransferCertificate].
func (tc *TransferCertificate) Attributes() transferCertificateAttributes {
	return transferCertificateAttributes{ref: terra.ReferenceResource(tc)}
}

// ImportState imports the given attribute values into [TransferCertificate]'s state.
func (tc *TransferCertificate) ImportState(av io.Reader) error {
	tc.state = &transferCertificateState{}
	if err := json.NewDecoder(av).Decode(tc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", tc.Type(), tc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TransferCertificate] has state.
func (tc *TransferCertificate) State() (*transferCertificateState, bool) {
	return tc.state, tc.state != nil
}

// StateMust returns the state for [TransferCertificate]. Panics if the state is nil.
func (tc *TransferCertificate) StateMust() *transferCertificateState {
	if tc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", tc.Type(), tc.LocalName()))
	}
	return tc.state
}

// TransferCertificateArgs contains the configurations for aws_transfer_certificate.
type TransferCertificateArgs struct {
	// Certificate: string, required
	Certificate terra.StringValue `hcl:"certificate,attr" validate:"required"`
	// CertificateChain: string, optional
	CertificateChain terra.StringValue `hcl:"certificate_chain,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PrivateKey: string, optional
	PrivateKey terra.StringValue `hcl:"private_key,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Usage: string, required
	Usage terra.StringValue `hcl:"usage,attr" validate:"required"`
}
type transferCertificateAttributes struct {
	ref terra.Reference
}

// ActiveDate returns a reference to field active_date of aws_transfer_certificate.
func (tc transferCertificateAttributes) ActiveDate() terra.StringValue {
	return terra.ReferenceAsString(tc.ref.Append("active_date"))
}

// Certificate returns a reference to field certificate of aws_transfer_certificate.
func (tc transferCertificateAttributes) Certificate() terra.StringValue {
	return terra.ReferenceAsString(tc.ref.Append("certificate"))
}

// CertificateChain returns a reference to field certificate_chain of aws_transfer_certificate.
func (tc transferCertificateAttributes) CertificateChain() terra.StringValue {
	return terra.ReferenceAsString(tc.ref.Append("certificate_chain"))
}

// CertificateId returns a reference to field certificate_id of aws_transfer_certificate.
func (tc transferCertificateAttributes) CertificateId() terra.StringValue {
	return terra.ReferenceAsString(tc.ref.Append("certificate_id"))
}

// Description returns a reference to field description of aws_transfer_certificate.
func (tc transferCertificateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(tc.ref.Append("description"))
}

// Id returns a reference to field id of aws_transfer_certificate.
func (tc transferCertificateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(tc.ref.Append("id"))
}

// InactiveDate returns a reference to field inactive_date of aws_transfer_certificate.
func (tc transferCertificateAttributes) InactiveDate() terra.StringValue {
	return terra.ReferenceAsString(tc.ref.Append("inactive_date"))
}

// PrivateKey returns a reference to field private_key of aws_transfer_certificate.
func (tc transferCertificateAttributes) PrivateKey() terra.StringValue {
	return terra.ReferenceAsString(tc.ref.Append("private_key"))
}

// Tags returns a reference to field tags of aws_transfer_certificate.
func (tc transferCertificateAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](tc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_transfer_certificate.
func (tc transferCertificateAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](tc.ref.Append("tags_all"))
}

// Usage returns a reference to field usage of aws_transfer_certificate.
func (tc transferCertificateAttributes) Usage() terra.StringValue {
	return terra.ReferenceAsString(tc.ref.Append("usage"))
}

type transferCertificateState struct {
	ActiveDate       string            `json:"active_date"`
	Certificate      string            `json:"certificate"`
	CertificateChain string            `json:"certificate_chain"`
	CertificateId    string            `json:"certificate_id"`
	Description      string            `json:"description"`
	Id               string            `json:"id"`
	InactiveDate     string            `json:"inactive_date"`
	PrivateKey       string            `json:"private_key"`
	Tags             map[string]string `json:"tags"`
	TagsAll          map[string]string `json:"tags_all"`
	Usage            string            `json:"usage"`
}
