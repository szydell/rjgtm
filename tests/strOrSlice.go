package main

import (
	"fmt"
)

func main() {

	response := "DEFAULT;SET:2,KIL:1,GET:12209,DTA:0,ORD:0,ZPR:0,QRY:0,LKS:0,LKF:0,CTN:5,DRD:89,DWT:9,NTW:4,NTR:15489,NBW:9,NBR:30986,NR0:0,NR1:0,NR2:0,NR3:0,TTW:0,TTR:0,TRB:0,TBW:0,TBR:0,TR0:0,TR1:0,TR2:0,TR3:0,TR4:0,TC0:0,TC1:0,TC2:0,TC3:0,TC4:0,ZTR:0,DFL:52,DFS:6,JFL:2,JFS:14,JBB:2464,JFB:40960,JFW:10,JRL:3,JRP:6,JRE:2,JRI:3,JRO:6,JEX:0,DEX:0,CAT:193,CFE:0,CFS:0,CFT:0,CQS:0,CQT:0,CYS:0,CYT:0,BTD:9|TMP;SET:0,KIL:0,GET:0,DTA:0,ORD:0,ZPR:0,QRY:0,LKS:0,LKF:0,CTN:1,DRD:0,DWT:0,NTW:0,NTR:0,NBW:0,NBR:0,NR0:0,NR1:0,NR2:0,NR3:0,TTW:0,TTR:0,TRB:0,TBW:0,TBR:0,TR0:0,TR1:0,TR2:0,TR3:0,TR4:0,TC0:0,TC1:0,TC2:0,TC3:0,TC4:0,ZTR:0,DFL:29,DFS:1,JFL:0,JFS:2,JBB:0,JFB:0,JFW:0,JRL:0,JRP:0,JRE:0,JRI:0,JRO:0,JEX:0,DEX:0,CAT:58,CFE:0,CFS:0,CFT:0,CQS:0,CQT:0,CYS:0,CYT:0,BTD:0"

	buildJSON := "[{\""
	for _, char := range response {
		switch char {
		case 44:
			buildJSON = buildJSON + ",\""
		case 58:
			buildJSON = buildJSON + "\":"
		case 59:
			buildJSON = buildJSON + "\":{\""
		case 124:
			buildJSON = buildJSON + "},\""
		default:
			buildJSON = buildJSON + (string(char))
		}
	}
	buildJSON = buildJSON + "}}]"
	fmt.Println(buildJSON)
}
