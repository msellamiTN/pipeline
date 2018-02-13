package defaults

import (
	"github.com/banzaicloud/banzai-types/constants"
	"github.com/banzaicloud/pipeline/model"
	"github.com/banzaicloud/banzai-types/components"
	"github.com/banzaicloud/banzai-types/components/amazon"
	"github.com/banzaicloud/banzai-types/components/azure"
	"github.com/banzaicloud/banzai-types/components/google"
)

type DefaultCreateRequestGKE struct {
	DefaultModel
	Location         string `gorm:"default:'us-central1-a'"`
	NodeInstanceType string `gorm:"default:'n1-standard-1'"`
	NodeCount        int    `gorm:"default:1"`
	NodeVersion      string `gorm:"default:'1.7.12-gke.1'"`
	MasterVersion    string `gorm:"default:'1.7.12-gke.1'"`
}

func (*DefaultCreateRequestGKE) TableName() string {
	return "google_default_create"
}

func (d *DefaultCreateRequestGKE) SaveDefaultInstance() error {
	return save(d)
}

func (d *DefaultCreateRequestGKE) IsDefinedBefore() bool {
	database := model.GetDB()
	database.First(&d)
	return d.ID != 0
}

func (d *DefaultCreateRequestGKE) GetType() string {
	return constants.Google
}

func (d *DefaultCreateRequestGKE) GetDefaultCreateClusterRequest() *components.CreateClusterRequest {
	loadFirst(&d)

	return &components.CreateClusterRequest{
		Name:             "", // todo good?
		Location:         d.Location,
		Cloud:            constants.Google,
		NodeInstanceType: d.NodeInstanceType,
		Properties: struct {
			CreateClusterAmazon *amazon.CreateClusterAmazon `json:"amazon,omitempty"`
			CreateClusterAzure  *azure.CreateClusterAzure   `json:"azure,omitempty"`
			CreateClusterGoogle *google.CreateClusterGoogle `json:"google,omitempty"`
		}{
			CreateClusterGoogle: &google.CreateClusterGoogle{
				Project: "", // todo good?
				Node: &google.GoogleNode{
					Count:   d.NodeCount,
					Version: d.NodeVersion,
				},
				Master: &google.GoogleMaster{
					Version: d.MasterVersion,
				},
			},
		},
	}
}