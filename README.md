experimenting with packages
===========================

- general packaging notes

    - note: this info is reflective of my developing understanding of go packaging -- not necessarily of go packaging   :/

    - useful source: <https://github.com/alco/gostart#faq3>

    - my goal is only to begin to organize many functions in a single main.go file -- into sensible files of conceptually-related functions that are called from the `main.go` file.

        - i note this because much of the info about packages understandably separates this purpose from the purpose of creating shared libraries which are to be 'installed' like other third-party packages.

    - (enforced?) convention: directories and filenames can be lowercase; if a function name is uppercase, it will be available to an external calling function; if lowercase, it will not be.

- `Some_func_a()`, called by `main.go`, shows how a function can be called from the `libs` directory via the import command at the top. It appears that the filename itself in the package is irrelevant; so I assume files themselves are a code-organizing convenience.
    - The 'normal', `go run ./main.go` and `go build ./main.go` work for this, and for `Some_func_b()`.

- `Some_func_b()`, called by `main.go`, is just like the above -- it confirms that (to my understandig):
    - the filename is irrelevant
    - as long as the file's package import statement is the same, and `main.go` imports that package, `main.go` has access to all of the functions in any of the files in the package.

- `Some_func_ac()`, called by `main.go`, shows how a function can be called from an external file.
    - `main.go` does not import anything extra.
    - a run or build require `go run ./*.go` and `go build ./*.go` instead of the 'normal' ways of calling the functions, or the following error will appear:

            # command-line-arguments
            ./main.go:19:17: undefined: Some_func_c


---
