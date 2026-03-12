## - abreviations -

- "m" is method
- "h" is header
- "j" is json
- "p" is path

## - callers -

always if you write this ``"@<endpoint>:<actionOfEndpoint>"`` you define a new 
"function of request", to execute run this:
```
  hc run @example:createNewExample 
``` 
can you see the [hello.hc file](./hello.hc) to example.
remember that callers must have unique names.
this is a caller.

## - main.hc -

when you use "hc run ..." the HC search a main.hc file in your pwd,
the main.hc is nescessary to define "baseurl" of your callers, and 
some headers ("h").