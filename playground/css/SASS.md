# Sass , Syntactically Awesome StyleSheets #
[SASS reference](http://sass-lang.com/documentation/file.SASS_REFERENCE.html)

defualt use : SCSS

### Variable ###

define `$<var>: <val>;` 
use `$<var>`

### Nesting ###
```
selector {
	selector {

	}
}
```

### Import ###

`@import <filename|url>`


### Mixins ###
```
@mixin <function>( $a: "default value" , $b : "default value" , ...) {
	
}
```

### extend = inheritance ###
```
selectorA {
	
}

selectorB {
	@extend selectorA
}

```
### Operators ###
`+` , `-` , `*` , `/` , '%'
comment `//` , `/**/`







===================================


NOTE
ref:http://blog.visioncan.com/2011/compass-sass-your-css/



教學手冊列表
ref:http://sam0512.blogspot.tw/2013/10/sass.html 


sass doc on github
https://github.com/hlb/sass-doc-zh/blob/master/SASS_REFERENCE_ZH_TW.md

sass
	sudo apt-get install ruby rubygem
	sudo gem install sass
	sudo gem install compass


% 跟mixin 的差益


susy and media 
https://speakerdeck.com/evenwu/rwd-xiao-shi-jiu-shang-shou


Compass


image-width



　sudo gem install compass

　compass create myproject



@import "compass/css3";



background: inline-image("images/fav_ico.png") no-repeat;
	會轉成data image


Sprite 
	
	@import "icons/*.png"; // only for png , load all imgs and merge to one . background position class name will be 
	.icons-imgname{}


reset模块
	通常，编写自己的样式之前，有必要重置浏览器的默认样式。
	写法是
	　　@import "compass/reset";


compass watch

config.rb
	output_style = :compressed     #css壓縮設定
	sass_options = {:debug_info => true}   #debug
