echo "Should just generate file and exit if command is not given"

export VAR="1234567890"

cat > _out/actual.template << 'END'
VAR = {{ .env.VAR }}
END

cat > _out/expect << END
VAR = ${VAR}
END

set -e

./reconf -w _out/actual
cmp _out/actual _out/expect
