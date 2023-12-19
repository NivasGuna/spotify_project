# MusictrackAPI

## Table of Contents
1. [Overview](#overview)
2. [Features](#features)
3. [Technologies Used](#technologies-used)
4. [Installation](#installation)
5. [Database Configuration](#database-configuration)
6. [Spotify Authentication](#spotify-authentication)
7. [Swagger Documentation](#swagger-documentation)
8. [Endpoints](#endpoints)



## Overview

MusictrackAPI is a Go-based API designed to manage music tracks. It allows users to search for tracks by ISRC or artist's name, store tracks in the database, and update track details.

## Features

- Fetch tracks by ISRC or artist name
- Store tracks in the database
- Update track details

## Technologies Used

- Go
- Gin (HTTP web framework)
- GORM (Object-Relational Mapping)
- PostgreSQL (Database)
- Swagger (API Documentation)
- Spotify API (for track information)

## Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/your_username/music-track-api.git
   cd music-track-api

2. install dependencies:

   go mod tidy

3. Run the application:
   
 go run main.go

## Database Configuration
  

The application uses a database to store track information. To configure the database:

1.Set up a database (MySQL, PostgreSQL, etc.).
2.Update the database credentials in the config or database package.
  

## Spotify Authentication

To use the Spotify features, you'll need to set up authentication with Spotify. Here are the steps:

1. Go to the Spotify Developer Dashboard (https://developer.spotify.com/dashboard/applications).
2. Create a new application.
3. Obtain the Client ID and Client Secret.
4. Set up environment variables or configuration files to store these keys securely.

## Swagger Documentation

Swagger is integrated into this API for easy documentation and testing. To access Swagger documentation:

1.Generate Swagger files using Swaggo:
  swag init

2.Run the application:
  
  go run main.go

3.Access Swagger UI in your browser:

 http://localhost:8080/swagger/index.html

 Swagger provides an interactive UI to explore and test API endpoints.


## API Endpoints

Available Endpoints

  1.GET /getisrc/{isrc}: Retrieves a track by its ISRC.
  2.GET /getartist/{artistname}: Retrieves tracks by the artist's name.
  3.POST /postisrc: Saves or searches for a track by its ISRC.
  4.PUT /updatetrack/{isrc}: Updates track details by ISRC.

  Make sure to check the Swagger documentation for detailed information about request and response payloads.
