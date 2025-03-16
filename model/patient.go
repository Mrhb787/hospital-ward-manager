package model

import "time"

type Patient struct {
	ID               int              `json:"id"`
	PatientID        string           `json:"patient_id"`
	FirstName        string           `json:"first_name"`
	LastName         string           `json:"last_name"`
	DateOfBirth      *time.Time       `json:"date_of_birth"`
	Gender           string           `json:"gender"` // e.g., Male, Female, Other
	Address          string           `json:"address"`
	City             string           `json:"city"`
	State            string           `json:"state"`
	ZipCode          string           `json:"zip_code"`
	Country          string           `json:"country"`
	PhoneNumber      string           `json:"phone_number"`
	Email            string           `json:"email"`
	EmergencyContact EmergencyContact `json:"emergency_contact"`
	MedicalHistory   MedicalHistory   `json:"medical_history,omitempty"`
	Allergies        string           `json:"allergies,omitempty"`
	Notes            string           `json:"notes,omitempty"`
	AdmittedAt       *time.Time       `json:"admitted_at,omitempty"`
	DischargedAt     *time.Time       `json:"discharged_at,omitempty"`
	WardID           int              `json:"ward_id,omitempty"`
	RoomID           int              `json:"room_id,omitempty"`
	BedID            int              `json:"bed_id,omitempty"`
}

type EmergencyContact struct {
	Name         string `json:"name"`
	Relationship string `json:"relationship"`
	PhoneNumber  string `json:"phone_number"`
	Email        string `json:"email,omitempty"`
}

type MedicalHistory struct {
	Conditions    []Condition     `json:"conditions,omitempty"`
	Surgeries     []Surgery       `json:"surgeries,omitempty"`
	Medications   []Medication    `json:"medications,omitempty"`
	Immunizations []Immunization  `json:"immunizations,omitempty"`
	FamilyHistory []FamilyHistory `json:"family_history,omitempty"`
}

type Condition struct {
	Name                 string       `json:"name"`
	DiagnosisDate        time.Time    `json:"diagnosis_date,omitempty"`
	DiagnosisReportLinks []ReportLink `json:"diagnosis_report_links,omitempty"`
	Notes                string       `json:"notes,omitempty"`
}

type Surgery struct {
	Name  string    `json:"name"`
	Date  time.Time `json:"date,omitempty"`
	Notes string    `json:"notes,omitempty"`
}

type Medication struct {
	Name      string    `json:"name"`
	Dosage    string    `json:"dosage,omitempty"`
	StartDate time.Time `json:"start_date,omitempty"`
	EndDate   time.Time `json:"end_date,omitempty"`
	Notes     string    `json:"notes,omitempty"`
}

type Immunization struct {
	Name string    `json:"name"`
	Date time.Time `json:"date,omitempty"`
}

type FamilyHistory struct {
	Relative  string `json:"relative"`
	Condition string `json:"condition"`
}

type ReportLink struct {
	ReportId   int    `json:"report_id"`
	ReportType string `json:"report_type"`
	ReportURL  string `json:"report_url"`
}
