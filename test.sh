echo Try to run "./push-swap"
./push-swap
echo Try to run "./push-swap "2 1 3 6 5 8""
./push-swap "2 1 3 6 5 8"
echo Try to run "./push-swap "0 1 2 3 4 5""
./push-swap "0 1 2 3 4 5"
echo Try to run "./push-swap "0 one 2 3""
./push-swap "0 one 2 3"
echo Try to run "./push-swap "1 2 2 3""
./push-swap "1 2 2 3"
echo Try to run "./push-swap "5 random numbers"" with 5 random numbers instead of the tag.
./push-swap "4 1 92 80 5"
echo Try to run "./checker" and input nothing.
./checker
echo Try to run "./checker "0 one 2 3"".
./checker "0 one 2 3"
echo Try to run "echo -e "sa\npb\nrrr\n" | ./checker "0 9 1 8 2 7 3 6 4 5"".
echo -e "sa\npb\nrrr\n" | ./checker "0 9 1 8 2 7 3 6 4 5"
echo Try to run "echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" | ./checker "0 9 1 8 2"".
echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" | ./checker "0 9 1 8 2"
echo Try to run "ARG=("4 67 3 87 23"); ./push-swap $ARG | ./checker $ARG".
ARG=("4 67 3 87 23"); ./push-swap "$ARG" | ./checker "$ARG"
