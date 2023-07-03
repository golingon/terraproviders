// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	datastreamstream "github.com/golingon/terraproviders/googlebeta/4.71.0/datastreamstream"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDatastreamStream creates a new instance of [DatastreamStream].
func NewDatastreamStream(name string, args DatastreamStreamArgs) *DatastreamStream {
	return &DatastreamStream{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatastreamStream)(nil)

// DatastreamStream represents the Terraform resource google_datastream_stream.
type DatastreamStream struct {
	Name      string
	Args      DatastreamStreamArgs
	state     *datastreamStreamState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DatastreamStream].
func (ds *DatastreamStream) Type() string {
	return "google_datastream_stream"
}

// LocalName returns the local name for [DatastreamStream].
func (ds *DatastreamStream) LocalName() string {
	return ds.Name
}

// Configuration returns the configuration (args) for [DatastreamStream].
func (ds *DatastreamStream) Configuration() interface{} {
	return ds.Args
}

// DependOn is used for other resources to depend on [DatastreamStream].
func (ds *DatastreamStream) DependOn() terra.Reference {
	return terra.ReferenceResource(ds)
}

// Dependencies returns the list of resources [DatastreamStream] depends_on.
func (ds *DatastreamStream) Dependencies() terra.Dependencies {
	return ds.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DatastreamStream].
func (ds *DatastreamStream) LifecycleManagement() *terra.Lifecycle {
	return ds.Lifecycle
}

// Attributes returns the attributes for [DatastreamStream].
func (ds *DatastreamStream) Attributes() datastreamStreamAttributes {
	return datastreamStreamAttributes{ref: terra.ReferenceResource(ds)}
}

// ImportState imports the given attribute values into [DatastreamStream]'s state.
func (ds *DatastreamStream) ImportState(av io.Reader) error {
	ds.state = &datastreamStreamState{}
	if err := json.NewDecoder(av).Decode(ds.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ds.Type(), ds.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DatastreamStream] has state.
func (ds *DatastreamStream) State() (*datastreamStreamState, bool) {
	return ds.state, ds.state != nil
}

// StateMust returns the state for [DatastreamStream]. Panics if the state is nil.
func (ds *DatastreamStream) StateMust() *datastreamStreamState {
	if ds.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ds.Type(), ds.LocalName()))
	}
	return ds.state
}

// DatastreamStreamArgs contains the configurations for google_datastream_stream.
type DatastreamStreamArgs struct {
	// CustomerManagedEncryptionKey: string, optional
	CustomerManagedEncryptionKey terra.StringValue `hcl:"customer_managed_encryption_key,attr"`
	// DesiredState: string, optional
	DesiredState terra.StringValue `hcl:"desired_state,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// StreamId: string, required
	StreamId terra.StringValue `hcl:"stream_id,attr" validate:"required"`
	// BackfillAll: optional
	BackfillAll *datastreamstream.BackfillAll `hcl:"backfill_all,block"`
	// BackfillNone: optional
	BackfillNone *datastreamstream.BackfillNone `hcl:"backfill_none,block"`
	// DestinationConfig: required
	DestinationConfig *datastreamstream.DestinationConfig `hcl:"destination_config,block" validate:"required"`
	// SourceConfig: required
	SourceConfig *datastreamstream.SourceConfig `hcl:"source_config,block" validate:"required"`
	// Timeouts: optional
	Timeouts *datastreamstream.Timeouts `hcl:"timeouts,block"`
}
type datastreamStreamAttributes struct {
	ref terra.Reference
}

// CustomerManagedEncryptionKey returns a reference to field customer_managed_encryption_key of google_datastream_stream.
func (ds datastreamStreamAttributes) CustomerManagedEncryptionKey() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("customer_managed_encryption_key"))
}

// DesiredState returns a reference to field desired_state of google_datastream_stream.
func (ds datastreamStreamAttributes) DesiredState() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("desired_state"))
}

// DisplayName returns a reference to field display_name of google_datastream_stream.
func (ds datastreamStreamAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("display_name"))
}

// Id returns a reference to field id of google_datastream_stream.
func (ds datastreamStreamAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("id"))
}

// Labels returns a reference to field labels of google_datastream_stream.
func (ds datastreamStreamAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ds.ref.Append("labels"))
}

// Location returns a reference to field location of google_datastream_stream.
func (ds datastreamStreamAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("location"))
}

// Name returns a reference to field name of google_datastream_stream.
func (ds datastreamStreamAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("name"))
}

// Project returns a reference to field project of google_datastream_stream.
func (ds datastreamStreamAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("project"))
}

// State returns a reference to field state of google_datastream_stream.
func (ds datastreamStreamAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("state"))
}

// StreamId returns a reference to field stream_id of google_datastream_stream.
func (ds datastreamStreamAttributes) StreamId() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("stream_id"))
}

func (ds datastreamStreamAttributes) BackfillAll() terra.ListValue[datastreamstream.BackfillAllAttributes] {
	return terra.ReferenceAsList[datastreamstream.BackfillAllAttributes](ds.ref.Append("backfill_all"))
}

func (ds datastreamStreamAttributes) BackfillNone() terra.ListValue[datastreamstream.BackfillNoneAttributes] {
	return terra.ReferenceAsList[datastreamstream.BackfillNoneAttributes](ds.ref.Append("backfill_none"))
}

func (ds datastreamStreamAttributes) DestinationConfig() terra.ListValue[datastreamstream.DestinationConfigAttributes] {
	return terra.ReferenceAsList[datastreamstream.DestinationConfigAttributes](ds.ref.Append("destination_config"))
}

func (ds datastreamStreamAttributes) SourceConfig() terra.ListValue[datastreamstream.SourceConfigAttributes] {
	return terra.ReferenceAsList[datastreamstream.SourceConfigAttributes](ds.ref.Append("source_config"))
}

func (ds datastreamStreamAttributes) Timeouts() datastreamstream.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datastreamstream.TimeoutsAttributes](ds.ref.Append("timeouts"))
}

type datastreamStreamState struct {
	CustomerManagedEncryptionKey string                                    `json:"customer_managed_encryption_key"`
	DesiredState                 string                                    `json:"desired_state"`
	DisplayName                  string                                    `json:"display_name"`
	Id                           string                                    `json:"id"`
	Labels                       map[string]string                         `json:"labels"`
	Location                     string                                    `json:"location"`
	Name                         string                                    `json:"name"`
	Project                      string                                    `json:"project"`
	State                        string                                    `json:"state"`
	StreamId                     string                                    `json:"stream_id"`
	BackfillAll                  []datastreamstream.BackfillAllState       `json:"backfill_all"`
	BackfillNone                 []datastreamstream.BackfillNoneState      `json:"backfill_none"`
	DestinationConfig            []datastreamstream.DestinationConfigState `json:"destination_config"`
	SourceConfig                 []datastreamstream.SourceConfigState      `json:"source_config"`
	Timeouts                     *datastreamstream.TimeoutsState           `json:"timeouts"`
}
