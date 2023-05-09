// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dnsrecordset

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type RoutingPolicy struct {
	// EnableGeoFencing: bool, optional
	EnableGeoFencing terra.BoolValue `hcl:"enable_geo_fencing,attr"`
	// Geo: min=0
	Geo []Geo `hcl:"geo,block" validate:"min=0"`
	// PrimaryBackup: optional
	PrimaryBackup *PrimaryBackup `hcl:"primary_backup,block"`
	// Wrr: min=0
	Wrr []Wrr `hcl:"wrr,block" validate:"min=0"`
}

type Geo struct {
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Rrdatas: list of string, optional
	Rrdatas terra.ListValue[terra.StringValue] `hcl:"rrdatas,attr"`
	// GeoHealthCheckedTargets: optional
	HealthCheckedTargets *GeoHealthCheckedTargets `hcl:"health_checked_targets,block"`
}

type GeoHealthCheckedTargets struct {
	// GeoHealthCheckedTargetsInternalLoadBalancers: min=1
	InternalLoadBalancers []GeoHealthCheckedTargetsInternalLoadBalancers `hcl:"internal_load_balancers,block" validate:"min=1"`
}

type GeoHealthCheckedTargetsInternalLoadBalancers struct {
	// IpAddress: string, required
	IpAddress terra.StringValue `hcl:"ip_address,attr" validate:"required"`
	// IpProtocol: string, required
	IpProtocol terra.StringValue `hcl:"ip_protocol,attr" validate:"required"`
	// LoadBalancerType: string, required
	LoadBalancerType terra.StringValue `hcl:"load_balancer_type,attr" validate:"required"`
	// NetworkUrl: string, required
	NetworkUrl terra.StringValue `hcl:"network_url,attr" validate:"required"`
	// Port: string, required
	Port terra.StringValue `hcl:"port,attr" validate:"required"`
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
}

type PrimaryBackup struct {
	// EnableGeoFencingForBackups: bool, optional
	EnableGeoFencingForBackups terra.BoolValue `hcl:"enable_geo_fencing_for_backups,attr"`
	// TrickleRatio: number, optional
	TrickleRatio terra.NumberValue `hcl:"trickle_ratio,attr"`
	// BackupGeo: min=1
	BackupGeo []BackupGeo `hcl:"backup_geo,block" validate:"min=1"`
	// Primary: required
	Primary *Primary `hcl:"primary,block" validate:"required"`
}

type BackupGeo struct {
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Rrdatas: list of string, optional
	Rrdatas terra.ListValue[terra.StringValue] `hcl:"rrdatas,attr"`
	// BackupGeoHealthCheckedTargets: optional
	HealthCheckedTargets *BackupGeoHealthCheckedTargets `hcl:"health_checked_targets,block"`
}

type BackupGeoHealthCheckedTargets struct {
	// BackupGeoHealthCheckedTargetsInternalLoadBalancers: min=1
	InternalLoadBalancers []BackupGeoHealthCheckedTargetsInternalLoadBalancers `hcl:"internal_load_balancers,block" validate:"min=1"`
}

type BackupGeoHealthCheckedTargetsInternalLoadBalancers struct {
	// IpAddress: string, required
	IpAddress terra.StringValue `hcl:"ip_address,attr" validate:"required"`
	// IpProtocol: string, required
	IpProtocol terra.StringValue `hcl:"ip_protocol,attr" validate:"required"`
	// LoadBalancerType: string, required
	LoadBalancerType terra.StringValue `hcl:"load_balancer_type,attr" validate:"required"`
	// NetworkUrl: string, required
	NetworkUrl terra.StringValue `hcl:"network_url,attr" validate:"required"`
	// Port: string, required
	Port terra.StringValue `hcl:"port,attr" validate:"required"`
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
}

type Primary struct {
	// PrimaryInternalLoadBalancers: min=1
	InternalLoadBalancers []PrimaryInternalLoadBalancers `hcl:"internal_load_balancers,block" validate:"min=1"`
}

type PrimaryInternalLoadBalancers struct {
	// IpAddress: string, required
	IpAddress terra.StringValue `hcl:"ip_address,attr" validate:"required"`
	// IpProtocol: string, required
	IpProtocol terra.StringValue `hcl:"ip_protocol,attr" validate:"required"`
	// LoadBalancerType: string, required
	LoadBalancerType terra.StringValue `hcl:"load_balancer_type,attr" validate:"required"`
	// NetworkUrl: string, required
	NetworkUrl terra.StringValue `hcl:"network_url,attr" validate:"required"`
	// Port: string, required
	Port terra.StringValue `hcl:"port,attr" validate:"required"`
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
}

type Wrr struct {
	// Rrdatas: list of string, optional
	Rrdatas terra.ListValue[terra.StringValue] `hcl:"rrdatas,attr"`
	// Weight: number, required
	Weight terra.NumberValue `hcl:"weight,attr" validate:"required"`
	// WrrHealthCheckedTargets: optional
	HealthCheckedTargets *WrrHealthCheckedTargets `hcl:"health_checked_targets,block"`
}

type WrrHealthCheckedTargets struct {
	// WrrHealthCheckedTargetsInternalLoadBalancers: min=1
	InternalLoadBalancers []WrrHealthCheckedTargetsInternalLoadBalancers `hcl:"internal_load_balancers,block" validate:"min=1"`
}

type WrrHealthCheckedTargetsInternalLoadBalancers struct {
	// IpAddress: string, required
	IpAddress terra.StringValue `hcl:"ip_address,attr" validate:"required"`
	// IpProtocol: string, required
	IpProtocol terra.StringValue `hcl:"ip_protocol,attr" validate:"required"`
	// LoadBalancerType: string, required
	LoadBalancerType terra.StringValue `hcl:"load_balancer_type,attr" validate:"required"`
	// NetworkUrl: string, required
	NetworkUrl terra.StringValue `hcl:"network_url,attr" validate:"required"`
	// Port: string, required
	Port terra.StringValue `hcl:"port,attr" validate:"required"`
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
}

type RoutingPolicyAttributes struct {
	ref terra.Reference
}

func (rp RoutingPolicyAttributes) InternalRef() (terra.Reference, error) {
	return rp.ref, nil
}

func (rp RoutingPolicyAttributes) InternalWithRef(ref terra.Reference) RoutingPolicyAttributes {
	return RoutingPolicyAttributes{ref: ref}
}

func (rp RoutingPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rp.ref.InternalTokens()
}

func (rp RoutingPolicyAttributes) EnableGeoFencing() terra.BoolValue {
	return terra.ReferenceAsBool(rp.ref.Append("enable_geo_fencing"))
}

func (rp RoutingPolicyAttributes) Geo() terra.ListValue[GeoAttributes] {
	return terra.ReferenceAsList[GeoAttributes](rp.ref.Append("geo"))
}

func (rp RoutingPolicyAttributes) PrimaryBackup() terra.ListValue[PrimaryBackupAttributes] {
	return terra.ReferenceAsList[PrimaryBackupAttributes](rp.ref.Append("primary_backup"))
}

func (rp RoutingPolicyAttributes) Wrr() terra.ListValue[WrrAttributes] {
	return terra.ReferenceAsList[WrrAttributes](rp.ref.Append("wrr"))
}

type GeoAttributes struct {
	ref terra.Reference
}

func (g GeoAttributes) InternalRef() (terra.Reference, error) {
	return g.ref, nil
}

func (g GeoAttributes) InternalWithRef(ref terra.Reference) GeoAttributes {
	return GeoAttributes{ref: ref}
}

func (g GeoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return g.ref.InternalTokens()
}

func (g GeoAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("location"))
}

func (g GeoAttributes) Rrdatas() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](g.ref.Append("rrdatas"))
}

func (g GeoAttributes) HealthCheckedTargets() terra.ListValue[GeoHealthCheckedTargetsAttributes] {
	return terra.ReferenceAsList[GeoHealthCheckedTargetsAttributes](g.ref.Append("health_checked_targets"))
}

type GeoHealthCheckedTargetsAttributes struct {
	ref terra.Reference
}

func (hct GeoHealthCheckedTargetsAttributes) InternalRef() (terra.Reference, error) {
	return hct.ref, nil
}

func (hct GeoHealthCheckedTargetsAttributes) InternalWithRef(ref terra.Reference) GeoHealthCheckedTargetsAttributes {
	return GeoHealthCheckedTargetsAttributes{ref: ref}
}

func (hct GeoHealthCheckedTargetsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hct.ref.InternalTokens()
}

func (hct GeoHealthCheckedTargetsAttributes) InternalLoadBalancers() terra.ListValue[GeoHealthCheckedTargetsInternalLoadBalancersAttributes] {
	return terra.ReferenceAsList[GeoHealthCheckedTargetsInternalLoadBalancersAttributes](hct.ref.Append("internal_load_balancers"))
}

type GeoHealthCheckedTargetsInternalLoadBalancersAttributes struct {
	ref terra.Reference
}

func (ilb GeoHealthCheckedTargetsInternalLoadBalancersAttributes) InternalRef() (terra.Reference, error) {
	return ilb.ref, nil
}

func (ilb GeoHealthCheckedTargetsInternalLoadBalancersAttributes) InternalWithRef(ref terra.Reference) GeoHealthCheckedTargetsInternalLoadBalancersAttributes {
	return GeoHealthCheckedTargetsInternalLoadBalancersAttributes{ref: ref}
}

func (ilb GeoHealthCheckedTargetsInternalLoadBalancersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ilb.ref.InternalTokens()
}

func (ilb GeoHealthCheckedTargetsInternalLoadBalancersAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("ip_address"))
}

func (ilb GeoHealthCheckedTargetsInternalLoadBalancersAttributes) IpProtocol() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("ip_protocol"))
}

func (ilb GeoHealthCheckedTargetsInternalLoadBalancersAttributes) LoadBalancerType() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("load_balancer_type"))
}

func (ilb GeoHealthCheckedTargetsInternalLoadBalancersAttributes) NetworkUrl() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("network_url"))
}

func (ilb GeoHealthCheckedTargetsInternalLoadBalancersAttributes) Port() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("port"))
}

func (ilb GeoHealthCheckedTargetsInternalLoadBalancersAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("project"))
}

func (ilb GeoHealthCheckedTargetsInternalLoadBalancersAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("region"))
}

type PrimaryBackupAttributes struct {
	ref terra.Reference
}

func (pb PrimaryBackupAttributes) InternalRef() (terra.Reference, error) {
	return pb.ref, nil
}

func (pb PrimaryBackupAttributes) InternalWithRef(ref terra.Reference) PrimaryBackupAttributes {
	return PrimaryBackupAttributes{ref: ref}
}

func (pb PrimaryBackupAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pb.ref.InternalTokens()
}

func (pb PrimaryBackupAttributes) EnableGeoFencingForBackups() terra.BoolValue {
	return terra.ReferenceAsBool(pb.ref.Append("enable_geo_fencing_for_backups"))
}

func (pb PrimaryBackupAttributes) TrickleRatio() terra.NumberValue {
	return terra.ReferenceAsNumber(pb.ref.Append("trickle_ratio"))
}

func (pb PrimaryBackupAttributes) BackupGeo() terra.ListValue[BackupGeoAttributes] {
	return terra.ReferenceAsList[BackupGeoAttributes](pb.ref.Append("backup_geo"))
}

func (pb PrimaryBackupAttributes) Primary() terra.ListValue[PrimaryAttributes] {
	return terra.ReferenceAsList[PrimaryAttributes](pb.ref.Append("primary"))
}

type BackupGeoAttributes struct {
	ref terra.Reference
}

func (bg BackupGeoAttributes) InternalRef() (terra.Reference, error) {
	return bg.ref, nil
}

func (bg BackupGeoAttributes) InternalWithRef(ref terra.Reference) BackupGeoAttributes {
	return BackupGeoAttributes{ref: ref}
}

func (bg BackupGeoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bg.ref.InternalTokens()
}

func (bg BackupGeoAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bg.ref.Append("location"))
}

func (bg BackupGeoAttributes) Rrdatas() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](bg.ref.Append("rrdatas"))
}

func (bg BackupGeoAttributes) HealthCheckedTargets() terra.ListValue[BackupGeoHealthCheckedTargetsAttributes] {
	return terra.ReferenceAsList[BackupGeoHealthCheckedTargetsAttributes](bg.ref.Append("health_checked_targets"))
}

type BackupGeoHealthCheckedTargetsAttributes struct {
	ref terra.Reference
}

func (hct BackupGeoHealthCheckedTargetsAttributes) InternalRef() (terra.Reference, error) {
	return hct.ref, nil
}

func (hct BackupGeoHealthCheckedTargetsAttributes) InternalWithRef(ref terra.Reference) BackupGeoHealthCheckedTargetsAttributes {
	return BackupGeoHealthCheckedTargetsAttributes{ref: ref}
}

func (hct BackupGeoHealthCheckedTargetsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hct.ref.InternalTokens()
}

func (hct BackupGeoHealthCheckedTargetsAttributes) InternalLoadBalancers() terra.ListValue[BackupGeoHealthCheckedTargetsInternalLoadBalancersAttributes] {
	return terra.ReferenceAsList[BackupGeoHealthCheckedTargetsInternalLoadBalancersAttributes](hct.ref.Append("internal_load_balancers"))
}

type BackupGeoHealthCheckedTargetsInternalLoadBalancersAttributes struct {
	ref terra.Reference
}

func (ilb BackupGeoHealthCheckedTargetsInternalLoadBalancersAttributes) InternalRef() (terra.Reference, error) {
	return ilb.ref, nil
}

func (ilb BackupGeoHealthCheckedTargetsInternalLoadBalancersAttributes) InternalWithRef(ref terra.Reference) BackupGeoHealthCheckedTargetsInternalLoadBalancersAttributes {
	return BackupGeoHealthCheckedTargetsInternalLoadBalancersAttributes{ref: ref}
}

func (ilb BackupGeoHealthCheckedTargetsInternalLoadBalancersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ilb.ref.InternalTokens()
}

func (ilb BackupGeoHealthCheckedTargetsInternalLoadBalancersAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("ip_address"))
}

func (ilb BackupGeoHealthCheckedTargetsInternalLoadBalancersAttributes) IpProtocol() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("ip_protocol"))
}

func (ilb BackupGeoHealthCheckedTargetsInternalLoadBalancersAttributes) LoadBalancerType() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("load_balancer_type"))
}

func (ilb BackupGeoHealthCheckedTargetsInternalLoadBalancersAttributes) NetworkUrl() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("network_url"))
}

func (ilb BackupGeoHealthCheckedTargetsInternalLoadBalancersAttributes) Port() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("port"))
}

func (ilb BackupGeoHealthCheckedTargetsInternalLoadBalancersAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("project"))
}

func (ilb BackupGeoHealthCheckedTargetsInternalLoadBalancersAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("region"))
}

type PrimaryAttributes struct {
	ref terra.Reference
}

func (p PrimaryAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PrimaryAttributes) InternalWithRef(ref terra.Reference) PrimaryAttributes {
	return PrimaryAttributes{ref: ref}
}

func (p PrimaryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PrimaryAttributes) InternalLoadBalancers() terra.ListValue[PrimaryInternalLoadBalancersAttributes] {
	return terra.ReferenceAsList[PrimaryInternalLoadBalancersAttributes](p.ref.Append("internal_load_balancers"))
}

type PrimaryInternalLoadBalancersAttributes struct {
	ref terra.Reference
}

func (ilb PrimaryInternalLoadBalancersAttributes) InternalRef() (terra.Reference, error) {
	return ilb.ref, nil
}

func (ilb PrimaryInternalLoadBalancersAttributes) InternalWithRef(ref terra.Reference) PrimaryInternalLoadBalancersAttributes {
	return PrimaryInternalLoadBalancersAttributes{ref: ref}
}

func (ilb PrimaryInternalLoadBalancersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ilb.ref.InternalTokens()
}

func (ilb PrimaryInternalLoadBalancersAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("ip_address"))
}

func (ilb PrimaryInternalLoadBalancersAttributes) IpProtocol() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("ip_protocol"))
}

func (ilb PrimaryInternalLoadBalancersAttributes) LoadBalancerType() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("load_balancer_type"))
}

func (ilb PrimaryInternalLoadBalancersAttributes) NetworkUrl() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("network_url"))
}

func (ilb PrimaryInternalLoadBalancersAttributes) Port() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("port"))
}

func (ilb PrimaryInternalLoadBalancersAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("project"))
}

func (ilb PrimaryInternalLoadBalancersAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("region"))
}

type WrrAttributes struct {
	ref terra.Reference
}

func (w WrrAttributes) InternalRef() (terra.Reference, error) {
	return w.ref, nil
}

func (w WrrAttributes) InternalWithRef(ref terra.Reference) WrrAttributes {
	return WrrAttributes{ref: ref}
}

func (w WrrAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return w.ref.InternalTokens()
}

func (w WrrAttributes) Rrdatas() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](w.ref.Append("rrdatas"))
}

func (w WrrAttributes) Weight() terra.NumberValue {
	return terra.ReferenceAsNumber(w.ref.Append("weight"))
}

func (w WrrAttributes) HealthCheckedTargets() terra.ListValue[WrrHealthCheckedTargetsAttributes] {
	return terra.ReferenceAsList[WrrHealthCheckedTargetsAttributes](w.ref.Append("health_checked_targets"))
}

type WrrHealthCheckedTargetsAttributes struct {
	ref terra.Reference
}

func (hct WrrHealthCheckedTargetsAttributes) InternalRef() (terra.Reference, error) {
	return hct.ref, nil
}

func (hct WrrHealthCheckedTargetsAttributes) InternalWithRef(ref terra.Reference) WrrHealthCheckedTargetsAttributes {
	return WrrHealthCheckedTargetsAttributes{ref: ref}
}

func (hct WrrHealthCheckedTargetsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hct.ref.InternalTokens()
}

func (hct WrrHealthCheckedTargetsAttributes) InternalLoadBalancers() terra.ListValue[WrrHealthCheckedTargetsInternalLoadBalancersAttributes] {
	return terra.ReferenceAsList[WrrHealthCheckedTargetsInternalLoadBalancersAttributes](hct.ref.Append("internal_load_balancers"))
}

type WrrHealthCheckedTargetsInternalLoadBalancersAttributes struct {
	ref terra.Reference
}

func (ilb WrrHealthCheckedTargetsInternalLoadBalancersAttributes) InternalRef() (terra.Reference, error) {
	return ilb.ref, nil
}

func (ilb WrrHealthCheckedTargetsInternalLoadBalancersAttributes) InternalWithRef(ref terra.Reference) WrrHealthCheckedTargetsInternalLoadBalancersAttributes {
	return WrrHealthCheckedTargetsInternalLoadBalancersAttributes{ref: ref}
}

func (ilb WrrHealthCheckedTargetsInternalLoadBalancersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ilb.ref.InternalTokens()
}

func (ilb WrrHealthCheckedTargetsInternalLoadBalancersAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("ip_address"))
}

func (ilb WrrHealthCheckedTargetsInternalLoadBalancersAttributes) IpProtocol() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("ip_protocol"))
}

func (ilb WrrHealthCheckedTargetsInternalLoadBalancersAttributes) LoadBalancerType() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("load_balancer_type"))
}

func (ilb WrrHealthCheckedTargetsInternalLoadBalancersAttributes) NetworkUrl() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("network_url"))
}

func (ilb WrrHealthCheckedTargetsInternalLoadBalancersAttributes) Port() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("port"))
}

func (ilb WrrHealthCheckedTargetsInternalLoadBalancersAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("project"))
}

func (ilb WrrHealthCheckedTargetsInternalLoadBalancersAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(ilb.ref.Append("region"))
}

type RoutingPolicyState struct {
	EnableGeoFencing bool                 `json:"enable_geo_fencing"`
	Geo              []GeoState           `json:"geo"`
	PrimaryBackup    []PrimaryBackupState `json:"primary_backup"`
	Wrr              []WrrState           `json:"wrr"`
}

type GeoState struct {
	Location             string                         `json:"location"`
	Rrdatas              []string                       `json:"rrdatas"`
	HealthCheckedTargets []GeoHealthCheckedTargetsState `json:"health_checked_targets"`
}

type GeoHealthCheckedTargetsState struct {
	InternalLoadBalancers []GeoHealthCheckedTargetsInternalLoadBalancersState `json:"internal_load_balancers"`
}

type GeoHealthCheckedTargetsInternalLoadBalancersState struct {
	IpAddress        string `json:"ip_address"`
	IpProtocol       string `json:"ip_protocol"`
	LoadBalancerType string `json:"load_balancer_type"`
	NetworkUrl       string `json:"network_url"`
	Port             string `json:"port"`
	Project          string `json:"project"`
	Region           string `json:"region"`
}

type PrimaryBackupState struct {
	EnableGeoFencingForBackups bool             `json:"enable_geo_fencing_for_backups"`
	TrickleRatio               float64          `json:"trickle_ratio"`
	BackupGeo                  []BackupGeoState `json:"backup_geo"`
	Primary                    []PrimaryState   `json:"primary"`
}

type BackupGeoState struct {
	Location             string                               `json:"location"`
	Rrdatas              []string                             `json:"rrdatas"`
	HealthCheckedTargets []BackupGeoHealthCheckedTargetsState `json:"health_checked_targets"`
}

type BackupGeoHealthCheckedTargetsState struct {
	InternalLoadBalancers []BackupGeoHealthCheckedTargetsInternalLoadBalancersState `json:"internal_load_balancers"`
}

type BackupGeoHealthCheckedTargetsInternalLoadBalancersState struct {
	IpAddress        string `json:"ip_address"`
	IpProtocol       string `json:"ip_protocol"`
	LoadBalancerType string `json:"load_balancer_type"`
	NetworkUrl       string `json:"network_url"`
	Port             string `json:"port"`
	Project          string `json:"project"`
	Region           string `json:"region"`
}

type PrimaryState struct {
	InternalLoadBalancers []PrimaryInternalLoadBalancersState `json:"internal_load_balancers"`
}

type PrimaryInternalLoadBalancersState struct {
	IpAddress        string `json:"ip_address"`
	IpProtocol       string `json:"ip_protocol"`
	LoadBalancerType string `json:"load_balancer_type"`
	NetworkUrl       string `json:"network_url"`
	Port             string `json:"port"`
	Project          string `json:"project"`
	Region           string `json:"region"`
}

type WrrState struct {
	Rrdatas              []string                       `json:"rrdatas"`
	Weight               float64                        `json:"weight"`
	HealthCheckedTargets []WrrHealthCheckedTargetsState `json:"health_checked_targets"`
}

type WrrHealthCheckedTargetsState struct {
	InternalLoadBalancers []WrrHealthCheckedTargetsInternalLoadBalancersState `json:"internal_load_balancers"`
}

type WrrHealthCheckedTargetsInternalLoadBalancersState struct {
	IpAddress        string `json:"ip_address"`
	IpProtocol       string `json:"ip_protocol"`
	LoadBalancerType string `json:"load_balancer_type"`
	NetworkUrl       string `json:"network_url"`
	Port             string `json:"port"`
	Project          string `json:"project"`
	Region           string `json:"region"`
}
