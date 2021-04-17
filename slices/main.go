package main

import (
	"fmt"
	s "github.com/inancgumus/prettyslice"
)

func main() {

	items := []string{"pacman", "mario", "tetris", "doom",
		"galaga", "frogger", "asteriods", "simcity",
		"metroid", "defender", "rayman", "tempest",
		"ultima"}

	const pageSize = 4

	l := len(items)
	for from := 0; from < l; from += pageSize {
		to := from + pageSize
		if to > l {
			to = l
		}
		fmt.Printf("%d:%d\n", from, to)
		currentPage := items[from:to]

		head := fmt.Sprintf("Page #%d", (from / pageSize) +1 )
		s.Show(head, currentPage)
	}

	/*s.MaxPerLine = 4
	s.Show("items", items)
	top3 := items[:3]
	s.Show("top 3 items", top3)

	l := len(items)

	last4 := items[l-4:]
	s.Show("last 4 items", last4)*/

	/*var todo []string
	todo = append(todo, "sing", "run", "code", "play")
	tomorrow := []string{"see mom","learn go"}
	todo = append(todo,tomorrow...)

	s.Show("todo", todo)*/

	/*rand.Seed(time.Now().UnixNano())
		max, _ := strconv.Atoi(os.Args[1])

		var uniques []int
	loop:
		for len(uniques) < max {
			n := rand.Intn(max) + 1

			fmt.Print(n, " ")

			for _, u := range uniques {
				if u == n {
					continue loop
				}
			}
			//uniques[found] = n
			uniques = append(uniques, n)

		}
		fmt.Println("\n\nuniques", uniques)
		sort.Ints(uniques)
		fmt.Println("\nsorted:", uniques)

		nums := [5]int{5, 4, 3, 2, 1}
		sort.Ints(nums[:]) // arrays can be used as slice like this
		fmt.Println("\nnums:", nums)*/

	/*games := []string{"kokemon", "sims"}
	newGames := []string{"pacman", "doom", "pong"}
	newGames = games
	var ok string
	for i, game := range games {
		if game != newGames[i] {
			ok = "not "
			break
		}
	}

	fmt.Printf("games and newGames are %sequal",ok)*/

	/*var books [5]string

	books[0] = "dracula"
	books[1] = "1984"
	books[2] = "island"

	games := []string{"kokemon", "sims"}

	fmt.Printf("books : %#v\n", books)
	fmt.Printf("books : %#v\n", games)
	fmt.Printf("books : %T\n", games)
	fmt.Printf("books : %d\n", len(games))
	fmt.Printf("nil? : %t\n", games == nil)*/

}
