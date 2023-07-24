// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	emrserverlessapplication "github.com/golingon/terraproviders/aws/5.9.0/emrserverlessapplication"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEmrserverlessApplication creates a new instance of [EmrserverlessApplication].
func NewEmrserverlessApplication(name string, args EmrserverlessApplicationArgs) *EmrserverlessApplication {
	return &EmrserverlessApplication{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EmrserverlessApplication)(nil)

// EmrserverlessApplication represents the Terraform resource aws_emrserverless_application.
type EmrserverlessApplication struct {
	Name      string
	Args      EmrserverlessApplicationArgs
	state     *emrserverlessApplicationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EmrserverlessApplication].
func (ea *EmrserverlessApplication) Type() string {
	return "aws_emrserverless_application"
}

// LocalName returns the local name for [EmrserverlessApplication].
func (ea *EmrserverlessApplication) LocalName() string {
	return ea.Name
}

// Configuration returns the configuration (args) for [EmrserverlessApplication].
func (ea *EmrserverlessApplication) Configuration() interface{} {
	return ea.Args
}

// DependOn is used for other resources to depend on [EmrserverlessApplication].
func (ea *EmrserverlessApplication) DependOn() terra.Reference {
	return terra.ReferenceResource(ea)
}

// Dependencies returns the list of resources [EmrserverlessApplication] depends_on.
func (ea *EmrserverlessApplication) Dependencies() terra.Dependencies {
	return ea.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EmrserverlessApplication].
func (ea *EmrserverlessApplication) LifecycleManagement() *terra.Lifecycle {
	return ea.Lifecycle
}

// Attributes returns the attributes for [EmrserverlessApplication].
func (ea *EmrserverlessApplication) Attributes() emrserverlessApplicationAttributes {
	return emrserverlessApplicationAttributes{ref: terra.ReferenceResource(ea)}
}

// ImportState imports the given attribute values into [EmrserverlessApplication]'s state.
func (ea *EmrserverlessApplication) ImportState(av io.Reader) error {
	ea.state = &emrserverlessApplicationState{}
	if err := json.NewDecoder(av).Decode(ea.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ea.Type(), ea.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EmrserverlessApplication] has state.
func (ea *EmrserverlessApplication) State() (*emrserverlessApplicationState, bool) {
	return ea.state, ea.state != nil
}

// StateMust returns the state for [EmrserverlessApplication]. Panics if the state is nil.
func (ea *EmrserverlessApplication) StateMust() *emrserverlessApplicationState {
	if ea.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ea.Type(), ea.LocalName()))
	}
	return ea.state
}

// EmrserverlessApplicationArgs contains the configurations for aws_emrserverless_application.
type EmrserverlessApplicationArgs struct {
	// Architecture: string, optional
	Architecture terra.StringValue `hcl:"architecture,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ReleaseLabel: string, required
	ReleaseLabel terra.StringValue `hcl:"release_label,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// AutoStartConfiguration: optional
	AutoStartConfiguration *emrserverlessapplication.AutoStartConfiguration `hcl:"auto_start_configuration,block"`
	// AutoStopConfiguration: optional
	AutoStopConfiguration *emrserverlessapplication.AutoStopConfiguration `hcl:"auto_stop_configuration,block"`
	// ImageConfiguration: optional
	ImageConfiguration *emrserverlessapplication.ImageConfiguration `hcl:"image_configuration,block"`
	// InitialCapacity: min=0
	InitialCapacity []emrserverlessapplication.InitialCapacity `hcl:"initial_capacity,block" validate:"min=0"`
	// MaximumCapacity: optional
	MaximumCapacity *emrserverlessapplication.MaximumCapacity `hcl:"maximum_capacity,block"`
	// NetworkConfiguration: optional
	NetworkConfiguration *emrserverlessapplication.NetworkConfiguration `hcl:"network_configuration,block"`
}
type emrserverlessApplicationAttributes struct {
	ref terra.Reference
}

// Architecture returns a reference to field architecture of aws_emrserverless_application.
func (ea emrserverlessApplicationAttributes) Architecture() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("architecture"))
}

// Arn returns a reference to field arn of aws_emrserverless_application.
func (ea emrserverlessApplicationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("arn"))
}

// Id returns a reference to field id of aws_emrserverless_application.
func (ea emrserverlessApplicationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("id"))
}

// Name returns a reference to field name of aws_emrserverless_application.
func (ea emrserverlessApplicationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("name"))
}

// ReleaseLabel returns a reference to field release_label of aws_emrserverless_application.
func (ea emrserverlessApplicationAttributes) ReleaseLabel() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("release_label"))
}

// Tags returns a reference to field tags of aws_emrserverless_application.
func (ea emrserverlessApplicationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ea.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_emrserverless_application.
func (ea emrserverlessApplicationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ea.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_emrserverless_application.
func (ea emrserverlessApplicationAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("type"))
}

func (ea emrserverlessApplicationAttributes) AutoStartConfiguration() terra.ListValue[emrserverlessapplication.AutoStartConfigurationAttributes] {
	return terra.ReferenceAsList[emrserverlessapplication.AutoStartConfigurationAttributes](ea.ref.Append("auto_start_configuration"))
}

func (ea emrserverlessApplicationAttributes) AutoStopConfiguration() terra.ListValue[emrserverlessapplication.AutoStopConfigurationAttributes] {
	return terra.ReferenceAsList[emrserverlessapplication.AutoStopConfigurationAttributes](ea.ref.Append("auto_stop_configuration"))
}

func (ea emrserverlessApplicationAttributes) ImageConfiguration() terra.ListValue[emrserverlessapplication.ImageConfigurationAttributes] {
	return terra.ReferenceAsList[emrserverlessapplication.ImageConfigurationAttributes](ea.ref.Append("image_configuration"))
}

func (ea emrserverlessApplicationAttributes) InitialCapacity() terra.SetValue[emrserverlessapplication.InitialCapacityAttributes] {
	return terra.ReferenceAsSet[emrserverlessapplication.InitialCapacityAttributes](ea.ref.Append("initial_capacity"))
}

func (ea emrserverlessApplicationAttributes) MaximumCapacity() terra.ListValue[emrserverlessapplication.MaximumCapacityAttributes] {
	return terra.ReferenceAsList[emrserverlessapplication.MaximumCapacityAttributes](ea.ref.Append("maximum_capacity"))
}

func (ea emrserverlessApplicationAttributes) NetworkConfiguration() terra.ListValue[emrserverlessapplication.NetworkConfigurationAttributes] {
	return terra.ReferenceAsList[emrserverlessapplication.NetworkConfigurationAttributes](ea.ref.Append("network_configuration"))
}

type emrserverlessApplicationState struct {
	Architecture           string                                                 `json:"architecture"`
	Arn                    string                                                 `json:"arn"`
	Id                     string                                                 `json:"id"`
	Name                   string                                                 `json:"name"`
	ReleaseLabel           string                                                 `json:"release_label"`
	Tags                   map[string]string                                      `json:"tags"`
	TagsAll                map[string]string                                      `json:"tags_all"`
	Type                   string                                                 `json:"type"`
	AutoStartConfiguration []emrserverlessapplication.AutoStartConfigurationState `json:"auto_start_configuration"`
	AutoStopConfiguration  []emrserverlessapplication.AutoStopConfigurationState  `json:"auto_stop_configuration"`
	ImageConfiguration     []emrserverlessapplication.ImageConfigurationState     `json:"image_configuration"`
	InitialCapacity        []emrserverlessapplication.InitialCapacityState        `json:"initial_capacity"`
	MaximumCapacity        []emrserverlessapplication.MaximumCapacityState        `json:"maximum_capacity"`
	NetworkConfiguration   []emrserverlessapplication.NetworkConfigurationState   `json:"network_configuration"`
}
