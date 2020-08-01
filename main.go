package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	year, collectorIP, gmt string
)

func init() {
	flag.StringVar(&gmt, "gmt", "+0500", "GMT time zone for the current collector")
	flag.StringVar(&year, "year", "2020", "Year when the ft file was recorded")
	flag.StringVar(&collectorIP, "ip", "192.168.65.1", "Ip address of the netflow collector")
	flag.Parse()

}

func main() {
	// line := "0623.09:57:16.569 0623.09:58:06.299 8     192.168.65.143  38978 13    87.250.251.119  443   6   2  553        612073"
	// outStr, err := parseNetFlowToSquidLine(line, year, collectorIP)
	// if err != nil {
	// 	fmt.Errorf("Error, %v", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(outStr)
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		outStr, err := parseNetFlowToSquidLine(line, year, collectorIP, gmt)
		if err != nil {
			fmt.Errorf("Error, %v", err)
			os.Exit(1)
		}
		fmt.Println(outStr)
	}

}

//Input - string in netflow format №5
//Start             End               Sif   SrcIPaddress    SrcP  DIf   DstIPaddress    DstP    P Fl Pkts       Octets
//Output - squid log format default

func parseNetFlowToSquidLine(strIn, year, collectorIP, gmt string) (string, error) {
	var protocol string
	strArray := strings.Fields(strIn)
	if len(strArray) <= 0 {
		return "", nil
	}
	unixStampStr := unixStampFromNetflowDateStr(strArray[0], year, gmt)
	startOfResponse := unixStampFromNetflowDate(strArray[0], year, gmt)
	endOfResponse := unixStampFromNetflowDate(strArray[1], year, gmt)
	delayStr := strconv.FormatInt((endOfResponse/1000 - startOfResponse/1000), 10)
	// user = "-"

	switch strArray[8] {
	case "6":
		protocol = "TCP_PACKET"
	case "17":
		protocol = "UDP_PACKET"
	case "1":
		protocol = "ICMP_PACKET"

	default:
		protocol = "OTHER_PACKET"

	}
	//Start             End               Sif   SrcIPaddress    SrcP  DIf   DstIPaddress    DstP    P Fl Pkts       Octets
	//
	out := fmt.Sprintf("%v %6v %v %v/- %v HEAD %v:%v - FIRSTUP_PARENT/%v packet_netflow/%v", unixStampStr, delayStr, strArray[3], protocol, strArray[len(strArray)-1], strArray[6], strArray[7], collectorIP, strArray[4])
	return out, nil
}

func unixStampFromNetflowDateStr(str, year, gmt string) string {
	str = year + str + gmt
	normalizedDate, err := time.Parse("20060102.15:04:05.000-0700", str)
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

func unixStampFromNetflowDate(str, year, gmt string) int64 {
	str = year + str + gmt
	normalizedDate, err := time.Parse("20060102.15:04:05.000-0700", str)
	if err != nil {
		return 0
	}

	timeUnixMili := normalizedDate.UnixNano() / 1000000
	return timeUnixMili
}
