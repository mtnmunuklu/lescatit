package service

import (
	"Lescatit/categorizer/classifiers"
	"Lescatit/categorizer/models/classifiersmdl"
	"Lescatit/categorizer/repositories/categorizersrps"
	"Lescatit/categorizer/repositories/classifiersrps"
	"Lescatit/categorizer/tokenizer"
	"Lescatit/categorizer/util"
	"Lescatit/pb"
	"Lescatit/security"
	"bytes"
	"context"
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// CatzeService provides categorizersRepository for categorizer service.
type CatzeService struct {
	categorizersRepository categorizersrps.CategorizersRepository
	classifiersRepository  classifiersrps.ClassifiersRepository
	tokenizer              tokenizer.Tokenizer
	nbClassifier           classifiers.NaiveBayesianClassifier
}

// NewCatzeService creates a new CatzeService instance.
func NewCatzeService(categorizersRepository categorizersrps.CategorizersRepository, classifiersRepository classifiersrps.ClassifiersRepository, tokenizer tokenizer.Tokenizer, nbClassifier classifiers.NaiveBayesianClassifier) pb.CatzeServiceServer {
	return &CatzeService{categorizersRepository: categorizersRepository, classifiersRepository: classifiersRepository, tokenizer: tokenizer, nbClassifier: nbClassifier}
}

// CategorizeURL performs categorize the url.
func (s *CatzeService) CategorizeURL(ctx context.Context, req *pb.CategorizeURLRequest) (*pb.CategorizeURLResponse, error) {
	err := util.ValidateURL(req.GetUrl())
	if err != nil {
		return nil, err
	}

	err = util.ValidateName(req.GetCmodel())
	if err != nil {
		return nil, err
	}

	err = util.ValidateData(req.GetData())
	if err != nil {
		return nil, err
	}

	base64DecodedData, err := security.Base64Decode(req.GetData())
	if err != nil {
		return nil, util.ErrDecodeData
	}

	if strings.HasSuffix(req.GetCmodel(), ".nbc") {
		bytesData := bytes.NewBuffer([]byte(base64DecodedData))
		tokenize := s.tokenizer.Tokenize(bytesData)
		tokens := make([]string, 0)
		for token := range tokenize {
			tokens = append(tokens, token)
		}

		classifer, err := s.classifiersRepository.GetByName(req.GetCmodel())
		if err != nil {
			return nil, util.ErrGetModel
		}

		classifer.Data, err = security.Base64Decode(classifer.Data)
		if err != nil {
			return nil, util.ErrDecodeData
		}

		err = s.nbClassifier.ReadClassifier(classifer.Data)
		if err != nil {
			return nil, util.ErrReadModel
		}

		category := s.nbClassifier.Predict(tokens)

		return &pb.CategorizeURLResponse{Url: req.GetUrl(), Category: category}, nil
	}

	return nil, util.ErrInvalidCategorizationModel
}

// CategorizeURLs performs categorize the urls.
func (s *CatzeService) CategorizeURLs(req *pb.CategorizeURLsRequest, stream pb.CatzeService_CategorizeURLsServer) error {
	err := util.ValidateURLs(req.GetUrls())
	if err != nil {
		return err
	}

	for _, url := range req.GetUrls() {
		err := util.ValidateURL(url.GetUrl())
		if err == nil {
			if len(url.GetData()) > 0 {
				base64DecodedData, err := security.Base64Decode(url.GetData())
				if err == nil {
					if strings.HasSuffix(url.GetCmodel(), ".nbc") {
						bytesData := bytes.NewBuffer([]byte(base64DecodedData))
						tokenize := s.tokenizer.Tokenize(bytesData)
						tokens := make([]string, 0)
						for token := range tokenize {
							tokens = append(tokens, token)
						}
						classifer, err := s.classifiersRepository.GetByName(url.GetCmodel())
						if err == nil {
							classifer.Data, err = security.Base64Decode(classifer.Data)
							if err == nil {
								err = s.nbClassifier.ReadClassifier(classifer.Data)
								if err == nil {
									category := s.nbClassifier.Predict(tokens)
									err = stream.Send(&pb.CategorizeURLResponse{Url: url.GetUrl(), Category: category})
									if err != nil {
										return err
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return nil
}

// GenerateClassificationModel performs generate a classification model.
func (s *CatzeService) GenerateClassificationModel(ctx context.Context, req *pb.GenerateClassificationModelRequest) (*pb.Classifier, error) {
	if req.GetCategory() == "NB" {
		model := make(map[string][]string)
		for _, cmodel := range req.GetModel() {
			base64DecodedData, err := security.Base64Decode(cmodel.GetData())
			if err != nil {
				return nil, util.ErrDecodeData
			}

			bytesData := bytes.NewBuffer([]byte(base64DecodedData))
			tokenize := s.tokenizer.Tokenize(bytesData)
			for token := range tokenize {
				model[cmodel.GetClass()] = append(model[cmodel.GetClass()], token)
			}
		}

		data, err := s.nbClassifier.Learn(model)
		if err != nil {
			return nil, util.ErrLearnModel
		}

		classifier := new(classifiersmdl.Classifier)
		classifier.Id = bson.NewObjectId()
		classifier.Name = util.GenerateRandomFileName("", ".nbc")
		classifier.Category = "NB"
		classifier.Created = time.Now()
		classifier.Updated = time.Now()
		classifier.Revision = "0"
		classifier.Data = security.Base64Encode(data)

		err = s.classifiersRepository.Save(classifier)
		if err != nil {
			return nil, util.ErrSaveModel
		}

		return classifier.ToProtoBuffer(), nil
	}

	return nil, util.ErrInvalidCategorizationModel
}

// GetClassificationModel performs return the classification model.
func (s *CatzeService) GetClassificationModel(ctx context.Context, req *pb.GetClassificationModelRequest) (*pb.Classifier, error) {
	err := util.ValidateName(req.GetName())
	if err != nil {
		return nil, err
	}

	classifier, err := s.classifiersRepository.GetByName(req.GetName())
	if err != nil {
		return nil, util.ErrGetModel
	}

	return classifier.ToProtoBuffer(), nil
}

// UpdateClassificationModel performs update the classification model.
func (s *CatzeService) UpdateClassificationModel(ctx context.Context, req *pb.UpdateClassificationModelRequest) (*pb.Classifier, error) {
	err := util.ValidateName(req.GetName())
	if err != nil {
		return nil, err
	}

	err = util.ValidateCategory(req.GetCategory())
	if err != nil {
		return nil, err
	}

	classifier, err := s.classifiersRepository.GetByName(req.GetName())
	if err != nil {
		return nil, util.ErrGetModel
	}

	if classifier.Category == req.GetCategory() {
		return classifier.ToProtoBuffer(), nil
	}

	classifier.Category = req.GetCategory()
	classifier.Updated = time.Now()
	revision, err := strconv.Atoi(classifier.Revision)
	if err == nil {
		classifier.Revision = strconv.Itoa((revision + 1))
	} else {
		classifier.Revision = "0"
	}

	err = s.classifiersRepository.Update(classifier)
	if err != nil {
		return nil, util.ErrUpdateModel
	}

	return classifier.ToProtoBuffer(), nil
}

// DeleteClassificationModel performs delete the classification model.
func (s *CatzeService) DeleteClassificationModel(ctx context.Context, req *pb.DeleteClassificationModelRequest) (*pb.DeleteClassificationModelResponse, error) {
	err := util.ValidateName(req.GetName())
	if err != nil {
		return nil, err
	}

	classifier, err := s.classifiersRepository.GetByName(req.GetName())
	if err != nil {
		return nil, util.ErrGetModel
	}

	err = s.classifiersRepository.DeleteById(classifier.Id.Hex())
	if err != nil {
		return nil, util.ErrDeleteModel
	}

	return &pb.DeleteClassificationModelResponse{Name: classifier.Name}, nil
}

// DeleteClassificationModels performs delete the classification models.
func (s *CatzeService) DeleteClassificationModels(req *pb.DeleteClassificationModelsRequest, stream pb.CatzeService_DeleteClassificationModelsServer) error {
	err := util.ValidateNames(req.GetNames())
	if err != nil {
		return err
	}

	for _, name := range req.GetNames() {
		classifier, err := s.classifiersRepository.GetByName(name)
		if err == nil {
			err = s.classifiersRepository.DeleteById(classifier.Id.Hex())
			if err == nil {
				err = stream.Send(&pb.DeleteClassificationModelResponse{Name: name})
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// ListClassificationModels performs list all classification models.
func (s *CatzeService) ListClassificationModels(req *pb.ListClassificationModelsRequest, stream pb.CatzeService_ListClassificationModelsServer) error {
	err := util.ValidateCategories(req.GetCategories())
	if err != nil {
		return err
	}

	count, err := util.ValidateCount(req.GetCount())
	if err != nil {
		return err
	}

	for _, category := range req.GetCategories() {
		classifiers, err := s.classifiersRepository.GetAllClassifiersByCategory(category, count)
		if err == nil {
			for _, classifier := range classifiers {
				err = stream.Send(classifier.ToProtoBuffer())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
