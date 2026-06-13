
function solution(B) {
    if (!(B instanceof Array)) return 0;
    // console.log(B)
    const A = [...B];
    let current, next, maxBreach=1, breachPosition=-1, result=0;

    for(let i=A.length-1; i>0; i--) {
        current = A[i];
        next = A[i-1];

        if (maxBreach > 0 && current >= next) result++;

        if (next > current) {
            result = 1;
            maxBreach--;
            A[i-1] = current;

            if (maxBreach<0) {
                result = 0;
                break;
            }
            breachPosition = i;
        }

    }

    if (maxBreach == 1 || // last element
        (maxBreach == 0 && breachPosition == A.length-1) || // delete breached one
        (breachPosition == B.length-1 && solution(B.slice(0,B.length-1))==B.length-1) // delete current
    ) result++;

    return result;
}

const L = [
    [[2,4,2],2],
    [[1,2,3],3],
    [[2,3,4,2],1],
    [[4,5,2,3,4],0],
    [[3,4,5,4],2],
    [[1],1],
    [[5,1,1],1],
    [[4,2],2],
    [[4,4,4,2],1],
    [[4,4,4,4],4],
    [[1,5,2,3,4],1],
    [[5,4,3,2,1],0],
    [[2,2,2,4,2,2],1]
]



L.forEach((e)=> {
    console.log(e[0],solution(e[0]),"expect:",e[1]);
})

