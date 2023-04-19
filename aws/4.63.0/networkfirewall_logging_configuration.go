// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	networkfirewallloggingconfiguration "github.com/golingon/terraproviders/aws/4.63.0/networkfirewallloggingconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkfirewallLoggingConfiguration creates a new instance of [NetworkfirewallLoggingConfiguration].
func NewNetworkfirewallLoggingConfiguration(name string, args NetworkfirewallLoggingConfigurationArgs) *NetworkfirewallLoggingConfiguration {
	return &NetworkfirewallLoggingConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkfirewallLoggingConfiguration)(nil)

// NetworkfirewallLoggingConfiguration represents the Terraform resource aws_networkfirewall_logging_configuration.
type NetworkfirewallLoggingConfiguration struct {
	Name      string
	Args      NetworkfirewallLoggingConfigurationArgs
	state     *networkfirewallLoggingConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkfirewallLoggingConfiguration].
func (nlc *NetworkfirewallLoggingConfiguration) Type() string {
	return "aws_networkfirewall_logging_configuration"
}

// LocalName returns the local name for [NetworkfirewallLoggingConfiguration].
func (nlc *NetworkfirewallLoggingConfiguration) LocalName() string {
	return nlc.Name
}

// Configuration returns the configuration (args) for [NetworkfirewallLoggingConfiguration].
func (nlc *NetworkfirewallLoggingConfiguration) Configuration() interface{} {
	return nlc.Args
}

// DependOn is used for other resources to depend on [NetworkfirewallLoggingConfiguration].
func (nlc *NetworkfirewallLoggingConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(nlc)
}

// Dependencies returns the list of resources [NetworkfirewallLoggingConfiguration] depends_on.
func (nlc *NetworkfirewallLoggingConfiguration) Dependencies() terra.Dependencies {
	return nlc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkfirewallLoggingConfiguration].
func (nlc *NetworkfirewallLoggingConfiguration) LifecycleManagement() *terra.Lifecycle {
	return nlc.Lifecycle
}

// Attributes returns the attributes for [NetworkfirewallLoggingConfiguration].
func (nlc *NetworkfirewallLoggingConfiguration) Attributes() networkfirewallLoggingConfigurationAttributes {
	return networkfirewallLoggingConfigurationAttributes{ref: terra.ReferenceResource(nlc)}
}

// ImportState imports the given attribute values into [NetworkfirewallLoggingConfiguration]'s state.
func (nlc *NetworkfirewallLoggingConfiguration) ImportState(av io.Reader) error {
	nlc.state = &networkfirewallLoggingConfigurationState{}
	if err := json.NewDecoder(av).Decode(nlc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nlc.Type(), nlc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkfirewallLoggingConfiguration] has state.
func (nlc *NetworkfirewallLoggingConfiguration) State() (*networkfirewallLoggingConfigurationState, bool) {
	return nlc.state, nlc.state != nil
}

// StateMust returns the state for [NetworkfirewallLoggingConfiguration]. Panics if the state is nil.
func (nlc *NetworkfirewallLoggingConfiguration) StateMust() *networkfirewallLoggingConfigurationState {
	if nlc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nlc.Type(), nlc.LocalName()))
	}
	return nlc.state
}

// NetworkfirewallLoggingConfigurationArgs contains the configurations for aws_networkfirewall_logging_configuration.
type NetworkfirewallLoggingConfigurationArgs struct {
	// FirewallArn: string, required
	FirewallArn terra.StringValue `hcl:"firewall_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LoggingConfiguration: required
	LoggingConfiguration *networkfirewallloggingconfiguration.LoggingConfiguration `hcl:"logging_configuration,block" validate:"required"`
}
type networkfirewallLoggingConfigurationAttributes struct {
	ref terra.Reference
}

// FirewallArn returns a reference to field firewall_arn of aws_networkfirewall_logging_configuration.
func (nlc networkfirewallLoggingConfigurationAttributes) FirewallArn() terra.StringValue {
	return terra.ReferenceAsString(nlc.ref.Append("firewall_arn"))
}

// Id returns a reference to field id of aws_networkfirewall_logging_configuration.
func (nlc networkfirewallLoggingConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nlc.ref.Append("id"))
}

func (nlc networkfirewallLoggingConfigurationAttributes) LoggingConfiguration() terra.ListValue[networkfirewallloggingconfiguration.LoggingConfigurationAttributes] {
	return terra.ReferenceAsList[networkfirewallloggingconfiguration.LoggingConfigurationAttributes](nlc.ref.Append("logging_configuration"))
}

type networkfirewallLoggingConfigurationState struct {
	FirewallArn          string                                                          `json:"firewall_arn"`
	Id                   string                                                          `json:"id"`
	LoggingConfiguration []networkfirewallloggingconfiguration.LoggingConfigurationState `json:"logging_configuration"`
}
