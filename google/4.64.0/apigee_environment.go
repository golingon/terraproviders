// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	apigeeenvironment "github.com/golingon/terraproviders/google/4.64.0/apigeeenvironment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApigeeEnvironment creates a new instance of [ApigeeEnvironment].
func NewApigeeEnvironment(name string, args ApigeeEnvironmentArgs) *ApigeeEnvironment {
	return &ApigeeEnvironment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApigeeEnvironment)(nil)

// ApigeeEnvironment represents the Terraform resource google_apigee_environment.
type ApigeeEnvironment struct {
	Name      string
	Args      ApigeeEnvironmentArgs
	state     *apigeeEnvironmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApigeeEnvironment].
func (ae *ApigeeEnvironment) Type() string {
	return "google_apigee_environment"
}

// LocalName returns the local name for [ApigeeEnvironment].
func (ae *ApigeeEnvironment) LocalName() string {
	return ae.Name
}

// Configuration returns the configuration (args) for [ApigeeEnvironment].
func (ae *ApigeeEnvironment) Configuration() interface{} {
	return ae.Args
}

// DependOn is used for other resources to depend on [ApigeeEnvironment].
func (ae *ApigeeEnvironment) DependOn() terra.Reference {
	return terra.ReferenceResource(ae)
}

// Dependencies returns the list of resources [ApigeeEnvironment] depends_on.
func (ae *ApigeeEnvironment) Dependencies() terra.Dependencies {
	return ae.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApigeeEnvironment].
func (ae *ApigeeEnvironment) LifecycleManagement() *terra.Lifecycle {
	return ae.Lifecycle
}

// Attributes returns the attributes for [ApigeeEnvironment].
func (ae *ApigeeEnvironment) Attributes() apigeeEnvironmentAttributes {
	return apigeeEnvironmentAttributes{ref: terra.ReferenceResource(ae)}
}

// ImportState imports the given attribute values into [ApigeeEnvironment]'s state.
func (ae *ApigeeEnvironment) ImportState(av io.Reader) error {
	ae.state = &apigeeEnvironmentState{}
	if err := json.NewDecoder(av).Decode(ae.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ae.Type(), ae.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApigeeEnvironment] has state.
func (ae *ApigeeEnvironment) State() (*apigeeEnvironmentState, bool) {
	return ae.state, ae.state != nil
}

// StateMust returns the state for [ApigeeEnvironment]. Panics if the state is nil.
func (ae *ApigeeEnvironment) StateMust() *apigeeEnvironmentState {
	if ae.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ae.Type(), ae.LocalName()))
	}
	return ae.state
}

// ApigeeEnvironmentArgs contains the configurations for google_apigee_environment.
type ApigeeEnvironmentArgs struct {
	// ApiProxyType: string, optional
	ApiProxyType terra.StringValue `hcl:"api_proxy_type,attr"`
	// DeploymentType: string, optional
	DeploymentType terra.StringValue `hcl:"deployment_type,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OrgId: string, required
	OrgId terra.StringValue `hcl:"org_id,attr" validate:"required"`
	// NodeConfig: optional
	NodeConfig *apigeeenvironment.NodeConfig `hcl:"node_config,block"`
	// Timeouts: optional
	Timeouts *apigeeenvironment.Timeouts `hcl:"timeouts,block"`
}
type apigeeEnvironmentAttributes struct {
	ref terra.Reference
}

// ApiProxyType returns a reference to field api_proxy_type of google_apigee_environment.
func (ae apigeeEnvironmentAttributes) ApiProxyType() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("api_proxy_type"))
}

// DeploymentType returns a reference to field deployment_type of google_apigee_environment.
func (ae apigeeEnvironmentAttributes) DeploymentType() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("deployment_type"))
}

// Description returns a reference to field description of google_apigee_environment.
func (ae apigeeEnvironmentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_apigee_environment.
func (ae apigeeEnvironmentAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("display_name"))
}

// Id returns a reference to field id of google_apigee_environment.
func (ae apigeeEnvironmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("id"))
}

// Name returns a reference to field name of google_apigee_environment.
func (ae apigeeEnvironmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("name"))
}

// OrgId returns a reference to field org_id of google_apigee_environment.
func (ae apigeeEnvironmentAttributes) OrgId() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("org_id"))
}

func (ae apigeeEnvironmentAttributes) NodeConfig() terra.ListValue[apigeeenvironment.NodeConfigAttributes] {
	return terra.ReferenceAsList[apigeeenvironment.NodeConfigAttributes](ae.ref.Append("node_config"))
}

func (ae apigeeEnvironmentAttributes) Timeouts() apigeeenvironment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apigeeenvironment.TimeoutsAttributes](ae.ref.Append("timeouts"))
}

type apigeeEnvironmentState struct {
	ApiProxyType   string                              `json:"api_proxy_type"`
	DeploymentType string                              `json:"deployment_type"`
	Description    string                              `json:"description"`
	DisplayName    string                              `json:"display_name"`
	Id             string                              `json:"id"`
	Name           string                              `json:"name"`
	OrgId          string                              `json:"org_id"`
	NodeConfig     []apigeeenvironment.NodeConfigState `json:"node_config"`
	Timeouts       *apigeeenvironment.TimeoutsState    `json:"timeouts"`
}
