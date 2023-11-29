package requests

type Address struct {
	AddressID			int     `json:"AddressID"`
	ValidityStartDate	string  `json:"ValidityStartDate"`
	ValidityEndDate		string	`json:"ValidityEndDate"`
	IsMarkedForDeletion *bool  `json:IsMarkedForDeletion`
}
