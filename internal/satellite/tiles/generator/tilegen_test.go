package generator

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
	url := "http://nds_til_dynamic/astc"
	fDir := "filedir"
	tmock := new(Interface_mock)
	a := NewFeaturesProcessor(url, tmock.BlobStorage, tmock.DownloadImages, fDir)
	assert.NotNil(t, a, "Success!")

}

func TestCreateMapFeatureResponse(t *testing.T) {
	var ctx = context.Background()
	featureproc := FeaturesProcessor{}
	bmock := new(Interface_mock)
	compty := "astc"
	mfstile := []*presenter.MfsTiles{}
	b, err := featureproc.CreateMapFeatureResponse(ctx, mfstile, compty, bmock.LogMessage)
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
	feature := NewFeaturesProcessor(bu, tmock.BlobStorage, tmock.DownloadImages, fDir)
	assert.Nil(t, s1)
	assert.NotNil(t, a, "Success!")

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

//error received while running  case 4 code
PS C:\Users\..\tiles\generator> go test -cover
--- FAIL: TestCreateMapFeatureResponse (0.00s)
    tilegen_test.go:243: FAIL:  DownloadBlob()
                        at: [C:\Users\..\tiles\generator\tilegen_test.go:233]
    tilegen_test.go:243: FAIL:  UploadBlob()
                        at: [C:\Users\..\tiles\generator\tilegen_test.go:234]
    tilegen_test.go:243: FAIL:  GetLastModifiedFromBlob()
                        at: [C:\Users\..\tiles\generator\tilegen_test.go:235]
    tilegen_test.go:243: FAIL:  GetDownloadJobs()
                        at: [C:\Users\..\tiles\generator\tilegen_test.go:237]


    tilegen_test.go:243: FAIL: 0 out of 4 expectation(s) were met.

	
                The code you are testing needs to make 4 more call(s).
                at: [C:\Users\..\tiles\generator\tilegen_test.go:243]
FAIL
coverage: 16.3% of statements
exit status 1
FAIL    abc/tiles/generator 4.894s