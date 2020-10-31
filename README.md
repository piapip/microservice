Lesson 5: Gorilla

Gorilla stores variable of the URL in mux.Vars(*http.Request)

**The request** comes into the server, gets picked up by the router, the router says "Ei, a PUT request", send **a signal** to the PUT subRouter The PUT subRouter says "And I've got this middleware, I'll execute the middleware, if it passes, then I will shot **a signal** to trigger the handler func"

Middleware is like a bunch of validation, checkup, or data conversion.

=========================================================================================

Lesson 4: RESTful 2

RESTful methods: POST, PUT.
Not much going on. Just showing off how to do this manually without using any high gear library. DELETE should be the same as PUT.

So with traditional, raw Golang, we can't define placeholder for the URL, like the localhost/products/1 for PUT request, it will immediately shot back to localhost/. We will have to parse the URL ourselves (wasteful codes) and it's really time-consuming. But with Gorilla's mux package, we can use placeholder, regex for those URL.

=========================================================================================

Lesson 3: RESTful

RESTful methods: GET.
Working with data. Just some half-ass created data. Not with proper database yet.
Using Encoder and working with JSON format.

=========================================================================================

Lesson 2: Introduction 2.

Refactor code to regroup handlers. Using ServeMux to use customized handler instead of defaultHandle in the ListenAndServe.
Learn how to config basic stat of a server like Address, handler, IdleTimeout,... and gracefully shutdown.

=========================================================================================

Lesson 1: Introduction.

Learn how to initiate a simple server that can send information. 
A very simple GET request. (doesn't specify that's a GET request though)
Learn what http package can do.
Remember that http also help you handle errors if there's any.

=========================================================================================