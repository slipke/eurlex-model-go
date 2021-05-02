# eurlex-model-go

# Overview

[eurlex-model-go](https://github.com/slipke/eurlex-model-go) implements the XML data model used by [EUR-Lex](https://eur-lex.europa.eu/homepage.html) in Golang. It is required by [eurlex-ws-go](https://github.com/slipke/eurlex-ws-go), which implements the Webservice (Search), and [Cellar](https://op.europa.eu/documents/10530/676542/ao10463_annex_17_cellar_dissemination_interface_en.pdf).

## Install

```bash
go get -u github.com/slipke/eurlex-model-go
```

## Usage

Retrieve the XML from either the Webservice or Cellar, or alternatively directly from the website (i.e. [celex:32008R1272](https://eur-lex.europa.eu/legal-content/EN/ALL/?uri=celex%3A32008R1272) and click on ["Download notice"](https://eur-lex.europa.eu/download-notice.html?legalContentId=cellar:6bf54b59-7673-461b-b8e1-f24c545cbd3c&noticeType=branch&callingUrl=%2Flegal-content%2FEN%2FALL%2F%3Furi%3Dcelex%253A32008R1272&lng=EN)). Afterwards, the resulting XML file can be parsed by using the exemplary code:

```go
import (
	"encoding/xml"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	model "github.com/slipke/eurlex-model-go"
)


func main() {
	xmlBytes, err := ioutil.ReadFile("cellar_6bf54b59-7673-461b-b8e1-f24c545cbd3c.xml")
	if err != nil {
		log.Fatalf("Failed to load file %s: %s", file, err)
	}

	var n *model.Notice

	err = xml.Unmarshal(xmlBytes, &n)
	if err != nil {
		log.Fatalf("Failed to parse xml: %s", err)
	}

	// Print notice
	log.Info(n)
}
```

## Resources

- [Cellar Interface](https://op.europa.eu/documents/10530/676542/ao10463_annex_17_cellar_dissemination_interface_en.pdf)
- [eurlex-ws-go](https://github.com/slipke/eurlex-ws-go)
- [EUR-Lex Help](https://eur-lex.europa.eu/content/help/webservice.html)
- [Web Service User Manual](https://eur-lex.europa.eu/content/tools/webservices/SearchWebServiceUserManual_v2.00.pdf)
- [Web Service Query Metadata](https://eur-lex.europa.eu/content/tools/webservices/WebServicesqueryMetadata.pdf)
- [Data Extraction Using Web Services](https://eur-lex.europa.eu/content/tools/webservices/DataExtractionUsingWebServices-v1.00.pdf)
