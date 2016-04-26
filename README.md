Go Config
---------
For development all go packages have to be located in the go path.
When this repository is not checkout under 

"$GOPATH/src/github.com/supul/uservault/"

you have to create a symbolic link to the repository location:

cd $GOPATH/src/github.com/supul/uservault/
ln -s $REPOSITORY/uservault

Dependency Management
---------------------
Dependencies are managed with Glide. You can init a
new project with "glide init". Dependencies are defined in the 
"glide.yaml". The "glide.lock" files locks the current installed
dependencies for all developers.

glide install - install all dependencies based on "glide.lock"
glide update - update dependencies and "glide.lock"

Important: sub dependencies are not installed automatically at the 
 moment. So you have to run "glide update" at least twice a time.
