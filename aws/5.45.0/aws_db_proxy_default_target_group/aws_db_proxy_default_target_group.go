// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_db_proxy_default_target_group

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_db_proxy_default_target_group.
type Resource struct {
	Name      string
	Args      Args
	state     *awsDbProxyDefaultTargetGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adpdtg *Resource) Type() string {
	return "aws_db_proxy_default_target_group"
}

// LocalName returns the local name for [Resource].
func (adpdtg *Resource) LocalName() string {
	return adpdtg.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adpdtg *Resource) Configuration() interface{} {
	return adpdtg.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adpdtg *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adpdtg)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adpdtg *Resource) Dependencies() terra.Dependencies {
	return adpdtg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adpdtg *Resource) LifecycleManagement() *terra.Lifecycle {
	return adpdtg.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adpdtg *Resource) Attributes() awsDbProxyDefaultTargetGroupAttributes {
	return awsDbProxyDefaultTargetGroupAttributes{ref: terra.ReferenceResource(adpdtg)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adpdtg *Resource) ImportState(state io.Reader) error {
	adpdtg.state = &awsDbProxyDefaultTargetGroupState{}
	if err := json.NewDecoder(state).Decode(adpdtg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adpdtg.Type(), adpdtg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adpdtg *Resource) State() (*awsDbProxyDefaultTargetGroupState, bool) {
	return adpdtg.state, adpdtg.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adpdtg *Resource) StateMust() *awsDbProxyDefaultTargetGroupState {
	if adpdtg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adpdtg.Type(), adpdtg.LocalName()))
	}
	return adpdtg.state
}

// Args contains the configurations for aws_db_proxy_default_target_group.
type Args struct {
	// DbProxyName: string, required
	DbProxyName terra.StringValue `hcl:"db_proxy_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ConnectionPoolConfig: optional
	ConnectionPoolConfig *ConnectionPoolConfig `hcl:"connection_pool_config,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsDbProxyDefaultTargetGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_db_proxy_default_target_group.
func (adpdtg awsDbProxyDefaultTargetGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(adpdtg.ref.Append("arn"))
}

// DbProxyName returns a reference to field db_proxy_name of aws_db_proxy_default_target_group.
func (adpdtg awsDbProxyDefaultTargetGroupAttributes) DbProxyName() terra.StringValue {
	return terra.ReferenceAsString(adpdtg.ref.Append("db_proxy_name"))
}

// Id returns a reference to field id of aws_db_proxy_default_target_group.
func (adpdtg awsDbProxyDefaultTargetGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adpdtg.ref.Append("id"))
}

// Name returns a reference to field name of aws_db_proxy_default_target_group.
func (adpdtg awsDbProxyDefaultTargetGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adpdtg.ref.Append("name"))
}

func (adpdtg awsDbProxyDefaultTargetGroupAttributes) ConnectionPoolConfig() terra.ListValue[ConnectionPoolConfigAttributes] {
	return terra.ReferenceAsList[ConnectionPoolConfigAttributes](adpdtg.ref.Append("connection_pool_config"))
}

func (adpdtg awsDbProxyDefaultTargetGroupAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](adpdtg.ref.Append("timeouts"))
}

type awsDbProxyDefaultTargetGroupState struct {
	Arn                  string                      `json:"arn"`
	DbProxyName          string                      `json:"db_proxy_name"`
	Id                   string                      `json:"id"`
	Name                 string                      `json:"name"`
	ConnectionPoolConfig []ConnectionPoolConfigState `json:"connection_pool_config"`
	Timeouts             *TimeoutsState              `json:"timeouts"`
}
