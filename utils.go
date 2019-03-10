package main

import (
	"math/rand"
	"time"
)

var (
	girls = []string{"Ana", "Camilla", "Sasha", "Madeline", "Julia", "Jackie", "Helen", "Mariah", "Diana", "Sarah", "Devin", "Amy", "Maria", "Angela", "Becky", "Jill", "Tessie"}
	boys  = []string{"Kirk", "Matt", "Nick", "Rob", "Mike", "Sam", "Joe", "John", "Caleb", "Tom", "Jake", "Jimmy", "Dan", "Chad", "Dave", "Zeph", "Will"}
)

//0 = girl name, 1 = boy name
func GenName(sex int) string {
	if sex == 0 {
		return girls[GenInt(0, len(girls))]
	} else {
		return boys[GenInt(0, len(boys))]
	}
}

//generates a weighted sexual interset
func GenInterest(sex int) string {
	r := GenInt(0, 10)

	if sex == 0 {
		switch {
		case r < 8:
			return "M"
		case r < 9:
			return "F"
		default:
			return "B"
		}
	} else {
		switch {
		case r < 8:
			return "F"
		case r < 9:
			return "M"
		default:
			return "B"
		}
	}
}

//filler float generator function. used for procedural gps creation
func GenFloat(min float64, max float64) float64 {
	rand.Seed(time.Now().UnixNano())

	return (max-min)*rand.Float64() + min
}

//Used for procedurally generating any integers
func GenInt(min int, max int) int {
	return int((float64(max)-float64(min))*rand.Float64() + float64(min))
}

//used for procedurally generating heights (60 = 5'0, 66 = 5'6, 72 = 6'0, 75 = 6'3)
func GenHeight(min int, max int) int {
	totalInches := int((float64(max)-float64(min))*rand.Float64() + float64(min))
	return (totalInches/12)*100 + (totalInches % 12)
}
