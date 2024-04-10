// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	apphubapplication "github.com/golingon/terraproviders/googlebeta/5.24.0/apphubapplication"
	"io"
)

// NewApphubApplication creates a new instance of [ApphubApplication].
func NewApphubApplication(name string, args ApphubApplicationArgs) *ApphubApplication {
	return &ApphubApplication{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApphubApplication)(nil)

// ApphubApplication represents the Terraform resource google_apphub_application.
type ApphubApplication struct {
	Name      string
	Args      ApphubApplicationArgs
	state     *apphubApplicationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApphubApplication].
func (aa *ApphubApplication) Type() string {
	return "google_apphub_application"
}

// LocalName returns the local name for [ApphubApplication].
func (aa *ApphubApplication) LocalName() string {
	return aa.Name
}

// Configuration returns the configuration (args) for [ApphubApplication].
func (aa *ApphubApplication) Configuration() interface{} {
	return aa.Args
}

// DependOn is used for other resources to depend on [ApphubApplication].
func (aa *ApphubApplication) DependOn() terra.Reference {
	return terra.ReferenceResource(aa)
}

// Dependencies returns the list of resources [ApphubApplication] depends_on.
func (aa *ApphubApplication) Dependencies() terra.Dependencies {
	return aa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApphubApplication].
func (aa *ApphubApplication) LifecycleManagement() *terra.Lifecycle {
	return aa.Lifecycle
}

// Attributes returns the attributes for [ApphubApplication].
func (aa *ApphubApplication) Attributes() apphubApplicationAttributes {
	return apphubApplicationAttributes{ref: terra.ReferenceResource(aa)}
}

// ImportState imports the given attribute values into [ApphubApplication]'s state.
func (aa *ApphubApplication) ImportState(av io.Reader) error {
	aa.state = &apphubApplicationState{}
	if err := json.NewDecoder(av).Decode(aa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aa.Type(), aa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApphubApplication] has state.
func (aa *ApphubApplication) State() (*apphubApplicationState, bool) {
	return aa.state, aa.state != nil
}

// StateMust returns the state for [ApphubApplication]. Panics if the state is nil.
func (aa *ApphubApplication) StateMust() *apphubApplicationState {
	if aa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aa.Type(), aa.LocalName()))
	}
	return aa.state
}

// ApphubApplicationArgs contains the configurations for google_apphub_application.
type ApphubApplicationArgs struct {
	// ApplicationId: string, required
	ApplicationId terra.StringValue `hcl:"application_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Attributes: optional
	Attributes *apphubapplication.Attributes `hcl:"attributes,block"`
	// Scope: required
	Scope *apphubapplication.Scope `hcl:"scope,block" validate:"required"`
	// Timeouts: optional
	Timeouts *apphubapplication.Timeouts `hcl:"timeouts,block"`
}
type apphubApplicationAttributes struct {
	ref terra.Reference
}

// ApplicationId returns a reference to field application_id of google_apphub_application.
func (aa apphubApplicationAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("application_id"))
}

// CreateTime returns a reference to field create_time of google_apphub_application.
func (aa apphubApplicationAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("create_time"))
}

// Description returns a reference to field description of google_apphub_application.
func (aa apphubApplicationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_apphub_application.
func (aa apphubApplicationAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("display_name"))
}

// Id returns a reference to field id of google_apphub_application.
func (aa apphubApplicationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("id"))
}

// Location returns a reference to field location of google_apphub_application.
func (aa apphubApplicationAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("location"))
}

// Name returns a reference to field name of google_apphub_application.
func (aa apphubApplicationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("name"))
}

// Project returns a reference to field project of google_apphub_application.
func (aa apphubApplicationAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("project"))
}

// State returns a reference to field state of google_apphub_application.
func (aa apphubApplicationAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("state"))
}

// Uid returns a reference to field uid of google_apphub_application.
func (aa apphubApplicationAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_apphub_application.
func (aa apphubApplicationAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("update_time"))
}

func (aa apphubApplicationAttributes) Attributes() terra.ListValue[apphubapplication.AttributesAttributes] {
	return terra.ReferenceAsList[apphubapplication.AttributesAttributes](aa.ref.Append("attributes"))
}

func (aa apphubApplicationAttributes) Scope() terra.ListValue[apphubapplication.ScopeAttributes] {
	return terra.ReferenceAsList[apphubapplication.ScopeAttributes](aa.ref.Append("scope"))
}

func (aa apphubApplicationAttributes) Timeouts() apphubapplication.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apphubapplication.TimeoutsAttributes](aa.ref.Append("timeouts"))
}

type apphubApplicationState struct {
	ApplicationId string                              `json:"application_id"`
	CreateTime    string                              `json:"create_time"`
	Description   string                              `json:"description"`
	DisplayName   string                              `json:"display_name"`
	Id            string                              `json:"id"`
	Location      string                              `json:"location"`
	Name          string                              `json:"name"`
	Project       string                              `json:"project"`
	State         string                              `json:"state"`
	Uid           string                              `json:"uid"`
	UpdateTime    string                              `json:"update_time"`
	Attributes    []apphubapplication.AttributesState `json:"attributes"`
	Scope         []apphubapplication.ScopeState      `json:"scope"`
	Timeouts      *apphubapplication.TimeoutsState    `json:"timeouts"`
}
