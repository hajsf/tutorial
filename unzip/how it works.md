1. Open the zip file
To unzip the file, first, open it with the zip.OpenReader(name string) function. As always, when working with files, remember to close it if you no longer need it, using the ReadCloser.Close() method in this case.
```go
// 1. Open the zip file
reader, err := zip.OpenReader(source)
if err != nil {
    return err
}
defer reader.Close()
```

2. Get the absolute destination path
Convert our relative destination path to the absolute representation, which will be needed in the step of Zip Slip vulnerability checking
```go
// 2. Get the absolute destination path
destination, err = filepath.Abs(destination)
if err != nil {
    return err
}
```

3. Iterate over zip files inside the archive and unzip each of them
The actual process of unzipping files in Go using archive/zip is to iterate through the files of the opened ZIP file and unpack each one individually to its final destination.
```go
// 3. Iterate over zip files inside the archive and unzip each of them
for _, f := range reader.File {
    err := unzipFile(f, destination)
    if err != nil {
        return err
    }
}
```

4. Check if file paths are not vulnerable to Zip Slip
The first step of an individual file unzipping function is to check whether the path of this file does not make use of the Zip Slip vulnerability, which was discovered in 2018 and affected thousands of projects. With a specially crafted archive that holds directory traversal filenames, e.g., ../../evil.sh, an attacker can gain access to parts of the file system outside of the target folder in which the unzipped files should reside. The attacker can then overwrite executable files and other sensitive resources, causing significant damage to the victim machine.
```go
func unzipFile(f *zip.File, destination string) error {
    // 4. Check if file paths are not vulnerable to Zip Slip
    filePath := filepath.Join(destination, f.Name)
    if !strings.HasPrefix(filePath, filepath.Clean(destination)+string(os.PathSeparator)) {
        return fmt.Errorf("invalid file path: %s", filePath)
    }
```

To detect this vulnerability, prepare the target file path by combining the destination and the name of the file inside the ZIP archive. It can be done using filepath.Join() function. Then we check if this final file path contains our destination path as a prefix. If not, the file may be trying to access the part of the file system other than the destination and should be rejected.

For example, when we want to unzip our file into the /a/b/ directory:
```go
err := unzipSource("testFolder.zip", "/a/b")
if err != nil {
    log.Fatal(err)
}
```

and in the archive there is a file with a name ../../../../evil.sh, then the output of
```go
filepath.Join("/a/b", "../../../../evil.sh")
```

is:
```go
/evil.sh
```
In this way, the attacker can unzip the evil.sh file in the root directory /, which should not be allowed with our check.

5. Create a directory tree
For each file or directory in the ZIP archive, we need to create a corresponding directory in the destination path, so that the resulting directory tree of the extracted files matches the directory tree inside the ZIP. We use os.MkdirAll() function to do this. For directories, we create the corresponding folder in the destination path, and for files, we create the base directory of the file. Note that we return from the function when the file is a directory as only files need to be unzipped, which we will do in the next steps.
```go
// 5. Create a directory tree
if f.FileInfo().IsDir() {
    if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
        return err
    }
    return nil
}

if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
    return err
}
```

6. Create a destination file for unzipped content
Before uncompressing a ZIP archive file, we need to create a target file where the extracted content could be saved. Since the mode of this target file should match the mode of the file inside the archive, we use os.OpenFile() function, where we can set the mode as an argument.
```go
// 6. Create a destination file for unzipped content
destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
if err != nil {
    return err
}
defer destinationFile.Close()
```

7. Unzip the content of a file and copy it to the destination file
In the last step, we open an individual ZIP file and copy its content to the file created in the previous step. Opening with zip.File.Open() gives access to the uncompressed data of the archive file while copying.
```go
// 7. Unzip the content of a file and copy it to the destination file
zippedFile, err := f.Open()
if err != nil {
    return err
}
defer zippedFile.Close()

if _, err := io.Copy(destinationFile, zippedFile); err != nil {
    return err
}
return nil
```