package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound
var IsDeleted = 1
var IsNotDeleted = 0
