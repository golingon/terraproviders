// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	grafanaroleassociation "github.com/golingon/terraproviders/aws/5.6.2/grafanaroleassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGrafanaRoleAssociation creates a new instance of [GrafanaRoleAssociation].
func NewGrafanaRoleAssociation(name string, args GrafanaRoleAssociationArgs) *GrafanaRoleAssociation {
	return &GrafanaRoleAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GrafanaRoleAssociation)(nil)

// GrafanaRoleAssociation represents the Terraform resource aws_grafana_role_association.
type GrafanaRoleAssociation struct {
	Name      string
	Args      GrafanaRoleAssociationArgs
	state     *grafanaRoleAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GrafanaRoleAssociation].
func (gra *GrafanaRoleAssociation) Type() string {
	return "aws_grafana_role_association"
}

// LocalName returns the local name for [GrafanaRoleAssociation].
func (gra *GrafanaRoleAssociation) LocalName() string {
	return gra.Name
}

// Configuration returns the configuration (args) for [GrafanaRoleAssociation].
func (gra *GrafanaRoleAssociation) Configuration() interface{} {
	return gra.Args
}

// DependOn is used for other resources to depend on [GrafanaRoleAssociation].
func (gra *GrafanaRoleAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(gra)
}

// Dependencies returns the list of resources [GrafanaRoleAssociation] depends_on.
func (gra *GrafanaRoleAssociation) Dependencies() terra.Dependencies {
	return gra.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GrafanaRoleAssociation].
func (gra *GrafanaRoleAssociation) LifecycleManagement() *terra.Lifecycle {
	return gra.Lifecycle
}

// Attributes returns the attributes for [GrafanaRoleAssociation].
func (gra *GrafanaRoleAssociation) Attributes() grafanaRoleAssociationAttributes {
	return grafanaRoleAssociationAttributes{ref: terra.ReferenceResource(gra)}
}

// ImportState imports the given attribute values into [GrafanaRoleAssociation]'s state.
func (gra *GrafanaRoleAssociation) ImportState(av io.Reader) error {
	gra.state = &grafanaRoleAssociationState{}
	if err := json.NewDecoder(av).Decode(gra.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gra.Type(), gra.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GrafanaRoleAssociation] has state.
func (gra *GrafanaRoleAssociation) State() (*grafanaRoleAssociationState, bool) {
	return gra.state, gra.state != nil
}

// StateMust returns the state for [GrafanaRoleAssociation]. Panics if the state is nil.
func (gra *GrafanaRoleAssociation) StateMust() *grafanaRoleAssociationState {
	if gra.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gra.Type(), gra.LocalName()))
	}
	return gra.state
}

// GrafanaRoleAssociationArgs contains the configurations for aws_grafana_role_association.
type GrafanaRoleAssociationArgs struct {
	// GroupIds: set of string, optional
	GroupIds terra.SetValue[terra.StringValue] `hcl:"group_ids,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// UserIds: set of string, optional
	UserIds terra.SetValue[terra.StringValue] `hcl:"user_ids,attr"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *grafanaroleassociation.Timeouts `hcl:"timeouts,block"`
}
type grafanaRoleAssociationAttributes struct {
	ref terra.Reference
}

// GroupIds returns a reference to field group_ids of aws_grafana_role_association.
func (gra grafanaRoleAssociationAttributes) GroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gra.ref.Append("group_ids"))
}

// Id returns a reference to field id of aws_grafana_role_association.
func (gra grafanaRoleAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gra.ref.Append("id"))
}

// Role returns a reference to field role of aws_grafana_role_association.
func (gra grafanaRoleAssociationAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(gra.ref.Append("role"))
}

// UserIds returns a reference to field user_ids of aws_grafana_role_association.
func (gra grafanaRoleAssociationAttributes) UserIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gra.ref.Append("user_ids"))
}

// WorkspaceId returns a reference to field workspace_id of aws_grafana_role_association.
func (gra grafanaRoleAssociationAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(gra.ref.Append("workspace_id"))
}

func (gra grafanaRoleAssociationAttributes) Timeouts() grafanaroleassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[grafanaroleassociation.TimeoutsAttributes](gra.ref.Append("timeouts"))
}

type grafanaRoleAssociationState struct {
	GroupIds    []string                              `json:"group_ids"`
	Id          string                                `json:"id"`
	Role        string                                `json:"role"`
	UserIds     []string                              `json:"user_ids"`
	WorkspaceId string                                `json:"workspace_id"`
	Timeouts    *grafanaroleassociation.TimeoutsState `json:"timeouts"`
}
