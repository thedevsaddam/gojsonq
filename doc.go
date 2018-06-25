// This work is covered by the MIT License.

// Package gojsonq lets you quickly query over JSON Data.
// This package provides very expressive queries with method chaining.
//
// Creating
// Start with New() and add a source with File(), JSONString() or Reader().
// You can copy existing JSONQs with the Copy() method.
// Reset() returns an object to its original state.
//
// Erros
// Check for errors with the Error() and Errors() methods.
//
// Query operations
// You can filter with the Where*() methods and the OrWhere() method.
// You sort with Sort() and SortBy().
// From(), Select() and GroupBy() provide additional selection mechanisms.
//
// Array operations
// Count(), Avg(), Sum(), Min(), Max()
//
// Receiving values
// First(), Last(), Nth(), Get(), Find(), Only(), Pluck()
//
// String
// String() formats a JSONQ as valid JSON without newlines, indentation and
// unnecessary white-space.
//
// For further information refer to the github-repository.
package gojsonq
