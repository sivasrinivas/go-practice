package main

import (
	"bytes"
	"io/ioutil"
	"strings"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	check(err)
	lines := strings.Split(string(data), "\n")
	var buffer bytes.Buffer
	for _, line := range lines {
		tokens := strings.Split(line, ",")
		q := fmt.Sprintf("update lillyjauat.Facility set geoLocationString='%s' where externalId='%s';", tokens[6]+","+tokens[4]+","+tokens[3],tokens[0])
		//q := fmt.Sprintf("Insert into lillyjauat.Facility (facilityName,externalId,isoCountryCode,timeZoneId,geoLocationString,isDeleted) " +
		//	" values('%s', '%s', '%s', '%s', '%s', %d);", "Facility"+strconv.Itoa(i), tokens[0], "JP", "Asia/Tokyo", tokens[3],0 )
		buffer.WriteString(q+"\n")
	}
	ioutil.WriteFile("output.txt", buffer.Bytes(), 0644)
}
