# Fiber

## Fiber application

Creating a new Fiber application with the default configuration

```go
import "github.com/gofiber/fiber/v2"

func main() {
  // Create a new Fiber instance, called 'app'
  app := fiber.New()

  // ...
}
```

## Configuration

Fiber has a special `fiber.Config` structure that contains many fields:

```go
type Config struct {
	// When set to true, this will spawn multiple Go processes listening on the same port.
	//
	// Default: false
	Prefork bool `json:"prefork"`

	// Enables the "Server: value" HTTP header.
	//
	// Default: ""
	ServerHeader string `json:"server_header"`

	// When set to true, the router treats "/foo" and "/foo/" as different.
	// By default this is disabled and both "/foo" and "/foo/" will execute the same handler.
	//
	// Default: false
	StrictRouting bool `json:"strict_routing"`

	// When set to true, enables case sensitive routing.
	// E.g. "/FoO" and "/foo" are treated as different routes.
	// By default this is disabled and both "/FoO" and "/foo" will execute the same handler.
	//
	// Default: false
	CaseSensitive bool `json:"case_sensitive"`

	// When set to true, this relinquishes the 0-allocation promise in certain
	// cases in order to access the handler values (e.g. request bodies) in an
	// immutable fashion so that these values are available even if you return
	// from handler.
	//
	// Default: false
	Immutable bool `json:"immutable"`

	// When set to true, converts all encoded characters in the route back
	// before setting the path for the context, so that the routing,
	// the returning of the current url from the context `ctx.Path()`
	// and the parameters `ctx.Params(%key%)` with decoded characters will work
	//
	// Default: false
	UnescapePath bool `json:"unescape_path"`

	// Enable or disable ETag header generation, since both weak and strong etags are generated
	// using the same hashing method (CRC-32). Weak ETags are the default when enabled.
	//
	// Default: false
	ETag bool `json:"etag"`

	// Max body size that the server accepts.
	// -1 will decline any body size
	//
	// Default: 4 * 1024 * 1024
	BodyLimit int `json:"body_limit"`

	// Maximum number of concurrent connections.
	//
	// Default: 256 * 1024
	Concurrency int `json:"concurrency"`

	// Views is the interface that wraps the Render function.
	//
	// Default: nil
	Views Views `json:"-"`

	// Views Layout is the global layout for all template render until override on Render function.
	//
	// Default: ""
	ViewsLayout string `json:"views_layout"`

	// PassLocalsToViews Enables passing of the locals set on a fiber.Ctx to the template engine
	//
	// Default: false
	PassLocalsToViews bool `json:"pass_locals_to_views"`

	// The amount of time allowed to read the full request including body.
	// It is reset after the request handler has returned.
	// The connection's read deadline is reset when the connection opens.
	//
	// Default: unlimited
	ReadTimeout time.Duration `json:"read_timeout"`

	// The maximum duration before timing out writes of the response.
	// It is reset after the request handler has returned.
	//
	// Default: unlimited
	WriteTimeout time.Duration `json:"write_timeout"`

	// The maximum amount of time to wait for the next request when keep-alive is enabled.
	// If IdleTimeout is zero, the value of ReadTimeout is used.
	//
	// Default: unlimited
	IdleTimeout time.Duration `json:"idle_timeout"`

	// Per-connection buffer size for requests' reading.
	// This also limits the maximum header size.
	// Increase this buffer if your clients send multi-KB RequestURIs
	// and/or multi-KB headers (for example, BIG cookies).
	//
	// Default: 4096
	ReadBufferSize int `json:"read_buffer_size"`

	// Per-connection buffer size for responses' writing.
	//
	// Default: 4096
	WriteBufferSize int `json:"write_buffer_size"`

	// CompressedFileSuffix adds suffix to the original file name and
	// tries saving the resulting compressed file under the new file name.
	//
	// Default: ".fiber.gz"
	CompressedFileSuffix string `json:"compressed_file_suffix"`

	// ProxyHeader will enable c.IP() to return the value of the given header key
	// By default c.IP() will return the Remote IP from the TCP connection
	// This property can be useful if you are behind a load balancer: X-Forwarded-*
	// NOTE: headers are easily spoofed and the detected IP addresses are unreliable.
	//
	// Default: ""
	ProxyHeader string `json:"proxy_header"`

	// GETOnly rejects all non-GET requests if set to true.
	// This option is useful as anti-DoS protection for servers
	// accepting only GET requests. The request size is limited
	// by ReadBufferSize if GETOnly is set.
	//
	// Default: false
	GETOnly bool `json:"get_only"`

	// ErrorHandler is executed when an error is returned from fiber.Handler.
	//
	// Default: DefaultErrorHandler
	ErrorHandler ErrorHandler `json:"-"`

	// When set to true, disables keep-alive connections.
	// The server will close incoming connections after sending the first response to client.
	//
	// Default: false
	DisableKeepalive bool `json:"disable_keepalive"`

	// When set to true, causes the default date header to be excluded from the response.
	//
	// Default: false
	DisableDefaultDate bool `json:"disable_default_date"`

	// When set to true, causes the default Content-Type header to be excluded from the response.
	//
	// Default: false
	DisableDefaultContentType bool `json:"disable_default_content_type"`

	// When set to true, disables header normalization.
	// By default all header names are normalized: conteNT-tYPE -> Content-Type.
	//
	// Default: false
	DisableHeaderNormalizing bool `json:"disable_header_normalizing"`

	// When set to true, it will not print out the «Fiber» ASCII art and listening address.
	//
	// Default: false
	DisableStartupMessage bool `json:"disable_startup_message"`

	// This function allows to setup app name for the app
	//
	// Default: nil
	AppName string `json:"app_name"`

	// StreamRequestBody enables request body streaming,
	// and calls the handler sooner when given body is
	// larger then the current limit.
	StreamRequestBody bool

	// Will not pre parse Multipart Form data if set to true.
	//
	// This option is useful for servers that desire to treat
	// multipart form data as a binary blob, or choose when to parse the data.
	//
	// Server pre parses multipart form data by default.
	DisablePreParseMultipartForm bool

	// Aggressively reduces memory usage at the cost of higher CPU usage
	// if set to true.
	//
	// Try enabling this option only if the server consumes too much memory
	// serving mostly idle keep-alive connections. This may reduce memory
	// usage by more than 50%.
	//
	// Default: false
	ReduceMemoryUsage bool `json:"reduce_memory_usage"`

	// FEATURE: v2.3.x
	// The router executes the same handler by default if StrictRouting or CaseSensitive is disabled.
	// Enabling RedirectFixedPath will change this behavior into a client redirect to the original route path.
	// Using the status code 301 for GET requests and 308 for all other request methods.
	//
	// Default: false
	// RedirectFixedPath bool

	// When set by an external client of Fiber it will use the provided implementation of a
	// JSONMarshal
	//
	// Allowing for flexibility in using another json library for encoding
	// Default: json.Marshal
	JSONEncoder utils.JSONMarshal `json:"-"`

	// When set by an external client of Fiber it will use the provided implementation of a
	// JSONUnmarshal
	//
	// Allowing for flexibility in using another json library for decoding
	// Default: json.Unmarshal
	JSONDecoder utils.JSONUnmarshal `json:"-"`

	// XMLEncoder set by an external client of Fiber it will use the provided implementation of a
	// XMLMarshal
	//
	// Allowing for flexibility in using another XML library for encoding
	// Default: xml.Marshal
	XMLEncoder utils.XMLMarshal `json:"-"`

	// Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only)
	// WARNING: When prefork is set to true, only "tcp4" and "tcp6" can be chose.
	//
	// Default: NetworkTCP4
	Network string

	// If you find yourself behind some sort of proxy, like a load balancer,
	// then certain header information may be sent to you using special X-Forwarded-* headers or the Forwarded header.
	// For example, the Host HTTP header is usually used to return the requested host.
	// But when you’re behind a proxy, the actual host may be stored in an X-Forwarded-Host header.
	//
	// If you are behind a proxy, you should enable TrustedProxyCheck to prevent header spoofing.
	// If you enable EnableTrustedProxyCheck and leave TrustedProxies empty Fiber will skip
	// all headers that could be spoofed.
	// If request ip in TrustedProxies whitelist then:
	//   1. c.Protocol() get value from X-Forwarded-Proto, X-Forwarded-Protocol, X-Forwarded-Ssl or X-Url-Scheme header
	//   2. c.IP() get value from ProxyHeader header.
	//   3. c.Hostname() get value from X-Forwarded-Host header
	// But if request ip NOT in Trusted Proxies whitelist then:
	//   1. c.Protocol() WON't get value from X-Forwarded-Proto, X-Forwarded-Protocol, X-Forwarded-Ssl or X-Url-Scheme header,
	//    will return https in case when tls connection is handled by the app, of http otherwise
	//   2. c.IP() WON'T get value from ProxyHeader header, will return RemoteIP() from fasthttp context
	//   3. c.Hostname() WON'T get value from X-Forwarded-Host header, fasthttp.Request.URI().Host()
	//    will be used to get the hostname.
	//
	// Default: false
	EnableTrustedProxyCheck bool `json:"enable_trusted_proxy_check"`

	// Read EnableTrustedProxyCheck doc.
	//
	// Default: []string
	TrustedProxies     []string `json:"trusted_proxies"`
	trustedProxiesMap  map[string]struct{}
	trustedProxyRanges []*net.IPNet

	// If set to true, c.IP() and c.IPs() will validate IP addresses before returning them.
	// Also, c.IP() will return only the first valid IP rather than just the raw header
	// WARNING: this has a performance cost associated with it.
	//
	// Default: false
	EnableIPValidation bool `json:"enable_ip_validation"`

	// If set to true, will print all routes with their method, path and handler.
	// Default: false
	EnablePrintRoutes bool `json:"enable_print_routes"`

	// You can define custom color scheme. They'll be used for startup message, route list and some middlewares.
	//
	// Optional. Default: DefaultColors
	ColorScheme Colors `json:"color_scheme"`

	// RequestMethods provides customizibility for HTTP methods. You can add/remove methods as you wish.
	//
	// Optional. Default: DefaultMethods
	RequestMethods []string
}
```

For example we can add a custom response header

```go
// Create config variable
config := fiber.Config{
  ServerHeader: "My Server", // add custom server header
}

// Create a new Fiber instance with custom config
app := fiber.New(config)
```

Will add the following header to each response

```go
HTTP/1.1 200 OK
Server: My Server
Date: Fri, 17 Feb 2023 09:43:41 GMT
Content-Type: application/json
Content-Length: 1622
```

### Prefork
Type: bool, default: false.

Enables use of the SO_REUSEPORT socket option. This will spawn multiple Go processes listening on the same port. Also, it's called socket sharding. If enabled, the application will need to be run through a shell because prefork mode sets environment variables.

### ServerHeader
Type: string, default: "" (empty string).

Enables the Server HTTP header with the given value.

### StrictRouting
Type: bool, default: false.

When enabled, the Fiber router treats /foo and /foo/ as different routes. This can be useful, if we want to improve the SEO (Search Engine Optimization) of the website.

### CaseSensitive
Type: bool, default: false.

When enabled, the Fiber router treats /Foo and /foo as different routes.

### Immutable
Type: bool, default: false.

When enabled, all values returned by context methods are immutable. By default, they are valid until you return from the handler.

### UnescapePath
Type: bool, default: false.

Converts all encoded characters in the route back before setting the path for the context, so that the routing can also work with URL encoded special characters.

### ETag
Type: bool, default: false.

Enable ETag header generation, since both weak and strong etags are generated using the same hashing method (CRC-32).

### BodyLimit
Type: int, default: 4 * 1024 * 1024.

Sets the maximum allowed size for a request body. If the size exceeds the configured limit, it sends HTTP 413 Payload Too Large response.

### Concurrency
Type: int, default: 256 * 1024.

Maximum number of concurrent connections.

### Views
Views is the interface to provide your own template engine and contains Load and Render methods.

The Load method is executed by Fiber on app initialization to load/parse the templates. And the Render method is linked to the ctx.Render function that accepts a template name and binding data.

The Fiber team supports template package that provides wrappers for multiple template engines:

### ReadTimeout
Type: time.Duration, default: nil.

The amount of time allowed to read the full request, including the body. Set to nil for unlimited timeout.

### WriteTimeout
Type: time.Duration, default: nil.

The maximum duration before timing out writes of the response. Set to nil for unlimited timeout.

### IdleTimeout
Type: time.Duration, default: nil.

The maximum amount of time to wait for the next request when keep-alive is enabled.

☝️ Note: If IdleTimeout is zero, the value of ReadTimeout is used.

### ReadBufferSize
Type: int, default: 4096.

Per-connection buffer size for requests' reading. This also limits the maximum header size. Increase this buffer, if your clients send multi-KB RequestURIs and/or multi-KB headers.

### WriteBufferSize
Type: int, default: 4096.

Per-connection buffer size for responses' writing.

### CompressedFileSuffix
Type: string, default: ".fiber.gz".

Adds a suffix to the original file name and tries saving the resulting compressed file under the new file name.

### ProxyHeader
Type: string, default: "" (empty string).

This will enable ctx.IP to return the value of the given header key. By default, ctx.IP will return the Remote IP from the TCP connection.

☝️ Note: This property can be useful if you are behind a load balancer, e.g. X-Forwarded-*.

### GETOnly
Type: bool, default: false.

Enables to rejects all non-GET requests. This option is useful as anti-DoS protection for servers accepting only GET requests.

☝️ Note: If GETOnly is set to true, the request size is limited by ReadBufferSize.

### ErrorHandler
Type: ErrorHandler, default: DefaultErrorHandler.

ErrorHandler is executed, when an error is returned from fiber.Handler.

### DisableKeepalive
Type: bool, default: false.

Disable keep-alive connections. The server will close incoming connections after sending the first response to the client.

### DisableDefaultDate
Type: bool, default: false.

When set to true causes the default date header to be excluded from the response.

### DisableDefaultContentType
Type: bool, default: false.

When set to true, causes the default Content-Type header to be excluded from the Response.

### DisableHeaderNormalizing
Type: bool, default: false.

By default, all header names are normalized. For example, header cOnteNT-tYPE will convert to more readable Content-Type.

### DisableStartupMessage
Type: bool, default: false.

When set to true, it will not print out debug information and startup message,

## Route Handlers
Use the following methods for register a route bound to a specific HTTP method.

```go
// Create a new route with GET method
app.Get("/", handler)

// Create a new route with POST method
app.Post("/article/add", handler)

// Create a new route with PUT method
app.Put("/article/:id/edit", handler)

// Create a new route with PATCH method
app.Patch("/article/:id/edit", handler)

// Create a new route with DELETE method
app.Delete("/article/:id", handler)

// ...
```

The Fiber Web Framework has the following helper methods for handling HTTP methods: Add and All.

The Add method allows you to use an HTTP method, as the value:

```go
// Set HTTP method to variable
method := "POST"

// Create a new route with this method
app.Add(method, "/form/create", handler)
```

And the All method opens up possibilities for us to register a route that can accept any valid HTTP methods, which can be handy in some cases:

```go
// Create a new route for all HTTP methods
app.All("/all", handler)
```

## Group

Use Group method for grouping routes by creating the `*Group` struct.

```go
package main

import "github.com/gofiber/fiber/v2"

func main() {
  // Create a new Fiber instance
  app := fiber.New()

  // Create a new route group '/api'
  api := app.Group("/api", handler)

  // Create a new route for API v1
  v1 := api.Group("/v1", handler)
  v1.Get("/list", handler)

  // Create a new route for API v1
  v2 := api.Group("/v2", handler)
  v2.Get("/list", handler)

  // Start server on port 3000
  app.Listen(":3000")
}
```
This built-in method is great for versioning our application's API.

## Stack / Debug

Use Stack method for return the original router stack.

```go
package main

import (
  "encoding/json"
  "fmt"

  "github.com/gofiber/fiber/v2"
)

func main() {
  // Create a new Fiber instance
  app := fiber.New()

  // Create new routes
  app.Get("/api/v1/book:id", handler)
  app.Post("/api/v1/book", handler)

  // Print the router stack in JSON format
  data, _ := json.MarshalIndent(app.Stack(), "", "  ")
  fmt.Println(string(data))

  // Start server on port 3000
  app.Listen(":3000")
}
```

The result will be a list of all routes in a nicely formatted JSON format:

```go
[
  [
    {
      "method": "GET",
      "path": "/api/v1/book:id",
      "params": [
        "age"
      ]
    }
  ],
  [
    {
      "method": "POST",
      "path": "/api/v1/book",
      "params": null
    }
  ]
]
```

## Fiber Context

### BodyParser

Binds the request body to a struct. BodyParser supports decoding query parameters and the following content types based on the Content-Type header:

- application/json
- application/xml
- application/x-www-form-urlencoded
- multipart/form-data

```go
import (
  "fmt"

  "github.com/gofiber/fiber/v2"
)

// Define the Person struct
type Book struct {
  Title  string `json:"title" form:"title"`
  Author string `json:"author" form:"author"`
}

func main() {
  // Create a new Fiber instance
  app := fiber.New()

  // Create a new route with POST method
  app.Post("/api/v1/book", func(c *fiber.Ctx) error {
    // Define a new Person struct
    book := new(Book)

    // Binds the request body to the Person struct
    if err := c.BodyParser(book); err != nil {
      return err
    }

    // Print data from the Person struct
    fmt.Println(book.Title, book.Author)

    return nil
  })

  // Start server on port 3000
  app.Listen(":3000")
}
```

If we run this application and send data from the form to the route http://localhost:3000/api/v1/book by POST, we will see in the console the data we sent.

☝️ Note: Field names in a struct should start with an uppercase letter.

## Resources

- https://dev.to/koddr/go-fiber-by-examples-delving-into-built-in-functions-1p3k

## TODO
- add Update
- add documentation api endpoint
- handle panics
- handle appliationjson headers etc and content type
- APIJson problem for responses
- Swagger
- HAL ?