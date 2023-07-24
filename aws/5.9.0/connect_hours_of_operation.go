// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	connecthoursofoperation "github.com/golingon/terraproviders/aws/5.9.0/connecthoursofoperation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewConnectHoursOfOperation creates a new instance of [ConnectHoursOfOperation].
func NewConnectHoursOfOperation(name string, args ConnectHoursOfOperationArgs) *ConnectHoursOfOperation {
	return &ConnectHoursOfOperation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConnectHoursOfOperation)(nil)

// ConnectHoursOfOperation represents the Terraform resource aws_connect_hours_of_operation.
type ConnectHoursOfOperation struct {
	Name      string
	Args      ConnectHoursOfOperationArgs
	state     *connectHoursOfOperationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConnectHoursOfOperation].
func (choo *ConnectHoursOfOperation) Type() string {
	return "aws_connect_hours_of_operation"
}

// LocalName returns the local name for [ConnectHoursOfOperation].
func (choo *ConnectHoursOfOperation) LocalName() string {
	return choo.Name
}

// Configuration returns the configuration (args) for [ConnectHoursOfOperation].
func (choo *ConnectHoursOfOperation) Configuration() interface{} {
	return choo.Args
}

// DependOn is used for other resources to depend on [ConnectHoursOfOperation].
func (choo *ConnectHoursOfOperation) DependOn() terra.Reference {
	return terra.ReferenceResource(choo)
}

// Dependencies returns the list of resources [ConnectHoursOfOperation] depends_on.
func (choo *ConnectHoursOfOperation) Dependencies() terra.Dependencies {
	return choo.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConnectHoursOfOperation].
func (choo *ConnectHoursOfOperation) LifecycleManagement() *terra.Lifecycle {
	return choo.Lifecycle
}

// Attributes returns the attributes for [ConnectHoursOfOperation].
func (choo *ConnectHoursOfOperation) Attributes() connectHoursOfOperationAttributes {
	return connectHoursOfOperationAttributes{ref: terra.ReferenceResource(choo)}
}

// ImportState imports the given attribute values into [ConnectHoursOfOperation]'s state.
func (choo *ConnectHoursOfOperation) ImportState(av io.Reader) error {
	choo.state = &connectHoursOfOperationState{}
	if err := json.NewDecoder(av).Decode(choo.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", choo.Type(), choo.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConnectHoursOfOperation] has state.
func (choo *ConnectHoursOfOperation) State() (*connectHoursOfOperationState, bool) {
	return choo.state, choo.state != nil
}

// StateMust returns the state for [ConnectHoursOfOperation]. Panics if the state is nil.
func (choo *ConnectHoursOfOperation) StateMust() *connectHoursOfOperationState {
	if choo.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", choo.Type(), choo.LocalName()))
	}
	return choo.state
}

// ConnectHoursOfOperationArgs contains the configurations for aws_connect_hours_of_operation.
type ConnectHoursOfOperationArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TimeZone: string, required
	TimeZone terra.StringValue `hcl:"time_zone,attr" validate:"required"`
	// Config: min=1
	Config []connecthoursofoperation.Config `hcl:"config,block" validate:"min=1"`
}
type connectHoursOfOperationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_connect_hours_of_operation.
func (choo connectHoursOfOperationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(choo.ref.Append("arn"))
}

// Description returns a reference to field description of aws_connect_hours_of_operation.
func (choo connectHoursOfOperationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(choo.ref.Append("description"))
}

// HoursOfOperationId returns a reference to field hours_of_operation_id of aws_connect_hours_of_operation.
func (choo connectHoursOfOperationAttributes) HoursOfOperationId() terra.StringValue {
	return terra.ReferenceAsString(choo.ref.Append("hours_of_operation_id"))
}

// Id returns a reference to field id of aws_connect_hours_of_operation.
func (choo connectHoursOfOperationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(choo.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_connect_hours_of_operation.
func (choo connectHoursOfOperationAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(choo.ref.Append("instance_id"))
}

// Name returns a reference to field name of aws_connect_hours_of_operation.
func (choo connectHoursOfOperationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(choo.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_connect_hours_of_operation.
func (choo connectHoursOfOperationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](choo.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_connect_hours_of_operation.
func (choo connectHoursOfOperationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](choo.ref.Append("tags_all"))
}

// TimeZone returns a reference to field time_zone of aws_connect_hours_of_operation.
func (choo connectHoursOfOperationAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(choo.ref.Append("time_zone"))
}

func (choo connectHoursOfOperationAttributes) Config() terra.SetValue[connecthoursofoperation.ConfigAttributes] {
	return terra.ReferenceAsSet[connecthoursofoperation.ConfigAttributes](choo.ref.Append("config"))
}

type connectHoursOfOperationState struct {
	Arn                string                                `json:"arn"`
	Description        string                                `json:"description"`
	HoursOfOperationId string                                `json:"hours_of_operation_id"`
	Id                 string                                `json:"id"`
	InstanceId         string                                `json:"instance_id"`
	Name               string                                `json:"name"`
	Tags               map[string]string                     `json:"tags"`
	TagsAll            map[string]string                     `json:"tags_all"`
	TimeZone           string                                `json:"time_zone"`
	Config             []connecthoursofoperation.ConfigState `json:"config"`
}
