package csv

import (
	"encoding/csv"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func DownloadFile(c *gin.Context, filePath string) error {
	fileTmp, err := os.Open(filePath)
	defer fileTmp.Close()

	fileName := path.Base(filePath)
	if err != nil {
		return err
	}

	c.Header("Content-Type", "application/octet-stream")
	//Force browser download
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	//Browser download or preview
	c.Header("Content-Disposition", "inline;filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")

	c.File(filePath)

	return nil
}

func ExporToCSV(c *gin.Context, data [][]string, nameFile string) error {
	f, err := os.Create(nameFile)
	defer f.Close()

	if err != nil {
		return err
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(data)

	if err != nil {
		return err
	}

	return nil
}
