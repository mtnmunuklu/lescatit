package service

import (
	"Lescatit/crawler/repository"
	"Lescatit/crawler/scraper"
	"Lescatit/crawler/validators"
	"Lescatit/pb"
	"Lescatit/security"
	"context"
	"time"
)

// CrawlService provides crawlersRepository for crawler service.
type CrawlService struct {
	crawlersRepository repository.CrawlersRepository
	collyScraper       scraper.CollyScraper
}

// NewCrawlService creates a new CrawlService instance.
func NewCrawlService(crawlersRepository repository.CrawlersRepository, collyScraper scraper.CollyScraper) pb.CrawlServiceServer {
	return &CrawlService{crawlersRepository: crawlersRepository, collyScraper: collyScraper}
}

//GetURLData provides to get the content in the url address.
func (s *CrawlService) GetURLData(ctx context.Context, req *pb.GetURLDataRequest) (*pb.GetURLDataResponse, error) {
	err := validators.ValidateURL(req.Url)
	if err != nil {
		return nil, err
	}
	base64URL := security.Base64Encode(req.Url)
	data, err := s.crawlersRepository.GetDataByURL(base64URL)
	if err != nil {
		currentData, err := s.collyScraper.GetData(req.Url)
		if err != nil {
			return nil, err
		}
		base64Data := security.Base64Encode(currentData)
		return &pb.GetURLDataResponse{Url: req.Url, Data: base64Data}, nil
	}
	currentTime := time.Now()
	diff := currentTime.Sub(data.Updated).Hours()
	if diff > 720 {
		currentData, err := s.collyScraper.GetData(req.Url)
		if err != nil {
			return nil, err
		}
		data.Data = security.Base64Encode(currentData)
	}
	return &pb.GetURLDataResponse{Url: req.Url, Data: data.Data}, nil
}

//GetURLsData provides to get the content in the url addresses.
func (s *CrawlService) GetURLsData(req *pb.GetURLsDataRequest, stream pb.CrawlService_GetURLsDataServer) error {
	err := validators.ValidateURLs(req.Urls)
	if err != nil {
		return err
	}
	for _, url := range req.Urls {
		err := validators.ValidateURL(url)
		if err == nil {
			base64URL := security.Base64Encode(url)
			data, err := s.crawlersRepository.GetDataByURL(base64URL)
			if err != nil {
				currentData, err := s.collyScraper.GetData(url)
				if err == nil {
					base64Data := security.Base64Encode(currentData)
					err = stream.Send(&pb.GetURLDataResponse{Url: url, Data: base64Data})
					if err != nil {
						return err
					}
				}
			} else {
				currentTime := time.Now()
				diff := currentTime.Sub(data.Updated).Hours()
				if diff > 720 {
					currentData, err := s.collyScraper.GetData(url)
					if err == nil {
						data.Data = security.Base64Encode(currentData)
					}
				}
				err = stream.Send(&pb.GetURLDataResponse{Url: url, Data: data.Data})
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// CrawlURL performs crawl the url
func (s *CrawlService) CrawlURL(ctx context.Context, req *pb.CrawlURLRequest) (*pb.CrawlURLResponse, error) {
	err := validators.ValidateURL(req.Url)
	if err != nil {
		return nil, err
	}
	s.collyScraper.FromProtoBuffer(req.GetCrawlRequest())
	links, err := s.collyScraper.GetLinks(req.Url)
	if err != nil {
		return nil, err
	}
	return &pb.CrawlURLResponse{Url: req.Url, Links: links}, nil
}

// CrawlUrls performs crawl the urls
func (s *CrawlService) CrawlURLs(req *pb.CrawlURLsRequest, stream pb.CrawlService_CrawlURLsServer) error {
	err := validators.ValidateURLs(req.Urls)
	if err != nil {
		return err
	}
	s.collyScraper.FromProtoBuffer(req.GetCrawlRequest())
	for _, url := range req.Urls {
		links, err := s.collyScraper.GetLinks(url)
		if err == nil {
			err := stream.Send(&pb.CrawlURLResponse{Url: url, Links: links})
			if err != nil {
				return err
			}
		}
	}
	return nil
}
