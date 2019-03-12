package main

import(
	"fmt"
)

func simpleMapCsvOutput(domainToDbmProviderMap DomainToDbmProviderMap, domainToIabVendorMap DomainToIabVendorMap){
	index := 1
	for domain, dbmVendor := range domainToDbmProviderMap {
		iabVendorId := 0
		iabPurposeIdInteger := 0
		iabVendor, found := domainToIabVendorMap[domain]		
		if found {
			iabVendorId = iabVendor.Id
			iabPurposeIdInteger = int(getPurposeIdInteger(iabVendor.PurposeIds))
		}
		if index == 1 {
			fmt.Println( "Index, DBM_Provider_Id, IAB_Vendor_Id, IAB_Purpose_Id_Integer, Company_Name, Domain" )
		}
        fmt.Println( index, ",", dbmVendor.ProviderId, ",", iabVendorId, ",", iabPurposeIdInteger, ",\"", dbmVendor.CompanyName, "\",", domain )
        index = index + 1
    }
}
