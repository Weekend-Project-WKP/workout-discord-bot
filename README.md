# Workout Discord Bot
A Discord bot written in Golang for tracking workouts over a period and generating basic reports to help users monitor their progress. Designed to be lightweight, easy to use, and fully integrated within your Discord server!

## Features
* Workout Tracking: Allows users to log various workout details like exercises, sets, reps, and time.
* Progress Reports: Generates summary reports on workout frequency, types, duration, and more, providing an overview of progress over a specified time period.
* User-Friendly Commands: Simple commands to add, edit, or delete workouts and view stats.
* Customizable Time Periods: Choose specific dates or periods to view workout summaries.

## Commands
* !logworkout [details] - Log a new workout with details like exercise name, reps, sets, and duration.
* !workoutsummary [time period] - Generate a report for a specific time period (e.g., weekly, monthly).
* !editworkout [workout ID] [new details] - Edit an existing workout entry.
* !deleteworkout [workout ID] - Delete a workout entry.

## Installation
1. Clone the Repository
```
$ git clone https://github.com/Weekend-Project-WKP/workout-discord-bot.git
$ cd workout-discord-bot
```

2. Set Up Environment: Configure your `.env` file with your Discord bot token.

3. Run the Bot:
```
go run main.go
```

## Requirements
* Golang version 1.23.3 or higher
* Discord Bot Token (available from the Discord Developer Portal)

## Future Improvements
* Comming soon!

## License
This project is licensed under the MIT License.