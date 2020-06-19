all:
	rm -f dict.txt
	deno run -A src/convert.ts

commit:
	git commit -am 'commit note'

push:
	git push gitee
	git push

pull:
	git pull gitee master
