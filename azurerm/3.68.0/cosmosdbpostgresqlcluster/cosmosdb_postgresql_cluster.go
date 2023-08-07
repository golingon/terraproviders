// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cosmosdbpostgresqlcluster

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type MaintenanceWindow struct {
	// DayOfWeek: number, optional
	DayOfWeek terra.NumberValue `hcl:"day_of_week,attr"`
	// StartHour: number, optional
	StartHour terra.NumberValue `hcl:"start_hour,attr"`
	// StartMinute: number, optional
	StartMinute terra.NumberValue `hcl:"start_minute,attr"`
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

type MaintenanceWindowAttributes struct {
	ref terra.Reference
}

func (mw MaintenanceWindowAttributes) InternalRef() (terra.Reference, error) {
	return mw.ref, nil
}

func (mw MaintenanceWindowAttributes) InternalWithRef(ref terra.Reference) MaintenanceWindowAttributes {
	return MaintenanceWindowAttributes{ref: ref}
}

func (mw MaintenanceWindowAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mw.ref.InternalTokens()
}

func (mw MaintenanceWindowAttributes) DayOfWeek() terra.NumberValue {
	return terra.ReferenceAsNumber(mw.ref.Append("day_of_week"))
}

func (mw MaintenanceWindowAttributes) StartHour() terra.NumberValue {
	return terra.ReferenceAsNumber(mw.ref.Append("start_hour"))
}

func (mw MaintenanceWindowAttributes) StartMinute() terra.NumberValue {
	return terra.ReferenceAsNumber(mw.ref.Append("start_minute"))
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

type MaintenanceWindowState struct {
	DayOfWeek   float64 `json:"day_of_week"`
	StartHour   float64 `json:"start_hour"`
	StartMinute float64 `json:"start_minute"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
