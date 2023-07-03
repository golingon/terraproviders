// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package kendraindex

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type IndexStatistics struct {
	// FaqStatistics: min=0
	FaqStatistics []FaqStatistics `hcl:"faq_statistics,block" validate:"min=0"`
	// TextDocumentStatistics: min=0
	TextDocumentStatistics []TextDocumentStatistics `hcl:"text_document_statistics,block" validate:"min=0"`
}

type FaqStatistics struct{}

type TextDocumentStatistics struct{}

type CapacityUnits struct {
	// QueryCapacityUnits: number, optional
	QueryCapacityUnits terra.NumberValue `hcl:"query_capacity_units,attr"`
	// StorageCapacityUnits: number, optional
	StorageCapacityUnits terra.NumberValue `hcl:"storage_capacity_units,attr"`
}

type DocumentMetadataConfigurationUpdates struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Relevance: optional
	Relevance *Relevance `hcl:"relevance,block"`
	// Search: optional
	Search *Search `hcl:"search,block"`
}

type Relevance struct {
	// Duration: string, optional
	Duration terra.StringValue `hcl:"duration,attr"`
	// Freshness: bool, optional
	Freshness terra.BoolValue `hcl:"freshness,attr"`
	// Importance: number, optional
	Importance terra.NumberValue `hcl:"importance,attr"`
	// RankOrder: string, optional
	RankOrder terra.StringValue `hcl:"rank_order,attr"`
	// ValuesImportanceMap: map of number, optional
	ValuesImportanceMap terra.MapValue[terra.NumberValue] `hcl:"values_importance_map,attr"`
}

type Search struct {
	// Displayable: bool, optional
	Displayable terra.BoolValue `hcl:"displayable,attr"`
	// Facetable: bool, optional
	Facetable terra.BoolValue `hcl:"facetable,attr"`
	// Searchable: bool, optional
	Searchable terra.BoolValue `hcl:"searchable,attr"`
	// Sortable: bool, optional
	Sortable terra.BoolValue `hcl:"sortable,attr"`
}

type ServerSideEncryptionConfiguration struct {
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type UserGroupResolutionConfiguration struct {
	// UserGroupResolutionMode: string, required
	UserGroupResolutionMode terra.StringValue `hcl:"user_group_resolution_mode,attr" validate:"required"`
}

type UserTokenConfigurations struct {
	// JsonTokenTypeConfiguration: optional
	JsonTokenTypeConfiguration *JsonTokenTypeConfiguration `hcl:"json_token_type_configuration,block"`
	// JwtTokenTypeConfiguration: optional
	JwtTokenTypeConfiguration *JwtTokenTypeConfiguration `hcl:"jwt_token_type_configuration,block"`
}

type JsonTokenTypeConfiguration struct {
	// GroupAttributeField: string, required
	GroupAttributeField terra.StringValue `hcl:"group_attribute_field,attr" validate:"required"`
	// UserNameAttributeField: string, required
	UserNameAttributeField terra.StringValue `hcl:"user_name_attribute_field,attr" validate:"required"`
}

type JwtTokenTypeConfiguration struct {
	// ClaimRegex: string, optional
	ClaimRegex terra.StringValue `hcl:"claim_regex,attr"`
	// GroupAttributeField: string, optional
	GroupAttributeField terra.StringValue `hcl:"group_attribute_field,attr"`
	// Issuer: string, optional
	Issuer terra.StringValue `hcl:"issuer,attr"`
	// KeyLocation: string, required
	KeyLocation terra.StringValue `hcl:"key_location,attr" validate:"required"`
	// SecretsManagerArn: string, optional
	SecretsManagerArn terra.StringValue `hcl:"secrets_manager_arn,attr"`
	// Url: string, optional
	Url terra.StringValue `hcl:"url,attr"`
	// UserNameAttributeField: string, optional
	UserNameAttributeField terra.StringValue `hcl:"user_name_attribute_field,attr"`
}

type IndexStatisticsAttributes struct {
	ref terra.Reference
}

func (is IndexStatisticsAttributes) InternalRef() (terra.Reference, error) {
	return is.ref, nil
}

func (is IndexStatisticsAttributes) InternalWithRef(ref terra.Reference) IndexStatisticsAttributes {
	return IndexStatisticsAttributes{ref: ref}
}

func (is IndexStatisticsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return is.ref.InternalTokens()
}

func (is IndexStatisticsAttributes) FaqStatistics() terra.ListValue[FaqStatisticsAttributes] {
	return terra.ReferenceAsList[FaqStatisticsAttributes](is.ref.Append("faq_statistics"))
}

func (is IndexStatisticsAttributes) TextDocumentStatistics() terra.ListValue[TextDocumentStatisticsAttributes] {
	return terra.ReferenceAsList[TextDocumentStatisticsAttributes](is.ref.Append("text_document_statistics"))
}

type FaqStatisticsAttributes struct {
	ref terra.Reference
}

func (fs FaqStatisticsAttributes) InternalRef() (terra.Reference, error) {
	return fs.ref, nil
}

func (fs FaqStatisticsAttributes) InternalWithRef(ref terra.Reference) FaqStatisticsAttributes {
	return FaqStatisticsAttributes{ref: ref}
}

func (fs FaqStatisticsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fs.ref.InternalTokens()
}

func (fs FaqStatisticsAttributes) IndexedQuestionAnswersCount() terra.NumberValue {
	return terra.ReferenceAsNumber(fs.ref.Append("indexed_question_answers_count"))
}

type TextDocumentStatisticsAttributes struct {
	ref terra.Reference
}

func (tds TextDocumentStatisticsAttributes) InternalRef() (terra.Reference, error) {
	return tds.ref, nil
}

func (tds TextDocumentStatisticsAttributes) InternalWithRef(ref terra.Reference) TextDocumentStatisticsAttributes {
	return TextDocumentStatisticsAttributes{ref: ref}
}

func (tds TextDocumentStatisticsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tds.ref.InternalTokens()
}

func (tds TextDocumentStatisticsAttributes) IndexedTextBytes() terra.NumberValue {
	return terra.ReferenceAsNumber(tds.ref.Append("indexed_text_bytes"))
}

func (tds TextDocumentStatisticsAttributes) IndexedTextDocumentsCount() terra.NumberValue {
	return terra.ReferenceAsNumber(tds.ref.Append("indexed_text_documents_count"))
}

type CapacityUnitsAttributes struct {
	ref terra.Reference
}

func (cu CapacityUnitsAttributes) InternalRef() (terra.Reference, error) {
	return cu.ref, nil
}

func (cu CapacityUnitsAttributes) InternalWithRef(ref terra.Reference) CapacityUnitsAttributes {
	return CapacityUnitsAttributes{ref: ref}
}

func (cu CapacityUnitsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cu.ref.InternalTokens()
}

func (cu CapacityUnitsAttributes) QueryCapacityUnits() terra.NumberValue {
	return terra.ReferenceAsNumber(cu.ref.Append("query_capacity_units"))
}

func (cu CapacityUnitsAttributes) StorageCapacityUnits() terra.NumberValue {
	return terra.ReferenceAsNumber(cu.ref.Append("storage_capacity_units"))
}

type DocumentMetadataConfigurationUpdatesAttributes struct {
	ref terra.Reference
}

func (dmcu DocumentMetadataConfigurationUpdatesAttributes) InternalRef() (terra.Reference, error) {
	return dmcu.ref, nil
}

func (dmcu DocumentMetadataConfigurationUpdatesAttributes) InternalWithRef(ref terra.Reference) DocumentMetadataConfigurationUpdatesAttributes {
	return DocumentMetadataConfigurationUpdatesAttributes{ref: ref}
}

func (dmcu DocumentMetadataConfigurationUpdatesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dmcu.ref.InternalTokens()
}

func (dmcu DocumentMetadataConfigurationUpdatesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dmcu.ref.Append("name"))
}

func (dmcu DocumentMetadataConfigurationUpdatesAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(dmcu.ref.Append("type"))
}

func (dmcu DocumentMetadataConfigurationUpdatesAttributes) Relevance() terra.ListValue[RelevanceAttributes] {
	return terra.ReferenceAsList[RelevanceAttributes](dmcu.ref.Append("relevance"))
}

func (dmcu DocumentMetadataConfigurationUpdatesAttributes) Search() terra.ListValue[SearchAttributes] {
	return terra.ReferenceAsList[SearchAttributes](dmcu.ref.Append("search"))
}

type RelevanceAttributes struct {
	ref terra.Reference
}

func (r RelevanceAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RelevanceAttributes) InternalWithRef(ref terra.Reference) RelevanceAttributes {
	return RelevanceAttributes{ref: ref}
}

func (r RelevanceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RelevanceAttributes) Duration() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("duration"))
}

func (r RelevanceAttributes) Freshness() terra.BoolValue {
	return terra.ReferenceAsBool(r.ref.Append("freshness"))
}

func (r RelevanceAttributes) Importance() terra.NumberValue {
	return terra.ReferenceAsNumber(r.ref.Append("importance"))
}

func (r RelevanceAttributes) RankOrder() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("rank_order"))
}

func (r RelevanceAttributes) ValuesImportanceMap() terra.MapValue[terra.NumberValue] {
	return terra.ReferenceAsMap[terra.NumberValue](r.ref.Append("values_importance_map"))
}

type SearchAttributes struct {
	ref terra.Reference
}

func (s SearchAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SearchAttributes) InternalWithRef(ref terra.Reference) SearchAttributes {
	return SearchAttributes{ref: ref}
}

func (s SearchAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SearchAttributes) Displayable() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("displayable"))
}

func (s SearchAttributes) Facetable() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("facetable"))
}

func (s SearchAttributes) Searchable() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("searchable"))
}

func (s SearchAttributes) Sortable() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("sortable"))
}

type ServerSideEncryptionConfigurationAttributes struct {
	ref terra.Reference
}

func (ssec ServerSideEncryptionConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ssec.ref, nil
}

func (ssec ServerSideEncryptionConfigurationAttributes) InternalWithRef(ref terra.Reference) ServerSideEncryptionConfigurationAttributes {
	return ServerSideEncryptionConfigurationAttributes{ref: ref}
}

func (ssec ServerSideEncryptionConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ssec.ref.InternalTokens()
}

func (ssec ServerSideEncryptionConfigurationAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(ssec.ref.Append("kms_key_id"))
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

type UserGroupResolutionConfigurationAttributes struct {
	ref terra.Reference
}

func (ugrc UserGroupResolutionConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ugrc.ref, nil
}

func (ugrc UserGroupResolutionConfigurationAttributes) InternalWithRef(ref terra.Reference) UserGroupResolutionConfigurationAttributes {
	return UserGroupResolutionConfigurationAttributes{ref: ref}
}

func (ugrc UserGroupResolutionConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ugrc.ref.InternalTokens()
}

func (ugrc UserGroupResolutionConfigurationAttributes) UserGroupResolutionMode() terra.StringValue {
	return terra.ReferenceAsString(ugrc.ref.Append("user_group_resolution_mode"))
}

type UserTokenConfigurationsAttributes struct {
	ref terra.Reference
}

func (utc UserTokenConfigurationsAttributes) InternalRef() (terra.Reference, error) {
	return utc.ref, nil
}

func (utc UserTokenConfigurationsAttributes) InternalWithRef(ref terra.Reference) UserTokenConfigurationsAttributes {
	return UserTokenConfigurationsAttributes{ref: ref}
}

func (utc UserTokenConfigurationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return utc.ref.InternalTokens()
}

func (utc UserTokenConfigurationsAttributes) JsonTokenTypeConfiguration() terra.ListValue[JsonTokenTypeConfigurationAttributes] {
	return terra.ReferenceAsList[JsonTokenTypeConfigurationAttributes](utc.ref.Append("json_token_type_configuration"))
}

func (utc UserTokenConfigurationsAttributes) JwtTokenTypeConfiguration() terra.ListValue[JwtTokenTypeConfigurationAttributes] {
	return terra.ReferenceAsList[JwtTokenTypeConfigurationAttributes](utc.ref.Append("jwt_token_type_configuration"))
}

type JsonTokenTypeConfigurationAttributes struct {
	ref terra.Reference
}

func (jttc JsonTokenTypeConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return jttc.ref, nil
}

func (jttc JsonTokenTypeConfigurationAttributes) InternalWithRef(ref terra.Reference) JsonTokenTypeConfigurationAttributes {
	return JsonTokenTypeConfigurationAttributes{ref: ref}
}

func (jttc JsonTokenTypeConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return jttc.ref.InternalTokens()
}

func (jttc JsonTokenTypeConfigurationAttributes) GroupAttributeField() terra.StringValue {
	return terra.ReferenceAsString(jttc.ref.Append("group_attribute_field"))
}

func (jttc JsonTokenTypeConfigurationAttributes) UserNameAttributeField() terra.StringValue {
	return terra.ReferenceAsString(jttc.ref.Append("user_name_attribute_field"))
}

type JwtTokenTypeConfigurationAttributes struct {
	ref terra.Reference
}

func (jttc JwtTokenTypeConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return jttc.ref, nil
}

func (jttc JwtTokenTypeConfigurationAttributes) InternalWithRef(ref terra.Reference) JwtTokenTypeConfigurationAttributes {
	return JwtTokenTypeConfigurationAttributes{ref: ref}
}

func (jttc JwtTokenTypeConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return jttc.ref.InternalTokens()
}

func (jttc JwtTokenTypeConfigurationAttributes) ClaimRegex() terra.StringValue {
	return terra.ReferenceAsString(jttc.ref.Append("claim_regex"))
}

func (jttc JwtTokenTypeConfigurationAttributes) GroupAttributeField() terra.StringValue {
	return terra.ReferenceAsString(jttc.ref.Append("group_attribute_field"))
}

func (jttc JwtTokenTypeConfigurationAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(jttc.ref.Append("issuer"))
}

func (jttc JwtTokenTypeConfigurationAttributes) KeyLocation() terra.StringValue {
	return terra.ReferenceAsString(jttc.ref.Append("key_location"))
}

func (jttc JwtTokenTypeConfigurationAttributes) SecretsManagerArn() terra.StringValue {
	return terra.ReferenceAsString(jttc.ref.Append("secrets_manager_arn"))
}

func (jttc JwtTokenTypeConfigurationAttributes) Url() terra.StringValue {
	return terra.ReferenceAsString(jttc.ref.Append("url"))
}

func (jttc JwtTokenTypeConfigurationAttributes) UserNameAttributeField() terra.StringValue {
	return terra.ReferenceAsString(jttc.ref.Append("user_name_attribute_field"))
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

type ServerSideEncryptionConfigurationState struct {
	KmsKeyId string `json:"kms_key_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
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
