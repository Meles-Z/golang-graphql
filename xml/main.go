// package main

// import (
// 	"bytes"
// 	"context"
// 	"encoding/xml"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"time"
// )

// type Envelope struct {
// 	XmlName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
// 	Body    Body     `xml:"Body"`
// }
// type Body struct {
// 	Response ListOfContinentsByNameResponse `xml:"http://www.oorsprong.org/websamples.countryinfo ListOfContinentsByNameResponse"`
// }

// type ListOfContinentsByNameResponse struct {
// 	Result ListOfContinentsByNameResult `xml:"ListOfContinentsByNameResult"`
// }

// type ListOfContinentsByNameResult struct {
// 	Continents []Continent `xml:"tContinent"`
// }
// type Continent struct {
// 	Code string `xml:"sCode"`
// 	Name string `xml:"sName"`
// }

// func main() {
// 	var env Envelope

// 	url := "http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso"
// 	reqBody := `<?xml version="1.0" encoding="utf-8"?>
// 	<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
// 	<soap:Body>
// 		<ListOfContinentsByName xmlns="http://www.oorsprong.org/websamples.countryinfo">
// 		</ListOfContinentsByName>
// 	</soap:Body>
// 	</soap:Envelope>`

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
// 	defer cancel()
// 	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBufferString(reqBody))
// 	if err != nil {
// 		fmt.Printf("Error to create request :%v\n", err)
// 		return
// 	}
// 	req.Header.Set("Content-Type", "text/xml; charset=utf-8")

// 	client := &http.Client{}

// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Printf("Failed to send request and get response: %v\n", err)
// 		return
// 	}
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Printf("Failed to read response body: %v\n", err)
// 		return
// 	}

// 	if err := xml.Unmarshal(body, &env); err != nil {
// 		fmt.Printf("Failed to unmarshal response body to custom struct:%v\n", err)
// 		return
// 	}

// 	for _, v := range env.Body.Response.Result.Continents {
// 		fmt.Printf("name->%s | code -> %s\n", v.Name, v.Code)
// 	}
// }

// package main

// import (
// 	"bytes"
// 	"context"
// 	"encoding/xml"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"time"
// )

// // Define SOAP Envelope Response structure
// type Envelope struct {
// 	XMLName xml.Name `xml:"Envelope"`
// 	Body    Body     `xml:"Body"`
// }

// type Body struct {
// 	Response CountryISOCodeResponse `xml:"CountryISOCodeResponse"`
// }

// type CountryISOCodeResponse struct {
// 	Result string `xml:"CountryISOCodeResult"`
// }

// func main() {
// 	countryName := "Ethiopia" // change this to any country

// 	// SOAP body with parameter
// 	soapBody := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
// 	<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
// 	  <soap:Body>
// 		<CountryISOCode xmlns="http://www.oorsprong.org/websamples.countryinfo">
// 		  <sCountryName>%s</sCountryName>
// 		</CountryISOCode>
// 	  </soap:Body>
// 	</soap:Envelope>`, countryName)

// 	url := "http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso"

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBufferString(soapBody))
// 	if err != nil {
// 		fmt.Println("Request creation error:", err)
// 		return
// 	}

// 	// Required headers
// 	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
// 	req.Header.Set("SOAPAction", `"http://www.oorsprong.org/websamples.countryinfo/CountryISOCode"`)

// 	client := &http.Client{}
// 	res, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Request failed:", err)
// 		return
// 	}
// 	defer res.Body.Close()

// 	respBody, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println("Failed to read response:", err)
// 		return
// 	}

// 	var envelope Envelope
// 	if err := xml.Unmarshal(respBody, &envelope); err != nil {
// 		fmt.Println("Unmarshal failed:", err)
// 		fmt.Println("Raw Response:", string(respBody)) // debug help
// 		return
// 	}

// 	fmt.Printf("ISO Code for %s: %s\n", countryName, envelope.Body.Response.Result)
// }

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
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	Response CountryISOCodeResponse `xml:"CountryISOCodeResponse"`
}

type CountryISOCodeResponse struct {
	Result string `xml:"CountryISOCodeResult"`
}

func main() {
	http.HandleFunc("/country", func(w http.ResponseWriter, r *http.Request) {
		country := r.URL.Query().Get("country")
		if country == "" {
			http.Error(w, "Missing 'country' parameter", http.StatusBadRequest)
			return
		}
		Handler(country, w)
	})

	fmt.Println("Server started on 8080")
	http.ListenAndServe(":8080", nil)
}

func Handler(name string, w http.ResponseWriter) {
	soapBody := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <CountryISOCode xmlns="http://www.oorsprong.org/websamples.countryinfo">
      <sCountryName>%s</sCountryName>
    </CountryISOCode>
  </soap:Body>
</soap:Envelope>`, name)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	url := "http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso"
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBufferString(soapBody))
	if err != nil {
		http.Error(w, "Request creation failed", http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", `"http://www.oorsprong.org/websamples.countryinfo/CountryISOCode"`)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, "SOAP request failed", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "Reading response failed", http.StatusInternalServerError)
		return
	}

	var envelope Envelope
	if err := xml.Unmarshal(body, &envelope); err != nil {
		http.Error(w, "Unmarshal failed", http.StatusInternalServerError)
		fmt.Println("Debug response body:\n", string(body)) // log for debug
		return
	}

	fmt.Fprintf(w, "ISO Code for %s: %s\n", name, envelope.Body.Response.Result)
}
