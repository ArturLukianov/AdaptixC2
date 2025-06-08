all: clean
	@ echo "    * Building agent_cog plugin"
	@ mkdir dist
	@ cp config.json ./dist/
	@ go build -buildmode=plugin -ldflags="-s -w" -o ./dist/agent_cog.so pl_main.go pl_agent.go
	@ echo "      done..."

	@ echo "    * Building agent"
	@ cp -r src_cog ./dist/src_cog
	@ echo "      done..."

clean:
	@ rm -rf dist
