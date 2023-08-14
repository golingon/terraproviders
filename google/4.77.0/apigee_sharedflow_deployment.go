// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	apigeesharedflowdeployment "github.com/golingon/terraproviders/google/4.77.0/apigeesharedflowdeployment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApigeeSharedflowDeployment creates a new instance of [ApigeeSharedflowDeployment].
func NewApigeeSharedflowDeployment(name string, args ApigeeSharedflowDeploymentArgs) *ApigeeSharedflowDeployment {
	return &ApigeeSharedflowDeployment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApigeeSharedflowDeployment)(nil)

// ApigeeSharedflowDeployment represents the Terraform resource google_apigee_sharedflow_deployment.
type ApigeeSharedflowDeployment struct {
	Name      string
	Args      ApigeeSharedflowDeploymentArgs
	state     *apigeeSharedflowDeploymentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApigeeSharedflowDeployment].
func (asd *ApigeeSharedflowDeployment) Type() string {
	return "google_apigee_sharedflow_deployment"
}

// LocalName returns the local name for [ApigeeSharedflowDeployment].
func (asd *ApigeeSharedflowDeployment) LocalName() string {
	return asd.Name
}

// Configuration returns the configuration (args) for [ApigeeSharedflowDeployment].
func (asd *ApigeeSharedflowDeployment) Configuration() interface{} {
	return asd.Args
}

// DependOn is used for other resources to depend on [ApigeeSharedflowDeployment].
func (asd *ApigeeSharedflowDeployment) DependOn() terra.Reference {
	return terra.ReferenceResource(asd)
}

// Dependencies returns the list of resources [ApigeeSharedflowDeployment] depends_on.
func (asd *ApigeeSharedflowDeployment) Dependencies() terra.Dependencies {
	return asd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApigeeSharedflowDeployment].
func (asd *ApigeeSharedflowDeployment) LifecycleManagement() *terra.Lifecycle {
	return asd.Lifecycle
}

// Attributes returns the attributes for [ApigeeSharedflowDeployment].
func (asd *ApigeeSharedflowDeployment) Attributes() apigeeSharedflowDeploymentAttributes {
	return apigeeSharedflowDeploymentAttributes{ref: terra.ReferenceResource(asd)}
}

// ImportState imports the given attribute values into [ApigeeSharedflowDeployment]'s state.
func (asd *ApigeeSharedflowDeployment) ImportState(av io.Reader) error {
	asd.state = &apigeeSharedflowDeploymentState{}
	if err := json.NewDecoder(av).Decode(asd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asd.Type(), asd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApigeeSharedflowDeployment] has state.
func (asd *ApigeeSharedflowDeployment) State() (*apigeeSharedflowDeploymentState, bool) {
	return asd.state, asd.state != nil
}

// StateMust returns the state for [ApigeeSharedflowDeployment]. Panics if the state is nil.
func (asd *ApigeeSharedflowDeployment) StateMust() *apigeeSharedflowDeploymentState {
	if asd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asd.Type(), asd.LocalName()))
	}
	return asd.state
}

// ApigeeSharedflowDeploymentArgs contains the configurations for google_apigee_sharedflow_deployment.
type ApigeeSharedflowDeploymentArgs struct {
	// Environment: string, required
	Environment terra.StringValue `hcl:"environment,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OrgId: string, required
	OrgId terra.StringValue `hcl:"org_id,attr" validate:"required"`
	// Revision: string, required
	Revision terra.StringValue `hcl:"revision,attr" validate:"required"`
	// ServiceAccount: string, optional
	ServiceAccount terra.StringValue `hcl:"service_account,attr"`
	// SharedflowId: string, required
	SharedflowId terra.StringValue `hcl:"sharedflow_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *apigeesharedflowdeployment.Timeouts `hcl:"timeouts,block"`
}
type apigeeSharedflowDeploymentAttributes struct {
	ref terra.Reference
}

// Environment returns a reference to field environment of google_apigee_sharedflow_deployment.
func (asd apigeeSharedflowDeploymentAttributes) Environment() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("environment"))
}

// Id returns a reference to field id of google_apigee_sharedflow_deployment.
func (asd apigeeSharedflowDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("id"))
}

// OrgId returns a reference to field org_id of google_apigee_sharedflow_deployment.
func (asd apigeeSharedflowDeploymentAttributes) OrgId() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("org_id"))
}

// Revision returns a reference to field revision of google_apigee_sharedflow_deployment.
func (asd apigeeSharedflowDeploymentAttributes) Revision() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("revision"))
}

// ServiceAccount returns a reference to field service_account of google_apigee_sharedflow_deployment.
func (asd apigeeSharedflowDeploymentAttributes) ServiceAccount() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("service_account"))
}

// SharedflowId returns a reference to field sharedflow_id of google_apigee_sharedflow_deployment.
func (asd apigeeSharedflowDeploymentAttributes) SharedflowId() terra.StringValue {
	return terra.ReferenceAsString(asd.ref.Append("sharedflow_id"))
}

func (asd apigeeSharedflowDeploymentAttributes) Timeouts() apigeesharedflowdeployment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apigeesharedflowdeployment.TimeoutsAttributes](asd.ref.Append("timeouts"))
}

type apigeeSharedflowDeploymentState struct {
	Environment    string                                    `json:"environment"`
	Id             string                                    `json:"id"`
	OrgId          string                                    `json:"org_id"`
	Revision       string                                    `json:"revision"`
	ServiceAccount string                                    `json:"service_account"`
	SharedflowId   string                                    `json:"sharedflow_id"`
	Timeouts       *apigeesharedflowdeployment.TimeoutsState `json:"timeouts"`
}