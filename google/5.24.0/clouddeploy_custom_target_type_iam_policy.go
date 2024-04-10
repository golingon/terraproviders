// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewClouddeployCustomTargetTypeIamPolicy creates a new instance of [ClouddeployCustomTargetTypeIamPolicy].
func NewClouddeployCustomTargetTypeIamPolicy(name string, args ClouddeployCustomTargetTypeIamPolicyArgs) *ClouddeployCustomTargetTypeIamPolicy {
	return &ClouddeployCustomTargetTypeIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ClouddeployCustomTargetTypeIamPolicy)(nil)

// ClouddeployCustomTargetTypeIamPolicy represents the Terraform resource google_clouddeploy_custom_target_type_iam_policy.
type ClouddeployCustomTargetTypeIamPolicy struct {
	Name      string
	Args      ClouddeployCustomTargetTypeIamPolicyArgs
	state     *clouddeployCustomTargetTypeIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ClouddeployCustomTargetTypeIamPolicy].
func (ccttip *ClouddeployCustomTargetTypeIamPolicy) Type() string {
	return "google_clouddeploy_custom_target_type_iam_policy"
}

// LocalName returns the local name for [ClouddeployCustomTargetTypeIamPolicy].
func (ccttip *ClouddeployCustomTargetTypeIamPolicy) LocalName() string {
	return ccttip.Name
}

// Configuration returns the configuration (args) for [ClouddeployCustomTargetTypeIamPolicy].
func (ccttip *ClouddeployCustomTargetTypeIamPolicy) Configuration() interface{} {
	return ccttip.Args
}

// DependOn is used for other resources to depend on [ClouddeployCustomTargetTypeIamPolicy].
func (ccttip *ClouddeployCustomTargetTypeIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(ccttip)
}

// Dependencies returns the list of resources [ClouddeployCustomTargetTypeIamPolicy] depends_on.
func (ccttip *ClouddeployCustomTargetTypeIamPolicy) Dependencies() terra.Dependencies {
	return ccttip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ClouddeployCustomTargetTypeIamPolicy].
func (ccttip *ClouddeployCustomTargetTypeIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return ccttip.Lifecycle
}

// Attributes returns the attributes for [ClouddeployCustomTargetTypeIamPolicy].
func (ccttip *ClouddeployCustomTargetTypeIamPolicy) Attributes() clouddeployCustomTargetTypeIamPolicyAttributes {
	return clouddeployCustomTargetTypeIamPolicyAttributes{ref: terra.ReferenceResource(ccttip)}
}

// ImportState imports the given attribute values into [ClouddeployCustomTargetTypeIamPolicy]'s state.
func (ccttip *ClouddeployCustomTargetTypeIamPolicy) ImportState(av io.Reader) error {
	ccttip.state = &clouddeployCustomTargetTypeIamPolicyState{}
	if err := json.NewDecoder(av).Decode(ccttip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ccttip.Type(), ccttip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ClouddeployCustomTargetTypeIamPolicy] has state.
func (ccttip *ClouddeployCustomTargetTypeIamPolicy) State() (*clouddeployCustomTargetTypeIamPolicyState, bool) {
	return ccttip.state, ccttip.state != nil
}

// StateMust returns the state for [ClouddeployCustomTargetTypeIamPolicy]. Panics if the state is nil.
func (ccttip *ClouddeployCustomTargetTypeIamPolicy) StateMust() *clouddeployCustomTargetTypeIamPolicyState {
	if ccttip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ccttip.Type(), ccttip.LocalName()))
	}
	return ccttip.state
}

// ClouddeployCustomTargetTypeIamPolicyArgs contains the configurations for google_clouddeploy_custom_target_type_iam_policy.
type ClouddeployCustomTargetTypeIamPolicyArgs struct {
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
type clouddeployCustomTargetTypeIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_clouddeploy_custom_target_type_iam_policy.
func (ccttip clouddeployCustomTargetTypeIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ccttip.ref.Append("etag"))
}

// Id returns a reference to field id of google_clouddeploy_custom_target_type_iam_policy.
func (ccttip clouddeployCustomTargetTypeIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ccttip.ref.Append("id"))
}

// Location returns a reference to field location of google_clouddeploy_custom_target_type_iam_policy.
func (ccttip clouddeployCustomTargetTypeIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ccttip.ref.Append("location"))
}

// Name returns a reference to field name of google_clouddeploy_custom_target_type_iam_policy.
func (ccttip clouddeployCustomTargetTypeIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ccttip.ref.Append("name"))
}

// PolicyData returns a reference to field policy_data of google_clouddeploy_custom_target_type_iam_policy.
func (ccttip clouddeployCustomTargetTypeIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(ccttip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_clouddeploy_custom_target_type_iam_policy.
func (ccttip clouddeployCustomTargetTypeIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ccttip.ref.Append("project"))
}

type clouddeployCustomTargetTypeIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
}
