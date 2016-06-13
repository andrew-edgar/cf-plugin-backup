package models

const UrlSuffix string = "_url"

type ResourceCollectionModel struct {
	TotalResults int               `json:"total_results"`
	TotalPages   int               `json:"total_pages"`
	NextURL      string            `json:"next_url,omitempty"`
	PrevURL      string            `json:"prev_url,omitempty"`
	Resources    *[]*ResourceModel `json:"resources"`
}

type ResourceModel struct {
	Metadata MetadataModel          `json:"metadata"`
	Entity   map[string]interface{} `json:"entity"`
}

type MetadataModel struct {
	Guid      string `json:"guid"`
	Url       string `json:"url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type BackupModel struct {
	Organizations interface{} `json:"organizations"`
}
