# bitwise
Dumb commandline expression evaluator thingy I use all the time but probably no one else would want to use, dunno.

```sh
go install github.com/soypat/bitwise/cmd/bw@latest
```


## Usage
**What this does**: Evaluate Go expressions.
#### Bithacking

```
$ bw '(^uint8(2)) + 1'
(^uint8(2)) + 1 = 254
```

#### Math
Go's [math](https://pkg.go.dev/math) library imported by default.
```
$ bw '1/2'
1/2 = 0

$ bw '1./2'
1./2 = 0.5

$ bw 'Sqrt(2)'
Sqrt(2) = 1.4142135623730951

$ bw 'Max(Inf(0), NaN())'
Max(Inf(0), NaN()) = +Inf
```

#### Formatting
Hex, formatting and octals follow Go's [fmt package rules](https://pkg.go.dev/fmt).
```
$ bw -f#0x 255
255 = 0xff

$ bw -fb 14
14 = 1110
```