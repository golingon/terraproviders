// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sfnalias "github.com/golingon/terraproviders/aws/5.8.0/sfnalias"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSfnAlias creates a new instance of [SfnAlias].
func NewSfnAlias(name string, args SfnAliasArgs) *SfnAlias {
	return &SfnAlias{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SfnAlias)(nil)

// SfnAlias represents the Terraform resource aws_sfn_alias.
type SfnAlias struct {
	Name      string
	Args      SfnAliasArgs
	state     *sfnAliasState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SfnAlias].
func (sa *SfnAlias) Type() string {
	return "aws_sfn_alias"
}

// LocalName returns the local name for [SfnAlias].
func (sa *SfnAlias) LocalName() string {
	return sa.Name
}

// Configuration returns the configuration (args) for [SfnAlias].
func (sa *SfnAlias) Configuration() interface{} {
	return sa.Args
}

// DependOn is used for other resources to depend on [SfnAlias].
func (sa *SfnAlias) DependOn() terra.Reference {
	return terra.ReferenceResource(sa)
}

// Dependencies returns the list of resources [SfnAlias] depends_on.
func (sa *SfnAlias) Dependencies() terra.Dependencies {
	return sa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SfnAlias].
func (sa *SfnAlias) LifecycleManagement() *terra.Lifecycle {
	return sa.Lifecycle
}

// Attributes returns the attributes for [SfnAlias].
func (sa *SfnAlias) Attributes() sfnAliasAttributes {
	return sfnAliasAttributes{ref: terra.ReferenceResource(sa)}
}

// ImportState imports the given attribute values into [SfnAlias]'s state.
func (sa *SfnAlias) ImportState(av io.Reader) error {
	sa.state = &sfnAliasState{}
	if err := json.NewDecoder(av).Decode(sa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sa.Type(), sa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SfnAlias] has state.
func (sa *SfnAlias) State() (*sfnAliasState, bool) {
	return sa.state, sa.state != nil
}

// StateMust returns the state for [SfnAlias]. Panics if the state is nil.
func (sa *SfnAlias) StateMust() *sfnAliasState {
	if sa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sa.Type(), sa.LocalName()))
	}
	return sa.state
}

// SfnAliasArgs contains the configurations for aws_sfn_alias.
type SfnAliasArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RoutingConfiguration: min=1
	RoutingConfiguration []sfnalias.RoutingConfiguration `hcl:"routing_configuration,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *sfnalias.Timeouts `hcl:"timeouts,block"`
}
type sfnAliasAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sfn_alias.
func (sa sfnAliasAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("arn"))
}

// CreationDate returns a reference to field creation_date of aws_sfn_alias.
func (sa sfnAliasAttributes) CreationDate() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("creation_date"))
}

// Description returns a reference to field description of aws_sfn_alias.
func (sa sfnAliasAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("description"))
}

// Id returns a reference to field id of aws_sfn_alias.
func (sa sfnAliasAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("id"))
}

// Name returns a reference to field name of aws_sfn_alias.
func (sa sfnAliasAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("name"))
}

func (sa sfnAliasAttributes) RoutingConfiguration() terra.ListValue[sfnalias.RoutingConfigurationAttributes] {
	return terra.ReferenceAsList[sfnalias.RoutingConfigurationAttributes](sa.ref.Append("routing_configuration"))
}

func (sa sfnAliasAttributes) Timeouts() sfnalias.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sfnalias.TimeoutsAttributes](sa.ref.Append("timeouts"))
}

type sfnAliasState struct {
	Arn                  string                               `json:"arn"`
	CreationDate         string                               `json:"creation_date"`
	Description          string                               `json:"description"`
	Id                   string                               `json:"id"`
	Name                 string                               `json:"name"`
	RoutingConfiguration []sfnalias.RoutingConfigurationState `json:"routing_configuration"`
	Timeouts             *sfnalias.TimeoutsState              `json:"timeouts"`
}
