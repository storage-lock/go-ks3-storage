package ks3_storage

import (
	"context"
	"os"
	"testing"

	s3_storage "github.com/storage-lock/go-s3-storage"
	storage_test_helper "github.com/storage-lock/go-storage-test-helper"
	"github.com/stretchr/testify/assert"
)

func TestKS3Storage(t *testing.T) {
	endpoint := os.Getenv("STORAGE_LOCK_KS3_ENDPOINT")
	accessKey := os.Getenv("STORAGE_LOCK_KS3_ACCESS_KEY")
	secretKey := os.Getenv("STORAGE_LOCK_KS3_SECRET_KEY")
	bucket := os.Getenv("STORAGE_LOCK_KS3_BUCKET")

	if endpoint == "" || accessKey == "" || secretKey == "" || bucket == "" {
		t.Skip("KS3 credentials not set, skipping test")
	}

	ctx := context.Background()
	cm := s3_storage.NewS3ConnectionManagerFromEndpoint(endpoint, accessKey, secretKey, true)
	s, err := NewKS3Storage(ctx, &s3_storage.S3StorageOptions{
		ConnectionManager: cm,
		BucketName:        bucket,
		KeyPrefix:         storage_test_helper.TestTableName,
	})
	assert.Nil(t, err)
	storage_test_helper.TestStorage(t, s)
}
