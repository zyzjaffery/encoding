package main

import (
	"encoding/xml"
	"fmt"
	"github.com/xtraclabs/xtrac-xsd-to-go/generator"
	"github.com/xtraclabs/xtrac-xsd-to-go/xsd"
	"io/ioutil"
	"os"
)

func buildCatalog(xsdDir string) (*xsd.Catalog, error) {
	loadXsd := func(xsdFileName string) (*xsd.Schema, error) {
		bs, err := ioutil.ReadFile(xsdFileName)
		if err != nil {
			return nil, err
		}

		s := new(xsd.Schema)
		err = xml.Unmarshal(bs, &s)
		if err != nil {
			return nil, err
		}
		return s, nil
	}

	elementsSchema, err := loadXsd(xsdDir + "/XtracElements.xsd")
	if err != nil {
		return nil, err
	}

	typesSchema, err := loadXsd(xsdDir + "/XtracTypes.xsd")
	if err != nil {
		return nil, err
	}

	catalog := xsd.NewCatalog()
	catalog.CatalogComplexTypes(elementsSchema)
	catalog.CatalogSimpleTypes(elementsSchema)
	catalog.CatalogSimpleTypes(typesSchema)

	return catalog, nil
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: genstructs struct-name type-name path-to-xsd")
		return
	}

	cat, err := buildCatalog(args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	gen := generator.NewGenerator(cat)
	err = gen.GenerateComplexType(args[0], args[1], make(map[string]string))
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

}
