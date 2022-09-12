package generator_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"abc/internal/satellite/api/presenter"
	"abc/internal/satellite/tiles/generator"
	"abc/internal/satellite/usecase"
)

type Interface_mock struct {
	usecase.BlobStorage
	usecase.DownloadImages
	mfsvc.LogMessage
	error
	mock.Mock
}

// case :1  ;;; successful
func TestNewFeaturesProcessor(t *testing.T) {

	bu := "test"
	fDir := "test"
	tmock := new(Interface_mock)
	a := generator.NewFeaturesProcessor(bu, tmock.BlobStorage, tmock.DownloadImages, fDir)
	assert.NotNil(t, a, "Success!")

}

func TestCreateMapFeatureResponse(t *testing.T) {
	var ctx = context.Background()
	feature := generator.FeaturesProcessor{}
	bmock := new(Interface_mock)
	cy := "type"
	cz := []*presenter.MfsTiles{}
	b, err := feature.CreateMapFeatureResponse(ctx, cz, cy, bmock.LogMessage)
	assert.NotNil(t, b, err, "Success!")
}

// Above is the code written for func NewFeaturesProcessor and CreateMapFeatureResponse.
// where blobstorage and DownloadImages are interfaces
// above testcase passses and gives code coverage percentage
// i have followed above code pattern for most of code, please check if this is right way to write






// case :2  ;;; unsuccessful
//below code i tried to write testcase using the "mocks" interface code auto generated via "mockery" command "mockery --all --keeptree"

func TestNewFeaturesProcessor(t *testing.T) {

	bu := "test"
	fDir := "test"
	tmock := new(Interface_mock)

	feature := generator.FeaturesProcessor{}

	s1 := &generator.FeaturesProcessor{
		BaseURL: "test",


		// //tried with postman response 
		// //but was not successful
		// Storage: (postman response)
		// DImages: (postman response)
		
		Storage: nil,
		DImages: nil,

		FileDir: "test",
	}
	// how to use tmock.ON function ?
	tmock.On("DownloadBlob", mock.Anything).Return(true) 
	// here the control dint go into the "DownloadBlob" func which is in the main go file.
	// so it test case got passed but coverage was zero percentage.. more over no point in writing test cases if control dosent flow to .go file
	tmock.On("GetLastModifiedFromBlob", mock.Anything).Return(true)
	tmock.On("UploadBlob", mock.Anything).Return(true)
	tmock.On("NewFeaturesProcessor", bu, fDir).Return(s1)

	
	
	// there might be miss in code here, 
	a := generator.NewFeaturesProcessor(bu, tmock.BlobStorage, tmock.DownloadImages, fDir)
	feature.CreateMapFeatureResponse()
	feature := NewFeaturesProcessor(bu, tmock.BlobStorage,tmock.DownloadImages , fDir)
	assert.Nil(t, s1)
	assert.NotNil(t, a, "Success!")

}



// case :3  ;;; unsuccessful
//Tried to implement test cases for channels, but getting error with the below code


func TestCheckTileStatus(t *testing.T) {
	var ctx = context.Background()
	tmock := new(Interface_mock)
	mapfet := presenter.MapFeatures{}
	mapfet2 := &presenter.MfsTiles{}
	str := "test"
	feature := &generator.FeaturesProcessor{}
	generator.CheckTileStatus(tmock.error, mapfet, mapfet2, ctx, str, tmock.LogMessage, feature)
	assert.NotNil(t, s, "Success!")

}

// case :4  ;;; unsuccessful

// not able to import these funtions from .go file to test case file... since it starts from small letter alphabet, How to get coverage for such codes.
func checkRefreshRequired(ctx context.Context, requestData *models.ResourceRequest, fp *FeaturesProcessor, logMsg mfsvc.LogMessage) error {
}
return nil
}

func (sa *StorageAccount) getBlobUrlIndex(filename string) int {
}
