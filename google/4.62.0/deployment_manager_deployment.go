// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	deploymentmanagerdeployment "github.com/golingon/terraproviders/google/4.62.0/deploymentmanagerdeployment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDeploymentManagerDeployment creates a new instance of [DeploymentManagerDeployment].
func NewDeploymentManagerDeployment(name string, args DeploymentManagerDeploymentArgs) *DeploymentManagerDeployment {
	return &DeploymentManagerDeployment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DeploymentManagerDeployment)(nil)

// DeploymentManagerDeployment represents the Terraform resource google_deployment_manager_deployment.
type DeploymentManagerDeployment struct {
	Name      string
	Args      DeploymentManagerDeploymentArgs
	state     *deploymentManagerDeploymentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DeploymentManagerDeployment].
func (dmd *DeploymentManagerDeployment) Type() string {
	return "google_deployment_manager_deployment"
}

// LocalName returns the local name for [DeploymentManagerDeployment].
func (dmd *DeploymentManagerDeployment) LocalName() string {
	return dmd.Name
}

// Configuration returns the configuration (args) for [DeploymentManagerDeployment].
func (dmd *DeploymentManagerDeployment) Configuration() interface{} {
	return dmd.Args
}

// DependOn is used for other resources to depend on [DeploymentManagerDeployment].
func (dmd *DeploymentManagerDeployment) DependOn() terra.Reference {
	return terra.ReferenceResource(dmd)
}

// Dependencies returns the list of resources [DeploymentManagerDeployment] depends_on.
func (dmd *DeploymentManagerDeployment) Dependencies() terra.Dependencies {
	return dmd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DeploymentManagerDeployment].
func (dmd *DeploymentManagerDeployment) LifecycleManagement() *terra.Lifecycle {
	return dmd.Lifecycle
}

// Attributes returns the attributes for [DeploymentManagerDeployment].
func (dmd *DeploymentManagerDeployment) Attributes() deploymentManagerDeploymentAttributes {
	return deploymentManagerDeploymentAttributes{ref: terra.ReferenceResource(dmd)}
}

// ImportState imports the given attribute values into [DeploymentManagerDeployment]'s state.
func (dmd *DeploymentManagerDeployment) ImportState(av io.Reader) error {
	dmd.state = &deploymentManagerDeploymentState{}
	if err := json.NewDecoder(av).Decode(dmd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dmd.Type(), dmd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DeploymentManagerDeployment] has state.
func (dmd *DeploymentManagerDeployment) State() (*deploymentManagerDeploymentState, bool) {
	return dmd.state, dmd.state != nil
}

// StateMust returns the state for [DeploymentManagerDeployment]. Panics if the state is nil.
func (dmd *DeploymentManagerDeployment) StateMust() *deploymentManagerDeploymentState {
	if dmd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dmd.Type(), dmd.LocalName()))
	}
	return dmd.state
}

// DeploymentManagerDeploymentArgs contains the configurations for google_deployment_manager_deployment.
type DeploymentManagerDeploymentArgs struct {
	// CreatePolicy: string, optional
	CreatePolicy terra.StringValue `hcl:"create_policy,attr"`
	// DeletePolicy: string, optional
	DeletePolicy terra.StringValue `hcl:"delete_policy,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Preview: bool, optional
	Preview terra.BoolValue `hcl:"preview,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Labels: min=0
	Labels []deploymentmanagerdeployment.Labels `hcl:"labels,block" validate:"min=0"`
	// Target: required
	Target *deploymentmanagerdeployment.Target `hcl:"target,block" validate:"required"`
	// Timeouts: optional
	Timeouts *deploymentmanagerdeployment.Timeouts `hcl:"timeouts,block"`
}
type deploymentManagerDeploymentAttributes struct {
	ref terra.Reference
}

// CreatePolicy returns a reference to field create_policy of google_deployment_manager_deployment.
func (dmd deploymentManagerDeploymentAttributes) CreatePolicy() terra.StringValue {
	return terra.ReferenceAsString(dmd.ref.Append("create_policy"))
}

// DeletePolicy returns a reference to field delete_policy of google_deployment_manager_deployment.
func (dmd deploymentManagerDeploymentAttributes) DeletePolicy() terra.StringValue {
	return terra.ReferenceAsString(dmd.ref.Append("delete_policy"))
}

// DeploymentId returns a reference to field deployment_id of google_deployment_manager_deployment.
func (dmd deploymentManagerDeploymentAttributes) DeploymentId() terra.StringValue {
	return terra.ReferenceAsString(dmd.ref.Append("deployment_id"))
}

// Description returns a reference to field description of google_deployment_manager_deployment.
func (dmd deploymentManagerDeploymentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dmd.ref.Append("description"))
}

// Id returns a reference to field id of google_deployment_manager_deployment.
func (dmd deploymentManagerDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dmd.ref.Append("id"))
}

// Manifest returns a reference to field manifest of google_deployment_manager_deployment.
func (dmd deploymentManagerDeploymentAttributes) Manifest() terra.StringValue {
	return terra.ReferenceAsString(dmd.ref.Append("manifest"))
}

// Name returns a reference to field name of google_deployment_manager_deployment.
func (dmd deploymentManagerDeploymentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dmd.ref.Append("name"))
}

// Preview returns a reference to field preview of google_deployment_manager_deployment.
func (dmd deploymentManagerDeploymentAttributes) Preview() terra.BoolValue {
	return terra.ReferenceAsBool(dmd.ref.Append("preview"))
}

// Project returns a reference to field project of google_deployment_manager_deployment.
func (dmd deploymentManagerDeploymentAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dmd.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_deployment_manager_deployment.
func (dmd deploymentManagerDeploymentAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(dmd.ref.Append("self_link"))
}

func (dmd deploymentManagerDeploymentAttributes) Labels() terra.SetValue[deploymentmanagerdeployment.LabelsAttributes] {
	return terra.ReferenceAsSet[deploymentmanagerdeployment.LabelsAttributes](dmd.ref.Append("labels"))
}

func (dmd deploymentManagerDeploymentAttributes) Target() terra.ListValue[deploymentmanagerdeployment.TargetAttributes] {
	return terra.ReferenceAsList[deploymentmanagerdeployment.TargetAttributes](dmd.ref.Append("target"))
}

func (dmd deploymentManagerDeploymentAttributes) Timeouts() deploymentmanagerdeployment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[deploymentmanagerdeployment.TimeoutsAttributes](dmd.ref.Append("timeouts"))
}

type deploymentManagerDeploymentState struct {
	CreatePolicy string                                     `json:"create_policy"`
	DeletePolicy string                                     `json:"delete_policy"`
	DeploymentId string                                     `json:"deployment_id"`
	Description  string                                     `json:"description"`
	Id           string                                     `json:"id"`
	Manifest     string                                     `json:"manifest"`
	Name         string                                     `json:"name"`
	Preview      bool                                       `json:"preview"`
	Project      string                                     `json:"project"`
	SelfLink     string                                     `json:"self_link"`
	Labels       []deploymentmanagerdeployment.LabelsState  `json:"labels"`
	Target       []deploymentmanagerdeployment.TargetState  `json:"target"`
	Timeouts     *deploymentmanagerdeployment.TimeoutsState `json:"timeouts"`
}
