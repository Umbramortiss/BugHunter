#!/bin/bash

#colors
red=`tput setaf 1`
green=`tput setaf 2`
yellow=`tput setaf 3`
cyan=`tput setaf 6`
magenta=`tput setaf 5`
white=`tput setaf 7`
reset=`tput sgr0`


read -p "Enter The Domain Name: " Domain

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

if [ -d ~/Bughunt/$DOM/Subdomains ]
then
    echo " "
else
mkdir ~/Bughunt/$DOM/Subdomains
fi

echo "${cyan} [+] Fetching URLS ${reset}"