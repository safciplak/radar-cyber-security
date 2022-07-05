package json

import "context"

type numbersApiClient interface {
	GetRandomText(ctx context.Context) (string, error)
}

type Service struct {
	Numbers numbersApiClient
}

func (s *Service) GetRandomText(ctx context.Context) (string, error) {
	return s.Numbers.GetRandomText(ctx)
}
