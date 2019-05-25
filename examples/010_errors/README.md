// Gotchas
we can use assigning and checking errors in 1 statement like below:

if u, err := url.Parse(s); err != nil {
	return err
} else {
	v.URL = u
}

But assignment should be done in else {} block, because u is now a local variable.

If we want to access to u outside if body, it will cause an errors

if u, err := url.Parse(s); err != nil {
	return err
} 

v.URL = u
