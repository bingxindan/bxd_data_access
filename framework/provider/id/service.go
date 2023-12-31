package id

import "github.com/rs/xid"

type BxdIDService struct {
}

func NewBxdIDService(params ...interface{}) (interface{}, error) {
	return &BxdIDService{}, nil
}

func (s *BxdIDService) NewID() string {
	return xid.New().String()
}
