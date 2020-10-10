package pet

// Pet defines pet attributes
type Pet struct {
	Nome  string `json:"nome"`
	Breed Breed  `json:"breed"`
	Size  Size   `json:"size"`
}
