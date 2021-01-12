# Practice 15: Refactoring

1. Retrieve real currency conversion rate and implement to the service. 

# Lesson 14: gRPC Client Connections

Using enum in protos. In many cases, we only take a handful of options as input, e.g: There're only so many currencies type out there in the world, so we won't accept any kind of string as inputs.\

In this lesson, we'll try to call api in ./currency from ./product-api. \
In order to do that, we will need to construct a client that allow us to call the currency server. What we kinda need to do is to use that proto file and use the client that it generated for us. The currency client in gRPC is generated in the file currency_grpc.pb.go (in the tutorial video, it's in the currency.pb.go).\ 
It's literally `type CurrencyClient interface {...}` and it has the method NewCurrencyClient right below for us to create a client server.

To create a client from grcp generated file, We'll need to initiate to a particular service, gRPC ClientConnection interface, then pass it to the New...Client(), the way we create connect in gRPC is kinda the exact way we create connection using "net".

Not much going on in this chapter tbh, it's just us calling generated client interface to our main code. Use what we've generated before.

gRPC has a thing called CallOption. Like cors, it won't let anything to call into it, so if we use gRPC.WithInsecure(), it's an option which says don't worry about any client or server side certificates, totally fine in development, totally not fine in production.

# Lesson 13: gRPC

With JSON based service and RESTful approach, HTTP requests, that way works great, easy to use, widely understood. But the problem is that it's not performance optimized. It's required a large amount of steps building client, server, intergration, ...\
gRPC is a new approach that Google came up with the intention behind which is, we are still using the standard protocols, but this time, it's going to be "HTTP 2" as opposed of "HTTP" and rather than JSON, we're going to use binary based message protocol called **PROTOBUF**.\
Protobuf, since they are already binary so it's gonna be faster to serialize and send it. And it ends up we defining interfaces - proto files - and anybody can generate a client based of these proto files that we created.\

In the new version of protoc plugins are not used anymore, to compile your .proto file to grpc you need to install protoc-gen-go-grpc. Basically you need both protoc-gen-go and protoc-gen-go-grpc.\

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
protoc -I protos/ --go-grpc_out=protos/currency --go_out=protos/currency protos/currency.proto

In this lesson, we will take a look at the generated code where it contains interface for us to implement to get ourselves a server.\
It will generate a `register...server` kinda function and we will have to implement it to get ourselves a server.

gRPC is not RESTful so things like cURL and Insomnia won't work if we want to test it out if it's running properly, we'll have to write test for it.


# Lesson 12: Gzipping audio

Though, we can transfer raw data normally between devices but when the file's size gets to a certain extent, we will have to think about Quality of Use. What if they have weak Internet, what if their device is not good enough to process raw data all at once. That's when we think of zipping data. Remind you, zipping and unzipping data is not free, it costs CPU but it's worth the effort.

Here's a command to try out with CURL, since I don't know how to use --compressed tag in Insomnia.<br />
curl -v http://localhost:9091/images/2/cosmopolitan.jpg --compressed -o file.jpg

# Lesson 11: HTTP multi-part Requests

It seems to be deprecated. It's used for uploading stuff, like images, or audios, or videos,... When browser were just static HTML, there weren't a lot of JS framework to use, if you want to send a request, you'll have a button and you'd be using HTML form, push the form and the browser will do the rest. Browser then sends given data in a data format called **multi-part form data**. It's HTTP data, combination of text and binary information and it would separate by using boundaries. It's not RESTful. NOTE: handler is designed to deal with REST.

It just happens so that I need to learn uploading file. antd Upload / desgin is so shit lol.

# Lesson 10: Files

First step towards microservice, file allocation, all the files must be positioned correctly. Microservice is pretty much all about structure. So if you messed up like I did, take some time and reorganize. Since I have no experience whatsoever to know if this structure will spawn any kind of bugs later on, I think it's safe to say that there's no weird bug or weird warning after generating swagger docs, and swagger sdk is a huge stepforward.

Learn how to upload images file. I should be able to upload other file types, like audio, or video. *audio and video can be saved as blob I suppose, I'm not too experienced about this stuff.

# Lesson 9: CORS

As usual, a problem during dev phase, here's how to handle cors with gorilla.\
import this: "github.com/gorilla/handlers"\
then this: corsHandler := gorilla_handlers.CORS(gorilla_handlers.AllowedOrigins([]string{"*"}))\
then in the server handler, change from \
    Handler:      serveMux     -----------------to------------------>     Handler:      corsHandler(serveMux)

REMEMBER: IN THE CLIENT SIDE, YOU HAVE TO SPECIFY THE BACKEND SERVER AS "http://loca..." NOT JUST "loca...", otherwise it won't work

# TRY 8: Generate code for client side with swagger 

Command\
swagger generate client -f ./swagger.yaml --target=sdk/ -A api \
the --target seems to be crucial because without it, it will generate code directly to this microservice folder and it would be messy and noone wants that.\
Without the --target, all the import will be messed up too.\
TLDR: --target is the must.\
It also has some bugs that related to not being able to find the spec file (my swagger.yaml file) so it's better for my head to just do it where the swagger.yaml file is.

DEBUGGING is important. Learn how to do that.

# Lesson 7: Swagger

Before we start this lesson, we need to fix the repository, got to do some rearrangement.

In this lesson, we will learn how to create a docs for our microservice by using swagger. Even though I say "learn", I can't really do it right away, it might take some time since I can only do it by myself after iterations of practices. Sitting by the documents and start grind through the content. Don't have to remember everything though.

The best way to understand all the parameters in swagger docs thingy is to actually put your hands into the jobs and start messing things up.

# Lesson 6: JSON Validation

Using Validator package to create some validators for our model Products and implement the validator to the middleware.\
Validator needs to have both built-in validation requirements (such as name, email,...) and customized requirements (in this lesson is SKU)

# Lesson 5: Gorilla

Gorilla stores variable of the URL in mux.Vars(*http.Request)

**The request** comes into the server, gets picked up by the router, the router says "Ei, a PUT request", send **a signal** to the PUT subRouter The PUT subRouter says "And I've got this middleware, I'll execute the middleware, if it passes, then I will shot **a signal** to trigger the handler func"

Middleware is like a bunch of validation, checkup, or data conversion.

# Lesson 4: RESTful 2

RESTful methods: POST, PUT.
Not much going on. Just showing off how to do this manually without using any high gear library. DELETE should be the same as PUT.

So with traditional, raw Golang, we can't define placeholder for the URL, like the localhost/products/1 for PUT request, it will immediately shot back to localhost/. We will have to parse the URL ourselves (wasteful codes) and it's really time-consuming. But with Gorilla's mux package, we can use placeholder, regex for those URL.

# Lesson 3: RESTful

RESTful methods: GET.\
Working with data. Just some half-ass created data. Not with proper database yet.\
Using Encoder and working with JSON format.

# Lesson 2: Introduction 2.

Refactor code to regroup handlers. Using ServeMux to use customized handler instead of defaultHandle in the ListenAndServe.\
Learn how to config basic stat of a server like Address, handler, IdleTimeout,... and gracefully shutdown.

# Lesson 1: Introduction.

Learn how to initiate a simple server that can send information.\
A very simple GET request. (doesn't specify that's a GET request though)\
Learn what http package can do.\
Remember that http also help you handle errors if there's any.
