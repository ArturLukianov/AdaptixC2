all: clean
	@ echo "    * Building listener_cog_stego_http plugin"
	@ mkdir dist
	@ cp config.json ./dist/
	@ go build -buildmode=plugin -ldflags="-s -w" -o ./dist/listener_cog_stego_http.so pl_main.go pl_listener.go pl_stego_http.go map.go
	@ echo "      done..."

clean:
	@ rm -rf dist
