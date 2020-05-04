package service

import "github.com/knry0329/go-di/repository"

type FileService interface {
	GetFileName(bucket string) (string, error)
}

type fileService struct {
	repo repository.FileRepository
}

func NewFileService(repo repository.FileRepository) FileService {
	return &fileService{
		repo: repo,
	}
}

func (s *fileService) GetFileName(bucket string) (string, error) {
	err := s.repo.GetLs(bucket)
	if err != nil {
		return "", err
	}
	return "okok", nil
}
