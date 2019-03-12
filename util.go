package main 

import(	
	"io/ioutil"
	"net"
	"net/url"
	"strconv"
)

func getInt(value string, defaultValue int) int {
	if tempValue, err := strconv.Atoi(value); err == nil {
		return tempValue
	}
	return defaultValue
}

func getDomainFromUrl(inputUrl string) (string) {
	u, err := url.Parse(inputUrl)
    if err != nil {    	
        return ""
    }
    host, _, e := net.SplitHostPort(u.Host) 
    if e != nil {
    	return u.Host
    }
    return host
}

func readFile(fileName string) string {
	dat, err := ioutil.ReadFile(fileName)
    if err != nil {
    	return ""
    }
    return string(dat)
}
