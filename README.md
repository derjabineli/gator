# Gator

Gator is CLI RSS Feed tool. Gator uses a local database to allow you to register and login various users. Users can add and follow RSS feeds that will aggregate based on a specified duration. Once a user is following an RSS feed they can see the aggregated posts in their CLI with a title, link to complete post, and a short description. Gator is intended to run continously in the background and will continue to aggregate as long as it is active.

## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)

## Requirements

[Go](https://go.dev/doc/install) (version 1.23+)

- Go is required to be installed so that you can run Go's install command and install Gator on your own machine

[Postgres](https://www.postgresql.org/download/)

- Postgres is used as the database for this project and will store all necessary information such as users, RSS feeds, posts data and more.

## Installation
