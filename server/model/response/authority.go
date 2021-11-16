package response

type ContextLinkResp struct {
	ContextLink string `json:"contextLink"`
}

type GetContextByLinkResp struct {
	ContextInfo
}