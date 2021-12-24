package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// input := ReadInput()
	input := "1,3,4,4,5,1,2,4"
	// input := "1,2,2"
	// input := "1,2,2,3,5"

	iUserGroups, err := GetIntsArray(input)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("\n Input Data: %v\n", iUserGroups)
	}

	layout := CreateLayout()

	PrintLayout(layout)

	layout, err = SeatUsers(layout, iUserGroups)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		PrintLayout(layout)
	}

}

func SeatUsers(layout [][]int, iUserGroups []int) ([][]int, error) {
	currentRow := 0
	currentSeat := 0

	rows := len(layout[0])
	seats := len(layout)

	sectionSize := rows * seats
	usersToSit := 0
	for _, users := range iUserGroups {
		usersToSit += users
	}

	if sectionSize < usersToSit {
		return layout, errors.New("seating: there are not enough sectionSize")
	}

	for index, users := range iUserGroups {

		for i := 0; i < users; i++ {
			layout[currentSeat][currentRow] = index + 1

			if currentSeat%2 != 1 {
				if currentRow == rows-1 {
					currentSeat++
				} else {
					currentRow++
				}
			} else {
				if currentRow == 0 {
					currentSeat++
				} else {
					currentRow--
				}
			}
		}
	}
	return layout, nil
}

func CreateLayout() [][]int {

	seats := 3
	rows := 8

	layout := make([][]int, seats)
	for i := 0; i < seats; i++ {
		layout[i] = make([]int, rows)
	}

	// https://stackoverflow.com/questions/39804861/what-is-a-concise-way-to-create-a-2d-slice-in-go/39806983
	// or this
	// var layout = make([][]int, seats)
	// seats := make([]int, seats*rows)
	// for i := 0; i < seats; i++ {
	// 	layout[i] = seats[i*rows : (i+1)*rows]
	// }

	return layout
}

func PrintLayout(layout [][]int) {
	fmt.Println()
	for i := 0; i < len(layout); i++ {
		fmt.Println(layout[i])
	}
}

func GetIntsArray(input string) ([]int, error) {
	sUserGroups := strings.Split(input, ",")
	iUserGroups := make([]int, len(sUserGroups))

	for i, v := range sUserGroups {

		userGroup, err := strconv.Atoi(v)
		if userGroup == 0 || err != nil {
			return iUserGroups, errors.New("input: user group contains invalid value")
		}
		iUserGroups[i], _ = userGroup, err
	}
	return iUserGroups, nil
}

func ReadInput() string {
	var input string
	fmt.Print("Enter a list of group of users (e.g. 1,3,4,4,5,1,2,4): ")
	fmt.Scanf("%s", &input)
	return input
}
