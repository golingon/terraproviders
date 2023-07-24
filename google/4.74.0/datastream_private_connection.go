// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	datastreamprivateconnection "github.com/golingon/terraproviders/google/4.74.0/datastreamprivateconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDatastreamPrivateConnection creates a new instance of [DatastreamPrivateConnection].
func NewDatastreamPrivateConnection(name string, args DatastreamPrivateConnectionArgs) *DatastreamPrivateConnection {
	return &DatastreamPrivateConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatastreamPrivateConnection)(nil)

// DatastreamPrivateConnection represents the Terraform resource google_datastream_private_connection.
type DatastreamPrivateConnection struct {
	Name      string
	Args      DatastreamPrivateConnectionArgs
	state     *datastreamPrivateConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DatastreamPrivateConnection].
func (dpc *DatastreamPrivateConnection) Type() string {
	return "google_datastream_private_connection"
}

// LocalName returns the local name for [DatastreamPrivateConnection].
func (dpc *DatastreamPrivateConnection) LocalName() string {
	return dpc.Name
}

// Configuration returns the configuration (args) for [DatastreamPrivateConnection].
func (dpc *DatastreamPrivateConnection) Configuration() interface{} {
	return dpc.Args
}

// DependOn is used for other resources to depend on [DatastreamPrivateConnection].
func (dpc *DatastreamPrivateConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(dpc)
}

// Dependencies returns the list of resources [DatastreamPrivateConnection] depends_on.
func (dpc *DatastreamPrivateConnection) Dependencies() terra.Dependencies {
	return dpc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DatastreamPrivateConnection].
func (dpc *DatastreamPrivateConnection) LifecycleManagement() *terra.Lifecycle {
	return dpc.Lifecycle
}

// Attributes returns the attributes for [DatastreamPrivateConnection].
func (dpc *DatastreamPrivateConnection) Attributes() datastreamPrivateConnectionAttributes {
	return datastreamPrivateConnectionAttributes{ref: terra.ReferenceResource(dpc)}
}

// ImportState imports the given attribute values into [DatastreamPrivateConnection]'s state.
func (dpc *DatastreamPrivateConnection) ImportState(av io.Reader) error {
	dpc.state = &datastreamPrivateConnectionState{}
	if err := json.NewDecoder(av).Decode(dpc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dpc.Type(), dpc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DatastreamPrivateConnection] has state.
func (dpc *DatastreamPrivateConnection) State() (*datastreamPrivateConnectionState, bool) {
	return dpc.state, dpc.state != nil
}

// StateMust returns the state for [DatastreamPrivateConnection]. Panics if the state is nil.
func (dpc *DatastreamPrivateConnection) StateMust() *datastreamPrivateConnectionState {
	if dpc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dpc.Type(), dpc.LocalName()))
	}
	return dpc.state
}

// DatastreamPrivateConnectionArgs contains the configurations for google_datastream_private_connection.
type DatastreamPrivateConnectionArgs struct {
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// PrivateConnectionId: string, required
	PrivateConnectionId terra.StringValue `hcl:"private_connection_id,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Error: min=0
	Error []datastreamprivateconnection.Error `hcl:"error,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datastreamprivateconnection.Timeouts `hcl:"timeouts,block"`
	// VpcPeeringConfig: required
	VpcPeeringConfig *datastreamprivateconnection.VpcPeeringConfig `hcl:"vpc_peering_config,block" validate:"required"`
}
type datastreamPrivateConnectionAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_datastream_private_connection.
func (dpc datastreamPrivateConnectionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dpc.ref.Append("display_name"))
}

// Id returns a reference to field id of google_datastream_private_connection.
func (dpc datastreamPrivateConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dpc.ref.Append("id"))
}

// Labels returns a reference to field labels of google_datastream_private_connection.
func (dpc datastreamPrivateConnectionAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dpc.ref.Append("labels"))
}

// Location returns a reference to field location of google_datastream_private_connection.
func (dpc datastreamPrivateConnectionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dpc.ref.Append("location"))
}

// Name returns a reference to field name of google_datastream_private_connection.
func (dpc datastreamPrivateConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dpc.ref.Append("name"))
}

// PrivateConnectionId returns a reference to field private_connection_id of google_datastream_private_connection.
func (dpc datastreamPrivateConnectionAttributes) PrivateConnectionId() terra.StringValue {
	return terra.ReferenceAsString(dpc.ref.Append("private_connection_id"))
}

// Project returns a reference to field project of google_datastream_private_connection.
func (dpc datastreamPrivateConnectionAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dpc.ref.Append("project"))
}

// State returns a reference to field state of google_datastream_private_connection.
func (dpc datastreamPrivateConnectionAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(dpc.ref.Append("state"))
}

func (dpc datastreamPrivateConnectionAttributes) Error() terra.ListValue[datastreamprivateconnection.ErrorAttributes] {
	return terra.ReferenceAsList[datastreamprivateconnection.ErrorAttributes](dpc.ref.Append("error"))
}

func (dpc datastreamPrivateConnectionAttributes) Timeouts() datastreamprivateconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datastreamprivateconnection.TimeoutsAttributes](dpc.ref.Append("timeouts"))
}

func (dpc datastreamPrivateConnectionAttributes) VpcPeeringConfig() terra.ListValue[datastreamprivateconnection.VpcPeeringConfigAttributes] {
	return terra.ReferenceAsList[datastreamprivateconnection.VpcPeeringConfigAttributes](dpc.ref.Append("vpc_peering_config"))
}

type datastreamPrivateConnectionState struct {
	DisplayName         string                                              `json:"display_name"`
	Id                  string                                              `json:"id"`
	Labels              map[string]string                                   `json:"labels"`
	Location            string                                              `json:"location"`
	Name                string                                              `json:"name"`
	PrivateConnectionId string                                              `json:"private_connection_id"`
	Project             string                                              `json:"project"`
	State               string                                              `json:"state"`
	Error               []datastreamprivateconnection.ErrorState            `json:"error"`
	Timeouts            *datastreamprivateconnection.TimeoutsState          `json:"timeouts"`
	VpcPeeringConfig    []datastreamprivateconnection.VpcPeeringConfigState `json:"vpc_peering_config"`
}
