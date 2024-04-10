// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package mssqlvirtualmachinegroup

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

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

type WsfcDomainProfile struct {
	// ClusterBootstrapAccountName: string, optional
	ClusterBootstrapAccountName terra.StringValue `hcl:"cluster_bootstrap_account_name,attr"`
	// ClusterOperatorAccountName: string, optional
	ClusterOperatorAccountName terra.StringValue `hcl:"cluster_operator_account_name,attr"`
	// ClusterSubnetType: string, required
	ClusterSubnetType terra.StringValue `hcl:"cluster_subnet_type,attr" validate:"required"`
	// Fqdn: string, required
	Fqdn terra.StringValue `hcl:"fqdn,attr" validate:"required"`
	// OrganizationalUnitPath: string, optional
	OrganizationalUnitPath terra.StringValue `hcl:"organizational_unit_path,attr"`
	// SqlServiceAccountName: string, optional
	SqlServiceAccountName terra.StringValue `hcl:"sql_service_account_name,attr"`
	// StorageAccountPrimaryKey: string, optional
	StorageAccountPrimaryKey terra.StringValue `hcl:"storage_account_primary_key,attr"`
	// StorageAccountUrl: string, optional
	StorageAccountUrl terra.StringValue `hcl:"storage_account_url,attr"`
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

type WsfcDomainProfileAttributes struct {
	ref terra.Reference
}

func (wdp WsfcDomainProfileAttributes) InternalRef() (terra.Reference, error) {
	return wdp.ref, nil
}

func (wdp WsfcDomainProfileAttributes) InternalWithRef(ref terra.Reference) WsfcDomainProfileAttributes {
	return WsfcDomainProfileAttributes{ref: ref}
}

func (wdp WsfcDomainProfileAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wdp.ref.InternalTokens()
}

func (wdp WsfcDomainProfileAttributes) ClusterBootstrapAccountName() terra.StringValue {
	return terra.ReferenceAsString(wdp.ref.Append("cluster_bootstrap_account_name"))
}

func (wdp WsfcDomainProfileAttributes) ClusterOperatorAccountName() terra.StringValue {
	return terra.ReferenceAsString(wdp.ref.Append("cluster_operator_account_name"))
}

func (wdp WsfcDomainProfileAttributes) ClusterSubnetType() terra.StringValue {
	return terra.ReferenceAsString(wdp.ref.Append("cluster_subnet_type"))
}

func (wdp WsfcDomainProfileAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(wdp.ref.Append("fqdn"))
}

func (wdp WsfcDomainProfileAttributes) OrganizationalUnitPath() terra.StringValue {
	return terra.ReferenceAsString(wdp.ref.Append("organizational_unit_path"))
}

func (wdp WsfcDomainProfileAttributes) SqlServiceAccountName() terra.StringValue {
	return terra.ReferenceAsString(wdp.ref.Append("sql_service_account_name"))
}

func (wdp WsfcDomainProfileAttributes) StorageAccountPrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(wdp.ref.Append("storage_account_primary_key"))
}

func (wdp WsfcDomainProfileAttributes) StorageAccountUrl() terra.StringValue {
	return terra.ReferenceAsString(wdp.ref.Append("storage_account_url"))
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

type WsfcDomainProfileState struct {
	ClusterBootstrapAccountName string `json:"cluster_bootstrap_account_name"`
	ClusterOperatorAccountName  string `json:"cluster_operator_account_name"`
	ClusterSubnetType           string `json:"cluster_subnet_type"`
	Fqdn                        string `json:"fqdn"`
	OrganizationalUnitPath      string `json:"organizational_unit_path"`
	SqlServiceAccountName       string `json:"sql_service_account_name"`
	StorageAccountPrimaryKey    string `json:"storage_account_primary_key"`
	StorageAccountUrl           string `json:"storage_account_url"`
}
