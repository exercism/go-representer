#!/usr/bin/env sh

# Synopsis:
# Run the representer on a solution.

# Arguments:
# $1: exercise slug
# $2: path to solution folder
# $3: path to output directory

# Output:
# Writes the representation to a representation.txt and representation.json file
# in the passed-in output directory.
# The output files are formatted according to the specifications at https://github.com/exercism/docs/blob/main/building/tooling/representers/interface.md

# Example:
# ./bin/run.sh two-fer path/to/solution/folder/ path/to/output/directory/

# If any required arguments is missing, print the usage and exit
if [ "$#" -lt 3 ]; then
    echo "usage: ./bin/run.sh exercise-slug path/to/solution/folder/ path/to/output/directory/"
    exit 1
fi

/opt/representer/bin/representer "${1}" "${2}" "${3}"
