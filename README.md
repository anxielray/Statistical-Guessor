# _*Statistics and Range Calculator*_

## _*Overview*_

- The progrm has been written in Golang
![alt text](Go-Logo_Yellow(1).png)
- This Go program reads a stream of integer values from standard input, calculates the mean and standard deviation of the values, and determines the range within which approximately 99.7% of the values should fall, based on a normal distribution assumption. The program maintains a sliding window of the most recent values to perform these calculations.

## _*Features*_

    Sliding Window: Only the most recent window number of values are considered in the calculations.
    Mean and Standard Deviation Calculation: Computes the mean and standard deviation of the current window of values.
    Range Calculation: Determines the lower and upper bounds within which approximately 99.7% of the values are expected to lie, using the empirical rule (3 standard deviations from the mean).
    Handles Input Dynamically: Continuously reads and processes input until terminated.

## _*Installation*_

- To use this program, you need to have Go installed on your machine. Follow the instructions below to install Go and run the program:

    - Install Go: Follow the installation instructions on the [Go website](https://golang.org).

    - Download the Code: Clone the repository or download the code file.

    ```sh
    git clone https://learn.zone01kisumu.ke/git/rogwel/guess-it-1.git
    cd guess-it-1
    ```
    - Input this on the command line for cloning, or, follow this [link](https://learn.zone01kisumu.ke/git/rogwel/guess-it-1.git) to download.

    - After download, we will have to open our terminal and build the program;

    ```sh
    go build -o main.go
    ```

- After building the program, we can run the program, either by;

```sh
./main
```

or using the shell script that we have in the  ``student/`` directory.

```sh
./script
```

## _*Usage*_

    - Provide Input: Enter integer values one per line. Press Ctrl+D (EOF) to indicate the end of input.

    - Output: The program outputs the range of values (lower and upper bounds) for the most recent window of numbers. If there is only one number in the window, it uses a simple range calculation instead.

## _*Example*_

```sh
go run . 
Input://This line doesn't in the console...
5
10

Output://This line does not appear on the console...
0 15
```

## _*Code Explanation*_

``calculateStatistics(numbers []float64) (mean, stddev float64)``

- Calculates the mean and standard deviation of a slice of float64 numbers.

    ``mean``: The average of the numbers.
    ``stddev``: The standard deviation, calculated as the square root of the variance.

``guess_it_1(mean, stddev float64) (lowerLimit, upperLimit float64)``

- Calculates the range of values within which approximately 99.7% of the data should fall, using the empirical rule (3 standard deviations from the mean).

    - mean: The average of the numbers.
    - stddev: The standard deviation of the numbers.
    - lowerLimit: Mean minus 3 times the standard deviation.
    - upperLimit: Mean plus 3 times the standard deviation.

``main()``

    - Reads integers from standard input.
    Maintains a sliding window of the most recent window number of values.
    Calculates the mean and standard deviation for the current window.
    Computes the range using guess_it_1 and prints it.

## _*Contributing*_

Contributions are welcome! If you have suggestions or find bugs, please create an issue or submit a pull request.

[Anxiel Math Guru](https://dev.to/anxiel_world_28c50ad32379/revolutionizing-mathematics-lessons-with-ai-the-future-of-home-tutoring-4c47)

## _*License*_

- This project is licensed under the MIT License - see the [LICENSE file](LICENSE) for details.
