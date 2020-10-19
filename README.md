# Systems Assignment 

Language : GO

## What is it?

This exercise is a follow-on my [General Assignment](https://github.com/Sreevatsava1/cloudflare-2020-general-engineering-assignment). Created a tool that makes requests to the endpoints using TLS sockets instead of using HTTP libraries.

## Instructions to RUN the program

- Install go and setup environment to run go programs. [Official_Link] (https://golang.org/doc/install)
- Unzip the archive and move into systems folder
- To execute the program enter the command
```
go run . --url <URL>
```
- Use help command to get more details
```
go run . --help
```
- If the run command didn't work try building the program
```
go build .
```
## Results

- Youtube and Pinterest took more time compared to others as they have to load lot of image data I guess.
- Cloudflare worker loaded fastest when compared to others.

## Screenshots

1. Cloudflare worker site links
![workersite](systems/screenshots/ss1.png)
2. Cloudflare worker site
![workersite](systems/screenshots/ss2.png)
3. Cloudflare website
![cloudflare](systems/screenshots/ss3.png)
4. Youtube website
![youtube](systems/screenshots/ss4.png)
5. Google.com
![google](systems/screenshots/ss5.png)
6. Pinterest 
![pinterest](systems/screenshots/ss6.png)