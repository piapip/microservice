# TRY 8: Generate code for client side with swagger 

This command: 
swagger generate client -f ../microservice/swagger.yaml -A api 
is bugged. I created a folder ../client and run that command from there. And instead of generating codes in that folder, the command created a new "client" folder in this thing, set up all the code here and leave that newly made ../client alone for some god damn reason who knows. But it's an open-source and who knows what's going on.

# Lesson 7: Swagger

Before we start this lesson, we need to fix the repository, got to do some rearrangement.

In this lesson, we will learn how to create a docs for our microservice by using swagger. Even though I say "learn", I can't really do it right away, it might take some time since I can only do it by myself after iterations of practices. Sitting by the documents and start grind through the content. Don't have to remember everything though.

The best way to understand all the parameters in swagger docs thingy is to actually put your hands into the jobs and start messing things up.

# Lesson 6: JSON Validation

Using Validator package to create some validators for our model Products and implement the validator to the middleware.
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

RESTful methods: GET.
Working with data. Just some half-ass created data. Not with proper database yet.
Using Encoder and working with JSON format.

# Lesson 2: Introduction 2.

Refactor code to regroup handlers. Using ServeMux to use customized handler instead of defaultHandle in the ListenAndServe.
Learn how to config basic stat of a server like Address, handler, IdleTimeout,... and gracefully shutdown.

# Lesson 1: Introduction.

Learn how to initiate a simple server that can send information. 
A very simple GET request. (doesn't specify that's a GET request though)
Learn what http package can do.
Remember that http also help you handle errors if there's any.