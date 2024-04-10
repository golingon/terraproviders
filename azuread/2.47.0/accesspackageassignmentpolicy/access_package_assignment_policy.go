// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package accesspackageassignmentpolicy

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ApprovalSettings struct {
	// ApprovalRequired: bool, optional
	ApprovalRequired terra.BoolValue `hcl:"approval_required,attr"`
	// ApprovalRequiredForExtension: bool, optional
	ApprovalRequiredForExtension terra.BoolValue `hcl:"approval_required_for_extension,attr"`
	// RequestorJustificationRequired: bool, optional
	RequestorJustificationRequired terra.BoolValue `hcl:"requestor_justification_required,attr"`
	// ApprovalStage: min=0
	ApprovalStage []ApprovalStage `hcl:"approval_stage,block" validate:"min=0"`
}

type ApprovalStage struct {
	// AlternativeApprovalEnabled: bool, optional
	AlternativeApprovalEnabled terra.BoolValue `hcl:"alternative_approval_enabled,attr"`
	// ApprovalTimeoutInDays: number, required
	ApprovalTimeoutInDays terra.NumberValue `hcl:"approval_timeout_in_days,attr" validate:"required"`
	// ApproverJustificationRequired: bool, optional
	ApproverJustificationRequired terra.BoolValue `hcl:"approver_justification_required,attr"`
	// EnableAlternativeApprovalInDays: number, optional
	EnableAlternativeApprovalInDays terra.NumberValue `hcl:"enable_alternative_approval_in_days,attr"`
	// AlternativeApprover: min=0
	AlternativeApprover []AlternativeApprover `hcl:"alternative_approver,block" validate:"min=0"`
	// PrimaryApprover: min=0
	PrimaryApprover []PrimaryApprover `hcl:"primary_approver,block" validate:"min=0"`
}

type AlternativeApprover struct {
	// Backup: bool, optional
	Backup terra.BoolValue `hcl:"backup,attr"`
	// ObjectId: string, optional
	ObjectId terra.StringValue `hcl:"object_id,attr"`
	// SubjectType: string, required
	SubjectType terra.StringValue `hcl:"subject_type,attr" validate:"required"`
}

type PrimaryApprover struct {
	// Backup: bool, optional
	Backup terra.BoolValue `hcl:"backup,attr"`
	// ObjectId: string, optional
	ObjectId terra.StringValue `hcl:"object_id,attr"`
	// SubjectType: string, required
	SubjectType terra.StringValue `hcl:"subject_type,attr" validate:"required"`
}

type AssignmentReviewSettings struct {
	// AccessRecommendationEnabled: bool, optional
	AccessRecommendationEnabled terra.BoolValue `hcl:"access_recommendation_enabled,attr"`
	// AccessReviewTimeoutBehavior: string, optional
	AccessReviewTimeoutBehavior terra.StringValue `hcl:"access_review_timeout_behavior,attr"`
	// ApproverJustificationRequired: bool, optional
	ApproverJustificationRequired terra.BoolValue `hcl:"approver_justification_required,attr"`
	// DurationInDays: number, optional
	DurationInDays terra.NumberValue `hcl:"duration_in_days,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// ReviewFrequency: string, optional
	ReviewFrequency terra.StringValue `hcl:"review_frequency,attr"`
	// ReviewType: string, optional
	ReviewType terra.StringValue `hcl:"review_type,attr"`
	// StartingOn: string, optional
	StartingOn terra.StringValue `hcl:"starting_on,attr"`
	// Reviewer: min=0
	Reviewer []Reviewer `hcl:"reviewer,block" validate:"min=0"`
}

type Reviewer struct {
	// Backup: bool, optional
	Backup terra.BoolValue `hcl:"backup,attr"`
	// ObjectId: string, optional
	ObjectId terra.StringValue `hcl:"object_id,attr"`
	// SubjectType: string, required
	SubjectType terra.StringValue `hcl:"subject_type,attr" validate:"required"`
}

type Question struct {
	// Required: bool, optional
	Required terra.BoolValue `hcl:"required,attr"`
	// Sequence: number, optional
	Sequence terra.NumberValue `hcl:"sequence,attr"`
	// Choice: min=0
	Choice []Choice `hcl:"choice,block" validate:"min=0"`
	// Text: required
	Text *Text `hcl:"text,block" validate:"required"`
}

type Choice struct {
	// ActualValue: string, required
	ActualValue terra.StringValue `hcl:"actual_value,attr" validate:"required"`
	// DisplayValue: required
	DisplayValue *DisplayValue `hcl:"display_value,block" validate:"required"`
}

type DisplayValue struct {
	// DefaultText: string, required
	DefaultText terra.StringValue `hcl:"default_text,attr" validate:"required"`
	// DisplayValueLocalizedText: min=0
	LocalizedText []DisplayValueLocalizedText `hcl:"localized_text,block" validate:"min=0"`
}

type DisplayValueLocalizedText struct {
	// Content: string, required
	Content terra.StringValue `hcl:"content,attr" validate:"required"`
	// LanguageCode: string, required
	LanguageCode terra.StringValue `hcl:"language_code,attr" validate:"required"`
}

type Text struct {
	// DefaultText: string, required
	DefaultText terra.StringValue `hcl:"default_text,attr" validate:"required"`
	// TextLocalizedText: min=0
	LocalizedText []TextLocalizedText `hcl:"localized_text,block" validate:"min=0"`
}

type TextLocalizedText struct {
	// Content: string, required
	Content terra.StringValue `hcl:"content,attr" validate:"required"`
	// LanguageCode: string, required
	LanguageCode terra.StringValue `hcl:"language_code,attr" validate:"required"`
}

type RequestorSettings struct {
	// RequestsAccepted: bool, optional
	RequestsAccepted terra.BoolValue `hcl:"requests_accepted,attr"`
	// ScopeType: string, optional
	ScopeType terra.StringValue `hcl:"scope_type,attr"`
	// Requestor: min=0
	Requestor []Requestor `hcl:"requestor,block" validate:"min=0"`
}

type Requestor struct {
	// Backup: bool, optional
	Backup terra.BoolValue `hcl:"backup,attr"`
	// ObjectId: string, optional
	ObjectId terra.StringValue `hcl:"object_id,attr"`
	// SubjectType: string, required
	SubjectType terra.StringValue `hcl:"subject_type,attr" validate:"required"`
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

type ApprovalSettingsAttributes struct {
	ref terra.Reference
}

func (as ApprovalSettingsAttributes) InternalRef() (terra.Reference, error) {
	return as.ref, nil
}

func (as ApprovalSettingsAttributes) InternalWithRef(ref terra.Reference) ApprovalSettingsAttributes {
	return ApprovalSettingsAttributes{ref: ref}
}

func (as ApprovalSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return as.ref.InternalTokens()
}

func (as ApprovalSettingsAttributes) ApprovalRequired() terra.BoolValue {
	return terra.ReferenceAsBool(as.ref.Append("approval_required"))
}

func (as ApprovalSettingsAttributes) ApprovalRequiredForExtension() terra.BoolValue {
	return terra.ReferenceAsBool(as.ref.Append("approval_required_for_extension"))
}

func (as ApprovalSettingsAttributes) RequestorJustificationRequired() terra.BoolValue {
	return terra.ReferenceAsBool(as.ref.Append("requestor_justification_required"))
}

func (as ApprovalSettingsAttributes) ApprovalStage() terra.ListValue[ApprovalStageAttributes] {
	return terra.ReferenceAsList[ApprovalStageAttributes](as.ref.Append("approval_stage"))
}

type ApprovalStageAttributes struct {
	ref terra.Reference
}

func (as ApprovalStageAttributes) InternalRef() (terra.Reference, error) {
	return as.ref, nil
}

func (as ApprovalStageAttributes) InternalWithRef(ref terra.Reference) ApprovalStageAttributes {
	return ApprovalStageAttributes{ref: ref}
}

func (as ApprovalStageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return as.ref.InternalTokens()
}

func (as ApprovalStageAttributes) AlternativeApprovalEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(as.ref.Append("alternative_approval_enabled"))
}

func (as ApprovalStageAttributes) ApprovalTimeoutInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(as.ref.Append("approval_timeout_in_days"))
}

func (as ApprovalStageAttributes) ApproverJustificationRequired() terra.BoolValue {
	return terra.ReferenceAsBool(as.ref.Append("approver_justification_required"))
}

func (as ApprovalStageAttributes) EnableAlternativeApprovalInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(as.ref.Append("enable_alternative_approval_in_days"))
}

func (as ApprovalStageAttributes) AlternativeApprover() terra.ListValue[AlternativeApproverAttributes] {
	return terra.ReferenceAsList[AlternativeApproverAttributes](as.ref.Append("alternative_approver"))
}

func (as ApprovalStageAttributes) PrimaryApprover() terra.ListValue[PrimaryApproverAttributes] {
	return terra.ReferenceAsList[PrimaryApproverAttributes](as.ref.Append("primary_approver"))
}

type AlternativeApproverAttributes struct {
	ref terra.Reference
}

func (aa AlternativeApproverAttributes) InternalRef() (terra.Reference, error) {
	return aa.ref, nil
}

func (aa AlternativeApproverAttributes) InternalWithRef(ref terra.Reference) AlternativeApproverAttributes {
	return AlternativeApproverAttributes{ref: ref}
}

func (aa AlternativeApproverAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aa.ref.InternalTokens()
}

func (aa AlternativeApproverAttributes) Backup() terra.BoolValue {
	return terra.ReferenceAsBool(aa.ref.Append("backup"))
}

func (aa AlternativeApproverAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("object_id"))
}

func (aa AlternativeApproverAttributes) SubjectType() terra.StringValue {
	return terra.ReferenceAsString(aa.ref.Append("subject_type"))
}

type PrimaryApproverAttributes struct {
	ref terra.Reference
}

func (pa PrimaryApproverAttributes) InternalRef() (terra.Reference, error) {
	return pa.ref, nil
}

func (pa PrimaryApproverAttributes) InternalWithRef(ref terra.Reference) PrimaryApproverAttributes {
	return PrimaryApproverAttributes{ref: ref}
}

func (pa PrimaryApproverAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pa.ref.InternalTokens()
}

func (pa PrimaryApproverAttributes) Backup() terra.BoolValue {
	return terra.ReferenceAsBool(pa.ref.Append("backup"))
}

func (pa PrimaryApproverAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("object_id"))
}

func (pa PrimaryApproverAttributes) SubjectType() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("subject_type"))
}

type AssignmentReviewSettingsAttributes struct {
	ref terra.Reference
}

func (ars AssignmentReviewSettingsAttributes) InternalRef() (terra.Reference, error) {
	return ars.ref, nil
}

func (ars AssignmentReviewSettingsAttributes) InternalWithRef(ref terra.Reference) AssignmentReviewSettingsAttributes {
	return AssignmentReviewSettingsAttributes{ref: ref}
}

func (ars AssignmentReviewSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ars.ref.InternalTokens()
}

func (ars AssignmentReviewSettingsAttributes) AccessRecommendationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ars.ref.Append("access_recommendation_enabled"))
}

func (ars AssignmentReviewSettingsAttributes) AccessReviewTimeoutBehavior() terra.StringValue {
	return terra.ReferenceAsString(ars.ref.Append("access_review_timeout_behavior"))
}

func (ars AssignmentReviewSettingsAttributes) ApproverJustificationRequired() terra.BoolValue {
	return terra.ReferenceAsBool(ars.ref.Append("approver_justification_required"))
}

func (ars AssignmentReviewSettingsAttributes) DurationInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(ars.ref.Append("duration_in_days"))
}

func (ars AssignmentReviewSettingsAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ars.ref.Append("enabled"))
}

func (ars AssignmentReviewSettingsAttributes) ReviewFrequency() terra.StringValue {
	return terra.ReferenceAsString(ars.ref.Append("review_frequency"))
}

func (ars AssignmentReviewSettingsAttributes) ReviewType() terra.StringValue {
	return terra.ReferenceAsString(ars.ref.Append("review_type"))
}

func (ars AssignmentReviewSettingsAttributes) StartingOn() terra.StringValue {
	return terra.ReferenceAsString(ars.ref.Append("starting_on"))
}

func (ars AssignmentReviewSettingsAttributes) Reviewer() terra.ListValue[ReviewerAttributes] {
	return terra.ReferenceAsList[ReviewerAttributes](ars.ref.Append("reviewer"))
}

type ReviewerAttributes struct {
	ref terra.Reference
}

func (r ReviewerAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r ReviewerAttributes) InternalWithRef(ref terra.Reference) ReviewerAttributes {
	return ReviewerAttributes{ref: ref}
}

func (r ReviewerAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r ReviewerAttributes) Backup() terra.BoolValue {
	return terra.ReferenceAsBool(r.ref.Append("backup"))
}

func (r ReviewerAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("object_id"))
}

func (r ReviewerAttributes) SubjectType() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("subject_type"))
}

type QuestionAttributes struct {
	ref terra.Reference
}

func (q QuestionAttributes) InternalRef() (terra.Reference, error) {
	return q.ref, nil
}

func (q QuestionAttributes) InternalWithRef(ref terra.Reference) QuestionAttributes {
	return QuestionAttributes{ref: ref}
}

func (q QuestionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return q.ref.InternalTokens()
}

func (q QuestionAttributes) Required() terra.BoolValue {
	return terra.ReferenceAsBool(q.ref.Append("required"))
}

func (q QuestionAttributes) Sequence() terra.NumberValue {
	return terra.ReferenceAsNumber(q.ref.Append("sequence"))
}

func (q QuestionAttributes) Choice() terra.ListValue[ChoiceAttributes] {
	return terra.ReferenceAsList[ChoiceAttributes](q.ref.Append("choice"))
}

func (q QuestionAttributes) Text() terra.ListValue[TextAttributes] {
	return terra.ReferenceAsList[TextAttributes](q.ref.Append("text"))
}

type ChoiceAttributes struct {
	ref terra.Reference
}

func (c ChoiceAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ChoiceAttributes) InternalWithRef(ref terra.Reference) ChoiceAttributes {
	return ChoiceAttributes{ref: ref}
}

func (c ChoiceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ChoiceAttributes) ActualValue() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("actual_value"))
}

func (c ChoiceAttributes) DisplayValue() terra.ListValue[DisplayValueAttributes] {
	return terra.ReferenceAsList[DisplayValueAttributes](c.ref.Append("display_value"))
}

type DisplayValueAttributes struct {
	ref terra.Reference
}

func (dv DisplayValueAttributes) InternalRef() (terra.Reference, error) {
	return dv.ref, nil
}

func (dv DisplayValueAttributes) InternalWithRef(ref terra.Reference) DisplayValueAttributes {
	return DisplayValueAttributes{ref: ref}
}

func (dv DisplayValueAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dv.ref.InternalTokens()
}

func (dv DisplayValueAttributes) DefaultText() terra.StringValue {
	return terra.ReferenceAsString(dv.ref.Append("default_text"))
}

func (dv DisplayValueAttributes) LocalizedText() terra.ListValue[DisplayValueLocalizedTextAttributes] {
	return terra.ReferenceAsList[DisplayValueLocalizedTextAttributes](dv.ref.Append("localized_text"))
}

type DisplayValueLocalizedTextAttributes struct {
	ref terra.Reference
}

func (lt DisplayValueLocalizedTextAttributes) InternalRef() (terra.Reference, error) {
	return lt.ref, nil
}

func (lt DisplayValueLocalizedTextAttributes) InternalWithRef(ref terra.Reference) DisplayValueLocalizedTextAttributes {
	return DisplayValueLocalizedTextAttributes{ref: ref}
}

func (lt DisplayValueLocalizedTextAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lt.ref.InternalTokens()
}

func (lt DisplayValueLocalizedTextAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("content"))
}

func (lt DisplayValueLocalizedTextAttributes) LanguageCode() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("language_code"))
}

type TextAttributes struct {
	ref terra.Reference
}

func (t TextAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TextAttributes) InternalWithRef(ref terra.Reference) TextAttributes {
	return TextAttributes{ref: ref}
}

func (t TextAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TextAttributes) DefaultText() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("default_text"))
}

func (t TextAttributes) LocalizedText() terra.ListValue[TextLocalizedTextAttributes] {
	return terra.ReferenceAsList[TextLocalizedTextAttributes](t.ref.Append("localized_text"))
}

type TextLocalizedTextAttributes struct {
	ref terra.Reference
}

func (lt TextLocalizedTextAttributes) InternalRef() (terra.Reference, error) {
	return lt.ref, nil
}

func (lt TextLocalizedTextAttributes) InternalWithRef(ref terra.Reference) TextLocalizedTextAttributes {
	return TextLocalizedTextAttributes{ref: ref}
}

func (lt TextLocalizedTextAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lt.ref.InternalTokens()
}

func (lt TextLocalizedTextAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("content"))
}

func (lt TextLocalizedTextAttributes) LanguageCode() terra.StringValue {
	return terra.ReferenceAsString(lt.ref.Append("language_code"))
}

type RequestorSettingsAttributes struct {
	ref terra.Reference
}

func (rs RequestorSettingsAttributes) InternalRef() (terra.Reference, error) {
	return rs.ref, nil
}

func (rs RequestorSettingsAttributes) InternalWithRef(ref terra.Reference) RequestorSettingsAttributes {
	return RequestorSettingsAttributes{ref: ref}
}

func (rs RequestorSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rs.ref.InternalTokens()
}

func (rs RequestorSettingsAttributes) RequestsAccepted() terra.BoolValue {
	return terra.ReferenceAsBool(rs.ref.Append("requests_accepted"))
}

func (rs RequestorSettingsAttributes) ScopeType() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("scope_type"))
}

func (rs RequestorSettingsAttributes) Requestor() terra.ListValue[RequestorAttributes] {
	return terra.ReferenceAsList[RequestorAttributes](rs.ref.Append("requestor"))
}

type RequestorAttributes struct {
	ref terra.Reference
}

func (r RequestorAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RequestorAttributes) InternalWithRef(ref terra.Reference) RequestorAttributes {
	return RequestorAttributes{ref: ref}
}

func (r RequestorAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RequestorAttributes) Backup() terra.BoolValue {
	return terra.ReferenceAsBool(r.ref.Append("backup"))
}

func (r RequestorAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("object_id"))
}

func (r RequestorAttributes) SubjectType() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("subject_type"))
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

type ApprovalSettingsState struct {
	ApprovalRequired               bool                 `json:"approval_required"`
	ApprovalRequiredForExtension   bool                 `json:"approval_required_for_extension"`
	RequestorJustificationRequired bool                 `json:"requestor_justification_required"`
	ApprovalStage                  []ApprovalStageState `json:"approval_stage"`
}

type ApprovalStageState struct {
	AlternativeApprovalEnabled      bool                       `json:"alternative_approval_enabled"`
	ApprovalTimeoutInDays           float64                    `json:"approval_timeout_in_days"`
	ApproverJustificationRequired   bool                       `json:"approver_justification_required"`
	EnableAlternativeApprovalInDays float64                    `json:"enable_alternative_approval_in_days"`
	AlternativeApprover             []AlternativeApproverState `json:"alternative_approver"`
	PrimaryApprover                 []PrimaryApproverState     `json:"primary_approver"`
}

type AlternativeApproverState struct {
	Backup      bool   `json:"backup"`
	ObjectId    string `json:"object_id"`
	SubjectType string `json:"subject_type"`
}

type PrimaryApproverState struct {
	Backup      bool   `json:"backup"`
	ObjectId    string `json:"object_id"`
	SubjectType string `json:"subject_type"`
}

type AssignmentReviewSettingsState struct {
	AccessRecommendationEnabled   bool            `json:"access_recommendation_enabled"`
	AccessReviewTimeoutBehavior   string          `json:"access_review_timeout_behavior"`
	ApproverJustificationRequired bool            `json:"approver_justification_required"`
	DurationInDays                float64         `json:"duration_in_days"`
	Enabled                       bool            `json:"enabled"`
	ReviewFrequency               string          `json:"review_frequency"`
	ReviewType                    string          `json:"review_type"`
	StartingOn                    string          `json:"starting_on"`
	Reviewer                      []ReviewerState `json:"reviewer"`
}

type ReviewerState struct {
	Backup      bool   `json:"backup"`
	ObjectId    string `json:"object_id"`
	SubjectType string `json:"subject_type"`
}

type QuestionState struct {
	Required bool          `json:"required"`
	Sequence float64       `json:"sequence"`
	Choice   []ChoiceState `json:"choice"`
	Text     []TextState   `json:"text"`
}

type ChoiceState struct {
	ActualValue  string              `json:"actual_value"`
	DisplayValue []DisplayValueState `json:"display_value"`
}

type DisplayValueState struct {
	DefaultText   string                           `json:"default_text"`
	LocalizedText []DisplayValueLocalizedTextState `json:"localized_text"`
}

type DisplayValueLocalizedTextState struct {
	Content      string `json:"content"`
	LanguageCode string `json:"language_code"`
}

type TextState struct {
	DefaultText   string                   `json:"default_text"`
	LocalizedText []TextLocalizedTextState `json:"localized_text"`
}

type TextLocalizedTextState struct {
	Content      string `json:"content"`
	LanguageCode string `json:"language_code"`
}

type RequestorSettingsState struct {
	RequestsAccepted bool             `json:"requests_accepted"`
	ScopeType        string           `json:"scope_type"`
	Requestor        []RequestorState `json:"requestor"`
}

type RequestorState struct {
	Backup      bool   `json:"backup"`
	ObjectId    string `json:"object_id"`
	SubjectType string `json:"subject_type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
