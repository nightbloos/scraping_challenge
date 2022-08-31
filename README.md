# Backend challenge

## Introduction
The goal of this challenge is to implement an application that scrapes a particular website and returns the extracted data in a clear format. There are some restrictions in terms of the stack to be used and also some bonus/optional tasks to do.

## Platform to scrape
* URL: [https://app.comet.co/freelancer](https://app.comet.co/freelancer/signin?to=%2Ffreelancer%2Fdashboard)

## Requirements
Implement a scraper to extract information from Comet:
* The implementation has to include the logic to go through the login of the site (like filling the credentials) 
and extract profile information (https://app.comet.co/freelancer/profile) such as general user profile
(like full name, profile picture url, etc), experiences and skills.
* The code has to be clean, well structured, easy to read and following good programming practices.
* A Readme containing instructions on how to run the application and documentation about design decisions has to be provided.

## Bonus points:
* Check if the login was successful or not (for example because the credentials were wrong)
and return the error from Comet to the end user
* Implement a simple REST API interface to do things like starting the scraping process,
getting the status of it (if it’s still running, finished, etc), getting the data, etc.
* Store the data in a database (it could be sqlite) 
* Use docker
  
## Stack
The challenge has to be solved using Golang and the Chromedp lib (https://github.com/chromedp/chromedp) for the scraping part.