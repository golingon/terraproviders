// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	servicecatalogorganizationsaccess "github.com/golingon/terraproviders/aws/5.11.0/servicecatalogorganizationsaccess"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServicecatalogOrganizationsAccess creates a new instance of [ServicecatalogOrganizationsAccess].
func NewServicecatalogOrganizationsAccess(name string, args ServicecatalogOrganizationsAccessArgs) *ServicecatalogOrganizationsAccess {
	return &ServicecatalogOrganizationsAccess{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServicecatalogOrganizationsAccess)(nil)

// ServicecatalogOrganizationsAccess represents the Terraform resource aws_servicecatalog_organizations_access.
type ServicecatalogOrganizationsAccess struct {
	Name      string
	Args      ServicecatalogOrganizationsAccessArgs
	state     *servicecatalogOrganizationsAccessState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServicecatalogOrganizationsAccess].
func (soa *ServicecatalogOrganizationsAccess) Type() string {
	return "aws_servicecatalog_organizations_access"
}

// LocalName returns the local name for [ServicecatalogOrganizationsAccess].
func (soa *ServicecatalogOrganizationsAccess) LocalName() string {
	return soa.Name
}

// Configuration returns the configuration (args) for [ServicecatalogOrganizationsAccess].
func (soa *ServicecatalogOrganizationsAccess) Configuration() interface{} {
	return soa.Args
}

// DependOn is used for other resources to depend on [ServicecatalogOrganizationsAccess].
func (soa *ServicecatalogOrganizationsAccess) DependOn() terra.Reference {
	return terra.ReferenceResource(soa)
}

// Dependencies returns the list of resources [ServicecatalogOrganizationsAccess] depends_on.
func (soa *ServicecatalogOrganizationsAccess) Dependencies() terra.Dependencies {
	return soa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServicecatalogOrganizationsAccess].
func (soa *ServicecatalogOrganizationsAccess) LifecycleManagement() *terra.Lifecycle {
	return soa.Lifecycle
}

// Attributes returns the attributes for [ServicecatalogOrganizationsAccess].
func (soa *ServicecatalogOrganizationsAccess) Attributes() servicecatalogOrganizationsAccessAttributes {
	return servicecatalogOrganizationsAccessAttributes{ref: terra.ReferenceResource(soa)}
}

// ImportState imports the given attribute values into [ServicecatalogOrganizationsAccess]'s state.
func (soa *ServicecatalogOrganizationsAccess) ImportState(av io.Reader) error {
	soa.state = &servicecatalogOrganizationsAccessState{}
	if err := json.NewDecoder(av).Decode(soa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", soa.Type(), soa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServicecatalogOrganizationsAccess] has state.
func (soa *ServicecatalogOrganizationsAccess) State() (*servicecatalogOrganizationsAccessState, bool) {
	return soa.state, soa.state != nil
}

// StateMust returns the state for [ServicecatalogOrganizationsAccess]. Panics if the state is nil.
func (soa *ServicecatalogOrganizationsAccess) StateMust() *servicecatalogOrganizationsAccessState {
	if soa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", soa.Type(), soa.LocalName()))
	}
	return soa.state
}

// ServicecatalogOrganizationsAccessArgs contains the configurations for aws_servicecatalog_organizations_access.
type ServicecatalogOrganizationsAccessArgs struct {
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *servicecatalogorganizationsaccess.Timeouts `hcl:"timeouts,block"`
}
type servicecatalogOrganizationsAccessAttributes struct {
	ref terra.Reference
}

// Enabled returns a reference to field enabled of aws_servicecatalog_organizations_access.
func (soa servicecatalogOrganizationsAccessAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(soa.ref.Append("enabled"))
}

// Id returns a reference to field id of aws_servicecatalog_organizations_access.
func (soa servicecatalogOrganizationsAccessAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(soa.ref.Append("id"))
}

func (soa servicecatalogOrganizationsAccessAttributes) Timeouts() servicecatalogorganizationsaccess.TimeoutsAttributes {
	return terra.ReferenceAsSingle[servicecatalogorganizationsaccess.TimeoutsAttributes](soa.ref.Append("timeouts"))
}

type servicecatalogOrganizationsAccessState struct {
	Enabled  bool                                             `json:"enabled"`
	Id       string                                           `json:"id"`
	Timeouts *servicecatalogorganizationsaccess.TimeoutsState `json:"timeouts"`
}
