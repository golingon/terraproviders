// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	athenaworkgroup "github.com/golingon/terraproviders/aws/5.7.0/athenaworkgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAthenaWorkgroup creates a new instance of [AthenaWorkgroup].
func NewAthenaWorkgroup(name string, args AthenaWorkgroupArgs) *AthenaWorkgroup {
	return &AthenaWorkgroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AthenaWorkgroup)(nil)

// AthenaWorkgroup represents the Terraform resource aws_athena_workgroup.
type AthenaWorkgroup struct {
	Name      string
	Args      AthenaWorkgroupArgs
	state     *athenaWorkgroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AthenaWorkgroup].
func (aw *AthenaWorkgroup) Type() string {
	return "aws_athena_workgroup"
}

// LocalName returns the local name for [AthenaWorkgroup].
func (aw *AthenaWorkgroup) LocalName() string {
	return aw.Name
}

// Configuration returns the configuration (args) for [AthenaWorkgroup].
func (aw *AthenaWorkgroup) Configuration() interface{} {
	return aw.Args
}

// DependOn is used for other resources to depend on [AthenaWorkgroup].
func (aw *AthenaWorkgroup) DependOn() terra.Reference {
	return terra.ReferenceResource(aw)
}

// Dependencies returns the list of resources [AthenaWorkgroup] depends_on.
func (aw *AthenaWorkgroup) Dependencies() terra.Dependencies {
	return aw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AthenaWorkgroup].
func (aw *AthenaWorkgroup) LifecycleManagement() *terra.Lifecycle {
	return aw.Lifecycle
}

// Attributes returns the attributes for [AthenaWorkgroup].
func (aw *AthenaWorkgroup) Attributes() athenaWorkgroupAttributes {
	return athenaWorkgroupAttributes{ref: terra.ReferenceResource(aw)}
}

// ImportState imports the given attribute values into [AthenaWorkgroup]'s state.
func (aw *AthenaWorkgroup) ImportState(av io.Reader) error {
	aw.state = &athenaWorkgroupState{}
	if err := json.NewDecoder(av).Decode(aw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aw.Type(), aw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AthenaWorkgroup] has state.
func (aw *AthenaWorkgroup) State() (*athenaWorkgroupState, bool) {
	return aw.state, aw.state != nil
}

// StateMust returns the state for [AthenaWorkgroup]. Panics if the state is nil.
func (aw *AthenaWorkgroup) StateMust() *athenaWorkgroupState {
	if aw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aw.Type(), aw.LocalName()))
	}
	return aw.state
}

// AthenaWorkgroupArgs contains the configurations for aws_athena_workgroup.
type AthenaWorkgroupArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// ForceDestroy: bool, optional
	ForceDestroy terra.BoolValue `hcl:"force_destroy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Configuration: optional
	Configuration *athenaworkgroup.Configuration `hcl:"configuration,block"`
}
type athenaWorkgroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_athena_workgroup.
func (aw athenaWorkgroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aw.ref.Append("arn"))
}

// Description returns a reference to field description of aws_athena_workgroup.
func (aw athenaWorkgroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aw.ref.Append("description"))
}

// ForceDestroy returns a reference to field force_destroy of aws_athena_workgroup.
func (aw athenaWorkgroupAttributes) ForceDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(aw.ref.Append("force_destroy"))
}

// Id returns a reference to field id of aws_athena_workgroup.
func (aw athenaWorkgroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aw.ref.Append("id"))
}

// Name returns a reference to field name of aws_athena_workgroup.
func (aw athenaWorkgroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aw.ref.Append("name"))
}

// State returns a reference to field state of aws_athena_workgroup.
func (aw athenaWorkgroupAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(aw.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_athena_workgroup.
func (aw athenaWorkgroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aw.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_athena_workgroup.
func (aw athenaWorkgroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aw.ref.Append("tags_all"))
}

func (aw athenaWorkgroupAttributes) Configuration() terra.ListValue[athenaworkgroup.ConfigurationAttributes] {
	return terra.ReferenceAsList[athenaworkgroup.ConfigurationAttributes](aw.ref.Append("configuration"))
}

type athenaWorkgroupState struct {
	Arn           string                               `json:"arn"`
	Description   string                               `json:"description"`
	ForceDestroy  bool                                 `json:"force_destroy"`
	Id            string                               `json:"id"`
	Name          string                               `json:"name"`
	State         string                               `json:"state"`
	Tags          map[string]string                    `json:"tags"`
	TagsAll       map[string]string                    `json:"tags_all"`
	Configuration []athenaworkgroup.ConfigurationState `json:"configuration"`
}
