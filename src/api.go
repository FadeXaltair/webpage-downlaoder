package src

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DownloadPage function is used to get the url from the request body.
func DownloadPage(c *gin.Context) {
	var body Body
	// getting the retrylimit from the path parameters
	retrylimit := c.Param("retry-limit")
	limit, _ := strconv.Atoi(retrylimit)
	if limit > 10 {
		limit = 10
	}
	//binding the request body
	err := c.Bind(&body)
	if err != nil {
		Error(err)
		return
	}
	data := Body{
		Url:      body.Url,
		Filename: body.Filename,
	}
	// downloading the webpage from the url
	file, err := File(data.Url, data.Filename, limit)
	if err != nil {
		Error(err)
		return
	}
	// response status
	resp := Response{
		Id:        data.Filename,
		Url:       data.Url,
		SourceUrl: file,
	}
	c.JSON(http.StatusOK, resp)

}

// File function will open the Url and download the webpage locally to the given directory
func File(url, filename string, limit int) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		Error(err)
	}

	f, err := os.Create(filename)
	if err != nil {
		return "File is not created after tries", err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Println(err)
		Error(err)
	}
	projectlocation := "/home/altair/go/src/webpage-download/" + filename
	Location := "/home/altair/Downloads/files/" + filename + ".html"
	err = os.Rename(projectlocation, Location)
	if err != nil {
		log.Fatal(err)
	}
	return Location, nil
}
