# MusictrackAPI

## Table of Contents
1. [Overview](#overview)
2. [Features](#features)
3. [Technologies Used](#technologies-used)
4. [Installation](#installation)
5. [Database Configuration](#database-configuration)
6. [Spotify Configuration](#spotify-configuration)
7. [Swagger Documentation](#swagger-documentation)
8. [API Endpoints](#api-endpoints)



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
   

2. install dependencies:

   ```
   go mod tidy

   ```


## Database Configuration
  
**PostgreSQL**

1. **Install PostgreSQL:**


   Make sure PostgreSQL is installed on your machine. You can download it from the official website or use a package manager.


2. Create a Database:

```
  createdb your_database_name
```
3. Connection string

```
  DATABASE_URL=postgres://your_database_user:your_database_password@localhost:5432/your_database_name?sslmode=disable
```
replace your_database_user, your_database_password and your_database_name with your actiual PostgreSQL credentials.
  

## Spotify Configuration

1. **Create a Spotify Developer Account**

    - Go to the Spotify Developer Dashboard.
    - Log in or sign up for a Spotify account.
  
2. **Create a New Application**

    - After logging in, go to the Dashboard and click on "Create an App."
    - Fill in the required information for your application (name, description, etc.).
    - Once created, you'll see your application listed on the Dashboard.
      
3. **Get Client ID, Client Secret, and Redirect URI**

    - Open your application from the Dashboard.
    - Note down the ***Client ID*** and ***Client Secret***. These will be used in your application for authentication.
    

4. **Set Environment Variables**

    In your project's .env file or environment configuration, set the following variables:

      
   ```
   SPOTIFY_CLIENT_ID=your_client_id
   
   SPOTIFY_CLIENT_SECRET=your_client_secret
   ```

   Replace   your_client_id,  and  your_client_secret  with the corresponding values obtained from the Spotify Developer 
   Dashboard.

## Swagger Documentation

Swagger is integrated into this API for easy documentation and testing. To access Swagger documentation:

1. **Generate Swagger files using Swaggo:**

```
swag init
```

2. **Run the application:**
  
  ```
go run main.go

```

3. **Access Swagger UI in your browser:**

 ```
  http://localhost:8080/swagger/index.html
```

 Swagger provides an interactive UI to explore and test API endpoints.


## API Endpoints

Available Endpoints

```
   GET /getisrc/{isrc}          //Retrieves a track by its ISRC.

   GET /getartist/{artistname}  //Retrieves tracks by the artist's name.

   POST /postisrc               //Saves or searches for a track by its ISRC.

   PUT /updatetrack/{isrc}      //Updates track details by ISRC.
```
  Make sure to check the Swagger documentation for detailed information about request and response payloads.

## Output
**Swagger**

![swaggerhome](https://github.com/NivasGuna/spotify_project/assets/152397268/13af0139-b46a-4ba9-9568-3cd6410f3bc2)

**GetByIsrc**

![getIsrc](https://github.com/NivasGuna/spotify_project/assets/152397268/7361c43a-ce9e-478c-9a9f-e26173661f52)

**GetByArtist**

![getartist](https://github.com/NivasGuna/spotify_project/assets/152397268/9d62e782-8043-4ec6-a2a7-8a879afa5909)

**PostIrsc**

![postIsrc](https://github.com/NivasGuna/spotify_project/assets/152397268/882ac682-1078-457a-a74a-6e61eca74628)

**PutIrsc**

![putIsrc](https://github.com/NivasGuna/spotify_project/assets/152397268/08b59439-8bb6-42bb-9d47-98a3c266dad8)

**Database details**

![databasedetail](https://github.com/NivasGuna/spotify_project/assets/152397268/35b1fd10-18c3-4165-8fc5-61b0a7cd39aa)
