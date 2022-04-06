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

if [ -d ~/Bughunt]
then
    echo " "
else
    mkdir ~/Bughunt
fi


if [ -d ~/Bughunt/$DOM ]
then
    echo " "
else
    mkdir ~/Bughunt/$DOM
fi

if [ -d ~/Bughunt/$DOM/XSS ]
then
    echo " "
else
mkdir ~/Bughunt/$DOM/XSS
fi

echo "${cyan} [+] Fetching URLS ${reset}"
echo " "

#gau
echo "${yellow}--------------------------xxxxxxx------------------------${reset}"

echo " "
if [ -f /usr/bin/gau ]
then 
    echo "${magenta} [+] Running Gau for retriving URLs ${reset}"
    gau --fc 404 --subs $DOM >> ~/Bughunt/$DOM/XSS/${DOM}-urls.txt
else
    echo "${cyan} [+] Installing Gau ${reset}"
    go g
    echo "${magenta} [+] Running Gau for retriving URLs ${reset}"
    gau --fc 404 --subs $DOM >> ~/Bughunt/$DOM/XSS/${DOM}-urls.txt
fi
echo " "
echo "${white} [+] Successfully saved as ${DOM}-url.txt ${reset}"
echo " "