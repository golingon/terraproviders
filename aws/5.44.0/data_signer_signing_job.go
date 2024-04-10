// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"github.com/golingon/lingon/pkg/terra"
	datasignersigningjob "github.com/golingon/terraproviders/aws/5.44.0/datasignersigningjob"
)

// NewDataSignerSigningJob creates a new instance of [DataSignerSigningJob].
func NewDataSignerSigningJob(name string, args DataSignerSigningJobArgs) *DataSignerSigningJob {
	return &DataSignerSigningJob{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSignerSigningJob)(nil)

// DataSignerSigningJob represents the Terraform data resource aws_signer_signing_job.
type DataSignerSigningJob struct {
	Name string
	Args DataSignerSigningJobArgs
}

// DataSource returns the Terraform object type for [DataSignerSigningJob].
func (ssj *DataSignerSigningJob) DataSource() string {
	return "aws_signer_signing_job"
}

// LocalName returns the local name for [DataSignerSigningJob].
func (ssj *DataSignerSigningJob) LocalName() string {
	return ssj.Name
}

// Configuration returns the configuration (args) for [DataSignerSigningJob].
func (ssj *DataSignerSigningJob) Configuration() interface{} {
	return ssj.Args
}

// Attributes returns the attributes for [DataSignerSigningJob].
func (ssj *DataSignerSigningJob) Attributes() dataSignerSigningJobAttributes {
	return dataSignerSigningJobAttributes{ref: terra.ReferenceDataResource(ssj)}
}

// DataSignerSigningJobArgs contains the configurations for aws_signer_signing_job.
type DataSignerSigningJobArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// JobId: string, required
	JobId terra.StringValue `hcl:"job_id,attr" validate:"required"`
	// RevocationRecord: min=0
	RevocationRecord []datasignersigningjob.RevocationRecord `hcl:"revocation_record,block" validate:"min=0"`
	// SignedObject: min=0
	SignedObject []datasignersigningjob.SignedObject `hcl:"signed_object,block" validate:"min=0"`
	// Source: min=0
	Source []datasignersigningjob.Source `hcl:"source,block" validate:"min=0"`
}
type dataSignerSigningJobAttributes struct {
	ref terra.Reference
}

// CompletedAt returns a reference to field completed_at of aws_signer_signing_job.
func (ssj dataSignerSigningJobAttributes) CompletedAt() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("completed_at"))
}

// CreatedAt returns a reference to field created_at of aws_signer_signing_job.
func (ssj dataSignerSigningJobAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("created_at"))
}

// Id returns a reference to field id of aws_signer_signing_job.
func (ssj dataSignerSigningJobAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("id"))
}

// JobId returns a reference to field job_id of aws_signer_signing_job.
func (ssj dataSignerSigningJobAttributes) JobId() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("job_id"))
}

// JobInvoker returns a reference to field job_invoker of aws_signer_signing_job.
func (ssj dataSignerSigningJobAttributes) JobInvoker() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("job_invoker"))
}

// JobOwner returns a reference to field job_owner of aws_signer_signing_job.
func (ssj dataSignerSigningJobAttributes) JobOwner() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("job_owner"))
}

// PlatformDisplayName returns a reference to field platform_display_name of aws_signer_signing_job.
func (ssj dataSignerSigningJobAttributes) PlatformDisplayName() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("platform_display_name"))
}

// PlatformId returns a reference to field platform_id of aws_signer_signing_job.
func (ssj dataSignerSigningJobAttributes) PlatformId() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("platform_id"))
}

// ProfileName returns a reference to field profile_name of aws_signer_signing_job.
func (ssj dataSignerSigningJobAttributes) ProfileName() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("profile_name"))
}

// ProfileVersion returns a reference to field profile_version of aws_signer_signing_job.
func (ssj dataSignerSigningJobAttributes) ProfileVersion() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("profile_version"))
}

// RequestedBy returns a reference to field requested_by of aws_signer_signing_job.
func (ssj dataSignerSigningJobAttributes) RequestedBy() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("requested_by"))
}

// SignatureExpiresAt returns a reference to field signature_expires_at of aws_signer_signing_job.
func (ssj dataSignerSigningJobAttributes) SignatureExpiresAt() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("signature_expires_at"))
}

// Status returns a reference to field status of aws_signer_signing_job.
func (ssj dataSignerSigningJobAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("status"))
}

// StatusReason returns a reference to field status_reason of aws_signer_signing_job.
func (ssj dataSignerSigningJobAttributes) StatusReason() terra.StringValue {
	return terra.ReferenceAsString(ssj.ref.Append("status_reason"))
}

func (ssj dataSignerSigningJobAttributes) RevocationRecord() terra.ListValue[datasignersigningjob.RevocationRecordAttributes] {
	return terra.ReferenceAsList[datasignersigningjob.RevocationRecordAttributes](ssj.ref.Append("revocation_record"))
}

func (ssj dataSignerSigningJobAttributes) SignedObject() terra.ListValue[datasignersigningjob.SignedObjectAttributes] {
	return terra.ReferenceAsList[datasignersigningjob.SignedObjectAttributes](ssj.ref.Append("signed_object"))
}

func (ssj dataSignerSigningJobAttributes) Source() terra.ListValue[datasignersigningjob.SourceAttributes] {
	return terra.ReferenceAsList[datasignersigningjob.SourceAttributes](ssj.ref.Append("source"))
}
