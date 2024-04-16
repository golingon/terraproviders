// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_healthcare_dicom_store_iam_policy

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_healthcare_dicom_store_iam_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *googleHealthcareDicomStoreIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ghdsip *Resource) Type() string {
	return "google_healthcare_dicom_store_iam_policy"
}

// LocalName returns the local name for [Resource].
func (ghdsip *Resource) LocalName() string {
	return ghdsip.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ghdsip *Resource) Configuration() interface{} {
	return ghdsip.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ghdsip *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ghdsip)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ghdsip *Resource) Dependencies() terra.Dependencies {
	return ghdsip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ghdsip *Resource) LifecycleManagement() *terra.Lifecycle {
	return ghdsip.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ghdsip *Resource) Attributes() googleHealthcareDicomStoreIamPolicyAttributes {
	return googleHealthcareDicomStoreIamPolicyAttributes{ref: terra.ReferenceResource(ghdsip)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ghdsip *Resource) ImportState(state io.Reader) error {
	ghdsip.state = &googleHealthcareDicomStoreIamPolicyState{}
	if err := json.NewDecoder(state).Decode(ghdsip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ghdsip.Type(), ghdsip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ghdsip *Resource) State() (*googleHealthcareDicomStoreIamPolicyState, bool) {
	return ghdsip.state, ghdsip.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ghdsip *Resource) StateMust() *googleHealthcareDicomStoreIamPolicyState {
	if ghdsip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ghdsip.Type(), ghdsip.LocalName()))
	}
	return ghdsip.state
}

// Args contains the configurations for google_healthcare_dicom_store_iam_policy.
type Args struct {
	// DicomStoreId: string, required
	DicomStoreId terra.StringValue `hcl:"dicom_store_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
}

type googleHealthcareDicomStoreIamPolicyAttributes struct {
	ref terra.Reference
}

// DicomStoreId returns a reference to field dicom_store_id of google_healthcare_dicom_store_iam_policy.
func (ghdsip googleHealthcareDicomStoreIamPolicyAttributes) DicomStoreId() terra.StringValue {
	return terra.ReferenceAsString(ghdsip.ref.Append("dicom_store_id"))
}

// Etag returns a reference to field etag of google_healthcare_dicom_store_iam_policy.
func (ghdsip googleHealthcareDicomStoreIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ghdsip.ref.Append("etag"))
}

// Id returns a reference to field id of google_healthcare_dicom_store_iam_policy.
func (ghdsip googleHealthcareDicomStoreIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ghdsip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_healthcare_dicom_store_iam_policy.
func (ghdsip googleHealthcareDicomStoreIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(ghdsip.ref.Append("policy_data"))
}

type googleHealthcareDicomStoreIamPolicyState struct {
	DicomStoreId string `json:"dicom_store_id"`
	Etag         string `json:"etag"`
	Id           string `json:"id"`
	PolicyData   string `json:"policy_data"`
}
