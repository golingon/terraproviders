// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	pubsubschema "github.com/golingon/terraproviders/googlebeta/4.62.0/pubsubschema"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPubsubSchema creates a new instance of [PubsubSchema].
func NewPubsubSchema(name string, args PubsubSchemaArgs) *PubsubSchema {
	return &PubsubSchema{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PubsubSchema)(nil)

// PubsubSchema represents the Terraform resource google_pubsub_schema.
type PubsubSchema struct {
	Name      string
	Args      PubsubSchemaArgs
	state     *pubsubSchemaState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PubsubSchema].
func (ps *PubsubSchema) Type() string {
	return "google_pubsub_schema"
}

// LocalName returns the local name for [PubsubSchema].
func (ps *PubsubSchema) LocalName() string {
	return ps.Name
}

// Configuration returns the configuration (args) for [PubsubSchema].
func (ps *PubsubSchema) Configuration() interface{} {
	return ps.Args
}

// DependOn is used for other resources to depend on [PubsubSchema].
func (ps *PubsubSchema) DependOn() terra.Reference {
	return terra.ReferenceResource(ps)
}

// Dependencies returns the list of resources [PubsubSchema] depends_on.
func (ps *PubsubSchema) Dependencies() terra.Dependencies {
	return ps.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PubsubSchema].
func (ps *PubsubSchema) LifecycleManagement() *terra.Lifecycle {
	return ps.Lifecycle
}

// Attributes returns the attributes for [PubsubSchema].
func (ps *PubsubSchema) Attributes() pubsubSchemaAttributes {
	return pubsubSchemaAttributes{ref: terra.ReferenceResource(ps)}
}

// ImportState imports the given attribute values into [PubsubSchema]'s state.
func (ps *PubsubSchema) ImportState(av io.Reader) error {
	ps.state = &pubsubSchemaState{}
	if err := json.NewDecoder(av).Decode(ps.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ps.Type(), ps.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PubsubSchema] has state.
func (ps *PubsubSchema) State() (*pubsubSchemaState, bool) {
	return ps.state, ps.state != nil
}

// StateMust returns the state for [PubsubSchema]. Panics if the state is nil.
func (ps *PubsubSchema) StateMust() *pubsubSchemaState {
	if ps.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ps.Type(), ps.LocalName()))
	}
	return ps.state
}

// PubsubSchemaArgs contains the configurations for google_pubsub_schema.
type PubsubSchemaArgs struct {
	// Definition: string, optional
	Definition terra.StringValue `hcl:"definition,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Timeouts: optional
	Timeouts *pubsubschema.Timeouts `hcl:"timeouts,block"`
}
type pubsubSchemaAttributes struct {
	ref terra.Reference
}

// Definition returns a reference to field definition of google_pubsub_schema.
func (ps pubsubSchemaAttributes) Definition() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("definition"))
}

// Id returns a reference to field id of google_pubsub_schema.
func (ps pubsubSchemaAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("id"))
}

// Name returns a reference to field name of google_pubsub_schema.
func (ps pubsubSchemaAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("name"))
}

// Project returns a reference to field project of google_pubsub_schema.
func (ps pubsubSchemaAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("project"))
}

// Type returns a reference to field type of google_pubsub_schema.
func (ps pubsubSchemaAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ps.ref.Append("type"))
}

func (ps pubsubSchemaAttributes) Timeouts() pubsubschema.TimeoutsAttributes {
	return terra.ReferenceAsSingle[pubsubschema.TimeoutsAttributes](ps.ref.Append("timeouts"))
}

type pubsubSchemaState struct {
	Definition string                      `json:"definition"`
	Id         string                      `json:"id"`
	Name       string                      `json:"name"`
	Project    string                      `json:"project"`
	Type       string                      `json:"type"`
	Timeouts   *pubsubschema.TimeoutsState `json:"timeouts"`
}
