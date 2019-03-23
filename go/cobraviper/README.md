# Cobra and Viper working together

It seems that its not all that clear how cobra and viper can work together so I have created a small example. First off this is how I expect it to work

1. Running the program with no params uses the default options
2. Specifying a option overrides the default option
3. Using a configuration file with no options overrides the default options (if specified in the config file)
4. Using a configuration file with options overrides the default options and the ones from the config files if specified on the command line
