// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	kmskeyringimportjob "github.com/golingon/terraproviders/googlebeta/5.24.0/kmskeyringimportjob"
	"io"
)

// NewKmsKeyRingImportJob creates a new instance of [KmsKeyRingImportJob].
func NewKmsKeyRingImportJob(name string, args KmsKeyRingImportJobArgs) *KmsKeyRingImportJob {
	return &KmsKeyRingImportJob{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KmsKeyRingImportJob)(nil)

// KmsKeyRingImportJob represents the Terraform resource google_kms_key_ring_import_job.
type KmsKeyRingImportJob struct {
	Name      string
	Args      KmsKeyRingImportJobArgs
	state     *kmsKeyRingImportJobState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KmsKeyRingImportJob].
func (kkrij *KmsKeyRingImportJob) Type() string {
	return "google_kms_key_ring_import_job"
}

// LocalName returns the local name for [KmsKeyRingImportJob].
func (kkrij *KmsKeyRingImportJob) LocalName() string {
	return kkrij.Name
}

// Configuration returns the configuration (args) for [KmsKeyRingImportJob].
func (kkrij *KmsKeyRingImportJob) Configuration() interface{} {
	return kkrij.Args
}

// DependOn is used for other resources to depend on [KmsKeyRingImportJob].
func (kkrij *KmsKeyRingImportJob) DependOn() terra.Reference {
	return terra.ReferenceResource(kkrij)
}

// Dependencies returns the list of resources [KmsKeyRingImportJob] depends_on.
func (kkrij *KmsKeyRingImportJob) Dependencies() terra.Dependencies {
	return kkrij.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KmsKeyRingImportJob].
func (kkrij *KmsKeyRingImportJob) LifecycleManagement() *terra.Lifecycle {
	return kkrij.Lifecycle
}

// Attributes returns the attributes for [KmsKeyRingImportJob].
func (kkrij *KmsKeyRingImportJob) Attributes() kmsKeyRingImportJobAttributes {
	return kmsKeyRingImportJobAttributes{ref: terra.ReferenceResource(kkrij)}
}

// ImportState imports the given attribute values into [KmsKeyRingImportJob]'s state.
func (kkrij *KmsKeyRingImportJob) ImportState(av io.Reader) error {
	kkrij.state = &kmsKeyRingImportJobState{}
	if err := json.NewDecoder(av).Decode(kkrij.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kkrij.Type(), kkrij.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KmsKeyRingImportJob] has state.
func (kkrij *KmsKeyRingImportJob) State() (*kmsKeyRingImportJobState, bool) {
	return kkrij.state, kkrij.state != nil
}

// StateMust returns the state for [KmsKeyRingImportJob]. Panics if the state is nil.
func (kkrij *KmsKeyRingImportJob) StateMust() *kmsKeyRingImportJobState {
	if kkrij.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kkrij.Type(), kkrij.LocalName()))
	}
	return kkrij.state
}

// KmsKeyRingImportJobArgs contains the configurations for google_kms_key_ring_import_job.
type KmsKeyRingImportJobArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImportJobId: string, required
	ImportJobId terra.StringValue `hcl:"import_job_id,attr" validate:"required"`
	// ImportMethod: string, required
	ImportMethod terra.StringValue `hcl:"import_method,attr" validate:"required"`
	// KeyRing: string, required
	KeyRing terra.StringValue `hcl:"key_ring,attr" validate:"required"`
	// ProtectionLevel: string, required
	ProtectionLevel terra.StringValue `hcl:"protection_level,attr" validate:"required"`
	// Attestation: min=0
	Attestation []kmskeyringimportjob.Attestation `hcl:"attestation,block" validate:"min=0"`
	// PublicKey: min=0
	PublicKey []kmskeyringimportjob.PublicKey `hcl:"public_key,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *kmskeyringimportjob.Timeouts `hcl:"timeouts,block"`
}
type kmsKeyRingImportJobAttributes struct {
	ref terra.Reference
}

// ExpireTime returns a reference to field expire_time of google_kms_key_ring_import_job.
func (kkrij kmsKeyRingImportJobAttributes) ExpireTime() terra.StringValue {
	return terra.ReferenceAsString(kkrij.ref.Append("expire_time"))
}

// Id returns a reference to field id of google_kms_key_ring_import_job.
func (kkrij kmsKeyRingImportJobAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kkrij.ref.Append("id"))
}

// ImportJobId returns a reference to field import_job_id of google_kms_key_ring_import_job.
func (kkrij kmsKeyRingImportJobAttributes) ImportJobId() terra.StringValue {
	return terra.ReferenceAsString(kkrij.ref.Append("import_job_id"))
}

// ImportMethod returns a reference to field import_method of google_kms_key_ring_import_job.
func (kkrij kmsKeyRingImportJobAttributes) ImportMethod() terra.StringValue {
	return terra.ReferenceAsString(kkrij.ref.Append("import_method"))
}

// KeyRing returns a reference to field key_ring of google_kms_key_ring_import_job.
func (kkrij kmsKeyRingImportJobAttributes) KeyRing() terra.StringValue {
	return terra.ReferenceAsString(kkrij.ref.Append("key_ring"))
}

// Name returns a reference to field name of google_kms_key_ring_import_job.
func (kkrij kmsKeyRingImportJobAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kkrij.ref.Append("name"))
}

// ProtectionLevel returns a reference to field protection_level of google_kms_key_ring_import_job.
func (kkrij kmsKeyRingImportJobAttributes) ProtectionLevel() terra.StringValue {
	return terra.ReferenceAsString(kkrij.ref.Append("protection_level"))
}

// State returns a reference to field state of google_kms_key_ring_import_job.
func (kkrij kmsKeyRingImportJobAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(kkrij.ref.Append("state"))
}

func (kkrij kmsKeyRingImportJobAttributes) Attestation() terra.ListValue[kmskeyringimportjob.AttestationAttributes] {
	return terra.ReferenceAsList[kmskeyringimportjob.AttestationAttributes](kkrij.ref.Append("attestation"))
}

func (kkrij kmsKeyRingImportJobAttributes) PublicKey() terra.ListValue[kmskeyringimportjob.PublicKeyAttributes] {
	return terra.ReferenceAsList[kmskeyringimportjob.PublicKeyAttributes](kkrij.ref.Append("public_key"))
}

func (kkrij kmsKeyRingImportJobAttributes) Timeouts() kmskeyringimportjob.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kmskeyringimportjob.TimeoutsAttributes](kkrij.ref.Append("timeouts"))
}

type kmsKeyRingImportJobState struct {
	ExpireTime      string                                 `json:"expire_time"`
	Id              string                                 `json:"id"`
	ImportJobId     string                                 `json:"import_job_id"`
	ImportMethod    string                                 `json:"import_method"`
	KeyRing         string                                 `json:"key_ring"`
	Name            string                                 `json:"name"`
	ProtectionLevel string                                 `json:"protection_level"`
	State           string                                 `json:"state"`
	Attestation     []kmskeyringimportjob.AttestationState `json:"attestation"`
	PublicKey       []kmskeyringimportjob.PublicKeyState   `json:"public_key"`
	Timeouts        *kmskeyringimportjob.TimeoutsState     `json:"timeouts"`
}
