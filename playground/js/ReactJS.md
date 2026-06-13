# React JS

[javascript-redux-generating-containers-with-connect-from-react-redux-visibletodolist](https://egghead.io/lessons/javascript-redux-generating-containers-with-connect-from-react-redux-visibletodolist)

### note bug
- target is not a DOM element => DOM haven't rendered
- Unknown DOM property class. Did you mean className.  => use className in ReactJS ,not class for css.
- Cannot read property 'setState' of undefined , => this.state = {} on constructor
- performUpdateIfNecessary: Unexpected batch number



1. React elements are immutable.
2. **pure** never change prop , and Props are read-only
3. one-way data flow (also called one-way binding) 



### basic

ReactDom.render

function == Component
class way


### === this.prop ===
input parameters from element attribute or ?
prop.children => sub elements in array

### === this.state ===
component statement
State is reserved only for interactivity. Don't use state at all to build this static version. 

##### this.setState()
https://facebook.github.io/react/docs/state-and-lifecycle.html#state-updates-may-be-asynchronous

use `this.setState((prevState, props) => ({ counter: prevState.counter + props.increment }));` , and this.setState is left merge.


##### life state up
two or more component need this value , you need to take it up to ancestor.



### === HOW TO prevent render this component ===
return null instead of return component

### === list component ===
to create list of component

```javascript
function NumberList(props) {
  const numbers = props.numbers;
  const listItems = numbers.map((number) =>
    <li key={number.toString()}>
      {number}
    </li>
  );
  return (
    <ul>{listItems}</ul>
  );
}
```

##### what is key
help React identify which items have changed , added , and removed. So the best key for each is unique.  You can use index as well, but not recommended. because it'll be slow [in-depth explanation about why keys are necessary ](https://facebook.github.io/react/docs/reconciliation.html#recursing-on-children) 

key don't pass it to this.prop value