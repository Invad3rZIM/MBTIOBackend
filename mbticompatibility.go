package main

type MBTIChart struct {
	compatibilities map[string][]int
	codes           map[string]int
}

func NewMBTIChart() *MBTIChart {
	return &MBTIChart{
		compatibilities: *GenCompatibilities(),
		codes:           *GenMBTICodes(),
	}
}

//returns the compatibity between Personallity p & Personallity q
func (mbti *MBTIChart) GetMatchup(p string, q string) float64 {
	return float64(mbti.compatibilities[p][mbti.Codify(q)])
}

//converts 4 Character MBTI to integer for hashing purposes
func (mbti *MBTIChart) Codify(str string) int {
	return mbti.codes[str]
}

//hardcoded - work into psql cloud db if time allows
func GenCompatibilities() *map[string][]int {
	mbti := make(map[string][]int)

	mbti["ENTJ"] = []int{0, 0, 2, 2, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}
	mbti["ENTP"] = []int{0, 0, 2, 1, 0, 0, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0}
	mbti["ENFJ"] = []int{2, 2, 2, 1, 0, 0, 1, 2, 1, 1, 1, 1, 0, 0, 1, 3}
	mbti["ENFP"] = []int{2, 1, 1, 2, 0, 0, 2, 2, 2, 1, 1, 1, 0, 0, 1, 1}
	mbti["ESTJ"] = []int{0, 0, 0, 0, 0, 0, 2, 2, 0, 0, 0, 0, 0, 0, 1, 1}
	mbti["ESTP"] = []int{0, 0, 0, 0, 0, 0, 2, 1, 0, 0, 0, 0, 0, 0, 3, 1}
	mbti["ESFJ"] = []int{0, 0, 1, 2, 2, 2, 2, 1, 0, 0, 1, 3, 1, 1, 1, 1}
	mbti["ESFP"] = []int{0, 0, 2, 2, 2, 1, 1, 2, 0, 0, 1, 1, 3, 1, 1, 1}
	mbti["INTJ"] = []int{0, 0, 1, 2, 0, 0, 0, 0, 0, 0, 1, 2, 0, 0, 0, 0}
	mbti["INTP"] = []int{0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 2, 2, 0, 0, 0, 0}
	mbti["INFJ"] = []int{1, 3, 1, 1, 0, 0, 1, 1, 1, 2, 2, 1, 0, 0, 2, 2}
	mbti["INFP"] = []int{1, 1, 1, 1, 0, 0, 3, 1, 2, 2, 1, 2, 0, 0, 2, 1}
	mbti["ISTJ"] = []int{0, 0, 0, 0, 0, 0, 1, 3, 0, 0, 0, 0, 0, 0, 1, 2}
	mbti["ISTP"] = []int{0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 2, 2}
	mbti["ISFJ"] = []int{0, 0, 1, 1, 1, 3, 1, 1, 0, 0, 2, 2, 1, 2, 2, 1}
	mbti["ISFP"] = []int{0, 0, 3, 1, 1, 1, 1, 1, 0, 0, 2, 1, 2, 2, 1, 2}

	return &mbti
}

//Creates Random MBTI Personallity... used for testing purposes
func (mbti *MBTIChart) RandomMBTI() string {
	r := GenInt(0, 16)

	for key, value := range mbti.codes {
		if value == r {
			return key
		}
	}

	return "XXXX"
}

//hardcoded - work into psql cloud db if time allows
func GenMBTICodes() *map[string]int {
	code := make(map[string]int)

	code["ENTJ"] = 0
	code["ENTP"] = 1
	code["ENFJ"] = 2
	code["ENFP"] = 3
	code["ESTJ"] = 4
	code["ESTP"] = 5
	code["ESFJ"] = 6
	code["ESFP"] = 7
	code["INTJ"] = 8
	code["INTP"] = 9
	code["INFJ"] = 10
	code["INFP"] = 11
	code["ISTJ"] = 12
	code["ISTP"] = 13
	code["ISFJ"] = 14
	code["ISFP"] = 15

	return &code
}
