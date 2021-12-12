var fs = require('fs')
let startTime = Date.now()
try {
    let contents = fs.readFileSync(__dirname + "/../resources/npm.json")
    let data = JSON.parse(contents)
    let results = data.map(ele => ({
        _id: ele,
        name: ele.toLowerCase(),
        collectedTime: Date.now()
    }))
    let resultsString = JSON.stringify(results)
    let endTime = Date.now()
    console.log(`resultsLength: ${results.length}, resultsSize: ${resultsString.length}, timeTaken: ${(endTime - startTime)} `)
} catch (e) {
    console.log('Exception while measuring', e)
}