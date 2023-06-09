// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	beyondcorpappconnection "github.com/golingon/terraproviders/googlebeta/4.66.0/beyondcorpappconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBeyondcorpAppConnection creates a new instance of [BeyondcorpAppConnection].
func NewBeyondcorpAppConnection(name string, args BeyondcorpAppConnectionArgs) *BeyondcorpAppConnection {
	return &BeyondcorpAppConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BeyondcorpAppConnection)(nil)

// BeyondcorpAppConnection represents the Terraform resource google_beyondcorp_app_connection.
type BeyondcorpAppConnection struct {
	Name      string
	Args      BeyondcorpAppConnectionArgs
	state     *beyondcorpAppConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BeyondcorpAppConnection].
func (bac *BeyondcorpAppConnection) Type() string {
	return "google_beyondcorp_app_connection"
}

// LocalName returns the local name for [BeyondcorpAppConnection].
func (bac *BeyondcorpAppConnection) LocalName() string {
	return bac.Name
}

// Configuration returns the configuration (args) for [BeyondcorpAppConnection].
func (bac *BeyondcorpAppConnection) Configuration() interface{} {
	return bac.Args
}

// DependOn is used for other resources to depend on [BeyondcorpAppConnection].
func (bac *BeyondcorpAppConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(bac)
}

// Dependencies returns the list of resources [BeyondcorpAppConnection] depends_on.
func (bac *BeyondcorpAppConnection) Dependencies() terra.Dependencies {
	return bac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BeyondcorpAppConnection].
func (bac *BeyondcorpAppConnection) LifecycleManagement() *terra.Lifecycle {
	return bac.Lifecycle
}

// Attributes returns the attributes for [BeyondcorpAppConnection].
func (bac *BeyondcorpAppConnection) Attributes() beyondcorpAppConnectionAttributes {
	return beyondcorpAppConnectionAttributes{ref: terra.ReferenceResource(bac)}
}

// ImportState imports the given attribute values into [BeyondcorpAppConnection]'s state.
func (bac *BeyondcorpAppConnection) ImportState(av io.Reader) error {
	bac.state = &beyondcorpAppConnectionState{}
	if err := json.NewDecoder(av).Decode(bac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bac.Type(), bac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BeyondcorpAppConnection] has state.
func (bac *BeyondcorpAppConnection) State() (*beyondcorpAppConnectionState, bool) {
	return bac.state, bac.state != nil
}

// StateMust returns the state for [BeyondcorpAppConnection]. Panics if the state is nil.
func (bac *BeyondcorpAppConnection) StateMust() *beyondcorpAppConnectionState {
	if bac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bac.Type(), bac.LocalName()))
	}
	return bac.state
}

// BeyondcorpAppConnectionArgs contains the configurations for google_beyondcorp_app_connection.
type BeyondcorpAppConnectionArgs struct {
	// Connectors: list of string, optional
	Connectors terra.ListValue[terra.StringValue] `hcl:"connectors,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
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
	// ApplicationEndpoint: required
	ApplicationEndpoint *beyondcorpappconnection.ApplicationEndpoint `hcl:"application_endpoint,block" validate:"required"`
	// Gateway: optional
	Gateway *beyondcorpappconnection.Gateway `hcl:"gateway,block"`
	// Timeouts: optional
	Timeouts *beyondcorpappconnection.Timeouts `hcl:"timeouts,block"`
}
type beyondcorpAppConnectionAttributes struct {
	ref terra.Reference
}

// Connectors returns a reference to field connectors of google_beyondcorp_app_connection.
func (bac beyondcorpAppConnectionAttributes) Connectors() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](bac.ref.Append("connectors"))
}

// DisplayName returns a reference to field display_name of google_beyondcorp_app_connection.
func (bac beyondcorpAppConnectionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(bac.ref.Append("display_name"))
}

// Id returns a reference to field id of google_beyondcorp_app_connection.
func (bac beyondcorpAppConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bac.ref.Append("id"))
}

// Labels returns a reference to field labels of google_beyondcorp_app_connection.
func (bac beyondcorpAppConnectionAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bac.ref.Append("labels"))
}

// Name returns a reference to field name of google_beyondcorp_app_connection.
func (bac beyondcorpAppConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bac.ref.Append("name"))
}

// Project returns a reference to field project of google_beyondcorp_app_connection.
func (bac beyondcorpAppConnectionAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bac.ref.Append("project"))
}

// Region returns a reference to field region of google_beyondcorp_app_connection.
func (bac beyondcorpAppConnectionAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(bac.ref.Append("region"))
}

// Type returns a reference to field type of google_beyondcorp_app_connection.
func (bac beyondcorpAppConnectionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(bac.ref.Append("type"))
}

func (bac beyondcorpAppConnectionAttributes) ApplicationEndpoint() terra.ListValue[beyondcorpappconnection.ApplicationEndpointAttributes] {
	return terra.ReferenceAsList[beyondcorpappconnection.ApplicationEndpointAttributes](bac.ref.Append("application_endpoint"))
}

func (bac beyondcorpAppConnectionAttributes) Gateway() terra.ListValue[beyondcorpappconnection.GatewayAttributes] {
	return terra.ReferenceAsList[beyondcorpappconnection.GatewayAttributes](bac.ref.Append("gateway"))
}

func (bac beyondcorpAppConnectionAttributes) Timeouts() beyondcorpappconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[beyondcorpappconnection.TimeoutsAttributes](bac.ref.Append("timeouts"))
}

type beyondcorpAppConnectionState struct {
	Connectors          []string                                           `json:"connectors"`
	DisplayName         string                                             `json:"display_name"`
	Id                  string                                             `json:"id"`
	Labels              map[string]string                                  `json:"labels"`
	Name                string                                             `json:"name"`
	Project             string                                             `json:"project"`
	Region              string                                             `json:"region"`
	Type                string                                             `json:"type"`
	ApplicationEndpoint []beyondcorpappconnection.ApplicationEndpointState `json:"application_endpoint"`
	Gateway             []beyondcorpappconnection.GatewayState             `json:"gateway"`
	Timeouts            *beyondcorpappconnection.TimeoutsState             `json:"timeouts"`
}
