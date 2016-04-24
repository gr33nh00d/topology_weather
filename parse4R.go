package main
import "fmt"
import "regexp"
// import "encoding/json"
// import "reflect"
import "strconv"
// import "io/ioutil"
import "os"
import "bufio"
import "time"
// import "log"

func main() {

  // args:=os.Args[1:]
  // var s string = args[0]
  // k, _ := strconv.Atoi(s)

  t:=time.Now()
  var path string = "/home/data/"+t.String()
  err:=os.MkdirAll(path,0711)
  if err !=nil{
    fmt.Println("error")
    return
  }

  file, _:= os.Open("output.txt")
  defer file.Close()
  scanner := bufio.NewScanner(file)
  var m int = 1

  fileSunrise, _ := os.Create(path+"/sunrise.txt")
  sunrise := bufio.NewWriter(fileSunrise)
  defer fileSunrise.Close()

  fileSunset, _ := os.Create(path+"/sunset.txt")
  sunset := bufio.NewWriter(fileSunset)
  defer fileSunset.Close()

  fileTemp, _ := os.Create(path+"/temp.txt")
  temp := bufio.NewWriter(fileTemp)
  defer fileTemp.Close()

  filePressure, _ := os.Create(path+"/pressure.txt")
  pressure := bufio.NewWriter(filePressure)
  defer filePressure.Close()

  fileHumidity, _ := os.Create(path+"/humidity.txt")
  humidity := bufio.NewWriter(fileHumidity)
  defer fileHumidity.Close()

  fileWindSpeed, _ := os.Create(path+"/windSpeed.txt")
  windSpeed := bufio.NewWriter(fileWindSpeed)
  defer fileWindSpeed.Close()

  fileAll, _ := os.Create(path+"/all.txt")
  all := bufio.NewWriter(fileAll)
  defer fileAll.Close()

  sunrise.WriteString("\"\",\"lat\",\"lon\",\"sunrise\"\n")
  sunset.WriteString("\"\",\"lat\",\"lon\",\"sunset\"\n")
  pressure.WriteString("\"\",\"lat\",\"lon\",\"pressure\"\n")
  humidity.WriteString("\"\",\"lat\",\"lon\",\"humidity\"\n")
  temp.WriteString("\"\",\"lat\",\"lon\",\"temp\"\n")
  windSpeed.WriteString("\"\",\"lat\",\"lon\",\"windSpeed\"\n")
  all.WriteString("\"\",\"lat\",\"lon\",\"sunrise\",\"sunset\",\"temp\",\"pressure\",\"humidity\",\"windSpeed\"\n")

  for scanner.Scan(){
    var strang string = scanner.Text()

    length:=len(strang)
    var start int = 0
    var end int = 0
    var chr string="o"
    // fmt.Println(strang)

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

    re := regexp.MustCompile("[0-9,.,-]+")
    numbers := re.FindAllString(strang, -1)
    var inc string = strconv.Itoa(m)
    m++
    //sunrise
    fmt.Println("["+numbers[1]+", "+numbers[0]+", "+numbers[5]+"],")
    sunrise.WriteString("\""+inc+"\""+","+numbers[1]+","+numbers[0]+","+numbers[5]+"\n")
    //sunset
    fmt.Println("["+numbers[1]+", "+numbers[0]+", "+numbers[6]+"],")
    sunset.WriteString("\""+inc+"\""+","+numbers[1]+","+numbers[0]+","+numbers[6]+"\n")
    //Temperature
    fmt.Println("["+numbers[1]+", "+numbers[0]+", "+numbers[7]+"],")
    temp.WriteString("\""+inc+"\""+","+numbers[1]+","+numbers[0]+","+numbers[7]+"\n")
    //pressure
    fmt.Println("["+numbers[1]+", "+numbers[0]+", "+numbers[10]+"],")
    pressure.WriteString("\""+inc+"\""+","+numbers[1]+","+numbers[0]+","+numbers[10]+"\n")
    //humidity
    fmt.Println("["+numbers[1]+", "+numbers[0]+", "+numbers[13]+"],")
    humidity.WriteString("\""+inc+"\""+","+numbers[1]+","+numbers[0]+","+numbers[13]+"\n")
    //windSpeed
    fmt.Println("["+numbers[1]+", "+numbers[0]+", "+numbers[15]+"],")
    windSpeed.WriteString("\""+inc+"\""+","+numbers[1]+","+numbers[0]+","+numbers[15]+"\n")
    all.WriteString("\""+inc+"\""+","+numbers[1]+","+numbers[0]+","+numbers[5]+","+numbers[6]+","+numbers[7]+","+numbers[10]+","+numbers[13]+","+numbers[15]+"\n")
  }

  all.Flush()
  humidity.Flush()
  pressure.Flush()
  temp.Flush()
  sunset.Flush()
  sunrise.Flush()
  windSpeed.Flush()
}

// lon,_:=strconv.ParseFloat(numbers[0],64)
// lat,_:=strconv.ParseFloat(numbers[1],64)
// sunrise,_:=strconv.Atoi(numbers[5])
// sunset,_:=strconv.Atoi(numbers[6])
// temp,_:=strconv.Atoi(numbers[7])
// pressure,_:=strconv.Atoi(numbers[10])
// humidity,_:=strconv.Atoi(numbers[13])
// windSpeed,_:=strconv.ParseFloat(numbers[14],64)
// windDir,_:=strconv.Atoi(numbers[15])
// id,_:=strconv.Atoi(numbers[18])
