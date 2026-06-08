package model

//RecordID defines a record id. Togethter with RecordType identifies unique records across all types.
type RecordID string

//RecordType defines a reacord type. Together with RecordID identifies unique records across all types.
type RecordType string

//Existing record types.
const (
	RecordTypeMovie = RecordType("movie")	//type conversion.
)

//UserID defines a user id.
type UserID string

//RatingValue defines a value of a rating record.
type RatingValue int

//Rating defines an individual rating created by a user for same record.
type Rating struct {
	RecordID string `json:"recordId"`
	RecordType string `json:"recordType"`
	UserID UserID `json:"userId"`
	Value RatingValue `json:"Value"`
}



