# Backend challenge

## Introduction

The goal of this challenge is to implement an application that scrapes a particular website and returns the extracted
data in a clear format. There are some restrictions in terms of the stack to be used and also some bonus/optional tasks
to do.

## Platform to scrape

* URL: [https://app.comet.co/freelancer](https://app.comet.co/freelancer/signin?to=%2Ffreelancer%2Fdashboard)

## Requirements

Implement a scraper to extract information from Comet:

* The implementation has to include the logic to go through the login of the site (like filling the credentials)
  and extract profile information (https://app.comet.co/freelancer/profile) such as general user profile
  (like full name, profile picture url, etc), experiences and skills.
* The code has to be clean, well structured, easy to read and following good programming practices.
* A Readme containing instructions on how to run the application and documentation about design decisions has to be
  provided.

## Bonus points:

* Check if the login was successful or not (for example because the credentials were wrong)
  and return the error from Comet to the end user
* Implement a simple REST API interface to do things like starting the scraping process,
  getting the status of it (if it’s still running, finished, etc), getting the data, etc.
* Store the data in a database (it could be sqlite)
* Use docker

## Stack

The challenge has to be solved using Golang and the Chromedp lib (https://github.com/chromedp/chromedp) for the scraping
part.

## Instructions

* copy `services/cometco-scraper/.env.example` to `services/cometco-scraper/.env` and fill the values
    * almost all fields can be used as is. Only exception are credentials, that are used as defaults for the login (
      described below, in `Inplementation details` section)
        * `SC_COMETCO_SCRAPER_PROFILE_CREDENTIALS_EMAIL`
        * `SC_COMETCO_SCRAPER_PROFILE_CREDENTIALS_PASSWORD`
* run `docker-compose up` to start the application
* use `api/comet.yaml` for importing into clients that support OpenAPI (like Postman, Insomnia, etc)
    * or, if you prefer curl, you still can use it as a guide for knowing the endpoints and their parameters
* Performing requests for creating tasks, checking their statuses, or getting `result_id`, for viewing scraped data.

## Implementation details

### Project structure

Basically, project code is divided into 3 parts:

- `api` : contains OpenAPI specification
- `services` : contains our application
- `common` : contains (potential) common code.

You may notice, that on project root level is not present `go.mod` file. Instead, there is `go.mod` in `common`
and `services/cometco-scraper` directories.
This is especially useful, when we have multiple services in the project, and we want to have more transparent
understanding of what dependencies are used by each service.

### cometco-scraper service

This service is responsible for scraping data from Comet.
As entry point, we have REST API, that is described in `api/comet.yaml` file.
This API, allows us to perform following actions:

* Create new task
* Get list of task ids, that previously were created
* Get task details, by its id
    * where will be present error/fail reason, if task failed
    * or result_id, if task was successful
* Get scraped data, by result_id

#### How is created task?

We are able to create new task by providing credentials, that will be used for login.
In case if both fields are empty (or will be sent empty object), then default credentials will be used.
Default credentials are stored in `.env` file, and are described in `Instructions` section.

After submitting credentials, we create new task, and return its id to the user.
Also, we start scraping process in background.

#### How is scraping process working?

Scraping process is divided into 4 steps:

* We try to login to Comet, using provided credentials
* If login was successful, we try to access profile page
* We try to extract data from profile page, into "local" model
* Convert local structure, into domain model

#### Why do we have so many models? (and converters)

At the moment, we have 3 models:

* models of data, that we are scraping
* domain model, that is used for storing data in database
* http model, that is used for returning data to the user via REST API

This is done, because we want to have clear separation of concerns.
Each part: HTTP, DB and Scraper, have it own responsibilities, and logic.
If something changes in one part, it should not affect other parts.
Also, this allows us to have more flexibility, when it comes to changing data structure. Or we can hide/manipulate data,
on different levels.

Also, if in the future we will decide to change/refactor something, this will be easier to do, because we have separated
already them. Under possible changes, I mean:

* Change Database from MongoDB to something else
* Change REST API to gRPC
* Change Scraper from Chromedp to something else
* Moving any of parts, to separate service

#### Why we consider that common package contains "potential" common code?

In actual code, probably everything inside `common` package can be moved to `services/cometco-scraper` package.
Originally, I was thinking to have our HTTP/Gateway as separate service, that just handles requests, and forwards them
to Scraper service.
And, in the future, if will be required to add support of scraping other websites, moving it to separate service, will
be a must. (and this will be easier to do, because we already have common package :) )

But, at the moment, I decided to not over-complicate things, and just keep HTTP in `cometco-scraper` service.
