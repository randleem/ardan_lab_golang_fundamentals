package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main(){
fmt.Println(UserInfo("ardanlabs"))
	
}


// UserInfo reutrns name and number of public repose from Githu api
func UserInfo(login string)(string, int, error){
url:= "https://api.github.com/users/"+login
	resp, err := http.Get(url)
	if err != nil{
		fmt.Println("Error: ", err)
		return "", 0, err
	}
	if resp.StatusCode != http.StatusOK{
		fmt.Printf("Error: bad status -%s\n", resp.Status)
		return "", 0, fmt.Errorf("%q - bad status: %s", url, resp.Status)
	}

	return parseResponse(resp.Body)
}

func parseResponse(r io.Reader)(string, int, error){
//io.Copy(os.Stdout, resp.Body)
	var reply struct{
		Name string

		NumRepos int `json:"public_repos"`
	}
	dec := json.NewDecoder(r)
	if err := dec.Decode(&reply); err!= nil{
		fmt.Print("Error: ", err)
		return "", 0, err
	}
	return reply.Name, reply.NumRepos, nil
}


/* JSON <-> Go

TYPES
string <-> string
true/false <-> bool
number <-> float64(default), 32, int, int8.... uint64
array <-> []T, []any
object <-> map[string]any, struct

encoding/json API
JSON -> []byte -> Go: Unmarshal
Go -> []byte -> JSON: Marshal
JSON -> io.Reader -> Go: Decoder
Go  -> io.Writer -> JSON: Encoder
*/