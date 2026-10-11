package ks3_storage

// KS3 Storage 金山云对象存储
// 金山云 KS3 兼容 S3 协议，复用 S3 实现

import (
	"context"

	s3_storage "github.com/storage-lock/go-s3-storage"
	"github.com/storage-lock/go-storage"
	)

// KS3Storage 金山云 KS3 对象存储实现
type KS3Storage struct {
	*s3_storage.S3Storage
}

var _ storage.Storage = &KS3Storage{}

func NewKS3Storage(ctx context.Context, options *s3_storage.S3StorageOptions) (*KS3Storage, error) {
	s3Storage, err := newS3(ctx, options)
	if err != nil {
		return nil, err
	}
	return &KS3Storage{S3Storage: s3Storage}, nil
}

var newS3 = s3_storage.NewS3Storage

const KS3StorageName = "ks3-storage"

