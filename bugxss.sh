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
    gau --fc 404 --subs $DOM >> ~/Bughunt/Bugxss/$DOM/Xss/${DOM}-urls.txt
else
    echo "${cyan} [+] Installing Gau ${reset}"
    go g
    echo "${magenta} [+] Running Gau for retriving URLs ${reset}"
    gau --fc 404 --subs $DOM >> ~/Bughunt/Bugxss/$DOM/Xss/${DOM}-urls.txt
fi
echo " "
echo "${white} [+] Successfully saved as ${DOM}-url.txt ${reset}"
echo " "

#Gxss
echo "${yellow}------------------------*********----------------------${reset}"

echo " "
if [ -f /usr/bin/Gxss ]
then 
    echo "${magenta} [+] Running Gxss"
    "cat" ~/Bughunt/Bugxss/$DOM/Xss/${DOM}-urls.txt | Gxss -c 100 -o 