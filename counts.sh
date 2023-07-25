
easy=$(ls -R | grep Easy -c)
echo "easy $easy"
medium=$(ls -R | grep Medium -c)
echo "medium $medium"
hard=$(ls -R | grep Hard -c)
echo "hard $hard"


echo "total $((easy + medium + hard))"

