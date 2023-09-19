package modals

type Question struct {
	Description string `json:"description" xml:"description"`
	Id          string `json:"id" xml:"id"`
	Private     bool   `json:"is_private" xml:"is_private"`
	CreatedBy   string `json:"created_by" xml:"created_by"`
}
