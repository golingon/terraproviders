// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_db_proxy

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

// Resource represents the Terraform resource aws_db_proxy.
type Resource struct {
	Name      string
	Args      Args
	state     *awsDbProxyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adp *Resource) Type() string {
	return "aws_db_proxy"
}

// LocalName returns the local name for [Resource].
func (adp *Resource) LocalName() string {
	return adp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adp *Resource) Configuration() interface{} {
	return adp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adp *Resource) Dependencies() terra.Dependencies {
	return adp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adp *Resource) LifecycleManagement() *terra.Lifecycle {
	return adp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adp *Resource) Attributes() awsDbProxyAttributes {
	return awsDbProxyAttributes{ref: terra.ReferenceResource(adp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adp *Resource) ImportState(state io.Reader) error {
	adp.state = &awsDbProxyState{}
	if err := json.NewDecoder(state).Decode(adp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adp.Type(), adp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adp *Resource) State() (*awsDbProxyState, bool) {
	return adp.state, adp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adp *Resource) StateMust() *awsDbProxyState {
	if adp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adp.Type(), adp.LocalName()))
	}
	return adp.state
}

// Args contains the configurations for aws_db_proxy.
type Args struct {
	// DebugLogging: bool, optional
	DebugLogging terra.BoolValue `hcl:"debug_logging,attr"`
	// EngineFamily: string, required
	EngineFamily terra.StringValue `hcl:"engine_family,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdleClientTimeout: number, optional
	IdleClientTimeout terra.NumberValue `hcl:"idle_client_timeout,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RequireTls: bool, optional
	RequireTls terra.BoolValue `hcl:"require_tls,attr"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VpcSecurityGroupIds: set of string, optional
	VpcSecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"vpc_security_group_ids,attr"`
	// VpcSubnetIds: set of string, required
	VpcSubnetIds terra.SetValue[terra.StringValue] `hcl:"vpc_subnet_ids,attr" validate:"required"`
	// Auth: min=1
	Auth []Auth `hcl:"auth,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsDbProxyAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_db_proxy.
func (adp awsDbProxyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(adp.ref.Append("arn"))
}

// DebugLogging returns a reference to field debug_logging of aws_db_proxy.
func (adp awsDbProxyAttributes) DebugLogging() terra.BoolValue {
	return terra.ReferenceAsBool(adp.ref.Append("debug_logging"))
}

// Endpoint returns a reference to field endpoint of aws_db_proxy.
func (adp awsDbProxyAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(adp.ref.Append("endpoint"))
}

// EngineFamily returns a reference to field engine_family of aws_db_proxy.
func (adp awsDbProxyAttributes) EngineFamily() terra.StringValue {
	return terra.ReferenceAsString(adp.ref.Append("engine_family"))
}

// Id returns a reference to field id of aws_db_proxy.
func (adp awsDbProxyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adp.ref.Append("id"))
}

// IdleClientTimeout returns a reference to field idle_client_timeout of aws_db_proxy.
func (adp awsDbProxyAttributes) IdleClientTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(adp.ref.Append("idle_client_timeout"))
}

// Name returns a reference to field name of aws_db_proxy.
func (adp awsDbProxyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adp.ref.Append("name"))
}

// RequireTls returns a reference to field require_tls of aws_db_proxy.
func (adp awsDbProxyAttributes) RequireTls() terra.BoolValue {
	return terra.ReferenceAsBool(adp.ref.Append("require_tls"))
}

// RoleArn returns a reference to field role_arn of aws_db_proxy.
func (adp awsDbProxyAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(adp.ref.Append("role_arn"))
}

// Tags returns a reference to field tags of aws_db_proxy.
func (adp awsDbProxyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_db_proxy.
func (adp awsDbProxyAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adp.ref.Append("tags_all"))
}

// VpcSecurityGroupIds returns a reference to field vpc_security_group_ids of aws_db_proxy.
func (adp awsDbProxyAttributes) VpcSecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](adp.ref.Append("vpc_security_group_ids"))
}

// VpcSubnetIds returns a reference to field vpc_subnet_ids of aws_db_proxy.
func (adp awsDbProxyAttributes) VpcSubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](adp.ref.Append("vpc_subnet_ids"))
}

func (adp awsDbProxyAttributes) Auth() terra.SetValue[AuthAttributes] {
	return terra.ReferenceAsSet[AuthAttributes](adp.ref.Append("auth"))
}

func (adp awsDbProxyAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](adp.ref.Append("timeouts"))
}

type awsDbProxyState struct {
	Arn                 string            `json:"arn"`
	DebugLogging        bool              `json:"debug_logging"`
	Endpoint            string            `json:"endpoint"`
	EngineFamily        string            `json:"engine_family"`
	Id                  string            `json:"id"`
	IdleClientTimeout   float64           `json:"idle_client_timeout"`
	Name                string            `json:"name"`
	RequireTls          bool              `json:"require_tls"`
	RoleArn             string            `json:"role_arn"`
	Tags                map[string]string `json:"tags"`
	TagsAll             map[string]string `json:"tags_all"`
	VpcSecurityGroupIds []string          `json:"vpc_security_group_ids"`
	VpcSubnetIds        []string          `json:"vpc_subnet_ids"`
	Auth                []AuthState       `json:"auth"`
	Timeouts            *TimeoutsState    `json:"timeouts"`
}
