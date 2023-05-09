// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	emrblockpublicaccessconfiguration "github.com/golingon/terraproviders/aws/4.66.1/emrblockpublicaccessconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEmrBlockPublicAccessConfiguration creates a new instance of [EmrBlockPublicAccessConfiguration].
func NewEmrBlockPublicAccessConfiguration(name string, args EmrBlockPublicAccessConfigurationArgs) *EmrBlockPublicAccessConfiguration {
	return &EmrBlockPublicAccessConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EmrBlockPublicAccessConfiguration)(nil)

// EmrBlockPublicAccessConfiguration represents the Terraform resource aws_emr_block_public_access_configuration.
type EmrBlockPublicAccessConfiguration struct {
	Name      string
	Args      EmrBlockPublicAccessConfigurationArgs
	state     *emrBlockPublicAccessConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EmrBlockPublicAccessConfiguration].
func (ebpac *EmrBlockPublicAccessConfiguration) Type() string {
	return "aws_emr_block_public_access_configuration"
}

// LocalName returns the local name for [EmrBlockPublicAccessConfiguration].
func (ebpac *EmrBlockPublicAccessConfiguration) LocalName() string {
	return ebpac.Name
}

// Configuration returns the configuration (args) for [EmrBlockPublicAccessConfiguration].
func (ebpac *EmrBlockPublicAccessConfiguration) Configuration() interface{} {
	return ebpac.Args
}

// DependOn is used for other resources to depend on [EmrBlockPublicAccessConfiguration].
func (ebpac *EmrBlockPublicAccessConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(ebpac)
}

// Dependencies returns the list of resources [EmrBlockPublicAccessConfiguration] depends_on.
func (ebpac *EmrBlockPublicAccessConfiguration) Dependencies() terra.Dependencies {
	return ebpac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EmrBlockPublicAccessConfiguration].
func (ebpac *EmrBlockPublicAccessConfiguration) LifecycleManagement() *terra.Lifecycle {
	return ebpac.Lifecycle
}

// Attributes returns the attributes for [EmrBlockPublicAccessConfiguration].
func (ebpac *EmrBlockPublicAccessConfiguration) Attributes() emrBlockPublicAccessConfigurationAttributes {
	return emrBlockPublicAccessConfigurationAttributes{ref: terra.ReferenceResource(ebpac)}
}

// ImportState imports the given attribute values into [EmrBlockPublicAccessConfiguration]'s state.
func (ebpac *EmrBlockPublicAccessConfiguration) ImportState(av io.Reader) error {
	ebpac.state = &emrBlockPublicAccessConfigurationState{}
	if err := json.NewDecoder(av).Decode(ebpac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ebpac.Type(), ebpac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EmrBlockPublicAccessConfiguration] has state.
func (ebpac *EmrBlockPublicAccessConfiguration) State() (*emrBlockPublicAccessConfigurationState, bool) {
	return ebpac.state, ebpac.state != nil
}

// StateMust returns the state for [EmrBlockPublicAccessConfiguration]. Panics if the state is nil.
func (ebpac *EmrBlockPublicAccessConfiguration) StateMust() *emrBlockPublicAccessConfigurationState {
	if ebpac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ebpac.Type(), ebpac.LocalName()))
	}
	return ebpac.state
}

// EmrBlockPublicAccessConfigurationArgs contains the configurations for aws_emr_block_public_access_configuration.
type EmrBlockPublicAccessConfigurationArgs struct {
	// BlockPublicSecurityGroupRules: bool, required
	BlockPublicSecurityGroupRules terra.BoolValue `hcl:"block_public_security_group_rules,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PermittedPublicSecurityGroupRuleRange: min=0
	PermittedPublicSecurityGroupRuleRange []emrblockpublicaccessconfiguration.PermittedPublicSecurityGroupRuleRange `hcl:"permitted_public_security_group_rule_range,block" validate:"min=0"`
}
type emrBlockPublicAccessConfigurationAttributes struct {
	ref terra.Reference
}

// BlockPublicSecurityGroupRules returns a reference to field block_public_security_group_rules of aws_emr_block_public_access_configuration.
func (ebpac emrBlockPublicAccessConfigurationAttributes) BlockPublicSecurityGroupRules() terra.BoolValue {
	return terra.ReferenceAsBool(ebpac.ref.Append("block_public_security_group_rules"))
}

// Id returns a reference to field id of aws_emr_block_public_access_configuration.
func (ebpac emrBlockPublicAccessConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ebpac.ref.Append("id"))
}

func (ebpac emrBlockPublicAccessConfigurationAttributes) PermittedPublicSecurityGroupRuleRange() terra.ListValue[emrblockpublicaccessconfiguration.PermittedPublicSecurityGroupRuleRangeAttributes] {
	return terra.ReferenceAsList[emrblockpublicaccessconfiguration.PermittedPublicSecurityGroupRuleRangeAttributes](ebpac.ref.Append("permitted_public_security_group_rule_range"))
}

type emrBlockPublicAccessConfigurationState struct {
	BlockPublicSecurityGroupRules         bool                                                                           `json:"block_public_security_group_rules"`
	Id                                    string                                                                         `json:"id"`
	PermittedPublicSecurityGroupRuleRange []emrblockpublicaccessconfiguration.PermittedPublicSecurityGroupRuleRangeState `json:"permitted_public_security_group_rule_range"`
}
