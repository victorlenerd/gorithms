package main

import "fmt"

/*
	Examples target 7, [2, 3] - [3, 2, 2]
	Examples target 7, [5, 3, 4, 7] - [4,3]
	Examples target 7, [2, 4] - null
	Examples target 8, [2, 3, 5] - [2,2,2,2]
	Examples target 300, [7, 14] - null
*/

func main()  {

	/*
		declare list for items 0....target

		set all the items to null

		set the first item to []

		loop for each item in the input arr

			loop for each item in the list of targets

				if item in input arr is greater than or equal to item in list

					get value of complement of item in input array and item in list

					set value of item in list to value of complement plus target
	*/

	target := 8

	denominations := []int{2, 3, 5}

	targetList := make([]*[]int, target + 1)

	for i := range targetList {
		targetList[uint(i)] = nil
	}

	targetList[0] = &[]int{}

	for _, denomination := range denominations {

		for localTargetIndex := 1;  localTargetIndex < len(targetList); localTargetIndex++ {

			if denomination <= localTargetIndex {

				complement := localTargetIndex - denomination

				localTargetValues := targetList[complement]

				newValues := []int{denomination}

				if localTargetValues != nil && targetList[localTargetIndex] == nil {

					for _, lT := range *localTargetValues {
						newValues = append(newValues, lT)
					}

					targetList[localTargetIndex] = &newValues
				}
			}

		}
	}


	fmt.Println(targetList[target])
}
