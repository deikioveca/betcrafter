BetCrafter
-
BetCrafter is a backend service for managing and analyzing football betting tickets.
It allows the user to store,update and track their football betting tickets, then extract useful insights to improve betting decisions over time.
The system provides core functionality for ticket management and an analyzer interface to generate actionable information from the data.


Features
-
* Ticket management
  * Create new football betting tickets
  * Retrieve pending tickets
  * Update pending tickets
  * Get ticket by status or id
* Data analysis
  * Analyze ticket history and performance
  * Provide insights and trends to improve future betting decisions
* User interaction
  * Supports both cli and web interfaces
  * Easy-to-use endpoints for ticket management and data analysis
* Database integration
  * Store tickets and analysis data in PostgreSQL
  * Maintain historical records for trend analysis


Technologies used
-
* Go
* GORM
* PostgreSQL
* Docker


Running guide
-
* Prerequisites
  * Go installed
  * Docker installed
  * PostgreSQL installed
* Locally
  * Clone the repo
  * Run go mod download
  * Start PostgreSQL and ensure that database betcrafter exist
  * Change .env file according to your setup
  * go run cmd/cli/main.go -> for the cli && go run cmd/web/main.go -> for the web
* Docker
  * Docker compose up


Workflow
-
* Create ticket              - store a new football betting ticket with stake and matches info
* Get pending ticket         - retrieve all tickets that are not yet resolved
* Update pending ticket      - modify ticket info: profit, cashout(true,false), status(won,lost), match result outcome if it's correct or wrong
* Get ticket by status or id - retrieve tickets by status or ticket by id
* Analyze tickets            - retrieve useful information about the tickets from date range: for example the 3 most profitable outcomes, overall statistic, etc.
