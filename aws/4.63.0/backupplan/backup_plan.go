// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package backupplan

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AdvancedBackupSetting struct {
	// BackupOptions: map of string, required
	BackupOptions terra.MapValue[terra.StringValue] `hcl:"backup_options,attr" validate:"required"`
	// ResourceType: string, required
	ResourceType terra.StringValue `hcl:"resource_type,attr" validate:"required"`
}

type Rule struct {
	// CompletionWindow: number, optional
	CompletionWindow terra.NumberValue `hcl:"completion_window,attr"`
	// EnableContinuousBackup: bool, optional
	EnableContinuousBackup terra.BoolValue `hcl:"enable_continuous_backup,attr"`
	// RecoveryPointTags: map of string, optional
	RecoveryPointTags terra.MapValue[terra.StringValue] `hcl:"recovery_point_tags,attr"`
	// RuleName: string, required
	RuleName terra.StringValue `hcl:"rule_name,attr" validate:"required"`
	// Schedule: string, optional
	Schedule terra.StringValue `hcl:"schedule,attr"`
	// StartWindow: number, optional
	StartWindow terra.NumberValue `hcl:"start_window,attr"`
	// TargetVaultName: string, required
	TargetVaultName terra.StringValue `hcl:"target_vault_name,attr" validate:"required"`
	// CopyAction: min=0
	CopyAction []CopyAction `hcl:"copy_action,block" validate:"min=0"`
	// RuleLifecycle: optional
	Lifecycle *RuleLifecycle `hcl:"lifecycle,block"`
}

type CopyAction struct {
	// DestinationVaultArn: string, required
	DestinationVaultArn terra.StringValue `hcl:"destination_vault_arn,attr" validate:"required"`
	// CopyActionLifecycle: optional
	Lifecycle *CopyActionLifecycle `hcl:"lifecycle,block"`
}

type CopyActionLifecycle struct {
	// ColdStorageAfter: number, optional
	ColdStorageAfter terra.NumberValue `hcl:"cold_storage_after,attr"`
	// DeleteAfter: number, optional
	DeleteAfter terra.NumberValue `hcl:"delete_after,attr"`
}

type RuleLifecycle struct {
	// ColdStorageAfter: number, optional
	ColdStorageAfter terra.NumberValue `hcl:"cold_storage_after,attr"`
	// DeleteAfter: number, optional
	DeleteAfter terra.NumberValue `hcl:"delete_after,attr"`
}

type AdvancedBackupSettingAttributes struct {
	ref terra.Reference
}

func (abs AdvancedBackupSettingAttributes) InternalRef() (terra.Reference, error) {
	return abs.ref, nil
}

func (abs AdvancedBackupSettingAttributes) InternalWithRef(ref terra.Reference) AdvancedBackupSettingAttributes {
	return AdvancedBackupSettingAttributes{ref: ref}
}

func (abs AdvancedBackupSettingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return abs.ref.InternalTokens()
}

func (abs AdvancedBackupSettingAttributes) BackupOptions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](abs.ref.Append("backup_options"))
}

func (abs AdvancedBackupSettingAttributes) ResourceType() terra.StringValue {
	return terra.ReferenceAsString(abs.ref.Append("resource_type"))
}

type RuleAttributes struct {
	ref terra.Reference
}

func (r RuleAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RuleAttributes) InternalWithRef(ref terra.Reference) RuleAttributes {
	return RuleAttributes{ref: ref}
}

func (r RuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RuleAttributes) CompletionWindow() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("completion_window"))
}

func (r RuleAttributes) EnableContinuousBackup() terra.BoolValue {
	return terra.ReferenceAsBool(r.ref.Append("enable_continuous_backup"))
}

func (r RuleAttributes) RecoveryPointTags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](r.ref.Append("recovery_point_tags"))
}

func (r RuleAttributes) RuleName() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("rule_name"))
}

func (r RuleAttributes) Schedule() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("schedule"))
}

func (r RuleAttributes) StartWindow() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("start_window"))
}

func (r RuleAttributes) TargetVaultName() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("target_vault_name"))
}

func (r RuleAttributes) CopyAction() terra.SetValue[CopyActionAttributes] {
	return terra.ReferenceAsSet[CopyActionAttributes](r.ref.Append("copy_action"))
}

func (r RuleAttributes) Lifecycle() terra.ListValue[RuleLifecycleAttributes] {
	return terra.ReferenceAsList[RuleLifecycleAttributes](r.ref.Append("lifecycle"))
}

type CopyActionAttributes struct {
	ref terra.Reference
}

func (ca CopyActionAttributes) InternalRef() (terra.Reference, error) {
	return ca.ref, nil
}

func (ca CopyActionAttributes) InternalWithRef(ref terra.Reference) CopyActionAttributes {
	return CopyActionAttributes{ref: ref}
}

func (ca CopyActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ca.ref.InternalTokens()
}

func (ca CopyActionAttributes) DestinationVaultArn() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("destination_vault_arn"))
}

func (ca CopyActionAttributes) Lifecycle() terra.ListValue[CopyActionLifecycleAttributes] {
	return terra.ReferenceAsList[CopyActionLifecycleAttributes](ca.ref.Append("lifecycle"))
}

type CopyActionLifecycleAttributes struct {
	ref terra.Reference
}

func (l CopyActionLifecycleAttributes) InternalRef() (terra.Reference, error) {
	return l.ref, nil
}

func (l CopyActionLifecycleAttributes) InternalWithRef(ref terra.Reference) CopyActionLifecycleAttributes {
	return CopyActionLifecycleAttributes{ref: ref}
}

func (l CopyActionLifecycleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return l.ref.InternalTokens()
}

func (l CopyActionLifecycleAttributes) ColdStorageAfter() terra.NumberValue {
	return terra.ReferenceAsNumber(l.ref.Append("cold_storage_after"))
}

func (l CopyActionLifecycleAttributes) DeleteAfter() terra.NumberValue {
	return terra.ReferenceAsNumber(l.ref.Append("delete_after"))
}

type RuleLifecycleAttributes struct {
	ref terra.Reference
}

func (l RuleLifecycleAttributes) InternalRef() (terra.Reference, error) {
	return l.ref, nil
}

func (l RuleLifecycleAttributes) InternalWithRef(ref terra.Reference) RuleLifecycleAttributes {
	return RuleLifecycleAttributes{ref: ref}
}

func (l RuleLifecycleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return l.ref.InternalTokens()
}

func (l RuleLifecycleAttributes) ColdStorageAfter() terra.NumberValue {
	return terra.ReferenceAsNumber(l.ref.Append("cold_storage_after"))
}

func (l RuleLifecycleAttributes) DeleteAfter() terra.NumberValue {
	return terra.ReferenceAsNumber(l.ref.Append("delete_after"))
}

type AdvancedBackupSettingState struct {
	BackupOptions map[string]string `json:"backup_options"`
	ResourceType  string            `json:"resource_type"`
}

type RuleState struct {
	CompletionWindow       float64              `json:"completion_window"`
	EnableContinuousBackup bool                 `json:"enable_continuous_backup"`
	RecoveryPointTags      map[string]string    `json:"recovery_point_tags"`
	RuleName               string               `json:"rule_name"`
	Schedule               string               `json:"schedule"`
	StartWindow            float64              `json:"start_window"`
	TargetVaultName        string               `json:"target_vault_name"`
	CopyAction             []CopyActionState    `json:"copy_action"`
	Lifecycle              []RuleLifecycleState `json:"lifecycle"`
}

type CopyActionState struct {
	DestinationVaultArn string                     `json:"destination_vault_arn"`
	Lifecycle           []CopyActionLifecycleState `json:"lifecycle"`
}

type CopyActionLifecycleState struct {
	ColdStorageAfter float64 `json:"cold_storage_after"`
	DeleteAfter      float64 `json:"delete_after"`
}

type RuleLifecycleState struct {
	ColdStorageAfter float64 `json:"cold_storage_after"`
	DeleteAfter      float64 `json:"delete_after"`
}
