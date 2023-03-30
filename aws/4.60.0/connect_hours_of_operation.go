// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	connecthoursofoperation "github.com/golingon/terraproviders/aws/4.60.0/connecthoursofoperation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewConnectHoursOfOperation(name string, args ConnectHoursOfOperationArgs) *ConnectHoursOfOperation {
	return &ConnectHoursOfOperation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConnectHoursOfOperation)(nil)

type ConnectHoursOfOperation struct {
	Name  string
	Args  ConnectHoursOfOperationArgs
	state *connectHoursOfOperationState
}

func (choo *ConnectHoursOfOperation) Type() string {
	return "aws_connect_hours_of_operation"
}

func (choo *ConnectHoursOfOperation) LocalName() string {
	return choo.Name
}

func (choo *ConnectHoursOfOperation) Configuration() interface{} {
	return choo.Args
}

func (choo *ConnectHoursOfOperation) Attributes() connectHoursOfOperationAttributes {
	return connectHoursOfOperationAttributes{ref: terra.ReferenceResource(choo)}
}

func (choo *ConnectHoursOfOperation) ImportState(av io.Reader) error {
	choo.state = &connectHoursOfOperationState{}
	if err := json.NewDecoder(av).Decode(choo.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", choo.Type(), choo.LocalName(), err)
	}
	return nil
}

func (choo *ConnectHoursOfOperation) State() (*connectHoursOfOperationState, bool) {
	return choo.state, choo.state != nil
}

func (choo *ConnectHoursOfOperation) StateMust() *connectHoursOfOperationState {
	if choo.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", choo.Type(), choo.LocalName()))
	}
	return choo.state
}

func (choo *ConnectHoursOfOperation) DependOn() terra.Reference {
	return terra.ReferenceResource(choo)
}

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
	// DependsOn contains resources that ConnectHoursOfOperation depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type connectHoursOfOperationAttributes struct {
	ref terra.Reference
}

func (choo connectHoursOfOperationAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(choo.ref.Append("arn"))
}

func (choo connectHoursOfOperationAttributes) Description() terra.StringValue {
	return terra.ReferenceString(choo.ref.Append("description"))
}

func (choo connectHoursOfOperationAttributes) HoursOfOperationArn() terra.StringValue {
	return terra.ReferenceString(choo.ref.Append("hours_of_operation_arn"))
}

func (choo connectHoursOfOperationAttributes) HoursOfOperationId() terra.StringValue {
	return terra.ReferenceString(choo.ref.Append("hours_of_operation_id"))
}

func (choo connectHoursOfOperationAttributes) Id() terra.StringValue {
	return terra.ReferenceString(choo.ref.Append("id"))
}

func (choo connectHoursOfOperationAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceString(choo.ref.Append("instance_id"))
}

func (choo connectHoursOfOperationAttributes) Name() terra.StringValue {
	return terra.ReferenceString(choo.ref.Append("name"))
}

func (choo connectHoursOfOperationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](choo.ref.Append("tags"))
}

func (choo connectHoursOfOperationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](choo.ref.Append("tags_all"))
}

func (choo connectHoursOfOperationAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceString(choo.ref.Append("time_zone"))
}

func (choo connectHoursOfOperationAttributes) Config() terra.SetValue[connecthoursofoperation.ConfigAttributes] {
	return terra.ReferenceSet[connecthoursofoperation.ConfigAttributes](choo.ref.Append("config"))
}

type connectHoursOfOperationState struct {
	Arn                 string                                `json:"arn"`
	Description         string                                `json:"description"`
	HoursOfOperationArn string                                `json:"hours_of_operation_arn"`
	HoursOfOperationId  string                                `json:"hours_of_operation_id"`
	Id                  string                                `json:"id"`
	InstanceId          string                                `json:"instance_id"`
	Name                string                                `json:"name"`
	Tags                map[string]string                     `json:"tags"`
	TagsAll             map[string]string                     `json:"tags_all"`
	TimeZone            string                                `json:"time_zone"`
	Config              []connecthoursofoperation.ConfigState `json:"config"`
}
