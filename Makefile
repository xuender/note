all:
	rm -rf dist
	mkdir -p dist
	deno run -A scripts/convert.ts

commit:
	git commit -am 'commit note'

push:
	git push gitee
	git push

pull:
	git pull gitee master
