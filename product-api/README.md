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