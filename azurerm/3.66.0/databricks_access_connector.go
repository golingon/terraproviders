// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	databricksaccessconnector "github.com/golingon/terraproviders/azurerm/3.66.0/databricksaccessconnector"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDatabricksAccessConnector creates a new instance of [DatabricksAccessConnector].
func NewDatabricksAccessConnector(name string, args DatabricksAccessConnectorArgs) *DatabricksAccessConnector {
	return &DatabricksAccessConnector{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatabricksAccessConnector)(nil)

// DatabricksAccessConnector represents the Terraform resource azurerm_databricks_access_connector.
type DatabricksAccessConnector struct {
	Name      string
	Args      DatabricksAccessConnectorArgs
	state     *databricksAccessConnectorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DatabricksAccessConnector].
func (dac *DatabricksAccessConnector) Type() string {
	return "azurerm_databricks_access_connector"
}

// LocalName returns the local name for [DatabricksAccessConnector].
func (dac *DatabricksAccessConnector) LocalName() string {
	return dac.Name
}

// Configuration returns the configuration (args) for [DatabricksAccessConnector].
func (dac *DatabricksAccessConnector) Configuration() interface{} {
	return dac.Args
}

// DependOn is used for other resources to depend on [DatabricksAccessConnector].
func (dac *DatabricksAccessConnector) DependOn() terra.Reference {
	return terra.ReferenceResource(dac)
}

// Dependencies returns the list of resources [DatabricksAccessConnector] depends_on.
func (dac *DatabricksAccessConnector) Dependencies() terra.Dependencies {
	return dac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DatabricksAccessConnector].
func (dac *DatabricksAccessConnector) LifecycleManagement() *terra.Lifecycle {
	return dac.Lifecycle
}

// Attributes returns the attributes for [DatabricksAccessConnector].
func (dac *DatabricksAccessConnector) Attributes() databricksAccessConnectorAttributes {
	return databricksAccessConnectorAttributes{ref: terra.ReferenceResource(dac)}
}

// ImportState imports the given attribute values into [DatabricksAccessConnector]'s state.
func (dac *DatabricksAccessConnector) ImportState(av io.Reader) error {
	dac.state = &databricksAccessConnectorState{}
	if err := json.NewDecoder(av).Decode(dac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dac.Type(), dac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DatabricksAccessConnector] has state.
func (dac *DatabricksAccessConnector) State() (*databricksAccessConnectorState, bool) {
	return dac.state, dac.state != nil
}

// StateMust returns the state for [DatabricksAccessConnector]. Panics if the state is nil.
func (dac *DatabricksAccessConnector) StateMust() *databricksAccessConnectorState {
	if dac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dac.Type(), dac.LocalName()))
	}
	return dac.state
}

// DatabricksAccessConnectorArgs contains the configurations for azurerm_databricks_access_connector.
type DatabricksAccessConnectorArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Identity: optional
	Identity *databricksaccessconnector.Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *databricksaccessconnector.Timeouts `hcl:"timeouts,block"`
}
type databricksAccessConnectorAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_databricks_access_connector.
func (dac databricksAccessConnectorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dac.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_databricks_access_connector.
func (dac databricksAccessConnectorAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dac.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_databricks_access_connector.
func (dac databricksAccessConnectorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dac.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_databricks_access_connector.
func (dac databricksAccessConnectorAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dac.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_databricks_access_connector.
func (dac databricksAccessConnectorAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dac.ref.Append("tags"))
}

func (dac databricksAccessConnectorAttributes) Identity() terra.ListValue[databricksaccessconnector.IdentityAttributes] {
	return terra.ReferenceAsList[databricksaccessconnector.IdentityAttributes](dac.ref.Append("identity"))
}

func (dac databricksAccessConnectorAttributes) Timeouts() databricksaccessconnector.TimeoutsAttributes {
	return terra.ReferenceAsSingle[databricksaccessconnector.TimeoutsAttributes](dac.ref.Append("timeouts"))
}

type databricksAccessConnectorState struct {
	Id                string                                    `json:"id"`
	Location          string                                    `json:"location"`
	Name              string                                    `json:"name"`
	ResourceGroupName string                                    `json:"resource_group_name"`
	Tags              map[string]string                         `json:"tags"`
	Identity          []databricksaccessconnector.IdentityState `json:"identity"`
	Timeouts          *databricksaccessconnector.TimeoutsState  `json:"timeouts"`
}
