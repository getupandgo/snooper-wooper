package validators

// todo: implement checking on empty param
type GetTopTokens struct {
	Limit  uint64 `form:"user" json:"user" xml:"user"  binding:"required"`
	Offset uint64 `form:"password" json:"password" xml:"password" binding:"required"`
}

//todo: merge validator and model somehow?
type Token struct {
	Text  string `json:"text"`
	Count uint64 `json:"count"`
}
