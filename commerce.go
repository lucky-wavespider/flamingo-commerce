package commerce

/*
	Flamingo Commerce Modules.
	The subpackage represent Flamingo Commerce modules - please refer to the documentations for the individual modules.
*/

//go:generate rm -rf docs/openapi
//go:generate go run github.com/swaggo/swag/cmd/swag init -p pascalcase --generalInfo=commerce.go --dir=./ --output=docs/openapi

// Swagger Documentation used for generator swag (https://github.com/swaggo/swag#declarative-comments-format)
// @title Flamingo Commerce API Spec
// @description Swagger (OpenAPI) Spec of all Flamingo Commerce modules
// @version 1.0
// @contact.name Flamingo
// @contact.url https://gitter.im/i-love-flamingo/community#
// @contact.email flamingo@aoe.com
// @license.name MIT
// @tag.name v1 Cart ajax API
// @tag.description This Cart APIs are most suitable to be called from a browser, because they rely on the session and cookie headers.
// @tag.name v1 Payment ajax API
// @tag.description This Payment APIs are most suitable to be called from a browser, because they rely on the session and cookie headers.
// @tag.name v1 Product API
// @tag.description Product API.
