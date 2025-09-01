
regen:
	@make -C ../tndr regen
	t1 generate

repolist:
	curl -s https://api.github.com/users/Senforsce/repos?per_page=1000 | grep -oP '(?<="git_url": ").*(?="\,)'

diffdir:
	diff <(tree /dir/one) <(tree /dir/two)
