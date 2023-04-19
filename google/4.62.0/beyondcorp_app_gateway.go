// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	beyondcorpappgateway "github.com/golingon/terraproviders/google/4.62.0/beyondcorpappgateway"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBeyondcorpAppGateway creates a new instance of [BeyondcorpAppGateway].
func NewBeyondcorpAppGateway(name string, args BeyondcorpAppGatewayArgs) *BeyondcorpAppGateway {
	return &BeyondcorpAppGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BeyondcorpAppGateway)(nil)

// BeyondcorpAppGateway represents the Terraform resource google_beyondcorp_app_gateway.
type BeyondcorpAppGateway struct {
	Name      string
	Args      BeyondcorpAppGatewayArgs
	state     *beyondcorpAppGatewayState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BeyondcorpAppGateway].
func (bag *BeyondcorpAppGateway) Type() string {
	return "google_beyondcorp_app_gateway"
}

// LocalName returns the local name for [BeyondcorpAppGateway].
func (bag *BeyondcorpAppGateway) LocalName() string {
	return bag.Name
}

// Configuration returns the configuration (args) for [BeyondcorpAppGateway].
func (bag *BeyondcorpAppGateway) Configuration() interface{} {
	return bag.Args
}

// DependOn is used for other resources to depend on [BeyondcorpAppGateway].
func (bag *BeyondcorpAppGateway) DependOn() terra.Reference {
	return terra.ReferenceResource(bag)
}

// Dependencies returns the list of resources [BeyondcorpAppGateway] depends_on.
func (bag *BeyondcorpAppGateway) Dependencies() terra.Dependencies {
	return bag.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BeyondcorpAppGateway].
func (bag *BeyondcorpAppGateway) LifecycleManagement() *terra.Lifecycle {
	return bag.Lifecycle
}

// Attributes returns the attributes for [BeyondcorpAppGateway].
func (bag *BeyondcorpAppGateway) Attributes() beyondcorpAppGatewayAttributes {
	return beyondcorpAppGatewayAttributes{ref: terra.ReferenceResource(bag)}
}

// ImportState imports the given attribute values into [BeyondcorpAppGateway]'s state.
func (bag *BeyondcorpAppGateway) ImportState(av io.Reader) error {
	bag.state = &beyondcorpAppGatewayState{}
	if err := json.NewDecoder(av).Decode(bag.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bag.Type(), bag.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BeyondcorpAppGateway] has state.
func (bag *BeyondcorpAppGateway) State() (*beyondcorpAppGatewayState, bool) {
	return bag.state, bag.state != nil
}

// StateMust returns the state for [BeyondcorpAppGateway]. Panics if the state is nil.
func (bag *BeyondcorpAppGateway) StateMust() *beyondcorpAppGatewayState {
	if bag.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bag.Type(), bag.LocalName()))
	}
	return bag.state
}

// BeyondcorpAppGatewayArgs contains the configurations for google_beyondcorp_app_gateway.
type BeyondcorpAppGatewayArgs struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// HostType: string, optional
	HostType terra.StringValue `hcl:"host_type,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// AllocatedConnections: min=0
	AllocatedConnections []beyondcorpappgateway.AllocatedConnections `hcl:"allocated_connections,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *beyondcorpappgateway.Timeouts `hcl:"timeouts,block"`
}
type beyondcorpAppGatewayAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_beyondcorp_app_gateway.
func (bag beyondcorpAppGatewayAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(bag.ref.Append("display_name"))
}

// HostType returns a reference to field host_type of google_beyondcorp_app_gateway.
func (bag beyondcorpAppGatewayAttributes) HostType() terra.StringValue {
	return terra.ReferenceAsString(bag.ref.Append("host_type"))
}

// Id returns a reference to field id of google_beyondcorp_app_gateway.
func (bag beyondcorpAppGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bag.ref.Append("id"))
}

// Labels returns a reference to field labels of google_beyondcorp_app_gateway.
func (bag beyondcorpAppGatewayAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bag.ref.Append("labels"))
}

// Name returns a reference to field name of google_beyondcorp_app_gateway.
func (bag beyondcorpAppGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bag.ref.Append("name"))
}

// Project returns a reference to field project of google_beyondcorp_app_gateway.
func (bag beyondcorpAppGatewayAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bag.ref.Append("project"))
}

// Region returns a reference to field region of google_beyondcorp_app_gateway.
func (bag beyondcorpAppGatewayAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(bag.ref.Append("region"))
}

// State returns a reference to field state of google_beyondcorp_app_gateway.
func (bag beyondcorpAppGatewayAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(bag.ref.Append("state"))
}

// Type returns a reference to field type of google_beyondcorp_app_gateway.
func (bag beyondcorpAppGatewayAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(bag.ref.Append("type"))
}

// Uri returns a reference to field uri of google_beyondcorp_app_gateway.
func (bag beyondcorpAppGatewayAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(bag.ref.Append("uri"))
}

func (bag beyondcorpAppGatewayAttributes) AllocatedConnections() terra.ListValue[beyondcorpappgateway.AllocatedConnectionsAttributes] {
	return terra.ReferenceAsList[beyondcorpappgateway.AllocatedConnectionsAttributes](bag.ref.Append("allocated_connections"))
}

func (bag beyondcorpAppGatewayAttributes) Timeouts() beyondcorpappgateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[beyondcorpappgateway.TimeoutsAttributes](bag.ref.Append("timeouts"))
}

type beyondcorpAppGatewayState struct {
	DisplayName          string                                           `json:"display_name"`
	HostType             string                                           `json:"host_type"`
	Id                   string                                           `json:"id"`
	Labels               map[string]string                                `json:"labels"`
	Name                 string                                           `json:"name"`
	Project              string                                           `json:"project"`
	Region               string                                           `json:"region"`
	State                string                                           `json:"state"`
	Type                 string                                           `json:"type"`
	Uri                  string                                           `json:"uri"`
	AllocatedConnections []beyondcorpappgateway.AllocatedConnectionsState `json:"allocated_connections"`
	Timeouts             *beyondcorpappgateway.TimeoutsState              `json:"timeouts"`
}
