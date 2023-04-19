// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	firebasedatabaseinstance "github.com/golingon/terraproviders/googlebeta/4.62.0/firebasedatabaseinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFirebaseDatabaseInstance creates a new instance of [FirebaseDatabaseInstance].
func NewFirebaseDatabaseInstance(name string, args FirebaseDatabaseInstanceArgs) *FirebaseDatabaseInstance {
	return &FirebaseDatabaseInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FirebaseDatabaseInstance)(nil)

// FirebaseDatabaseInstance represents the Terraform resource google_firebase_database_instance.
type FirebaseDatabaseInstance struct {
	Name      string
	Args      FirebaseDatabaseInstanceArgs
	state     *firebaseDatabaseInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FirebaseDatabaseInstance].
func (fdi *FirebaseDatabaseInstance) Type() string {
	return "google_firebase_database_instance"
}

// LocalName returns the local name for [FirebaseDatabaseInstance].
func (fdi *FirebaseDatabaseInstance) LocalName() string {
	return fdi.Name
}

// Configuration returns the configuration (args) for [FirebaseDatabaseInstance].
func (fdi *FirebaseDatabaseInstance) Configuration() interface{} {
	return fdi.Args
}

// DependOn is used for other resources to depend on [FirebaseDatabaseInstance].
func (fdi *FirebaseDatabaseInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(fdi)
}

// Dependencies returns the list of resources [FirebaseDatabaseInstance] depends_on.
func (fdi *FirebaseDatabaseInstance) Dependencies() terra.Dependencies {
	return fdi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FirebaseDatabaseInstance].
func (fdi *FirebaseDatabaseInstance) LifecycleManagement() *terra.Lifecycle {
	return fdi.Lifecycle
}

// Attributes returns the attributes for [FirebaseDatabaseInstance].
func (fdi *FirebaseDatabaseInstance) Attributes() firebaseDatabaseInstanceAttributes {
	return firebaseDatabaseInstanceAttributes{ref: terra.ReferenceResource(fdi)}
}

// ImportState imports the given attribute values into [FirebaseDatabaseInstance]'s state.
func (fdi *FirebaseDatabaseInstance) ImportState(av io.Reader) error {
	fdi.state = &firebaseDatabaseInstanceState{}
	if err := json.NewDecoder(av).Decode(fdi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fdi.Type(), fdi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FirebaseDatabaseInstance] has state.
func (fdi *FirebaseDatabaseInstance) State() (*firebaseDatabaseInstanceState, bool) {
	return fdi.state, fdi.state != nil
}

// StateMust returns the state for [FirebaseDatabaseInstance]. Panics if the state is nil.
func (fdi *FirebaseDatabaseInstance) StateMust() *firebaseDatabaseInstanceState {
	if fdi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fdi.Type(), fdi.LocalName()))
	}
	return fdi.state
}

// FirebaseDatabaseInstanceArgs contains the configurations for google_firebase_database_instance.
type FirebaseDatabaseInstanceArgs struct {
	// DesiredState: string, optional
	DesiredState terra.StringValue `hcl:"desired_state,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, required
	Region terra.StringValue `hcl:"region,attr" validate:"required"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Timeouts: optional
	Timeouts *firebasedatabaseinstance.Timeouts `hcl:"timeouts,block"`
}
type firebaseDatabaseInstanceAttributes struct {
	ref terra.Reference
}

// DatabaseUrl returns a reference to field database_url of google_firebase_database_instance.
func (fdi firebaseDatabaseInstanceAttributes) DatabaseUrl() terra.StringValue {
	return terra.ReferenceAsString(fdi.ref.Append("database_url"))
}

// DesiredState returns a reference to field desired_state of google_firebase_database_instance.
func (fdi firebaseDatabaseInstanceAttributes) DesiredState() terra.StringValue {
	return terra.ReferenceAsString(fdi.ref.Append("desired_state"))
}

// Id returns a reference to field id of google_firebase_database_instance.
func (fdi firebaseDatabaseInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fdi.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of google_firebase_database_instance.
func (fdi firebaseDatabaseInstanceAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(fdi.ref.Append("instance_id"))
}

// Name returns a reference to field name of google_firebase_database_instance.
func (fdi firebaseDatabaseInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fdi.ref.Append("name"))
}

// Project returns a reference to field project of google_firebase_database_instance.
func (fdi firebaseDatabaseInstanceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(fdi.ref.Append("project"))
}

// Region returns a reference to field region of google_firebase_database_instance.
func (fdi firebaseDatabaseInstanceAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(fdi.ref.Append("region"))
}

// State returns a reference to field state of google_firebase_database_instance.
func (fdi firebaseDatabaseInstanceAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(fdi.ref.Append("state"))
}

// Type returns a reference to field type of google_firebase_database_instance.
func (fdi firebaseDatabaseInstanceAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(fdi.ref.Append("type"))
}

func (fdi firebaseDatabaseInstanceAttributes) Timeouts() firebasedatabaseinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[firebasedatabaseinstance.TimeoutsAttributes](fdi.ref.Append("timeouts"))
}

type firebaseDatabaseInstanceState struct {
	DatabaseUrl  string                                  `json:"database_url"`
	DesiredState string                                  `json:"desired_state"`
	Id           string                                  `json:"id"`
	InstanceId   string                                  `json:"instance_id"`
	Name         string                                  `json:"name"`
	Project      string                                  `json:"project"`
	Region       string                                  `json:"region"`
	State        string                                  `json:"state"`
	Type         string                                  `json:"type"`
	Timeouts     *firebasedatabaseinstance.TimeoutsState `json:"timeouts"`
}
