package requests

type PersonDocumntRequest struct {
	File string `json:"file"'`
	Filetype string `json:"filetype"`
	PersonId string `json:"person_id"`
	Name string	`json:"name"`
	Description string `json:"description"`
}

//func (r PersonDocumntRequest)Convert(c echo.Context)*PersonDocumntRequest  {
//
//}