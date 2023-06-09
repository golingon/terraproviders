// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServiceDirectoryServiceIamPolicy creates a new instance of [ServiceDirectoryServiceIamPolicy].
func NewServiceDirectoryServiceIamPolicy(name string, args ServiceDirectoryServiceIamPolicyArgs) *ServiceDirectoryServiceIamPolicy {
	return &ServiceDirectoryServiceIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServiceDirectoryServiceIamPolicy)(nil)

// ServiceDirectoryServiceIamPolicy represents the Terraform resource google_service_directory_service_iam_policy.
type ServiceDirectoryServiceIamPolicy struct {
	Name      string
	Args      ServiceDirectoryServiceIamPolicyArgs
	state     *serviceDirectoryServiceIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServiceDirectoryServiceIamPolicy].
func (sdsip *ServiceDirectoryServiceIamPolicy) Type() string {
	return "google_service_directory_service_iam_policy"
}

// LocalName returns the local name for [ServiceDirectoryServiceIamPolicy].
func (sdsip *ServiceDirectoryServiceIamPolicy) LocalName() string {
	return sdsip.Name
}

// Configuration returns the configuration (args) for [ServiceDirectoryServiceIamPolicy].
func (sdsip *ServiceDirectoryServiceIamPolicy) Configuration() interface{} {
	return sdsip.Args
}

// DependOn is used for other resources to depend on [ServiceDirectoryServiceIamPolicy].
func (sdsip *ServiceDirectoryServiceIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(sdsip)
}

// Dependencies returns the list of resources [ServiceDirectoryServiceIamPolicy] depends_on.
func (sdsip *ServiceDirectoryServiceIamPolicy) Dependencies() terra.Dependencies {
	return sdsip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServiceDirectoryServiceIamPolicy].
func (sdsip *ServiceDirectoryServiceIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return sdsip.Lifecycle
}

// Attributes returns the attributes for [ServiceDirectoryServiceIamPolicy].
func (sdsip *ServiceDirectoryServiceIamPolicy) Attributes() serviceDirectoryServiceIamPolicyAttributes {
	return serviceDirectoryServiceIamPolicyAttributes{ref: terra.ReferenceResource(sdsip)}
}

// ImportState imports the given attribute values into [ServiceDirectoryServiceIamPolicy]'s state.
func (sdsip *ServiceDirectoryServiceIamPolicy) ImportState(av io.Reader) error {
	sdsip.state = &serviceDirectoryServiceIamPolicyState{}
	if err := json.NewDecoder(av).Decode(sdsip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdsip.Type(), sdsip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServiceDirectoryServiceIamPolicy] has state.
func (sdsip *ServiceDirectoryServiceIamPolicy) State() (*serviceDirectoryServiceIamPolicyState, bool) {
	return sdsip.state, sdsip.state != nil
}

// StateMust returns the state for [ServiceDirectoryServiceIamPolicy]. Panics if the state is nil.
func (sdsip *ServiceDirectoryServiceIamPolicy) StateMust() *serviceDirectoryServiceIamPolicyState {
	if sdsip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdsip.Type(), sdsip.LocalName()))
	}
	return sdsip.state
}

// ServiceDirectoryServiceIamPolicyArgs contains the configurations for google_service_directory_service_iam_policy.
type ServiceDirectoryServiceIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
}
type serviceDirectoryServiceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_service_directory_service_iam_policy.
func (sdsip serviceDirectoryServiceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(sdsip.ref.Append("etag"))
}

// Id returns a reference to field id of google_service_directory_service_iam_policy.
func (sdsip serviceDirectoryServiceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdsip.ref.Append("id"))
}

// Name returns a reference to field name of google_service_directory_service_iam_policy.
func (sdsip serviceDirectoryServiceIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdsip.ref.Append("name"))
}

// PolicyData returns a reference to field policy_data of google_service_directory_service_iam_policy.
func (sdsip serviceDirectoryServiceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(sdsip.ref.Append("policy_data"))
}

type serviceDirectoryServiceIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	PolicyData string `json:"policy_data"`
}
