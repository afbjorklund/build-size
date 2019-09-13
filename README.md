
Here is a small script to measure the impact on code size by dependencies.

When running `go build -work`, the results are kept and a `WORK` is output.

Exporting this variable and running the script, will show a list of sizes.

It shows the size (in KB) of each static library built for each module.
