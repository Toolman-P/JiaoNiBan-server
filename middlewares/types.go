package middlewares

type IndexQuery struct {
	Author string `form:"author"`
}

type DescQuery struct {
	Author string `form:"author"`
	Page   int    `form:"page"`
}

type ContentQuery struct {
	Author string `form:"author"`
	Hash   string `form:"id"`
}
