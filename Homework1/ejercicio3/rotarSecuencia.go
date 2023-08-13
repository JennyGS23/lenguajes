package main

import "fmt"

func rotateSequenceElements(pointerList *[]int, numberMovements int, direction string) {
	tempList := make([]int, len(*pointerList))
	lengthList := len(*pointerList)
	for i := 0; i < len(*pointerList); i++ {
		if direction == "right" {
			addPosition := i + numberMovements
			if addPosition > lengthList-1 {
				newPosition := addPosition - lengthList
				tempList[newPosition] = (*pointerList)[i]

			} else {
				tempList[addPosition] = (*pointerList)[i]
			}

		} else {
			addPosition := i - numberMovements
			if addPosition < 0 {
				newPosition := addPosition + lengthList
				tempList[newPosition] = (*pointerList)[i]

			} else {
				tempList[addPosition] = (*pointerList)[i]
			}

		}

	}
	*pointerList = tempList
}

func main() {
	intList := []int{1, 2, 3, 4, 5}

	pointerList := &intList

	//Use "right" or "left"
	rotateSequenceElements(pointerList, 2, "right")
	fmt.Print(intList)

}
