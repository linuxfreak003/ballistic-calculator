package repository

// PageInfo is an interface the represents the information
// needed for a page
type PageInfo interface {
	GetPageSize() uint32
	GetPageToken() string
}

// ListRequest represents a list request.
type ListRequest struct {
	PageSize  uint32
	PageToken string
}

// FromProto sets ListRequest values from anything
// with a page size and page token.
func (r *ListRequest) FromProto(info PageInfo) *ListRequest {
	if r == nil {
		r = new(ListRequest)
	}

	r.PageSize = info.GetPageSize()
	r.PageToken = info.GetPageToken()
	return r
}

// GetPageSize is a nil-safe getter for page size
func (r *ListRequest) GetPageSize() uint32 {
	if r != nil {
		return r.PageSize
	}
	return 0
}

// GetPageToken is a nil-safe getter for page token
func (r *ListRequest) GetPageToken() string {
	if r != nil {
		return r.PageToken
	}
	return ""
}

// ListResponse represents a response
type ListResponse struct {
	NextPageToken string
}

// GetNextPageToken is a nil-safe getter for next page token
func (r *ListResponse) GetNextPageToken() string {
	if r != nil {
		return r.NextPageToken
	}
	return ""
}
