# Observer pattern

Subject holds all observers. If the subject has been updated, it will call all observers to do something.




function Subject(){
    this.list = [];
    this.state = null;
    add (observer o) {
        this.list.add(o)
    }

    setState(state s) {
        this.state = s
        this.notifyAllObservers();
    }

    notifyAllObservers() {
        for(Observer o: this.list) {
            o.update();
        }
    }
}

function Observer() {
    update() {
        System.out.Println("do something");
    }
}
