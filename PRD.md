## Guess-it — Product Requirements Document

Status: agreed initial approach, before implementation.
Implementation language: Go.

# 1. Goal

The program receives numbers one at a time and, after each number, predicts a range for the next number. It uses only numbers already received.

The goal is to predict the next number successfully as often as possible while keeping the range narrow. A wider range earns a lower score for a correct prediction. Individual predictions are not guaranteed to succeed.

# 2. Input and output

Read one integer per line from standard input (stdin).

Keep running while input continues.

After each valid number, immediately print one line to standard output (stdout).

Each output line contains only two integer bounds separated by one space: lower_bound upper_bound.

Do not print headings, prompts, or diagnostic messages to stdout.

Terminate normally at end of input (EOF). Produce no predictions if no numbers are received.

Output format example, not a specific calculation:

70 90

# 3. Sliding window

Use up to the latest N numbers. Select N through testing.

Add each new number before calculating the prediction. When the window exceeds N, remove the oldest number.

Illustrative example only, with N = 3:

New number

Numbers used

10

10

12

10, 12

11

10, 12, 11

15

12, 11, 15

Produce the first prediction as soon as the first number arrives. Do not wait for the window to fill, and do not retain the first number forever.

# 4. Initial prediction method

Reuse the logic of average(), variance(), and standardDeviation() from math-skills.

For the numbers in the current window:

Calculate the mean.

Calculate the variance, dividing by the number of values as in the existing math-skills implementation.

Calculate the standard deviation as the square root of the variance.

Use the standard deviation as the distance from the mean, subject to a minimum margin M.

margin = max(standard deviation, M)
lower bound = floor(mean - margin)
upper bound = ceil(mean + margin)

M is the minimum distance from the center on each side, not the total range width. Select its value through testing.

Round the lower bound downward and the upper bound upward. Do not round intermediate calculations.

With a single number or identical numbers, the standard deviation is zero. The minimum margin prevents a zero-width range.

This is a baseline method to evaluate, not a proven optimal solution.

# 5. Error handling

For invalid input, write an explanatory message to stderr and terminate with failure. Do not produce a prediction for the invalid line.

The message explains the problem and includes the line number and offending value where available. Example:

Invalid input on line 4: "hello" is not a valid integer.

Report input-reading errors to stderr as well. Exact message wording may be improved later.

# 6. Organization and performance

Keep main() clean and focused on coordination.

Organize input handling and calculations into separate functions and files.

Leave the existing math-skills project unchanged.

Retain only the required window, not the entire input history.

Make each prediction available to the tester before waiting for the next number.

Start with a simple implementation and optimize where testing demonstrates a need.

# 7. Testing and acceptance

Verify that:

Each valid number produces exactly one correctly formatted prediction.

Predictions use only numbers already received.

The window correctly removes the oldest number when full.

The first value and repeated identical values use the minimum margin.

Bounds are rounded outward, and the lower bound does not exceed the upper bound.

Invalid input produces an error on stderr, not stdout.

EOF terminates the program normally.

Evaluate the program against Data 1, Data 2, and Data 3 in the provided tester. Record settings and scores. Change one setting at a time to identify its effect.

No numerical target score has been defined yet, and the tester has not been run as part of this PRD.

# 8. Delivery

Provide a student/ folder at the root of the supplied tester, containing the files required to run the program.

Provide an executable script.sh containing the program startup commands.

Script paths must work when invoked from the tester root directory.

Confirm the exact script location and invocation against the tester instructions during integration.

# 9. Open settings

Window size N.

Minimum per-side margin M.

N = 3 is not a final choice; it was only a teaching example. A graphical interface and other extra features are outside the initial scope.