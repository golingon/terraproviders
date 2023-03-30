// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datakendraindex

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type CapacityUnits struct{}

type DocumentMetadataConfigurationUpdates struct {
	// Relevance: min=0
	Relevance []Relevance `hcl:"relevance,block" validate:"min=0"`
	// Search: min=0
	Search []Search `hcl:"search,block" validate:"min=0"`
}

type Relevance struct{}

type Search struct{}

type IndexStatistics struct {
	// FaqStatistics: min=0
	FaqStatistics []FaqStatistics `hcl:"faq_statistics,block" validate:"min=0"`
	// TextDocumentStatistics: min=0
	TextDocumentStatistics []TextDocumentStatistics `hcl:"text_document_statistics,block" validate:"min=0"`
}

type FaqStatistics struct{}

type TextDocumentStatistics struct{}

type ServerSideEncryptionConfiguration struct{}

type UserGroupResolutionConfiguration struct{}

type UserTokenConfigurations struct {
	// JsonTokenTypeConfiguration: min=0
	JsonTokenTypeConfiguration []JsonTokenTypeConfiguration `hcl:"json_token_type_configuration,block" validate:"min=0"`
	// JwtTokenTypeConfiguration: min=0
	JwtTokenTypeConfiguration []JwtTokenTypeConfiguration `hcl:"jwt_token_type_configuration,block" validate:"min=0"`
}

type JsonTokenTypeConfiguration struct{}

type JwtTokenTypeConfiguration struct{}

type CapacityUnitsAttributes struct {
	ref terra.Reference
}

func (cu CapacityUnitsAttributes) InternalRef() terra.Reference {
	return cu.ref
}

func (cu CapacityUnitsAttributes) InternalWithRef(ref terra.Reference) CapacityUnitsAttributes {
	return CapacityUnitsAttributes{ref: ref}
}

func (cu CapacityUnitsAttributes) InternalTokens() hclwrite.Tokens {
	return cu.ref.InternalTokens()
}

func (cu CapacityUnitsAttributes) QueryCapacityUnits() terra.NumberValue {
	return terra.ReferenceNumber(cu.ref.Append("query_capacity_units"))
}

func (cu CapacityUnitsAttributes) StorageCapacityUnits() terra.NumberValue {
	return terra.ReferenceNumber(cu.ref.Append("storage_capacity_units"))
}

type DocumentMetadataConfigurationUpdatesAttributes struct {
	ref terra.Reference
}

func (dmcu DocumentMetadataConfigurationUpdatesAttributes) InternalRef() terra.Reference {
	return dmcu.ref
}

func (dmcu DocumentMetadataConfigurationUpdatesAttributes) InternalWithRef(ref terra.Reference) DocumentMetadataConfigurationUpdatesAttributes {
	return DocumentMetadataConfigurationUpdatesAttributes{ref: ref}
}

func (dmcu DocumentMetadataConfigurationUpdatesAttributes) InternalTokens() hclwrite.Tokens {
	return dmcu.ref.InternalTokens()
}

func (dmcu DocumentMetadataConfigurationUpdatesAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dmcu.ref.Append("name"))
}

func (dmcu DocumentMetadataConfigurationUpdatesAttributes) Type() terra.StringValue {
	return terra.ReferenceString(dmcu.ref.Append("type"))
}

func (dmcu DocumentMetadataConfigurationUpdatesAttributes) Relevance() terra.ListValue[RelevanceAttributes] {
	return terra.ReferenceList[RelevanceAttributes](dmcu.ref.Append("relevance"))
}

func (dmcu DocumentMetadataConfigurationUpdatesAttributes) Search() terra.ListValue[SearchAttributes] {
	return terra.ReferenceList[SearchAttributes](dmcu.ref.Append("search"))
}

type RelevanceAttributes struct {
	ref terra.Reference
}

func (r RelevanceAttributes) InternalRef() terra.Reference {
	return r.ref
}

func (r RelevanceAttributes) InternalWithRef(ref terra.Reference) RelevanceAttributes {
	return RelevanceAttributes{ref: ref}
}

func (r RelevanceAttributes) InternalTokens() hclwrite.Tokens {
	return r.ref.InternalTokens()
}

func (r RelevanceAttributes) Duration() terra.StringValue {
	return terra.ReferenceString(r.ref.Append("duration"))
}

func (r RelevanceAttributes) Freshness() terra.BoolValue {
	return terra.ReferenceBool(r.ref.Append("freshness"))
}

func (r RelevanceAttributes) Importance() terra.NumberValue {
	return terra.ReferenceNumber(r.ref.Append("importance"))
}

func (r RelevanceAttributes) RankOrder() terra.StringValue {
	return terra.ReferenceString(r.ref.Append("rank_order"))
}

func (r RelevanceAttributes) ValuesImportanceMap() terra.MapValue[terra.NumberValue] {
	return terra.ReferenceMap[terra.NumberValue](r.ref.Append("values_importance_map"))
}

type SearchAttributes struct {
	ref terra.Reference
}

func (s SearchAttributes) InternalRef() terra.Reference {
	return s.ref
}

func (s SearchAttributes) InternalWithRef(ref terra.Reference) SearchAttributes {
	return SearchAttributes{ref: ref}
}

func (s SearchAttributes) InternalTokens() hclwrite.Tokens {
	return s.ref.InternalTokens()
}

func (s SearchAttributes) Displayable() terra.BoolValue {
	return terra.ReferenceBool(s.ref.Append("displayable"))
}

func (s SearchAttributes) Facetable() terra.BoolValue {
	return terra.ReferenceBool(s.ref.Append("facetable"))
}

func (s SearchAttributes) Searchable() terra.BoolValue {
	return terra.ReferenceBool(s.ref.Append("searchable"))
}

func (s SearchAttributes) Sortable() terra.BoolValue {
	return terra.ReferenceBool(s.ref.Append("sortable"))
}

type IndexStatisticsAttributes struct {
	ref terra.Reference
}

func (is IndexStatisticsAttributes) InternalRef() terra.Reference {
	return is.ref
}

func (is IndexStatisticsAttributes) InternalWithRef(ref terra.Reference) IndexStatisticsAttributes {
	return IndexStatisticsAttributes{ref: ref}
}

func (is IndexStatisticsAttributes) InternalTokens() hclwrite.Tokens {
	return is.ref.InternalTokens()
}

func (is IndexStatisticsAttributes) FaqStatistics() terra.ListValue[FaqStatisticsAttributes] {
	return terra.ReferenceList[FaqStatisticsAttributes](is.ref.Append("faq_statistics"))
}

func (is IndexStatisticsAttributes) TextDocumentStatistics() terra.ListValue[TextDocumentStatisticsAttributes] {
	return terra.ReferenceList[TextDocumentStatisticsAttributes](is.ref.Append("text_document_statistics"))
}

type FaqStatisticsAttributes struct {
	ref terra.Reference
}

func (fs FaqStatisticsAttributes) InternalRef() terra.Reference {
	return fs.ref
}

func (fs FaqStatisticsAttributes) InternalWithRef(ref terra.Reference) FaqStatisticsAttributes {
	return FaqStatisticsAttributes{ref: ref}
}

func (fs FaqStatisticsAttributes) InternalTokens() hclwrite.Tokens {
	return fs.ref.InternalTokens()
}

func (fs FaqStatisticsAttributes) IndexedQuestionAnswersCount() terra.NumberValue {
	return terra.ReferenceNumber(fs.ref.Append("indexed_question_answers_count"))
}

type TextDocumentStatisticsAttributes struct {
	ref terra.Reference
}

func (tds TextDocumentStatisticsAttributes) InternalRef() terra.Reference {
	return tds.ref
}

func (tds TextDocumentStatisticsAttributes) InternalWithRef(ref terra.Reference) TextDocumentStatisticsAttributes {
	return TextDocumentStatisticsAttributes{ref: ref}
}

func (tds TextDocumentStatisticsAttributes) InternalTokens() hclwrite.Tokens {
	return tds.ref.InternalTokens()
}

func (tds TextDocumentStatisticsAttributes) IndexedTextBytes() terra.NumberValue {
	return terra.ReferenceNumber(tds.ref.Append("indexed_text_bytes"))
}

func (tds TextDocumentStatisticsAttributes) IndexedTextDocumentsCount() terra.NumberValue {
	return terra.ReferenceNumber(tds.ref.Append("indexed_text_documents_count"))
}

type ServerSideEncryptionConfigurationAttributes struct {
	ref terra.Reference
}

func (ssec ServerSideEncryptionConfigurationAttributes) InternalRef() terra.Reference {
	return ssec.ref
}

func (ssec ServerSideEncryptionConfigurationAttributes) InternalWithRef(ref terra.Reference) ServerSideEncryptionConfigurationAttributes {
	return ServerSideEncryptionConfigurationAttributes{ref: ref}
}

func (ssec ServerSideEncryptionConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return ssec.ref.InternalTokens()
}

func (ssec ServerSideEncryptionConfigurationAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceString(ssec.ref.Append("kms_key_id"))
}

type UserGroupResolutionConfigurationAttributes struct {
	ref terra.Reference
}

func (ugrc UserGroupResolutionConfigurationAttributes) InternalRef() terra.Reference {
	return ugrc.ref
}

func (ugrc UserGroupResolutionConfigurationAttributes) InternalWithRef(ref terra.Reference) UserGroupResolutionConfigurationAttributes {
	return UserGroupResolutionConfigurationAttributes{ref: ref}
}

func (ugrc UserGroupResolutionConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return ugrc.ref.InternalTokens()
}

func (ugrc UserGroupResolutionConfigurationAttributes) UserGroupResolutionMode() terra.StringValue {
	return terra.ReferenceString(ugrc.ref.Append("user_group_resolution_mode"))
}

type UserTokenConfigurationsAttributes struct {
	ref terra.Reference
}

func (utc UserTokenConfigurationsAttributes) InternalRef() terra.Reference {
	return utc.ref
}

func (utc UserTokenConfigurationsAttributes) InternalWithRef(ref terra.Reference) UserTokenConfigurationsAttributes {
	return UserTokenConfigurationsAttributes{ref: ref}
}

func (utc UserTokenConfigurationsAttributes) InternalTokens() hclwrite.Tokens {
	return utc.ref.InternalTokens()
}

func (utc UserTokenConfigurationsAttributes) JsonTokenTypeConfiguration() terra.ListValue[JsonTokenTypeConfigurationAttributes] {
	return terra.ReferenceList[JsonTokenTypeConfigurationAttributes](utc.ref.Append("json_token_type_configuration"))
}

func (utc UserTokenConfigurationsAttributes) JwtTokenTypeConfiguration() terra.ListValue[JwtTokenTypeConfigurationAttributes] {
	return terra.ReferenceList[JwtTokenTypeConfigurationAttributes](utc.ref.Append("jwt_token_type_configuration"))
}

type JsonTokenTypeConfigurationAttributes struct {
	ref terra.Reference
}

func (jttc JsonTokenTypeConfigurationAttributes) InternalRef() terra.Reference {
	return jttc.ref
}

func (jttc JsonTokenTypeConfigurationAttributes) InternalWithRef(ref terra.Reference) JsonTokenTypeConfigurationAttributes {
	return JsonTokenTypeConfigurationAttributes{ref: ref}
}

func (jttc JsonTokenTypeConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return jttc.ref.InternalTokens()
}

func (jttc JsonTokenTypeConfigurationAttributes) GroupAttributeField() terra.StringValue {
	return terra.ReferenceString(jttc.ref.Append("group_attribute_field"))
}

func (jttc JsonTokenTypeConfigurationAttributes) UserNameAttributeField() terra.StringValue {
	return terra.ReferenceString(jttc.ref.Append("user_name_attribute_field"))
}

type JwtTokenTypeConfigurationAttributes struct {
	ref terra.Reference
}

func (jttc JwtTokenTypeConfigurationAttributes) InternalRef() terra.Reference {
	return jttc.ref
}

func (jttc JwtTokenTypeConfigurationAttributes) InternalWithRef(ref terra.Reference) JwtTokenTypeConfigurationAttributes {
	return JwtTokenTypeConfigurationAttributes{ref: ref}
}

func (jttc JwtTokenTypeConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return jttc.ref.InternalTokens()
}

func (jttc JwtTokenTypeConfigurationAttributes) ClaimRegex() terra.StringValue {
	return terra.ReferenceString(jttc.ref.Append("claim_regex"))
}

func (jttc JwtTokenTypeConfigurationAttributes) GroupAttributeField() terra.StringValue {
	return terra.ReferenceString(jttc.ref.Append("group_attribute_field"))
}

func (jttc JwtTokenTypeConfigurationAttributes) Issuer() terra.StringValue {
	return terra.ReferenceString(jttc.ref.Append("issuer"))
}

func (jttc JwtTokenTypeConfigurationAttributes) KeyLocation() terra.StringValue {
	return terra.ReferenceString(jttc.ref.Append("key_location"))
}

func (jttc JwtTokenTypeConfigurationAttributes) SecretsManagerArn() terra.StringValue {
	return terra.ReferenceString(jttc.ref.Append("secrets_manager_arn"))
}

func (jttc JwtTokenTypeConfigurationAttributes) Url() terra.StringValue {
	return terra.ReferenceString(jttc.ref.Append("url"))
}

func (jttc JwtTokenTypeConfigurationAttributes) UserNameAttributeField() terra.StringValue {
	return terra.ReferenceString(jttc.ref.Append("user_name_attribute_field"))
}

type CapacityUnitsState struct {
	QueryCapacityUnits   float64 `json:"query_capacity_units"`
	StorageCapacityUnits float64 `json:"storage_capacity_units"`
}

type DocumentMetadataConfigurationUpdatesState struct {
	Name      string           `json:"name"`
	Type      string           `json:"type"`
	Relevance []RelevanceState `json:"relevance"`
	Search    []SearchState    `json:"search"`
}

type RelevanceState struct {
	Duration            string             `json:"duration"`
	Freshness           bool               `json:"freshness"`
	Importance          float64            `json:"importance"`
	RankOrder           string             `json:"rank_order"`
	ValuesImportanceMap map[string]float64 `json:"values_importance_map"`
}

type SearchState struct {
	Displayable bool `json:"displayable"`
	Facetable   bool `json:"facetable"`
	Searchable  bool `json:"searchable"`
	Sortable    bool `json:"sortable"`
}

type IndexStatisticsState struct {
	FaqStatistics          []FaqStatisticsState          `json:"faq_statistics"`
	TextDocumentStatistics []TextDocumentStatisticsState `json:"text_document_statistics"`
}

type FaqStatisticsState struct {
	IndexedQuestionAnswersCount float64 `json:"indexed_question_answers_count"`
}

type TextDocumentStatisticsState struct {
	IndexedTextBytes          float64 `json:"indexed_text_bytes"`
	IndexedTextDocumentsCount float64 `json:"indexed_text_documents_count"`
}

type ServerSideEncryptionConfigurationState struct {
	KmsKeyId string `json:"kms_key_id"`
}

type UserGroupResolutionConfigurationState struct {
	UserGroupResolutionMode string `json:"user_group_resolution_mode"`
}

type UserTokenConfigurationsState struct {
	JsonTokenTypeConfiguration []JsonTokenTypeConfigurationState `json:"json_token_type_configuration"`
	JwtTokenTypeConfiguration  []JwtTokenTypeConfigurationState  `json:"jwt_token_type_configuration"`
}

type JsonTokenTypeConfigurationState struct {
	GroupAttributeField    string `json:"group_attribute_field"`
	UserNameAttributeField string `json:"user_name_attribute_field"`
}

type JwtTokenTypeConfigurationState struct {
	ClaimRegex             string `json:"claim_regex"`
	GroupAttributeField    string `json:"group_attribute_field"`
	Issuer                 string `json:"issuer"`
	KeyLocation            string `json:"key_location"`
	SecretsManagerArn      string `json:"secrets_manager_arn"`
	Url                    string `json:"url"`
	UserNameAttributeField string `json:"user_name_attribute_field"`
}
