package rpc

import (
	"context"
	"errors"
	"hex/internal/adapters/inbound/grpc/pb"
	"hex/models/dto"

	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/types/known/durationpb"
)

// GetAddition gets the result of adding parameters a and b
func (grpca *Adapter) ShortenUrl(ctx context.Context, req *pb.ShortnerRequest) (*pb.ShortnerResponse, error) {
	pr, ok := peer.FromContext(ctx)
	if ok {
		resp, err := grpca.api.GetUrlShortingService().CreateShorty(ctx, &dto.CustomShortRequest{
			URL:         req.Url,
			CustomShort: req.CustomShort,
			Expiry:      req.Expiry.AsDuration(),
		}, pr.Addr.String())
		if err != nil {
			return nil, err
		}
		return &pb.ShortnerResponse{
			Url:             resp.URL,
			CustomShort:     resp.CustomShort,
			Expiry:          durationpb.New(resp.Expiry),
			XRateRemaining:  resp.XRateRemaining,
			XRateLimitReset: durationpb.New(resp.XRateLimitReset),
		}, nil
	}
	return nil, errors.New("failed to get client IP")

}

// GetSubtraction gets the result of subtracting parameters a and b
func (grpca *Adapter) ResolveUrl(ctx context.Context, req *pb.UrlRequest) (*pb.UrlResponse, error) {
	url, err := grpca.api.GetUrlShortingService().ResolveShorty(ctx, req.ShortUrl)
	if err != nil {
		return nil, err
	}

	// http.Redirect(httptest.NewRecorder(), &http.Request{}, url, 301)

	return &pb.UrlResponse{ActualUrl: url}, nil
}
