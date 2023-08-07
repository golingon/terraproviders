// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	monitordatacollectionruleassociation "github.com/golingon/terraproviders/azurerm/3.68.0/monitordatacollectionruleassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMonitorDataCollectionRuleAssociation creates a new instance of [MonitorDataCollectionRuleAssociation].
func NewMonitorDataCollectionRuleAssociation(name string, args MonitorDataCollectionRuleAssociationArgs) *MonitorDataCollectionRuleAssociation {
	return &MonitorDataCollectionRuleAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitorDataCollectionRuleAssociation)(nil)

// MonitorDataCollectionRuleAssociation represents the Terraform resource azurerm_monitor_data_collection_rule_association.
type MonitorDataCollectionRuleAssociation struct {
	Name      string
	Args      MonitorDataCollectionRuleAssociationArgs
	state     *monitorDataCollectionRuleAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitorDataCollectionRuleAssociation].
func (mdcra *MonitorDataCollectionRuleAssociation) Type() string {
	return "azurerm_monitor_data_collection_rule_association"
}

// LocalName returns the local name for [MonitorDataCollectionRuleAssociation].
func (mdcra *MonitorDataCollectionRuleAssociation) LocalName() string {
	return mdcra.Name
}

// Configuration returns the configuration (args) for [MonitorDataCollectionRuleAssociation].
func (mdcra *MonitorDataCollectionRuleAssociation) Configuration() interface{} {
	return mdcra.Args
}

// DependOn is used for other resources to depend on [MonitorDataCollectionRuleAssociation].
func (mdcra *MonitorDataCollectionRuleAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(mdcra)
}

// Dependencies returns the list of resources [MonitorDataCollectionRuleAssociation] depends_on.
func (mdcra *MonitorDataCollectionRuleAssociation) Dependencies() terra.Dependencies {
	return mdcra.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitorDataCollectionRuleAssociation].
func (mdcra *MonitorDataCollectionRuleAssociation) LifecycleManagement() *terra.Lifecycle {
	return mdcra.Lifecycle
}

// Attributes returns the attributes for [MonitorDataCollectionRuleAssociation].
func (mdcra *MonitorDataCollectionRuleAssociation) Attributes() monitorDataCollectionRuleAssociationAttributes {
	return monitorDataCollectionRuleAssociationAttributes{ref: terra.ReferenceResource(mdcra)}
}

// ImportState imports the given attribute values into [MonitorDataCollectionRuleAssociation]'s state.
func (mdcra *MonitorDataCollectionRuleAssociation) ImportState(av io.Reader) error {
	mdcra.state = &monitorDataCollectionRuleAssociationState{}
	if err := json.NewDecoder(av).Decode(mdcra.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mdcra.Type(), mdcra.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitorDataCollectionRuleAssociation] has state.
func (mdcra *MonitorDataCollectionRuleAssociation) State() (*monitorDataCollectionRuleAssociationState, bool) {
	return mdcra.state, mdcra.state != nil
}

// StateMust returns the state for [MonitorDataCollectionRuleAssociation]. Panics if the state is nil.
func (mdcra *MonitorDataCollectionRuleAssociation) StateMust() *monitorDataCollectionRuleAssociationState {
	if mdcra.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mdcra.Type(), mdcra.LocalName()))
	}
	return mdcra.state
}

// MonitorDataCollectionRuleAssociationArgs contains the configurations for azurerm_monitor_data_collection_rule_association.
type MonitorDataCollectionRuleAssociationArgs struct {
	// DataCollectionEndpointId: string, optional
	DataCollectionEndpointId terra.StringValue `hcl:"data_collection_endpoint_id,attr"`
	// DataCollectionRuleId: string, optional
	DataCollectionRuleId terra.StringValue `hcl:"data_collection_rule_id,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// TargetResourceId: string, required
	TargetResourceId terra.StringValue `hcl:"target_resource_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *monitordatacollectionruleassociation.Timeouts `hcl:"timeouts,block"`
}
type monitorDataCollectionRuleAssociationAttributes struct {
	ref terra.Reference
}

// DataCollectionEndpointId returns a reference to field data_collection_endpoint_id of azurerm_monitor_data_collection_rule_association.
func (mdcra monitorDataCollectionRuleAssociationAttributes) DataCollectionEndpointId() terra.StringValue {
	return terra.ReferenceAsString(mdcra.ref.Append("data_collection_endpoint_id"))
}

// DataCollectionRuleId returns a reference to field data_collection_rule_id of azurerm_monitor_data_collection_rule_association.
func (mdcra monitorDataCollectionRuleAssociationAttributes) DataCollectionRuleId() terra.StringValue {
	return terra.ReferenceAsString(mdcra.ref.Append("data_collection_rule_id"))
}

// Description returns a reference to field description of azurerm_monitor_data_collection_rule_association.
func (mdcra monitorDataCollectionRuleAssociationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mdcra.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_monitor_data_collection_rule_association.
func (mdcra monitorDataCollectionRuleAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mdcra.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_monitor_data_collection_rule_association.
func (mdcra monitorDataCollectionRuleAssociationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mdcra.ref.Append("name"))
}

// TargetResourceId returns a reference to field target_resource_id of azurerm_monitor_data_collection_rule_association.
func (mdcra monitorDataCollectionRuleAssociationAttributes) TargetResourceId() terra.StringValue {
	return terra.ReferenceAsString(mdcra.ref.Append("target_resource_id"))
}

func (mdcra monitorDataCollectionRuleAssociationAttributes) Timeouts() monitordatacollectionruleassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitordatacollectionruleassociation.TimeoutsAttributes](mdcra.ref.Append("timeouts"))
}

type monitorDataCollectionRuleAssociationState struct {
	DataCollectionEndpointId string                                              `json:"data_collection_endpoint_id"`
	DataCollectionRuleId     string                                              `json:"data_collection_rule_id"`
	Description              string                                              `json:"description"`
	Id                       string                                              `json:"id"`
	Name                     string                                              `json:"name"`
	TargetResourceId         string                                              `json:"target_resource_id"`
	Timeouts                 *monitordatacollectionruleassociation.TimeoutsState `json:"timeouts"`
}
