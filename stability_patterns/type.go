package stability_pattern

import "context"

type Circuit func(context context.Context) (string, error)
