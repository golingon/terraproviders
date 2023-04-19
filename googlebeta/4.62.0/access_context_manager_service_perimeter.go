// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	accesscontextmanagerserviceperimeter "github.com/golingon/terraproviders/googlebeta/4.62.0/accesscontextmanagerserviceperimeter"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAccessContextManagerServicePerimeter creates a new instance of [AccessContextManagerServicePerimeter].
func NewAccessContextManagerServicePerimeter(name string, args AccessContextManagerServicePerimeterArgs) *AccessContextManagerServicePerimeter {
	return &AccessContextManagerServicePerimeter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AccessContextManagerServicePerimeter)(nil)

// AccessContextManagerServicePerimeter represents the Terraform resource google_access_context_manager_service_perimeter.
type AccessContextManagerServicePerimeter struct {
	Name      string
	Args      AccessContextManagerServicePerimeterArgs
	state     *accessContextManagerServicePerimeterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AccessContextManagerServicePerimeter].
func (acmsp *AccessContextManagerServicePerimeter) Type() string {
	return "google_access_context_manager_service_perimeter"
}

// LocalName returns the local name for [AccessContextManagerServicePerimeter].
func (acmsp *AccessContextManagerServicePerimeter) LocalName() string {
	return acmsp.Name
}

// Configuration returns the configuration (args) for [AccessContextManagerServicePerimeter].
func (acmsp *AccessContextManagerServicePerimeter) Configuration() interface{} {
	return acmsp.Args
}

// DependOn is used for other resources to depend on [AccessContextManagerServicePerimeter].
func (acmsp *AccessContextManagerServicePerimeter) DependOn() terra.Reference {
	return terra.ReferenceResource(acmsp)
}

// Dependencies returns the list of resources [AccessContextManagerServicePerimeter] depends_on.
func (acmsp *AccessContextManagerServicePerimeter) Dependencies() terra.Dependencies {
	return acmsp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AccessContextManagerServicePerimeter].
func (acmsp *AccessContextManagerServicePerimeter) LifecycleManagement() *terra.Lifecycle {
	return acmsp.Lifecycle
}

// Attributes returns the attributes for [AccessContextManagerServicePerimeter].
func (acmsp *AccessContextManagerServicePerimeter) Attributes() accessContextManagerServicePerimeterAttributes {
	return accessContextManagerServicePerimeterAttributes{ref: terra.ReferenceResource(acmsp)}
}

// ImportState imports the given attribute values into [AccessContextManagerServicePerimeter]'s state.
func (acmsp *AccessContextManagerServicePerimeter) ImportState(av io.Reader) error {
	acmsp.state = &accessContextManagerServicePerimeterState{}
	if err := json.NewDecoder(av).Decode(acmsp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acmsp.Type(), acmsp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AccessContextManagerServicePerimeter] has state.
func (acmsp *AccessContextManagerServicePerimeter) State() (*accessContextManagerServicePerimeterState, bool) {
	return acmsp.state, acmsp.state != nil
}

// StateMust returns the state for [AccessContextManagerServicePerimeter]. Panics if the state is nil.
func (acmsp *AccessContextManagerServicePerimeter) StateMust() *accessContextManagerServicePerimeterState {
	if acmsp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acmsp.Type(), acmsp.LocalName()))
	}
	return acmsp.state
}

// AccessContextManagerServicePerimeterArgs contains the configurations for google_access_context_manager_service_perimeter.
type AccessContextManagerServicePerimeterArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// PerimeterType: string, optional
	PerimeterType terra.StringValue `hcl:"perimeter_type,attr"`
	// Title: string, required
	Title terra.StringValue `hcl:"title,attr" validate:"required"`
	// UseExplicitDryRunSpec: bool, optional
	UseExplicitDryRunSpec terra.BoolValue `hcl:"use_explicit_dry_run_spec,attr"`
	// Spec: optional
	Spec *accesscontextmanagerserviceperimeter.Spec `hcl:"spec,block"`
	// Status: optional
	Status *accesscontextmanagerserviceperimeter.Status `hcl:"status,block"`
	// Timeouts: optional
	Timeouts *accesscontextmanagerserviceperimeter.Timeouts `hcl:"timeouts,block"`
}
type accessContextManagerServicePerimeterAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_access_context_manager_service_perimeter.
func (acmsp accessContextManagerServicePerimeterAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(acmsp.ref.Append("create_time"))
}

// Description returns a reference to field description of google_access_context_manager_service_perimeter.
func (acmsp accessContextManagerServicePerimeterAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(acmsp.ref.Append("description"))
}

// Id returns a reference to field id of google_access_context_manager_service_perimeter.
func (acmsp accessContextManagerServicePerimeterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acmsp.ref.Append("id"))
}

// Name returns a reference to field name of google_access_context_manager_service_perimeter.
func (acmsp accessContextManagerServicePerimeterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acmsp.ref.Append("name"))
}

// Parent returns a reference to field parent of google_access_context_manager_service_perimeter.
func (acmsp accessContextManagerServicePerimeterAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(acmsp.ref.Append("parent"))
}

// PerimeterType returns a reference to field perimeter_type of google_access_context_manager_service_perimeter.
func (acmsp accessContextManagerServicePerimeterAttributes) PerimeterType() terra.StringValue {
	return terra.ReferenceAsString(acmsp.ref.Append("perimeter_type"))
}

// Title returns a reference to field title of google_access_context_manager_service_perimeter.
func (acmsp accessContextManagerServicePerimeterAttributes) Title() terra.StringValue {
	return terra.ReferenceAsString(acmsp.ref.Append("title"))
}

// UpdateTime returns a reference to field update_time of google_access_context_manager_service_perimeter.
func (acmsp accessContextManagerServicePerimeterAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(acmsp.ref.Append("update_time"))
}

// UseExplicitDryRunSpec returns a reference to field use_explicit_dry_run_spec of google_access_context_manager_service_perimeter.
func (acmsp accessContextManagerServicePerimeterAttributes) UseExplicitDryRunSpec() terra.BoolValue {
	return terra.ReferenceAsBool(acmsp.ref.Append("use_explicit_dry_run_spec"))
}

func (acmsp accessContextManagerServicePerimeterAttributes) Spec() terra.ListValue[accesscontextmanagerserviceperimeter.SpecAttributes] {
	return terra.ReferenceAsList[accesscontextmanagerserviceperimeter.SpecAttributes](acmsp.ref.Append("spec"))
}

func (acmsp accessContextManagerServicePerimeterAttributes) Status() terra.ListValue[accesscontextmanagerserviceperimeter.StatusAttributes] {
	return terra.ReferenceAsList[accesscontextmanagerserviceperimeter.StatusAttributes](acmsp.ref.Append("status"))
}

func (acmsp accessContextManagerServicePerimeterAttributes) Timeouts() accesscontextmanagerserviceperimeter.TimeoutsAttributes {
	return terra.ReferenceAsSingle[accesscontextmanagerserviceperimeter.TimeoutsAttributes](acmsp.ref.Append("timeouts"))
}

type accessContextManagerServicePerimeterState struct {
	CreateTime            string                                              `json:"create_time"`
	Description           string                                              `json:"description"`
	Id                    string                                              `json:"id"`
	Name                  string                                              `json:"name"`
	Parent                string                                              `json:"parent"`
	PerimeterType         string                                              `json:"perimeter_type"`
	Title                 string                                              `json:"title"`
	UpdateTime            string                                              `json:"update_time"`
	UseExplicitDryRunSpec bool                                                `json:"use_explicit_dry_run_spec"`
	Spec                  []accesscontextmanagerserviceperimeter.SpecState    `json:"spec"`
	Status                []accesscontextmanagerserviceperimeter.StatusState  `json:"status"`
	Timeouts              *accesscontextmanagerserviceperimeter.TimeoutsState `json:"timeouts"`
}
