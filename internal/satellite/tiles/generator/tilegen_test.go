package generator_test

import (
	"context"
	"testing"

	"abc/mocks"

	"github.com/stretchr/testify/assert"

	"abc/internal/satellite/api/presenter"
	"abc/internal/satellite/tiles/generator"
	"abc/internal/satellite/usecase/mocks"
)

func TestFeaturesProcessor_CreateMapFeatureResponse(t *testing.T) {
	ctx := context.Background()
	url := "https://some.url.org/path"
	blobStorage := new(mocks.BlobStorage)
	imageDownloader := new(mocks.DownloadImages)
	fileDirectory := "/path/to/file"
	featureproc := generator.NewFeaturesProcessor(url, blobStorage, imageDownloader, fileDirectory)

	logger := new(mocks.LogMessage)
	compty := "astc"
	mfstile := presenter.MfsTiles{}
	b, err := featureproc.CreateMapFeatureResponse(ctx, &mfstile, compty, logger)
	assert.Error(t, b, err, "expected at error at creating feature response")
	assert.Nil(t, b, "expected nil map feature response")
}

// Above is the code written for func NewFeaturesProcessor and CreateMapFeatureResponse.
// where blobstorage and DownloadImages are interfaces
// above testcase passses and gives code coverage percentage
// i have followed above code pattern for most of code, please check if this is right way to write

// case :2  ;;; unsuccessful
//below code i tried to write testcase using the "mocks" interface code auto generated via "mockery" command "mockery --all --keeptree"

func TestNewFeaturesProcessor(t *testing.T) {
	ctx := context.Background()
	url := "https://some.url.org/path"
	blobStorage := new(mocks.BlobStorage)
	imageDownloader := new(mocks.DownloadImages)
	fileDirectory := "/path/to/file"
	feature := generator.NewFeaturesProcessor(url, blobStorage, imageDownloader, fileDirectory)

	// Create a dummy BlobDetails, in whichever package it is in
	blobDetails := BlobDetails{
		// ...
		// Assign fields in exported.
		// If not, call the constructor
	}

	logger := new(mocks.LogMessage)
	lastModefied := "12345678"
	mfTiles := []presenter.MfsTiles{
		{
			// ...
		},
	}
	compressionType := "something"

	blobStorage.On("DownloadBlob", ctx, fileDirectory, logger).Return(blobDetails, nil)
	blobStorage.On("GetLastModifiedFromBlob", ctx, fileDirectory, logger).Return(lastModefied, nil)
	blobStorage.On("UploadBlob", fileDirectory, blobDetails).Return()

	resp, err := feature.CreateMapFeatureResponse(ctx, mfTiles, compressionType, logger)
	assert.NoError(t, err, "unexpected error")
	assert.NotNil(t, resp, "Unexpected nil from map feature response")

}

// case :3
//Tried to implement test cases for channels.. is this the correct way
func TestCheckTileStatus(t *testing.T) {
	var ctx = context.Background()
	tmock := new(Interface_mock)
	var mapfet3 chan presenter.MapFeatures
	mapfet2 := &presenter.MfsTiles{}
	str := "astc"
	feature := &FeaturesProcessor{}
	var a chan error
	CheckTileStatus(a, mapfet3, mapfet2, ctx, str, tmock.LogMessage, feature)
}

// case :4 ;
// below code i tried to write testcase using the "mocks" one more try in different way
func TestCreateMapFeatureResponse(t *testing.T) {
	mockRepo := new(mocks.BlobStorage)
	mockRepo2 := new(mocks.DownloadImages)

	//downblob := &mocks.BlobStorage.DownloadBlob{ctx, str, mockRepo.LogMessage}

	mockRepo.On("DownloadBlob").Return(storageblob.BlobDetails{}, nil)
	mockRepo.On("UploadBlob").Return(storageblob.BlobDetails{}, nil)
	mockRepo.On("GetLastModifiedFromBlob").Return(storageblob.BlobDetails{}, nil)
	mockRepo.On("GetDownloadJobs").Return(&download.DJobs{}, nil)
	tst := NewFeaturesProcessor("test", mockRepo, mockRepo2, "test2")
	compty := "astc"
	mfstile := []*presenter.MfsTiles{}
	var ctx = context.Background()
	result, _ := tst.CreateMapFeatureResponse(ctx, mfstile, compty, mfsvc.NewLogMessage("test"))
	mockRepo.AssertExpectations(t)
	assert.NotNil(t, result, "Success!")
}
