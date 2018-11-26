global.sorcu = {};
require("./parsed-en.js");

const fs = require("fs");

fs.writeFile("en-jbo.json", JSON.stringify(global.sorcu["en"]), (err) => {
    if (err != null) {
        throw err;
    }

    console.log("see en-jbo.json");
});
