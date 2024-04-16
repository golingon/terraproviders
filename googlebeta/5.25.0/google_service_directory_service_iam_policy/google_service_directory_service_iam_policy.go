// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_service_directory_service_iam_policy

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

// Resource represents the Terraform resource google_service_directory_service_iam_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *googleServiceDirectoryServiceIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gsdsip *Resource) Type() string {
	return "google_service_directory_service_iam_policy"
}

// LocalName returns the local name for [Resource].
func (gsdsip *Resource) LocalName() string {
	return gsdsip.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gsdsip *Resource) Configuration() interface{} {
	return gsdsip.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gsdsip *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gsdsip)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gsdsip *Resource) Dependencies() terra.Dependencies {
	return gsdsip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gsdsip *Resource) LifecycleManagement() *terra.Lifecycle {
	return gsdsip.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gsdsip *Resource) Attributes() googleServiceDirectoryServiceIamPolicyAttributes {
	return googleServiceDirectoryServiceIamPolicyAttributes{ref: terra.ReferenceResource(gsdsip)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gsdsip *Resource) ImportState(state io.Reader) error {
	gsdsip.state = &googleServiceDirectoryServiceIamPolicyState{}
	if err := json.NewDecoder(state).Decode(gsdsip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gsdsip.Type(), gsdsip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gsdsip *Resource) State() (*googleServiceDirectoryServiceIamPolicyState, bool) {
	return gsdsip.state, gsdsip.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gsdsip *Resource) StateMust() *googleServiceDirectoryServiceIamPolicyState {
	if gsdsip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gsdsip.Type(), gsdsip.LocalName()))
	}
	return gsdsip.state
}

// Args contains the configurations for google_service_directory_service_iam_policy.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
}

type googleServiceDirectoryServiceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_service_directory_service_iam_policy.
func (gsdsip googleServiceDirectoryServiceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gsdsip.ref.Append("etag"))
}

// Id returns a reference to field id of google_service_directory_service_iam_policy.
func (gsdsip googleServiceDirectoryServiceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gsdsip.ref.Append("id"))
}

// Name returns a reference to field name of google_service_directory_service_iam_policy.
func (gsdsip googleServiceDirectoryServiceIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gsdsip.ref.Append("name"))
}

// PolicyData returns a reference to field policy_data of google_service_directory_service_iam_policy.
func (gsdsip googleServiceDirectoryServiceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(gsdsip.ref.Append("policy_data"))
}

type googleServiceDirectoryServiceIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	PolicyData string `json:"policy_data"`
}
