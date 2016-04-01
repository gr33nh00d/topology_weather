package main
import "fmt"
import "regexp"
import "encoding/json"
// import "reflect"
import "strconv"
import "io/ioutil"

type datum struct{
  CityID int
  Latitude float64
  Longitude float64
  Sunrise int
  Sunset int
  Temperature int
  Pressure int
  Humidity int
  WindSpeed float64
  WindDir int
}

func main() {

  dat, _:= ioutil.ReadFile("output.txt")
  var strang string = string(dat)
  length:=len(strang)
  var start int = 0
  var end int = 0
  var chr string="o"

  //removing the cmd clouds thing bc 8c2 and 23d will return different number of numbers
  for i := 0; i < length; i++{
    chr=strang[i:i+1]
    if chr=="[" {
      start=i
      i=length+1
    }
  }
  for i := 0; i < length; i++{
    chr=strang[i:i+1]
    if chr=="]"{
      end=i
      i=length
    }
  }
  strang=strang[0:start]+strang[end:]

//0,1,7,10,13,14,15
  re := regexp.MustCompile("[0-9,.,-]+")
  numbers := re.FindAllString(strang, -1)

  lon,_:=strconv.ParseFloat(numbers[0],64)
  lat,_:=strconv.ParseFloat(numbers[1],64)
  sunrise,_:=strconv.Atoi(numbers[5])
  sunset,_:=strconv.Atoi(numbers[6])
  temp,_:=strconv.Atoi(numbers[7])
  pressure,_:=strconv.Atoi(numbers[10])
  humidity,_:=strconv.Atoi(numbers[13])
  windSpeed,_:=strconv.ParseFloat(numbers[14],64)
  windDir,_:=strconv.Atoi(numbers[15])
  id,_:=strconv.Atoi(numbers[18])

  foo := &datum{
    CityID: id,
    Latitude: lat,
    Longitude: lon,
    Sunrise: sunrise,
    Sunset: sunset,
    Temperature: temp,
    Pressure: pressure,
    Humidity: humidity,
    WindSpeed: windSpeed,
    WindDir: windDir}
  bar,_:=json.Marshal(foo)
  fmt.Println(string(bar))
}
