echo "Should propagate signal-exit status of the command"

signal=2
expect=$((128 + signal))
./reconf sh -c "kill -${signal} \$\$"
actual=$?

test ${actual} -eq ${expect} || exit 1
