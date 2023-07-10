// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	appconfigextension "github.com/golingon/terraproviders/aws/5.7.0/appconfigextension"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppconfigExtension creates a new instance of [AppconfigExtension].
func NewAppconfigExtension(name string, args AppconfigExtensionArgs) *AppconfigExtension {
	return &AppconfigExtension{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppconfigExtension)(nil)

// AppconfigExtension represents the Terraform resource aws_appconfig_extension.
type AppconfigExtension struct {
	Name      string
	Args      AppconfigExtensionArgs
	state     *appconfigExtensionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppconfigExtension].
func (ae *AppconfigExtension) Type() string {
	return "aws_appconfig_extension"
}

// LocalName returns the local name for [AppconfigExtension].
func (ae *AppconfigExtension) LocalName() string {
	return ae.Name
}

// Configuration returns the configuration (args) for [AppconfigExtension].
func (ae *AppconfigExtension) Configuration() interface{} {
	return ae.Args
}

// DependOn is used for other resources to depend on [AppconfigExtension].
func (ae *AppconfigExtension) DependOn() terra.Reference {
	return terra.ReferenceResource(ae)
}

// Dependencies returns the list of resources [AppconfigExtension] depends_on.
func (ae *AppconfigExtension) Dependencies() terra.Dependencies {
	return ae.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppconfigExtension].
func (ae *AppconfigExtension) LifecycleManagement() *terra.Lifecycle {
	return ae.Lifecycle
}

// Attributes returns the attributes for [AppconfigExtension].
func (ae *AppconfigExtension) Attributes() appconfigExtensionAttributes {
	return appconfigExtensionAttributes{ref: terra.ReferenceResource(ae)}
}

// ImportState imports the given attribute values into [AppconfigExtension]'s state.
func (ae *AppconfigExtension) ImportState(av io.Reader) error {
	ae.state = &appconfigExtensionState{}
	if err := json.NewDecoder(av).Decode(ae.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ae.Type(), ae.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppconfigExtension] has state.
func (ae *AppconfigExtension) State() (*appconfigExtensionState, bool) {
	return ae.state, ae.state != nil
}

// StateMust returns the state for [AppconfigExtension]. Panics if the state is nil.
func (ae *AppconfigExtension) StateMust() *appconfigExtensionState {
	if ae.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ae.Type(), ae.LocalName()))
	}
	return ae.state
}

// AppconfigExtensionArgs contains the configurations for aws_appconfig_extension.
type AppconfigExtensionArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// ActionPoint: min=1
	ActionPoint []appconfigextension.ActionPoint `hcl:"action_point,block" validate:"min=1"`
	// Parameter: min=0
	Parameter []appconfigextension.Parameter `hcl:"parameter,block" validate:"min=0"`
}
type appconfigExtensionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_appconfig_extension.
func (ae appconfigExtensionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("arn"))
}

// Description returns a reference to field description of aws_appconfig_extension.
func (ae appconfigExtensionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("description"))
}

// Id returns a reference to field id of aws_appconfig_extension.
func (ae appconfigExtensionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("id"))
}

// Name returns a reference to field name of aws_appconfig_extension.
func (ae appconfigExtensionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_appconfig_extension.
func (ae appconfigExtensionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ae.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_appconfig_extension.
func (ae appconfigExtensionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ae.ref.Append("tags_all"))
}

// Version returns a reference to field version of aws_appconfig_extension.
func (ae appconfigExtensionAttributes) Version() terra.NumberValue {
	return terra.ReferenceAsNumber(ae.ref.Append("version"))
}

func (ae appconfigExtensionAttributes) ActionPoint() terra.SetValue[appconfigextension.ActionPointAttributes] {
	return terra.ReferenceAsSet[appconfigextension.ActionPointAttributes](ae.ref.Append("action_point"))
}

func (ae appconfigExtensionAttributes) Parameter() terra.SetValue[appconfigextension.ParameterAttributes] {
	return terra.ReferenceAsSet[appconfigextension.ParameterAttributes](ae.ref.Append("parameter"))
}

type appconfigExtensionState struct {
	Arn         string                                `json:"arn"`
	Description string                                `json:"description"`
	Id          string                                `json:"id"`
	Name        string                                `json:"name"`
	Tags        map[string]string                     `json:"tags"`
	TagsAll     map[string]string                     `json:"tags_all"`
	Version     float64                               `json:"version"`
	ActionPoint []appconfigextension.ActionPointState `json:"action_point"`
	Parameter   []appconfigextension.ParameterState   `json:"parameter"`
}
