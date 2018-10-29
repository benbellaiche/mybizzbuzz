package main

//import requeried packages
import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"
)

// Contains the JSON to send
type Response struct {
  Values []string
}

// contain query args
type Query struct {
    String1 string
    String2 string
    Int1 int
    Int2 int
    Limit int

}

//Let to generate set of values with numbers and parameters String
func generate(userQuery *Query)([]string) {
    tab := []string{}

    for i := 1; i <= userQuery.Limit; i++ {
        if (i%userQuery.Int1 == 0) && (i%userQuery.Int2 == 0){
            tab = append(tab, userQuery.String1+userQuery.String2)
        } else if (i%userQuery.Int1 == 0){
            tab = append(tab, userQuery.String1)
        } else if (i%userQuery.Int2 == 0){
            tab = append(tab, userQuery.String2)
        } else{
            tab = append(tab, strconv.Itoa(i))
        }
    }

    return tab
}

// Entrypoint for API REST
func main() {
    http.HandleFunc("/", handler)

    error := http.ListenAndServe(":8080", nil)

    if error != nil {
        log.Fatal("ListenAndServe ", error)
    }
}

// Function call on client request
func handler(w http.ResponseWriter, r *http.Request) {

    //Parsing args
    string1, ok1 := r.URL.Query()["string1"]
    string2, ok2 := r.URL.Query()["string2"]
    int1, ok3 := r.URL.Query()["int1"]
    int2, ok4 := r.URL.Query()["int2"]
    limit, ok5 := r.URL.Query()["limit"]

    //Check for requeried fields
    if !ok1 || len(string1[0]) < 1 {
        log.Println("Url Param 'string1' is missing")
        return
    }

    if !ok2 || len(string2[0]) < 1 {
        log.Println("Url Param 'string2' is missing")
        return
    }

    if !ok3 || len(int1[0]) < 1 {
        log.Println("Url Param 'int1' is missing")
        return
    }

    if !ok4 || len(int2[0]) < 1 {
        log.Println("Url Param 'int2' is missing")
        return
    }

    if !ok5 || len(limit[0]) < 1 {
        log.Println("Url Param 'limit' is missing")
        return
    }

    //Convert to int
    i1, err1 := strconv.Atoi(int1[0])
    i2, err2 := strconv.Atoi(int2[0])
    l, err3 := strconv.Atoi(limit[0])


    if err1 != nil {
        log.Fatal("Field int1 is not integer")
        return
    }

    if err2 != nil {
        log.Fatal("Field int2 is not integer")
        return
    }

    if err3 != nil {
        log.Fatal("Field limit is not integer")
        return
    }

    if (i1 < 1) {
        log.Fatal("Field int1 has to be over 0")
        return
    }

    if (i2 < 1) {
        log.Fatal("Field int2 has to be over 0")
        return
    }

    if (l < 1) {
        log.Fatal("Field limit has to be over 0")
        return
    }

    //Create new struct to store args values
    userquery := &Query{}

    //store values
    userquery.String1 = string1[0]
    userquery.String2 = string2[0]
    userquery.Int1 = i1
    userquery.Int2 = i2
    userquery.Limit = l

    //Generate fizzbuzz values
    tab := generate(userquery)

    //Store fizzbuzz values in new structure
    profile := Response{tab}

    //Generate json
    js, err := json.Marshal(profile)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    //Response to client
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}
