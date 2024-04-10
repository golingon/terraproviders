// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewIotThing creates a new instance of [IotThing].
func NewIotThing(name string, args IotThingArgs) *IotThing {
	return &IotThing{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IotThing)(nil)

// IotThing represents the Terraform resource aws_iot_thing.
type IotThing struct {
	Name      string
	Args      IotThingArgs
	state     *iotThingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IotThing].
func (it *IotThing) Type() string {
	return "aws_iot_thing"
}

// LocalName returns the local name for [IotThing].
func (it *IotThing) LocalName() string {
	return it.Name
}

// Configuration returns the configuration (args) for [IotThing].
func (it *IotThing) Configuration() interface{} {
	return it.Args
}

// DependOn is used for other resources to depend on [IotThing].
func (it *IotThing) DependOn() terra.Reference {
	return terra.ReferenceResource(it)
}

// Dependencies returns the list of resources [IotThing] depends_on.
func (it *IotThing) Dependencies() terra.Dependencies {
	return it.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IotThing].
func (it *IotThing) LifecycleManagement() *terra.Lifecycle {
	return it.Lifecycle
}

// Attributes returns the attributes for [IotThing].
func (it *IotThing) Attributes() iotThingAttributes {
	return iotThingAttributes{ref: terra.ReferenceResource(it)}
}

// ImportState imports the given attribute values into [IotThing]'s state.
func (it *IotThing) ImportState(av io.Reader) error {
	it.state = &iotThingState{}
	if err := json.NewDecoder(av).Decode(it.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", it.Type(), it.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IotThing] has state.
func (it *IotThing) State() (*iotThingState, bool) {
	return it.state, it.state != nil
}

// StateMust returns the state for [IotThing]. Panics if the state is nil.
func (it *IotThing) StateMust() *iotThingState {
	if it.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", it.Type(), it.LocalName()))
	}
	return it.state
}

// IotThingArgs contains the configurations for aws_iot_thing.
type IotThingArgs struct {
	// Attributes: map of string, optional
	Attributes terra.MapValue[terra.StringValue] `hcl:"attributes,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ThingTypeName: string, optional
	ThingTypeName terra.StringValue `hcl:"thing_type_name,attr"`
}
type iotThingAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_iot_thing.
func (it iotThingAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(it.ref.Append("arn"))
}

// Attributes returns a reference to field attributes of aws_iot_thing.
func (it iotThingAttributes) Attributes() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](it.ref.Append("attributes"))
}

// DefaultClientId returns a reference to field default_client_id of aws_iot_thing.
func (it iotThingAttributes) DefaultClientId() terra.StringValue {
	return terra.ReferenceAsString(it.ref.Append("default_client_id"))
}

// Id returns a reference to field id of aws_iot_thing.
func (it iotThingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(it.ref.Append("id"))
}

// Name returns a reference to field name of aws_iot_thing.
func (it iotThingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(it.ref.Append("name"))
}

// ThingTypeName returns a reference to field thing_type_name of aws_iot_thing.
func (it iotThingAttributes) ThingTypeName() terra.StringValue {
	return terra.ReferenceAsString(it.ref.Append("thing_type_name"))
}

// Version returns a reference to field version of aws_iot_thing.
func (it iotThingAttributes) Version() terra.NumberValue {
	return terra.ReferenceAsNumber(it.ref.Append("version"))
}

type iotThingState struct {
	Arn             string            `json:"arn"`
	Attributes      map[string]string `json:"attributes"`
	DefaultClientId string            `json:"default_client_id"`
	Id              string            `json:"id"`
	Name            string            `json:"name"`
	ThingTypeName   string            `json:"thing_type_name"`
	Version         float64           `json:"version"`
}
