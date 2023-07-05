package storage

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"os"
)

type Storage struct {
	ProjectURL         string
	ProjectAPIKey      string
	ProjectStorageName string
	StorageFolder      string
}

func newSupabaseStorage() (*Storage, error) {
	var err error
	cfgPayment := Storage{
		ProjectURL:         os.Getenv("SUPABASE_PROJECT_URL"),
		ProjectAPIKey:      os.Getenv("SUPABASE_PROJECT_API_KEY"),
		ProjectStorageName: os.Getenv("SUPABASE_PROJECT_STORAGE_NAME"),
		StorageFolder:      os.Getenv("SUPABASE_STORAGE_FOLDER"),
	}

	err = cfgPayment.validate()
	if err != nil {
		return nil, err
	}

	return &cfgPayment, err
}

func (p Storage) validate() error {
	return validation.ValidateStruct(
		&p,
		validation.Field(&p.ProjectURL, validation.Required),
		validation.Field(&p.ProjectAPIKey, validation.Required),
		validation.Field(&p.ProjectStorageName, validation.Required),
		validation.Field(&p.StorageFolder, validation.Required),
	)
}
