// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	securityhuborganizationconfiguration "github.com/golingon/terraproviders/aws/5.44.0/securityhuborganizationconfiguration"
	"io"
)

// NewSecurityhubOrganizationConfiguration creates a new instance of [SecurityhubOrganizationConfiguration].
func NewSecurityhubOrganizationConfiguration(name string, args SecurityhubOrganizationConfigurationArgs) *SecurityhubOrganizationConfiguration {
	return &SecurityhubOrganizationConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SecurityhubOrganizationConfiguration)(nil)

// SecurityhubOrganizationConfiguration represents the Terraform resource aws_securityhub_organization_configuration.
type SecurityhubOrganizationConfiguration struct {
	Name      string
	Args      SecurityhubOrganizationConfigurationArgs
	state     *securityhubOrganizationConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SecurityhubOrganizationConfiguration].
func (soc *SecurityhubOrganizationConfiguration) Type() string {
	return "aws_securityhub_organization_configuration"
}

// LocalName returns the local name for [SecurityhubOrganizationConfiguration].
func (soc *SecurityhubOrganizationConfiguration) LocalName() string {
	return soc.Name
}

// Configuration returns the configuration (args) for [SecurityhubOrganizationConfiguration].
func (soc *SecurityhubOrganizationConfiguration) Configuration() interface{} {
	return soc.Args
}

// DependOn is used for other resources to depend on [SecurityhubOrganizationConfiguration].
func (soc *SecurityhubOrganizationConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(soc)
}

// Dependencies returns the list of resources [SecurityhubOrganizationConfiguration] depends_on.
func (soc *SecurityhubOrganizationConfiguration) Dependencies() terra.Dependencies {
	return soc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SecurityhubOrganizationConfiguration].
func (soc *SecurityhubOrganizationConfiguration) LifecycleManagement() *terra.Lifecycle {
	return soc.Lifecycle
}

// Attributes returns the attributes for [SecurityhubOrganizationConfiguration].
func (soc *SecurityhubOrganizationConfiguration) Attributes() securityhubOrganizationConfigurationAttributes {
	return securityhubOrganizationConfigurationAttributes{ref: terra.ReferenceResource(soc)}
}

// ImportState imports the given attribute values into [SecurityhubOrganizationConfiguration]'s state.
func (soc *SecurityhubOrganizationConfiguration) ImportState(av io.Reader) error {
	soc.state = &securityhubOrganizationConfigurationState{}
	if err := json.NewDecoder(av).Decode(soc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", soc.Type(), soc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SecurityhubOrganizationConfiguration] has state.
func (soc *SecurityhubOrganizationConfiguration) State() (*securityhubOrganizationConfigurationState, bool) {
	return soc.state, soc.state != nil
}

// StateMust returns the state for [SecurityhubOrganizationConfiguration]. Panics if the state is nil.
func (soc *SecurityhubOrganizationConfiguration) StateMust() *securityhubOrganizationConfigurationState {
	if soc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", soc.Type(), soc.LocalName()))
	}
	return soc.state
}

// SecurityhubOrganizationConfigurationArgs contains the configurations for aws_securityhub_organization_configuration.
type SecurityhubOrganizationConfigurationArgs struct {
	// AutoEnable: bool, required
	AutoEnable terra.BoolValue `hcl:"auto_enable,attr" validate:"required"`
	// AutoEnableStandards: string, optional
	AutoEnableStandards terra.StringValue `hcl:"auto_enable_standards,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OrganizationConfiguration: optional
	OrganizationConfiguration *securityhuborganizationconfiguration.OrganizationConfiguration `hcl:"organization_configuration,block"`
	// Timeouts: optional
	Timeouts *securityhuborganizationconfiguration.Timeouts `hcl:"timeouts,block"`
}
type securityhubOrganizationConfigurationAttributes struct {
	ref terra.Reference
}

// AutoEnable returns a reference to field auto_enable of aws_securityhub_organization_configuration.
func (soc securityhubOrganizationConfigurationAttributes) AutoEnable() terra.BoolValue {
	return terra.ReferenceAsBool(soc.ref.Append("auto_enable"))
}

// AutoEnableStandards returns a reference to field auto_enable_standards of aws_securityhub_organization_configuration.
func (soc securityhubOrganizationConfigurationAttributes) AutoEnableStandards() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("auto_enable_standards"))
}

// Id returns a reference to field id of aws_securityhub_organization_configuration.
func (soc securityhubOrganizationConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("id"))
}

func (soc securityhubOrganizationConfigurationAttributes) OrganizationConfiguration() terra.ListValue[securityhuborganizationconfiguration.OrganizationConfigurationAttributes] {
	return terra.ReferenceAsList[securityhuborganizationconfiguration.OrganizationConfigurationAttributes](soc.ref.Append("organization_configuration"))
}

func (soc securityhubOrganizationConfigurationAttributes) Timeouts() securityhuborganizationconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[securityhuborganizationconfiguration.TimeoutsAttributes](soc.ref.Append("timeouts"))
}

type securityhubOrganizationConfigurationState struct {
	AutoEnable                bool                                                                  `json:"auto_enable"`
	AutoEnableStandards       string                                                                `json:"auto_enable_standards"`
	Id                        string                                                                `json:"id"`
	OrganizationConfiguration []securityhuborganizationconfiguration.OrganizationConfigurationState `json:"organization_configuration"`
	Timeouts                  *securityhuborganizationconfiguration.TimeoutsState                   `json:"timeouts"`
}
