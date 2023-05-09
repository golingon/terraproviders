// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	vpcipampool "github.com/golingon/terraproviders/aws/4.66.1/vpcipampool"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpcIpamPool creates a new instance of [VpcIpamPool].
func NewVpcIpamPool(name string, args VpcIpamPoolArgs) *VpcIpamPool {
	return &VpcIpamPool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpcIpamPool)(nil)

// VpcIpamPool represents the Terraform resource aws_vpc_ipam_pool.
type VpcIpamPool struct {
	Name      string
	Args      VpcIpamPoolArgs
	state     *vpcIpamPoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpcIpamPool].
func (vip *VpcIpamPool) Type() string {
	return "aws_vpc_ipam_pool"
}

// LocalName returns the local name for [VpcIpamPool].
func (vip *VpcIpamPool) LocalName() string {
	return vip.Name
}

// Configuration returns the configuration (args) for [VpcIpamPool].
func (vip *VpcIpamPool) Configuration() interface{} {
	return vip.Args
}

// DependOn is used for other resources to depend on [VpcIpamPool].
func (vip *VpcIpamPool) DependOn() terra.Reference {
	return terra.ReferenceResource(vip)
}

// Dependencies returns the list of resources [VpcIpamPool] depends_on.
func (vip *VpcIpamPool) Dependencies() terra.Dependencies {
	return vip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpcIpamPool].
func (vip *VpcIpamPool) LifecycleManagement() *terra.Lifecycle {
	return vip.Lifecycle
}

// Attributes returns the attributes for [VpcIpamPool].
func (vip *VpcIpamPool) Attributes() vpcIpamPoolAttributes {
	return vpcIpamPoolAttributes{ref: terra.ReferenceResource(vip)}
}

// ImportState imports the given attribute values into [VpcIpamPool]'s state.
func (vip *VpcIpamPool) ImportState(av io.Reader) error {
	vip.state = &vpcIpamPoolState{}
	if err := json.NewDecoder(av).Decode(vip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vip.Type(), vip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpcIpamPool] has state.
func (vip *VpcIpamPool) State() (*vpcIpamPoolState, bool) {
	return vip.state, vip.state != nil
}

// StateMust returns the state for [VpcIpamPool]. Panics if the state is nil.
func (vip *VpcIpamPool) StateMust() *vpcIpamPoolState {
	if vip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vip.Type(), vip.LocalName()))
	}
	return vip.state
}

// VpcIpamPoolArgs contains the configurations for aws_vpc_ipam_pool.
type VpcIpamPoolArgs struct {
	// AddressFamily: string, required
	AddressFamily terra.StringValue `hcl:"address_family,attr" validate:"required"`
	// AllocationDefaultNetmaskLength: number, optional
	AllocationDefaultNetmaskLength terra.NumberValue `hcl:"allocation_default_netmask_length,attr"`
	// AllocationMaxNetmaskLength: number, optional
	AllocationMaxNetmaskLength terra.NumberValue `hcl:"allocation_max_netmask_length,attr"`
	// AllocationMinNetmaskLength: number, optional
	AllocationMinNetmaskLength terra.NumberValue `hcl:"allocation_min_netmask_length,attr"`
	// AllocationResourceTags: map of string, optional
	AllocationResourceTags terra.MapValue[terra.StringValue] `hcl:"allocation_resource_tags,attr"`
	// AutoImport: bool, optional
	AutoImport terra.BoolValue `hcl:"auto_import,attr"`
	// AwsService: string, optional
	AwsService terra.StringValue `hcl:"aws_service,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpamScopeId: string, required
	IpamScopeId terra.StringValue `hcl:"ipam_scope_id,attr" validate:"required"`
	// Locale: string, optional
	Locale terra.StringValue `hcl:"locale,attr"`
	// PublicIpSource: string, optional
	PublicIpSource terra.StringValue `hcl:"public_ip_source,attr"`
	// PubliclyAdvertisable: bool, optional
	PubliclyAdvertisable terra.BoolValue `hcl:"publicly_advertisable,attr"`
	// SourceIpamPoolId: string, optional
	SourceIpamPoolId terra.StringValue `hcl:"source_ipam_pool_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *vpcipampool.Timeouts `hcl:"timeouts,block"`
}
type vpcIpamPoolAttributes struct {
	ref terra.Reference
}

// AddressFamily returns a reference to field address_family of aws_vpc_ipam_pool.
func (vip vpcIpamPoolAttributes) AddressFamily() terra.StringValue {
	return terra.ReferenceAsString(vip.ref.Append("address_family"))
}

// AllocationDefaultNetmaskLength returns a reference to field allocation_default_netmask_length of aws_vpc_ipam_pool.
func (vip vpcIpamPoolAttributes) AllocationDefaultNetmaskLength() terra.NumberValue {
	return terra.ReferenceAsNumber(vip.ref.Append("allocation_default_netmask_length"))
}

// AllocationMaxNetmaskLength returns a reference to field allocation_max_netmask_length of aws_vpc_ipam_pool.
func (vip vpcIpamPoolAttributes) AllocationMaxNetmaskLength() terra.NumberValue {
	return terra.ReferenceAsNumber(vip.ref.Append("allocation_max_netmask_length"))
}

// AllocationMinNetmaskLength returns a reference to field allocation_min_netmask_length of aws_vpc_ipam_pool.
func (vip vpcIpamPoolAttributes) AllocationMinNetmaskLength() terra.NumberValue {
	return terra.ReferenceAsNumber(vip.ref.Append("allocation_min_netmask_length"))
}

// AllocationResourceTags returns a reference to field allocation_resource_tags of aws_vpc_ipam_pool.
func (vip vpcIpamPoolAttributes) AllocationResourceTags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vip.ref.Append("allocation_resource_tags"))
}

// Arn returns a reference to field arn of aws_vpc_ipam_pool.
func (vip vpcIpamPoolAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(vip.ref.Append("arn"))
}

// AutoImport returns a reference to field auto_import of aws_vpc_ipam_pool.
func (vip vpcIpamPoolAttributes) AutoImport() terra.BoolValue {
	return terra.ReferenceAsBool(vip.ref.Append("auto_import"))
}

// AwsService returns a reference to field aws_service of aws_vpc_ipam_pool.
func (vip vpcIpamPoolAttributes) AwsService() terra.StringValue {
	return terra.ReferenceAsString(vip.ref.Append("aws_service"))
}

// Description returns a reference to field description of aws_vpc_ipam_pool.
func (vip vpcIpamPoolAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(vip.ref.Append("description"))
}

// Id returns a reference to field id of aws_vpc_ipam_pool.
func (vip vpcIpamPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vip.ref.Append("id"))
}

// IpamScopeId returns a reference to field ipam_scope_id of aws_vpc_ipam_pool.
func (vip vpcIpamPoolAttributes) IpamScopeId() terra.StringValue {
	return terra.ReferenceAsString(vip.ref.Append("ipam_scope_id"))
}

// IpamScopeType returns a reference to field ipam_scope_type of aws_vpc_ipam_pool.
func (vip vpcIpamPoolAttributes) IpamScopeType() terra.StringValue {
	return terra.ReferenceAsString(vip.ref.Append("ipam_scope_type"))
}

// Locale returns a reference to field locale of aws_vpc_ipam_pool.
func (vip vpcIpamPoolAttributes) Locale() terra.StringValue {
	return terra.ReferenceAsString(vip.ref.Append("locale"))
}

// PoolDepth returns a reference to field pool_depth of aws_vpc_ipam_pool.
func (vip vpcIpamPoolAttributes) PoolDepth() terra.NumberValue {
	return terra.ReferenceAsNumber(vip.ref.Append("pool_depth"))
}

// PublicIpSource returns a reference to field public_ip_source of aws_vpc_ipam_pool.
func (vip vpcIpamPoolAttributes) PublicIpSource() terra.StringValue {
	return terra.ReferenceAsString(vip.ref.Append("public_ip_source"))
}

// PubliclyAdvertisable returns a reference to field publicly_advertisable of aws_vpc_ipam_pool.
func (vip vpcIpamPoolAttributes) PubliclyAdvertisable() terra.BoolValue {
	return terra.ReferenceAsBool(vip.ref.Append("publicly_advertisable"))
}

// SourceIpamPoolId returns a reference to field source_ipam_pool_id of aws_vpc_ipam_pool.
func (vip vpcIpamPoolAttributes) SourceIpamPoolId() terra.StringValue {
	return terra.ReferenceAsString(vip.ref.Append("source_ipam_pool_id"))
}

// State returns a reference to field state of aws_vpc_ipam_pool.
func (vip vpcIpamPoolAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(vip.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_vpc_ipam_pool.
func (vip vpcIpamPoolAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vip.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_vpc_ipam_pool.
func (vip vpcIpamPoolAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vip.ref.Append("tags_all"))
}

func (vip vpcIpamPoolAttributes) Timeouts() vpcipampool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpcipampool.TimeoutsAttributes](vip.ref.Append("timeouts"))
}

type vpcIpamPoolState struct {
	AddressFamily                  string                     `json:"address_family"`
	AllocationDefaultNetmaskLength float64                    `json:"allocation_default_netmask_length"`
	AllocationMaxNetmaskLength     float64                    `json:"allocation_max_netmask_length"`
	AllocationMinNetmaskLength     float64                    `json:"allocation_min_netmask_length"`
	AllocationResourceTags         map[string]string          `json:"allocation_resource_tags"`
	Arn                            string                     `json:"arn"`
	AutoImport                     bool                       `json:"auto_import"`
	AwsService                     string                     `json:"aws_service"`
	Description                    string                     `json:"description"`
	Id                             string                     `json:"id"`
	IpamScopeId                    string                     `json:"ipam_scope_id"`
	IpamScopeType                  string                     `json:"ipam_scope_type"`
	Locale                         string                     `json:"locale"`
	PoolDepth                      float64                    `json:"pool_depth"`
	PublicIpSource                 string                     `json:"public_ip_source"`
	PubliclyAdvertisable           bool                       `json:"publicly_advertisable"`
	SourceIpamPoolId               string                     `json:"source_ipam_pool_id"`
	State                          string                     `json:"state"`
	Tags                           map[string]string          `json:"tags"`
	TagsAll                        map[string]string          `json:"tags_all"`
	Timeouts                       *vpcipampool.TimeoutsState `json:"timeouts"`
}
