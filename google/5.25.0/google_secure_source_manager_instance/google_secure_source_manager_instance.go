// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_secure_source_manager_instance

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_secure_source_manager_instance.
type Resource struct {
	Name      string
	Args      Args
	state     *googleSecureSourceManagerInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gssmi *Resource) Type() string {
	return "google_secure_source_manager_instance"
}

// LocalName returns the local name for [Resource].
func (gssmi *Resource) LocalName() string {
	return gssmi.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gssmi *Resource) Configuration() interface{} {
	return gssmi.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gssmi *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gssmi)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gssmi *Resource) Dependencies() terra.Dependencies {
	return gssmi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gssmi *Resource) LifecycleManagement() *terra.Lifecycle {
	return gssmi.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gssmi *Resource) Attributes() googleSecureSourceManagerInstanceAttributes {
	return googleSecureSourceManagerInstanceAttributes{ref: terra.ReferenceResource(gssmi)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gssmi *Resource) ImportState(state io.Reader) error {
	gssmi.state = &googleSecureSourceManagerInstanceState{}
	if err := json.NewDecoder(state).Decode(gssmi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gssmi.Type(), gssmi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gssmi *Resource) State() (*googleSecureSourceManagerInstanceState, bool) {
	return gssmi.state, gssmi.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gssmi *Resource) StateMust() *googleSecureSourceManagerInstanceState {
	if gssmi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gssmi.Type(), gssmi.LocalName()))
	}
	return gssmi.state
}

// Args contains the configurations for google_secure_source_manager_instance.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// KmsKey: string, optional
	KmsKey terra.StringValue `hcl:"kms_key,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// PrivateConfig: optional
	PrivateConfig *PrivateConfig `hcl:"private_config,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleSecureSourceManagerInstanceAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_secure_source_manager_instance.
func (gssmi googleSecureSourceManagerInstanceAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(gssmi.ref.Append("create_time"))
}

// EffectiveLabels returns a reference to field effective_labels of google_secure_source_manager_instance.
func (gssmi googleSecureSourceManagerInstanceAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gssmi.ref.Append("effective_labels"))
}

// Id returns a reference to field id of google_secure_source_manager_instance.
func (gssmi googleSecureSourceManagerInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gssmi.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of google_secure_source_manager_instance.
func (gssmi googleSecureSourceManagerInstanceAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(gssmi.ref.Append("instance_id"))
}

// KmsKey returns a reference to field kms_key of google_secure_source_manager_instance.
func (gssmi googleSecureSourceManagerInstanceAttributes) KmsKey() terra.StringValue {
	return terra.ReferenceAsString(gssmi.ref.Append("kms_key"))
}

// Labels returns a reference to field labels of google_secure_source_manager_instance.
func (gssmi googleSecureSourceManagerInstanceAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gssmi.ref.Append("labels"))
}

// Location returns a reference to field location of google_secure_source_manager_instance.
func (gssmi googleSecureSourceManagerInstanceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gssmi.ref.Append("location"))
}

// Name returns a reference to field name of google_secure_source_manager_instance.
func (gssmi googleSecureSourceManagerInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gssmi.ref.Append("name"))
}

// Project returns a reference to field project of google_secure_source_manager_instance.
func (gssmi googleSecureSourceManagerInstanceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gssmi.ref.Append("project"))
}

// State returns a reference to field state of google_secure_source_manager_instance.
func (gssmi googleSecureSourceManagerInstanceAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(gssmi.ref.Append("state"))
}

// StateNote returns a reference to field state_note of google_secure_source_manager_instance.
func (gssmi googleSecureSourceManagerInstanceAttributes) StateNote() terra.StringValue {
	return terra.ReferenceAsString(gssmi.ref.Append("state_note"))
}

// TerraformLabels returns a reference to field terraform_labels of google_secure_source_manager_instance.
func (gssmi googleSecureSourceManagerInstanceAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gssmi.ref.Append("terraform_labels"))
}

// UpdateTime returns a reference to field update_time of google_secure_source_manager_instance.
func (gssmi googleSecureSourceManagerInstanceAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(gssmi.ref.Append("update_time"))
}

func (gssmi googleSecureSourceManagerInstanceAttributes) HostConfig() terra.ListValue[HostConfigAttributes] {
	return terra.ReferenceAsList[HostConfigAttributes](gssmi.ref.Append("host_config"))
}

func (gssmi googleSecureSourceManagerInstanceAttributes) PrivateConfig() terra.ListValue[PrivateConfigAttributes] {
	return terra.ReferenceAsList[PrivateConfigAttributes](gssmi.ref.Append("private_config"))
}

func (gssmi googleSecureSourceManagerInstanceAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gssmi.ref.Append("timeouts"))
}

type googleSecureSourceManagerInstanceState struct {
	CreateTime      string               `json:"create_time"`
	EffectiveLabels map[string]string    `json:"effective_labels"`
	Id              string               `json:"id"`
	InstanceId      string               `json:"instance_id"`
	KmsKey          string               `json:"kms_key"`
	Labels          map[string]string    `json:"labels"`
	Location        string               `json:"location"`
	Name            string               `json:"name"`
	Project         string               `json:"project"`
	State           string               `json:"state"`
	StateNote       string               `json:"state_note"`
	TerraformLabels map[string]string    `json:"terraform_labels"`
	UpdateTime      string               `json:"update_time"`
	HostConfig      []HostConfigState    `json:"host_config"`
	PrivateConfig   []PrivateConfigState `json:"private_config"`
	Timeouts        *TimeoutsState       `json:"timeouts"`
}
