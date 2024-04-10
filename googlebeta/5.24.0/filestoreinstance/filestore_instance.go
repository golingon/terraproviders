// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package filestoreinstance

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type FileShares struct {
	// CapacityGb: number, required
	CapacityGb terra.NumberValue `hcl:"capacity_gb,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SourceBackup: string, optional
	SourceBackup terra.StringValue `hcl:"source_backup,attr"`
	// NfsExportOptions: min=0,max=10
	NfsExportOptions []NfsExportOptions `hcl:"nfs_export_options,block" validate:"min=0,max=10"`
}

type NfsExportOptions struct {
	// AccessMode: string, optional
	AccessMode terra.StringValue `hcl:"access_mode,attr"`
	// AnonGid: number, optional
	AnonGid terra.NumberValue `hcl:"anon_gid,attr"`
	// AnonUid: number, optional
	AnonUid terra.NumberValue `hcl:"anon_uid,attr"`
	// IpRanges: list of string, optional
	IpRanges terra.ListValue[terra.StringValue] `hcl:"ip_ranges,attr"`
	// SquashMode: string, optional
	SquashMode terra.StringValue `hcl:"squash_mode,attr"`
}

type Networks struct {
	// ConnectMode: string, optional
	ConnectMode terra.StringValue `hcl:"connect_mode,attr"`
	// Modes: list of string, required
	Modes terra.ListValue[terra.StringValue] `hcl:"modes,attr" validate:"required"`
	// Network: string, required
	Network terra.StringValue `hcl:"network,attr" validate:"required"`
	// ReservedIpRange: string, optional
	ReservedIpRange terra.StringValue `hcl:"reserved_ip_range,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type FileSharesAttributes struct {
	ref terra.Reference
}

func (fs FileSharesAttributes) InternalRef() (terra.Reference, error) {
	return fs.ref, nil
}

func (fs FileSharesAttributes) InternalWithRef(ref terra.Reference) FileSharesAttributes {
	return FileSharesAttributes{ref: ref}
}

func (fs FileSharesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fs.ref.InternalTokens()
}

func (fs FileSharesAttributes) CapacityGb() terra.NumberValue {
	return terra.ReferenceAsNumber(fs.ref.Append("capacity_gb"))
}

func (fs FileSharesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fs.ref.Append("name"))
}

func (fs FileSharesAttributes) SourceBackup() terra.StringValue {
	return terra.ReferenceAsString(fs.ref.Append("source_backup"))
}

func (fs FileSharesAttributes) NfsExportOptions() terra.ListValue[NfsExportOptionsAttributes] {
	return terra.ReferenceAsList[NfsExportOptionsAttributes](fs.ref.Append("nfs_export_options"))
}

type NfsExportOptionsAttributes struct {
	ref terra.Reference
}

func (neo NfsExportOptionsAttributes) InternalRef() (terra.Reference, error) {
	return neo.ref, nil
}

func (neo NfsExportOptionsAttributes) InternalWithRef(ref terra.Reference) NfsExportOptionsAttributes {
	return NfsExportOptionsAttributes{ref: ref}
}

func (neo NfsExportOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return neo.ref.InternalTokens()
}

func (neo NfsExportOptionsAttributes) AccessMode() terra.StringValue {
	return terra.ReferenceAsString(neo.ref.Append("access_mode"))
}

func (neo NfsExportOptionsAttributes) AnonGid() terra.NumberValue {
	return terra.ReferenceAsNumber(neo.ref.Append("anon_gid"))
}

func (neo NfsExportOptionsAttributes) AnonUid() terra.NumberValue {
	return terra.ReferenceAsNumber(neo.ref.Append("anon_uid"))
}

func (neo NfsExportOptionsAttributes) IpRanges() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](neo.ref.Append("ip_ranges"))
}

func (neo NfsExportOptionsAttributes) SquashMode() terra.StringValue {
	return terra.ReferenceAsString(neo.ref.Append("squash_mode"))
}

type NetworksAttributes struct {
	ref terra.Reference
}

func (n NetworksAttributes) InternalRef() (terra.Reference, error) {
	return n.ref, nil
}

func (n NetworksAttributes) InternalWithRef(ref terra.Reference) NetworksAttributes {
	return NetworksAttributes{ref: ref}
}

func (n NetworksAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return n.ref.InternalTokens()
}

func (n NetworksAttributes) ConnectMode() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("connect_mode"))
}

func (n NetworksAttributes) IpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](n.ref.Append("ip_addresses"))
}

func (n NetworksAttributes) Modes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](n.ref.Append("modes"))
}

func (n NetworksAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("network"))
}

func (n NetworksAttributes) ReservedIpRange() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("reserved_ip_range"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type FileSharesState struct {
	CapacityGb       float64                 `json:"capacity_gb"`
	Name             string                  `json:"name"`
	SourceBackup     string                  `json:"source_backup"`
	NfsExportOptions []NfsExportOptionsState `json:"nfs_export_options"`
}

type NfsExportOptionsState struct {
	AccessMode string   `json:"access_mode"`
	AnonGid    float64  `json:"anon_gid"`
	AnonUid    float64  `json:"anon_uid"`
	IpRanges   []string `json:"ip_ranges"`
	SquashMode string   `json:"squash_mode"`
}

type NetworksState struct {
	ConnectMode     string   `json:"connect_mode"`
	IpAddresses     []string `json:"ip_addresses"`
	Modes           []string `json:"modes"`
	Network         string   `json:"network"`
	ReservedIpRange string   `json:"reserved_ip_range"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
