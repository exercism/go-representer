#!/usr/bin/env sh
set -e

# Synopsis:
# Run the representer on a solution using the representer Docker image.
# The representer Docker image is built automatically.

# Arguments:
# $1: exercise slug
# $2: path to solution folder
# $3: path to output directory

# Output:
# Writes the representation to a representation.txt and representation.json file
# in the passed-in output directory.
# The output files are formatted according to the specifications at https://github.com/exercism/docs/blob/main/building/tooling/representers/interface.md

# Example:
# ./bin/run-in-docker.sh two-fer path/to/solution/folder/ path/to/output/directory/

# If any required arguments is missing, print the usage and exit
if [ "$#" -lt 3 ]; then
    echo "usage: ./bin/run-in-docker.sh exercise-slug path/to/solution/folder/ path/to/output/directory/"
    exit 1
fi

slug="$1"
solution_dir=$(realpath "${2%/}")
output_dir=$(realpath "${3%/}")

# Create the output directory if it doesn't exist
mkdir -p "${output_dir}"

# Build the Docker image
docker build --rm -t exercism/go-representer .

# Run the Docker image using the settings mimicking the production environment
docker run \
    --rm \
    --network none \
    --read-only \
    --mount type=bind,src="${solution_dir}",dst=/solution \
    --mount type=bind,src="${output_dir}",dst=/output \
    --mount type=tmpfs,dst=/tmp \
    exercism/go-representer "${slug}" /solution /output 
