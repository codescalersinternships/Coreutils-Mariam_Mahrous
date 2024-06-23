# Coreutils-Mariam_Mahrous

 This repo covers 9 basic linux commands implemented using Go 

To run each command, you can use `go run` followed by the Go file name and any arguments required by the command. Here’s how you can run each one:

 1. Head :
    Print the first 10 lines of input by default. Use the `-n` flag to specify a different number of lines to print.
    ```sh
    go run head.go [filename] -n [number]
    ```
 2. Tail :
    Print the last 10 lines of input by default. Use the `-n` flag to specify a different number of lines to print.
    ```sh
    go run tail.go [filename] -n [number]
    ```
 3. Echo :
    Print arguments to standard output. use -n to omit the trailing newline.
    ```sh
    go run echo.go -n [text]
    ```
 4. Cat :
    print files. Use -n to output lines numbers.
    ```sh
    go run cat.go -n [filename]
    ```
 5. wc : 
    Count lines, words, and characters in the input.
    Add -l, -w, and -c flags to display only lines, words, or characters respectively.
    ```sh
    go run wc.go [filename] -wlc
    ```
 6. true :
    A command that does nothing successfully.
    ```sh
    go run true.go
    ```
 7. false :
    A command that does nothing unsuccessfully.
    ```sh
    go run false.go
    ```

 8. tree :
    Print all files and folders inside a directory recursively. Use the `-l` flag to specify the maximum depth of printing files.
    ```sh
    go run tree.go -l [number]
    ```
    
 9. env :
  Print all environment variables.
    ```sh
    go run env.go
    ```
 