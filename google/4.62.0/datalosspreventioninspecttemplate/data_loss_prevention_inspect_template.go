// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datalosspreventioninspecttemplate

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type InspectConfig struct {
	// ContentOptions: list of string, optional
	ContentOptions terra.ListValue[terra.StringValue] `hcl:"content_options,attr"`
	// ExcludeInfoTypes: bool, optional
	ExcludeInfoTypes terra.BoolValue `hcl:"exclude_info_types,attr"`
	// IncludeQuote: bool, optional
	IncludeQuote terra.BoolValue `hcl:"include_quote,attr"`
	// MinLikelihood: string, optional
	MinLikelihood terra.StringValue `hcl:"min_likelihood,attr"`
	// CustomInfoTypes: min=0
	CustomInfoTypes []CustomInfoTypes `hcl:"custom_info_types,block" validate:"min=0"`
	// InspectConfigInfoTypes: min=0
	InfoTypes []InspectConfigInfoTypes `hcl:"info_types,block" validate:"min=0"`
	// Limits: optional
	Limits *Limits `hcl:"limits,block"`
	// RuleSet: min=0
	RuleSet []RuleSet `hcl:"rule_set,block" validate:"min=0"`
}

type CustomInfoTypes struct {
	// ExclusionType: string, optional
	ExclusionType terra.StringValue `hcl:"exclusion_type,attr"`
	// Likelihood: string, optional
	Likelihood terra.StringValue `hcl:"likelihood,attr"`
	// CustomInfoTypesDictionary: optional
	Dictionary *CustomInfoTypesDictionary `hcl:"dictionary,block"`
	// CustomInfoTypesInfoType: required
	InfoType *CustomInfoTypesInfoType `hcl:"info_type,block" validate:"required"`
	// CustomInfoTypesRegex: optional
	Regex *CustomInfoTypesRegex `hcl:"regex,block"`
	// StoredType: optional
	StoredType *StoredType `hcl:"stored_type,block"`
}

type CustomInfoTypesDictionary struct {
	// CustomInfoTypesDictionaryCloudStoragePath: optional
	CloudStoragePath *CustomInfoTypesDictionaryCloudStoragePath `hcl:"cloud_storage_path,block"`
	// CustomInfoTypesDictionaryWordList: optional
	WordList *CustomInfoTypesDictionaryWordList `hcl:"word_list,block"`
}

type CustomInfoTypesDictionaryCloudStoragePath struct {
	// Path: string, required
	Path terra.StringValue `hcl:"path,attr" validate:"required"`
}

type CustomInfoTypesDictionaryWordList struct {
	// Words: list of string, required
	Words terra.ListValue[terra.StringValue] `hcl:"words,attr" validate:"required"`
}

type CustomInfoTypesInfoType struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type CustomInfoTypesRegex struct {
	// GroupIndexes: list of number, optional
	GroupIndexes terra.ListValue[terra.NumberValue] `hcl:"group_indexes,attr"`
	// Pattern: string, required
	Pattern terra.StringValue `hcl:"pattern,attr" validate:"required"`
}

type StoredType struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type InspectConfigInfoTypes struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
}

type Limits struct {
	// MaxFindingsPerItem: number, required
	MaxFindingsPerItem terra.NumberValue `hcl:"max_findings_per_item,attr" validate:"required"`
	// MaxFindingsPerRequest: number, required
	MaxFindingsPerRequest terra.NumberValue `hcl:"max_findings_per_request,attr" validate:"required"`
	// MaxFindingsPerInfoType: min=0
	MaxFindingsPerInfoType []MaxFindingsPerInfoType `hcl:"max_findings_per_info_type,block" validate:"min=0"`
}

type MaxFindingsPerInfoType struct {
	// MaxFindings: number, required
	MaxFindings terra.NumberValue `hcl:"max_findings,attr" validate:"required"`
	// MaxFindingsPerInfoTypeInfoType: required
	InfoType *MaxFindingsPerInfoTypeInfoType `hcl:"info_type,block" validate:"required"`
}

type MaxFindingsPerInfoTypeInfoType struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type RuleSet struct {
	// RuleSetInfoTypes: min=1
	InfoTypes []RuleSetInfoTypes `hcl:"info_types,block" validate:"min=1"`
	// Rules: min=1
	Rules []Rules `hcl:"rules,block" validate:"min=1"`
}

type RuleSetInfoTypes struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type Rules struct {
	// ExclusionRule: optional
	ExclusionRule *ExclusionRule `hcl:"exclusion_rule,block"`
	// HotwordRule: optional
	HotwordRule *HotwordRule `hcl:"hotword_rule,block"`
}

type ExclusionRule struct {
	// MatchingType: string, required
	MatchingType terra.StringValue `hcl:"matching_type,attr" validate:"required"`
	// ExclusionRuleDictionary: optional
	Dictionary *ExclusionRuleDictionary `hcl:"dictionary,block"`
	// ExcludeInfoTypes: optional
	ExcludeInfoTypes *ExcludeInfoTypes `hcl:"exclude_info_types,block"`
	// ExclusionRuleRegex: optional
	Regex *ExclusionRuleRegex `hcl:"regex,block"`
}

type ExclusionRuleDictionary struct {
	// ExclusionRuleDictionaryCloudStoragePath: optional
	CloudStoragePath *ExclusionRuleDictionaryCloudStoragePath `hcl:"cloud_storage_path,block"`
	// ExclusionRuleDictionaryWordList: optional
	WordList *ExclusionRuleDictionaryWordList `hcl:"word_list,block"`
}

type ExclusionRuleDictionaryCloudStoragePath struct {
	// Path: string, required
	Path terra.StringValue `hcl:"path,attr" validate:"required"`
}

type ExclusionRuleDictionaryWordList struct {
	// Words: list of string, required
	Words terra.ListValue[terra.StringValue] `hcl:"words,attr" validate:"required"`
}

type ExcludeInfoTypes struct {
	// ExcludeInfoTypesInfoTypes: min=1
	InfoTypes []ExcludeInfoTypesInfoTypes `hcl:"info_types,block" validate:"min=1"`
}

type ExcludeInfoTypesInfoTypes struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type ExclusionRuleRegex struct {
	// GroupIndexes: list of number, optional
	GroupIndexes terra.ListValue[terra.NumberValue] `hcl:"group_indexes,attr"`
	// Pattern: string, required
	Pattern terra.StringValue `hcl:"pattern,attr" validate:"required"`
}

type HotwordRule struct {
	// HotwordRegex: required
	HotwordRegex *HotwordRegex `hcl:"hotword_regex,block" validate:"required"`
	// LikelihoodAdjustment: required
	LikelihoodAdjustment *LikelihoodAdjustment `hcl:"likelihood_adjustment,block" validate:"required"`
	// Proximity: required
	Proximity *Proximity `hcl:"proximity,block" validate:"required"`
}

type HotwordRegex struct {
	// GroupIndexes: list of number, optional
	GroupIndexes terra.ListValue[terra.NumberValue] `hcl:"group_indexes,attr"`
	// Pattern: string, required
	Pattern terra.StringValue `hcl:"pattern,attr" validate:"required"`
}

type LikelihoodAdjustment struct {
	// FixedLikelihood: string, optional
	FixedLikelihood terra.StringValue `hcl:"fixed_likelihood,attr"`
	// RelativeLikelihood: number, optional
	RelativeLikelihood terra.NumberValue `hcl:"relative_likelihood,attr"`
}

type Proximity struct {
	// WindowAfter: number, optional
	WindowAfter terra.NumberValue `hcl:"window_after,attr"`
	// WindowBefore: number, optional
	WindowBefore terra.NumberValue `hcl:"window_before,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type InspectConfigAttributes struct {
	ref terra.Reference
}

func (ic InspectConfigAttributes) InternalRef() (terra.Reference, error) {
	return ic.ref, nil
}

func (ic InspectConfigAttributes) InternalWithRef(ref terra.Reference) InspectConfigAttributes {
	return InspectConfigAttributes{ref: ref}
}

func (ic InspectConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ic.ref.InternalTokens()
}

func (ic InspectConfigAttributes) ContentOptions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ic.ref.Append("content_options"))
}

func (ic InspectConfigAttributes) ExcludeInfoTypes() terra.BoolValue {
	return terra.ReferenceAsBool(ic.ref.Append("exclude_info_types"))
}

func (ic InspectConfigAttributes) IncludeQuote() terra.BoolValue {
	return terra.ReferenceAsBool(ic.ref.Append("include_quote"))
}

func (ic InspectConfigAttributes) MinLikelihood() terra.StringValue {
	return terra.ReferenceAsString(ic.ref.Append("min_likelihood"))
}

func (ic InspectConfigAttributes) CustomInfoTypes() terra.ListValue[CustomInfoTypesAttributes] {
	return terra.ReferenceAsList[CustomInfoTypesAttributes](ic.ref.Append("custom_info_types"))
}

func (ic InspectConfigAttributes) InfoTypes() terra.ListValue[InspectConfigInfoTypesAttributes] {
	return terra.ReferenceAsList[InspectConfigInfoTypesAttributes](ic.ref.Append("info_types"))
}

func (ic InspectConfigAttributes) Limits() terra.ListValue[LimitsAttributes] {
	return terra.ReferenceAsList[LimitsAttributes](ic.ref.Append("limits"))
}

func (ic InspectConfigAttributes) RuleSet() terra.ListValue[RuleSetAttributes] {
	return terra.ReferenceAsList[RuleSetAttributes](ic.ref.Append("rule_set"))
}

type CustomInfoTypesAttributes struct {
	ref terra.Reference
}

func (cit CustomInfoTypesAttributes) InternalRef() (terra.Reference, error) {
	return cit.ref, nil
}

func (cit CustomInfoTypesAttributes) InternalWithRef(ref terra.Reference) CustomInfoTypesAttributes {
	return CustomInfoTypesAttributes{ref: ref}
}

func (cit CustomInfoTypesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cit.ref.InternalTokens()
}

func (cit CustomInfoTypesAttributes) ExclusionType() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("exclusion_type"))
}

func (cit CustomInfoTypesAttributes) Likelihood() terra.StringValue {
	return terra.ReferenceAsString(cit.ref.Append("likelihood"))
}

func (cit CustomInfoTypesAttributes) Dictionary() terra.ListValue[CustomInfoTypesDictionaryAttributes] {
	return terra.ReferenceAsList[CustomInfoTypesDictionaryAttributes](cit.ref.Append("dictionary"))
}

func (cit CustomInfoTypesAttributes) InfoType() terra.ListValue[CustomInfoTypesInfoTypeAttributes] {
	return terra.ReferenceAsList[CustomInfoTypesInfoTypeAttributes](cit.ref.Append("info_type"))
}

func (cit CustomInfoTypesAttributes) Regex() terra.ListValue[CustomInfoTypesRegexAttributes] {
	return terra.ReferenceAsList[CustomInfoTypesRegexAttributes](cit.ref.Append("regex"))
}

func (cit CustomInfoTypesAttributes) StoredType() terra.ListValue[StoredTypeAttributes] {
	return terra.ReferenceAsList[StoredTypeAttributes](cit.ref.Append("stored_type"))
}

type CustomInfoTypesDictionaryAttributes struct {
	ref terra.Reference
}

func (d CustomInfoTypesDictionaryAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d CustomInfoTypesDictionaryAttributes) InternalWithRef(ref terra.Reference) CustomInfoTypesDictionaryAttributes {
	return CustomInfoTypesDictionaryAttributes{ref: ref}
}

func (d CustomInfoTypesDictionaryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d CustomInfoTypesDictionaryAttributes) CloudStoragePath() terra.ListValue[CustomInfoTypesDictionaryCloudStoragePathAttributes] {
	return terra.ReferenceAsList[CustomInfoTypesDictionaryCloudStoragePathAttributes](d.ref.Append("cloud_storage_path"))
}

func (d CustomInfoTypesDictionaryAttributes) WordList() terra.ListValue[CustomInfoTypesDictionaryWordListAttributes] {
	return terra.ReferenceAsList[CustomInfoTypesDictionaryWordListAttributes](d.ref.Append("word_list"))
}

type CustomInfoTypesDictionaryCloudStoragePathAttributes struct {
	ref terra.Reference
}

func (csp CustomInfoTypesDictionaryCloudStoragePathAttributes) InternalRef() (terra.Reference, error) {
	return csp.ref, nil
}

func (csp CustomInfoTypesDictionaryCloudStoragePathAttributes) InternalWithRef(ref terra.Reference) CustomInfoTypesDictionaryCloudStoragePathAttributes {
	return CustomInfoTypesDictionaryCloudStoragePathAttributes{ref: ref}
}

func (csp CustomInfoTypesDictionaryCloudStoragePathAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return csp.ref.InternalTokens()
}

func (csp CustomInfoTypesDictionaryCloudStoragePathAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(csp.ref.Append("path"))
}

type CustomInfoTypesDictionaryWordListAttributes struct {
	ref terra.Reference
}

func (wl CustomInfoTypesDictionaryWordListAttributes) InternalRef() (terra.Reference, error) {
	return wl.ref, nil
}

func (wl CustomInfoTypesDictionaryWordListAttributes) InternalWithRef(ref terra.Reference) CustomInfoTypesDictionaryWordListAttributes {
	return CustomInfoTypesDictionaryWordListAttributes{ref: ref}
}

func (wl CustomInfoTypesDictionaryWordListAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wl.ref.InternalTokens()
}

func (wl CustomInfoTypesDictionaryWordListAttributes) Words() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](wl.ref.Append("words"))
}

type CustomInfoTypesInfoTypeAttributes struct {
	ref terra.Reference
}

func (it CustomInfoTypesInfoTypeAttributes) InternalRef() (terra.Reference, error) {
	return it.ref, nil
}

func (it CustomInfoTypesInfoTypeAttributes) InternalWithRef(ref terra.Reference) CustomInfoTypesInfoTypeAttributes {
	return CustomInfoTypesInfoTypeAttributes{ref: ref}
}

func (it CustomInfoTypesInfoTypeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return it.ref.InternalTokens()
}

func (it CustomInfoTypesInfoTypeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(it.ref.Append("name"))
}

type CustomInfoTypesRegexAttributes struct {
	ref terra.Reference
}

func (r CustomInfoTypesRegexAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r CustomInfoTypesRegexAttributes) InternalWithRef(ref terra.Reference) CustomInfoTypesRegexAttributes {
	return CustomInfoTypesRegexAttributes{ref: ref}
}

func (r CustomInfoTypesRegexAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r CustomInfoTypesRegexAttributes) GroupIndexes() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](r.ref.Append("group_indexes"))
}

func (r CustomInfoTypesRegexAttributes) Pattern() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("pattern"))
}

type StoredTypeAttributes struct {
	ref terra.Reference
}

func (st StoredTypeAttributes) InternalRef() (terra.Reference, error) {
	return st.ref, nil
}

func (st StoredTypeAttributes) InternalWithRef(ref terra.Reference) StoredTypeAttributes {
	return StoredTypeAttributes{ref: ref}
}

func (st StoredTypeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return st.ref.InternalTokens()
}

func (st StoredTypeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("name"))
}

type InspectConfigInfoTypesAttributes struct {
	ref terra.Reference
}

func (it InspectConfigInfoTypesAttributes) InternalRef() (terra.Reference, error) {
	return it.ref, nil
}

func (it InspectConfigInfoTypesAttributes) InternalWithRef(ref terra.Reference) InspectConfigInfoTypesAttributes {
	return InspectConfigInfoTypesAttributes{ref: ref}
}

func (it InspectConfigInfoTypesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return it.ref.InternalTokens()
}

func (it InspectConfigInfoTypesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(it.ref.Append("name"))
}

func (it InspectConfigInfoTypesAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(it.ref.Append("version"))
}

type LimitsAttributes struct {
	ref terra.Reference
}

func (l LimitsAttributes) InternalRef() (terra.Reference, error) {
	return l.ref, nil
}

func (l LimitsAttributes) InternalWithRef(ref terra.Reference) LimitsAttributes {
	return LimitsAttributes{ref: ref}
}

func (l LimitsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return l.ref.InternalTokens()
}

func (l LimitsAttributes) MaxFindingsPerItem() terra.NumberValue {
	return terra.ReferenceAsNumber(l.ref.Append("max_findings_per_item"))
}

func (l LimitsAttributes) MaxFindingsPerRequest() terra.NumberValue {
	return terra.ReferenceAsNumber(l.ref.Append("max_findings_per_request"))
}

func (l LimitsAttributes) MaxFindingsPerInfoType() terra.ListValue[MaxFindingsPerInfoTypeAttributes] {
	return terra.ReferenceAsList[MaxFindingsPerInfoTypeAttributes](l.ref.Append("max_findings_per_info_type"))
}

type MaxFindingsPerInfoTypeAttributes struct {
	ref terra.Reference
}

func (mfpit MaxFindingsPerInfoTypeAttributes) InternalRef() (terra.Reference, error) {
	return mfpit.ref, nil
}

func (mfpit MaxFindingsPerInfoTypeAttributes) InternalWithRef(ref terra.Reference) MaxFindingsPerInfoTypeAttributes {
	return MaxFindingsPerInfoTypeAttributes{ref: ref}
}

func (mfpit MaxFindingsPerInfoTypeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mfpit.ref.InternalTokens()
}

func (mfpit MaxFindingsPerInfoTypeAttributes) MaxFindings() terra.NumberValue {
	return terra.ReferenceAsNumber(mfpit.ref.Append("max_findings"))
}

func (mfpit MaxFindingsPerInfoTypeAttributes) InfoType() terra.ListValue[MaxFindingsPerInfoTypeInfoTypeAttributes] {
	return terra.ReferenceAsList[MaxFindingsPerInfoTypeInfoTypeAttributes](mfpit.ref.Append("info_type"))
}

type MaxFindingsPerInfoTypeInfoTypeAttributes struct {
	ref terra.Reference
}

func (it MaxFindingsPerInfoTypeInfoTypeAttributes) InternalRef() (terra.Reference, error) {
	return it.ref, nil
}

func (it MaxFindingsPerInfoTypeInfoTypeAttributes) InternalWithRef(ref terra.Reference) MaxFindingsPerInfoTypeInfoTypeAttributes {
	return MaxFindingsPerInfoTypeInfoTypeAttributes{ref: ref}
}

func (it MaxFindingsPerInfoTypeInfoTypeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return it.ref.InternalTokens()
}

func (it MaxFindingsPerInfoTypeInfoTypeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(it.ref.Append("name"))
}

type RuleSetAttributes struct {
	ref terra.Reference
}

func (rs RuleSetAttributes) InternalRef() (terra.Reference, error) {
	return rs.ref, nil
}

func (rs RuleSetAttributes) InternalWithRef(ref terra.Reference) RuleSetAttributes {
	return RuleSetAttributes{ref: ref}
}

func (rs RuleSetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rs.ref.InternalTokens()
}

func (rs RuleSetAttributes) InfoTypes() terra.ListValue[RuleSetInfoTypesAttributes] {
	return terra.ReferenceAsList[RuleSetInfoTypesAttributes](rs.ref.Append("info_types"))
}

func (rs RuleSetAttributes) Rules() terra.ListValue[RulesAttributes] {
	return terra.ReferenceAsList[RulesAttributes](rs.ref.Append("rules"))
}

type RuleSetInfoTypesAttributes struct {
	ref terra.Reference
}

func (it RuleSetInfoTypesAttributes) InternalRef() (terra.Reference, error) {
	return it.ref, nil
}

func (it RuleSetInfoTypesAttributes) InternalWithRef(ref terra.Reference) RuleSetInfoTypesAttributes {
	return RuleSetInfoTypesAttributes{ref: ref}
}

func (it RuleSetInfoTypesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return it.ref.InternalTokens()
}

func (it RuleSetInfoTypesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(it.ref.Append("name"))
}

type RulesAttributes struct {
	ref terra.Reference
}

func (r RulesAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RulesAttributes) InternalWithRef(ref terra.Reference) RulesAttributes {
	return RulesAttributes{ref: ref}
}

func (r RulesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RulesAttributes) ExclusionRule() terra.ListValue[ExclusionRuleAttributes] {
	return terra.ReferenceAsList[ExclusionRuleAttributes](r.ref.Append("exclusion_rule"))
}

func (r RulesAttributes) HotwordRule() terra.ListValue[HotwordRuleAttributes] {
	return terra.ReferenceAsList[HotwordRuleAttributes](r.ref.Append("hotword_rule"))
}

type ExclusionRuleAttributes struct {
	ref terra.Reference
}

func (er ExclusionRuleAttributes) InternalRef() (terra.Reference, error) {
	return er.ref, nil
}

func (er ExclusionRuleAttributes) InternalWithRef(ref terra.Reference) ExclusionRuleAttributes {
	return ExclusionRuleAttributes{ref: ref}
}

func (er ExclusionRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return er.ref.InternalTokens()
}

func (er ExclusionRuleAttributes) MatchingType() terra.StringValue {
	return terra.ReferenceAsString(er.ref.Append("matching_type"))
}

func (er ExclusionRuleAttributes) Dictionary() terra.ListValue[ExclusionRuleDictionaryAttributes] {
	return terra.ReferenceAsList[ExclusionRuleDictionaryAttributes](er.ref.Append("dictionary"))
}

func (er ExclusionRuleAttributes) ExcludeInfoTypes() terra.ListValue[ExcludeInfoTypesAttributes] {
	return terra.ReferenceAsList[ExcludeInfoTypesAttributes](er.ref.Append("exclude_info_types"))
}

func (er ExclusionRuleAttributes) Regex() terra.ListValue[ExclusionRuleRegexAttributes] {
	return terra.ReferenceAsList[ExclusionRuleRegexAttributes](er.ref.Append("regex"))
}

type ExclusionRuleDictionaryAttributes struct {
	ref terra.Reference
}

func (d ExclusionRuleDictionaryAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d ExclusionRuleDictionaryAttributes) InternalWithRef(ref terra.Reference) ExclusionRuleDictionaryAttributes {
	return ExclusionRuleDictionaryAttributes{ref: ref}
}

func (d ExclusionRuleDictionaryAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d ExclusionRuleDictionaryAttributes) CloudStoragePath() terra.ListValue[ExclusionRuleDictionaryCloudStoragePathAttributes] {
	return terra.ReferenceAsList[ExclusionRuleDictionaryCloudStoragePathAttributes](d.ref.Append("cloud_storage_path"))
}

func (d ExclusionRuleDictionaryAttributes) WordList() terra.ListValue[ExclusionRuleDictionaryWordListAttributes] {
	return terra.ReferenceAsList[ExclusionRuleDictionaryWordListAttributes](d.ref.Append("word_list"))
}

type ExclusionRuleDictionaryCloudStoragePathAttributes struct {
	ref terra.Reference
}

func (csp ExclusionRuleDictionaryCloudStoragePathAttributes) InternalRef() (terra.Reference, error) {
	return csp.ref, nil
}

func (csp ExclusionRuleDictionaryCloudStoragePathAttributes) InternalWithRef(ref terra.Reference) ExclusionRuleDictionaryCloudStoragePathAttributes {
	return ExclusionRuleDictionaryCloudStoragePathAttributes{ref: ref}
}

func (csp ExclusionRuleDictionaryCloudStoragePathAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return csp.ref.InternalTokens()
}

func (csp ExclusionRuleDictionaryCloudStoragePathAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(csp.ref.Append("path"))
}

type ExclusionRuleDictionaryWordListAttributes struct {
	ref terra.Reference
}

func (wl ExclusionRuleDictionaryWordListAttributes) InternalRef() (terra.Reference, error) {
	return wl.ref, nil
}

func (wl ExclusionRuleDictionaryWordListAttributes) InternalWithRef(ref terra.Reference) ExclusionRuleDictionaryWordListAttributes {
	return ExclusionRuleDictionaryWordListAttributes{ref: ref}
}

func (wl ExclusionRuleDictionaryWordListAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wl.ref.InternalTokens()
}

func (wl ExclusionRuleDictionaryWordListAttributes) Words() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](wl.ref.Append("words"))
}

type ExcludeInfoTypesAttributes struct {
	ref terra.Reference
}

func (eit ExcludeInfoTypesAttributes) InternalRef() (terra.Reference, error) {
	return eit.ref, nil
}

func (eit ExcludeInfoTypesAttributes) InternalWithRef(ref terra.Reference) ExcludeInfoTypesAttributes {
	return ExcludeInfoTypesAttributes{ref: ref}
}

func (eit ExcludeInfoTypesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return eit.ref.InternalTokens()
}

func (eit ExcludeInfoTypesAttributes) InfoTypes() terra.ListValue[ExcludeInfoTypesInfoTypesAttributes] {
	return terra.ReferenceAsList[ExcludeInfoTypesInfoTypesAttributes](eit.ref.Append("info_types"))
}

type ExcludeInfoTypesInfoTypesAttributes struct {
	ref terra.Reference
}

func (it ExcludeInfoTypesInfoTypesAttributes) InternalRef() (terra.Reference, error) {
	return it.ref, nil
}

func (it ExcludeInfoTypesInfoTypesAttributes) InternalWithRef(ref terra.Reference) ExcludeInfoTypesInfoTypesAttributes {
	return ExcludeInfoTypesInfoTypesAttributes{ref: ref}
}

func (it ExcludeInfoTypesInfoTypesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return it.ref.InternalTokens()
}

func (it ExcludeInfoTypesInfoTypesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(it.ref.Append("name"))
}

type ExclusionRuleRegexAttributes struct {
	ref terra.Reference
}

func (r ExclusionRuleRegexAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r ExclusionRuleRegexAttributes) InternalWithRef(ref terra.Reference) ExclusionRuleRegexAttributes {
	return ExclusionRuleRegexAttributes{ref: ref}
}

func (r ExclusionRuleRegexAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r ExclusionRuleRegexAttributes) GroupIndexes() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](r.ref.Append("group_indexes"))
}

func (r ExclusionRuleRegexAttributes) Pattern() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("pattern"))
}

type HotwordRuleAttributes struct {
	ref terra.Reference
}

func (hr HotwordRuleAttributes) InternalRef() (terra.Reference, error) {
	return hr.ref, nil
}

func (hr HotwordRuleAttributes) InternalWithRef(ref terra.Reference) HotwordRuleAttributes {
	return HotwordRuleAttributes{ref: ref}
}

func (hr HotwordRuleAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hr.ref.InternalTokens()
}

func (hr HotwordRuleAttributes) HotwordRegex() terra.ListValue[HotwordRegexAttributes] {
	return terra.ReferenceAsList[HotwordRegexAttributes](hr.ref.Append("hotword_regex"))
}

func (hr HotwordRuleAttributes) LikelihoodAdjustment() terra.ListValue[LikelihoodAdjustmentAttributes] {
	return terra.ReferenceAsList[LikelihoodAdjustmentAttributes](hr.ref.Append("likelihood_adjustment"))
}

func (hr HotwordRuleAttributes) Proximity() terra.ListValue[ProximityAttributes] {
	return terra.ReferenceAsList[ProximityAttributes](hr.ref.Append("proximity"))
}

type HotwordRegexAttributes struct {
	ref terra.Reference
}

func (hr HotwordRegexAttributes) InternalRef() (terra.Reference, error) {
	return hr.ref, nil
}

func (hr HotwordRegexAttributes) InternalWithRef(ref terra.Reference) HotwordRegexAttributes {
	return HotwordRegexAttributes{ref: ref}
}

func (hr HotwordRegexAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hr.ref.InternalTokens()
}

func (hr HotwordRegexAttributes) GroupIndexes() terra.ListValue[terra.NumberValue] {
	return terra.ReferenceAsList[terra.NumberValue](hr.ref.Append("group_indexes"))
}

func (hr HotwordRegexAttributes) Pattern() terra.StringValue {
	return terra.ReferenceAsString(hr.ref.Append("pattern"))
}

type LikelihoodAdjustmentAttributes struct {
	ref terra.Reference
}

func (la LikelihoodAdjustmentAttributes) InternalRef() (terra.Reference, error) {
	return la.ref, nil
}

func (la LikelihoodAdjustmentAttributes) InternalWithRef(ref terra.Reference) LikelihoodAdjustmentAttributes {
	return LikelihoodAdjustmentAttributes{ref: ref}
}

func (la LikelihoodAdjustmentAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return la.ref.InternalTokens()
}

func (la LikelihoodAdjustmentAttributes) FixedLikelihood() terra.StringValue {
	return terra.ReferenceAsString(la.ref.Append("fixed_likelihood"))
}

func (la LikelihoodAdjustmentAttributes) RelativeLikelihood() terra.NumberValue {
	return terra.ReferenceAsNumber(la.ref.Append("relative_likelihood"))
}

type ProximityAttributes struct {
	ref terra.Reference
}

func (p ProximityAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p ProximityAttributes) InternalWithRef(ref terra.Reference) ProximityAttributes {
	return ProximityAttributes{ref: ref}
}

func (p ProximityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p ProximityAttributes) WindowAfter() terra.NumberValue {
	return terra.ReferenceAsNumber(p.ref.Append("window_after"))
}

func (p ProximityAttributes) WindowBefore() terra.NumberValue {
	return terra.ReferenceAsNumber(p.ref.Append("window_before"))
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

type InspectConfigState struct {
	ContentOptions   []string                      `json:"content_options"`
	ExcludeInfoTypes bool                          `json:"exclude_info_types"`
	IncludeQuote     bool                          `json:"include_quote"`
	MinLikelihood    string                        `json:"min_likelihood"`
	CustomInfoTypes  []CustomInfoTypesState        `json:"custom_info_types"`
	InfoTypes        []InspectConfigInfoTypesState `json:"info_types"`
	Limits           []LimitsState                 `json:"limits"`
	RuleSet          []RuleSetState                `json:"rule_set"`
}

type CustomInfoTypesState struct {
	ExclusionType string                           `json:"exclusion_type"`
	Likelihood    string                           `json:"likelihood"`
	Dictionary    []CustomInfoTypesDictionaryState `json:"dictionary"`
	InfoType      []CustomInfoTypesInfoTypeState   `json:"info_type"`
	Regex         []CustomInfoTypesRegexState      `json:"regex"`
	StoredType    []StoredTypeState                `json:"stored_type"`
}

type CustomInfoTypesDictionaryState struct {
	CloudStoragePath []CustomInfoTypesDictionaryCloudStoragePathState `json:"cloud_storage_path"`
	WordList         []CustomInfoTypesDictionaryWordListState         `json:"word_list"`
}

type CustomInfoTypesDictionaryCloudStoragePathState struct {
	Path string `json:"path"`
}

type CustomInfoTypesDictionaryWordListState struct {
	Words []string `json:"words"`
}

type CustomInfoTypesInfoTypeState struct {
	Name string `json:"name"`
}

type CustomInfoTypesRegexState struct {
	GroupIndexes []float64 `json:"group_indexes"`
	Pattern      string    `json:"pattern"`
}

type StoredTypeState struct {
	Name string `json:"name"`
}

type InspectConfigInfoTypesState struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type LimitsState struct {
	MaxFindingsPerItem     float64                       `json:"max_findings_per_item"`
	MaxFindingsPerRequest  float64                       `json:"max_findings_per_request"`
	MaxFindingsPerInfoType []MaxFindingsPerInfoTypeState `json:"max_findings_per_info_type"`
}

type MaxFindingsPerInfoTypeState struct {
	MaxFindings float64                               `json:"max_findings"`
	InfoType    []MaxFindingsPerInfoTypeInfoTypeState `json:"info_type"`
}

type MaxFindingsPerInfoTypeInfoTypeState struct {
	Name string `json:"name"`
}

type RuleSetState struct {
	InfoTypes []RuleSetInfoTypesState `json:"info_types"`
	Rules     []RulesState            `json:"rules"`
}

type RuleSetInfoTypesState struct {
	Name string `json:"name"`
}

type RulesState struct {
	ExclusionRule []ExclusionRuleState `json:"exclusion_rule"`
	HotwordRule   []HotwordRuleState   `json:"hotword_rule"`
}

type ExclusionRuleState struct {
	MatchingType     string                         `json:"matching_type"`
	Dictionary       []ExclusionRuleDictionaryState `json:"dictionary"`
	ExcludeInfoTypes []ExcludeInfoTypesState        `json:"exclude_info_types"`
	Regex            []ExclusionRuleRegexState      `json:"regex"`
}

type ExclusionRuleDictionaryState struct {
	CloudStoragePath []ExclusionRuleDictionaryCloudStoragePathState `json:"cloud_storage_path"`
	WordList         []ExclusionRuleDictionaryWordListState         `json:"word_list"`
}

type ExclusionRuleDictionaryCloudStoragePathState struct {
	Path string `json:"path"`
}

type ExclusionRuleDictionaryWordListState struct {
	Words []string `json:"words"`
}

type ExcludeInfoTypesState struct {
	InfoTypes []ExcludeInfoTypesInfoTypesState `json:"info_types"`
}

type ExcludeInfoTypesInfoTypesState struct {
	Name string `json:"name"`
}

type ExclusionRuleRegexState struct {
	GroupIndexes []float64 `json:"group_indexes"`
	Pattern      string    `json:"pattern"`
}

type HotwordRuleState struct {
	HotwordRegex         []HotwordRegexState         `json:"hotword_regex"`
	LikelihoodAdjustment []LikelihoodAdjustmentState `json:"likelihood_adjustment"`
	Proximity            []ProximityState            `json:"proximity"`
}

type HotwordRegexState struct {
	GroupIndexes []float64 `json:"group_indexes"`
	Pattern      string    `json:"pattern"`
}

type LikelihoodAdjustmentState struct {
	FixedLikelihood    string  `json:"fixed_likelihood"`
	RelativeLikelihood float64 `json:"relative_likelihood"`
}

type ProximityState struct {
	WindowAfter  float64 `json:"window_after"`
	WindowBefore float64 `json:"window_before"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
