package main

func main() {
	domainToDbmProviderMap := generateDbmMap(readFile("DataFiles/DBM_Providers.csv") )
	domainToIabVendorMap := generateIabMap(readFile("DataFiles/vendor-list.json"))
	domainToDbm250ProviderMap := generateDbmMap(readFile("DataFiles/DBM_PM_250_Providers.csv") )
	//simpleMapCsvOutput(domainToDbmProviderMap, domainToIabVendorMap)
	staticListMapCsvOutput(domainToDbmProviderMap, domainToIabVendorMap, domainToDbm250ProviderMap)	
}
