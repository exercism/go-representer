#!/usr/bin/env sh

# Synopsis:
# Test the representer by running it against a predefined set of solutions 
# with an expected output.

# Output:
# Outputs the diff of the expected representation files against the 
# actual representation files generated by the test runner.

# Example:
# ./bin/run-tests.sh

exit_code=0
filenames="representation.txt mapping.json"

# Iterate over all test directories
for test_dir in testdata/*/*; do
    test_dir_name=$(basename "${test_dir}")
    test_dir_path=$(realpath "${test_dir}")
    output_dir_path="/tmp/${test_dir_name}"

    mkdir -p "${output_dir_path}"

    bin/run.sh "${test_dir_name}" "${test_dir_path}" "${output_dir_path}"

    for filename in $filenames; do
        actual_filepath="${output_dir_path}/${filename}"
        expected_filepath="${test_dir_path}/expected_${filename}"
        
        echo "${test_dir_name}: comparing ${filename} to expected_${filename}"
        diff "${actual_filepath}" "${expected_filepath}"

        if [ $? -ne 0 ]; then
            exit_code=1
        fi
    done
done

exit ${exit_code}
