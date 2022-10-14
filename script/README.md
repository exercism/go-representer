# Batch Execution Script

This folder contains a Go program to run the representer against a set of student solutions.
The program prints some information on how many solutions had the same representation and which ones were unique.

As input, the program expects a extracted (not zipped) directory with 500 student solutions in folders with names from 0 to 499.
You can get this data set via the endpoint `https://exercism.org/api/v2/tracks/{track}/exercises/{exercise-slug}/export_solutions`.
The endpoint is only available for maintainers and only allows to download a couple of data sets per day.

After you downloaded and unzipped the solutions, you can run the analysis by execting the following command in the root folder of this repository.
```bash
go run ./script/batch_exec.go <path-to-solutions-folder>
```

If you make improvements to the script that others could benefit from as well, please create a PR.

Example outout of the program (shortened for readability):
```text
101 solutions have the same representation as solution 2.
47 solutions have the same representation as solution 5.
26 solutions have the same representation as solution 7.
19 solutions have the same representation as solution 15.
14 solutions have the same representation as solution 21.
12 solutions have the same representation as solution 17.
[...]

120 unique solutions: [0 4 8 10 29 30 31 32 34 40 55 ... 348 350 361 362 366 375 376 377 380 381 386 389 393 394 398 419 420 423 427 428 438 440 441 446 449 471 474 475 479 480 484 486 496 497]

Number of solutions that did not compile or failed the tests:  4
```