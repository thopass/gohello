package gohello

import (
	"fmt"
	"strconv"
)

func PrintTitle(title string) {
	if title == "" {
		title = "Uknown application"
	}

	printStringWithFrame(title)
}

func PrintVersion(major, minor, build int) {
	strMajor := strconv.Itoa(major)
	strMinor := strconv.Itoa(minor)
	strBuild := strconv.Itoa(build)
	// let's assume that there will be no more than 9999 builds for version
	switch len(strBuild) {
	case 1:
		strBuild = "000" + strBuild
	case 2:
		strBuild = "00" + strBuild
	case 3:
		strBuild = "0" + strBuild
	default:
	}

	title := "Version: "
	strFullver := strMajor + "." + strMinor + "." + strBuild

	printStringWithFrame(title + strFullver)

}

func printStringWithFrame(content string) {
	textLen := len(content)
	buf := make([]byte, textLen+4)

	for x := 0; x < (textLen + 4); x++ {
		buf[x] = '*'
	}

	fmt.Println(string(buf))
	fmt.Println("*", content, "*")
	fmt.Println(string(buf))
}
