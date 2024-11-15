package geo

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/hailocab/go-geoindex"
	"github.com/opentracing/opentracing-go"
	pb "github.com/ucy-coast/hotel-app/internal/geo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

// Geo implements the geo service
type Geo struct {
	port      int
	addr      string
	dbsession *DatabaseSession
	tracer    opentracing.Tracer
	geoidx    *geoindex.ClusteringIndex
}

const (
	maxSearchRadius  = 10
	maxSearchResults = 5
)

// NewGeo returns a new server
func NewGeo(a string, p int, db *DatabaseSession, tr opentracing.Tracer) *Geo {
	return &Geo{
		addr:      a,
		port:      p,
		dbsession: db,
		tracer:    tr,
		geoidx:    db.newGeoIndex(),
	}
}

// Run starts the server
func (s *Geo) Run() error {
	// TODO: Implement me
	// Run starts the server
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.addr, s.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	serverOptions := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
		}),
	}

	// Add OpenTracing interceptor for distributed tracing
	if s.tracer != nil {
		serverOptions = append(serverOptions, grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(s.tracer)))
	}

	grpcServer := grpc.NewServer(serverOptions...)
	pb.RegisterGeoServer(grpcServer, s)

	// Register reflection service on gRPC server to support reflection requests
	reflection.Register(grpcServer)

	log.Printf("Geo service listening on %s:%d", s.addr, s.port)
	return grpcServer.Serve(lis)

}

// Nearby returns all hotels within a given distance.
func (s *Geo) Nearby(ctx context.Context, req *pb.Request) (*pb.Result, error) {
	// TODO: Implement me
	// HINT: Reuse the implementation from the monolithic implementation 
	// HINT: and modify as needed.
	points := s.getNearbyPoints(float64(req.Lat), float64(req.Lon))
	res := &pb.Result{}

	for _, p := range points {
		res.HotelIds = append(res.HotelIds, p.Id())
	}

	return res, nil
}

func (s *Geo) getNearbyPoints(lat, lon float64) []geoindex.Point {
	center := &geoindex.GeoPoint{
		Pid:  "",
		Plat: lat,
		Plon: lon,
	}

	return s.geoidx.KNearest(
		center,
		maxSearchResults,
		geoindex.Km(maxSearchRadius), func(p geoindex.Point) bool {
			return true
		},
	)
}

// Implement Point interface
func (p *point) Lat() float64 { return p.Plat }
func (p *point) Lon() float64 { return p.Plon }
func (p *point) Id() string   { return p.Pid }
