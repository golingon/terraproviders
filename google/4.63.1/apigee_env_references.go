// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	apigeeenvreferences "github.com/golingon/terraproviders/google/4.63.1/apigeeenvreferences"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApigeeEnvReferences creates a new instance of [ApigeeEnvReferences].
func NewApigeeEnvReferences(name string, args ApigeeEnvReferencesArgs) *ApigeeEnvReferences {
	return &ApigeeEnvReferences{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApigeeEnvReferences)(nil)

// ApigeeEnvReferences represents the Terraform resource google_apigee_env_references.
type ApigeeEnvReferences struct {
	Name      string
	Args      ApigeeEnvReferencesArgs
	state     *apigeeEnvReferencesState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApigeeEnvReferences].
func (aer *ApigeeEnvReferences) Type() string {
	return "google_apigee_env_references"
}

// LocalName returns the local name for [ApigeeEnvReferences].
func (aer *ApigeeEnvReferences) LocalName() string {
	return aer.Name
}

// Configuration returns the configuration (args) for [ApigeeEnvReferences].
func (aer *ApigeeEnvReferences) Configuration() interface{} {
	return aer.Args
}

// DependOn is used for other resources to depend on [ApigeeEnvReferences].
func (aer *ApigeeEnvReferences) DependOn() terra.Reference {
	return terra.ReferenceResource(aer)
}

// Dependencies returns the list of resources [ApigeeEnvReferences] depends_on.
func (aer *ApigeeEnvReferences) Dependencies() terra.Dependencies {
	return aer.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApigeeEnvReferences].
func (aer *ApigeeEnvReferences) LifecycleManagement() *terra.Lifecycle {
	return aer.Lifecycle
}

// Attributes returns the attributes for [ApigeeEnvReferences].
func (aer *ApigeeEnvReferences) Attributes() apigeeEnvReferencesAttributes {
	return apigeeEnvReferencesAttributes{ref: terra.ReferenceResource(aer)}
}

// ImportState imports the given attribute values into [ApigeeEnvReferences]'s state.
func (aer *ApigeeEnvReferences) ImportState(av io.Reader) error {
	aer.state = &apigeeEnvReferencesState{}
	if err := json.NewDecoder(av).Decode(aer.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aer.Type(), aer.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApigeeEnvReferences] has state.
func (aer *ApigeeEnvReferences) State() (*apigeeEnvReferencesState, bool) {
	return aer.state, aer.state != nil
}

// StateMust returns the state for [ApigeeEnvReferences]. Panics if the state is nil.
func (aer *ApigeeEnvReferences) StateMust() *apigeeEnvReferencesState {
	if aer.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aer.Type(), aer.LocalName()))
	}
	return aer.state
}

// ApigeeEnvReferencesArgs contains the configurations for google_apigee_env_references.
type ApigeeEnvReferencesArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EnvId: string, required
	EnvId terra.StringValue `hcl:"env_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Refers: string, required
	Refers terra.StringValue `hcl:"refers,attr" validate:"required"`
	// ResourceType: string, required
	ResourceType terra.StringValue `hcl:"resource_type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *apigeeenvreferences.Timeouts `hcl:"timeouts,block"`
}
type apigeeEnvReferencesAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_apigee_env_references.
func (aer apigeeEnvReferencesAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aer.ref.Append("description"))
}

// EnvId returns a reference to field env_id of google_apigee_env_references.
func (aer apigeeEnvReferencesAttributes) EnvId() terra.StringValue {
	return terra.ReferenceAsString(aer.ref.Append("env_id"))
}

// Id returns a reference to field id of google_apigee_env_references.
func (aer apigeeEnvReferencesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aer.ref.Append("id"))
}

// Name returns a reference to field name of google_apigee_env_references.
func (aer apigeeEnvReferencesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aer.ref.Append("name"))
}

// Refers returns a reference to field refers of google_apigee_env_references.
func (aer apigeeEnvReferencesAttributes) Refers() terra.StringValue {
	return terra.ReferenceAsString(aer.ref.Append("refers"))
}

// ResourceType returns a reference to field resource_type of google_apigee_env_references.
func (aer apigeeEnvReferencesAttributes) ResourceType() terra.StringValue {
	return terra.ReferenceAsString(aer.ref.Append("resource_type"))
}

func (aer apigeeEnvReferencesAttributes) Timeouts() apigeeenvreferences.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apigeeenvreferences.TimeoutsAttributes](aer.ref.Append("timeouts"))
}

type apigeeEnvReferencesState struct {
	Description  string                             `json:"description"`
	EnvId        string                             `json:"env_id"`
	Id           string                             `json:"id"`
	Name         string                             `json:"name"`
	Refers       string                             `json:"refers"`
	ResourceType string                             `json:"resource_type"`
	Timeouts     *apigeeenvreferences.TimeoutsState `json:"timeouts"`
}
