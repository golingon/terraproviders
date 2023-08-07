// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	iotthingtype "github.com/golingon/terraproviders/aws/5.11.0/iotthingtype"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIotThingType creates a new instance of [IotThingType].
func NewIotThingType(name string, args IotThingTypeArgs) *IotThingType {
	return &IotThingType{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IotThingType)(nil)

// IotThingType represents the Terraform resource aws_iot_thing_type.
type IotThingType struct {
	Name      string
	Args      IotThingTypeArgs
	state     *iotThingTypeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IotThingType].
func (itt *IotThingType) Type() string {
	return "aws_iot_thing_type"
}

// LocalName returns the local name for [IotThingType].
func (itt *IotThingType) LocalName() string {
	return itt.Name
}

// Configuration returns the configuration (args) for [IotThingType].
func (itt *IotThingType) Configuration() interface{} {
	return itt.Args
}

// DependOn is used for other resources to depend on [IotThingType].
func (itt *IotThingType) DependOn() terra.Reference {
	return terra.ReferenceResource(itt)
}

// Dependencies returns the list of resources [IotThingType] depends_on.
func (itt *IotThingType) Dependencies() terra.Dependencies {
	return itt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IotThingType].
func (itt *IotThingType) LifecycleManagement() *terra.Lifecycle {
	return itt.Lifecycle
}

// Attributes returns the attributes for [IotThingType].
func (itt *IotThingType) Attributes() iotThingTypeAttributes {
	return iotThingTypeAttributes{ref: terra.ReferenceResource(itt)}
}

// ImportState imports the given attribute values into [IotThingType]'s state.
func (itt *IotThingType) ImportState(av io.Reader) error {
	itt.state = &iotThingTypeState{}
	if err := json.NewDecoder(av).Decode(itt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", itt.Type(), itt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IotThingType] has state.
func (itt *IotThingType) State() (*iotThingTypeState, bool) {
	return itt.state, itt.state != nil
}

// StateMust returns the state for [IotThingType]. Panics if the state is nil.
func (itt *IotThingType) StateMust() *iotThingTypeState {
	if itt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", itt.Type(), itt.LocalName()))
	}
	return itt.state
}

// IotThingTypeArgs contains the configurations for aws_iot_thing_type.
type IotThingTypeArgs struct {
	// Deprecated: bool, optional
	Deprecated terra.BoolValue `hcl:"deprecated,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Properties: optional
	Properties *iotthingtype.Properties `hcl:"properties,block"`
}
type iotThingTypeAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_iot_thing_type.
func (itt iotThingTypeAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(itt.ref.Append("arn"))
}

// Deprecated returns a reference to field deprecated of aws_iot_thing_type.
func (itt iotThingTypeAttributes) Deprecated() terra.BoolValue {
	return terra.ReferenceAsBool(itt.ref.Append("deprecated"))
}

// Id returns a reference to field id of aws_iot_thing_type.
func (itt iotThingTypeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(itt.ref.Append("id"))
}

// Name returns a reference to field name of aws_iot_thing_type.
func (itt iotThingTypeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(itt.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_iot_thing_type.
func (itt iotThingTypeAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](itt.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_iot_thing_type.
func (itt iotThingTypeAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](itt.ref.Append("tags_all"))
}

func (itt iotThingTypeAttributes) Properties() terra.ListValue[iotthingtype.PropertiesAttributes] {
	return terra.ReferenceAsList[iotthingtype.PropertiesAttributes](itt.ref.Append("properties"))
}

type iotThingTypeState struct {
	Arn        string                         `json:"arn"`
	Deprecated bool                           `json:"deprecated"`
	Id         string                         `json:"id"`
	Name       string                         `json:"name"`
	Tags       map[string]string              `json:"tags"`
	TagsAll    map[string]string              `json:"tags_all"`
	Properties []iotthingtype.PropertiesState `json:"properties"`
}
