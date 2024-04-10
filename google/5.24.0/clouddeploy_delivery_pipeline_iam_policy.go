// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewClouddeployDeliveryPipelineIamPolicy creates a new instance of [ClouddeployDeliveryPipelineIamPolicy].
func NewClouddeployDeliveryPipelineIamPolicy(name string, args ClouddeployDeliveryPipelineIamPolicyArgs) *ClouddeployDeliveryPipelineIamPolicy {
	return &ClouddeployDeliveryPipelineIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ClouddeployDeliveryPipelineIamPolicy)(nil)

// ClouddeployDeliveryPipelineIamPolicy represents the Terraform resource google_clouddeploy_delivery_pipeline_iam_policy.
type ClouddeployDeliveryPipelineIamPolicy struct {
	Name      string
	Args      ClouddeployDeliveryPipelineIamPolicyArgs
	state     *clouddeployDeliveryPipelineIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ClouddeployDeliveryPipelineIamPolicy].
func (cdpip *ClouddeployDeliveryPipelineIamPolicy) Type() string {
	return "google_clouddeploy_delivery_pipeline_iam_policy"
}

// LocalName returns the local name for [ClouddeployDeliveryPipelineIamPolicy].
func (cdpip *ClouddeployDeliveryPipelineIamPolicy) LocalName() string {
	return cdpip.Name
}

// Configuration returns the configuration (args) for [ClouddeployDeliveryPipelineIamPolicy].
func (cdpip *ClouddeployDeliveryPipelineIamPolicy) Configuration() interface{} {
	return cdpip.Args
}

// DependOn is used for other resources to depend on [ClouddeployDeliveryPipelineIamPolicy].
func (cdpip *ClouddeployDeliveryPipelineIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(cdpip)
}

// Dependencies returns the list of resources [ClouddeployDeliveryPipelineIamPolicy] depends_on.
func (cdpip *ClouddeployDeliveryPipelineIamPolicy) Dependencies() terra.Dependencies {
	return cdpip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ClouddeployDeliveryPipelineIamPolicy].
func (cdpip *ClouddeployDeliveryPipelineIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return cdpip.Lifecycle
}

// Attributes returns the attributes for [ClouddeployDeliveryPipelineIamPolicy].
func (cdpip *ClouddeployDeliveryPipelineIamPolicy) Attributes() clouddeployDeliveryPipelineIamPolicyAttributes {
	return clouddeployDeliveryPipelineIamPolicyAttributes{ref: terra.ReferenceResource(cdpip)}
}

// ImportState imports the given attribute values into [ClouddeployDeliveryPipelineIamPolicy]'s state.
func (cdpip *ClouddeployDeliveryPipelineIamPolicy) ImportState(av io.Reader) error {
	cdpip.state = &clouddeployDeliveryPipelineIamPolicyState{}
	if err := json.NewDecoder(av).Decode(cdpip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cdpip.Type(), cdpip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ClouddeployDeliveryPipelineIamPolicy] has state.
func (cdpip *ClouddeployDeliveryPipelineIamPolicy) State() (*clouddeployDeliveryPipelineIamPolicyState, bool) {
	return cdpip.state, cdpip.state != nil
}

// StateMust returns the state for [ClouddeployDeliveryPipelineIamPolicy]. Panics if the state is nil.
func (cdpip *ClouddeployDeliveryPipelineIamPolicy) StateMust() *clouddeployDeliveryPipelineIamPolicyState {
	if cdpip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cdpip.Type(), cdpip.LocalName()))
	}
	return cdpip.state
}

// ClouddeployDeliveryPipelineIamPolicyArgs contains the configurations for google_clouddeploy_delivery_pipeline_iam_policy.
type ClouddeployDeliveryPipelineIamPolicyArgs struct {
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
type clouddeployDeliveryPipelineIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_clouddeploy_delivery_pipeline_iam_policy.
func (cdpip clouddeployDeliveryPipelineIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cdpip.ref.Append("etag"))
}

// Id returns a reference to field id of google_clouddeploy_delivery_pipeline_iam_policy.
func (cdpip clouddeployDeliveryPipelineIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cdpip.ref.Append("id"))
}

// Location returns a reference to field location of google_clouddeploy_delivery_pipeline_iam_policy.
func (cdpip clouddeployDeliveryPipelineIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cdpip.ref.Append("location"))
}

// Name returns a reference to field name of google_clouddeploy_delivery_pipeline_iam_policy.
func (cdpip clouddeployDeliveryPipelineIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cdpip.ref.Append("name"))
}

// PolicyData returns a reference to field policy_data of google_clouddeploy_delivery_pipeline_iam_policy.
func (cdpip clouddeployDeliveryPipelineIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(cdpip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_clouddeploy_delivery_pipeline_iam_policy.
func (cdpip clouddeployDeliveryPipelineIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cdpip.ref.Append("project"))
}

type clouddeployDeliveryPipelineIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
}
