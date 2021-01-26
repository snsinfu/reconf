echo "Should render multiple templates"

export VAR_A="1234567890"
export VAR_B="abcdefghijklmnopqrstuvwxyz"

cat > _out/a.txt.template << 'END'
VAR_A = {{ .env.VAR_A }}
END

cat > _out/b.txt.template << 'END'
VAR_B = {{ .env.VAR_B }}
END

cat > _out/expect << END
VAR_A = ${VAR_A}
VAR_B = ${VAR_B}
END

set -e

./reconf -w _out/a.txt -w _out/b.txt cat _out/a.txt _out/b.txt > _out/actual
cmp _out/actual _out/expect
