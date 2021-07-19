package service

import (
	"CWS/categorization/models"
	"CWS/categorization/repository"
	"CWS/categorization/validators"
	"CWS/pb"
	"context"
	"time"

	"gopkg.in/mgo.v2"
)

type categoryService struct {
	categoriesRepository repository.CategoriesRepository
}

func NewCategorySevice(categoriesRepository repository.CategoriesRepository) pb.CategoryServiceServer {
	return &categoryService{categoriesRepository: categoriesRepository}
}

func (s *categoryService) GetCategory(ctx context.Context, req *pb.GetCategoryRequest) (*pb.Category, error) {
	err := validators.ValidateUrl(req.Url)
	if err != nil {
		return nil, err
	}
	found, err := s.categoriesRepository.GetCategoryByUrl(req.Url)
	if err != nil {
		return nil, err
	}
	return found.ToProtoBuffer(), nil
}

func (s *categoryService) ReportMiscategorization(ctx context.Context, req *pb.GetCategoryRequest) (*pb.Category, error) {
	err := validators.ValidateUrl(req.Url)
	if err != nil {
		return nil, err
	}
	found, err := s.categoriesRepository.GetCategoryByUrl(req.Url)
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
	found.Revision += 1
	err = s.categoriesRepository.Update(found)
	return found.ToProtoBuffer(), err
}

func (s *categoryService) ListUrls(req *pb.ListUrlsRequest, stream pb.CategoryService_ListUrlsServer) error {
	err := validators.ValidateCategory(req.Category)
	if err != nil {
		return err
	}
	count, err := validators.ValidateCount(req.Count)
	if err != nil {
		return err
	}

	for _, category := range req.Category {

		urls, err := s.categoriesRepository.GetAllUrlsByCategory(category, count)
		if err != nil {
			return err
		}
		for _, url := range urls {
			err := stream.Send(url.ToProtoBuffer())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *categoryService) UpdateCategory(ctx context.Context, req *pb.Category) (*pb.Category, error) {
	err := validators.ValidateUrl(req.Url)
	if err != nil {
		return nil, err
	}
	found, err := s.categoriesRepository.GetCategoryByUrl(req.Url)
	if err != nil {
		return nil, err
	}
	if found.Category == req.Category {
		return found.ToProtoBuffer(), nil
	}
	found.Category = req.Category
	found.Updated = time.Now()
	found.Revision += 1
	err = s.categoriesRepository.Update(found)
	return found.ToProtoBuffer(), err
}

func (s *categoryService) AddUrl(ctx context.Context, req *pb.Category) (*pb.Category, error) {
	err := validators.ValidateUrl(req.Url)
	if err != nil {
		return nil, err
	}
	err = validators.ValidateId(req.Id)
	if err != nil {
		return nil, err
	}
	found, err := s.categoriesRepository.GetCategoryByUrl(req.Url)
	if err == mgo.ErrNotFound {
		//send to categorizer
		//use returned value to update category
		category := "NewCategory"
		req.Category = category
		url := new(models.Category)
		url.FromProtoBuffer(req)
		err := s.categoriesRepository.Save(url)
		if err != nil {
			return nil, err
		}
		return url.ToProtoBuffer(), nil
	}
	if found == nil {
		return nil, err
	}
	return nil, validators.ErrUrlAlreadyExist
}

func (s *categoryService) DeleteUrl(ctx context.Context, req *pb.GetCategoryRequest) (*pb.DeleteUrlResponse, error) {
	err := validators.ValidateUrl(req.Url)
	if err != nil {
		return nil, err
	}
	err = s.categoriesRepository.Delete(req.Url)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUrlResponse{Url: req.Url}, nil
}
