<div style="display: flex;
            justify-content: space-between;
            gap: 40px;">
	<div>
    	<h1>Birdnest - Reaktor Assignment</h1>

        <p>
            My solution to Reaktor's summer job <a href="https://assignments.reaktor.com/birdnest">pre-assignment</a>.
        </p>
    
        <h3>
            What?
        </h3>
        <p>
            For summer-trainee role Reaktor requires a solution for their custom-made pre-assignment which you can get to from the link above. 
            As I was highly interested in the challenge, I decided to take a look at it.
            This repository contains a few words of mine in addition to the code and solution to the challenge.
        </p>
        <h3>
            A few simple words
        </h3>
        <p>
            The challenge was itself quite fun and I really enjoyed doing it. That said, unarguably the best part was getting to solve a problem using technologies that I knew well.
        </p>
    </div>
    
    <img style="width: 30%;
                border-radius: 10px;
                " src="/home/toke/.config/Typora/typora-user-images/image-20221227230951332.png"></img>
</div>

## Running this app

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

- The website is not the most optimized as it always rerenders all pilots on each update.
