Read File Line by Line

We may use a handy bufio.Scanner structure to read a file line by line. 
Its constructor, NewScanner(), accepts an open file (remember to close the file after the procedure is complete)
and allows you to read the following lines using the Scan() and Text() methods. 
You can inspect errors found during file reading using the Err() function.

 We first read the file test.txt using the os.Open() utility and store it in the file variable. 
 Then we use the file to create a new scanner. 
 The Scan() function reads the next line of the file, which is accessible via the Text() function. 
 The Err() function will return any errors that occurred during scanning after Scan returns false. 
 Err() will return nil if the error is EOF (End of File).