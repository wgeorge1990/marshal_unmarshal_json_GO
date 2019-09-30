package main

import (
	"encoding/json"
	"fmt"
)

// Type art is not exportable to other packages because naming is lowercased. If export is needed, upercase custom type name.
// Fields are uppercase which allow Marshal and Unmarshal accesss from the encoding/json pkg.

// type artist struct {
// 	Name string `json:"Name"`
// 	Age  int    `json:"Age"`
// 	Art  []struct {
// 		Name   string `json:"Name"`
// 		Medium string `json:"Medium"`
// 		Price  int    `json:"Price"`
// 	} `json:"Art"`
// }

//------------OR--------------//

type art struct {
	Name   string `json:"Name"`
	Medium string `json:"Medium"`
	Price  int    `json:"Price"`
}

type artist struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
	Art  []art  `json:"Art"`
}

func convertToJSON(v interface{}) []byte {
	bs, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
	}
	return bs
}

func unMarshalJSON(bs []byte, a artist{}) {
	err := json.Unmarshal(bs, a)
}

func main() {
	//Generate sample data to Marshal into JSON.
	artOne := art{Name: "paintingOne", Medium: "paint", Price: 1000}
	artTwo := art{Name: "paintingTwo", Medium: "paint", Price: 1000}
	artThree := art{Name: "paintingThree", Medium: "paint", Price: 1000}
	artistOne := artist{Name: "will", Age: 29, Art: []art{artOne, artTwo, artThree}}
	//call function and pass in artistOne variable which returns []bytes and is then converted into
	// a string inside of the print function on line 56 : string([]byte)
	jsonOne := convertToJSON(artistOne)
	// fmt.Println(string(jsonOne))
	//Print type []uint8 which is same as []byte
	// fmt.Printf("%T\n", jsonOne)
	//Print value int bytes
	//fmt.Printf("%v\n", json)

	// var singleArtist artist
	//         OR
	//______name________type__
	 singleArtist := artist{}
	err := json.Unmarshal(jsonOne, &singleArtist)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("all of the data", singleArtist)


}
