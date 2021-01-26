echo "Should propagate exit status of the command"

expect=15
./reconf sh -c "exit ${expect}"
actual=$?

test ${actual} -eq ${expect} || exit 1
