# Explaining the test steps

I started reading about the language, and after read a lot of documentation, I decided to buy two courses indicated on e-mail (Udemy).

I had to abandon the course to start coding, but I'm really excited with this new tool =)

So, first I putted the code inside a main package to have how to run it by `go run .`

I saw that with a simple implementation I could create a very simple route admin on `main.go`

Each of the routes pointing to the action file.

At the first time I didn't split the code that generates the challenge, but after start writing some teste I decided to create the `performX` functions, to have a way to call it inside the tests files.

I found a problem dealing with big.Int values and operations, on `multiply` endpoint that I decided to left it to deliver the code on time, but I'll continue the challenge to get more knowlegde on this language. 

About the tests, I don't know if the best approach is to send the local file through the service by `multipart/form-data`, or simply create a Slice to teste the features/challenges. But the challenge was really nice hehe

well, this code came out of approximately 10h of study in the language, I am sure that some improvements will be evident for a more experienced programmer.

What I noticed can be improvements:

- make more use of language resources such as Receiver Functions
- test structure could contain less code repetition, and maybe different test cases as well
- use `big.Int` in the `multiply` service
- maybe the types on `sum` and` multiply` could be `int` and` * Int` respectively


# League Backend Challenge

In main.go you will find a basic web server written in GoLang. It accepts a single request _/echo_. Extend the webservice with the ability to perform the following operations

Given an uploaded csv file
```
1,2,3
4,5,6
7,8,9
```

1. Echo (given)
    - Return the matrix as a string in matrix format.
    
    ```
    // Expected output
    1,2,3
    4,5,6
    7,8,9
    ``` 
2. Invert
    - Return the matrix as a string in matrix format where the columns and rows are inverted
    ```
    // Expected output
    1,4,7
    2,5,8
    3,6,9
    ``` 
3. Flatten
    - Return the matrix as a 1 line string, with values separated by commas.
    ```
    // Expected output
    1,2,3,4,5,6,7,8,9
    ``` 
4. Sum
    - Return the sum of the integers in the matrix
    ```
    // Expected output
    45
    ``` 
5. Multiply
    - Return the product of the integers in the matrix
    ```
    // Expected output
    362880
    ``` 

The input file to these functions is a matrix, of any dimension where the number of rows are equal to the number of columns (square). Each value is an integer, and there is no header row. matrix.csv is example valid input.  

Run web server
```
go run .
```

Send request
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
```

## What we're looking for

- The solution runs
- The solution performs all cases correctly
- The code is easy to read
- The code is reasonably documented
- The code is tested
- The code is robust and handles invalid input and provides helpful error messages
