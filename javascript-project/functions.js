//FUNCTIONS

const power = function (base,exponent) {
    let result= 1;
    for (count=0;count<exponent;count++) {
        result *= base;
    }

    return result;
};

console.log(power(2,4))

function future()  {
    return "there will never be flying cars";
}
console.log("the future says",future())

function isitEven(n) {
    if (n%2==0) {
        return "true"
    }else {
        return "false"
    }

}
 console.log(isitEven(50))

 //Bean Counter Excercise

 function countbs(word) {
     counter = 0
     for (let nletter=0;nletter<=word.length;nletter++) {
         if (word[nletter] =="B") {
             counter+=1
         }

     }
     return counter;
 }

 console.log(countbs("BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"))

 function countchar(word,letter) {
    counter = 0
    for (let nletter=0;nletter<=word.length;nletter++) {
        if (word[nletter] == letter) {
            counter+=1
        }

    }
    return counter;
}

console.log(countchar("pneumonoultramicroscopicsilicovolcanoconiosis","n"))

//

