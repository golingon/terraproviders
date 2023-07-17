// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	vpcipamscope "github.com/golingon/terraproviders/aws/5.8.0/vpcipamscope"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpcIpamScope creates a new instance of [VpcIpamScope].
func NewVpcIpamScope(name string, args VpcIpamScopeArgs) *VpcIpamScope {
	return &VpcIpamScope{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpcIpamScope)(nil)

// VpcIpamScope represents the Terraform resource aws_vpc_ipam_scope.
type VpcIpamScope struct {
	Name      string
	Args      VpcIpamScopeArgs
	state     *vpcIpamScopeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpcIpamScope].
func (vis *VpcIpamScope) Type() string {
	return "aws_vpc_ipam_scope"
}

// LocalName returns the local name for [VpcIpamScope].
func (vis *VpcIpamScope) LocalName() string {
	return vis.Name
}

// Configuration returns the configuration (args) for [VpcIpamScope].
func (vis *VpcIpamScope) Configuration() interface{} {
	return vis.Args
}

// DependOn is used for other resources to depend on [VpcIpamScope].
func (vis *VpcIpamScope) DependOn() terra.Reference {
	return terra.ReferenceResource(vis)
}

// Dependencies returns the list of resources [VpcIpamScope] depends_on.
func (vis *VpcIpamScope) Dependencies() terra.Dependencies {
	return vis.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpcIpamScope].
func (vis *VpcIpamScope) LifecycleManagement() *terra.Lifecycle {
	return vis.Lifecycle
}

// Attributes returns the attributes for [VpcIpamScope].
func (vis *VpcIpamScope) Attributes() vpcIpamScopeAttributes {
	return vpcIpamScopeAttributes{ref: terra.ReferenceResource(vis)}
}

// ImportState imports the given attribute values into [VpcIpamScope]'s state.
func (vis *VpcIpamScope) ImportState(av io.Reader) error {
	vis.state = &vpcIpamScopeState{}
	if err := json.NewDecoder(av).Decode(vis.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vis.Type(), vis.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpcIpamScope] has state.
func (vis *VpcIpamScope) State() (*vpcIpamScopeState, bool) {
	return vis.state, vis.state != nil
}

// StateMust returns the state for [VpcIpamScope]. Panics if the state is nil.
func (vis *VpcIpamScope) StateMust() *vpcIpamScopeState {
	if vis.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vis.Type(), vis.LocalName()))
	}
	return vis.state
}

// VpcIpamScopeArgs contains the configurations for aws_vpc_ipam_scope.
type VpcIpamScopeArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpamId: string, required
	IpamId terra.StringValue `hcl:"ipam_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *vpcipamscope.Timeouts `hcl:"timeouts,block"`
}
type vpcIpamScopeAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_vpc_ipam_scope.
func (vis vpcIpamScopeAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(vis.ref.Append("arn"))
}

// Description returns a reference to field description of aws_vpc_ipam_scope.
func (vis vpcIpamScopeAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(vis.ref.Append("description"))
}

// Id returns a reference to field id of aws_vpc_ipam_scope.
func (vis vpcIpamScopeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vis.ref.Append("id"))
}

// IpamArn returns a reference to field ipam_arn of aws_vpc_ipam_scope.
func (vis vpcIpamScopeAttributes) IpamArn() terra.StringValue {
	return terra.ReferenceAsString(vis.ref.Append("ipam_arn"))
}

// IpamId returns a reference to field ipam_id of aws_vpc_ipam_scope.
func (vis vpcIpamScopeAttributes) IpamId() terra.StringValue {
	return terra.ReferenceAsString(vis.ref.Append("ipam_id"))
}

// IpamScopeType returns a reference to field ipam_scope_type of aws_vpc_ipam_scope.
func (vis vpcIpamScopeAttributes) IpamScopeType() terra.StringValue {
	return terra.ReferenceAsString(vis.ref.Append("ipam_scope_type"))
}

// IsDefault returns a reference to field is_default of aws_vpc_ipam_scope.
func (vis vpcIpamScopeAttributes) IsDefault() terra.BoolValue {
	return terra.ReferenceAsBool(vis.ref.Append("is_default"))
}

// PoolCount returns a reference to field pool_count of aws_vpc_ipam_scope.
func (vis vpcIpamScopeAttributes) PoolCount() terra.NumberValue {
	return terra.ReferenceAsNumber(vis.ref.Append("pool_count"))
}

// Tags returns a reference to field tags of aws_vpc_ipam_scope.
func (vis vpcIpamScopeAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vis.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_vpc_ipam_scope.
func (vis vpcIpamScopeAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vis.ref.Append("tags_all"))
}

func (vis vpcIpamScopeAttributes) Timeouts() vpcipamscope.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpcipamscope.TimeoutsAttributes](vis.ref.Append("timeouts"))
}

type vpcIpamScopeState struct {
	Arn           string                      `json:"arn"`
	Description   string                      `json:"description"`
	Id            string                      `json:"id"`
	IpamArn       string                      `json:"ipam_arn"`
	IpamId        string                      `json:"ipam_id"`
	IpamScopeType string                      `json:"ipam_scope_type"`
	IsDefault     bool                        `json:"is_default"`
	PoolCount     float64                     `json:"pool_count"`
	Tags          map[string]string           `json:"tags"`
	TagsAll       map[string]string           `json:"tags_all"`
	Timeouts      *vpcipamscope.TimeoutsState `json:"timeouts"`
}
