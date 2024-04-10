// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewHealthcareDicomStoreIamPolicy creates a new instance of [HealthcareDicomStoreIamPolicy].
func NewHealthcareDicomStoreIamPolicy(name string, args HealthcareDicomStoreIamPolicyArgs) *HealthcareDicomStoreIamPolicy {
	return &HealthcareDicomStoreIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareDicomStoreIamPolicy)(nil)

// HealthcareDicomStoreIamPolicy represents the Terraform resource google_healthcare_dicom_store_iam_policy.
type HealthcareDicomStoreIamPolicy struct {
	Name      string
	Args      HealthcareDicomStoreIamPolicyArgs
	state     *healthcareDicomStoreIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HealthcareDicomStoreIamPolicy].
func (hdsip *HealthcareDicomStoreIamPolicy) Type() string {
	return "google_healthcare_dicom_store_iam_policy"
}

// LocalName returns the local name for [HealthcareDicomStoreIamPolicy].
func (hdsip *HealthcareDicomStoreIamPolicy) LocalName() string {
	return hdsip.Name
}

// Configuration returns the configuration (args) for [HealthcareDicomStoreIamPolicy].
func (hdsip *HealthcareDicomStoreIamPolicy) Configuration() interface{} {
	return hdsip.Args
}

// DependOn is used for other resources to depend on [HealthcareDicomStoreIamPolicy].
func (hdsip *HealthcareDicomStoreIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(hdsip)
}

// Dependencies returns the list of resources [HealthcareDicomStoreIamPolicy] depends_on.
func (hdsip *HealthcareDicomStoreIamPolicy) Dependencies() terra.Dependencies {
	return hdsip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HealthcareDicomStoreIamPolicy].
func (hdsip *HealthcareDicomStoreIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return hdsip.Lifecycle
}

// Attributes returns the attributes for [HealthcareDicomStoreIamPolicy].
func (hdsip *HealthcareDicomStoreIamPolicy) Attributes() healthcareDicomStoreIamPolicyAttributes {
	return healthcareDicomStoreIamPolicyAttributes{ref: terra.ReferenceResource(hdsip)}
}

// ImportState imports the given attribute values into [HealthcareDicomStoreIamPolicy]'s state.
func (hdsip *HealthcareDicomStoreIamPolicy) ImportState(av io.Reader) error {
	hdsip.state = &healthcareDicomStoreIamPolicyState{}
	if err := json.NewDecoder(av).Decode(hdsip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hdsip.Type(), hdsip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HealthcareDicomStoreIamPolicy] has state.
func (hdsip *HealthcareDicomStoreIamPolicy) State() (*healthcareDicomStoreIamPolicyState, bool) {
	return hdsip.state, hdsip.state != nil
}

// StateMust returns the state for [HealthcareDicomStoreIamPolicy]. Panics if the state is nil.
func (hdsip *HealthcareDicomStoreIamPolicy) StateMust() *healthcareDicomStoreIamPolicyState {
	if hdsip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hdsip.Type(), hdsip.LocalName()))
	}
	return hdsip.state
}

// HealthcareDicomStoreIamPolicyArgs contains the configurations for google_healthcare_dicom_store_iam_policy.
type HealthcareDicomStoreIamPolicyArgs struct {
	// DicomStoreId: string, required
	DicomStoreId terra.StringValue `hcl:"dicom_store_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
}
type healthcareDicomStoreIamPolicyAttributes struct {
	ref terra.Reference
}

// DicomStoreId returns a reference to field dicom_store_id of google_healthcare_dicom_store_iam_policy.
func (hdsip healthcareDicomStoreIamPolicyAttributes) DicomStoreId() terra.StringValue {
	return terra.ReferenceAsString(hdsip.ref.Append("dicom_store_id"))
}

// Etag returns a reference to field etag of google_healthcare_dicom_store_iam_policy.
func (hdsip healthcareDicomStoreIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(hdsip.ref.Append("etag"))
}

// Id returns a reference to field id of google_healthcare_dicom_store_iam_policy.
func (hdsip healthcareDicomStoreIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hdsip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_healthcare_dicom_store_iam_policy.
func (hdsip healthcareDicomStoreIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(hdsip.ref.Append("policy_data"))
}

type healthcareDicomStoreIamPolicyState struct {
	DicomStoreId string `json:"dicom_store_id"`
	Etag         string `json:"etag"`
	Id           string `json:"id"`
	PolicyData   string `json:"policy_data"`
}
