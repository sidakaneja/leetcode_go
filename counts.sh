
easy=$(ls -R | grep Easy -c)
medium=$(ls -R | grep Medium -c)
hard=$(ls -R | grep Hard -c)
total=$((easy + medium + hard)) 

easy_percent=$(echo "scale=2 ; $easy / $total" | bc)
medium_percent=$(echo "scale=2 ; $medium / $total" | bc)
hard_percent=$(echo "scale=2 ; $hard / $total" | bc)


echo "easy $easy $easy_percent"
echo "medium $medium $medium_percent"
echo "hard $hard $hard_percent"
echo "total $total"
