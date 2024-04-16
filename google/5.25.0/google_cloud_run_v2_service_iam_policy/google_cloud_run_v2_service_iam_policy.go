// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_cloud_run_v2_service_iam_policy

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

// Resource represents the Terraform resource google_cloud_run_v2_service_iam_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *googleCloudRunV2ServiceIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcrvsip *Resource) Type() string {
	return "google_cloud_run_v2_service_iam_policy"
}

// LocalName returns the local name for [Resource].
func (gcrvsip *Resource) LocalName() string {
	return gcrvsip.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcrvsip *Resource) Configuration() interface{} {
	return gcrvsip.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcrvsip *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcrvsip)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcrvsip *Resource) Dependencies() terra.Dependencies {
	return gcrvsip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcrvsip *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcrvsip.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcrvsip *Resource) Attributes() googleCloudRunV2ServiceIamPolicyAttributes {
	return googleCloudRunV2ServiceIamPolicyAttributes{ref: terra.ReferenceResource(gcrvsip)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcrvsip *Resource) ImportState(state io.Reader) error {
	gcrvsip.state = &googleCloudRunV2ServiceIamPolicyState{}
	if err := json.NewDecoder(state).Decode(gcrvsip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcrvsip.Type(), gcrvsip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcrvsip *Resource) State() (*googleCloudRunV2ServiceIamPolicyState, bool) {
	return gcrvsip.state, gcrvsip.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcrvsip *Resource) StateMust() *googleCloudRunV2ServiceIamPolicyState {
	if gcrvsip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcrvsip.Type(), gcrvsip.LocalName()))
	}
	return gcrvsip.state
}

// Args contains the configurations for google_cloud_run_v2_service_iam_policy.
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

type googleCloudRunV2ServiceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_cloud_run_v2_service_iam_policy.
func (gcrvsip googleCloudRunV2ServiceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gcrvsip.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloud_run_v2_service_iam_policy.
func (gcrvsip googleCloudRunV2ServiceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcrvsip.ref.Append("id"))
}

// Location returns a reference to field location of google_cloud_run_v2_service_iam_policy.
func (gcrvsip googleCloudRunV2ServiceIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gcrvsip.ref.Append("location"))
}

// Name returns a reference to field name of google_cloud_run_v2_service_iam_policy.
func (gcrvsip googleCloudRunV2ServiceIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcrvsip.ref.Append("name"))
}

// PolicyData returns a reference to field policy_data of google_cloud_run_v2_service_iam_policy.
func (gcrvsip googleCloudRunV2ServiceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(gcrvsip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_cloud_run_v2_service_iam_policy.
func (gcrvsip googleCloudRunV2ServiceIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcrvsip.ref.Append("project"))
}

type googleCloudRunV2ServiceIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
}
