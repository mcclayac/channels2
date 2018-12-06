package main

import (
	"fmt"
	"strconv"
	"time"
)

var frameId = 0
var frameName = ""
var assemblyArrangement[3] string


func main () {

	fmt.Println("Go Channels 2")

	assemblyArrangement[0] = "frame"
	assemblyArrangement[1] = "body"
	assemblyArrangement[2] = "interior"


	framesToCreate := len(assemblyArrangement)
	frameInforChan := make(chan string)

	for stageNumber := 0; stageNumber < framesToCreate; stageNumber++ {
		go assemblyStage(frameInforChan, assemblyArrangement[stageNumber],
			stageNumber, framesToCreate)
		time.Sleep(time.Millisecond * 1000)
		//fmt.Println("Interaction complete")
	}

	/* for i :=0; i < framesToCreate; i++ {
		go assembleFrame(frameInforChan)
		go addBody(frameInforChan)
		go addInterior(frameInforChan)

		time.Sleep(time.Microsecond * 1000)
	}*/

}


func assemblyStage(framInfoChan chan string, stage string, stageNumber int, framesToCreate int ) {
		nextStage := "paint"
		if stageNumber < framesToCreate {
			frameName = "Frame ID" + strconv.Itoa(stageNumber)
			if stageNumber != framesToCreate-1 {
				nextStage = assemblyArrangement[stageNumber+1]
			}
		}
		fmt.Println("Add", stage,"and proceed to",nextStage)
		framInfoChan <- frameName
		time.Sleep(time.Microsecond * 10)
}

func assembleFrame(frameInfoChan chan string) {

	frameId++
	frameName = "Frame ID" + strconv.Itoa(frameId)
	fmt.Println("Frame assembly complete", frameName, "Proceed to body")
	frameInfoChan <- frameName
	time.Sleep(time.Microsecond * 5)
}

func addBody(frameInfoChan chan string) {
	body := <- frameInfoChan
	fmt.Println("Add Body to", body,"and proceed to interior" )
	frameInfoChan <- frameName
	time.Sleep(time.Microsecond * 5)

}

func addInterior(frameInfoChan chan string) {
	interior := <- frameInfoChan
	fmt.Println("Add Interior to", interior,"and proceed to paint" )
	time.Sleep(time.Microsecond * 5)
}