package main

type Vector2 struct {
	x int
	y int
}

func nextPosition(curr Vector2, d rune) Vector2 {
	if d == '^' {
		curr.y -= 1
	} else if d == 'v' {
		curr.y += 1
	} else if d == '>' {
		curr.x += 1
	} else if d == '<' {
		curr.x -= 1
	} else {
		panic("Wrong direction")
	}
	return curr
}

func CountHouse(directions string) int {

	houseDict := make(map[(Vector2)]bool)

	curr := Vector2{}
	houseDict[curr] = true

	for _, d := range directions {
		curr = nextPosition(curr, d)
		houseDict[curr] = true
	}

	return len(houseDict)
}

func CountHouseRobo(directions string) int {

	houseDict := make(map[(Vector2)]bool)

	curr := Vector2{}
	currRobo := Vector2{}
	houseDict[curr] = true
	houseDict[currRobo] = true

	idx := 0
	list := []rune(directions)
	for idx < len(list){
		curr = nextPosition(curr, list[idx])
		houseDict[curr] = true
		idx += 1

		if idx < len(list) {
			currRobo = nextPosition(currRobo, list[idx])
			houseDict[currRobo] = true
			idx += 1
		}
	}

	return len(houseDict)
}