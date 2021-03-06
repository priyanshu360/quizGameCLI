# quizGameCLI
Timed quiz game via the command line implemented in GO

A program that will read in a quiz provided via a CSV file (more details below) and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.

The CSV file should default to `problems.csv` (example shown below), but the user should be able to customize the filename via a flag.

The CSV file will be in a format like below, where the first column is a question and the second column in the same row is the answer to that question.

```
5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7
```

At the end of the quiz the program output the total number of questions correct and how many questions there were in total. Questions given invalid answers are considered incorrect.

The default time limit is 30 seconds, but also customizable via a flag. Quiz stop as soon as the time limit has exceeded. 

![Screenshot from 2021-06-05 16-39-28](https://user-images.githubusercontent.com/20378236/120889839-af750d80-c61c-11eb-9d22-c42f5267631d.png)
![Screenshot from 2021-06-05 16-39-33](https://user-images.githubusercontent.com/20378236/120889841-b13ed100-c61c-11eb-9320-902c5dcdd7d1.png)
![Screenshot from 2021-06-05 16-39-38](https://user-images.githubusercontent.com/20378236/120889842-b1d76780-c61c-11eb-8348-1fe28ac6083a.png)

