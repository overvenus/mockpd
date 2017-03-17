package server

import (
	"errors"
	"net"
	"net/url"

	"github.com/ngaut/log"
	"github.com/overvenus/mockpd/cases"
	pb "github.com/pingcap/kvproto/pkg/pdpb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Serve creates a mock PD server and listens on eps.
func Serve(eps []string, c cases.Case) {
	for _, e := range eps {
		u, err := url.Parse(e)
		if err != nil {
			log.Fatalf("%v", err)
		}

		lis, err := net.Listen("tcp", u.Host)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		var opts []grpc.ServerOption
		opts = append(opts, grpc.UnaryInterceptor(c.GetUnaryServerInterceptor()))
		// TODO: StreamServerInterceptor
		gsvr := grpc.NewServer(opts...)
		pb.RegisterPDServer(gsvr, new(mockPD))

		go func() {
			if err := gsvr.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
		}()
	}
}

var errUnimpl = errors.New("unimpl")

type mockPD struct{}

func (m *mockPD) GetMembers(ctx context.Context, req *pb.GetMembersRequest) (*pb.GetMembersResponse, error) {
	if f, ok := ctx.Value(cases.GetMembers).(func(*pb.GetMembersRequest) (*pb.GetMembersResponse, error)); ok {
		return f(req)
	}
	return nil, errUnimpl
}

func (m *mockPD) Tso(svr pb.PD_TsoServer) error {
	return errUnimpl
}

func (m *mockPD) Bootstrap(ctx context.Context, req *pb.BootstrapRequest) (*pb.BootstrapResponse, error) {
	if f, ok := ctx.Value(cases.Bootstrap).(func(*pb.BootstrapRequest) (*pb.BootstrapResponse, error)); ok {
		return f(req)
	}
	return nil, errUnimpl
}

func (m *mockPD) IsBootstrapped(ctx context.Context, req *pb.IsBootstrappedRequest) (*pb.IsBootstrappedResponse, error) {
	if f, ok := ctx.Value(cases.IsBootstrapped).(func(*pb.IsBootstrappedRequest) (*pb.IsBootstrappedResponse, error)); ok {
		return f(req)
	}
	return nil, errUnimpl
}

func (m *mockPD) AllocID(ctx context.Context, req *pb.AllocIDRequest) (*pb.AllocIDResponse, error) {
	if f, ok := ctx.Value(cases.AllocID).(func(*pb.AllocIDRequest) (*pb.AllocIDResponse, error)); ok {
		return f(req)
	}
	return nil, errUnimpl
}

func (m *mockPD) GetStore(ctx context.Context, req *pb.GetStoreRequest) (*pb.GetStoreResponse, error) {
	return nil, errUnimpl
}

func (m *mockPD) PutStore(ctx context.Context, req *pb.PutStoreRequest) (*pb.PutStoreResponse, error) {
	return nil, errUnimpl
}

func (m *mockPD) StoreHeartbeat(ctx context.Context, req *pb.StoreHeartbeatRequest) (*pb.StoreHeartbeatResponse, error) {
	return nil, errUnimpl
}

func (m *mockPD) RegionHeartbeat(ctx context.Context, req *pb.RegionHeartbeatRequest) (*pb.RegionHeartbeatResponse, error) {
	return nil, errUnimpl
}

func (m *mockPD) GetRegion(ctx context.Context, req *pb.GetRegionRequest) (*pb.GetRegionResponse, error) {
	return nil, errUnimpl
}

func (m *mockPD) GetRegionByID(ctx context.Context, req *pb.GetRegionByIDRequest) (*pb.GetRegionResponse, error) {
	return nil, errUnimpl
}

func (m *mockPD) AskSplit(ctx context.Context, req *pb.AskSplitRequest) (*pb.AskSplitResponse, error) {
	return nil, errUnimpl
}

func (m *mockPD) ReportSplit(ctx context.Context, req *pb.ReportSplitRequest) (*pb.ReportSplitResponse, error) {
	return nil, errUnimpl
}

func (m *mockPD) GetClusterConfig(ctx context.Context, req *pb.GetClusterConfigRequest) (*pb.GetClusterConfigResponse, error) {
	return nil, errUnimpl
}

func (m *mockPD) PutClusterConfig(ctx context.Context, req *pb.PutClusterConfigRequest) (*pb.PutClusterConfigResponse, error) {
	return nil, errUnimpl
}
