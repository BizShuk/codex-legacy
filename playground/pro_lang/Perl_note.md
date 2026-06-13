

[Perl doc](http://irw.ncut.edu.tw/peterju/perl.html)
======




======

- deal with "ngx.print" and "print"

  ans: (?<!ngx\.)print

-


- not start with
	`^(?!debug)(.)` replace with `\t\1\2`



pcre jit
======

- ref. http://sljit.sourceforge.net/pcre.html
- ref. https://redmine.openinfosecfoundation.org/projects/suricata/wiki/Installation_from_GIT_with_PCRE-JIT

./configure --enable-utf8 --enable-unicode-properties --enable-jit