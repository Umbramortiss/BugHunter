#!/bin/bash

#colors
red='tput setaf 1'
green='tput setaf 2'
yellow='tput setaf 3'
cyan='tput setaf 6'
magenta='tput setaf 5'
reset='tput sgr0'

read -p "Enter The Domain Name: " DOM 

if [ -d ~/Bugsubs ]
then
    echo " "
else
    mkdir ~/Bugsubs
fi 

if [ -d ~/Bugsubs/$DOM ]
then
    echo " "
else 
    mkdir ~/Bugsubs/$DOM
fi

if [ -d ~/Bugsubs/$DOM/Subdomains ]
then
    echo " "
else 
mkdir ~/Bugsubs/$DOM/Subdomains
fi

echo "${cyan} [+] Starting Subdomain Enumeration ${reset}"
echo " "

#assetfinder
echo "${yellow} ----------------------xxxxxxx-------------------------${reset}"

echo " "
if [ -f /usr/bin/assetfinder ]
then
    echo "${magenta} [+] Running Assetfinder for Subdomain Enuneration ${reset}"
    assetfinder -subs-only $DOM >> ~/Bugsubs/$DOM/Subdomains/assetfinder.txt
else
    echo "${cyan} [+] Installing Assetfinder ${reset}"
    go get -u github.com/tomnomnom/Assetfinder
    echo "${magenta} [+] Running Assetfinder for subdomain enumeration ${reset}"
    assetfinder -subs-only $DOM >> ~/Bugsubs/$DOM/Subdomains/assetfinder.txt
fi
echo " "
echo "${cyan} [+] Successfully saved as assetfinder.txt ${reset}"
echo " "

#amass
echo "${yellow}-------------------------xxxxxxx-------------------${reset}"
echo " "
if [ -f /usr/bin/amass ]
then
    echo "${magenta} [+] Running Amass for subdomain enumeration ${reset}"
    amass enum --passive -d $DOM > ~/Bugsubs/$DOM/Subdomains/amass.txt
else
    echo "${cyan} [+] Installing Amass ${reset}"
    echo "${cyan} [+] This may take few minutes hang tight... ${reset}"
    go get -u github.com/OWASP/Amass/...
    echo "${magenta} [+] Running Amass for subm' enumeration ${reset}"
    amass enum --passive -d $DOM > ~/Bugsubs/$DOM/Subdomains/amass.txt
fi
echo " "
echo "${cyan} [+] Successfully saved as amass.txt ${reset}"
echo " "

echo " "
if [ -f /usr/local/bin/subfinder ]
then
    echo "${magenta} [+] Running Subfinder for subdomain enumeration ${reset}"
    subfinder -d $DOM -o ~/Bugsubs/$DOM/Subdomains/subfinder.txt
else
    echo "${cyan} [+] Installing Subfinder ${reset}"
    go get -u -v github.com/projectdiscovery/subfinder/v2/cmd/subfinder
    echo "${magenta} [+] Running Subfinder for subdomain enumeration ${reset}"
    subfinder -d $DOM -o ~/Bugsubs/$DOM/Subdomains/subfinder.txt
fi
echo " "
echo "${cyan} [+] Successfully saved as subfinder.txt ${reset}"
echo " "

#uniquesubdomains
echo "${yellow} --------------------------xxxxxxxx--------------------- ${reset}"
echo " "
echo "${magenta} [+] Fetching unique domains ${reset}"
echo " "
cat ~/Bugsubs/$DOM/Subdomains/*.txt | sort -u >> ~/Bugsubs/$DOM/Subdomains/unique.txt
echo "${blue} [+] Successfully saved as unique.txt ${reset}"
echo " "

#sorting alive subdomains
echo "${yellow} ---------------------------xxxxxxxx-------------------- ${reset}"
echo  " "
if [ -f /usr/local/bin/httpx ]
then
    echo "${magenta} [+] Running Httpx for sorting alive subdomains ${reset}"
    cat ~/Bugsubs/$DOM/Subdomains/unique.txt | httpx >> ~/Bugsubs/$DOM/Subdomains/all-alive-subs.txt
    cat ~/Bugsubs/$DOM/Subdomains/all-alive-subs.txt | sed 's/http\(.?*\)*:\/\///g' | sort -u ~/Bugsubs/$DOM/Subdomains/protoless-all-alive-subs.txt
else
    echo "${cyan} [+] Installing Httpx ${reset}"
    go get -u github.com/projectdiscovery/httpx/cmd/httpx
    echo "${magenta} [+] Running Httpx for sorting alive subdomains ${reset}"
    cat ~/Bugsubs/$DOM/Subdomains/unique.txt | httpx >> ~/Bugsubs/$DOM/Subdomains/all-alive-subs.txt
    cat ~/Bugsubs/$DOM/Subdomains/all-alive-subs.txt | sed 's/http\(.?*\)*:\/\///g' | sort -u > ~/Bugsubs/$DOM/Subdomains/protoless-all-alive-subs.txt
fi
echo " "
echo "${cyan} [+] Successfully saved the results"
echo " "

