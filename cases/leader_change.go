package cases

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/ngaut/log"
	pb "github.com/pingcap/kvproto/pkg/pdpb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func NewLeaderChange(eps []string) Case {
	var ms []*pb.Member
	for i, e := range eps {
		m := &pb.Member{
			Name:       fmt.Sprintf("pd%d", i),
			MemberId:   uint64(100 + i),
			PeerUrls:   []string{e},
			ClientUrls: []string{e},
		}
		ms = append(ms, m)
	}

	// Add a dead PD
	deadEp := "http://127.0.0.1:65534"
	ms = append(ms, &pb.Member{
		Name:       "pd_dead",
		MemberId:   uint64(1000),
		PeerUrls:   []string{deadEp},
		ClientUrls: []string{deadEp},
	})

	var resps []*pb.GetMembersResponse
	for i := range ms {
		r := &pb.GetMembersResponse{
			Header:  &pb.ResponseHeader{ClusterId: uint64(1)},
			Members: ms,
			Leader:  ms[i],
		}
		resps = append(resps, r)
	}
	lc := &LeaderChange{
		resps: resps,
	}

	go func() {
		// TODO: quit.
		ticker := time.NewTicker(5 * time.Second)
		for {
			select {
			case <-ticker.C:
				atomic.AddUint64(&lc.idx, 1)
				resp, _ := lc.getMembers(nil)
				log.Debugf("[LeaderChange] switch to GetMembersResponse: %+v", resp)
			}
		}
	}()

	return lc
}

type LeaderChange struct {
	resps []*pb.GetMembersResponse
	idx   uint64
}

func (lc *LeaderChange) GetUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return lc.interceptor
}

func (lc *LeaderChange) getMembers(_ *pb.GetMembersRequest) (*pb.GetMembersResponse, error) {
	return lc.resps[int(atomic.LoadUint64(&lc.idx))%len(lc.resps)], nil
}

func (lc *LeaderChange) interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	ctx = context.WithValue(ctx, GetMembers, lc.getMembers)
	return handler(ctx, req)
}
