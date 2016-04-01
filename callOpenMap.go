package main

import(
  "os"
  "log"
  "fmt"
  "strconv"
  owm "github.com/briandowns/openweathermap"
)


func parser(strang string) {
  fmt.Println(strang)
}

func main()  {
  args:=os.Args[1:]
  var s string = args[0]
  k, _ := strconv.Atoi(s)
  // fmt.Println(err)

  os.Setenv("OWM_API_KEY", "0a09d24cb9d40fa80c05d5011f6eca28")
  w, err := owm.NewCurrent("C", "en")
  if err != nil {
    log.Fatalln(err)
  }
  w.CurrentByID(k)
  // fmt.Println(reflect.TypeOf(w))
  fmt.Println(w)
  // var m []M
  // reading, err := json.Unmarshal(w, &m)
  // fmt.Println(reflect.TypeOf(reading))
}
