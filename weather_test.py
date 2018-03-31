from weather import Weather
weather = Weather()
lookup = weather.lookup(12770747)

condition = lookup.condition()

print(condition.text())
print(condition.temp())

