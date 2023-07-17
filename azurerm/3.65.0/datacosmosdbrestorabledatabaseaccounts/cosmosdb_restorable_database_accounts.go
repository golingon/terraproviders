// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datacosmosdbrestorabledatabaseaccounts

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Accounts struct {
	// RestorableLocations: min=0
	RestorableLocations []RestorableLocations `hcl:"restorable_locations,block" validate:"min=0"`
}

type RestorableLocations struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type AccountsAttributes struct {
	ref terra.Reference
}

func (a AccountsAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AccountsAttributes) InternalWithRef(ref terra.Reference) AccountsAttributes {
	return AccountsAttributes{ref: ref}
}

func (a AccountsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AccountsAttributes) ApiType() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("api_type"))
}

func (a AccountsAttributes) CreationTime() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("creation_time"))
}

func (a AccountsAttributes) DeletionTime() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("deletion_time"))
}

func (a AccountsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("id"))
}

func (a AccountsAttributes) RestorableLocations() terra.ListValue[RestorableLocationsAttributes] {
	return terra.ReferenceAsList[RestorableLocationsAttributes](a.ref.Append("restorable_locations"))
}

type RestorableLocationsAttributes struct {
	ref terra.Reference
}

func (rl RestorableLocationsAttributes) InternalRef() (terra.Reference, error) {
	return rl.ref, nil
}

func (rl RestorableLocationsAttributes) InternalWithRef(ref terra.Reference) RestorableLocationsAttributes {
	return RestorableLocationsAttributes{ref: ref}
}

func (rl RestorableLocationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rl.ref.InternalTokens()
}

func (rl RestorableLocationsAttributes) CreationTime() terra.StringValue {
	return terra.ReferenceAsString(rl.ref.Append("creation_time"))
}

func (rl RestorableLocationsAttributes) DeletionTime() terra.StringValue {
	return terra.ReferenceAsString(rl.ref.Append("deletion_time"))
}

func (rl RestorableLocationsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(rl.ref.Append("location"))
}

func (rl RestorableLocationsAttributes) RegionalDatabaseAccountInstanceId() terra.StringValue {
	return terra.ReferenceAsString(rl.ref.Append("regional_database_account_instance_id"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type AccountsState struct {
	ApiType             string                     `json:"api_type"`
	CreationTime        string                     `json:"creation_time"`
	DeletionTime        string                     `json:"deletion_time"`
	Id                  string                     `json:"id"`
	RestorableLocations []RestorableLocationsState `json:"restorable_locations"`
}

type RestorableLocationsState struct {
	CreationTime                      string `json:"creation_time"`
	DeletionTime                      string `json:"deletion_time"`
	Location                          string `json:"location"`
	RegionalDatabaseAccountInstanceId string `json:"regional_database_account_instance_id"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
