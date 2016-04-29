package server

import (
	"io/ioutil"
	"net/http"

	"path/filepath"

	"github.com/EMC-CMD/cf-persist-service-broker/storage"
	"github.com/gin-gonic/gin"
)

func CatalogHandler(c *gin.Context) {
	c.Status(http.StatusOK)
	p, _ := filepath.Abs("templates/catalog.json")
	c.File(p)
}

func ProvisioningHandler(c *gin.Context) {
	_, err := CreateVolume(&storage.ScaleIODriver{})
	if err != nil {
		c.JSON(500, gin.H{})
	}
	c.JSON(http.StatusCreated, gin.H{})
}

func CreateVolume(driver storage.StorageDriver) (*storage.Volume, error) {
	return driver.VolumeCreate(storage.Context{}, "", &storage.VolumeCreateOpts{})
}

func DeprovisionHandler(c *gin.Context) {
	err := RemoveVolume(&storage.ScaleIODriver{})
	if err != nil {
		c.JSON(500, gin.H{})
	}
	c.JSON(http.StatusOK, gin.H{})
}

func RemoveVolume(driver storage.StorageDriver) error {
	return driver.VolumeRemove(storage.Context{}, "", &storage.VolumeCreateOpts{})
}

func BindingHandler(c *gin.Context) {
	body, _ := ioutil.ReadFile("fixtures/create_binding_response.json")
	c.String(http.StatusCreated, string(body))
}

func UnbindingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
