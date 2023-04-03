// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	databoxedgeorder "github.com/golingon/terraproviders/azurerm/3.49.0/databoxedgeorder"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataboxEdgeOrder creates a new instance of [DataboxEdgeOrder].
func NewDataboxEdgeOrder(name string, args DataboxEdgeOrderArgs) *DataboxEdgeOrder {
	return &DataboxEdgeOrder{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataboxEdgeOrder)(nil)

// DataboxEdgeOrder represents the Terraform resource azurerm_databox_edge_order.
type DataboxEdgeOrder struct {
	Name      string
	Args      DataboxEdgeOrderArgs
	state     *databoxEdgeOrderState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataboxEdgeOrder].
func (deo *DataboxEdgeOrder) Type() string {
	return "azurerm_databox_edge_order"
}

// LocalName returns the local name for [DataboxEdgeOrder].
func (deo *DataboxEdgeOrder) LocalName() string {
	return deo.Name
}

// Configuration returns the configuration (args) for [DataboxEdgeOrder].
func (deo *DataboxEdgeOrder) Configuration() interface{} {
	return deo.Args
}

// DependOn is used for other resources to depend on [DataboxEdgeOrder].
func (deo *DataboxEdgeOrder) DependOn() terra.Reference {
	return terra.ReferenceResource(deo)
}

// Dependencies returns the list of resources [DataboxEdgeOrder] depends_on.
func (deo *DataboxEdgeOrder) Dependencies() terra.Dependencies {
	return deo.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataboxEdgeOrder].
func (deo *DataboxEdgeOrder) LifecycleManagement() *terra.Lifecycle {
	return deo.Lifecycle
}

// Attributes returns the attributes for [DataboxEdgeOrder].
func (deo *DataboxEdgeOrder) Attributes() databoxEdgeOrderAttributes {
	return databoxEdgeOrderAttributes{ref: terra.ReferenceResource(deo)}
}

// ImportState imports the given attribute values into [DataboxEdgeOrder]'s state.
func (deo *DataboxEdgeOrder) ImportState(av io.Reader) error {
	deo.state = &databoxEdgeOrderState{}
	if err := json.NewDecoder(av).Decode(deo.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", deo.Type(), deo.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataboxEdgeOrder] has state.
func (deo *DataboxEdgeOrder) State() (*databoxEdgeOrderState, bool) {
	return deo.state, deo.state != nil
}

// StateMust returns the state for [DataboxEdgeOrder]. Panics if the state is nil.
func (deo *DataboxEdgeOrder) StateMust() *databoxEdgeOrderState {
	if deo.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", deo.Type(), deo.LocalName()))
	}
	return deo.state
}

// DataboxEdgeOrderArgs contains the configurations for azurerm_databox_edge_order.
type DataboxEdgeOrderArgs struct {
	// DeviceName: string, required
	DeviceName terra.StringValue `hcl:"device_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ReturnTracking: min=0
	ReturnTracking []databoxedgeorder.ReturnTracking `hcl:"return_tracking,block" validate:"min=0"`
	// ShipmentHistory: min=0
	ShipmentHistory []databoxedgeorder.ShipmentHistory `hcl:"shipment_history,block" validate:"min=0"`
	// ShipmentTracking: min=0
	ShipmentTracking []databoxedgeorder.ShipmentTracking `hcl:"shipment_tracking,block" validate:"min=0"`
	// Status: min=0
	Status []databoxedgeorder.Status `hcl:"status,block" validate:"min=0"`
	// Contact: required
	Contact *databoxedgeorder.Contact `hcl:"contact,block" validate:"required"`
	// ShipmentAddress: required
	ShipmentAddress *databoxedgeorder.ShipmentAddress `hcl:"shipment_address,block" validate:"required"`
	// Timeouts: optional
	Timeouts *databoxedgeorder.Timeouts `hcl:"timeouts,block"`
}
type databoxEdgeOrderAttributes struct {
	ref terra.Reference
}

// DeviceName returns a reference to field device_name of azurerm_databox_edge_order.
func (deo databoxEdgeOrderAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceAsString(deo.ref.Append("device_name"))
}

// Id returns a reference to field id of azurerm_databox_edge_order.
func (deo databoxEdgeOrderAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(deo.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_databox_edge_order.
func (deo databoxEdgeOrderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(deo.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_databox_edge_order.
func (deo databoxEdgeOrderAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(deo.ref.Append("resource_group_name"))
}

// SerialNumber returns a reference to field serial_number of azurerm_databox_edge_order.
func (deo databoxEdgeOrderAttributes) SerialNumber() terra.StringValue {
	return terra.ReferenceAsString(deo.ref.Append("serial_number"))
}

func (deo databoxEdgeOrderAttributes) ReturnTracking() terra.SetValue[databoxedgeorder.ReturnTrackingAttributes] {
	return terra.ReferenceAsSet[databoxedgeorder.ReturnTrackingAttributes](deo.ref.Append("return_tracking"))
}

func (deo databoxEdgeOrderAttributes) ShipmentHistory() terra.SetValue[databoxedgeorder.ShipmentHistoryAttributes] {
	return terra.ReferenceAsSet[databoxedgeorder.ShipmentHistoryAttributes](deo.ref.Append("shipment_history"))
}

func (deo databoxEdgeOrderAttributes) ShipmentTracking() terra.SetValue[databoxedgeorder.ShipmentTrackingAttributes] {
	return terra.ReferenceAsSet[databoxedgeorder.ShipmentTrackingAttributes](deo.ref.Append("shipment_tracking"))
}

func (deo databoxEdgeOrderAttributes) Status() terra.ListValue[databoxedgeorder.StatusAttributes] {
	return terra.ReferenceAsList[databoxedgeorder.StatusAttributes](deo.ref.Append("status"))
}

func (deo databoxEdgeOrderAttributes) Contact() terra.ListValue[databoxedgeorder.ContactAttributes] {
	return terra.ReferenceAsList[databoxedgeorder.ContactAttributes](deo.ref.Append("contact"))
}

func (deo databoxEdgeOrderAttributes) ShipmentAddress() terra.ListValue[databoxedgeorder.ShipmentAddressAttributes] {
	return terra.ReferenceAsList[databoxedgeorder.ShipmentAddressAttributes](deo.ref.Append("shipment_address"))
}

func (deo databoxEdgeOrderAttributes) Timeouts() databoxedgeorder.TimeoutsAttributes {
	return terra.ReferenceAsSingle[databoxedgeorder.TimeoutsAttributes](deo.ref.Append("timeouts"))
}

type databoxEdgeOrderState struct {
	DeviceName        string                                   `json:"device_name"`
	Id                string                                   `json:"id"`
	Name              string                                   `json:"name"`
	ResourceGroupName string                                   `json:"resource_group_name"`
	SerialNumber      string                                   `json:"serial_number"`
	ReturnTracking    []databoxedgeorder.ReturnTrackingState   `json:"return_tracking"`
	ShipmentHistory   []databoxedgeorder.ShipmentHistoryState  `json:"shipment_history"`
	ShipmentTracking  []databoxedgeorder.ShipmentTrackingState `json:"shipment_tracking"`
	Status            []databoxedgeorder.StatusState           `json:"status"`
	Contact           []databoxedgeorder.ContactState          `json:"contact"`
	ShipmentAddress   []databoxedgeorder.ShipmentAddressState  `json:"shipment_address"`
	Timeouts          *databoxedgeorder.TimeoutsState          `json:"timeouts"`
}
