# DbmProviderToIabVendor

This prject will help you create a CSV of map of DBM Provider Id to Iab Vendor Id
DataFiles/vendor-list.json is taken from https://vendorlist.consensu.org/vendorlist.json
DataFiles/DBM_Providers.csv is taken from https://developers.google.com/authorized-buyers/rtb/data#reference-tables-click-to-download-csvs

# Usage
Checkout the code in $GOPATH/src
cd $GOPATH/src/DbmProviderToIabVendor
sudo go build
./DbmProviderToIabVendor > output.csv

This will create a csv file with columns like following.
```
Index, DBM_Provider_Id, IAB_Vendor_Id, IAB_Purpose_Id_Integer, Company_Name, Domain
```
