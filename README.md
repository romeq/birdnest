# Birdnest
My solution to Reaktor's summer job <a href="https://assignments.reaktor.com/birdnest">pre-assignment</a>.

### What?
For summer-trainee role Reaktor requires a solution for their custom-made pre-assignment which you can get to from the link above. 
As I was highly interested in the challenge, I decided to take a look at it.
This repository contains a few words of mine in addition to the code and solution to the challenge.

### A few simple words
The challenge was itself quite fun and I really enjoyed doing it. That said, unarguably the best part was getting to solve a problem using technologies that I knew well.


### Running this app

**NOTE**: The webserver will start at port 8080 except if configured otherwise

You're able to run this program by compiling the binary or building a docker image:

- Building the binary

  ```bash
  $ go build . -o birdnest-server
  $ ./birdnest-server
  ```

- Building and running docker image

  ```bash
  $ sudo docker build -t birdnest
  
  # with docker-compose (don't detach)
  $ sudo docker-compose up
  # with docker-compose (detach, a.k.a. run in the background)
  $ sudo docker-compose up -d
  ```


### Caveats

- The website is not the most optimized as it always re-renders all pilots on each view update.
