package main


type Tests []*Test
 
type Test struct {
    PostID int    `json:"postId"`
    ID     int    `json:"id"`
    Name   string `json:"name"`
    Email  string `json:"email"`
    Body   string `json:"body"`
}
 
func (Test) TableName() string {
    return "test"
}
 
resp, err := http.Get("https://jsonplaceholder.typicode.com/comments?postId=1")
    if err != nil {
        log.Fatal(err)
    }
 
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }
    defer resp.Body.Close()
    tests:= make(Tests,0)
    err := json.Unmarshal(body, &tests )
    if err != nil {
        fmt.Println(err)
    }
    for test := tests {
         res = gormDb.Create(test)
         if res.Error != nil  {
            panic(res.Error)
          }
    }