PUBG Player Statistic Comparison
================================
PUBGstats is is simple project that is designed to pull player statistics from the PUBG player statistic [api](https://pubgtracker.com/site-api) and compare player skill ratings, kill-death ratios, and win rates to determine the better player.  

Documentation
-------------
This project requires a TRN API key which can be had by following the link to the site, registering for an account and requesting an API key.

Getting Started
---------------
Build the binary from source using the provided `Makefile`:  
`make`  

Run the binary:  
`go run main.go -k <your TRN api key>`  
or  
`./pubgstats -k <your TRN api key>`  

Developing PUBGstats 
--------------------
Feel free!
