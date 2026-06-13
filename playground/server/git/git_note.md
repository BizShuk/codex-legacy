# git_note.md

##### # ref.
- [github readme.md syntax](markdown_syntax.js)
- [git hook deployment](https://www.digitalocean.com/community/tutorials/how-to-use-git-hooks-to-automate-development-and-deployment-tasks)
- [gitlab structure](https://github.com/gitlabhq/gitlabhq/blob/master/doc/development/architecture.md)

## download revision package
http://github.com/<username>/<repository name>/archive/<revision|master|branch_name>.zip


### git cmd

##### # env
- `$GIT_DIR` , 	default  `.`
-  ``



##### git checkout
	-f , force
	--work-tree , 要checkout的目錄
	--git-dir , git 來源目錄
	ex: git --work-tree=/tmp --git-dir=/home/shuk/note.git checkout -f [commit]

	從 note.git checkout一份到/tmp 會造成note.git HEAD跑到commit的位置 =>detached HEAD => modified content will make 類似branch的效果 但沒有開branch 可能會造成麻煩


copy HEAD to another dir

	# need create dir first
	git --work-tree=/dir/.... co HEAD
	
	git ls-tree --name-only -r HEAD
	# and do for loop of it


##### # git clean -f -d
	remove untracked file and directory
	- -d , for dir
	- -f , force

##### # [git log](http://git-scm.com/docs/git-log)
	-p 	, 與上次版本的差異 diff commit commit~1
	--graph , 顯示branch線圖
	--stat , 該commit改變的統計 --shortstat 只有總數
	--name-only , 只顯示有改變的檔案名稱
	-n number , 顯示幾筆commit
	--pretty=format:"...." , 依照格式顯示log
		%H	該更新的SHA1雜湊值
		%h	該更新的簡短SHA1雜湊值
		%T	存放該更新的根目錄的Tree物件的SHA1雜湊值
		%t	存放該更新的根目錄的Tree物件的簡短SHA1雜湊值
		%P	該更新的父更新的SHA1雜湊值
		%p	該更新的父更新的簡短SHA1雜湊值
		%an	作者名字
		%ae	作者電子郵件
		%ad	作者的日期 (格式依據 date 選項而不同)
		%ar	相對於目前時間的作者的日期
		%cn	提交者的名字
		%ce	提交者的電子郵件
		%cd	提交的日期
		%cr	相對於目前時間的提交的日期
		%s	標題

	-relative-date , 相對日期 2 days ago

	條件過濾
	--since , --after
	--until , --before
	--auther , 過濾作者
	--committer , 過濾commiter



	git log

##### # git pull
	 --git-dir=/home/shuk/note/.git 	// pull "note" repository
	 --git-dir=../shuk/note/.git 		// pull "note" repository


##### # git ls-tree
	-r , recurse into **subtree** , not include repo root to current dir
	-l , show object size
	--name-only , show file name name-only
	--full-name , include path from repo root
	--abbrev=n , n is digits  that show how long of rev sha-1

	output:
		? filetype rev_sha1 filename

	ex: `git ls-tree -r --full-name --abbrev=10 $rev`
		100644 blob d4702c8c73  README.md
		100644 blob 190a18037c  d/d/d/d/d/d/in_d
		100644 blob 2847afafe2  index.html
		100644 blob 89f9bb6d46  test.tmp

##### # git ls-file
	same with git ls-tree --name-only , but it can filter file stage , ex : cached , modified , deleted , ignored

	-s , simular git ls-tree

##### # git cat-file
	-t , show file type only
	-s , show file size only
	-p , print file content
	??? --textconv , --batch ,--batch-check

	ex: `git cat-file -p $file_rev > filename`


##### # git rev-parese
	--show-toplevel , show work tree dir
	--is-bare-repository , show  true or false of  repository


##### # git rev-list $rev1 $rev2
show all commits between $rev1 and $rev2


### # git config
	--list  	,	show all configs
	--global



for Git 1.8+
git branch --set-upstream-to=remote_name/remote_branch
git branch -u $remote_name/$remote_branch [$local_branch]

	set branch with remote



for Git 1.7
	git br --setp-upstream $local_br_name $remote_name/$remote_branch



### undoc
git remote -v
git remote set-url origin https://github.com/USERNAME/OTHERREPOSITORY.git
git config --global credential.helper 'cache --timeout=3600'


- git config
http://git-scm.com/docs/git-config
https://git-scm.com/book/zh-tw/v1/Git-%E5%9F%BA%E7%A4%8E-%E6%8F%90%E4%BA%A4%E6%9B%B4%E6%96%B0%E5%88%B0%E5%84%B2%E5%AD%98%E5%BA%AB


- mirror repo
    https://help.github.com/articles/duplicating-a-repository
- git bind username and passwd by ssh key



### git server setup
clone from repo to a bare repo by`git clone --bare --share "repo_url"/.git "repo_name.git" ` **it can be pushed to master branch in default , and must check user authorization**



### Gitlab note
- [install guide](https://github.com/gitlabhq/gitlabhq/blob/master/doc/install/installation.md)
- [easily install guide](https://about.gitlab.com/downloads/)




##### # error
fatal: unable to access 'https://github.com/BizShuk/note.git/': Problem with the SSL CA cert (path? access rights?)
