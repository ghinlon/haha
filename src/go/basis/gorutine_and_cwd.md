# gorutine_and_cwd 

# Links

* [Changing directories in go routines - Stack Overflow](https://stackoverflow.com/questions/33001411/changing-directories-in-go-routines)


it's a property of the thread on unix systems, and since goroutines can be run on multiple different threads you can find the CWD changed out from under you if your goroutine gets moved. 
