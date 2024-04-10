// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package maintenanceconfiguration

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type InstallPatches struct {
	// Reboot: string, optional
	Reboot terra.StringValue `hcl:"reboot,attr"`
	// Linux: min=0
	Linux []Linux `hcl:"linux,block" validate:"min=0"`
	// Windows: min=0
	Windows []Windows `hcl:"windows,block" validate:"min=0"`
}

type Linux struct {
	// ClassificationsToInclude: list of string, optional
	ClassificationsToInclude terra.ListValue[terra.StringValue] `hcl:"classifications_to_include,attr"`
	// PackageNamesMaskToExclude: list of string, optional
	PackageNamesMaskToExclude terra.ListValue[terra.StringValue] `hcl:"package_names_mask_to_exclude,attr"`
	// PackageNamesMaskToInclude: list of string, optional
	PackageNamesMaskToInclude terra.ListValue[terra.StringValue] `hcl:"package_names_mask_to_include,attr"`
}

type Windows struct {
	// ClassificationsToInclude: list of string, optional
	ClassificationsToInclude terra.ListValue[terra.StringValue] `hcl:"classifications_to_include,attr"`
	// KbNumbersToExclude: list of string, optional
	KbNumbersToExclude terra.ListValue[terra.StringValue] `hcl:"kb_numbers_to_exclude,attr"`
	// KbNumbersToInclude: list of string, optional
	KbNumbersToInclude terra.ListValue[terra.StringValue] `hcl:"kb_numbers_to_include,attr"`
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

type Window struct {
	// Duration: string, optional
	Duration terra.StringValue `hcl:"duration,attr"`
	// ExpirationDateTime: string, optional
	ExpirationDateTime terra.StringValue `hcl:"expiration_date_time,attr"`
	// RecurEvery: string, optional
	RecurEvery terra.StringValue `hcl:"recur_every,attr"`
	// StartDateTime: string, required
	StartDateTime terra.StringValue `hcl:"start_date_time,attr" validate:"required"`
	// TimeZone: string, required
	TimeZone terra.StringValue `hcl:"time_zone,attr" validate:"required"`
}

type InstallPatchesAttributes struct {
	ref terra.Reference
}

func (ip InstallPatchesAttributes) InternalRef() (terra.Reference, error) {
	return ip.ref, nil
}

func (ip InstallPatchesAttributes) InternalWithRef(ref terra.Reference) InstallPatchesAttributes {
	return InstallPatchesAttributes{ref: ref}
}

func (ip InstallPatchesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ip.ref.InternalTokens()
}

func (ip InstallPatchesAttributes) Reboot() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("reboot"))
}

func (ip InstallPatchesAttributes) Linux() terra.ListValue[LinuxAttributes] {
	return terra.ReferenceAsList[LinuxAttributes](ip.ref.Append("linux"))
}

func (ip InstallPatchesAttributes) Windows() terra.ListValue[WindowsAttributes] {
	return terra.ReferenceAsList[WindowsAttributes](ip.ref.Append("windows"))
}

type LinuxAttributes struct {
	ref terra.Reference
}

func (l LinuxAttributes) InternalRef() (terra.Reference, error) {
	return l.ref, nil
}

func (l LinuxAttributes) InternalWithRef(ref terra.Reference) LinuxAttributes {
	return LinuxAttributes{ref: ref}
}

func (l LinuxAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return l.ref.InternalTokens()
}

func (l LinuxAttributes) ClassificationsToInclude() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](l.ref.Append("classifications_to_include"))
}

func (l LinuxAttributes) PackageNamesMaskToExclude() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](l.ref.Append("package_names_mask_to_exclude"))
}

func (l LinuxAttributes) PackageNamesMaskToInclude() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](l.ref.Append("package_names_mask_to_include"))
}

type WindowsAttributes struct {
	ref terra.Reference
}

func (w WindowsAttributes) InternalRef() (terra.Reference, error) {
	return w.ref, nil
}

func (w WindowsAttributes) InternalWithRef(ref terra.Reference) WindowsAttributes {
	return WindowsAttributes{ref: ref}
}

func (w WindowsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return w.ref.InternalTokens()
}

func (w WindowsAttributes) ClassificationsToInclude() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](w.ref.Append("classifications_to_include"))
}

func (w WindowsAttributes) KbNumbersToExclude() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](w.ref.Append("kb_numbers_to_exclude"))
}

func (w WindowsAttributes) KbNumbersToInclude() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](w.ref.Append("kb_numbers_to_include"))
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

type WindowAttributes struct {
	ref terra.Reference
}

func (w WindowAttributes) InternalRef() (terra.Reference, error) {
	return w.ref, nil
}

func (w WindowAttributes) InternalWithRef(ref terra.Reference) WindowAttributes {
	return WindowAttributes{ref: ref}
}

func (w WindowAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return w.ref.InternalTokens()
}

func (w WindowAttributes) Duration() terra.StringValue {
	return terra.ReferenceAsString(w.ref.Append("duration"))
}

func (w WindowAttributes) ExpirationDateTime() terra.StringValue {
	return terra.ReferenceAsString(w.ref.Append("expiration_date_time"))
}

func (w WindowAttributes) RecurEvery() terra.StringValue {
	return terra.ReferenceAsString(w.ref.Append("recur_every"))
}

func (w WindowAttributes) StartDateTime() terra.StringValue {
	return terra.ReferenceAsString(w.ref.Append("start_date_time"))
}

func (w WindowAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(w.ref.Append("time_zone"))
}

type InstallPatchesState struct {
	Reboot  string         `json:"reboot"`
	Linux   []LinuxState   `json:"linux"`
	Windows []WindowsState `json:"windows"`
}

type LinuxState struct {
	ClassificationsToInclude  []string `json:"classifications_to_include"`
	PackageNamesMaskToExclude []string `json:"package_names_mask_to_exclude"`
	PackageNamesMaskToInclude []string `json:"package_names_mask_to_include"`
}

type WindowsState struct {
	ClassificationsToInclude []string `json:"classifications_to_include"`
	KbNumbersToExclude       []string `json:"kb_numbers_to_exclude"`
	KbNumbersToInclude       []string `json:"kb_numbers_to_include"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

type WindowState struct {
	Duration           string `json:"duration"`
	ExpirationDateTime string `json:"expiration_date_time"`
	RecurEvery         string `json:"recur_every"`
	StartDateTime      string `json:"start_date_time"`
	TimeZone           string `json:"time_zone"`
}
