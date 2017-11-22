package opendota

import (
	"net/http"

	"github.com/dghubble/sling"
)

func newMetadataService(sling *sling.Sling) *MetadataService {
	return &MetadataService{
		sling: sling.Path("metadata"),
	}
}

// MetadataService provides methods for the site metadata.
type MetadataService struct {
	sling *sling.Sling
}

// Metadata describes information about the site
type Metadata struct {
	Banner string `json:"banner"`
	Cheese cheese `json:"cheese"`
}

type cheese struct {
	Cheese string `json:"cheese"`
	Goal   string `json:"goal"`
}

// Metadata returns a collection of metadata about the site.
// https://docs.opendota.com/#tag/metadata%2Fpaths%2F~1metadata%2Fget
func (s *MetadataService) Metadata() (Metadata, *http.Response, error) {
	metadata := new(Metadata)
	apiError := new(APIError)
	resp, err := s.sling.New().Receive(metadata, apiError)
	return *metadata, resp, relevantError(err, *apiError)
}
