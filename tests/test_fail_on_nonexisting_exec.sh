echo "Should fail if executable is not found"
echo "* This test may print an error message"

if PATH= ./reconf notfound; then
    exit 1
else
    exit 0
fi
