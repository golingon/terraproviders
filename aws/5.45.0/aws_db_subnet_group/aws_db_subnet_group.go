// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_db_subnet_group

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

// Resource represents the Terraform resource aws_db_subnet_group.
type Resource struct {
	Name      string
	Args      Args
	state     *awsDbSubnetGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adsg *Resource) Type() string {
	return "aws_db_subnet_group"
}

// LocalName returns the local name for [Resource].
func (adsg *Resource) LocalName() string {
	return adsg.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adsg *Resource) Configuration() interface{} {
	return adsg.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adsg *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adsg)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adsg *Resource) Dependencies() terra.Dependencies {
	return adsg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adsg *Resource) LifecycleManagement() *terra.Lifecycle {
	return adsg.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adsg *Resource) Attributes() awsDbSubnetGroupAttributes {
	return awsDbSubnetGroupAttributes{ref: terra.ReferenceResource(adsg)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adsg *Resource) ImportState(state io.Reader) error {
	adsg.state = &awsDbSubnetGroupState{}
	if err := json.NewDecoder(state).Decode(adsg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adsg.Type(), adsg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adsg *Resource) State() (*awsDbSubnetGroupState, bool) {
	return adsg.state, adsg.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adsg *Resource) StateMust() *awsDbSubnetGroupState {
	if adsg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adsg.Type(), adsg.LocalName()))
	}
	return adsg.state
}

// Args contains the configurations for aws_db_subnet_group.
type Args struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// SubnetIds: set of string, required
	SubnetIds terra.SetValue[terra.StringValue] `hcl:"subnet_ids,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}

type awsDbSubnetGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_db_subnet_group.
func (adsg awsDbSubnetGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(adsg.ref.Append("arn"))
}

// Description returns a reference to field description of aws_db_subnet_group.
func (adsg awsDbSubnetGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(adsg.ref.Append("description"))
}

// Id returns a reference to field id of aws_db_subnet_group.
func (adsg awsDbSubnetGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adsg.ref.Append("id"))
}

// Name returns a reference to field name of aws_db_subnet_group.
func (adsg awsDbSubnetGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adsg.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_db_subnet_group.
func (adsg awsDbSubnetGroupAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(adsg.ref.Append("name_prefix"))
}

// SubnetIds returns a reference to field subnet_ids of aws_db_subnet_group.
func (adsg awsDbSubnetGroupAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](adsg.ref.Append("subnet_ids"))
}

// SupportedNetworkTypes returns a reference to field supported_network_types of aws_db_subnet_group.
func (adsg awsDbSubnetGroupAttributes) SupportedNetworkTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](adsg.ref.Append("supported_network_types"))
}

// Tags returns a reference to field tags of aws_db_subnet_group.
func (adsg awsDbSubnetGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adsg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_db_subnet_group.
func (adsg awsDbSubnetGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adsg.ref.Append("tags_all"))
}

// VpcId returns a reference to field vpc_id of aws_db_subnet_group.
func (adsg awsDbSubnetGroupAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(adsg.ref.Append("vpc_id"))
}

type awsDbSubnetGroupState struct {
	Arn                   string            `json:"arn"`
	Description           string            `json:"description"`
	Id                    string            `json:"id"`
	Name                  string            `json:"name"`
	NamePrefix            string            `json:"name_prefix"`
	SubnetIds             []string          `json:"subnet_ids"`
	SupportedNetworkTypes []string          `json:"supported_network_types"`
	Tags                  map[string]string `json:"tags"`
	TagsAll               map[string]string `json:"tags_all"`
	VpcId                 string            `json:"vpc_id"`
}
