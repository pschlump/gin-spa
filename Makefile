
all:
	go build

setup_git:
	git init
	git add README.md Makefile
	git commit -m "first commit"
	git branch -M main
	git remote add origin https://github.com/pschlump/gin-spa.git
	git push -u origin main
