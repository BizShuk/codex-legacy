# Client js

- PageX , PageY
    according page , no on IE
- clientX , clientY
    according browser
- offsetX , offsetY
    only on IE , according box model
- layerX , layerY
    only on FF 
- offsetLeft , offsetTop , offsetWidth , offsetHeight
    box relative to offsetParent(node)
    use recursion to get document offsetLeft and offsetTop
- scrollLeft , scrollTop