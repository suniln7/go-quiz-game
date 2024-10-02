#  Quiz Game

A program that will read in a quiz provided via a CSV file (more details below) and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.


The CSV file is default to `problems.csv` , but the user is be able to customize the filename via a flag.

The CSV file will be in a format , where the first column is a question and the second column in the same row is the answer to that question.

At the end of the quiz the program will output the total number of questions correct and how many questions there were in total. Questions given invalid answers are considered incorrect.
