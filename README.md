# AddNumbersEvenOdd
A simple Function in Golang that sums an array of numbers, returns the even numbers and the odd numbers from the array. 

##
Run Program Manually 

```
go run addNumbersEvenOdd.go
```
  or 
  
```
go build -o addNumbersEvenOdd
./addNumbersEvenOdd 
```
##
Using Docker

[Install Docker](https://docs.docker.com/get-docker/)  on your machine

 Build your container: 
  ``` 
  docker build -t addnumbers_even_odd ./ 
  ```
  
  Run your container: 
  ``` 
  docker run -p 4000:4000 addnumbers_even_odd 
  ```
  
  You can view your images created with 
  ```
  docker images
  ```
  
  ## Terminal Display using asciinema
  
  [![asciicast](https://asciinema.org/a/530186.svg)](https://asciinema.org/a/530186)
  
  
  
