# javascript

<<<<<<< Updated upstream
- [basic js](basic_js.md)
- [old browser js](browser_note.js)
- [es6 feature](es6.md)
- [ReactJS](ReactJS.md)
- [AngularJS]()
- 
=======
### function arguments passed by 
Arguments are passed by Value
Objects are passed by Reference

### variable assignment
copy by reference

### Memory leak
- [js info](http://javascript.info/tutorial/memory-leaks)
- [MSDN](https://msdn.microsoft.com/en-us/library/ms976398.aspx)
- [js mem leak](http://blogger.gtwang.org/2014/01/javascript-memory-leak-patterns.html)

issue:

will cause leak or not?
```
    function A(){}
        A.prototype.x = function(){
            var y = copy of A
            return y
    }
    var a = new A();
    a = a.x()

```
>>>>>>> Stashed changes


### tools
[check here for document and setup script](https://github.com/BizShuk/env_setup/tree/master/pkg/nodejs)

- **npm**: Node JS package manager
- **Webpack**: 
- **Gulp**: Automation just like Grunt but instead of configurations you can write JavaScript with streams like it's a node application. Much better than Grunt.
- **Express**: Node JS web application framework. Could be used for routing and anything else needed is cover through middle wares. Very popular and beautifully designed so if you want to create a web application project with node, you are probably using it.
- Slush and Yeoman: Project scaffolding systems. You can create starter projects with them. Not good that much, use a Github boilerplate instead.
- **Babel** : transcompiler for ES6 to ES5



