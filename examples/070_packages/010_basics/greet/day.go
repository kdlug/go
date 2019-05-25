// Go community recommends to use plain and simple names for packages. For example, strutils for string utility functions or http for HTTP requests related functions.
// A package names with under_scores, hy-phens or mixedCaps should be avoided.
package greet

// Go exports a variable if a variable name starts with Uppercase.
// All other variables not starting with an uppercase letter is private to the package.
var morning = "Good Morning"
var Morning = "Hey, " + morning
