// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package vault

import "github.com/golingon/lingon/pkg/terra"

// NewDataIdentityGroup creates a new instance of [DataIdentityGroup].
func NewDataIdentityGroup(name string, args DataIdentityGroupArgs) *DataIdentityGroup {
	return &DataIdentityGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIdentityGroup)(nil)

// DataIdentityGroup represents the Terraform data resource vault_identity_group.
type DataIdentityGroup struct {
	Name string
	Args DataIdentityGroupArgs
}

// DataSource returns the Terraform object type for [DataIdentityGroup].
func (ig *DataIdentityGroup) DataSource() string {
	return "vault_identity_group"
}

// LocalName returns the local name for [DataIdentityGroup].
func (ig *DataIdentityGroup) LocalName() string {
	return ig.Name
}

// Configuration returns the configuration (args) for [DataIdentityGroup].
func (ig *DataIdentityGroup) Configuration() interface{} {
	return ig.Args
}

// Attributes returns the attributes for [DataIdentityGroup].
func (ig *DataIdentityGroup) Attributes() dataIdentityGroupAttributes {
	return dataIdentityGroupAttributes{ref: terra.ReferenceDataResource(ig)}
}

// DataIdentityGroupArgs contains the configurations for vault_identity_group.
type DataIdentityGroupArgs struct {
	// AliasId: string, optional
	AliasId terra.StringValue `hcl:"alias_id,attr"`
	// AliasMountAccessor: string, optional
	AliasMountAccessor terra.StringValue `hcl:"alias_mount_accessor,attr"`
	// AliasName: string, optional
	AliasName terra.StringValue `hcl:"alias_name,attr"`
	// GroupId: string, optional
	GroupId terra.StringValue `hcl:"group_id,attr"`
	// GroupName: string, optional
	GroupName terra.StringValue `hcl:"group_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
}
type dataIdentityGroupAttributes struct {
	ref terra.Reference
}

// AliasCanonicalId returns a reference to field alias_canonical_id of vault_identity_group.
func (ig dataIdentityGroupAttributes) AliasCanonicalId() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("alias_canonical_id"))
}

// AliasCreationTime returns a reference to field alias_creation_time of vault_identity_group.
func (ig dataIdentityGroupAttributes) AliasCreationTime() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("alias_creation_time"))
}

// AliasId returns a reference to field alias_id of vault_identity_group.
func (ig dataIdentityGroupAttributes) AliasId() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("alias_id"))
}

// AliasLastUpdateTime returns a reference to field alias_last_update_time of vault_identity_group.
func (ig dataIdentityGroupAttributes) AliasLastUpdateTime() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("alias_last_update_time"))
}

// AliasMergedFromCanonicalIds returns a reference to field alias_merged_from_canonical_ids of vault_identity_group.
func (ig dataIdentityGroupAttributes) AliasMergedFromCanonicalIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ig.ref.Append("alias_merged_from_canonical_ids"))
}

// AliasMetadata returns a reference to field alias_metadata of vault_identity_group.
func (ig dataIdentityGroupAttributes) AliasMetadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ig.ref.Append("alias_metadata"))
}

// AliasMountAccessor returns a reference to field alias_mount_accessor of vault_identity_group.
func (ig dataIdentityGroupAttributes) AliasMountAccessor() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("alias_mount_accessor"))
}

// AliasMountPath returns a reference to field alias_mount_path of vault_identity_group.
func (ig dataIdentityGroupAttributes) AliasMountPath() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("alias_mount_path"))
}

// AliasMountType returns a reference to field alias_mount_type of vault_identity_group.
func (ig dataIdentityGroupAttributes) AliasMountType() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("alias_mount_type"))
}

// AliasName returns a reference to field alias_name of vault_identity_group.
func (ig dataIdentityGroupAttributes) AliasName() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("alias_name"))
}

// CreationTime returns a reference to field creation_time of vault_identity_group.
func (ig dataIdentityGroupAttributes) CreationTime() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("creation_time"))
}

// DataJson returns a reference to field data_json of vault_identity_group.
func (ig dataIdentityGroupAttributes) DataJson() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("data_json"))
}

// GroupId returns a reference to field group_id of vault_identity_group.
func (ig dataIdentityGroupAttributes) GroupId() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("group_id"))
}

// GroupName returns a reference to field group_name of vault_identity_group.
func (ig dataIdentityGroupAttributes) GroupName() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("group_name"))
}

// Id returns a reference to field id of vault_identity_group.
func (ig dataIdentityGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("id"))
}

// LastUpdateTime returns a reference to field last_update_time of vault_identity_group.
func (ig dataIdentityGroupAttributes) LastUpdateTime() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("last_update_time"))
}

// MemberEntityIds returns a reference to field member_entity_ids of vault_identity_group.
func (ig dataIdentityGroupAttributes) MemberEntityIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ig.ref.Append("member_entity_ids"))
}

// MemberGroupIds returns a reference to field member_group_ids of vault_identity_group.
func (ig dataIdentityGroupAttributes) MemberGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ig.ref.Append("member_group_ids"))
}

// Metadata returns a reference to field metadata of vault_identity_group.
func (ig dataIdentityGroupAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ig.ref.Append("metadata"))
}

// ModifyIndex returns a reference to field modify_index of vault_identity_group.
func (ig dataIdentityGroupAttributes) ModifyIndex() terra.NumberValue {
	return terra.ReferenceAsNumber(ig.ref.Append("modify_index"))
}

// Namespace returns a reference to field namespace of vault_identity_group.
func (ig dataIdentityGroupAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("namespace"))
}

// NamespaceId returns a reference to field namespace_id of vault_identity_group.
func (ig dataIdentityGroupAttributes) NamespaceId() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("namespace_id"))
}

// ParentGroupIds returns a reference to field parent_group_ids of vault_identity_group.
func (ig dataIdentityGroupAttributes) ParentGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ig.ref.Append("parent_group_ids"))
}

// Policies returns a reference to field policies of vault_identity_group.
func (ig dataIdentityGroupAttributes) Policies() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ig.ref.Append("policies"))
}

// Type returns a reference to field type of vault_identity_group.
func (ig dataIdentityGroupAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("type"))
}
