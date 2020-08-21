package models

import (
	"time"
)

// Model is the same as gorm.Model, but includes json information
type Model struct {
	ID        uint       `gorm:"primary_key;column:id" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

// User expands Model and includes user information
type User struct {
	Model
	FirstName string `gorm:"column:first_name" json:"first_name"`
	LastName  string `gorm:"column:last_name" json:"last_name"`
	Address   string `gorm:"column:address" json:"address"`
	Email     string `gorm:"column:email" json:"email"`
}

// Survey includes data to pass to SurveyJS on the frontend
type Survey struct {
	Model
	Title             string `json:"title"`
	Pages             []Page `json:"pages"`
	ShowCompletedPage bool   `gorm:"column:show_completed_page" json:"show_completed_page"`
	NavigateToURL     string `gorm:"column:navigate_to_url" json:"navigate_to_url"`
}

// Page represents a page of the survey
type Page struct {
	Model
	SurveyID uint      `gorm:"column:survey_id" json:"survey_id"`
	Name     string    `gorm:"column:name" json:"name"`
	Elements []Element `json:"elements"`
}

// Element represents each element (i.e. question) on a page for a survey
type Element struct {
	Model
	PageID           uint     `gorm:"page_id" json:"page_id"`
	Type             string   `json:"type"`
	Name             string   `json:"name"`
	Title            string   `json:"title"`
	IsRequired       bool     `gorm:"column:is_required" json:"isRequired"`
	Choices          []Choice `json:"choices,omitempty"`
	HideNumber       bool     `gorm:"column:hide_number" json:"hideNumber,omitempty"`
	Columns          []Column `json:"columns,omitempty"`
	Rows             []Row    `json:"rows,omitempty"`
	IsAllRowRequired bool     `gorm:"column:is_all_row_required" json:"isAllRowRequired,omitempty"`
	Description      string   `json:"description,omitempty"`
}

// Choice represents different choices (i.e., answers) for an element
type Choice struct {
	Model
	ElementID uint   `gorm:"element_id"`
	Value     string `json:"value"`
	Text      string `json:"text"`
}

// Column represents options for Matrix-style questions
type Column struct {
	Model
	ElementID uint   `gorm:"element_id"`
	Value     string `json:"value"`
	Text      string `json:"text"`
}

// Row represents options for Matrix-style questions
type Row struct {
	Model
	ElementID uint   `gorm:"element_id"`
	Value     string `json:"value"`
	Text      string `json:"text"`
}
