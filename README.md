# go-mime-multipart-upload-srv

 A simple example of an HTTP mime/multipart upload files in Go 

 Start with

 ```
 go run main.go
 ```

 And now upload multiply files in one CURL request

  ```
 curl -X POST -H 'Authorization: Token 123' \
-F "file=@sample.jpg" \
-F "file=@sample2.jpg" \
http://127.0.0.1:4141/v1/upload
 ```

 or 
   ```
 curl -X POST -H 'Authorization: Token 123' \
-F "file1=@sample.jpg" \
-F "file2=@sample2.jpg" \
http://127.0.0.1:4141/v1/upload
 ```