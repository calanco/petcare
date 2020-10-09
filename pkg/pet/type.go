package pet

// Pet defines pet attributes
type Pet struct {
	Nome   string `json:"nome"`
	Taglia string `json:"size,omitempty"`
	Razza  string `json:"razza,omitempty"`
}
