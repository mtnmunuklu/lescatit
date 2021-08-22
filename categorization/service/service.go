package service

import (
	"Lescatit/categorization/models"
	"Lescatit/categorization/repository"
	"Lescatit/categorization/validators"
	"Lescatit/pb"
	"Lescatit/security"
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
	err := validators.ValidateURL(req.Url)
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
	err := validators.ValidateURL(req.Url)
	if err != nil {
		return nil, err
	}
	base64URL := security.Base64Encode(req.Url)
	found, err := s.categoriesRepository.GetCategoryByURL(base64URL)
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
	err := validators.ValidateURL(req.Url)
	if err != nil {
		return nil, err
	}
	base64URL := security.Base64Encode(req.Url)
	found, err := s.categoriesRepository.GetCategoryByURL(base64URL)
	if err != nil {
		return nil, err
	}
	//if update time is not current send url to crawler
	//get data of the url
	data := "Data"
	//send data of the url to categorizer
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
	found.Data = security.Base64Encode(data)
	err = s.categoriesRepository.Update(found)
	if err != nil {
		return nil, err
	}
	found.Url = req.Url
	return found.ToProtoBuffer(), nil
}

// AddURLs performs add the urls.
func (s *CatService) AddURLs(req *pb.AddURLsRequest, stream pb.CatService_AddURLsServer) error {
	err := validators.ValidateURLs(req.Urls)
	if err != nil {
		return err
	}
	for _, url := range req.Urls {
		err := validators.ValidateURL(url)
		if err == nil {
			base64URL := security.Base64Encode(url)
			_, err := s.categoriesRepository.GetCategoryByURL(base64URL)
			if err == mgo.ErrNotFound {
				category := new(models.Category)
				category.Url = base64URL
				//send url to crawler
				//get data of the url
				data := "Data"
				//send data of the url to categorizer
				//use returned value to update category
				category.Category = "Category"
				category.Created = time.Now()
				category.Updated = time.Now()
				category.Id = bson.NewObjectId()
				category.Revision = "0"
				category.Data = security.Base64Encode(data)
				err := s.categoriesRepository.Save(category)
				if err == nil {
					category.Url = url
					err = stream.Send(category.ToProtoBuffer())
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

// AddURL performs add the url.
func (s *CatService) AddURL(ctx context.Context, req *pb.AddURLRequest) (*pb.Category, error) {
	err := validators.ValidateURL(req.Url)
	if err != nil {
		return nil, err
	}
	base64URL := security.Base64Encode(req.Url)
	found, err := s.categoriesRepository.GetCategoryByURL(base64URL)
	if err == mgo.ErrNotFound {
		url := new(models.Category)
		url.Url = base64URL
		//send url to crawler
		//get data of the url
		data := "Data"
		//send data of the url to categorizer
		//use returned value to update category
		url.Category = "Category"
		url.Created = time.Now()
		url.Updated = time.Now()
		url.Id = bson.NewObjectId()
		url.Revision = "0"
		url.Data = security.Base64Encode(data)
		err := s.categoriesRepository.Save(url)
		if err != nil {
			return nil, err
		}
		url.Url = req.Url
		return url.ToProtoBuffer(), nil
	}
	if found == nil {
		return nil, err
	}
	return nil, validators.ErrURLAlreadyExist
}

// DeleteURLs performs delete the urls.
func (s *CatService) DeleteURLs(req *pb.DeleteURLsRequest, stream pb.CatService_DeleteURLsServer) error {
	err := validators.ValidateURLs(req.Urls)
	if err != nil {
		return err
	}
	for _, url := range req.Urls {
		err := validators.ValidateURL(url)
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

// DeleteURL performs delete the url.
func (s *CatService) DeleteURL(ctx context.Context, req *pb.DeleteURLRequest) (*pb.DeleteURLResponse, error) {
	err := validators.ValidateURL(req.Url)
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

// ListURLs performs list the urls based on categories and count.
func (s *CatService) ListURLs(req *pb.ListURLsRequest, stream pb.CatService_ListURLsServer) error {
	err := validators.ValidateCategories(req.Categories)
	if err != nil {
		return err
	}
	count, err := validators.ValidateCount(req.Count)
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
