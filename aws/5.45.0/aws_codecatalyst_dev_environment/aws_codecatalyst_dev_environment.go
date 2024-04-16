// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_codecatalyst_dev_environment

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

// Resource represents the Terraform resource aws_codecatalyst_dev_environment.
type Resource struct {
	Name      string
	Args      Args
	state     *awsCodecatalystDevEnvironmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (acde *Resource) Type() string {
	return "aws_codecatalyst_dev_environment"
}

// LocalName returns the local name for [Resource].
func (acde *Resource) LocalName() string {
	return acde.Name
}

// Configuration returns the configuration (args) for [Resource].
func (acde *Resource) Configuration() interface{} {
	return acde.Args
}

// DependOn is used for other resources to depend on [Resource].
func (acde *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(acde)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (acde *Resource) Dependencies() terra.Dependencies {
	return acde.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (acde *Resource) LifecycleManagement() *terra.Lifecycle {
	return acde.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (acde *Resource) Attributes() awsCodecatalystDevEnvironmentAttributes {
	return awsCodecatalystDevEnvironmentAttributes{ref: terra.ReferenceResource(acde)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (acde *Resource) ImportState(state io.Reader) error {
	acde.state = &awsCodecatalystDevEnvironmentState{}
	if err := json.NewDecoder(state).Decode(acde.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acde.Type(), acde.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (acde *Resource) State() (*awsCodecatalystDevEnvironmentState, bool) {
	return acde.state, acde.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (acde *Resource) StateMust() *awsCodecatalystDevEnvironmentState {
	if acde.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acde.Type(), acde.LocalName()))
	}
	return acde.state
}

// Args contains the configurations for aws_codecatalyst_dev_environment.
type Args struct {
	// Alias: string, optional
	Alias terra.StringValue `hcl:"alias,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InactivityTimeoutMinutes: number, optional
	InactivityTimeoutMinutes terra.NumberValue `hcl:"inactivity_timeout_minutes,attr"`
	// InstanceType: string, required
	InstanceType terra.StringValue `hcl:"instance_type,attr" validate:"required"`
	// ProjectName: string, required
	ProjectName terra.StringValue `hcl:"project_name,attr" validate:"required"`
	// SpaceName: string, required
	SpaceName terra.StringValue `hcl:"space_name,attr" validate:"required"`
	// Ides: required
	Ides *Ides `hcl:"ides,block" validate:"required"`
	// PersistentStorage: required
	PersistentStorage *PersistentStorage `hcl:"persistent_storage,block" validate:"required"`
	// Repositories: min=0,max=100
	Repositories []Repositories `hcl:"repositories,block" validate:"min=0,max=100"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsCodecatalystDevEnvironmentAttributes struct {
	ref terra.Reference
}

// Alias returns a reference to field alias of aws_codecatalyst_dev_environment.
func (acde awsCodecatalystDevEnvironmentAttributes) Alias() terra.StringValue {
	return terra.ReferenceAsString(acde.ref.Append("alias"))
}

// Id returns a reference to field id of aws_codecatalyst_dev_environment.
func (acde awsCodecatalystDevEnvironmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acde.ref.Append("id"))
}

// InactivityTimeoutMinutes returns a reference to field inactivity_timeout_minutes of aws_codecatalyst_dev_environment.
func (acde awsCodecatalystDevEnvironmentAttributes) InactivityTimeoutMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(acde.ref.Append("inactivity_timeout_minutes"))
}

// InstanceType returns a reference to field instance_type of aws_codecatalyst_dev_environment.
func (acde awsCodecatalystDevEnvironmentAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(acde.ref.Append("instance_type"))
}

// ProjectName returns a reference to field project_name of aws_codecatalyst_dev_environment.
func (acde awsCodecatalystDevEnvironmentAttributes) ProjectName() terra.StringValue {
	return terra.ReferenceAsString(acde.ref.Append("project_name"))
}

// SpaceName returns a reference to field space_name of aws_codecatalyst_dev_environment.
func (acde awsCodecatalystDevEnvironmentAttributes) SpaceName() terra.StringValue {
	return terra.ReferenceAsString(acde.ref.Append("space_name"))
}

func (acde awsCodecatalystDevEnvironmentAttributes) Ides() terra.ListValue[IdesAttributes] {
	return terra.ReferenceAsList[IdesAttributes](acde.ref.Append("ides"))
}

func (acde awsCodecatalystDevEnvironmentAttributes) PersistentStorage() terra.ListValue[PersistentStorageAttributes] {
	return terra.ReferenceAsList[PersistentStorageAttributes](acde.ref.Append("persistent_storage"))
}

func (acde awsCodecatalystDevEnvironmentAttributes) Repositories() terra.ListValue[RepositoriesAttributes] {
	return terra.ReferenceAsList[RepositoriesAttributes](acde.ref.Append("repositories"))
}

func (acde awsCodecatalystDevEnvironmentAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](acde.ref.Append("timeouts"))
}

type awsCodecatalystDevEnvironmentState struct {
	Alias                    string                   `json:"alias"`
	Id                       string                   `json:"id"`
	InactivityTimeoutMinutes float64                  `json:"inactivity_timeout_minutes"`
	InstanceType             string                   `json:"instance_type"`
	ProjectName              string                   `json:"project_name"`
	SpaceName                string                   `json:"space_name"`
	Ides                     []IdesState              `json:"ides"`
	PersistentStorage        []PersistentStorageState `json:"persistent_storage"`
	Repositories             []RepositoriesState      `json:"repositories"`
	Timeouts                 *TimeoutsState           `json:"timeouts"`
}
