package alpr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Client :
type Client struct {
	SecretKey string
}

var (
	instanceClient *Client
)

// NewClient :
func NewClient() *Client {
	if instanceClient == nil {
		instanceClient = new(Client)
	}
	return instanceClient
}

// Initialize :
func (c *Client) Initialize(secretKey string) *Client {
	c.SecretKey = secretKey
	return c
}

// Execute :
func (c *Client) Execute(imageBytes string) (Response, error) {
	var response Response
	var rpr ResponsePlateRecognition

	body := strings.NewReader(imageBytes)
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.openalpr.com/v2/recognize_bytes?secret_key=%s&recognize_vehicle=1&country=id&state=id&return_image=0&topn=10", c.SecretKey), body)
	if err != nil {
		return response, err
	}

	/**
	 * ------------------
	 * NOTE : set header
	 * ------------------
	 */
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	/**
	 * ----------------------------------------------
	 * NOTE : invoke api.openalpr untuk mendapatkan
	 * detail plat nomor kendaraan (mobil)
	 * ----------------------------------------------
	 */
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()

	responseData, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(responseData, &rpr)

	result := rpr.Results[0]

	return Response{
		Plate:     result.Plate,
		Region:    result.Region,
		Color:     result.Vehicle.Color[0].Name,
		Make:      result.Vehicle.Make[0].Name,
		BodyType:  result.Vehicle.BodyType[0].Name,
		MakeModel: result.Vehicle.MakeModel[0].Name,
	}, nil
}
