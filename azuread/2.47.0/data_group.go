// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azuread

import (
	"github.com/golingon/lingon/pkg/terra"
	datagroup "github.com/golingon/terraproviders/azuread/2.47.0/datagroup"
)

// NewDataGroup creates a new instance of [DataGroup].
func NewDataGroup(name string, args DataGroupArgs) *DataGroup {
	return &DataGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataGroup)(nil)

// DataGroup represents the Terraform data resource azuread_group.
type DataGroup struct {
	Name string
	Args DataGroupArgs
}

// DataSource returns the Terraform object type for [DataGroup].
func (g *DataGroup) DataSource() string {
	return "azuread_group"
}

// LocalName returns the local name for [DataGroup].
func (g *DataGroup) LocalName() string {
	return g.Name
}

// Configuration returns the configuration (args) for [DataGroup].
func (g *DataGroup) Configuration() interface{} {
	return g.Args
}

// Attributes returns the attributes for [DataGroup].
func (g *DataGroup) Attributes() dataGroupAttributes {
	return dataGroupAttributes{ref: terra.ReferenceDataResource(g)}
}

// DataGroupArgs contains the configurations for azuread_group.
type DataGroupArgs struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MailEnabled: bool, optional
	MailEnabled terra.BoolValue `hcl:"mail_enabled,attr"`
	// MailNickname: string, optional
	MailNickname terra.StringValue `hcl:"mail_nickname,attr"`
	// ObjectId: string, optional
	ObjectId terra.StringValue `hcl:"object_id,attr"`
	// SecurityEnabled: bool, optional
	SecurityEnabled terra.BoolValue `hcl:"security_enabled,attr"`
	// DynamicMembership: min=0
	DynamicMembership []datagroup.DynamicMembership `hcl:"dynamic_membership,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datagroup.Timeouts `hcl:"timeouts,block"`
}
type dataGroupAttributes struct {
	ref terra.Reference
}

// AssignableToRole returns a reference to field assignable_to_role of azuread_group.
func (g dataGroupAttributes) AssignableToRole() terra.BoolValue {
	return terra.ReferenceAsBool(g.ref.Append("assignable_to_role"))
}

// AutoSubscribeNewMembers returns a reference to field auto_subscribe_new_members of azuread_group.
func (g dataGroupAttributes) AutoSubscribeNewMembers() terra.BoolValue {
	return terra.ReferenceAsBool(g.ref.Append("auto_subscribe_new_members"))
}

// Behaviors returns a reference to field behaviors of azuread_group.
func (g dataGroupAttributes) Behaviors() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](g.ref.Append("behaviors"))
}

// Description returns a reference to field description of azuread_group.
func (g dataGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azuread_group.
func (g dataGroupAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("display_name"))
}

// ExternalSendersAllowed returns a reference to field external_senders_allowed of azuread_group.
func (g dataGroupAttributes) ExternalSendersAllowed() terra.BoolValue {
	return terra.ReferenceAsBool(g.ref.Append("external_senders_allowed"))
}

// HideFromAddressLists returns a reference to field hide_from_address_lists of azuread_group.
func (g dataGroupAttributes) HideFromAddressLists() terra.BoolValue {
	return terra.ReferenceAsBool(g.ref.Append("hide_from_address_lists"))
}

// HideFromOutlookClients returns a reference to field hide_from_outlook_clients of azuread_group.
func (g dataGroupAttributes) HideFromOutlookClients() terra.BoolValue {
	return terra.ReferenceAsBool(g.ref.Append("hide_from_outlook_clients"))
}

// Id returns a reference to field id of azuread_group.
func (g dataGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("id"))
}

// Mail returns a reference to field mail of azuread_group.
func (g dataGroupAttributes) Mail() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("mail"))
}

// MailEnabled returns a reference to field mail_enabled of azuread_group.
func (g dataGroupAttributes) MailEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(g.ref.Append("mail_enabled"))
}

// MailNickname returns a reference to field mail_nickname of azuread_group.
func (g dataGroupAttributes) MailNickname() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("mail_nickname"))
}

// Members returns a reference to field members of azuread_group.
func (g dataGroupAttributes) Members() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](g.ref.Append("members"))
}

// ObjectId returns a reference to field object_id of azuread_group.
func (g dataGroupAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("object_id"))
}

// OnpremisesDomainName returns a reference to field onpremises_domain_name of azuread_group.
func (g dataGroupAttributes) OnpremisesDomainName() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("onpremises_domain_name"))
}

// OnpremisesGroupType returns a reference to field onpremises_group_type of azuread_group.
func (g dataGroupAttributes) OnpremisesGroupType() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("onpremises_group_type"))
}

// OnpremisesNetbiosName returns a reference to field onpremises_netbios_name of azuread_group.
func (g dataGroupAttributes) OnpremisesNetbiosName() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("onpremises_netbios_name"))
}

// OnpremisesSamAccountName returns a reference to field onpremises_sam_account_name of azuread_group.
func (g dataGroupAttributes) OnpremisesSamAccountName() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("onpremises_sam_account_name"))
}

// OnpremisesSecurityIdentifier returns a reference to field onpremises_security_identifier of azuread_group.
func (g dataGroupAttributes) OnpremisesSecurityIdentifier() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("onpremises_security_identifier"))
}

// OnpremisesSyncEnabled returns a reference to field onpremises_sync_enabled of azuread_group.
func (g dataGroupAttributes) OnpremisesSyncEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(g.ref.Append("onpremises_sync_enabled"))
}

// Owners returns a reference to field owners of azuread_group.
func (g dataGroupAttributes) Owners() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](g.ref.Append("owners"))
}

// PreferredLanguage returns a reference to field preferred_language of azuread_group.
func (g dataGroupAttributes) PreferredLanguage() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("preferred_language"))
}

// ProvisioningOptions returns a reference to field provisioning_options of azuread_group.
func (g dataGroupAttributes) ProvisioningOptions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](g.ref.Append("provisioning_options"))
}

// ProxyAddresses returns a reference to field proxy_addresses of azuread_group.
func (g dataGroupAttributes) ProxyAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](g.ref.Append("proxy_addresses"))
}

// SecurityEnabled returns a reference to field security_enabled of azuread_group.
func (g dataGroupAttributes) SecurityEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(g.ref.Append("security_enabled"))
}

// Theme returns a reference to field theme of azuread_group.
func (g dataGroupAttributes) Theme() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("theme"))
}

// Types returns a reference to field types of azuread_group.
func (g dataGroupAttributes) Types() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](g.ref.Append("types"))
}

// Visibility returns a reference to field visibility of azuread_group.
func (g dataGroupAttributes) Visibility() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("visibility"))
}

// WritebackEnabled returns a reference to field writeback_enabled of azuread_group.
func (g dataGroupAttributes) WritebackEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(g.ref.Append("writeback_enabled"))
}

func (g dataGroupAttributes) DynamicMembership() terra.ListValue[datagroup.DynamicMembershipAttributes] {
	return terra.ReferenceAsList[datagroup.DynamicMembershipAttributes](g.ref.Append("dynamic_membership"))
}

func (g dataGroupAttributes) Timeouts() datagroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datagroup.TimeoutsAttributes](g.ref.Append("timeouts"))
}
