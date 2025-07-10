
package main

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Envelope struct {
	XmlName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    Body     `xml:"Body"`
}
type Body struct {
	Response ListOfContinentsByNameResponse `xml:"http://www.oorsprong.org/websamples.countryinfo ListOfContinentsByNameResponse"`
}

type ListOfContinentsByNameResponse struct {
	Result ListOfContinentsByNameResult `xml:"ListOfContinentsByNameResult"`
}

type ListOfContinentsByNameResult struct {
	Continents []Continent `xml:"tContinent"`
}
type Continent struct {
	Code string `xml:"sCode"`
	Name string `xml:"sName"`
}

func main() {
	var env Envelope

	url := "http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso"
	reqBody := `<?xml version="1.0" encoding="utf-8"?>
	<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Body>
		<ListOfContinentsByName xmlns="http://www.oorsprong.org/websamples.countryinfo">
		</ListOfContinentsByName>
	</soap:Body>
	</soap:Envelope>`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBufferString(reqBody))
	if err != nil {
		fmt.Printf("Error to create request :%v\n", err)
		return
	}
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to send request and get response: %v\n", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %v\n", err)
		return
	}

	if err := xml.Unmarshal(body, &env); err != nil {
		fmt.Printf("Failed to unmarshal response body to custom struct:%v\n", err)
		return
	}

	for _, v := range env.Body.Response.Result.Continents {
		fmt.Printf("name->%s | code -> %s\n", v.Name, v.Code)
	}
}
