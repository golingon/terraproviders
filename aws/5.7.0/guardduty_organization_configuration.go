// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	guarddutyorganizationconfiguration "github.com/golingon/terraproviders/aws/5.7.0/guarddutyorganizationconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGuarddutyOrganizationConfiguration creates a new instance of [GuarddutyOrganizationConfiguration].
func NewGuarddutyOrganizationConfiguration(name string, args GuarddutyOrganizationConfigurationArgs) *GuarddutyOrganizationConfiguration {
	return &GuarddutyOrganizationConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GuarddutyOrganizationConfiguration)(nil)

// GuarddutyOrganizationConfiguration represents the Terraform resource aws_guardduty_organization_configuration.
type GuarddutyOrganizationConfiguration struct {
	Name      string
	Args      GuarddutyOrganizationConfigurationArgs
	state     *guarddutyOrganizationConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GuarddutyOrganizationConfiguration].
func (goc *GuarddutyOrganizationConfiguration) Type() string {
	return "aws_guardduty_organization_configuration"
}

// LocalName returns the local name for [GuarddutyOrganizationConfiguration].
func (goc *GuarddutyOrganizationConfiguration) LocalName() string {
	return goc.Name
}

// Configuration returns the configuration (args) for [GuarddutyOrganizationConfiguration].
func (goc *GuarddutyOrganizationConfiguration) Configuration() interface{} {
	return goc.Args
}

// DependOn is used for other resources to depend on [GuarddutyOrganizationConfiguration].
func (goc *GuarddutyOrganizationConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(goc)
}

// Dependencies returns the list of resources [GuarddutyOrganizationConfiguration] depends_on.
func (goc *GuarddutyOrganizationConfiguration) Dependencies() terra.Dependencies {
	return goc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GuarddutyOrganizationConfiguration].
func (goc *GuarddutyOrganizationConfiguration) LifecycleManagement() *terra.Lifecycle {
	return goc.Lifecycle
}

// Attributes returns the attributes for [GuarddutyOrganizationConfiguration].
func (goc *GuarddutyOrganizationConfiguration) Attributes() guarddutyOrganizationConfigurationAttributes {
	return guarddutyOrganizationConfigurationAttributes{ref: terra.ReferenceResource(goc)}
}

// ImportState imports the given attribute values into [GuarddutyOrganizationConfiguration]'s state.
func (goc *GuarddutyOrganizationConfiguration) ImportState(av io.Reader) error {
	goc.state = &guarddutyOrganizationConfigurationState{}
	if err := json.NewDecoder(av).Decode(goc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", goc.Type(), goc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GuarddutyOrganizationConfiguration] has state.
func (goc *GuarddutyOrganizationConfiguration) State() (*guarddutyOrganizationConfigurationState, bool) {
	return goc.state, goc.state != nil
}

// StateMust returns the state for [GuarddutyOrganizationConfiguration]. Panics if the state is nil.
func (goc *GuarddutyOrganizationConfiguration) StateMust() *guarddutyOrganizationConfigurationState {
	if goc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", goc.Type(), goc.LocalName()))
	}
	return goc.state
}

// GuarddutyOrganizationConfigurationArgs contains the configurations for aws_guardduty_organization_configuration.
type GuarddutyOrganizationConfigurationArgs struct {
	// AutoEnable: bool, optional
	AutoEnable terra.BoolValue `hcl:"auto_enable,attr"`
	// AutoEnableOrganizationMembers: string, optional
	AutoEnableOrganizationMembers terra.StringValue `hcl:"auto_enable_organization_members,attr"`
	// DetectorId: string, required
	DetectorId terra.StringValue `hcl:"detector_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Datasources: optional
	Datasources *guarddutyorganizationconfiguration.Datasources `hcl:"datasources,block"`
}
type guarddutyOrganizationConfigurationAttributes struct {
	ref terra.Reference
}

// AutoEnable returns a reference to field auto_enable of aws_guardduty_organization_configuration.
func (goc guarddutyOrganizationConfigurationAttributes) AutoEnable() terra.BoolValue {
	return terra.ReferenceAsBool(goc.ref.Append("auto_enable"))
}

// AutoEnableOrganizationMembers returns a reference to field auto_enable_organization_members of aws_guardduty_organization_configuration.
func (goc guarddutyOrganizationConfigurationAttributes) AutoEnableOrganizationMembers() terra.StringValue {
	return terra.ReferenceAsString(goc.ref.Append("auto_enable_organization_members"))
}

// DetectorId returns a reference to field detector_id of aws_guardduty_organization_configuration.
func (goc guarddutyOrganizationConfigurationAttributes) DetectorId() terra.StringValue {
	return terra.ReferenceAsString(goc.ref.Append("detector_id"))
}

// Id returns a reference to field id of aws_guardduty_organization_configuration.
func (goc guarddutyOrganizationConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(goc.ref.Append("id"))
}

func (goc guarddutyOrganizationConfigurationAttributes) Datasources() terra.ListValue[guarddutyorganizationconfiguration.DatasourcesAttributes] {
	return terra.ReferenceAsList[guarddutyorganizationconfiguration.DatasourcesAttributes](goc.ref.Append("datasources"))
}

type guarddutyOrganizationConfigurationState struct {
	AutoEnable                    bool                                                  `json:"auto_enable"`
	AutoEnableOrganizationMembers string                                                `json:"auto_enable_organization_members"`
	DetectorId                    string                                                `json:"detector_id"`
	Id                            string                                                `json:"id"`
	Datasources                   []guarddutyorganizationconfiguration.DatasourcesState `json:"datasources"`
}
