package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"os"
)

type doc struct {
	XMLName    xml.Name `xml:"urlset"`
	Xmlns      string   `xml:"xmlns,attr"`
	XmlnsXhtml string   `xml:"xmlns:xhtml,attr"`
	Urls       []url
}

type url struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
	Links   []xhtml
}

type xhtml struct {
	XMLName  xml.Name `xml:"xhtml:link"`
	Rel      string   `xml:"rel,attr"`
	Hreflang string   `xml:"hreflang,attr"`
	Href     string   `xml:"href,attr"`
}

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf(err.Error())
	}

	defer file.Close()

	reader := csv.NewReader(file)

	data, err := reader.ReadAll()

	output := doc{
		Xmlns:      "http://www.sitemaps.org/schemas/sitemap/0.9",
		XmlnsXhtml: "http://www.w3.org/1999/xhtml",
	}

	output.Urls = make([]url, 0)

	if err != nil {
		fmt.Printf(err.Error())
	}

	for i := range data {
		for k := range data[i] {
			if i != 0 && len(data[i][k]) > 0 {
				l := []xhtml{
					xhtml{
						Rel:      "alternate",
						Hreflang: data[0][0],
						Href:     data[i][0],
					},
					xhtml{
						Rel:      "alternate",
						Hreflang: data[0][1],
						Href:     data[i][1],
					},
					xhtml{
						Rel:      "alternate",
						Hreflang: data[0][2],
						Href:     data[i][2],
					},
				}

				u := url{
					Loc:   data[i][k],
					Links: l,
				}

				output.Urls = append(output.Urls, u)
			}
		}
	}

	outFile, err := os.Create("sitemap.xml")
	defer outFile.Close()

	outXML, err := xml.MarshalIndent(output, " ", " ")
	outXML = []byte(xml.Header + string(outXML))

	outFile.Write(outXML)
}
