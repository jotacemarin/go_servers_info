# Go server info
Technical test for truora write by Julio Marin

## SQL cockroach database
### For initialize cluster: _necesary to backend run_
    $ cockroach start --insecure --listen-addr=localhost

## Run project
    $ go run main.go

## Third party dependencies
* `github.com/PuerkitoBio/goquery` for manipulate DOM elements
* `github.com/ebuchman/go-shell-pipes` for execute shell commands with pipes
* `github.com/go-chi/chi` for routing restful services
* `github.com/lib/pq` postgres driver use to connect with cockroach db