// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	siterecoveryfabric "github.com/golingon/terraproviders/azurerm/3.49.0/siterecoveryfabric"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSiteRecoveryFabric(name string, args SiteRecoveryFabricArgs) *SiteRecoveryFabric {
	return &SiteRecoveryFabric{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SiteRecoveryFabric)(nil)

type SiteRecoveryFabric struct {
	Name  string
	Args  SiteRecoveryFabricArgs
	state *siteRecoveryFabricState
}

func (srf *SiteRecoveryFabric) Type() string {
	return "azurerm_site_recovery_fabric"
}

func (srf *SiteRecoveryFabric) LocalName() string {
	return srf.Name
}

func (srf *SiteRecoveryFabric) Configuration() interface{} {
	return srf.Args
}

func (srf *SiteRecoveryFabric) Attributes() siteRecoveryFabricAttributes {
	return siteRecoveryFabricAttributes{ref: terra.ReferenceResource(srf)}
}

func (srf *SiteRecoveryFabric) ImportState(av io.Reader) error {
	srf.state = &siteRecoveryFabricState{}
	if err := json.NewDecoder(av).Decode(srf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", srf.Type(), srf.LocalName(), err)
	}
	return nil
}

func (srf *SiteRecoveryFabric) State() (*siteRecoveryFabricState, bool) {
	return srf.state, srf.state != nil
}

func (srf *SiteRecoveryFabric) StateMust() *siteRecoveryFabricState {
	if srf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", srf.Type(), srf.LocalName()))
	}
	return srf.state
}

func (srf *SiteRecoveryFabric) DependOn() terra.Reference {
	return terra.ReferenceResource(srf)
}

type SiteRecoveryFabricArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RecoveryVaultName: string, required
	RecoveryVaultName terra.StringValue `hcl:"recovery_vault_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *siterecoveryfabric.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that SiteRecoveryFabric depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type siteRecoveryFabricAttributes struct {
	ref terra.Reference
}

func (srf siteRecoveryFabricAttributes) Id() terra.StringValue {
	return terra.ReferenceString(srf.ref.Append("id"))
}

func (srf siteRecoveryFabricAttributes) Location() terra.StringValue {
	return terra.ReferenceString(srf.ref.Append("location"))
}

func (srf siteRecoveryFabricAttributes) Name() terra.StringValue {
	return terra.ReferenceString(srf.ref.Append("name"))
}

func (srf siteRecoveryFabricAttributes) RecoveryVaultName() terra.StringValue {
	return terra.ReferenceString(srf.ref.Append("recovery_vault_name"))
}

func (srf siteRecoveryFabricAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(srf.ref.Append("resource_group_name"))
}

func (srf siteRecoveryFabricAttributes) Timeouts() siterecoveryfabric.TimeoutsAttributes {
	return terra.ReferenceSingle[siterecoveryfabric.TimeoutsAttributes](srf.ref.Append("timeouts"))
}

type siteRecoveryFabricState struct {
	Id                string                            `json:"id"`
	Location          string                            `json:"location"`
	Name              string                            `json:"name"`
	RecoveryVaultName string                            `json:"recovery_vault_name"`
	ResourceGroupName string                            `json:"resource_group_name"`
	Timeouts          *siterecoveryfabric.TimeoutsState `json:"timeouts"`
}
