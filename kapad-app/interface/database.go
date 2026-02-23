package kapadinteface

import "context"

type Database interface {
	Create(context context.Context, collection string, data interface{})
	Get(context context.Context, collection string, docId string, dest interface{})
}
