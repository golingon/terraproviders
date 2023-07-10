// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ecrreplicationconfiguration "github.com/golingon/terraproviders/aws/5.7.0/ecrreplicationconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEcrReplicationConfiguration creates a new instance of [EcrReplicationConfiguration].
func NewEcrReplicationConfiguration(name string, args EcrReplicationConfigurationArgs) *EcrReplicationConfiguration {
	return &EcrReplicationConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EcrReplicationConfiguration)(nil)

// EcrReplicationConfiguration represents the Terraform resource aws_ecr_replication_configuration.
type EcrReplicationConfiguration struct {
	Name      string
	Args      EcrReplicationConfigurationArgs
	state     *ecrReplicationConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EcrReplicationConfiguration].
func (erc *EcrReplicationConfiguration) Type() string {
	return "aws_ecr_replication_configuration"
}

// LocalName returns the local name for [EcrReplicationConfiguration].
func (erc *EcrReplicationConfiguration) LocalName() string {
	return erc.Name
}

// Configuration returns the configuration (args) for [EcrReplicationConfiguration].
func (erc *EcrReplicationConfiguration) Configuration() interface{} {
	return erc.Args
}

// DependOn is used for other resources to depend on [EcrReplicationConfiguration].
func (erc *EcrReplicationConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(erc)
}

// Dependencies returns the list of resources [EcrReplicationConfiguration] depends_on.
func (erc *EcrReplicationConfiguration) Dependencies() terra.Dependencies {
	return erc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EcrReplicationConfiguration].
func (erc *EcrReplicationConfiguration) LifecycleManagement() *terra.Lifecycle {
	return erc.Lifecycle
}

// Attributes returns the attributes for [EcrReplicationConfiguration].
func (erc *EcrReplicationConfiguration) Attributes() ecrReplicationConfigurationAttributes {
	return ecrReplicationConfigurationAttributes{ref: terra.ReferenceResource(erc)}
}

// ImportState imports the given attribute values into [EcrReplicationConfiguration]'s state.
func (erc *EcrReplicationConfiguration) ImportState(av io.Reader) error {
	erc.state = &ecrReplicationConfigurationState{}
	if err := json.NewDecoder(av).Decode(erc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", erc.Type(), erc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EcrReplicationConfiguration] has state.
func (erc *EcrReplicationConfiguration) State() (*ecrReplicationConfigurationState, bool) {
	return erc.state, erc.state != nil
}

// StateMust returns the state for [EcrReplicationConfiguration]. Panics if the state is nil.
func (erc *EcrReplicationConfiguration) StateMust() *ecrReplicationConfigurationState {
	if erc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", erc.Type(), erc.LocalName()))
	}
	return erc.state
}

// EcrReplicationConfigurationArgs contains the configurations for aws_ecr_replication_configuration.
type EcrReplicationConfigurationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ReplicationConfiguration: optional
	ReplicationConfiguration *ecrreplicationconfiguration.ReplicationConfiguration `hcl:"replication_configuration,block"`
}
type ecrReplicationConfigurationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ecr_replication_configuration.
func (erc ecrReplicationConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(erc.ref.Append("id"))
}

// RegistryId returns a reference to field registry_id of aws_ecr_replication_configuration.
func (erc ecrReplicationConfigurationAttributes) RegistryId() terra.StringValue {
	return terra.ReferenceAsString(erc.ref.Append("registry_id"))
}

func (erc ecrReplicationConfigurationAttributes) ReplicationConfiguration() terra.ListValue[ecrreplicationconfiguration.ReplicationConfigurationAttributes] {
	return terra.ReferenceAsList[ecrreplicationconfiguration.ReplicationConfigurationAttributes](erc.ref.Append("replication_configuration"))
}

type ecrReplicationConfigurationState struct {
	Id                       string                                                      `json:"id"`
	RegistryId               string                                                      `json:"registry_id"`
	ReplicationConfiguration []ecrreplicationconfiguration.ReplicationConfigurationState `json:"replication_configuration"`
}
