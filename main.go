package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)
// Convert JSON to Go Data Structures:
// https://mholt.github.io/json-to-go/

// Type art is not exportable to other packages because naming is lowercased. If export is needed, uppercase custom type name.
// Fields are uppercase which allow Marshal and Unmarshal access from the encoding/json pkg.

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

func unMarshalJSON(bs []byte, a *artist) {
	//use type pointer
	err := json.Unmarshal(bs, a)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// Generate sample data to Marshal into JSON.
	artOne := art{Name: "paintingOne", Medium: "paint", Price: 1000}
	artTwo := art{Name: "paintingTwo", Medium: "paint", Price: 1000}
	artThree := art{Name: "paintingThree", Medium: "paint", Price: 1000}
	artistOne := artist{Name: "will", Age: 29, Art: []art{artOne, artTwo, artThree}}

	// Call function and pass in artistOne variable which returns []bytes and is then converted into
	// A string inside of the print function on line 56 : string([]byte)

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

	 unMarshalJSON(jsonOne, &singleArtist)
	 fmt.Println(singleArtist)

	 // Or do inline bellow -->

	// err := json.Unmarshal(jsonOne, &singleArtist)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("all of the data", singleArtist)

	s := `mypassword`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("password before encryption:", s)
	fmt.Println("password after encryption:", bs)
	fmt.Printf("%T: %v\n", bs, bs)

	loginPassword := `mypassword`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword))
	if err != nil {
		fmt.Println("error, You Cant Login, Did You Forget who you are?..Try again please, or dont.", err)
		return
	}
	fmt.Println("your logged in")
}

// Links to code snippets on GoPlayground
// https://play.golang.org/p/ys3y3xzYLdL
// https://play.golang.org/p/rLDwNiILa2A
// SORT ints and strings: https://play.golang.org/p/at0dKQXleEh
// https://play.golang.org/p/dJUaNK03MCE
// 
// 