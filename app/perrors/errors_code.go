package perrors

type ErrorCode int

// Rule
// The error code has 6 charactors
// 2 first charactors is for handler group, example:
// next 2 charactors is for package: handler = 01, service = 02, repository = 03
// 2 last charactors is for error no
// Example: 100301
// 01: Article
// 03: Repository
// 01: NotFound

const (
	// Common Error
	NotFound ErrorCode = 404
	// Article
	ArticleInvalidIDRepository ErrorCode = 100301
)
