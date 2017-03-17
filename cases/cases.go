package cases

import (
	"google.golang.org/grpc"
)

var (
	GetMembers       = &struct{}{}
	Tso              = &struct{}{}
	Bootstrap        = &struct{}{}
	IsBootstrapped   = &struct{}{}
	AllocID          = &struct{}{}
	GetStore         = &struct{}{}
	PutStore         = &struct{}{}
	StoreHeartbeat   = &struct{}{}
	RegionHeartbeat  = &struct{}{}
	GetRegion        = &struct{}{}
	GetRegionByID    = &struct{}{}
	AskSplit         = &struct{}{}
	ReportSplit      = &struct{}{}
	GetClusterConfig = &struct{}{}
	PutClusterConfig = &struct{}{}
)

// Case is a mock case.
type Case interface {
	GetUnaryServerInterceptor() grpc.UnaryServerInterceptor
}
