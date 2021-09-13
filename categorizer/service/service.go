package service

import (
	"Lescatit/categorizer/repository"
	"Lescatit/categorizer/validators"
	"Lescatit/pb"
	"context"
)

// CatzeService provides categorizersRepository for categorizer service.
type CatzeService struct {
	categorizersRepository repository.CategorizersRepository
}

// NewCatzeService creates a new CatzeService instance.
func NewCatzeService(categorizersRepository repository.CategorizersRepository) pb.CatzeServiceServer {
	return &CatzeService{categorizersRepository: categorizersRepository}
}

// CategorizeURL provides to categorize the url.
func (s *CatzeService) CategorizeURL(ctx context.Context, req *pb.CategorizeURLRequest) (*pb.CategorizeURLResponse, error) {
	if len(req.Data) > 0 {
		return &pb.CategorizeURLResponse{Url: req.Url, Category: "uncategorized"}, nil
	}
	return nil, validators.ErrEmptyData
}

// CategorizeURLs provides to categorize the urls.
func (s *CatzeService) CategorizeURLs(req *pb.CategorizeURLsRequest, stream pb.CatzeService_CategorizeURLsServer) error {
	for _, url := range req.CategorizeURLRequest {
		if len(url.Data) > 0 {
			err := stream.Send(&pb.CategorizeURLResponse{Url: url.Url, Category: "uncategorized"})
			if err != nil {
				return err
			}
		}
	}
	return nil
}
