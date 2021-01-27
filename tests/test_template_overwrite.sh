echo "Should overwrite file if forced"

export VAR

cat > _out/actual.template << 'END'
VAR = {{ .env.VAR }}
END

# Initial run
VAR="1234567890"

./reconf -w _out/actual

# Second, forced run with different environment
VAR="abcdef"

./reconf -fw _out/actual

# The second run should have overwritten the file with the new environment
cat > _out/expect << END
VAR = ${VAR}
END

cmp _out/actual _out/expect
