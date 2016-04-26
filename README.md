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

glide update --all-dependencies

Important: sub dependencies are not installed automatically at the 
 moment. So you have to run "glide update" at least twice a time.

Configuration
-------------

The app is configured by viper and cobra. You can use EnvVars, Config 
files and flags. All flags are bound to the viper configuration.

The most important flag is "env" which defines which configuration
file is loaded. Currently allowed values are "release", "debug" and 
"test". The config filename used for release is "config". All other
environments use "config_{env_name}" as config file. Allowed files 
formats are JSON, TOML, YAML, HCL, or Java properties. 

The files are searched in the following folders and order:
- "/etc/uservault"
- "$HOME/.uservault"
- "."

All EnvVars prefixed with "USERVAULT" are also loaded into the viper
configuration. 

The global order for setting configuration values in viper is (where
default has the lowest priority):
- explicit call to Set
- flag
- env
- config
- key/value store
- default

You can access any configuration set by any of these mechanics by
"viper.get...(name)" for getting the "env" config you can use:
viper.getString("env") 
Viper also supports nested config like
viper.getInt("server.port")



