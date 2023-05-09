// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	cdnfrontdoorcustomdomainassociation "github.com/golingon/terraproviders/azurerm/3.55.0/cdnfrontdoorcustomdomainassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCdnFrontdoorCustomDomainAssociation creates a new instance of [CdnFrontdoorCustomDomainAssociation].
func NewCdnFrontdoorCustomDomainAssociation(name string, args CdnFrontdoorCustomDomainAssociationArgs) *CdnFrontdoorCustomDomainAssociation {
	return &CdnFrontdoorCustomDomainAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CdnFrontdoorCustomDomainAssociation)(nil)

// CdnFrontdoorCustomDomainAssociation represents the Terraform resource azurerm_cdn_frontdoor_custom_domain_association.
type CdnFrontdoorCustomDomainAssociation struct {
	Name      string
	Args      CdnFrontdoorCustomDomainAssociationArgs
	state     *cdnFrontdoorCustomDomainAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CdnFrontdoorCustomDomainAssociation].
func (cfcda *CdnFrontdoorCustomDomainAssociation) Type() string {
	return "azurerm_cdn_frontdoor_custom_domain_association"
}

// LocalName returns the local name for [CdnFrontdoorCustomDomainAssociation].
func (cfcda *CdnFrontdoorCustomDomainAssociation) LocalName() string {
	return cfcda.Name
}

// Configuration returns the configuration (args) for [CdnFrontdoorCustomDomainAssociation].
func (cfcda *CdnFrontdoorCustomDomainAssociation) Configuration() interface{} {
	return cfcda.Args
}

// DependOn is used for other resources to depend on [CdnFrontdoorCustomDomainAssociation].
func (cfcda *CdnFrontdoorCustomDomainAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(cfcda)
}

// Dependencies returns the list of resources [CdnFrontdoorCustomDomainAssociation] depends_on.
func (cfcda *CdnFrontdoorCustomDomainAssociation) Dependencies() terra.Dependencies {
	return cfcda.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CdnFrontdoorCustomDomainAssociation].
func (cfcda *CdnFrontdoorCustomDomainAssociation) LifecycleManagement() *terra.Lifecycle {
	return cfcda.Lifecycle
}

// Attributes returns the attributes for [CdnFrontdoorCustomDomainAssociation].
func (cfcda *CdnFrontdoorCustomDomainAssociation) Attributes() cdnFrontdoorCustomDomainAssociationAttributes {
	return cdnFrontdoorCustomDomainAssociationAttributes{ref: terra.ReferenceResource(cfcda)}
}

// ImportState imports the given attribute values into [CdnFrontdoorCustomDomainAssociation]'s state.
func (cfcda *CdnFrontdoorCustomDomainAssociation) ImportState(av io.Reader) error {
	cfcda.state = &cdnFrontdoorCustomDomainAssociationState{}
	if err := json.NewDecoder(av).Decode(cfcda.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cfcda.Type(), cfcda.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CdnFrontdoorCustomDomainAssociation] has state.
func (cfcda *CdnFrontdoorCustomDomainAssociation) State() (*cdnFrontdoorCustomDomainAssociationState, bool) {
	return cfcda.state, cfcda.state != nil
}

// StateMust returns the state for [CdnFrontdoorCustomDomainAssociation]. Panics if the state is nil.
func (cfcda *CdnFrontdoorCustomDomainAssociation) StateMust() *cdnFrontdoorCustomDomainAssociationState {
	if cfcda.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cfcda.Type(), cfcda.LocalName()))
	}
	return cfcda.state
}

// CdnFrontdoorCustomDomainAssociationArgs contains the configurations for azurerm_cdn_frontdoor_custom_domain_association.
type CdnFrontdoorCustomDomainAssociationArgs struct {
	// CdnFrontdoorCustomDomainId: string, required
	CdnFrontdoorCustomDomainId terra.StringValue `hcl:"cdn_frontdoor_custom_domain_id,attr" validate:"required"`
	// CdnFrontdoorRouteIds: list of string, required
	CdnFrontdoorRouteIds terra.ListValue[terra.StringValue] `hcl:"cdn_frontdoor_route_ids,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *cdnfrontdoorcustomdomainassociation.Timeouts `hcl:"timeouts,block"`
}
type cdnFrontdoorCustomDomainAssociationAttributes struct {
	ref terra.Reference
}

// CdnFrontdoorCustomDomainId returns a reference to field cdn_frontdoor_custom_domain_id of azurerm_cdn_frontdoor_custom_domain_association.
func (cfcda cdnFrontdoorCustomDomainAssociationAttributes) CdnFrontdoorCustomDomainId() terra.StringValue {
	return terra.ReferenceAsString(cfcda.ref.Append("cdn_frontdoor_custom_domain_id"))
}

// CdnFrontdoorRouteIds returns a reference to field cdn_frontdoor_route_ids of azurerm_cdn_frontdoor_custom_domain_association.
func (cfcda cdnFrontdoorCustomDomainAssociationAttributes) CdnFrontdoorRouteIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cfcda.ref.Append("cdn_frontdoor_route_ids"))
}

// Id returns a reference to field id of azurerm_cdn_frontdoor_custom_domain_association.
func (cfcda cdnFrontdoorCustomDomainAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cfcda.ref.Append("id"))
}

func (cfcda cdnFrontdoorCustomDomainAssociationAttributes) Timeouts() cdnfrontdoorcustomdomainassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cdnfrontdoorcustomdomainassociation.TimeoutsAttributes](cfcda.ref.Append("timeouts"))
}

type cdnFrontdoorCustomDomainAssociationState struct {
	CdnFrontdoorCustomDomainId string                                             `json:"cdn_frontdoor_custom_domain_id"`
	CdnFrontdoorRouteIds       []string                                           `json:"cdn_frontdoor_route_ids"`
	Id                         string                                             `json:"id"`
	Timeouts                   *cdnfrontdoorcustomdomainassociation.TimeoutsState `json:"timeouts"`
}
