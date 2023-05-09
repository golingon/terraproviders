// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorylinkedservicesftp "github.com/golingon/terraproviders/azurerm/3.55.0/datafactorylinkedservicesftp"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryLinkedServiceSftp creates a new instance of [DataFactoryLinkedServiceSftp].
func NewDataFactoryLinkedServiceSftp(name string, args DataFactoryLinkedServiceSftpArgs) *DataFactoryLinkedServiceSftp {
	return &DataFactoryLinkedServiceSftp{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryLinkedServiceSftp)(nil)

// DataFactoryLinkedServiceSftp represents the Terraform resource azurerm_data_factory_linked_service_sftp.
type DataFactoryLinkedServiceSftp struct {
	Name      string
	Args      DataFactoryLinkedServiceSftpArgs
	state     *dataFactoryLinkedServiceSftpState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryLinkedServiceSftp].
func (dflss *DataFactoryLinkedServiceSftp) Type() string {
	return "azurerm_data_factory_linked_service_sftp"
}

// LocalName returns the local name for [DataFactoryLinkedServiceSftp].
func (dflss *DataFactoryLinkedServiceSftp) LocalName() string {
	return dflss.Name
}

// Configuration returns the configuration (args) for [DataFactoryLinkedServiceSftp].
func (dflss *DataFactoryLinkedServiceSftp) Configuration() interface{} {
	return dflss.Args
}

// DependOn is used for other resources to depend on [DataFactoryLinkedServiceSftp].
func (dflss *DataFactoryLinkedServiceSftp) DependOn() terra.Reference {
	return terra.ReferenceResource(dflss)
}

// Dependencies returns the list of resources [DataFactoryLinkedServiceSftp] depends_on.
func (dflss *DataFactoryLinkedServiceSftp) Dependencies() terra.Dependencies {
	return dflss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryLinkedServiceSftp].
func (dflss *DataFactoryLinkedServiceSftp) LifecycleManagement() *terra.Lifecycle {
	return dflss.Lifecycle
}

// Attributes returns the attributes for [DataFactoryLinkedServiceSftp].
func (dflss *DataFactoryLinkedServiceSftp) Attributes() dataFactoryLinkedServiceSftpAttributes {
	return dataFactoryLinkedServiceSftpAttributes{ref: terra.ReferenceResource(dflss)}
}

// ImportState imports the given attribute values into [DataFactoryLinkedServiceSftp]'s state.
func (dflss *DataFactoryLinkedServiceSftp) ImportState(av io.Reader) error {
	dflss.state = &dataFactoryLinkedServiceSftpState{}
	if err := json.NewDecoder(av).Decode(dflss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dflss.Type(), dflss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryLinkedServiceSftp] has state.
func (dflss *DataFactoryLinkedServiceSftp) State() (*dataFactoryLinkedServiceSftpState, bool) {
	return dflss.state, dflss.state != nil
}

// StateMust returns the state for [DataFactoryLinkedServiceSftp]. Panics if the state is nil.
func (dflss *DataFactoryLinkedServiceSftp) StateMust() *dataFactoryLinkedServiceSftpState {
	if dflss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dflss.Type(), dflss.LocalName()))
	}
	return dflss.state
}

// DataFactoryLinkedServiceSftpArgs contains the configurations for azurerm_data_factory_linked_service_sftp.
type DataFactoryLinkedServiceSftpArgs struct {
	// AdditionalProperties: map of string, optional
	AdditionalProperties terra.MapValue[terra.StringValue] `hcl:"additional_properties,attr"`
	// Annotations: list of string, optional
	Annotations terra.ListValue[terra.StringValue] `hcl:"annotations,attr"`
	// AuthenticationType: string, required
	AuthenticationType terra.StringValue `hcl:"authentication_type,attr" validate:"required"`
	// DataFactoryId: string, required
	DataFactoryId terra.StringValue `hcl:"data_factory_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Host: string, required
	Host terra.StringValue `hcl:"host,attr" validate:"required"`
	// HostKeyFingerprint: string, optional
	HostKeyFingerprint terra.StringValue `hcl:"host_key_fingerprint,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IntegrationRuntimeName: string, optional
	IntegrationRuntimeName terra.StringValue `hcl:"integration_runtime_name,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// Port: number, required
	Port terra.NumberValue `hcl:"port,attr" validate:"required"`
	// SkipHostKeyValidation: bool, optional
	SkipHostKeyValidation terra.BoolValue `hcl:"skip_host_key_validation,attr"`
	// Username: string, required
	Username terra.StringValue `hcl:"username,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datafactorylinkedservicesftp.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryLinkedServiceSftpAttributes struct {
	ref terra.Reference
}

// AdditionalProperties returns a reference to field additional_properties of azurerm_data_factory_linked_service_sftp.
func (dflss dataFactoryLinkedServiceSftpAttributes) AdditionalProperties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflss.ref.Append("additional_properties"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_linked_service_sftp.
func (dflss dataFactoryLinkedServiceSftpAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dflss.ref.Append("annotations"))
}

// AuthenticationType returns a reference to field authentication_type of azurerm_data_factory_linked_service_sftp.
func (dflss dataFactoryLinkedServiceSftpAttributes) AuthenticationType() terra.StringValue {
	return terra.ReferenceAsString(dflss.ref.Append("authentication_type"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_linked_service_sftp.
func (dflss dataFactoryLinkedServiceSftpAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dflss.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_linked_service_sftp.
func (dflss dataFactoryLinkedServiceSftpAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dflss.ref.Append("description"))
}

// Host returns a reference to field host of azurerm_data_factory_linked_service_sftp.
func (dflss dataFactoryLinkedServiceSftpAttributes) Host() terra.StringValue {
	return terra.ReferenceAsString(dflss.ref.Append("host"))
}

// HostKeyFingerprint returns a reference to field host_key_fingerprint of azurerm_data_factory_linked_service_sftp.
func (dflss dataFactoryLinkedServiceSftpAttributes) HostKeyFingerprint() terra.StringValue {
	return terra.ReferenceAsString(dflss.ref.Append("host_key_fingerprint"))
}

// Id returns a reference to field id of azurerm_data_factory_linked_service_sftp.
func (dflss dataFactoryLinkedServiceSftpAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dflss.ref.Append("id"))
}

// IntegrationRuntimeName returns a reference to field integration_runtime_name of azurerm_data_factory_linked_service_sftp.
func (dflss dataFactoryLinkedServiceSftpAttributes) IntegrationRuntimeName() terra.StringValue {
	return terra.ReferenceAsString(dflss.ref.Append("integration_runtime_name"))
}

// Name returns a reference to field name of azurerm_data_factory_linked_service_sftp.
func (dflss dataFactoryLinkedServiceSftpAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dflss.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_data_factory_linked_service_sftp.
func (dflss dataFactoryLinkedServiceSftpAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dflss.ref.Append("parameters"))
}

// Password returns a reference to field password of azurerm_data_factory_linked_service_sftp.
func (dflss dataFactoryLinkedServiceSftpAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(dflss.ref.Append("password"))
}

// Port returns a reference to field port of azurerm_data_factory_linked_service_sftp.
func (dflss dataFactoryLinkedServiceSftpAttributes) Port() terra.NumberValue {
	return terra.ReferenceAsNumber(dflss.ref.Append("port"))
}

// SkipHostKeyValidation returns a reference to field skip_host_key_validation of azurerm_data_factory_linked_service_sftp.
func (dflss dataFactoryLinkedServiceSftpAttributes) SkipHostKeyValidation() terra.BoolValue {
	return terra.ReferenceAsBool(dflss.ref.Append("skip_host_key_validation"))
}

// Username returns a reference to field username of azurerm_data_factory_linked_service_sftp.
func (dflss dataFactoryLinkedServiceSftpAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(dflss.ref.Append("username"))
}

func (dflss dataFactoryLinkedServiceSftpAttributes) Timeouts() datafactorylinkedservicesftp.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorylinkedservicesftp.TimeoutsAttributes](dflss.ref.Append("timeouts"))
}

type dataFactoryLinkedServiceSftpState struct {
	AdditionalProperties   map[string]string                           `json:"additional_properties"`
	Annotations            []string                                    `json:"annotations"`
	AuthenticationType     string                                      `json:"authentication_type"`
	DataFactoryId          string                                      `json:"data_factory_id"`
	Description            string                                      `json:"description"`
	Host                   string                                      `json:"host"`
	HostKeyFingerprint     string                                      `json:"host_key_fingerprint"`
	Id                     string                                      `json:"id"`
	IntegrationRuntimeName string                                      `json:"integration_runtime_name"`
	Name                   string                                      `json:"name"`
	Parameters             map[string]string                           `json:"parameters"`
	Password               string                                      `json:"password"`
	Port                   float64                                     `json:"port"`
	SkipHostKeyValidation  bool                                        `json:"skip_host_key_validation"`
	Username               string                                      `json:"username"`
	Timeouts               *datafactorylinkedservicesftp.TimeoutsState `json:"timeouts"`
}
