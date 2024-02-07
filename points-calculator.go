package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

/*
PROBLEM:

When calculating points for a high-school test, students can be graded on a scale from 0-4 with .5 increments. These points
are calculated based on the number of points a student gets in each section. For example, if a section of a test has
2 questions and the student got 4's on both of them, their overall point is 4. However, there are a few exceptions to this
rule which is as follows:

0 points = 0
Any points above 0 but less than the overallPoint of 1 = 1
Any points in-between two consecutive overallPoint will be rounded down except the rule above.

For ease of grading, we will need an output that says if a student got x points in a section that has y questions, their
overallPoint = some number.

Example output:
$ ./points-calculator.exe 3
0:   0
1:   1, 2, 3, 4
1.5: 5
2:   6, 7
2.5: 8
3:   9, 10
3.5: 11
4:   12

As you can see, if a student gets 9 points in a section that has 3 questions, their overallPoint = 3
*/

func main() {
	arguments := os.Args[1:]
	numberOfQuestion, _ := strconv.Atoi(arguments[0])

	// Since each question can have 4 points associated, the total number of points on a given section is 4 * numberOfQuestion
	totalPoints := 4 * numberOfQuestion

	// pointsStructure will hold the map that will store the answers to the calculation
	pointsStructure := map[float64][]int{}

	// We exclude 0 because of the behavior of how numbers less than 1 will go to 1
	possibleOverallPoints := []float64{1, 1.5, 2, 2.5, 3, 3.5, 4}

	// Initialize 0 since we know what that is
	pointsStructure[0] = []int{0}

	// Iterate over total number of points with skipping 0 since it's a special case
	for point := 1; point <= totalPoints; point++ {
		var pointValue float64

		// pointValue will take the current point that it's iterating over and divide it by the number of totalPoints. This
		// will allow us to have a pointValue < 1 and be able to compare the value to the possible range since the range
		// needs to be divided by 4 for the number of points a question can have
		pointValue = float64(point) / float64(totalPoints)

		// We will iterate over all possible overall points in order to determine where a certain point will land in the
		// overall points
		for pos, possibleOverallPoint := range possibleOverallPoints {
			// Divide the possible point by 4 to determine to scale in which we will compare the pointValue against
			scale := possibleOverallPoint / 4
			// This case will capture all the values less than 1 and make it 1
			if pointValue < scale {
				pointsStructure[possibleOverallPoint] = append(pointsStructure[possibleOverallPoint], point)
				break
			} else if pointValue == scale {
				// This case will capture all the values that are exactly equal to that scale and set it to the possiblePoint
				pointsStructure[possibleOverallPoint] = append(pointsStructure[possibleOverallPoint], point)
				break
			} else if pointValue > scale && pointValue < possibleOverallPoints[pos+1]/4 {
				// This case will capture all the values that are in-between the current number and the next number. We need
				// to check this because numbers round down.
				// We can do this even at the end of the possibleOverallPoints without a nil pointer exception because
				// the earlier rule of pointValue == scale will catch the last possible number in totalPoints
				pointsStructure[possibleOverallPoint] = append(pointsStructure[possibleOverallPoint], point)
				break
			}
		}
	}

	// Since golang will iterate over a map without an order, we will force the order by iterating through a slice
	listOfPointsInOrder := []float64{0, 1, 1.5, 2, 2.5, 3, 3.5, 4}
	for _, orderItem := range listOfPointsInOrder {
		// Format output based on if it has a .5 or not.
		if orderItem == math.Trunc(orderItem) {
			fmt.Printf("%+v:   ", orderItem)
		} else {
			fmt.Printf("%+v: ", orderItem)
		}

		// Beautify the result by adding commas in between the numbers in the map
		commaSeparatedScores := ""
		for _, arrOfScores := range pointsStructure[orderItem] {
			commaSeparatedScores = commaSeparatedScores + strconv.Itoa(arrOfScores) + ", "
		}

		// Remove the last , from the string before printing it.
		commaSeparatedScores = commaSeparatedScores[:len(commaSeparatedScores)-2]
		fmt.Println(commaSeparatedScores)
	}
}
