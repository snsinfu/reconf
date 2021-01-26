echo "Should render complex template"

export PORTS="22/tcp,53/udp,80/tcp,443/tcp"

cat > _out/actual.template << 'END'
{{ range $port := .env.PORTS | split "," -}}
  {{ $number := $port | before "/" -}}
  {{ $proto := $port | after "/" -}}
-p {{ $proto }} --dport {{ $number }}
{{ end -}}
END

cat > _out/expect << END
-p tcp --dport 22
-p udp --dport 53
-p tcp --dport 80
-p tcp --dport 443
END

set -e

./reconf -w _out/actual true
cmp _out/actual _out/expect
