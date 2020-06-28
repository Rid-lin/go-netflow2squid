package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// var dataLine = "0623.09:52:34.249 0623.09:53:20.879 13    217.20.152.247  443   8     192.168.65.82   40192 6   2  47         41724"

func init() {

}

func main() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		// fmt.Println(line)
		outStr, _ := parseNetFlowToSquidLine(line, "2020", "192.168.65.1")
		fmt.Println(outStr)
	}

}

// Input - string in netflow format №5
//Start             End               Sif   SrcIPaddress    SrcP  DIf   DstIPaddress    DstP    P Fl Pkts       Octets
//Output - squid log format default

func parseNetFlowToSquidLine(strIn, year, collectorIP string) (string, error) {
	var protocol string
	strArray := strings.Fields(strIn)
	// fmt.Println(str2)
	// str := strings.Join(strings.Fields(strIn), " ")
	// strArray := strings.Split(str, " ")
	if len(strArray) <= 0 {
		return "", nil
	}
	unixStampStr := unixStampFromNetflowDateStr(strArray[0], year)
	startOfResponse := unixStampFromNetflowDate(strArray[0], year)
	endOfResponse := unixStampFromNetflowDate(strArray[1], year)
	delayStr := strconv.FormatInt((endOfResponse/1000 - startOfResponse/1000), 10)
	// user = "-"

	switch strArray[8] {
	case "6":
		protocol = "TCP_PACKET"
	case "17":
		protocol = "UDP_PACKET"
	default:
		protocol = "OTHER_PACKET"

	}
	//Start             End               Sif   SrcIPaddress    SrcP  DIf   DstIPaddress    DstP    P Fl Pkts       Octets
	//
	out := fmt.Sprintf("%v %6v %v %v:%v/200 %v HEAD %v:%v - FIRSTUP_PARENT/%v packet/netflow", unixStampStr, delayStr, strArray[3], protocol, strArray[4], strArray[len(strArray)-1], strArray[6], strArray[7], collectorIP)
	return out, nil
}

func unixStampFromNetflowDateStr(str, year string) string {
	str = year + str
	normalizedDate, err := time.Parse("20060102.15:04:05.000", str)
	if err != nil {
		return ""
	}

	timeUnix := normalizedDate.Unix()
	timeUnixStr := strconv.FormatInt(timeUnix, 10)
	timeUnixNanoStr := strconv.FormatInt(((normalizedDate.UnixNano() - timeUnix*1000000000) / 1000000), 10)
	if len(timeUnixNanoStr) == 1 {
		timeUnixNanoStr = timeUnixNanoStr + "00"
	} else if len(timeUnixNanoStr) == 2 {
		timeUnixNanoStr = timeUnixNanoStr + "0"
	}

	out := fmt.Sprintf("%v.%v", timeUnixStr, timeUnixNanoStr)
	return out
}

func unixStampFromNetflowDate(str, year string) int64 {
	str = year + str
	normalizedDate, err := time.Parse("20060102.15:04:05.000", str)
	if err != nil {
		return 0
	}

	timeUnixMili := normalizedDate.UnixNano() / 1000000
	return timeUnixMili
}
