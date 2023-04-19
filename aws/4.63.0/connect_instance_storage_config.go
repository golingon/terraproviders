// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	connectinstancestorageconfig "github.com/golingon/terraproviders/aws/4.63.0/connectinstancestorageconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewConnectInstanceStorageConfig creates a new instance of [ConnectInstanceStorageConfig].
func NewConnectInstanceStorageConfig(name string, args ConnectInstanceStorageConfigArgs) *ConnectInstanceStorageConfig {
	return &ConnectInstanceStorageConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConnectInstanceStorageConfig)(nil)

// ConnectInstanceStorageConfig represents the Terraform resource aws_connect_instance_storage_config.
type ConnectInstanceStorageConfig struct {
	Name      string
	Args      ConnectInstanceStorageConfigArgs
	state     *connectInstanceStorageConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConnectInstanceStorageConfig].
func (cisc *ConnectInstanceStorageConfig) Type() string {
	return "aws_connect_instance_storage_config"
}

// LocalName returns the local name for [ConnectInstanceStorageConfig].
func (cisc *ConnectInstanceStorageConfig) LocalName() string {
	return cisc.Name
}

// Configuration returns the configuration (args) for [ConnectInstanceStorageConfig].
func (cisc *ConnectInstanceStorageConfig) Configuration() interface{} {
	return cisc.Args
}

// DependOn is used for other resources to depend on [ConnectInstanceStorageConfig].
func (cisc *ConnectInstanceStorageConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(cisc)
}

// Dependencies returns the list of resources [ConnectInstanceStorageConfig] depends_on.
func (cisc *ConnectInstanceStorageConfig) Dependencies() terra.Dependencies {
	return cisc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConnectInstanceStorageConfig].
func (cisc *ConnectInstanceStorageConfig) LifecycleManagement() *terra.Lifecycle {
	return cisc.Lifecycle
}

// Attributes returns the attributes for [ConnectInstanceStorageConfig].
func (cisc *ConnectInstanceStorageConfig) Attributes() connectInstanceStorageConfigAttributes {
	return connectInstanceStorageConfigAttributes{ref: terra.ReferenceResource(cisc)}
}

// ImportState imports the given attribute values into [ConnectInstanceStorageConfig]'s state.
func (cisc *ConnectInstanceStorageConfig) ImportState(av io.Reader) error {
	cisc.state = &connectInstanceStorageConfigState{}
	if err := json.NewDecoder(av).Decode(cisc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cisc.Type(), cisc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConnectInstanceStorageConfig] has state.
func (cisc *ConnectInstanceStorageConfig) State() (*connectInstanceStorageConfigState, bool) {
	return cisc.state, cisc.state != nil
}

// StateMust returns the state for [ConnectInstanceStorageConfig]. Panics if the state is nil.
func (cisc *ConnectInstanceStorageConfig) StateMust() *connectInstanceStorageConfigState {
	if cisc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cisc.Type(), cisc.LocalName()))
	}
	return cisc.state
}

// ConnectInstanceStorageConfigArgs contains the configurations for aws_connect_instance_storage_config.
type ConnectInstanceStorageConfigArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// ResourceType: string, required
	ResourceType terra.StringValue `hcl:"resource_type,attr" validate:"required"`
	// StorageConfig: required
	StorageConfig *connectinstancestorageconfig.StorageConfig `hcl:"storage_config,block" validate:"required"`
}
type connectInstanceStorageConfigAttributes struct {
	ref terra.Reference
}

// AssociationId returns a reference to field association_id of aws_connect_instance_storage_config.
func (cisc connectInstanceStorageConfigAttributes) AssociationId() terra.StringValue {
	return terra.ReferenceAsString(cisc.ref.Append("association_id"))
}

// Id returns a reference to field id of aws_connect_instance_storage_config.
func (cisc connectInstanceStorageConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cisc.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_connect_instance_storage_config.
func (cisc connectInstanceStorageConfigAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(cisc.ref.Append("instance_id"))
}

// ResourceType returns a reference to field resource_type of aws_connect_instance_storage_config.
func (cisc connectInstanceStorageConfigAttributes) ResourceType() terra.StringValue {
	return terra.ReferenceAsString(cisc.ref.Append("resource_type"))
}

func (cisc connectInstanceStorageConfigAttributes) StorageConfig() terra.ListValue[connectinstancestorageconfig.StorageConfigAttributes] {
	return terra.ReferenceAsList[connectinstancestorageconfig.StorageConfigAttributes](cisc.ref.Append("storage_config"))
}

type connectInstanceStorageConfigState struct {
	AssociationId string                                            `json:"association_id"`
	Id            string                                            `json:"id"`
	InstanceId    string                                            `json:"instance_id"`
	ResourceType  string                                            `json:"resource_type"`
	StorageConfig []connectinstancestorageconfig.StorageConfigState `json:"storage_config"`
}
