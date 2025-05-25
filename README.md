It's a CLI tool that allows users to:

- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post

# Commands

### Register
`gator register [name]`

Register a new user in the database and add it to the `~/.gatorconfig.json` file.

### Login 

`gator login [name]`

Set the user (if exists) in the config file.

### Reset

`gator reset`

Delete all users on the database.

### Get all users

`gator users`

Get all users on the database.

### Get feed

`gator agg`

[WIP] Currently displaying a mock from https://www.wagslane.dev/index.xml
Get feed.