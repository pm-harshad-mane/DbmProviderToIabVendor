package main

func main() {
	domainToDbmProviderMap := generateDbmMap(readFile("DataFiles/DBM_Providers.csv") )
	domainToIabVendorMap := generateIabMap(readFile("DataFiles/vendor-list.json"))
	simpleMapCsvOutput(domainToDbmProviderMap, domainToIabVendorMap)
}
