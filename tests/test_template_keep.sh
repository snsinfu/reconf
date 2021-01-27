echo "Should not overwrite file if not forced"

export VAR="1234567890"

cat > _out/actual.template << 'END'
VAR = {{ .env.VAR }}
END

cat > _out/expect << END
VAR = ${VAR}
END

# Initial run
./reconf -w _out/actual

# Second run does nothing becasue the file already exists
VAR="abcdef"
./reconf -w _out/actual

cmp _out/actual _out/expect
