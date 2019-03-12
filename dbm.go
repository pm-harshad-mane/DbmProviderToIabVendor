package main

import(
	"encoding/csv"
	"io"
	"strings"
)

type DbmProvider struct {
	ProviderId	int
	CompanyName	string
	PolicyURL	string
	// we have ignored "domains" from CSV
}

type DomainToDbmProviderMap map[string]DbmProvider

func generateDbmMap(theCsvString string) (DomainToDbmProviderMap) {
	aDomainToDbmProviderMap := make(DomainToDbmProviderMap)
	r := csv.NewReader(strings.NewReader(theCsvString))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			// TODO: need to handle				
			continue
		}
		aDbmProvider := DbmProvider{
			getInt(record[0], 0),
			record[1],
			record[2],
		}
		aDomainToDbmProviderMap[ getDomainFromUrl(record[2]) ] = aDbmProvider		
	}
	return aDomainToDbmProviderMap
}
