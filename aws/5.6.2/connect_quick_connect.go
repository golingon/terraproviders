// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	connectquickconnect "github.com/golingon/terraproviders/aws/5.6.2/connectquickconnect"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewConnectQuickConnect creates a new instance of [ConnectQuickConnect].
func NewConnectQuickConnect(name string, args ConnectQuickConnectArgs) *ConnectQuickConnect {
	return &ConnectQuickConnect{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConnectQuickConnect)(nil)

// ConnectQuickConnect represents the Terraform resource aws_connect_quick_connect.
type ConnectQuickConnect struct {
	Name      string
	Args      ConnectQuickConnectArgs
	state     *connectQuickConnectState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConnectQuickConnect].
func (cqc *ConnectQuickConnect) Type() string {
	return "aws_connect_quick_connect"
}

// LocalName returns the local name for [ConnectQuickConnect].
func (cqc *ConnectQuickConnect) LocalName() string {
	return cqc.Name
}

// Configuration returns the configuration (args) for [ConnectQuickConnect].
func (cqc *ConnectQuickConnect) Configuration() interface{} {
	return cqc.Args
}

// DependOn is used for other resources to depend on [ConnectQuickConnect].
func (cqc *ConnectQuickConnect) DependOn() terra.Reference {
	return terra.ReferenceResource(cqc)
}

// Dependencies returns the list of resources [ConnectQuickConnect] depends_on.
func (cqc *ConnectQuickConnect) Dependencies() terra.Dependencies {
	return cqc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConnectQuickConnect].
func (cqc *ConnectQuickConnect) LifecycleManagement() *terra.Lifecycle {
	return cqc.Lifecycle
}

// Attributes returns the attributes for [ConnectQuickConnect].
func (cqc *ConnectQuickConnect) Attributes() connectQuickConnectAttributes {
	return connectQuickConnectAttributes{ref: terra.ReferenceResource(cqc)}
}

// ImportState imports the given attribute values into [ConnectQuickConnect]'s state.
func (cqc *ConnectQuickConnect) ImportState(av io.Reader) error {
	cqc.state = &connectQuickConnectState{}
	if err := json.NewDecoder(av).Decode(cqc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cqc.Type(), cqc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConnectQuickConnect] has state.
func (cqc *ConnectQuickConnect) State() (*connectQuickConnectState, bool) {
	return cqc.state, cqc.state != nil
}

// StateMust returns the state for [ConnectQuickConnect]. Panics if the state is nil.
func (cqc *ConnectQuickConnect) StateMust() *connectQuickConnectState {
	if cqc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cqc.Type(), cqc.LocalName()))
	}
	return cqc.state
}

// ConnectQuickConnectArgs contains the configurations for aws_connect_quick_connect.
type ConnectQuickConnectArgs struct {
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
	// QuickConnectConfig: required
	QuickConnectConfig *connectquickconnect.QuickConnectConfig `hcl:"quick_connect_config,block" validate:"required"`
}
type connectQuickConnectAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_connect_quick_connect.
func (cqc connectQuickConnectAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cqc.ref.Append("arn"))
}

// Description returns a reference to field description of aws_connect_quick_connect.
func (cqc connectQuickConnectAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cqc.ref.Append("description"))
}

// Id returns a reference to field id of aws_connect_quick_connect.
func (cqc connectQuickConnectAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cqc.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_connect_quick_connect.
func (cqc connectQuickConnectAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(cqc.ref.Append("instance_id"))
}

// Name returns a reference to field name of aws_connect_quick_connect.
func (cqc connectQuickConnectAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cqc.ref.Append("name"))
}

// QuickConnectId returns a reference to field quick_connect_id of aws_connect_quick_connect.
func (cqc connectQuickConnectAttributes) QuickConnectId() terra.StringValue {
	return terra.ReferenceAsString(cqc.ref.Append("quick_connect_id"))
}

// Tags returns a reference to field tags of aws_connect_quick_connect.
func (cqc connectQuickConnectAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cqc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_connect_quick_connect.
func (cqc connectQuickConnectAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cqc.ref.Append("tags_all"))
}

func (cqc connectQuickConnectAttributes) QuickConnectConfig() terra.ListValue[connectquickconnect.QuickConnectConfigAttributes] {
	return terra.ReferenceAsList[connectquickconnect.QuickConnectConfigAttributes](cqc.ref.Append("quick_connect_config"))
}

type connectQuickConnectState struct {
	Arn                string                                        `json:"arn"`
	Description        string                                        `json:"description"`
	Id                 string                                        `json:"id"`
	InstanceId         string                                        `json:"instance_id"`
	Name               string                                        `json:"name"`
	QuickConnectId     string                                        `json:"quick_connect_id"`
	Tags               map[string]string                             `json:"tags"`
	TagsAll            map[string]string                             `json:"tags_all"`
	QuickConnectConfig []connectquickconnect.QuickConnectConfigState `json:"quick_connect_config"`
}
