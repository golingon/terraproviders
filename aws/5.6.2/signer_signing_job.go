// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	signersigningjob "github.com/golingon/terraproviders/aws/5.6.2/signersigningjob"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSignerSigningJob creates a new instance of [SignerSigningJob].
func NewSignerSigningJob(name string, args SignerSigningJobArgs) *SignerSigningJob {
	return &SignerSigningJob{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SignerSigningJob)(nil)

// SignerSigningJob represents the Terraform resource aws_signer_signing_job.
type SignerSigningJob struct {
	Name      string
	Args      SignerSigningJobArgs
	state     *signerSigningJobState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SignerSigningJob].
func (ssj *SignerSigningJob) Type() string {
	return "aws_signer_signing_job"
}

// LocalName returns the local name for [SignerSigningJob].
func (ssj *SignerSigningJob) LocalName() string {
	return ssj.Name
}

// Configuration returns the configuration (args) for [SignerSigningJob].
func (ssj *SignerSigningJob) Configuration() interface{} {
	return ssj.Args
}

// DependOn is used for other resources to depend on [SignerSigningJob].
func (ssj *SignerSigningJob) DependOn() terra.Reference {
	return terra.ReferenceResource(ssj)
}

// Dependencies returns the list of resources [SignerSigningJob] depends_on.
func (ssj *SignerSigningJob) Dependencies() terra.Dependencies {
	return ssj.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SignerSigningJob].
func (ssj *SignerSigningJob) LifecycleManagement() *terra.Lifecycle {
	return ssj.Lifecycle
}

// Attributes returns the attributes for [SignerSigningJob].
func (ssj *SignerSigningJob) Attributes() signerSigningJobAttributes {
	return signerSigningJobAttributes{ref: terra.ReferenceResource(ssj)}
}

// ImportState imports the given attribute values into [SignerSigningJob]'s state.
func (ssj *SignerSigningJob) ImportState(av io.Reader) error {
	ssj.state = &signerSigningJobState{}
	if err := json.NewDecoder(av).Decode(ssj.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ssj.Type(), ssj.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SignerSigningJob] has state.
func (ssj *SignerSigningJob) State() (*signerSigningJobState, bool) {
	return ssj.state, ssj.state != nil
}

// StateMust returns the state for [SignerSigningJob]. Panics if the state is nil.
func (ssj *SignerSigningJob) StateMust() *signerSigningJobState {
	if ssj.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ssj.Type(), ssj.LocalName()))
	}
	return ssj.state
}

// SignerSigningJobArgs contains the configurations for aws_signer_signing_job.
type SignerSigningJobArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IgnoreSigningJobFailure: bool, optional
	IgnoreSigningJobFailure terra.BoolValue `hcl:"ignore_signing_job_failure,attr"`
	// ProfileName: string, required
	ProfileName terra.StringValue `hcl:"profile_name,attr" validate:"required"`
	// RevocationRecord: min=0
	RevocationRecord []signersigningjob.RevocationRecord `hcl:"revocation_record,block" validate:"min=0"`
	// SignedObject: min=0
	SignedObject []signersigningjob.SignedObject `hcl:"signed_object,block" validate:"min=0"`
	// Destination: required
	Destination *signersigningjob.Destination `hcl:"destination,block" validate:"required"`
	// Source: required
	Source *signersigningjob.Source `hcl:"source,block" validate:"required"`
}
type signerSigningJobAttributes struct {
	ref terra.Reference
}

// CompletedAt returns a reference to field completed_at of aws_signer_signing_job.
func (ssj signerSigningJobAttributes) CompletedAt() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("completed_at"))
}

// CreatedAt returns a reference to field created_at of aws_signer_signing_job.
func (ssj signerSigningJobAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("created_at"))
}

// Id returns a reference to field id of aws_signer_signing_job.
func (ssj signerSigningJobAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("id"))
}

// IgnoreSigningJobFailure returns a reference to field ignore_signing_job_failure of aws_signer_signing_job.
func (ssj signerSigningJobAttributes) IgnoreSigningJobFailure() terra.BoolValue {
	return terra.ReferenceAsBool(ssj.ref.Append("ignore_signing_job_failure"))
}

// JobId returns a reference to field job_id of aws_signer_signing_job.
func (ssj signerSigningJobAttributes) JobId() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("job_id"))
}

// JobInvoker returns a reference to field job_invoker of aws_signer_signing_job.
func (ssj signerSigningJobAttributes) JobInvoker() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("job_invoker"))
}

// JobOwner returns a reference to field job_owner of aws_signer_signing_job.
func (ssj signerSigningJobAttributes) JobOwner() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("job_owner"))
}

// PlatformDisplayName returns a reference to field platform_display_name of aws_signer_signing_job.
func (ssj signerSigningJobAttributes) PlatformDisplayName() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("platform_display_name"))
}

// PlatformId returns a reference to field platform_id of aws_signer_signing_job.
func (ssj signerSigningJobAttributes) PlatformId() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("platform_id"))
}

// ProfileName returns a reference to field profile_name of aws_signer_signing_job.
func (ssj signerSigningJobAttributes) ProfileName() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("profile_name"))
}

// ProfileVersion returns a reference to field profile_version of aws_signer_signing_job.
func (ssj signerSigningJobAttributes) ProfileVersion() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("profile_version"))
}

// RequestedBy returns a reference to field requested_by of aws_signer_signing_job.
func (ssj signerSigningJobAttributes) RequestedBy() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("requested_by"))
}

// SignatureExpiresAt returns a reference to field signature_expires_at of aws_signer_signing_job.
func (ssj signerSigningJobAttributes) SignatureExpiresAt() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("signature_expires_at"))
}

// Status returns a reference to field status of aws_signer_signing_job.
func (ssj signerSigningJobAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("status"))
}

// StatusReason returns a reference to field status_reason of aws_signer_signing_job.
func (ssj signerSigningJobAttributes) StatusReason() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("status_reason"))
}

func (ssj signerSigningJobAttributes) RevocationRecord() terra.ListValue[signersigningjob.RevocationRecordAttributes] {
	return terra.ReferenceAsList[signersigningjob.RevocationRecordAttributes](ssj.ref.Append("revocation_record"))
}

func (ssj signerSigningJobAttributes) SignedObject() terra.ListValue[signersigningjob.SignedObjectAttributes] {
	return terra.ReferenceAsList[signersigningjob.SignedObjectAttributes](ssj.ref.Append("signed_object"))
}

func (ssj signerSigningJobAttributes) Destination() terra.ListValue[signersigningjob.DestinationAttributes] {
	return terra.ReferenceAsList[signersigningjob.DestinationAttributes](ssj.ref.Append("destination"))
}

func (ssj signerSigningJobAttributes) Source() terra.ListValue[signersigningjob.SourceAttributes] {
	return terra.ReferenceAsList[signersigningjob.SourceAttributes](ssj.ref.Append("source"))
}

type signerSigningJobState struct {
	CompletedAt             string                                   `json:"completed_at"`
	CreatedAt               string                                   `json:"created_at"`
	Id                      string                                   `json:"id"`
	IgnoreSigningJobFailure bool                                     `json:"ignore_signing_job_failure"`
	JobId                   string                                   `json:"job_id"`
	JobInvoker              string                                   `json:"job_invoker"`
	JobOwner                string                                   `json:"job_owner"`
	PlatformDisplayName     string                                   `json:"platform_display_name"`
	PlatformId              string                                   `json:"platform_id"`
	ProfileName             string                                   `json:"profile_name"`
	ProfileVersion          string                                   `json:"profile_version"`
	RequestedBy             string                                   `json:"requested_by"`
	SignatureExpiresAt      string                                   `json:"signature_expires_at"`
	Status                  string                                   `json:"status"`
	StatusReason            string                                   `json:"status_reason"`
	RevocationRecord        []signersigningjob.RevocationRecordState `json:"revocation_record"`
	SignedObject            []signersigningjob.SignedObjectState     `json:"signed_object"`
	Destination             []signersigningjob.DestinationState      `json:"destination"`
	Source                  []signersigningjob.SourceState           `json:"source"`
}
