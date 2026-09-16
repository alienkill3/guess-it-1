# Guess-it

A Go program for the Zone01 Guess-it project. It reads integers from standard input and, after each number, prints a range predicting the next number.

The prediction uses a moving window, the mean, and the standard deviation. The goal is to balance successful predictions with narrow ranges, since the tester awards more points for a correct prediction within a smaller range.

## Requirements

Go 1.25.0 or later, as declared in go.mod.

Docker with Docker Compose to run the supplied tester.

The Go program uses only the standard library.

Run the program

From the project root:

go run .

Enter one integer per line. The program prints two integers separated by a space:

lower upper

Each output predicts the next input, not the number just entered. The program keeps reading until EOF or an error. In a macOS or Linux terminal, press Ctrl+D on an empty line to send EOF.

## Example

printf '10\n14\n12\n' | go run .

## Output:

9 11
11 13
11 13

## You can also read numbers from a file containing one integer per line:

go run . < numbers.txt

## Prediction method

After receiving a valid integer, the program:

Adds it to the moving window.

Removes the oldest number if the window contains more than 3 numbers.

Calculates the mean, population variance, and standard deviation.

Calculates a margin and rounds the range boundaries outward.

margin = max(standardDeviation × 0.1, 1)
lower  = floor(mean - margin)
upper  = ceil(mean + margin)

Setting

Value

Window size

3 numbers

Standard deviation multiplier

0.1

Minimum margin on each side

1

Before the window fills, the program uses all numbers received so far. Statistical calculations use float64; the final boundaries are integers.

These settings were selected through tester comparisons. They do not guarantee a correct prediction for every input sequence.

## Error handling

Invalid input produces an error on stderr containing the line number and invalid value.

Input-reading errors are also reported on stderr.

Errors terminate the program with exit code 1.

EOF ends the program normally.

Standard output contains only prediction ranges.

For example, entering hello as the second line produces:

Invalid input on line 2: "hello" is not a valid integer.

When running through go run, Go may also display exit status 1.

## Project files

Path

Purpose

main.go

Calls the input handler and sets the failure exit code.

input.go

Reads and validates input, maintains the window, and prints predictions.

stats.go

Calculates the mean, variance, and standard deviation.

prediction.go

Calculates the prediction boundaries.

go.mod

Declares the Go module and version.

PRD.md

Project requirements and design document.

guess-it-dockerized/

Supplied Docker tester.

guess-it-dockerized/student/

Copy of the Go program used by the tester.

## Docker tester

Start Docker Desktop or your Docker daemon. From the project root:

cd guess-it-dockerized
docker compose up --build

Keep that terminal running and open the tester against big-range.

Change the guesser URL parameter to select an opponent:

big-range

average

median

nic (bonus)

Select Test Data 1, 2, or 3 and use Quick to display the result. Use Clean between runs. Each dataset button randomly selects one of five input files, so scores can vary between runs.

Compare Final Result scores to determine the winner. Correct Guesses is the success percentage; a higher percentage alone does not mean a higher score.

The audit requires three runs for each dataset against each opponent, with at least two wins out of three for each dataset separately. The required opponents are big-range, average, and median; nic is the bonus opponent.

Keep the two code copies synchronized

The tester runs the files in guess-it-dockerized/student/. A Docker rebuild does not copy files from the project root into that folder.

After editing the main project, run this from the project root to update the tester copy:

cp go.mod main.go input.go stats.go prediction.go guess-it-dockerized/student/

This replaces the corresponding files in student/ with the root versions. If you experimented inside student/, first copy the changes you want to keep back to the root.

## Then stop the running tester with Ctrl+C and rebuild from its folder:

cd guess-it-dockerized
docker compose up --build

Refresh the browser after rebuilding. The supplied Compose configuration has no bind mount, so local edits are not automatically reflected in the running container.

Test error handling inside Docker

With the tester running, open a second terminal in guess-it-dockerized:

docker compose exec -T node sh -c 'printf "10\nhello\n" | ./student/script.sh'

The program should print 9 11, report invalid input on line 2, and exit with an error. The displayed order of the range and error may vary because they use separate output streams.