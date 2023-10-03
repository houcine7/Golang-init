package main

import "fmt"

func main() {
	// we can declare and init a map
	var mp map[string]uint32 = map[string]uint32{"Houcine": 45, "Jhon": 65}
	fmt.Println(mp)

	// we can use make
	var mpMake = make(map[string]int32)
	mpMake["Key"] = 654
	mpMake["Key2"] = 654
	fmt.Println(mpMake)

	// we can get the value of a key
	fmt.Println(mp["Houcine"])
	// but if the key doesn't exist we get the default value of the values
	fmt.Println(mp["Houcine2"]) // 0 the default value of uint32

	// we can check if a key exists by
	var val, exist = mp["Houcine4"]
	if exist {
		fmt.Println(val)
	} else {
		fmt.Println("key doesn't exist")
	}
	// delete a element in the map
	delete(mp, "Houcine564")
	fmt.Println(mp)

	// to iterate over an array
	for key, value := range mp {
		fmt.Println(key, value)
	}

}

