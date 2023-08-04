# Net/Http
 It is a standard package provided by golang where it  is used to build web application,The net/http package in Go provides functionalities for building HTTP servers and clients. 
 
 The net/http package in Go is a powerful tool for building HTTP servers and clients. It provides various functionalities and features that make it suitable for a wide range of use cases. 
 
 ****Building HTTP servers****: You can use the net/http package to create HTTP servers to handle incoming requests. This is useful for building web applications, APIs, or microservices
 
****Creating RESTful APIs****: With the net/http package, you can define routes and handlers to implement RESTful APIs. It allows you to handle different HTTP methods (GET, POST, PUT, DELETE) and perform CRUD operations on resources.

***Serving static files***: The net/http package can serve static files such as HTML, CSS, JavaScript, images, and other assets. This is useful for building web applications that require serving static content.

These are just a few examples of the uses of the net/http package in Go. It provides a flexible and powerful framework for building HTTP-based applications, both on the server and client side.

# What is  ListenAndServe(":portnumber",nil)
ListenAndServe is a function in the net/http package in Go that starts an HTTP server and begins listening for incoming HTTP requests on a specified network address.

The ListenAndServe function takes two parameters:

The network address to listen on, specified as a string in the format "host:port". For example, "localhost:8080" or ":8000" (listen on all available network interfaces).

An optional parameter of type http.Handler, which specifies the handler to use for incoming requests. If nil is provided, the default DefaultServeMux will be used.

When ListenAndServe is called, it blocks the execution of your Go program and starts a continuous loop to handle incoming HTTP requests. It dispatches each request to the specified handler based on the request's URL path.

****Remember you always give ListenAndServe at the end of the program****


# What is http.Get(url)
In Go's net/http package, the http.Get function is used to send an HTTP GET request to the specified URL and retrieve the response. It is a convenient way to make simple GET requests without explicitly creating an http.Client and http.Request objects.

In the above example, http.Get(url) sends a GET request to the specified url ("https://go.dev/doc/" in this case) and returns the response as an http.Response object. We then read the response body using ioutil.ReadAll(response.Body) and print it to the console.

Note that http.Get is a simplified way to make GET requests, but it does not provide as much control and flexibility as manually creating an http.Client and http.Request objects. For more complex scenarios, it is recommended to use the http.Client and construct requests using http.NewRequest.

# What is url.Parse("URL")

When you parse a URL using url.Parse() in Go, it breaks down the URL string into its individual components and provides you with easy access to those components. Here's what each component represents:

Scheme: The scheme specifies the protocol or application layer protocol used in the URL, such as "http", "https", "ftp", etc.

****Path****: The path component represents the hierarchical path of the resource within the server's file system or routing structure.

****Hostname:**** The hostname is the domain name or IP address of the server.
Port: The port component specifies the port number on the server to which the client should connect. If no port is specified, it will return the default port for the scheme.

****RawQuery:**** The raw query represents the query string of the URL, which typically consists of key-value pairs used to pass parameters to the server.


By parsing the URL, you can access these individual components separately, allowing you to work with them in your code as needed. It provides a convenient way to extract specific information from a URL and utilize it in your application logic.