package service

import (
	"CWS/categorization/models"
	"CWS/categorization/repository"
	"CWS/categorization/validators"
	"CWS/pb"
	"CWS/security"
	"context"
	"strconv"
	"time"

	"gopkg.in/mgo.v2"
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
	err := validators.ValidateUrl(req.Url)
	if err != nil {
		return nil, err
	}
	base64Url := security.Base64Encode(req.Url)
	found, err := s.categoriesRepository.GetCategoryByUrl(base64Url)
	if err != nil {
		return nil, err
	}
	found.Url = req.Url
	return found.ToProtoBuffer(), nil
}

// UpdateCategory performs update the category.
func (s *CatService) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.Category, error) {
	err := validators.ValidateUrl(req.Url)
	if err != nil {
		return nil, err
	}
	base64Url := security.Base64Encode(req.Url)
	found, err := s.categoriesRepository.GetCategoryByUrl(base64Url)
	if err != nil {
		return nil, err
	}
	if found.Category == req.Category {
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

// ReportMiscategorization reports miscategorization.
func (s *CatService) ReportMiscategorization(ctx context.Context, req *pb.GetCategoryRequest) (*pb.Category, error) {
	err := validators.ValidateUrl(req.Url)
	if err != nil {
		return nil, err
	}
	base64Url := security.Base64Encode(req.Url)
	found, err := s.categoriesRepository.GetCategoryByUrl(base64Url)
	if err != nil {
		return nil, err
	}
	//send to categorizer
	//use returned value to update category
	category := "NewCategory"
	if found.Category == category {
		return found.ToProtoBuffer(), nil
	}
	found.Category = category
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

// AddUrls performs add the urls.
func (s *CatService) AddUrls(req *pb.AddUrlsRequest, stream pb.CatService_AddUrlsServer) error {
	err := validators.ValidateUrls(req.Urls)
	if err != nil {
		return err
	}
	for _, url := range req.Urls {
		err := validators.ValidateUrl(url)
		if err != nil {
			return err
		}
		base64Url := security.Base64Encode(url)
		found, err := s.categoriesRepository.GetCategoryByUrl(base64Url)
		if err == mgo.ErrNotFound {
			//send to categorizer
			//use returned value to update category
			category := new(models.Category)
			category.Url = base64Url
			category.Category = "NewCategory"
			category.Created = time.Now()
			category.Updated = time.Now()
			category.Id = bson.NewObjectId()
			category.Revision = "0"
			err := s.categoriesRepository.Save(category)
			if err != nil {
				return err
			}
			category.Url = url
			err = stream.Send(category.ToProtoBuffer())
			if err != nil {
				return err
			}
		} else if found == nil {
			return err
		} else {
			return validators.ErrUrlAlreadyExist
		}
	}
	return nil
}

// AddUrl performs add the url.
func (s *CatService) AddUrl(ctx context.Context, req *pb.AddUrlRequest) (*pb.Category, error) {
	err := validators.ValidateUrl(req.Url)
	if err != nil {
		return nil, err
	}
	base64Url := security.Base64Encode(req.Url)
	found, err := s.categoriesRepository.GetCategoryByUrl(base64Url)
	if err == mgo.ErrNotFound {
		//send to categorizer
		//use returned value to update category
		category := new(models.Category)
		category.Url = base64Url
		category.Category = "NewCategory"
		category.Created = time.Now()
		category.Updated = time.Now()
		category.Id = bson.NewObjectId()
		category.Revision = "0"
		err := s.categoriesRepository.Save(category)
		if err != nil {
			return nil, err
		}
		category.Url = req.Url
		return category.ToProtoBuffer(), nil
	}
	if found == nil {
		return nil, err
	}
	return nil, validators.ErrUrlAlreadyExist
}

// DeleteUrls performs delete the urls.
func (s *CatService) DeleteUrls(req *pb.DeleteUrlsRequest, stream pb.CatService_DeleteUrlsServer) error {
	err := validators.ValidateUrls(req.Urls)
	if err != nil {
		return err
	}
	for _, url := range req.Urls {
		err := validators.ValidateUrl(url)
		if err != nil {
			return err
		}
		base64Url := security.Base64Encode(url)
		category, err := s.categoriesRepository.GetCategoryByUrl(base64Url)
		if err != nil {
			return err
		}
		err = s.categoriesRepository.Delete(category.Id.Hex())
		if err != nil {
			return err
		}
		category.Url = url
		err = stream.Send(category.ToProtoBuffer())
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteUrl performs delete the url.
func (s *CatService) DeleteUrl(ctx context.Context, req *pb.DeleteUrlRequest) (*pb.DeleteUrlResponse, error) {
	err := validators.ValidateUrl(req.Url)
	if err != nil {
		return nil, err
	}
	base64Url := security.Base64Encode(req.Url)
	category, err := s.categoriesRepository.GetCategoryByUrl(base64Url)
	if err != nil {
		return nil, err
	}
	err = s.categoriesRepository.Delete(category.Id.Hex())
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUrlResponse{Url: req.Url}, nil
}

// ListUrls performs list the urls based on categories and count.
func (s *CatService) ListUrls(req *pb.ListUrlsRequest, stream pb.CatService_ListUrlsServer) error {
	err := validators.ValidateCategories(req.Categories)
	if err != nil {
		return err
	}
	count, err := validators.ValidateCount(req.Count)
	if err != nil {
		return err
	}

	for _, category := range req.Categories {

		urls, err := s.categoriesRepository.GetAllUrlsByCategory(category, count)
		if err != nil {
			return err
		}
		for _, url := range urls {
			url.Url, err = security.Base64Decode(url.Url)
			if err != nil {
				return err
			}
			err := stream.Send(url.ToProtoBuffer())
			if err != nil {
				return err
			}
		}
	}

	return nil
}
