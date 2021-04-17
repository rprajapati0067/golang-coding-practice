package main

import "fmt"

func mapPractice() {
	dict := map[string]string{"101": "ravi", "102": "amit"}
	mapl := make(map[string]string)
	mapl["ravi"] = "101"
	for k, v := range dict {
		fmt.Printf("key : %s value: %s ", k, v)
	}
//	fmt.Println(dict)
}

func main() {

	mapPractice()
	/*args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[english word] -> [turkish word]")
		return
	}
	query := args[0]

	dist := map[string]string{"good": "kotu", "great": "harika", "perfect": "mukemmel"}

	//dist["up"] = "yukari"
	//dist["down"] = "asagi"

	value := dist[query]

	fmt.Printf("%q Zero value: %#v\n", query, value)
	fmt.Printf("# of key: %d\n", len(dist))*/
}
