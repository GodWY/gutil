package gutil

import "context"

type Log func(ctx context.Context, info ...interface{})

var logHandler Log

func InstallLog(ll Log) {
	logHandler = ll
}
