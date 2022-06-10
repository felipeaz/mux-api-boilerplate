package models

// SampleModel is a model example
type SampleModel struct {
	ID          string
	Name        string
	Description string
}

// NewSample returns an instance of SampleModel
func NewSample() SampleModel {
	return SampleModel{
		ID:          "sample-id",
		Name:        "Sample Model",
		Description: "This is a Sample Object",
	}
}

// NewSampleArray returns an instance of SampleModel
func NewSampleArray() []SampleModel {
	return []SampleModel{
		{
			ID:          "sample-id",
			Name:        "Sample Model",
			Description: "This is a Sample Object",
		},
		{
			ID:          "sample-id-2",
			Name:        "Sample Model 2",
			Description: "This is a Sample Object 2",
		},
		{
			ID:          "sample-id-3",
			Name:        "Sample Model 3",
			Description: "This is a Sample Object 3",
		},
	}
}
