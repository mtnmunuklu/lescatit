package service

import (
	"Lescatit/categorization/models"
	"Lescatit/categorization/repository"
	"Lescatit/categorization/util"
	"Lescatit/pb"
	"Lescatit/security"
	"context"
	"strconv"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// CatService provides categoriesRepository for categorization service.
type CatService struct {
	categoriesRepository repository.CategoriesRepository
}

// NewCatService creates a new CatService instance.
func NewCatSevice(categoriesRepository repository.CategoriesRepository) pb.CatServiceServer {
	return &CatService{categoriesRepository: categoriesRepository}
}

// GetCategory performs return the category by url.
func (s *CatService) GetCategory(ctx context.Context, req *pb.GetCategoryRequest) (*pb.Category, error) {
	err := util.ValidateURL(req.Url)
	if err != nil {
		return nil, err
	}
	base64URL := security.Base64Encode(req.Url)
	found, err := s.categoriesRepository.GetCategoryByURL(base64URL)
	if err != nil {
		return nil, err
	}
	found.Url = req.Url
	return found.ToProtoBuffer(), nil
}

// UpdateCategory performs update the category.
func (s *CatService) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.Category, error) {
	err := util.ValidateURL(req.Url)
	if err != nil {
		return nil, err
	}
	base64URL := security.Base64Encode(req.Url)
	found, err := s.categoriesRepository.GetCategoryByURL(base64URL)
	if err != nil {
		return nil, err
	}
	if found.Category == req.Category {
		found.Url, err = security.Base64Decode(found.Url)
		if err != nil {
			return nil, err
		}
		return found.ToProtoBuffer(), nil
	}
	found.Category = req.Category
	found.Updated = time.Now()
	revision, err := strconv.Atoi(found.Revision)
	if err == nil {
		found.Revision = strconv.Itoa((revision + 1))
	} else {
		found.Revision = "0"
	}
	err = s.categoriesRepository.Update(found)
	if err != nil {
		return nil, err
	}
	found.Url = req.Url
	return found.ToProtoBuffer(), nil
}

// AddURL performs add the url.
func (s *CatService) AddURL(ctx context.Context, req *pb.AddURLRequest) (*pb.Category, error) {

	base64URL := security.Base64Encode(req.GetUrl())
	// New url
	if req.GetStatus() == "New" {
		addedURL := new(models.Category)
		addedURL.Url = base64URL
		addedURL.Category = req.GetCategory()
		addedURL.Created = time.Now()
		addedURL.Updated = time.Now()
		addedURL.Id = bson.NewObjectId()
		addedURL.Revision = "0"
		addedURL.Data = req.GetData()
		err := s.categoriesRepository.Save(addedURL)
		if err != nil {
			return nil, err
		}
		addedURL.Url = req.GetUrl()
		return addedURL.ToProtoBuffer(), nil
	}
	// Updated url
	updatedURL, err := s.categoriesRepository.GetCategoryByURL(base64URL)
	if err != nil {
		return nil, err
	}
	updatedURL.Category = req.GetCategory()
	updatedURL.Updated = time.Now()
	revision, err := strconv.Atoi(updatedURL.Revision)
	if err == nil {
		updatedURL.Revision = strconv.Itoa((revision + 1))
	} else {
		updatedURL.Revision = "0"
	}
	updatedURL.Data = req.GetData()
	err = s.categoriesRepository.Update(updatedURL)
	if err != nil {
		return nil, err
	}
	updatedURL.Url = req.GetUrl()
	return updatedURL.ToProtoBuffer(), nil
}

// DeleteURL performs delete the url.
func (s *CatService) DeleteURL(ctx context.Context, req *pb.DeleteURLRequest) (*pb.DeleteURLResponse, error) {
	err := util.ValidateURL(req.Url)
	if err != nil {
		return nil, err
	}
	base64URL := security.Base64Encode(req.Url)
	found, err := s.categoriesRepository.GetCategoryByURL(base64URL)
	if err != nil {
		return nil, err
	}
	err = s.categoriesRepository.Delete(found.Id.Hex())
	if err != nil {
		return nil, err
	}
	return &pb.DeleteURLResponse{Url: req.Url}, nil
}

// ReportMiscategorization reports miscategorization.
func (s *CatService) ReportMiscategorization(ctx context.Context, req *pb.ReportMiscategorizationRequest) (*pb.Category, error) {
	base64URL := security.Base64Encode(req.GetUrl())
	reportedURL, err := s.categoriesRepository.GetCategoryByURL(base64URL)
	if err != nil {
		return nil, err
	}
	if reportedURL.Category == req.GetCategory() {
		reportedURL.Url = req.Url
		return reportedURL.ToProtoBuffer(), nil
	}
	reportedURL.Category = req.GetCategory()
	reportedURL.Updated = time.Now()
	revision, err := strconv.Atoi(reportedURL.Revision)
	if err == nil {
		reportedURL.Revision = strconv.Itoa((revision + 1))
	} else {
		reportedURL.Revision = "0"
	}
	reportedURL.Data = req.GetData()
	err = s.categoriesRepository.Update(reportedURL)
	if err != nil {
		return nil, err
	}
	reportedURL.Url = req.Url
	return reportedURL.ToProtoBuffer(), nil
}

// DeleteURLs performs delete the urls.
func (s *CatService) DeleteURLs(req *pb.DeleteURLsRequest, stream pb.CatService_DeleteURLsServer) error {
	err := util.ValidateURLs(req.Urls)
	if err != nil {
		return err
	}
	for _, url := range req.Urls {
		err := util.ValidateURL(url)
		if err == nil {
			base64URL := security.Base64Encode(url)
			found, err := s.categoriesRepository.GetCategoryByURL(base64URL)
			if err == nil {
				err = s.categoriesRepository.Delete(found.Id.Hex())
				if err == nil {
					err = stream.Send(&pb.DeleteURLResponse{Url: url})
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

// ListURLs performs list the urls based on categories and count.
func (s *CatService) ListURLs(req *pb.ListURLsRequest, stream pb.CatService_ListURLsServer) error {
	err := util.ValidateCategories(req.Categories)
	if err != nil {
		return err
	}
	count, err := util.ValidateCount(req.Count)
	if err != nil {
		return err
	}
	for _, category := range req.Categories {
		urls, err := s.categoriesRepository.GetAllURLsByCategory(category, count)
		if err == nil {
			for _, url := range urls {
				url.Url, err = security.Base64Decode(url.Url)
				if err == nil {
					err := stream.Send(url.ToProtoBuffer())
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}
