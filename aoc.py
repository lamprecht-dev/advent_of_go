import sys
import requests
import os
from datetime import date
import browser_cookie3
import shutil

year = 2025

### WRITEN BY antooro - https://github.com/antooro/advent-of-code-2019/blob/master/startDay.py

# Get cookies from the browser
# cj = browser_cookie3.chrome(cookie_file="/mnt/c/Users/david/AppData/Local/Google/Chrome/User Data/Profile 1/Network/Cookies")
cj = browser_cookie3.firefox(cookie_file="/mnt/c/Users/david/AppData/Roaming/Mozilla/Firefox/Profiles/get91m7f.default-release/cookies.sqlite")

# Get today number of day
day_today = date.today().strftime("%d").lstrip("0")

if len(sys.argv) > 1:
    day = int(sys.argv[1])
    if day < 0 or day > 31:
        exit("Day is not valid")
else:
    day = day_today

print(f"Initializing day {day}")

if not os.path.exists(f"day{day}"):
    shutil.copytree(f"{year}/template", f"{year}/day{day}")
    os.chdir(f"{year}/day{day}")

    r = requests.get(f"https://adventofcode.com/{year}/day/{day}/input", cookies=cj)

    with open(f"input.txt", "w") as f:
        f.write(r.text)
