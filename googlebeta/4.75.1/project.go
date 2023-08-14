// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	project "github.com/golingon/terraproviders/googlebeta/4.75.1/project"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewProject creates a new instance of [Project].
func NewProject(name string, args ProjectArgs) *Project {
	return &Project{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Project)(nil)

// Project represents the Terraform resource google_project.
type Project struct {
	Name      string
	Args      ProjectArgs
	state     *projectState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Project].
func (p *Project) Type() string {
	return "google_project"
}

// LocalName returns the local name for [Project].
func (p *Project) LocalName() string {
	return p.Name
}

// Configuration returns the configuration (args) for [Project].
func (p *Project) Configuration() interface{} {
	return p.Args
}

// DependOn is used for other resources to depend on [Project].
func (p *Project) DependOn() terra.Reference {
	return terra.ReferenceResource(p)
}

// Dependencies returns the list of resources [Project] depends_on.
func (p *Project) Dependencies() terra.Dependencies {
	return p.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Project].
func (p *Project) LifecycleManagement() *terra.Lifecycle {
	return p.Lifecycle
}

// Attributes returns the attributes for [Project].
func (p *Project) Attributes() projectAttributes {
	return projectAttributes{ref: terra.ReferenceResource(p)}
}

// ImportState imports the given attribute values into [Project]'s state.
func (p *Project) ImportState(av io.Reader) error {
	p.state = &projectState{}
	if err := json.NewDecoder(av).Decode(p.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", p.Type(), p.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Project] has state.
func (p *Project) State() (*projectState, bool) {
	return p.state, p.state != nil
}

// StateMust returns the state for [Project]. Panics if the state is nil.
func (p *Project) StateMust() *projectState {
	if p.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", p.Type(), p.LocalName()))
	}
	return p.state
}

// ProjectArgs contains the configurations for google_project.
type ProjectArgs struct {
	// AutoCreateNetwork: bool, optional
	AutoCreateNetwork terra.BoolValue `hcl:"auto_create_network,attr"`
	// BillingAccount: string, optional
	BillingAccount terra.StringValue `hcl:"billing_account,attr"`
	// FolderId: string, optional
	FolderId terra.StringValue `hcl:"folder_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OrgId: string, optional
	OrgId terra.StringValue `hcl:"org_id,attr"`
	// ProjectId: string, required
	ProjectId terra.StringValue `hcl:"project_id,attr" validate:"required"`
	// SkipDelete: bool, optional
	SkipDelete terra.BoolValue `hcl:"skip_delete,attr"`
	// Timeouts: optional
	Timeouts *project.Timeouts `hcl:"timeouts,block"`
}
type projectAttributes struct {
	ref terra.Reference
}

// AutoCreateNetwork returns a reference to field auto_create_network of google_project.
func (p projectAttributes) AutoCreateNetwork() terra.BoolValue {
	return terra.ReferenceAsBool(p.ref.Append("auto_create_network"))
}

// BillingAccount returns a reference to field billing_account of google_project.
func (p projectAttributes) BillingAccount() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("billing_account"))
}

// FolderId returns a reference to field folder_id of google_project.
func (p projectAttributes) FolderId() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("folder_id"))
}

// Id returns a reference to field id of google_project.
func (p projectAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("id"))
}

// Labels returns a reference to field labels of google_project.
func (p projectAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](p.ref.Append("labels"))
}

// Name returns a reference to field name of google_project.
func (p projectAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("name"))
}

// Number returns a reference to field number of google_project.
func (p projectAttributes) Number() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("number"))
}

// OrgId returns a reference to field org_id of google_project.
func (p projectAttributes) OrgId() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("org_id"))
}

// ProjectId returns a reference to field project_id of google_project.
func (p projectAttributes) ProjectId() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("project_id"))
}

// SkipDelete returns a reference to field skip_delete of google_project.
func (p projectAttributes) SkipDelete() terra.BoolValue {
	return terra.ReferenceAsBool(p.ref.Append("skip_delete"))
}

func (p projectAttributes) Timeouts() project.TimeoutsAttributes {
	return terra.ReferenceAsSingle[project.TimeoutsAttributes](p.ref.Append("timeouts"))
}

type projectState struct {
	AutoCreateNetwork bool                   `json:"auto_create_network"`
	BillingAccount    string                 `json:"billing_account"`
	FolderId          string                 `json:"folder_id"`
	Id                string                 `json:"id"`
	Labels            map[string]string      `json:"labels"`
	Name              string                 `json:"name"`
	Number            string                 `json:"number"`
	OrgId             string                 `json:"org_id"`
	ProjectId         string                 `json:"project_id"`
	SkipDelete        bool                   `json:"skip_delete"`
	Timeouts          *project.TimeoutsState `json:"timeouts"`
}