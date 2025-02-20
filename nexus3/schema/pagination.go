package schema

type PaginationResult[T any] struct {
	Items             []T     `json:"items"`
	ContinuationToken *string `json:"continuationToken"`
}
