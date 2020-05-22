from weather import Weather
import os
import datetime
from newsapi import NewsApiClient
import time

def print_paper():
    api = NewsApiClient(api_key='bbd8b254ae1a4ae1abee5c695f63fb43')
    weather = Weather()
    lookup = weather.lookup(12770747)
    condition = lookup.condition()
    clear = [31,32]
    fair = [33,34]
    partly_cloudy = [44]
    cloudy = range(19,31)
    rain = range(0,19)+range(35,44)+[45,46,47]
    conditions = [clear,fair,partly_cloudy,cloudy,rain]
    condition_files = ["clear.png","fair.png","partly_cloudy.png","cloudy.png","rain.png"]

    def image_lookup(cond_code):
        cond_code = int(cond_code)
        for i in range(0,len(conditions)):
            cond = conditions[i]
            if cond_code in cond:
                return condition_files[i]

    articles = api.get_top_headlines(sources='bbc-news')["articles"]
    printable = []
    temp = min(len(articles),5)
    for i in range(temp):
        printable.append(articles[i]["title"]+"\n")

    print_file = open("pf.txt","w")
    print_file.write(condition.text()+" "+condition.temp()+"C\n")
    print_file.write("Good Morning!\n")
    print_file.write(datetime.datetime.now().strftime("%H:%M %m/%d/%y")+"\n")
    print_file.write("\nHere is your news for today!\n")
    for line in printable:
        print_file.write("-----------------\n")
        linemod = ""
        while(len(line)>18):
            linemod = linemod+line[:18]+"\n"
            line = line[18:]
        linemod+=line
        print_file.write(linemod)
    print_file.close()

    #os.system('lpr -o scaling=50 '+image_lookup(condition.code()))
    os.system('lpr pf.txt')



hour = datetime.datetime.now().hour
minute = datetime.datetime.now().minute
hour_alarm = 9
minute_alarm = 0

while((not hour_alarm==hour) or (not minute_alarm==minute)):
    t_minutes = hour_alarm*60+minute_alarm-hour*60-minute
    print(t_minutes," Minutes Until Alarm")
    time.sleep(5)
    hour = datetime.datetime.now().hour
    minute = datetime.datetime.now().minute

print_paper()