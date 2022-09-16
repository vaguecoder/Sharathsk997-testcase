package generator

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"abc/internal/satellite/tiles/download"
	"abc/internal/satellite/usecase"
	"abc/internal/satellite/api/presenter"
	"abc/internal/satellite/models"
	"abc/internal/satellite/tiles/storageblob"

)

// FeaturesProcessor struct
type FeaturesProcessor struct {
	baseURL string
	storage usecase.BlobStorage
	dImages usecase.DownloadImages
	fileDir string
}

// NewFeaturesProcessor d
func NewFeaturesProcessor(url string, bc usecase.BlobStorage, si usecase.DownloadImages, fileDir string) *FeaturesProcessor {
	baseURL := url + "/resources"
	return &FeaturesProcessor{baseURL: baseURL, storage: bc, dImages: si, fileDir: fileDir}
}

// CreateMapFeatureResponse f
func (fp *FeaturesProcessor) CreateMapFeatureResponse(ctx context.Context, foundTiles []*presenter.MfsTiles, compressionType string, logMsg mfsvc.LogMessage) (*presenter.MapfeatureResponse, error) {

	//.................
	// mapFeatures := make([]*presenter.MapFeatures, 0)
	// chanErr := make(chan error)
	// chanResp := make(chan presenter.MapFeatures)
	// mapFeatureRes := presenter.MapfeatureResponse{
	// 	Features: presenter.Features{
	// 		MapFeatures: mapFeatures,
	// 	},
	// }.........

	return &mapFeatureRes, nil
}

func CheckTileStatus(chanErr chan error, chanResp chan presenter.MapFeatures, tile *presenter.MfsTiles, ctx context.Context, compressionType string, logMsg mfsvc.LogMessage, fp *FeaturesProcessor) {
	
	................
	requestData := models.ResourceRequest{Level: tile.Level, TileID: tile.ID, TextureFormat: compressionType}
	err := checkRefreshRequired(ctx, &requestData, fp, logMsg)
	if err != nil {

		chanErr <- err
		logMsg.Error("SA_Error", err.Error())
		fmt.Print(err.Error())
	}.........


	chanResp <- mapFeature
}

func checkRefreshRequired(ctx context.Context, requestData *models.ResourceRequest, fp *FeaturesProcessor, logMsg mfsvc.LogMessage) error {
	//...............
	}
	return nil
}


/// these are code from different .go file.
//here under  UploadBlob function getBlobUrlIndex is called.. do do we need to write test case for getBlobUrlIndex?
//and test case written for UploadBlob dosent cover getBlobUrlIndex function.
// test case written for UploadBlob is similar to Case:1 test case
func (sa *StorageAccount) UploadBlob(fileName string, blobDetails BlobDetails) {
	// Here's how to upload a blob.
	blobURL := sa.containerURL[sa.getBlobUrlIndex(fileName)].NewBlockBlobURL(fileName)
	///	..........
}

func (sa *StorageAccount) getBlobUrlIndex(filename string) int {
//.........
}