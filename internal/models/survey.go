package models

type SurveyInfo struct {
	SurveyID    string `json:"surveyID,omitempty" db:"survey_id"`
	SurveyTitle string `json:"surveyTitle,omitempty" db:"survey_title"`
	CompanyLogo string `json:"companyLogo,omitempty" db:"company_logo"`
}

type Survey struct {
	SurveyID    string      `json:"surveyID" db:"survey_id"`
	SurveyTitle string      `json:"surveyTitle,omitempty" db:"survey_title"`
	CompanyLogo string      `json:"company_logo" db:"company_logo"`
	Questions   []*Question `json:"questions,omitempty" db:"questions"`
}

type Question struct {
	ID      uint64    `json:"id,omitempty"`
	Title   string    `json:"title,omitempty"`
	Options []*Option `json:"options,omitempty"`
}

type Option struct {
	ID   uint64 `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
}
