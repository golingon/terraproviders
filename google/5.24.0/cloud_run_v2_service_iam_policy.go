// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewCloudRunV2ServiceIamPolicy creates a new instance of [CloudRunV2ServiceIamPolicy].
func NewCloudRunV2ServiceIamPolicy(name string, args CloudRunV2ServiceIamPolicyArgs) *CloudRunV2ServiceIamPolicy {
	return &CloudRunV2ServiceIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudRunV2ServiceIamPolicy)(nil)

// CloudRunV2ServiceIamPolicy represents the Terraform resource google_cloud_run_v2_service_iam_policy.
type CloudRunV2ServiceIamPolicy struct {
	Name      string
	Args      CloudRunV2ServiceIamPolicyArgs
	state     *cloudRunV2ServiceIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudRunV2ServiceIamPolicy].
func (crvsip *CloudRunV2ServiceIamPolicy) Type() string {
	return "google_cloud_run_v2_service_iam_policy"
}

// LocalName returns the local name for [CloudRunV2ServiceIamPolicy].
func (crvsip *CloudRunV2ServiceIamPolicy) LocalName() string {
	return crvsip.Name
}

// Configuration returns the configuration (args) for [CloudRunV2ServiceIamPolicy].
func (crvsip *CloudRunV2ServiceIamPolicy) Configuration() interface{} {
	return crvsip.Args
}

// DependOn is used for other resources to depend on [CloudRunV2ServiceIamPolicy].
func (crvsip *CloudRunV2ServiceIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(crvsip)
}

// Dependencies returns the list of resources [CloudRunV2ServiceIamPolicy] depends_on.
func (crvsip *CloudRunV2ServiceIamPolicy) Dependencies() terra.Dependencies {
	return crvsip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudRunV2ServiceIamPolicy].
func (crvsip *CloudRunV2ServiceIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return crvsip.Lifecycle
}

// Attributes returns the attributes for [CloudRunV2ServiceIamPolicy].
func (crvsip *CloudRunV2ServiceIamPolicy) Attributes() cloudRunV2ServiceIamPolicyAttributes {
	return cloudRunV2ServiceIamPolicyAttributes{ref: terra.ReferenceResource(crvsip)}
}

// ImportState imports the given attribute values into [CloudRunV2ServiceIamPolicy]'s state.
func (crvsip *CloudRunV2ServiceIamPolicy) ImportState(av io.Reader) error {
	crvsip.state = &cloudRunV2ServiceIamPolicyState{}
	if err := json.NewDecoder(av).Decode(crvsip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crvsip.Type(), crvsip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudRunV2ServiceIamPolicy] has state.
func (crvsip *CloudRunV2ServiceIamPolicy) State() (*cloudRunV2ServiceIamPolicyState, bool) {
	return crvsip.state, crvsip.state != nil
}

// StateMust returns the state for [CloudRunV2ServiceIamPolicy]. Panics if the state is nil.
func (crvsip *CloudRunV2ServiceIamPolicy) StateMust() *cloudRunV2ServiceIamPolicyState {
	if crvsip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crvsip.Type(), crvsip.LocalName()))
	}
	return crvsip.state
}

// CloudRunV2ServiceIamPolicyArgs contains the configurations for google_cloud_run_v2_service_iam_policy.
type CloudRunV2ServiceIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type cloudRunV2ServiceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_cloud_run_v2_service_iam_policy.
func (crvsip cloudRunV2ServiceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(crvsip.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloud_run_v2_service_iam_policy.
func (crvsip cloudRunV2ServiceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crvsip.ref.Append("id"))
}

// Location returns a reference to field location of google_cloud_run_v2_service_iam_policy.
func (crvsip cloudRunV2ServiceIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(crvsip.ref.Append("location"))
}

// Name returns a reference to field name of google_cloud_run_v2_service_iam_policy.
func (crvsip cloudRunV2ServiceIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crvsip.ref.Append("name"))
}

// PolicyData returns a reference to field policy_data of google_cloud_run_v2_service_iam_policy.
func (crvsip cloudRunV2ServiceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(crvsip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_cloud_run_v2_service_iam_policy.
func (crvsip cloudRunV2ServiceIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crvsip.ref.Append("project"))
}

type cloudRunV2ServiceIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
}
