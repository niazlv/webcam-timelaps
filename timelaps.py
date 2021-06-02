#author github.com/niazlv

import os
import time
from datetime import datetime

from config import *

url="http://"+login+":"+"passwd"+"@"+ip+"/axis-cgi/jpg/image.cgi?"+resolution+"&"+compression

urls="'"
urls+=url
urls+="'"

def log(str):
    print(str)

def cam(filename='cam.jpg'):
    os.system('wget '+urls+' -q -O '+filename)
    #os.system('fswebcam -q -r 640x480 --no-banner --save '+filename)
    log('cam1: '+filename)

while True:
    now=datetime.now()
    todayDIR=now.strftime("%m-%d-%Y")
    todays=now.strftime("%H-%M-%S")
    name=todayDIR+"/cam1_"+todays+".jpg"
    #print("today's date:",today)

    if os.path.isdir(todayDIR):
        #log("найдено")
        #name=todayDIR+"/cam1_"+todays+".jpg"
        cam(name)

    else:
    	os.system("mkdir "+todayDIR)
    	log("dir "+todayDIR+" not found, and now created")
    time.sleep(delay)
