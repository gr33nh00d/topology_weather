import json
from collections import OrderedDict

in_file=open("temp.txt", "r")
out_file=open("data.txt", "a")
text = in_file.read()
#out_file.write(text)
parsed = json.loads(text)
current=parsed['current_observation']

cityInfo = current['display_location']

city = cityInfo['city']
latitude = cityInfo['latitude']
longitude = cityInfo['longitude']
elevation = cityInfo['elevation']
temp = current['temp_c']
humidity = current['relative_humidity']
windDir = current['wind_degrees']
windSpeed = current['wind_kph']
pressure = current['pressure_in']
pressureTrend = current['pressure_trend']
time = current['observation_time_rfc822']


output = json.dumps(OrderedDict([("city", city), ("time", time), ("latitude", latitude), ("longitude", longitude), ("elevation", elevation), ("temp", temp), ("humidity", humidity), ("Wind_Direction",windDir), ("windSpeed",windSpeed), ("pressure",pressure), ("Presure_Trend",pressureTrend)]))

out_file.write(output+",\n")
