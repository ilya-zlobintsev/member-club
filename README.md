This is a basic demo app for displaying a list of club members.

https://ilyaz-member-club.herokuapp.com/

![image](https://user-images.githubusercontent.com/22796665/146759421-88092874-77ad-41f8-9c9c-5d95f4e6ab53.png)


## Building locally
Requirements:
- golang 1.16+

Commands:

`go run .`

It's also possible to run in "debug" mode, which means web assets will be loaded dynamically from files and you can change them without restarting the app (by default they're embedded at compile time):

`go run -tags debug .`

## Building with Docker

`docker build -t member-club .`

`docker run -p 8080:8080 member-club`
