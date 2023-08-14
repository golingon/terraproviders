// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ecrregistryscanningconfiguration "github.com/golingon/terraproviders/aws/5.12.0/ecrregistryscanningconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEcrRegistryScanningConfiguration creates a new instance of [EcrRegistryScanningConfiguration].
func NewEcrRegistryScanningConfiguration(name string, args EcrRegistryScanningConfigurationArgs) *EcrRegistryScanningConfiguration {
	return &EcrRegistryScanningConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EcrRegistryScanningConfiguration)(nil)

// EcrRegistryScanningConfiguration represents the Terraform resource aws_ecr_registry_scanning_configuration.
type EcrRegistryScanningConfiguration struct {
	Name      string
	Args      EcrRegistryScanningConfigurationArgs
	state     *ecrRegistryScanningConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EcrRegistryScanningConfiguration].
func (ersc *EcrRegistryScanningConfiguration) Type() string {
	return "aws_ecr_registry_scanning_configuration"
}

// LocalName returns the local name for [EcrRegistryScanningConfiguration].
func (ersc *EcrRegistryScanningConfiguration) LocalName() string {
	return ersc.Name
}

// Configuration returns the configuration (args) for [EcrRegistryScanningConfiguration].
func (ersc *EcrRegistryScanningConfiguration) Configuration() interface{} {
	return ersc.Args
}

// DependOn is used for other resources to depend on [EcrRegistryScanningConfiguration].
func (ersc *EcrRegistryScanningConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(ersc)
}

// Dependencies returns the list of resources [EcrRegistryScanningConfiguration] depends_on.
func (ersc *EcrRegistryScanningConfiguration) Dependencies() terra.Dependencies {
	return ersc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EcrRegistryScanningConfiguration].
func (ersc *EcrRegistryScanningConfiguration) LifecycleManagement() *terra.Lifecycle {
	return ersc.Lifecycle
}

// Attributes returns the attributes for [EcrRegistryScanningConfiguration].
func (ersc *EcrRegistryScanningConfiguration) Attributes() ecrRegistryScanningConfigurationAttributes {
	return ecrRegistryScanningConfigurationAttributes{ref: terra.ReferenceResource(ersc)}
}

// ImportState imports the given attribute values into [EcrRegistryScanningConfiguration]'s state.
func (ersc *EcrRegistryScanningConfiguration) ImportState(av io.Reader) error {
	ersc.state = &ecrRegistryScanningConfigurationState{}
	if err := json.NewDecoder(av).Decode(ersc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ersc.Type(), ersc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EcrRegistryScanningConfiguration] has state.
func (ersc *EcrRegistryScanningConfiguration) State() (*ecrRegistryScanningConfigurationState, bool) {
	return ersc.state, ersc.state != nil
}

// StateMust returns the state for [EcrRegistryScanningConfiguration]. Panics if the state is nil.
func (ersc *EcrRegistryScanningConfiguration) StateMust() *ecrRegistryScanningConfigurationState {
	if ersc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ersc.Type(), ersc.LocalName()))
	}
	return ersc.state
}

// EcrRegistryScanningConfigurationArgs contains the configurations for aws_ecr_registry_scanning_configuration.
type EcrRegistryScanningConfigurationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ScanType: string, required
	ScanType terra.StringValue `hcl:"scan_type,attr" validate:"required"`
	// Rule: min=0,max=100
	Rule []ecrregistryscanningconfiguration.Rule `hcl:"rule,block" validate:"min=0,max=100"`
}
type ecrRegistryScanningConfigurationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ecr_registry_scanning_configuration.
func (ersc ecrRegistryScanningConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ersc.ref.Append("id"))
}

// RegistryId returns a reference to field registry_id of aws_ecr_registry_scanning_configuration.
func (ersc ecrRegistryScanningConfigurationAttributes) RegistryId() terra.StringValue {
	return terra.ReferenceAsString(ersc.ref.Append("registry_id"))
}

// ScanType returns a reference to field scan_type of aws_ecr_registry_scanning_configuration.
func (ersc ecrRegistryScanningConfigurationAttributes) ScanType() terra.StringValue {
	return terra.ReferenceAsString(ersc.ref.Append("scan_type"))
}

func (ersc ecrRegistryScanningConfigurationAttributes) Rule() terra.SetValue[ecrregistryscanningconfiguration.RuleAttributes] {
	return terra.ReferenceAsSet[ecrregistryscanningconfiguration.RuleAttributes](ersc.ref.Append("rule"))
}

type ecrRegistryScanningConfigurationState struct {
	Id         string                                       `json:"id"`
	RegistryId string                                       `json:"registry_id"`
	ScanType   string                                       `json:"scan_type"`
	Rule       []ecrregistryscanningconfiguration.RuleState `json:"rule"`
}
