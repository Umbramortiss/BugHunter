#!/bin/bash

#colors
red=`tput setaf 1`
green=`tput setaf 2`
yellow=`tput setaf 3`
cyan=`tput setaf 6`
magenta=`tput setaf 5`
white=`tput setaf 7`
reset=`tput sgr0`


read -p "Enter The Domain Name: " DOM

if [ -d ~/Bughunt/Bugxss ]
then
    echo " "
else
    mkdir ~/Bughunt/Bugxss
fi


if [ -d ~/Bughunt/Bugxss/$DOM ]
then
    echo " "
else
    mkdir ~/Bughunt/Bugxss/$DOM
fi

if [ -d ~/Bughunt/Bugxss/$DOM/Xss ]
then
    echo " "
else
mkdir ~/Bughunt/Bugxss/$DOM/Xss
fi

echo "${cyan} [+] Fetching URLS ${reset}"
echo " "


#gau
echo "${yellow}------------------------*********-----------------------${reset}"

echo " "
if [ -f /usr/bin/gau ]
then 
    echo "${magenta} [+] Running Gau for retriving URLs ${reset}"
    gau --fc 404 --subs $DOM > ~/Bughunt/Bugxss/$DOM/Xss/${DOM}-urls.txt
else
    echo "${cyan} [+] Installing Gau ${reset}"
    go get -u github.com/lc/gau/v2/cmd/gau
    echo "${magenta} [+] Running Gau for retriving URLs ${reset}"
    gau --fc 404 --subs $DOM > ~/Bughunt/Bugxss/$DOM/Xss/${DOM}-urls.txt
fi
echo " "
echo "${white} [+] Successfully saved as ${DOM}-url.txt ${reset}"
echo " "


#httpx
echo "${yellow}-------------------------*********---------------------${reset}"

echo " "
if [ -f /usr/bin/local/httpx ]
then
    echo "{magenta} [+] Running httpx for sorting alive urls"
    cat ~/Bughunt/Bugxss/$DOM/Xss/${DOM}-urls.txt  | httpx -silent >> ~/Bughunt/Bugxss/$DOM/Xss/${DOM}-alive-urls.txt
else
    echo "${cyan} [+] Installing Httpx ${reset}"
    go get -u github.com/projectdiscovery/httpx/cmd/httpx
    echo "{magenta} [+] Running httpx for sorting alive urls"
    cat ~/Bughunt/Bugxss/$DOM/Xss/${DOM}-urls.txt  | httpx  -silent >> ~/Bughunt/Bugxss/$DOM/Xss/${DOM}-alive-urls.txt
fi
echo " "
echo "${white} [+] Successfully Saved As ${DOM}-alive-urls.txt"
echo " "

#Gxss
echo "${yellow}------------------------*********----------------------${reset}"

echo " "
if [ -f /usr/bin/Gxss ]
then 
    echo "${magenta} [+] Running Gxss"
    cat ~/Bughunt/Bugxss/$DOM/Xss/${DOM}-alive-urls.txt | Gxss -c 100 | sed 's/=.*/=/'| sed 's/URL: //' > ~/Bughunt/Bugxss/$DOM/Xss/${DOM}-reflecparm.txt

else 
    echo "${cyan} [+] Installing Gxss"
    go get -u github.com/KathanP19/Gxss
    echo "${magenta} [+] Running Gxss for Refleced Parameters"
    cat ~/Bughunt/Bugxss/$DOM/Xss/${DOM}-alive-urls.txt | Gxss -c 100 | sed 's/=.*/=/'| sed 's/URL: //' | sort -u > ~/Bughunt/Bugxss/$DOM/Xss/${DOM}-reflecparm.txt
fi
echo " "
echo "${white} [+] Successfully Saved As ${DOM}-reflecparm.txt ${reset}"
echo " "


#Dalfox
echo "${yellow}-----------------------*********-------------------${reset}"

echo " "
if [ -f /usr/bin/dalfox ]
then
    echo "${magenta} [+] Dalfox scanning for Xss"
    cat ~/Bughunt/Bugxss/$DOM/Xss/${DOM}-reflecparm.txt | dalfox pipe > ~/Bughunt/Bugxss/Xss/${DOM}-xssScan.txt
    
else
    echo "${cyan} [+] Installing Dalfox"
    go get -u github.com/hahwul/dalfox
    echo "${magenta} [+] Dalfox scanning for Xss"
    cat ~/Bughunt/Bugxss/$DOM/Xss/${DOM}-reflecparm.txt | dalfox pipe > ~/Bughunt/Bugxss/Xss/${DOM}-xssScan.txt

fi
echo " "
echo "${cyan} [+] Successfully Saved Results As ${DOM}-xssScan.txt"
echo " "