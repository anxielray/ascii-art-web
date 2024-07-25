# _*Ascii-art-web*_


## _*Description*_


Ascii-art-web is a web application that provides a graphical user interface (GUI) for creating ASCII art banners. The application allows users to input text and select a banner style (shadow, standard, or thinkertoy) to generate an ASCII art banner.


## _*Authors*_


* [<Joan](http://github.com/joan254)
* [Cherryl](http://github.com/Cherrypick14)
* [Raymond](http://github.com/anxielray)

### _*Installation and setup*_

- To install we can use the commandline `git` arguments to clone it locally to our machine. This can be achieved by running the command;

```bash
git clone https://learn.zone01kisumu.ke/git/jwambugu/ascii-art-web.git
```

- As well we also have an option to download the project directly from `Gitea` and we can manually access it through the file manager of your machine.

## _*Usage*_


- To run the application, follow these steps:


1. Navigate to the project directory: 

```bash
cd ascii-art-web
```


3. Run the application: 

```bash
go run main.go
```

4. Open a web browser at the port provided: `http://localhost:8080`


## _*Implementation Details*_


- The application uses the Go programming language to create an HTTP server that handles GET and POST requests. The server uses HTML templates to display data to the user.


### _*Algorithm*_


1. The user sends a GET request to the root URL (`/`) to retrieve the main page.

2. The server responds with an HTML template that includes a text input, radio buttons to select a banner style, and a button to submit the form.

3. When the user submits the form, the client sends a POST request to the `/ascii-art` endpoint with the text and banner style as form data.

4. The server processes the request and generates an ASCII art banner using the selected style. The page opens with the banner set tp standard set on a base default.

5. The server responds with an HTML template that displays the generated ASCII art banner.


### _*Code Structure*_


The code is organized into the following directories:


* `templates`: contains HTML templates for the main page and ASCII art banner display

* `main.go`: contains the HTTP server implementation


### _*HTTP Endpoints*_


* `GET /`: returns the main page HTML template

* `POST /ascii-art`: processes the form data and generates an ASCII art banner


### _*HTTP Status Codes*_


* `200 OK`: returned when the request is successful

* `404 Not Found`: returned when a requested resource is not found

* `400 Bad Request`: returned when the request is invalid

* `500 Internal Server Error`: returned when an unhandled error occurs


## _*Notes*_


* The application uses Go templates to display data to the user.

* The application uses form data to send the text and banner style to the server.

* The application uses an HTTP server to handle GET and POST requests.
