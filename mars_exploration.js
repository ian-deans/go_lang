
let input = "SOSSPSSQSSOR"

/* 
  - split string into lengths of 3.
  - compare each chunk to SOS
  - count the number of differences
  - return count 
*/

const sliceMessage = (inputString, messageArray) => {
  var substring = inputString.slice(0, 3)
  inputString = inputString.slice(3)
  return messageArray.push(substring)
}

const messages = []

sliceMessage(input, messages)

console.log(messages)
console.log(input)



