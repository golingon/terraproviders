// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package macie2classificationjob

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type UserPausedDetails struct{}

type S3JobDefinition struct {
	// BucketCriteria: optional
	BucketCriteria *BucketCriteria `hcl:"bucket_criteria,block"`
	// BucketDefinitions: min=0
	BucketDefinitions []BucketDefinitions `hcl:"bucket_definitions,block" validate:"min=0"`
	// Scoping: optional
	Scoping *Scoping `hcl:"scoping,block"`
}

type BucketCriteria struct {
	// BucketCriteriaExcludes: optional
	Excludes *BucketCriteriaExcludes `hcl:"excludes,block"`
	// BucketCriteriaIncludes: optional
	Includes *BucketCriteriaIncludes `hcl:"includes,block"`
}

type BucketCriteriaExcludes struct {
	// BucketCriteriaExcludesAnd: min=0
	And []BucketCriteriaExcludesAnd `hcl:"and,block" validate:"min=0"`
}

type BucketCriteriaExcludesAnd struct {
	// ExcludesAndSimpleCriterion: optional
	SimpleCriterion *ExcludesAndSimpleCriterion `hcl:"simple_criterion,block"`
	// ExcludesAndTagCriterion: optional
	TagCriterion *ExcludesAndTagCriterion `hcl:"tag_criterion,block"`
}

type ExcludesAndSimpleCriterion struct {
	// Comparator: string, optional
	Comparator terra.StringValue `hcl:"comparator,attr"`
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// Values: list of string, optional
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr"`
}

type ExcludesAndTagCriterion struct {
	// Comparator: string, optional
	Comparator terra.StringValue `hcl:"comparator,attr"`
	// ExcludesAndTagCriterionTagValues: min=0
	TagValues []ExcludesAndTagCriterionTagValues `hcl:"tag_values,block" validate:"min=0"`
}

type ExcludesAndTagCriterionTagValues struct {
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type BucketCriteriaIncludes struct {
	// BucketCriteriaIncludesAnd: min=0
	And []BucketCriteriaIncludesAnd `hcl:"and,block" validate:"min=0"`
}

type BucketCriteriaIncludesAnd struct {
	// IncludesAndSimpleCriterion: optional
	SimpleCriterion *IncludesAndSimpleCriterion `hcl:"simple_criterion,block"`
	// IncludesAndTagCriterion: optional
	TagCriterion *IncludesAndTagCriterion `hcl:"tag_criterion,block"`
}

type IncludesAndSimpleCriterion struct {
	// Comparator: string, optional
	Comparator terra.StringValue `hcl:"comparator,attr"`
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// Values: list of string, optional
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr"`
}

type IncludesAndTagCriterion struct {
	// Comparator: string, optional
	Comparator terra.StringValue `hcl:"comparator,attr"`
	// IncludesAndTagCriterionTagValues: min=0
	TagValues []IncludesAndTagCriterionTagValues `hcl:"tag_values,block" validate:"min=0"`
}

type IncludesAndTagCriterionTagValues struct {
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type BucketDefinitions struct {
	// AccountId: string, required
	AccountId terra.StringValue `hcl:"account_id,attr" validate:"required"`
	// Buckets: list of string, required
	Buckets terra.ListValue[terra.StringValue] `hcl:"buckets,attr" validate:"required"`
}

type Scoping struct {
	// ScopingExcludes: optional
	Excludes *ScopingExcludes `hcl:"excludes,block"`
	// ScopingIncludes: optional
	Includes *ScopingIncludes `hcl:"includes,block"`
}

type ScopingExcludes struct {
	// ScopingExcludesAnd: min=0
	And []ScopingExcludesAnd `hcl:"and,block" validate:"min=0"`
}

type ScopingExcludesAnd struct {
	// ExcludesAndSimpleScopeTerm: optional
	SimpleScopeTerm *ExcludesAndSimpleScopeTerm `hcl:"simple_scope_term,block"`
	// ExcludesAndTagScopeTerm: optional
	TagScopeTerm *ExcludesAndTagScopeTerm `hcl:"tag_scope_term,block"`
}

type ExcludesAndSimpleScopeTerm struct {
	// Comparator: string, optional
	Comparator terra.StringValue `hcl:"comparator,attr"`
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// Values: list of string, optional
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr"`
}

type ExcludesAndTagScopeTerm struct {
	// Comparator: string, optional
	Comparator terra.StringValue `hcl:"comparator,attr"`
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// Target: string, optional
	Target terra.StringValue `hcl:"target,attr"`
	// ExcludesAndTagScopeTermTagValues: min=0
	TagValues []ExcludesAndTagScopeTermTagValues `hcl:"tag_values,block" validate:"min=0"`
}

type ExcludesAndTagScopeTermTagValues struct {
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type ScopingIncludes struct {
	// ScopingIncludesAnd: min=0
	And []ScopingIncludesAnd `hcl:"and,block" validate:"min=0"`
}

type ScopingIncludesAnd struct {
	// IncludesAndSimpleScopeTerm: optional
	SimpleScopeTerm *IncludesAndSimpleScopeTerm `hcl:"simple_scope_term,block"`
	// IncludesAndTagScopeTerm: optional
	TagScopeTerm *IncludesAndTagScopeTerm `hcl:"tag_scope_term,block"`
}

type IncludesAndSimpleScopeTerm struct {
	// Comparator: string, optional
	Comparator terra.StringValue `hcl:"comparator,attr"`
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// Values: list of string, optional
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr"`
}

type IncludesAndTagScopeTerm struct {
	// Comparator: string, optional
	Comparator terra.StringValue `hcl:"comparator,attr"`
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// Target: string, optional
	Target terra.StringValue `hcl:"target,attr"`
	// IncludesAndTagScopeTermTagValues: min=0
	TagValues []IncludesAndTagScopeTermTagValues `hcl:"tag_values,block" validate:"min=0"`
}

type IncludesAndTagScopeTermTagValues struct {
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
}

type ScheduleFrequency struct {
	// DailySchedule: bool, optional
	DailySchedule terra.BoolValue `hcl:"daily_schedule,attr"`
	// MonthlySchedule: number, optional
	MonthlySchedule terra.NumberValue `hcl:"monthly_schedule,attr"`
	// WeeklySchedule: string, optional
	WeeklySchedule terra.StringValue `hcl:"weekly_schedule,attr"`
}

type UserPausedDetailsAttributes struct {
	ref terra.Reference
}

func (upd UserPausedDetailsAttributes) InternalRef() (terra.Reference, error) {
	return upd.ref, nil
}

func (upd UserPausedDetailsAttributes) InternalWithRef(ref terra.Reference) UserPausedDetailsAttributes {
	return UserPausedDetailsAttributes{ref: ref}
}

func (upd UserPausedDetailsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return upd.ref.InternalTokens()
}

func (upd UserPausedDetailsAttributes) JobExpiresAt() terra.StringValue {
	return terra.ReferenceAsString(upd.ref.Append("job_expires_at"))
}

func (upd UserPausedDetailsAttributes) JobImminentExpirationHealthEventArn() terra.StringValue {
	return terra.ReferenceAsString(upd.ref.Append("job_imminent_expiration_health_event_arn"))
}

func (upd UserPausedDetailsAttributes) JobPausedAt() terra.StringValue {
	return terra.ReferenceAsString(upd.ref.Append("job_paused_at"))
}

type S3JobDefinitionAttributes struct {
	ref terra.Reference
}

func (sjd S3JobDefinitionAttributes) InternalRef() (terra.Reference, error) {
	return sjd.ref, nil
}

func (sjd S3JobDefinitionAttributes) InternalWithRef(ref terra.Reference) S3JobDefinitionAttributes {
	return S3JobDefinitionAttributes{ref: ref}
}

func (sjd S3JobDefinitionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sjd.ref.InternalTokens()
}

func (sjd S3JobDefinitionAttributes) BucketCriteria() terra.ListValue[BucketCriteriaAttributes] {
	return terra.ReferenceAsList[BucketCriteriaAttributes](sjd.ref.Append("bucket_criteria"))
}

func (sjd S3JobDefinitionAttributes) BucketDefinitions() terra.ListValue[BucketDefinitionsAttributes] {
	return terra.ReferenceAsList[BucketDefinitionsAttributes](sjd.ref.Append("bucket_definitions"))
}

func (sjd S3JobDefinitionAttributes) Scoping() terra.ListValue[ScopingAttributes] {
	return terra.ReferenceAsList[ScopingAttributes](sjd.ref.Append("scoping"))
}

type BucketCriteriaAttributes struct {
	ref terra.Reference
}

func (bc BucketCriteriaAttributes) InternalRef() (terra.Reference, error) {
	return bc.ref, nil
}

func (bc BucketCriteriaAttributes) InternalWithRef(ref terra.Reference) BucketCriteriaAttributes {
	return BucketCriteriaAttributes{ref: ref}
}

func (bc BucketCriteriaAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bc.ref.InternalTokens()
}

func (bc BucketCriteriaAttributes) Excludes() terra.ListValue[BucketCriteriaExcludesAttributes] {
	return terra.ReferenceAsList[BucketCriteriaExcludesAttributes](bc.ref.Append("excludes"))
}

func (bc BucketCriteriaAttributes) Includes() terra.ListValue[BucketCriteriaIncludesAttributes] {
	return terra.ReferenceAsList[BucketCriteriaIncludesAttributes](bc.ref.Append("includes"))
}

type BucketCriteriaExcludesAttributes struct {
	ref terra.Reference
}

func (e BucketCriteriaExcludesAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e BucketCriteriaExcludesAttributes) InternalWithRef(ref terra.Reference) BucketCriteriaExcludesAttributes {
	return BucketCriteriaExcludesAttributes{ref: ref}
}

func (e BucketCriteriaExcludesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e BucketCriteriaExcludesAttributes) And() terra.ListValue[BucketCriteriaExcludesAndAttributes] {
	return terra.ReferenceAsList[BucketCriteriaExcludesAndAttributes](e.ref.Append("and"))
}

type BucketCriteriaExcludesAndAttributes struct {
	ref terra.Reference
}

func (a BucketCriteriaExcludesAndAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a BucketCriteriaExcludesAndAttributes) InternalWithRef(ref terra.Reference) BucketCriteriaExcludesAndAttributes {
	return BucketCriteriaExcludesAndAttributes{ref: ref}
}

func (a BucketCriteriaExcludesAndAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a BucketCriteriaExcludesAndAttributes) SimpleCriterion() terra.ListValue[ExcludesAndSimpleCriterionAttributes] {
	return terra.ReferenceAsList[ExcludesAndSimpleCriterionAttributes](a.ref.Append("simple_criterion"))
}

func (a BucketCriteriaExcludesAndAttributes) TagCriterion() terra.ListValue[ExcludesAndTagCriterionAttributes] {
	return terra.ReferenceAsList[ExcludesAndTagCriterionAttributes](a.ref.Append("tag_criterion"))
}

type ExcludesAndSimpleCriterionAttributes struct {
	ref terra.Reference
}

func (sc ExcludesAndSimpleCriterionAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc ExcludesAndSimpleCriterionAttributes) InternalWithRef(ref terra.Reference) ExcludesAndSimpleCriterionAttributes {
	return ExcludesAndSimpleCriterionAttributes{ref: ref}
}

func (sc ExcludesAndSimpleCriterionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc ExcludesAndSimpleCriterionAttributes) Comparator() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("comparator"))
}

func (sc ExcludesAndSimpleCriterionAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("key"))
}

func (sc ExcludesAndSimpleCriterionAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sc.ref.Append("values"))
}

type ExcludesAndTagCriterionAttributes struct {
	ref terra.Reference
}

func (tc ExcludesAndTagCriterionAttributes) InternalRef() (terra.Reference, error) {
	return tc.ref, nil
}

func (tc ExcludesAndTagCriterionAttributes) InternalWithRef(ref terra.Reference) ExcludesAndTagCriterionAttributes {
	return ExcludesAndTagCriterionAttributes{ref: ref}
}

func (tc ExcludesAndTagCriterionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tc.ref.InternalTokens()
}

func (tc ExcludesAndTagCriterionAttributes) Comparator() terra.StringValue {
	return terra.ReferenceAsString(tc.ref.Append("comparator"))
}

func (tc ExcludesAndTagCriterionAttributes) TagValues() terra.ListValue[ExcludesAndTagCriterionTagValuesAttributes] {
	return terra.ReferenceAsList[ExcludesAndTagCriterionTagValuesAttributes](tc.ref.Append("tag_values"))
}

type ExcludesAndTagCriterionTagValuesAttributes struct {
	ref terra.Reference
}

func (tv ExcludesAndTagCriterionTagValuesAttributes) InternalRef() (terra.Reference, error) {
	return tv.ref, nil
}

func (tv ExcludesAndTagCriterionTagValuesAttributes) InternalWithRef(ref terra.Reference) ExcludesAndTagCriterionTagValuesAttributes {
	return ExcludesAndTagCriterionTagValuesAttributes{ref: ref}
}

func (tv ExcludesAndTagCriterionTagValuesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tv.ref.InternalTokens()
}

func (tv ExcludesAndTagCriterionTagValuesAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(tv.ref.Append("key"))
}

func (tv ExcludesAndTagCriterionTagValuesAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(tv.ref.Append("value"))
}

type BucketCriteriaIncludesAttributes struct {
	ref terra.Reference
}

func (i BucketCriteriaIncludesAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i BucketCriteriaIncludesAttributes) InternalWithRef(ref terra.Reference) BucketCriteriaIncludesAttributes {
	return BucketCriteriaIncludesAttributes{ref: ref}
}

func (i BucketCriteriaIncludesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i BucketCriteriaIncludesAttributes) And() terra.ListValue[BucketCriteriaIncludesAndAttributes] {
	return terra.ReferenceAsList[BucketCriteriaIncludesAndAttributes](i.ref.Append("and"))
}

type BucketCriteriaIncludesAndAttributes struct {
	ref terra.Reference
}

func (a BucketCriteriaIncludesAndAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a BucketCriteriaIncludesAndAttributes) InternalWithRef(ref terra.Reference) BucketCriteriaIncludesAndAttributes {
	return BucketCriteriaIncludesAndAttributes{ref: ref}
}

func (a BucketCriteriaIncludesAndAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a BucketCriteriaIncludesAndAttributes) SimpleCriterion() terra.ListValue[IncludesAndSimpleCriterionAttributes] {
	return terra.ReferenceAsList[IncludesAndSimpleCriterionAttributes](a.ref.Append("simple_criterion"))
}

func (a BucketCriteriaIncludesAndAttributes) TagCriterion() terra.ListValue[IncludesAndTagCriterionAttributes] {
	return terra.ReferenceAsList[IncludesAndTagCriterionAttributes](a.ref.Append("tag_criterion"))
}

type IncludesAndSimpleCriterionAttributes struct {
	ref terra.Reference
}

func (sc IncludesAndSimpleCriterionAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc IncludesAndSimpleCriterionAttributes) InternalWithRef(ref terra.Reference) IncludesAndSimpleCriterionAttributes {
	return IncludesAndSimpleCriterionAttributes{ref: ref}
}

func (sc IncludesAndSimpleCriterionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc IncludesAndSimpleCriterionAttributes) Comparator() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("comparator"))
}

func (sc IncludesAndSimpleCriterionAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("key"))
}

func (sc IncludesAndSimpleCriterionAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sc.ref.Append("values"))
}

type IncludesAndTagCriterionAttributes struct {
	ref terra.Reference
}

func (tc IncludesAndTagCriterionAttributes) InternalRef() (terra.Reference, error) {
	return tc.ref, nil
}

func (tc IncludesAndTagCriterionAttributes) InternalWithRef(ref terra.Reference) IncludesAndTagCriterionAttributes {
	return IncludesAndTagCriterionAttributes{ref: ref}
}

func (tc IncludesAndTagCriterionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tc.ref.InternalTokens()
}

func (tc IncludesAndTagCriterionAttributes) Comparator() terra.StringValue {
	return terra.ReferenceAsString(tc.ref.Append("comparator"))
}

func (tc IncludesAndTagCriterionAttributes) TagValues() terra.ListValue[IncludesAndTagCriterionTagValuesAttributes] {
	return terra.ReferenceAsList[IncludesAndTagCriterionTagValuesAttributes](tc.ref.Append("tag_values"))
}

type IncludesAndTagCriterionTagValuesAttributes struct {
	ref terra.Reference
}

func (tv IncludesAndTagCriterionTagValuesAttributes) InternalRef() (terra.Reference, error) {
	return tv.ref, nil
}

func (tv IncludesAndTagCriterionTagValuesAttributes) InternalWithRef(ref terra.Reference) IncludesAndTagCriterionTagValuesAttributes {
	return IncludesAndTagCriterionTagValuesAttributes{ref: ref}
}

func (tv IncludesAndTagCriterionTagValuesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tv.ref.InternalTokens()
}

func (tv IncludesAndTagCriterionTagValuesAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(tv.ref.Append("key"))
}

func (tv IncludesAndTagCriterionTagValuesAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(tv.ref.Append("value"))
}

type BucketDefinitionsAttributes struct {
	ref terra.Reference
}

func (bd BucketDefinitionsAttributes) InternalRef() (terra.Reference, error) {
	return bd.ref, nil
}

func (bd BucketDefinitionsAttributes) InternalWithRef(ref terra.Reference) BucketDefinitionsAttributes {
	return BucketDefinitionsAttributes{ref: ref}
}

func (bd BucketDefinitionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bd.ref.InternalTokens()
}

func (bd BucketDefinitionsAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(bd.ref.Append("account_id"))
}

func (bd BucketDefinitionsAttributes) Buckets() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](bd.ref.Append("buckets"))
}

type ScopingAttributes struct {
	ref terra.Reference
}

func (s ScopingAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s ScopingAttributes) InternalWithRef(ref terra.Reference) ScopingAttributes {
	return ScopingAttributes{ref: ref}
}

func (s ScopingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s ScopingAttributes) Excludes() terra.ListValue[ScopingExcludesAttributes] {
	return terra.ReferenceAsList[ScopingExcludesAttributes](s.ref.Append("excludes"))
}

func (s ScopingAttributes) Includes() terra.ListValue[ScopingIncludesAttributes] {
	return terra.ReferenceAsList[ScopingIncludesAttributes](s.ref.Append("includes"))
}

type ScopingExcludesAttributes struct {
	ref terra.Reference
}

func (e ScopingExcludesAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e ScopingExcludesAttributes) InternalWithRef(ref terra.Reference) ScopingExcludesAttributes {
	return ScopingExcludesAttributes{ref: ref}
}

func (e ScopingExcludesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e ScopingExcludesAttributes) And() terra.ListValue[ScopingExcludesAndAttributes] {
	return terra.ReferenceAsList[ScopingExcludesAndAttributes](e.ref.Append("and"))
}

type ScopingExcludesAndAttributes struct {
	ref terra.Reference
}

func (a ScopingExcludesAndAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a ScopingExcludesAndAttributes) InternalWithRef(ref terra.Reference) ScopingExcludesAndAttributes {
	return ScopingExcludesAndAttributes{ref: ref}
}

func (a ScopingExcludesAndAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a ScopingExcludesAndAttributes) SimpleScopeTerm() terra.ListValue[ExcludesAndSimpleScopeTermAttributes] {
	return terra.ReferenceAsList[ExcludesAndSimpleScopeTermAttributes](a.ref.Append("simple_scope_term"))
}

func (a ScopingExcludesAndAttributes) TagScopeTerm() terra.ListValue[ExcludesAndTagScopeTermAttributes] {
	return terra.ReferenceAsList[ExcludesAndTagScopeTermAttributes](a.ref.Append("tag_scope_term"))
}

type ExcludesAndSimpleScopeTermAttributes struct {
	ref terra.Reference
}

func (sst ExcludesAndSimpleScopeTermAttributes) InternalRef() (terra.Reference, error) {
	return sst.ref, nil
}

func (sst ExcludesAndSimpleScopeTermAttributes) InternalWithRef(ref terra.Reference) ExcludesAndSimpleScopeTermAttributes {
	return ExcludesAndSimpleScopeTermAttributes{ref: ref}
}

func (sst ExcludesAndSimpleScopeTermAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sst.ref.InternalTokens()
}

func (sst ExcludesAndSimpleScopeTermAttributes) Comparator() terra.StringValue {
	return terra.ReferenceAsString(sst.ref.Append("comparator"))
}

func (sst ExcludesAndSimpleScopeTermAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(sst.ref.Append("key"))
}

func (sst ExcludesAndSimpleScopeTermAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sst.ref.Append("values"))
}

type ExcludesAndTagScopeTermAttributes struct {
	ref terra.Reference
}

func (tst ExcludesAndTagScopeTermAttributes) InternalRef() (terra.Reference, error) {
	return tst.ref, nil
}

func (tst ExcludesAndTagScopeTermAttributes) InternalWithRef(ref terra.Reference) ExcludesAndTagScopeTermAttributes {
	return ExcludesAndTagScopeTermAttributes{ref: ref}
}

func (tst ExcludesAndTagScopeTermAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tst.ref.InternalTokens()
}

func (tst ExcludesAndTagScopeTermAttributes) Comparator() terra.StringValue {
	return terra.ReferenceAsString(tst.ref.Append("comparator"))
}

func (tst ExcludesAndTagScopeTermAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(tst.ref.Append("key"))
}

func (tst ExcludesAndTagScopeTermAttributes) Target() terra.StringValue {
	return terra.ReferenceAsString(tst.ref.Append("target"))
}

func (tst ExcludesAndTagScopeTermAttributes) TagValues() terra.ListValue[ExcludesAndTagScopeTermTagValuesAttributes] {
	return terra.ReferenceAsList[ExcludesAndTagScopeTermTagValuesAttributes](tst.ref.Append("tag_values"))
}

type ExcludesAndTagScopeTermTagValuesAttributes struct {
	ref terra.Reference
}

func (tv ExcludesAndTagScopeTermTagValuesAttributes) InternalRef() (terra.Reference, error) {
	return tv.ref, nil
}

func (tv ExcludesAndTagScopeTermTagValuesAttributes) InternalWithRef(ref terra.Reference) ExcludesAndTagScopeTermTagValuesAttributes {
	return ExcludesAndTagScopeTermTagValuesAttributes{ref: ref}
}

func (tv ExcludesAndTagScopeTermTagValuesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tv.ref.InternalTokens()
}

func (tv ExcludesAndTagScopeTermTagValuesAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(tv.ref.Append("key"))
}

func (tv ExcludesAndTagScopeTermTagValuesAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(tv.ref.Append("value"))
}

type ScopingIncludesAttributes struct {
	ref terra.Reference
}

func (i ScopingIncludesAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i ScopingIncludesAttributes) InternalWithRef(ref terra.Reference) ScopingIncludesAttributes {
	return ScopingIncludesAttributes{ref: ref}
}

func (i ScopingIncludesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i ScopingIncludesAttributes) And() terra.ListValue[ScopingIncludesAndAttributes] {
	return terra.ReferenceAsList[ScopingIncludesAndAttributes](i.ref.Append("and"))
}

type ScopingIncludesAndAttributes struct {
	ref terra.Reference
}

func (a ScopingIncludesAndAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a ScopingIncludesAndAttributes) InternalWithRef(ref terra.Reference) ScopingIncludesAndAttributes {
	return ScopingIncludesAndAttributes{ref: ref}
}

func (a ScopingIncludesAndAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a ScopingIncludesAndAttributes) SimpleScopeTerm() terra.ListValue[IncludesAndSimpleScopeTermAttributes] {
	return terra.ReferenceAsList[IncludesAndSimpleScopeTermAttributes](a.ref.Append("simple_scope_term"))
}

func (a ScopingIncludesAndAttributes) TagScopeTerm() terra.ListValue[IncludesAndTagScopeTermAttributes] {
	return terra.ReferenceAsList[IncludesAndTagScopeTermAttributes](a.ref.Append("tag_scope_term"))
}

type IncludesAndSimpleScopeTermAttributes struct {
	ref terra.Reference
}

func (sst IncludesAndSimpleScopeTermAttributes) InternalRef() (terra.Reference, error) {
	return sst.ref, nil
}

func (sst IncludesAndSimpleScopeTermAttributes) InternalWithRef(ref terra.Reference) IncludesAndSimpleScopeTermAttributes {
	return IncludesAndSimpleScopeTermAttributes{ref: ref}
}

func (sst IncludesAndSimpleScopeTermAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sst.ref.InternalTokens()
}

func (sst IncludesAndSimpleScopeTermAttributes) Comparator() terra.StringValue {
	return terra.ReferenceAsString(sst.ref.Append("comparator"))
}

func (sst IncludesAndSimpleScopeTermAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(sst.ref.Append("key"))
}

func (sst IncludesAndSimpleScopeTermAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sst.ref.Append("values"))
}

type IncludesAndTagScopeTermAttributes struct {
	ref terra.Reference
}

func (tst IncludesAndTagScopeTermAttributes) InternalRef() (terra.Reference, error) {
	return tst.ref, nil
}

func (tst IncludesAndTagScopeTermAttributes) InternalWithRef(ref terra.Reference) IncludesAndTagScopeTermAttributes {
	return IncludesAndTagScopeTermAttributes{ref: ref}
}

func (tst IncludesAndTagScopeTermAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tst.ref.InternalTokens()
}

func (tst IncludesAndTagScopeTermAttributes) Comparator() terra.StringValue {
	return terra.ReferenceAsString(tst.ref.Append("comparator"))
}

func (tst IncludesAndTagScopeTermAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(tst.ref.Append("key"))
}

func (tst IncludesAndTagScopeTermAttributes) Target() terra.StringValue {
	return terra.ReferenceAsString(tst.ref.Append("target"))
}

func (tst IncludesAndTagScopeTermAttributes) TagValues() terra.ListValue[IncludesAndTagScopeTermTagValuesAttributes] {
	return terra.ReferenceAsList[IncludesAndTagScopeTermTagValuesAttributes](tst.ref.Append("tag_values"))
}

type IncludesAndTagScopeTermTagValuesAttributes struct {
	ref terra.Reference
}

func (tv IncludesAndTagScopeTermTagValuesAttributes) InternalRef() (terra.Reference, error) {
	return tv.ref, nil
}

func (tv IncludesAndTagScopeTermTagValuesAttributes) InternalWithRef(ref terra.Reference) IncludesAndTagScopeTermTagValuesAttributes {
	return IncludesAndTagScopeTermTagValuesAttributes{ref: ref}
}

func (tv IncludesAndTagScopeTermTagValuesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tv.ref.InternalTokens()
}

func (tv IncludesAndTagScopeTermTagValuesAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(tv.ref.Append("key"))
}

func (tv IncludesAndTagScopeTermTagValuesAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(tv.ref.Append("value"))
}

type ScheduleFrequencyAttributes struct {
	ref terra.Reference
}

func (sf ScheduleFrequencyAttributes) InternalRef() (terra.Reference, error) {
	return sf.ref, nil
}

func (sf ScheduleFrequencyAttributes) InternalWithRef(ref terra.Reference) ScheduleFrequencyAttributes {
	return ScheduleFrequencyAttributes{ref: ref}
}

func (sf ScheduleFrequencyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sf.ref.InternalTokens()
}

func (sf ScheduleFrequencyAttributes) DailySchedule() terra.BoolValue {
	return terra.ReferenceAsBool(sf.ref.Append("daily_schedule"))
}

func (sf ScheduleFrequencyAttributes) MonthlySchedule() terra.NumberValue {
	return terra.ReferenceAsNumber(sf.ref.Append("monthly_schedule"))
}

func (sf ScheduleFrequencyAttributes) WeeklySchedule() terra.StringValue {
	return terra.ReferenceAsString(sf.ref.Append("weekly_schedule"))
}

type UserPausedDetailsState struct {
	JobExpiresAt                        string `json:"job_expires_at"`
	JobImminentExpirationHealthEventArn string `json:"job_imminent_expiration_health_event_arn"`
	JobPausedAt                         string `json:"job_paused_at"`
}

type S3JobDefinitionState struct {
	BucketCriteria    []BucketCriteriaState    `json:"bucket_criteria"`
	BucketDefinitions []BucketDefinitionsState `json:"bucket_definitions"`
	Scoping           []ScopingState           `json:"scoping"`
}

type BucketCriteriaState struct {
	Excludes []BucketCriteriaExcludesState `json:"excludes"`
	Includes []BucketCriteriaIncludesState `json:"includes"`
}

type BucketCriteriaExcludesState struct {
	And []BucketCriteriaExcludesAndState `json:"and"`
}

type BucketCriteriaExcludesAndState struct {
	SimpleCriterion []ExcludesAndSimpleCriterionState `json:"simple_criterion"`
	TagCriterion    []ExcludesAndTagCriterionState    `json:"tag_criterion"`
}

type ExcludesAndSimpleCriterionState struct {
	Comparator string   `json:"comparator"`
	Key        string   `json:"key"`
	Values     []string `json:"values"`
}

type ExcludesAndTagCriterionState struct {
	Comparator string                                  `json:"comparator"`
	TagValues  []ExcludesAndTagCriterionTagValuesState `json:"tag_values"`
}

type ExcludesAndTagCriterionTagValuesState struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type BucketCriteriaIncludesState struct {
	And []BucketCriteriaIncludesAndState `json:"and"`
}

type BucketCriteriaIncludesAndState struct {
	SimpleCriterion []IncludesAndSimpleCriterionState `json:"simple_criterion"`
	TagCriterion    []IncludesAndTagCriterionState    `json:"tag_criterion"`
}

type IncludesAndSimpleCriterionState struct {
	Comparator string   `json:"comparator"`
	Key        string   `json:"key"`
	Values     []string `json:"values"`
}

type IncludesAndTagCriterionState struct {
	Comparator string                                  `json:"comparator"`
	TagValues  []IncludesAndTagCriterionTagValuesState `json:"tag_values"`
}

type IncludesAndTagCriterionTagValuesState struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type BucketDefinitionsState struct {
	AccountId string   `json:"account_id"`
	Buckets   []string `json:"buckets"`
}

type ScopingState struct {
	Excludes []ScopingExcludesState `json:"excludes"`
	Includes []ScopingIncludesState `json:"includes"`
}

type ScopingExcludesState struct {
	And []ScopingExcludesAndState `json:"and"`
}

type ScopingExcludesAndState struct {
	SimpleScopeTerm []ExcludesAndSimpleScopeTermState `json:"simple_scope_term"`
	TagScopeTerm    []ExcludesAndTagScopeTermState    `json:"tag_scope_term"`
}

type ExcludesAndSimpleScopeTermState struct {
	Comparator string   `json:"comparator"`
	Key        string   `json:"key"`
	Values     []string `json:"values"`
}

type ExcludesAndTagScopeTermState struct {
	Comparator string                                  `json:"comparator"`
	Key        string                                  `json:"key"`
	Target     string                                  `json:"target"`
	TagValues  []ExcludesAndTagScopeTermTagValuesState `json:"tag_values"`
}

type ExcludesAndTagScopeTermTagValuesState struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ScopingIncludesState struct {
	And []ScopingIncludesAndState `json:"and"`
}

type ScopingIncludesAndState struct {
	SimpleScopeTerm []IncludesAndSimpleScopeTermState `json:"simple_scope_term"`
	TagScopeTerm    []IncludesAndTagScopeTermState    `json:"tag_scope_term"`
}

type IncludesAndSimpleScopeTermState struct {
	Comparator string   `json:"comparator"`
	Key        string   `json:"key"`
	Values     []string `json:"values"`
}

type IncludesAndTagScopeTermState struct {
	Comparator string                                  `json:"comparator"`
	Key        string                                  `json:"key"`
	Target     string                                  `json:"target"`
	TagValues  []IncludesAndTagScopeTermTagValuesState `json:"tag_values"`
}

type IncludesAndTagScopeTermTagValuesState struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ScheduleFrequencyState struct {
	DailySchedule   bool    `json:"daily_schedule"`
	MonthlySchedule float64 `json:"monthly_schedule"`
	WeeklySchedule  string  `json:"weekly_schedule"`
}
