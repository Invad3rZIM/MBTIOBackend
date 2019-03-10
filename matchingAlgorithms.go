package main

import (
	"container/heap"
	"math"
)

func (h *Handler) GetMatches(uid int, max int) ([]*User, error) {
	matchedUsers := []*User{}
	p, err := h.GetUser(uid)

	if err != nil {
		return nil, err
	}

	uc := h.UserCache

	for _, q := range uc.cache {
		if p == q { //cannot match with self
			continue
		}

		//exclude all the people who haven't actually made their profile
		if !q.ProfileReady {
			continue
		}

		//if sexual interests align, add them to the dating pool
		if p.IsSexuallyCompatible(q) {
			//calculate score between p & q
			score := h.GetMatchup(p.MBTI, q.MBTI)
			score += p.AgeMatchup(q)
			score += p.HeightMatchup(q)

			//omit really, *really* bad scores
			if score < -1 {
				continue
			}

			r := *q
			r.Score = score
			q := &r

			matchedUsers = append(matchedUsers, q)
		}
	}

	priorityQ := make(PriorityQueue, len(matchedUsers))

	for index, q := range matchedUsers {
		priorityQ[index] = &Node{
			Distance: p.CalcDistance(q),
			Pointer:  q,
		}
	}

	//sorts priorityQ by distance
	heap.Init(&priorityQ)

	matchedUsers = []*User{}

	// Take the items out; they arrive in decreasing priority order.
	for len(priorityQ) > 0 && len(matchedUsers) < max {
		node := heap.Pop(&priorityQ).(*Node)
		user := *node.Pointer
		user.Distance = int(node.Distance)
		user.Pin = 0

		matchedUsers = append(matchedUsers, &user)
	}

	return matchedUsers, nil
}

//returns true if the u and p are sexually compatible (compatible sex AND sexual interest)
func (u User) IsSexuallyCompatible(p *User) bool {
	if u.Sex == "F" {
		if u.Interest == "F" {
			return p.Sex == "F" && p.Interest == "F"
		}

		if u.Interest == "M" {
			return p.Sex == "M" && (p.Interest == "F" || p.Interest == "B")
		}

		if u.Interest == "B" {
			return p.Interest == "F" || p.Interest == "B"
		}
	}

	if u.Sex == "M" {
		if u.Interest == "M" {
			return p.Sex == "M" && p.Interest == "M"
		}

		if u.Interest == "F" {
			return p.Sex == "F" && (p.Interest == "M" || p.Interest == "B")
		}

		if u.Interest == "B" {
			return p.Interest == "M" || p.Interest == "B"
		}
	}

	return false
}

//age logic will be expounded upon if time permits. basically you shouldn't be wayyy old or wayy younger than the person
//also you get a slight advantage if you're 21 and a guy, cause then they're all like "look at you! you're so much more mature than normal guys :)"
func (p *User) AgeMatchup(q *User) float64 {
	if !p.legal || !q.legal { //no one under 18 allowed in the pool!
		return -10
	}

	twentyOneAdvantage := 0.

	if p.twentyOne && !q.twentyOne {
		twentyOneAdvantage = .5
	}

	both21 := p.twentyOne && q.twentyOne
	onlyP21 := p.twentyOne && !q.twentyOne

	diff := p.Age - q.Age
	adiff := math.Abs(float64(diff))

	if p.Sex == "F" {
		if q.Sex == "F" { //FF Homo
			if diff > 5 { //too old
				return -1
			}

			if diff > 2 { //slightly older, there's stuff to teach
				return .5
			}

			if diff >= -2 && diff <= 2 { //same age range
				return 1
			}

			if diff < -2 { //too young,
				return -1
			}
		}
		if q.Sex == "M" { //FM Hetero
			if onlyP21 {
				if diff > 6 { //too predatory, guy's shouldn't be 6+ years older than a girl who isn't 21
					return -2
				}

				if diff >= 4 {
					return -1
				}

				if diff >= 3 { //m is 21, w isn't, mAge < wAge +4; this is an ideal
					return -1
				}

				return -.5 //girl is older than guy by 1-2 years, and she's 21 or 22
			}

			if both21 {
				if diff >= 6 { //g is 6+ years older... too old!
					return -2
				}
				if diff < 6 && diff >= 1 { //g is 1-5 years older, both 21
					return -1
				}
				if diff == 0 { //same age
					return .5
				}
				if diff < 0 && adiff <= 3 { //girl is young, both 21+
					return .5
				}
				if diff < 0 && adiff > 3 { //girl is a bit younger, both 21+
					return 0
				}
				if diff < 0 && adiff > 5 { //girl is 21+ as is guy, guy is 5+ years older
					return -2
				}
			}

			return -2 //guy's too young... girl is 21 and guy is younger than that. too immature
		}
	}
	if p.Sex == "M" { //MF Hetero
		if q.Sex == "F" {
			if onlyP21 {
				if diff >= 6 { //too predatory, guy's shouldn't be 6+ years older than a girl who isn't 21
					return -2
				}

				if diff == 5 {
					return -1
				}

				if diff <= 4 { //m is 21, w isn't, mAge < wAge +4; this is an ideal
					return .5
				}
			}

			if both21 {
				if diff >= 6 { //m is 6+ years older... too old!
					return -2
				}
				if diff < 6 && diff >= 1 {
					return .5
				}
				if diff == 0 {
					return 0
				}
				if diff < 0 && adiff < 2 { //guy is young
					return -1
				}
				if diff < 0 && adiff >= 2 {
					return -2
				}
			}
			return -2 //guy's too young... girl is 21 and guy is younger than that. too immature
		}
		if q.Sex == "M" { //MM Homosexual pairing
			if diff >= 10 { //too old, predatory
				return -2
			}
			if diff >= 7 { //still a little predatory
				return -1
			}
			if diff >= 2 { //this works
				return .5 + twentyOneAdvantage
			}

			if diff >= -2 { //around same age
				return 0
			}

			if diff < -2 { //too young
				return -1
			}
		}
	}

	return 0
}

//HeightMatchup logic will be expounded upon if time allows later
func (p *User) HeightMatchup(q *User) float64 {
	diff := p.InchHeight() - q.InchHeight() //overall height difference between p & q

	if p.Sex == "F" {
		if q.Sex == "F" { //FF Homo
			return .5 //women are less superficial relative to each other than they are to men; height doesn't matter much here
		}
		if q.Sex == "M" { //FM Hetero
			if diff <= -4 || q.InchHeight() >= 72 { //guy is 4+ inches OR 6'
				return .5
			}
			if diff <= -2 { //guy is 2-3 inches taller
				return 0
			}
			if diff <= 0 { //guy is 0-1 inches taller
				return -.25
			}

			return -1 //if male is shorter than woman, take off 1 point
		}
	}
	if p.Sex == "M" { //MF Hetero
		if q.Sex == "F" {
			if diff >= 4 || p.InchHeight() >= 72 { //guy is 4+ inches OR 6'
				return .5
			}
			if diff >= 2 { //guy is 2-3 inches taller
				return 0
			}
			if diff >= 0 { //guy is 0-1 inches taller
				return -.25
			}

			return -1 //if male is shorter than woman, take off 1 point
		}
		if q.Sex == "M" { //MM Homo
			if diff >= 4 || p.InchHeight() >= 72 { //4+ inches OR 6'
				return 1
			}
			if diff >= 2 { //M1 is 2-3 inches taller than M2
				return .5
			}
			if diff >= -1 { //M1 & M2 are roughly equal height
				return 0
			}
			if diff >= -3 { //M1 is 2-3 inches shorter than M2
				return -.5
			}
			return -2
		}
	}

	return 0
}

//returns distance between p & q in kilometers
func (p *User) CalcDistance(q *User) float64 {

	dist := 0.5 - math.Cos((q.lat-p.lat)*pi)/2 + math.Cos(p.lat*pi)*math.Cos(q.lat*pi)*(1-math.Cos((q.long-p.long)*pi))/2
	dist = 12742 * math.Asin(math.Sqrt(dist))

	return dist
}
