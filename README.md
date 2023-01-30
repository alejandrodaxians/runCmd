## RUNNING SHELL COMMANDS WITH SCRIPTS
<br>
This repo presents a script, written in Python and GoLang respectively, that allows us to run any number of shell commands just running one file. This specific example runs both docker build and docker pùsh to create and upload an image to the GiHub container registry. 

<br>

The usage is easy, you must pass the three arguments needed for the docker commands to succesfully run: the GitHub account that'll store the package, the project name inside said GitHub account, and the package version.

Say we want to upload the package to a repo named "myproject", inside Elon Musk's account, in its version v1.1.2. Then the command will read as follows:

<br>


### Python:

```python
python runCmd_python.py elonmuskgithub myproject 1.1.2
```

### GoLang:

```golang
go run runCmd_golang.go elonmuskgithub myproject 1.1.2
```

If there's a valid Dockerfile in the folder you are running this program from, the image will be built and pushed.