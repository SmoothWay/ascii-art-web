## Description

Project was developed during the study in [alem.school](https://alem.school)

Project serves to draw ascii-art in the web page with 3 different fontTypes.

## Usage: how to run?
Clone this repository to your local machine. `cd ascii-art-web` and type `go run .`  
After you can go to the `localhost:8080` or your specified port in code.  
Then enter the text and choose fontype and press submit button. After that you will see the result.

## Implementation

Script based on ascii-art terminal version and implemented for browser usage.  
It uses `maps` for drawing ascii-art, which it gets from txt file that contains ascii symbol's drawing versions.  
Web page handles only POST requests and logs status codes in terminal.
