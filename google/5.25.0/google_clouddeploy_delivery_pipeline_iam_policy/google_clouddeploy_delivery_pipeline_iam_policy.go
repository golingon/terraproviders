// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_clouddeploy_delivery_pipeline_iam_policy

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

// Resource represents the Terraform resource google_clouddeploy_delivery_pipeline_iam_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *googleClouddeployDeliveryPipelineIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcdpip *Resource) Type() string {
	return "google_clouddeploy_delivery_pipeline_iam_policy"
}

// LocalName returns the local name for [Resource].
func (gcdpip *Resource) LocalName() string {
	return gcdpip.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcdpip *Resource) Configuration() interface{} {
	return gcdpip.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcdpip *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcdpip)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcdpip *Resource) Dependencies() terra.Dependencies {
	return gcdpip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcdpip *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcdpip.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcdpip *Resource) Attributes() googleClouddeployDeliveryPipelineIamPolicyAttributes {
	return googleClouddeployDeliveryPipelineIamPolicyAttributes{ref: terra.ReferenceResource(gcdpip)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcdpip *Resource) ImportState(state io.Reader) error {
	gcdpip.state = &googleClouddeployDeliveryPipelineIamPolicyState{}
	if err := json.NewDecoder(state).Decode(gcdpip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcdpip.Type(), gcdpip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcdpip *Resource) State() (*googleClouddeployDeliveryPipelineIamPolicyState, bool) {
	return gcdpip.state, gcdpip.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcdpip *Resource) StateMust() *googleClouddeployDeliveryPipelineIamPolicyState {
	if gcdpip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcdpip.Type(), gcdpip.LocalName()))
	}
	return gcdpip.state
}

// Args contains the configurations for google_clouddeploy_delivery_pipeline_iam_policy.
type Args struct {
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

type googleClouddeployDeliveryPipelineIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_clouddeploy_delivery_pipeline_iam_policy.
func (gcdpip googleClouddeployDeliveryPipelineIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gcdpip.ref.Append("etag"))
}

// Id returns a reference to field id of google_clouddeploy_delivery_pipeline_iam_policy.
func (gcdpip googleClouddeployDeliveryPipelineIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcdpip.ref.Append("id"))
}

// Location returns a reference to field location of google_clouddeploy_delivery_pipeline_iam_policy.
func (gcdpip googleClouddeployDeliveryPipelineIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gcdpip.ref.Append("location"))
}

// Name returns a reference to field name of google_clouddeploy_delivery_pipeline_iam_policy.
func (gcdpip googleClouddeployDeliveryPipelineIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcdpip.ref.Append("name"))
}

// PolicyData returns a reference to field policy_data of google_clouddeploy_delivery_pipeline_iam_policy.
func (gcdpip googleClouddeployDeliveryPipelineIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(gcdpip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_clouddeploy_delivery_pipeline_iam_policy.
func (gcdpip googleClouddeployDeliveryPipelineIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcdpip.ref.Append("project"))
}

type googleClouddeployDeliveryPipelineIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
}
