// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	codecatalystdevenvironment "github.com/golingon/terraproviders/aws/5.44.0/codecatalystdevenvironment"
	"io"
)

// NewCodecatalystDevEnvironment creates a new instance of [CodecatalystDevEnvironment].
func NewCodecatalystDevEnvironment(name string, args CodecatalystDevEnvironmentArgs) *CodecatalystDevEnvironment {
	return &CodecatalystDevEnvironment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CodecatalystDevEnvironment)(nil)

// CodecatalystDevEnvironment represents the Terraform resource aws_codecatalyst_dev_environment.
type CodecatalystDevEnvironment struct {
	Name      string
	Args      CodecatalystDevEnvironmentArgs
	state     *codecatalystDevEnvironmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CodecatalystDevEnvironment].
func (cde *CodecatalystDevEnvironment) Type() string {
	return "aws_codecatalyst_dev_environment"
}

// LocalName returns the local name for [CodecatalystDevEnvironment].
func (cde *CodecatalystDevEnvironment) LocalName() string {
	return cde.Name
}

// Configuration returns the configuration (args) for [CodecatalystDevEnvironment].
func (cde *CodecatalystDevEnvironment) Configuration() interface{} {
	return cde.Args
}

// DependOn is used for other resources to depend on [CodecatalystDevEnvironment].
func (cde *CodecatalystDevEnvironment) DependOn() terra.Reference {
	return terra.ReferenceResource(cde)
}

// Dependencies returns the list of resources [CodecatalystDevEnvironment] depends_on.
func (cde *CodecatalystDevEnvironment) Dependencies() terra.Dependencies {
	return cde.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CodecatalystDevEnvironment].
func (cde *CodecatalystDevEnvironment) LifecycleManagement() *terra.Lifecycle {
	return cde.Lifecycle
}

// Attributes returns the attributes for [CodecatalystDevEnvironment].
func (cde *CodecatalystDevEnvironment) Attributes() codecatalystDevEnvironmentAttributes {
	return codecatalystDevEnvironmentAttributes{ref: terra.ReferenceResource(cde)}
}

// ImportState imports the given attribute values into [CodecatalystDevEnvironment]'s state.
func (cde *CodecatalystDevEnvironment) ImportState(av io.Reader) error {
	cde.state = &codecatalystDevEnvironmentState{}
	if err := json.NewDecoder(av).Decode(cde.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cde.Type(), cde.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CodecatalystDevEnvironment] has state.
func (cde *CodecatalystDevEnvironment) State() (*codecatalystDevEnvironmentState, bool) {
	return cde.state, cde.state != nil
}

// StateMust returns the state for [CodecatalystDevEnvironment]. Panics if the state is nil.
func (cde *CodecatalystDevEnvironment) StateMust() *codecatalystDevEnvironmentState {
	if cde.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cde.Type(), cde.LocalName()))
	}
	return cde.state
}

// CodecatalystDevEnvironmentArgs contains the configurations for aws_codecatalyst_dev_environment.
type CodecatalystDevEnvironmentArgs struct {
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
	Ides *codecatalystdevenvironment.Ides `hcl:"ides,block" validate:"required"`
	// PersistentStorage: required
	PersistentStorage *codecatalystdevenvironment.PersistentStorage `hcl:"persistent_storage,block" validate:"required"`
	// Repositories: min=0,max=100
	Repositories []codecatalystdevenvironment.Repositories `hcl:"repositories,block" validate:"min=0,max=100"`
	// Timeouts: optional
	Timeouts *codecatalystdevenvironment.Timeouts `hcl:"timeouts,block"`
}
type codecatalystDevEnvironmentAttributes struct {
	ref terra.Reference
}

// Alias returns a reference to field alias of aws_codecatalyst_dev_environment.
func (cde codecatalystDevEnvironmentAttributes) Alias() terra.StringValue {
	return terra.ReferenceAsString(cde.ref.Append("alias"))
}

// Id returns a reference to field id of aws_codecatalyst_dev_environment.
func (cde codecatalystDevEnvironmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cde.ref.Append("id"))
}

// InactivityTimeoutMinutes returns a reference to field inactivity_timeout_minutes of aws_codecatalyst_dev_environment.
func (cde codecatalystDevEnvironmentAttributes) InactivityTimeoutMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(cde.ref.Append("inactivity_timeout_minutes"))
}

// InstanceType returns a reference to field instance_type of aws_codecatalyst_dev_environment.
func (cde codecatalystDevEnvironmentAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(cde.ref.Append("instance_type"))
}

// ProjectName returns a reference to field project_name of aws_codecatalyst_dev_environment.
func (cde codecatalystDevEnvironmentAttributes) ProjectName() terra.StringValue {
	return terra.ReferenceAsString(cde.ref.Append("project_name"))
}

// SpaceName returns a reference to field space_name of aws_codecatalyst_dev_environment.
func (cde codecatalystDevEnvironmentAttributes) SpaceName() terra.StringValue {
	return terra.ReferenceAsString(cde.ref.Append("space_name"))
}

func (cde codecatalystDevEnvironmentAttributes) Ides() terra.ListValue[codecatalystdevenvironment.IdesAttributes] {
	return terra.ReferenceAsList[codecatalystdevenvironment.IdesAttributes](cde.ref.Append("ides"))
}

func (cde codecatalystDevEnvironmentAttributes) PersistentStorage() terra.ListValue[codecatalystdevenvironment.PersistentStorageAttributes] {
	return terra.ReferenceAsList[codecatalystdevenvironment.PersistentStorageAttributes](cde.ref.Append("persistent_storage"))
}

func (cde codecatalystDevEnvironmentAttributes) Repositories() terra.ListValue[codecatalystdevenvironment.RepositoriesAttributes] {
	return terra.ReferenceAsList[codecatalystdevenvironment.RepositoriesAttributes](cde.ref.Append("repositories"))
}

func (cde codecatalystDevEnvironmentAttributes) Timeouts() codecatalystdevenvironment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[codecatalystdevenvironment.TimeoutsAttributes](cde.ref.Append("timeouts"))
}

type codecatalystDevEnvironmentState struct {
	Alias                    string                                              `json:"alias"`
	Id                       string                                              `json:"id"`
	InactivityTimeoutMinutes float64                                             `json:"inactivity_timeout_minutes"`
	InstanceType             string                                              `json:"instance_type"`
	ProjectName              string                                              `json:"project_name"`
	SpaceName                string                                              `json:"space_name"`
	Ides                     []codecatalystdevenvironment.IdesState              `json:"ides"`
	PersistentStorage        []codecatalystdevenvironment.PersistentStorageState `json:"persistent_storage"`
	Repositories             []codecatalystdevenvironment.RepositoriesState      `json:"repositories"`
	Timeouts                 *codecatalystdevenvironment.TimeoutsState           `json:"timeouts"`
}
