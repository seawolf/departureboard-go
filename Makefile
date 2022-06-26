.PHONY: build clean run entities

bin/departure_board: build
	@ ( cd src && ../bin/departure_board )

build:
	@ ( cd src && go build -o ../bin/departure_board . )
	@ chmod +x bin/departure_board

clean:
	@ [ -d bin ] && rm -r bin || true

doc/models.png: _FORCE
	@ dot -Tpng -o doc/models.png doc/models.dot

entities:
	@ ( cd src && go run -mod=mod entgo.io/ent/cmd/ent generate --target=./entities ./entities_schema )

doc/entities.txt: _FORCE
	@ ( cd src && go run -mod=mod entgo.io/ent/cmd/ent describe ./entities_schema > ../doc/entities.txt )

_FORCE: # used to make a target always trigger, e.g. re-generate docs
