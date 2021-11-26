package gcache

import (
	"github.com/localrivet/galaxycache"
	galaxGrpc "github.com/localrivet/galaxycache/grpc"
	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
)

// create the universe!
// NewUniverse
func NewUniverse(listenOn string) *galaxycache.Universe {
	grpcFetchProtocol := galaxGrpc.NewGRPCFetchProtocol(grpc.WithInsecure())
	universe := galaxycache.NewUniverse(grpcFetchProtocol, listenOn)
	grpcServer := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))
	galaxGrpc.RegisterGRPCServer(universe, grpcServer)
	return universe
}

// register a new galaxy
// RegisterGalaxyFunc
func RegisterGalaxyFunc(key string, universe *galaxycache.Universe, getter galaxycache.GetterFunc) *galaxycache.Galaxy {
	return universe.NewGalaxy(key, 1<<20, getter)
}
