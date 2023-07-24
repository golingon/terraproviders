// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHealthcareConsentStoreIamPolicy creates a new instance of [HealthcareConsentStoreIamPolicy].
func NewHealthcareConsentStoreIamPolicy(name string, args HealthcareConsentStoreIamPolicyArgs) *HealthcareConsentStoreIamPolicy {
	return &HealthcareConsentStoreIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HealthcareConsentStoreIamPolicy)(nil)

// HealthcareConsentStoreIamPolicy represents the Terraform resource google_healthcare_consent_store_iam_policy.
type HealthcareConsentStoreIamPolicy struct {
	Name      string
	Args      HealthcareConsentStoreIamPolicyArgs
	state     *healthcareConsentStoreIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HealthcareConsentStoreIamPolicy].
func (hcsip *HealthcareConsentStoreIamPolicy) Type() string {
	return "google_healthcare_consent_store_iam_policy"
}

// LocalName returns the local name for [HealthcareConsentStoreIamPolicy].
func (hcsip *HealthcareConsentStoreIamPolicy) LocalName() string {
	return hcsip.Name
}

// Configuration returns the configuration (args) for [HealthcareConsentStoreIamPolicy].
func (hcsip *HealthcareConsentStoreIamPolicy) Configuration() interface{} {
	return hcsip.Args
}

// DependOn is used for other resources to depend on [HealthcareConsentStoreIamPolicy].
func (hcsip *HealthcareConsentStoreIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(hcsip)
}

// Dependencies returns the list of resources [HealthcareConsentStoreIamPolicy] depends_on.
func (hcsip *HealthcareConsentStoreIamPolicy) Dependencies() terra.Dependencies {
	return hcsip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HealthcareConsentStoreIamPolicy].
func (hcsip *HealthcareConsentStoreIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return hcsip.Lifecycle
}

// Attributes returns the attributes for [HealthcareConsentStoreIamPolicy].
func (hcsip *HealthcareConsentStoreIamPolicy) Attributes() healthcareConsentStoreIamPolicyAttributes {
	return healthcareConsentStoreIamPolicyAttributes{ref: terra.ReferenceResource(hcsip)}
}

// ImportState imports the given attribute values into [HealthcareConsentStoreIamPolicy]'s state.
func (hcsip *HealthcareConsentStoreIamPolicy) ImportState(av io.Reader) error {
	hcsip.state = &healthcareConsentStoreIamPolicyState{}
	if err := json.NewDecoder(av).Decode(hcsip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hcsip.Type(), hcsip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HealthcareConsentStoreIamPolicy] has state.
func (hcsip *HealthcareConsentStoreIamPolicy) State() (*healthcareConsentStoreIamPolicyState, bool) {
	return hcsip.state, hcsip.state != nil
}

// StateMust returns the state for [HealthcareConsentStoreIamPolicy]. Panics if the state is nil.
func (hcsip *HealthcareConsentStoreIamPolicy) StateMust() *healthcareConsentStoreIamPolicyState {
	if hcsip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hcsip.Type(), hcsip.LocalName()))
	}
	return hcsip.state
}

// HealthcareConsentStoreIamPolicyArgs contains the configurations for google_healthcare_consent_store_iam_policy.
type HealthcareConsentStoreIamPolicyArgs struct {
	// ConsentStoreId: string, required
	ConsentStoreId terra.StringValue `hcl:"consent_store_id,attr" validate:"required"`
	// Dataset: string, required
	Dataset terra.StringValue `hcl:"dataset,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
}
type healthcareConsentStoreIamPolicyAttributes struct {
	ref terra.Reference
}

// ConsentStoreId returns a reference to field consent_store_id of google_healthcare_consent_store_iam_policy.
func (hcsip healthcareConsentStoreIamPolicyAttributes) ConsentStoreId() terra.StringValue {
	return terra.ReferenceAsString(hcsip.ref.Append("consent_store_id"))
}

// Dataset returns a reference to field dataset of google_healthcare_consent_store_iam_policy.
func (hcsip healthcareConsentStoreIamPolicyAttributes) Dataset() terra.StringValue {
	return terra.ReferenceAsString(hcsip.ref.Append("dataset"))
}

// Etag returns a reference to field etag of google_healthcare_consent_store_iam_policy.
func (hcsip healthcareConsentStoreIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(hcsip.ref.Append("etag"))
}

// Id returns a reference to field id of google_healthcare_consent_store_iam_policy.
func (hcsip healthcareConsentStoreIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hcsip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_healthcare_consent_store_iam_policy.
func (hcsip healthcareConsentStoreIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(hcsip.ref.Append("policy_data"))
}

type healthcareConsentStoreIamPolicyState struct {
	ConsentStoreId string `json:"consent_store_id"`
	Dataset        string `json:"dataset"`
	Etag           string `json:"etag"`
	Id             string `json:"id"`
	PolicyData     string `json:"policy_data"`
}
