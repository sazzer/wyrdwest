package database

// RecordNotFoundError represents when querying a single record and none could be found
type RecordNotFoundError struct{}

// Error returns the error message for the RecordNotFoundError
func (e RecordNotFoundError) Error() string {
	return "Expected exactly one record but none found"
}

// MultipleRecordFoundError represents when querying a single record and 2+ were found
type MultipleRecordFoundError struct{}

// Error returns the error message for the MultipleRecordFoundError
func (e MultipleRecordFoundError) Error() string {
	return "Expected exactly one record but 2+ found"
}
