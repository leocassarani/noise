# Noisy

Noisy is a command-line utility that injects a customisable amount of
random noise into a stream of bytes coming from standard input.

Noisy accomplishes this by flipping a user-defined amount (e.g. 25%) of
bits in the input stream, chosen at random. The resulting (noisy) stream
is written to standard output.

## Installation

As long as you have a working installation of Go, you can use `go get`:

    go get github.com/leocassarani/noisy

## Usage

Pipe some text into noisy:

    $ echo -n 'Hello World!' | noisy | xxd
    00000000: cb7d 4c2d 5b99 f55e 5ff8 6435            .}L-[..^_.d5

Use the `-rate` command-line flag to specify a floating-point number
between 0 and 1 representing how many errors you would like. 0 will
leave the input stream untouched, while 1 will flip every bit in your
input. The default value is 0.25.

    $ echo 'Hello World!' | noisy -rate 0.05
    @ello`WkrlD!*
