// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	networkmanageradminrulecollection "github.com/golingon/terraproviders/azurerm/3.49.0/networkmanageradminrulecollection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewNetworkManagerAdminRuleCollection(name string, args NetworkManagerAdminRuleCollectionArgs) *NetworkManagerAdminRuleCollection {
	return &NetworkManagerAdminRuleCollection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkManagerAdminRuleCollection)(nil)

type NetworkManagerAdminRuleCollection struct {
	Name  string
	Args  NetworkManagerAdminRuleCollectionArgs
	state *networkManagerAdminRuleCollectionState
}

func (nmarc *NetworkManagerAdminRuleCollection) Type() string {
	return "azurerm_network_manager_admin_rule_collection"
}

func (nmarc *NetworkManagerAdminRuleCollection) LocalName() string {
	return nmarc.Name
}

func (nmarc *NetworkManagerAdminRuleCollection) Configuration() interface{} {
	return nmarc.Args
}

func (nmarc *NetworkManagerAdminRuleCollection) Attributes() networkManagerAdminRuleCollectionAttributes {
	return networkManagerAdminRuleCollectionAttributes{ref: terra.ReferenceResource(nmarc)}
}

func (nmarc *NetworkManagerAdminRuleCollection) ImportState(av io.Reader) error {
	nmarc.state = &networkManagerAdminRuleCollectionState{}
	if err := json.NewDecoder(av).Decode(nmarc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nmarc.Type(), nmarc.LocalName(), err)
	}
	return nil
}

func (nmarc *NetworkManagerAdminRuleCollection) State() (*networkManagerAdminRuleCollectionState, bool) {
	return nmarc.state, nmarc.state != nil
}

func (nmarc *NetworkManagerAdminRuleCollection) StateMust() *networkManagerAdminRuleCollectionState {
	if nmarc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nmarc.Type(), nmarc.LocalName()))
	}
	return nmarc.state
}

func (nmarc *NetworkManagerAdminRuleCollection) DependOn() terra.Reference {
	return terra.ReferenceResource(nmarc)
}

type NetworkManagerAdminRuleCollectionArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkGroupIds: list of string, required
	NetworkGroupIds terra.ListValue[terra.StringValue] `hcl:"network_group_ids,attr" validate:"required"`
	// SecurityAdminConfigurationId: string, required
	SecurityAdminConfigurationId terra.StringValue `hcl:"security_admin_configuration_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *networkmanageradminrulecollection.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that NetworkManagerAdminRuleCollection depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type networkManagerAdminRuleCollectionAttributes struct {
	ref terra.Reference
}

func (nmarc networkManagerAdminRuleCollectionAttributes) Description() terra.StringValue {
	return terra.ReferenceString(nmarc.ref.Append("description"))
}

func (nmarc networkManagerAdminRuleCollectionAttributes) Id() terra.StringValue {
	return terra.ReferenceString(nmarc.ref.Append("id"))
}

func (nmarc networkManagerAdminRuleCollectionAttributes) Name() terra.StringValue {
	return terra.ReferenceString(nmarc.ref.Append("name"))
}

func (nmarc networkManagerAdminRuleCollectionAttributes) NetworkGroupIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](nmarc.ref.Append("network_group_ids"))
}

func (nmarc networkManagerAdminRuleCollectionAttributes) SecurityAdminConfigurationId() terra.StringValue {
	return terra.ReferenceString(nmarc.ref.Append("security_admin_configuration_id"))
}

func (nmarc networkManagerAdminRuleCollectionAttributes) Timeouts() networkmanageradminrulecollection.TimeoutsAttributes {
	return terra.ReferenceSingle[networkmanageradminrulecollection.TimeoutsAttributes](nmarc.ref.Append("timeouts"))
}

type networkManagerAdminRuleCollectionState struct {
	Description                  string                                           `json:"description"`
	Id                           string                                           `json:"id"`
	Name                         string                                           `json:"name"`
	NetworkGroupIds              []string                                         `json:"network_group_ids"`
	SecurityAdminConfigurationId string                                           `json:"security_admin_configuration_id"`
	Timeouts                     *networkmanageradminrulecollection.TimeoutsState `json:"timeouts"`
}
