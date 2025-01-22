[![Go](https://github.com/Weekend-Project-WKP/workout-discord-bot/actions/workflows/go.yml/badge.svg)](https://github.com/Weekend-Project-WKP/workout-discord-bot/actions/workflows/go.yml)
[![CodeQL](https://github.com/Weekend-Project-WKP/workout-discord-bot/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/Weekend-Project-WKP/workout-discord-bot/actions/workflows/github-code-scanning/codeql)
# Workout Discord Bot
A Discord bot written in Golang for tracking workouts over a period and generating basic reports to help users monitor their progress. Designed to be lightweight, easy to use, and fully integrated within your Discord server!

## Features
* Workout Tracking: Allows users to log various workout details like exercises, sets, reps, and time.
* Progress Reports: Generates summary reports on workout frequency, types, duration, and more, providing an overview of progress over a specified time period.
* User-Friendly Commands: Simple commands to add, edit, or delete workouts and view stats.
* Customizable Time Periods: Choose specific dates or periods to view workout summaries.

## Commands
* !workoutsummary [time period] - Generate a report for a specific time period (e.g., weekly, monthly).
* !editworkout [workout ID] [new details] - Edit an existing workout entry.
* !deleteworkout [workout ID] - Delete a workout entry.

## Slash Commands
* /workout [workout-category][workout-duration] - Log a new workout for the current day for the current user

## Emoji Reactions
* 🧪 - Triggers Ai to interpret the reacted message (with or without images) to provide a workout score
* ✅ - Calculates points and Adds a workout summary to the DB
* 💪🏿 - Adds and removes the "Workout Challenge role" to the user that the emoji was reacted to

## Installation
1. Clone the Repository
```
$ git clone https://github.com/Weekend-Project-WKP/workout-discord-bot.git
$ cd workout-discord-bot
```

2. Set Up Environment: Configure your `.env` file with the required tokens.
```
DISCORD_TOKEN = <insert token>
GEMINI_API_KEY = <insert key>
MONGODB_PASSWORD = <insert password>
MONGODB_USERID = <insert userid>
```

3. Run the Bot:
```
go run main.go
```

## Requirements
* Golang version 1.23.3 or higher
* Discord Bot Token (available from the Discord Developer Portal)
* Gemini Api Key (obtainable from https://aistudio.google.com/apikey)
* MongoDB connection string credentials

## Future Improvements
* Make Emoji Reactions skin tone insensitive
* Make the Ai Prompt grab its scores dynamically by the Scoring Structure sheet
* Refine the Ai Prompt

## License
This project is licensed under the MIT License.
