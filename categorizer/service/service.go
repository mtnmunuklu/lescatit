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
	return &CatzeService{categorizersRepository: categorizersRepository, tokenizer: tokenizer, nbClassifier: nbClassifier}
}

// CategorizeURL provides to categorize the url.
func (s *CatzeService) CategorizeURL(ctx context.Context, req *pb.CategorizeURLRequest) (*pb.CategorizeURLResponse, error) {
	err := util.ValidateURL(req.GetUrl())
	if err != nil {
		return nil, err
	}
	if len(req.GetData()) > 0 {
		base64DecodedData, err := security.Base64Decode(req.GetData())
		if err != nil {
			return nil, err
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
				return nil, util.ErrFailedModelGet
			}
			classifer.Data, err = security.Base64Decode(classifer.Data)
			if err != nil {
				return nil, err
			}
			err = s.nbClassifier.ReadClassifier(classifer.Data)
			if err != nil {
				return nil, util.ErrFailedModelRead
			}
			category := s.nbClassifier.Predict(tokens)
			return &pb.CategorizeURLResponse{Url: req.GetUrl(), Category: category}, nil
		}
	}
	return nil, util.ErrEmptyData
}

// CategorizeURLs provides to categorize the urls.
func (s *CatzeService) CategorizeURLs(req *pb.CategorizeURLsRequest, stream pb.CatzeService_CategorizeURLsServer) error {
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

func (s *CatzeService) GenerateClassificationModel(ctx context.Context, req *pb.GenerateClassificationModelRequest) (*pb.Classifier, error) {
	if req.GetCategory() == "NB" {
		model := make(map[string][]string)
		for _, cmodel := range req.GetModel() {
			base64DecodedData, err := security.Base64Decode(cmodel.GetData())
			if err != nil {
				return nil, err
			}
			bytesData := bytes.NewBuffer([]byte(base64DecodedData))
			tokenize := s.tokenizer.Tokenize(bytesData)
			for token := range tokenize {
				model[cmodel.GetClass()] = append(model[cmodel.GetClass()], token)
			}
		}
		data, err := s.nbClassifier.Learn(model)
		if err != nil {
			return nil, util.ErrFailedModelLearn
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
			return nil, util.ErrFailedModelSave
		}
		return classifier.ToProtoBuffer(), nil
	}
	return nil, util.ErrInvalidCategorizationModel
}
