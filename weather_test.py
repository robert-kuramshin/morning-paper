from weather import Weather
import os
import datetime

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

print_file = open("pf.txt","w")
print_file.write("Good Morning!\n")
print_file.write("It is "+datetime.datetime.now().strftime("%H:%M %m/%d/%y")+"\n")
print_file.write(condition.text()+"\n"+condition.temp()+"C")
print_file.close()

os.system('lpr -o fit-to-page '+image_lookup(condition.code()))
os.system('lpr pf.txt')

