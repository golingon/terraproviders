// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package nginxdeployment

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type AutoScaleProfile struct {
	// MaxCapacity: number, required
	MaxCapacity terra.NumberValue `hcl:"max_capacity,attr" validate:"required"`
	// MinCapacity: number, required
	MinCapacity terra.NumberValue `hcl:"min_capacity,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type Configuration struct {
	// PackageData: string, optional
	PackageData terra.StringValue `hcl:"package_data,attr"`
	// RootFile: string, required
	RootFile terra.StringValue `hcl:"root_file,attr" validate:"required"`
	// ConfigFile: min=0
	ConfigFile []ConfigFile `hcl:"config_file,block" validate:"min=0"`
	// ProtectedFile: min=0
	ProtectedFile []ProtectedFile `hcl:"protected_file,block" validate:"min=0"`
}

type ConfigFile struct {
	// Content: string, required
	Content terra.StringValue `hcl:"content,attr" validate:"required"`
	// VirtualPath: string, required
	VirtualPath terra.StringValue `hcl:"virtual_path,attr" validate:"required"`
}

type ProtectedFile struct {
	// Content: string, required
	Content terra.StringValue `hcl:"content,attr" validate:"required"`
	// VirtualPath: string, required
	VirtualPath terra.StringValue `hcl:"virtual_path,attr" validate:"required"`
}

type FrontendPrivate struct {
	// AllocationMethod: string, required
	AllocationMethod terra.StringValue `hcl:"allocation_method,attr" validate:"required"`
	// IpAddress: string, required
	IpAddress terra.StringValue `hcl:"ip_address,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
}

type FrontendPublic struct {
	// IpAddress: list of string, optional
	IpAddress terra.ListValue[terra.StringValue] `hcl:"ip_address,attr"`
}

type Identity struct {
	// IdentityIds: set of string, optional
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type LoggingStorageAccount struct {
	// ContainerName: string, optional
	ContainerName terra.StringValue `hcl:"container_name,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
}

type NetworkInterface struct {
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AutoScaleProfileAttributes struct {
	ref terra.Reference
}

func (asp AutoScaleProfileAttributes) InternalRef() (terra.Reference, error) {
	return asp.ref, nil
}

func (asp AutoScaleProfileAttributes) InternalWithRef(ref terra.Reference) AutoScaleProfileAttributes {
	return AutoScaleProfileAttributes{ref: ref}
}

func (asp AutoScaleProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return asp.ref.InternalTokens()
}

func (asp AutoScaleProfileAttributes) MaxCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(asp.ref.Append("max_capacity"))
}

func (asp AutoScaleProfileAttributes) MinCapacity() terra.NumberValue {
	return terra.ReferenceAsNumber(asp.ref.Append("min_capacity"))
}

func (asp AutoScaleProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asp.ref.Append("name"))
}

type ConfigurationAttributes struct {
	ref terra.Reference
}

func (c ConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConfigurationAttributes) InternalWithRef(ref terra.Reference) ConfigurationAttributes {
	return ConfigurationAttributes{ref: ref}
}

func (c ConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConfigurationAttributes) PackageData() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("package_data"))
}

func (c ConfigurationAttributes) RootFile() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("root_file"))
}

func (c ConfigurationAttributes) ConfigFile() terra.SetValue[ConfigFileAttributes] {
	return terra.ReferenceAsSet[ConfigFileAttributes](c.ref.Append("config_file"))
}

func (c ConfigurationAttributes) ProtectedFile() terra.SetValue[ProtectedFileAttributes] {
	return terra.ReferenceAsSet[ProtectedFileAttributes](c.ref.Append("protected_file"))
}

type ConfigFileAttributes struct {
	ref terra.Reference
}

func (cf ConfigFileAttributes) InternalRef() (terra.Reference, error) {
	return cf.ref, nil
}

func (cf ConfigFileAttributes) InternalWithRef(ref terra.Reference) ConfigFileAttributes {
	return ConfigFileAttributes{ref: ref}
}

func (cf ConfigFileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cf.ref.InternalTokens()
}

func (cf ConfigFileAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("content"))
}

func (cf ConfigFileAttributes) VirtualPath() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("virtual_path"))
}

type ProtectedFileAttributes struct {
	ref terra.Reference
}

func (pf ProtectedFileAttributes) InternalRef() (terra.Reference, error) {
	return pf.ref, nil
}

func (pf ProtectedFileAttributes) InternalWithRef(ref terra.Reference) ProtectedFileAttributes {
	return ProtectedFileAttributes{ref: ref}
}

func (pf ProtectedFileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pf.ref.InternalTokens()
}

func (pf ProtectedFileAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(pf.ref.Append("content"))
}

func (pf ProtectedFileAttributes) VirtualPath() terra.StringValue {
	return terra.ReferenceAsString(pf.ref.Append("virtual_path"))
}

type FrontendPrivateAttributes struct {
	ref terra.Reference
}

func (fp FrontendPrivateAttributes) InternalRef() (terra.Reference, error) {
	return fp.ref, nil
}

func (fp FrontendPrivateAttributes) InternalWithRef(ref terra.Reference) FrontendPrivateAttributes {
	return FrontendPrivateAttributes{ref: ref}
}

func (fp FrontendPrivateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fp.ref.InternalTokens()
}

func (fp FrontendPrivateAttributes) AllocationMethod() terra.StringValue {
	return terra.ReferenceAsString(fp.ref.Append("allocation_method"))
}

func (fp FrontendPrivateAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(fp.ref.Append("ip_address"))
}

func (fp FrontendPrivateAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(fp.ref.Append("subnet_id"))
}

type FrontendPublicAttributes struct {
	ref terra.Reference
}

func (fp FrontendPublicAttributes) InternalRef() (terra.Reference, error) {
	return fp.ref, nil
}

func (fp FrontendPublicAttributes) InternalWithRef(ref terra.Reference) FrontendPublicAttributes {
	return FrontendPublicAttributes{ref: ref}
}

func (fp FrontendPublicAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fp.ref.InternalTokens()
}

func (fp FrontendPublicAttributes) IpAddress() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](fp.ref.Append("ip_address"))
}

type IdentityAttributes struct {
	ref terra.Reference
}

func (i IdentityAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IdentityAttributes) InternalWithRef(ref terra.Reference) IdentityAttributes {
	return IdentityAttributes{ref: ref}
}

func (i IdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IdentityAttributes) IdentityIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i IdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("principal_id"))
}

func (i IdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenant_id"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type LoggingStorageAccountAttributes struct {
	ref terra.Reference
}

func (lsa LoggingStorageAccountAttributes) InternalRef() (terra.Reference, error) {
	return lsa.ref, nil
}

func (lsa LoggingStorageAccountAttributes) InternalWithRef(ref terra.Reference) LoggingStorageAccountAttributes {
	return LoggingStorageAccountAttributes{ref: ref}
}

func (lsa LoggingStorageAccountAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lsa.ref.InternalTokens()
}

func (lsa LoggingStorageAccountAttributes) ContainerName() terra.StringValue {
	return terra.ReferenceAsString(lsa.ref.Append("container_name"))
}

func (lsa LoggingStorageAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lsa.ref.Append("name"))
}

type NetworkInterfaceAttributes struct {
	ref terra.Reference
}

func (ni NetworkInterfaceAttributes) InternalRef() (terra.Reference, error) {
	return ni.ref, nil
}

func (ni NetworkInterfaceAttributes) InternalWithRef(ref terra.Reference) NetworkInterfaceAttributes {
	return NetworkInterfaceAttributes{ref: ref}
}

func (ni NetworkInterfaceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ni.ref.InternalTokens()
}

func (ni NetworkInterfaceAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(ni.ref.Append("subnet_id"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type AutoScaleProfileState struct {
	MaxCapacity float64 `json:"max_capacity"`
	MinCapacity float64 `json:"min_capacity"`
	Name        string  `json:"name"`
}

type ConfigurationState struct {
	PackageData   string               `json:"package_data"`
	RootFile      string               `json:"root_file"`
	ConfigFile    []ConfigFileState    `json:"config_file"`
	ProtectedFile []ProtectedFileState `json:"protected_file"`
}

type ConfigFileState struct {
	Content     string `json:"content"`
	VirtualPath string `json:"virtual_path"`
}

type ProtectedFileState struct {
	Content     string `json:"content"`
	VirtualPath string `json:"virtual_path"`
}

type FrontendPrivateState struct {
	AllocationMethod string `json:"allocation_method"`
	IpAddress        string `json:"ip_address"`
	SubnetId         string `json:"subnet_id"`
}

type FrontendPublicState struct {
	IpAddress []string `json:"ip_address"`
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type LoggingStorageAccountState struct {
	ContainerName string `json:"container_name"`
	Name          string `json:"name"`
}

type NetworkInterfaceState struct {
	SubnetId string `json:"subnet_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
