# ----------------
# Command examples
# ----------------
# APPNAME ARGUMENT <COMMAND | SUBCOMMAND> --FLAG
# docker container ls
# ----------------

# App called with no param can have different behavior
#   If init was never called (configuration doesn't exits) --> Prompts the user (yes or no) to execute the 'init' command
#   If configuration does exists --> Launch the viewer
cv

# App called with 'help' flag will print the full help doc
cv --help

# Init command to create default configurations
cv init
