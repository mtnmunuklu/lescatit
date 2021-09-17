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
	err := validators.ValidateURL(req.GetUrl())
	if err != nil {
		return nil, err
	}
	if len(req.Data) > 0 {
		return &pb.CategorizeURLResponse{Url: req.GetUrl(), Category: "uncategorized"}, nil
	}
	return nil, validators.ErrEmptyData
}

// CategorizeURLs provides to categorize the urls.
func (s *CatzeService) CategorizeURLs(req *pb.CategorizeURLsRequest, stream pb.CatzeService_CategorizeURLsServer) error {
	for _, url := range req.CategorizeURLsRequest {
		err := validators.ValidateURL(url.GetUrl())
		if err == nil {
			if len(url.GetData()) > 0 {
				err := stream.Send(&pb.CategorizeURLResponse{Url: url.GetUrl(), Category: "uncategorized"})
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
