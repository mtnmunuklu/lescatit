package service

import (
	"context"
	"time"

	"github.com/mtnmunuklu/lescatit/crawler/repository"
	"github.com/mtnmunuklu/lescatit/crawler/scraper"
	"github.com/mtnmunuklu/lescatit/crawler/util"
	"github.com/mtnmunuklu/lescatit/pb"
	"github.com/mtnmunuklu/lescatit/security"
)

// CrawlService provides crawlersRepository and collyScraper for crawler service.
type CrawlService struct {
	crawlersRepository repository.CrawlersRepository
	collyScraper       scraper.CollyScraper
}

// NewCrawlService creates a new CrawlService instance.
func NewCrawlService(crawlersRepository repository.CrawlersRepository, collyScraper scraper.CollyScraper) pb.CrawlServiceServer {
	return &CrawlService{crawlersRepository: crawlersRepository, collyScraper: collyScraper}
}

// GetURLData provides to get the content in the url address.
func (s *CrawlService) GetURLData(ctx context.Context, req *pb.GetURLDataRequest) (*pb.GetURLDataResponse, error) {
	err := util.ValidateURL(req.GetUrl())
	if err != nil {
		return nil, err
	}

	base64URL := security.Base64Encode(req.GetUrl())
	data, err := s.crawlersRepository.GetDataByURL(base64URL)
	if err != nil {
		if req.GetType() == "notnew" {
			return nil, util.ErrNotExistURL
		} else {
			currentData, err := s.collyScraper.GetData(req.GetUrl())
			if err != nil {
				return nil, util.ErrGetData
			}

			base64Data := security.Base64Encode(currentData)

			return &pb.GetURLDataResponse{Url: req.GetUrl(), Data: base64Data, Status: "New"}, nil
		}
	}

	currentTime := time.Now()
	diff := currentTime.Sub(data.Updated).Hours()
	if diff > 720 {
		currentData, err := s.collyScraper.GetData(req.GetUrl())
		if err != nil {
			return nil, util.ErrGetData
		}

		base64Data := security.Base64Encode(currentData)

		return &pb.GetURLDataResponse{Url: req.GetUrl(), Data: base64Data, Status: "Updated"}, nil
	}

	return &pb.GetURLDataResponse{Url: req.GetUrl(), Data: data.Data}, nil
}

// GetURLsData provides to get the content in the url addresses.
func (s *CrawlService) GetURLsData(req *pb.GetURLsDataRequest, stream pb.CrawlService_GetURLsDataServer) error {

	for _, url := range req.GetURLsDataRequest {
		err := util.ValidateURL(url.GetUrl())
		if err == nil {
			base64URL := security.Base64Encode(url.GetUrl())
			data, err := s.crawlersRepository.GetDataByURL(base64URL)
			if err != nil {
				if url.GetType() != "notnew" {
					currentData, err := s.collyScraper.GetData(url.GetUrl())
					if err == nil {
						base64Data := security.Base64Encode(currentData)
						err = stream.Send(&pb.GetURLDataResponse{Url: url.GetUrl(), Data: base64Data, Status: "New"})
						if err != nil {
							return err
						}
					}
				}
			} else if data != nil {
				currentTime := time.Now()
				diff := currentTime.Sub(data.Updated).Hours()
				if diff > 720 {
					currentData, err := s.collyScraper.GetData(url.GetUrl())
					if err == nil {
						base64Data := security.Base64Encode(currentData)
						err = stream.Send(&pb.GetURLDataResponse{Url: url.GetUrl(), Data: base64Data, Status: "Updated"})
						if err != nil {
							return util.ErrNotPerformedOperation
						}
					}
				}

				err = stream.Send(&pb.GetURLDataResponse{Url: url.GetUrl(), Data: data.Data})
				if err != nil {
					return util.ErrNotPerformedOperation
				}
			}
		}
	}

	return nil
}

// CrawlURL performs crawl the url.
func (s *CrawlService) CrawlURL(ctx context.Context, req *pb.CrawlURLRequest) (*pb.CrawlURLResponse, error) {
	err := util.ValidateURL(req.Url)
	if err != nil {
		return nil, err
	}

	s.collyScraper.FromProtoBuffer(req.GetCrawlRequest())

	links, err := s.collyScraper.GetLinks(req.GetUrl())
	if err != nil {
		return nil, util.ErrGetLinks
	}

	return &pb.CrawlURLResponse{Url: req.GetUrl(), Links: links}, nil
}

// CrawlUrls performs crawl the urls.
func (s *CrawlService) CrawlURLs(req *pb.CrawlURLsRequest, stream pb.CrawlService_CrawlURLsServer) error {
	err := util.ValidateURLs(req.GetUrls())
	if err != nil {
		return err
	}

	s.collyScraper.FromProtoBuffer(req.GetCrawlRequest())

	for _, url := range req.GetUrls() {
		links, err := s.collyScraper.GetLinks(url)
		if err == nil {
			err := stream.Send(&pb.CrawlURLResponse{Url: url, Links: links})
			if err != nil {
				return util.ErrNotPerformedOperation
			}
		}
	}

	return nil
}
