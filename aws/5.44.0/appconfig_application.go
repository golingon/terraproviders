// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewAppconfigApplication creates a new instance of [AppconfigApplication].
func NewAppconfigApplication(name string, args AppconfigApplicationArgs) *AppconfigApplication {
	return &AppconfigApplication{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppconfigApplication)(nil)

// AppconfigApplication represents the Terraform resource aws_appconfig_application.
type AppconfigApplication struct {
	Name      string
	Args      AppconfigApplicationArgs
	state     *appconfigApplicationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppconfigApplication].
func (aa *AppconfigApplication) Type() string {
	return "aws_appconfig_application"
}

// LocalName returns the local name for [AppconfigApplication].
func (aa *AppconfigApplication) LocalName() string {
	return aa.Name
}

// Configuration returns the configuration (args) for [AppconfigApplication].
func (aa *AppconfigApplication) Configuration() interface{} {
	return aa.Args
}

// DependOn is used for other resources to depend on [AppconfigApplication].
func (aa *AppconfigApplication) DependOn() terra.Reference {
	return terra.ReferenceResource(aa)
}

// Dependencies returns the list of resources [AppconfigApplication] depends_on.
func (aa *AppconfigApplication) Dependencies() terra.Dependencies {
	return aa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppconfigApplication].
func (aa *AppconfigApplication) LifecycleManagement() *terra.Lifecycle {
	return aa.Lifecycle
}

// Attributes returns the attributes for [AppconfigApplication].
func (aa *AppconfigApplication) Attributes() appconfigApplicationAttributes {
	return appconfigApplicationAttributes{ref: terra.ReferenceResource(aa)}
}

// ImportState imports the given attribute values into [AppconfigApplication]'s state.
func (aa *AppconfigApplication) ImportState(av io.Reader) error {
	aa.state = &appconfigApplicationState{}
	if err := json.NewDecoder(av).Decode(aa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aa.Type(), aa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppconfigApplication] has state.
func (aa *AppconfigApplication) State() (*appconfigApplicationState, bool) {
	return aa.state, aa.state != nil
}

// StateMust returns the state for [AppconfigApplication]. Panics if the state is nil.
func (aa *AppconfigApplication) StateMust() *appconfigApplicationState {
	if aa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aa.Type(), aa.LocalName()))
	}
	return aa.state
}

// AppconfigApplicationArgs contains the configurations for aws_appconfig_application.
type AppconfigApplicationArgs struct {
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
}
type appconfigApplicationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_appconfig_application.
func (aa appconfigApplicationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("arn"))
}

// Description returns a reference to field description of aws_appconfig_application.
func (aa appconfigApplicationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("description"))
}

// Id returns a reference to field id of aws_appconfig_application.
func (aa appconfigApplicationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("id"))
}

// Name returns a reference to field name of aws_appconfig_application.
func (aa appconfigApplicationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_appconfig_application.
func (aa appconfigApplicationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aa.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_appconfig_application.
func (aa appconfigApplicationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aa.ref.Append("tags_all"))
}

type appconfigApplicationState struct {
	Arn         string            `json:"arn"`
	Description string            `json:"description"`
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Tags        map[string]string `json:"tags"`
	TagsAll     map[string]string `json:"tags_all"`
}
