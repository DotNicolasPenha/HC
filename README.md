# HCA (http calls)
write simple scripts to request APIs
## Abreviations 

- "m" is method
- "h" is header
- "j" is json
- "p" is path

## How to use?

first create a `main.hc` to define the `base url` to request and if you want, the headers.

```
baseurl: http://localhost:3000
h:{}
```

second, create a any file for create the script, follow this example: `hello.hc`.

now write in this format in the file:
```
@<namefile>:<callername>

m: get
p: /path
h: {} 
j: {}
```

here is an example in the file hello.hc:

```
@hello:getHello

m: get
p: /hello
h: {}
j: {}
```

then, run in your terminal this:
```
hca run @hello:getHello
```  
The expected output:
```
Status: 200
Body: {msg:"hello"}
```